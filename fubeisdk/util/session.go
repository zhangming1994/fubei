package util

import (
	"bytes"
	"compress/gzip"
	"crypto/tls"
	"encoding/json"
	"errors"
	"io"
	"io/ioutil"
	"net"
	"net/http"
	"net/url"
	"strings"
	"time"
)

type Session struct {
	Timeout         time.Duration // 过期时间
	Client          *http.Client
	Userinfo        *url.Userinfo
	Header          *http.Header
	Params          *Params
	IgnoreEmptyBody bool   // 是否忽略空body
	IgnoreRedirect  bool   // 是否忽略重定向
	Datatype        string // 请求数据传输类型
	Ipaddr          string
}

func (s *Session) Post(url string, p *Params, payload, result, err interface{}) (string, error) {
	r := Request{
		Method:  "POST",
		URL:     url,
		Params:  p,
		Payload: payload,
		Result:  result,
		Error:   err,
	}
	return s.send(&r)
}

func (s *Session) SetHeader(key, value string) {
	if s.Header == nil {
		s.Header = &http.Header{}
	}
	if len(s.Header.Get(key)) == 0 {
		s.Header.Add(key, value)
	} else {
		s.Header.Set(key, value)
	}
}

func (s *Session) send(r *Request) (body string, err error) {
	r.Method = strings.ToUpper(r.Method)

	u, err := url.Parse(r.URL)
	if err != nil {
		return
	}
	//连接的Params
	p := Params{} //默认的
	if s.Params != nil {
		for k, v := range *s.Params {
			p[k] = v
		}
	}
	if r.Params != nil {
		//参数带进来的
		for k, v := range *r.Params {
			p[k] = v
		}
	}
	vals := u.Query()
	for k, v := range p {
		vals.Set(k, v)
	}
	u.RawQuery = vals.Encode()
	//
	// Create a Request object;
	header := http.Header{}
	if s.Header != nil {
		for k, _ := range *s.Header {
			v := s.Header.Get(k)
			header.Set(k, v)
		}
	}
	var req *http.Request
	var str []byte
	if r.Payload != nil {
		if s.Datatype == "json" {
			header.Set("Content-Type", "application/json; charset=utf-8")
		}
		str, err = creactPlayload2(s.Datatype, r.Payload)
		if err != nil {
			return
		}
		buf := bytes.NewBuffer(str)
		req, err = http.NewRequest(r.Method, u.String(), buf)
		if err != nil {
			return
		}
	} else {
		req, err = http.NewRequest(r.Method, u.String(), nil)
		if err != nil {
			return
		}
	}

	//
	// Merge Session and Request options
	//
	var userinfo *url.Userinfo
	if u.User != nil {
		userinfo = u.User
	}
	if s.Userinfo != nil {
		userinfo = s.Userinfo
	}
	// Prefer Request's user credentials
	if r.Userinfo != nil {
		userinfo = r.Userinfo
	}
	if r.Header != nil {
		for k, v := range *r.Header {
			header.Set(k, v[0]) // Is there always guarnateed to be at least one value for a header?
		}
	}
	header.Set("Connection", "close")
	req.Header = header
	if userinfo != nil {
		pwd, _ := userinfo.Password()
		req.SetBasicAuth(userinfo.Username(), pwd)
	}
	r.timestamp = time.Now()
	var client *http.Client
	if s.Client != nil {
		client = s.Client
	} else {
		client = &http.Client{}
	}

	//设置请求超时时间
	client.Transport = &http.Transport{
		Dial: func(netw, addr string) (net.Conn, error) {
			return net.DialTimeout(netw, addr, s.Timeout*time.Second)
		},
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}

	//发送请求
	resp, err := client.Do(req)
	if err != nil {
		return
	}
	defer resp.Body.Close()

	if resp.Header.Get("Content-Encoding") == "gzip" {
		r.body, err = dumpGZIP(resp.Body)
		if err != nil {
			return "", err
		}
	} else {
		r.body, err = ioutil.ReadAll(resp.Body)
		if err != nil {
			return "", err
		}
	}
	//解析body
	if string(r.body) != "" {
		if resp.StatusCode < 300 && r.Result != nil {
			if s.Datatype == "json" {
				err = json.Unmarshal(r.body, r.Result)
			}
		}
		if resp.StatusCode >= 400 && r.Error != nil {
			err = json.Unmarshal(r.body, r.Error) // ignore unmarshall error?
		}
	} else {
		if s.Datatype == "json" {
			err = errors.New("body is empty")
		}
	}
	body = string(r.body)
	return
}

func creactPlayload2(dataType string, payload interface{}) (body []byte, err error) {
	if dataType == "json" {
		body, err = json.Marshal(&payload)
		if err != nil {
			return body, err
		}
	}
	return body, nil
}

func dumpGZIP(r io.Reader) ([]byte, error) {
	reader, err := gzip.NewReader(r)
	if err != nil {
		return []byte{}, err
	}
	defer reader.Close()
	body, err := ioutil.ReadAll(reader)
	if err != nil {
		return []byte{}, err
	}
	return body, nil
}
