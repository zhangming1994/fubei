package merchant

import (
	"encoding/json"
	"fmt"

	"git.myarena7.com/arena/fubeisdk/core/config"
	"git.myarena7.com/arena/fubeisdk/util"
)

type Channel struct {
	config config.Config
}

func NewChannel(cfg config.Config) *Channel {
	return &Channel{config: cfg}
}

type CommonChannelData struct {
	SubMchID string `json:"sub_mch_id"` // 对应的三方商户号 微信就是微信商户号码 支付宝就是
	Msg      string `jso:"msg"`         // 上报失败时的失败原因
}

type WechatChannel struct {
	CommonChannelData
	ChannelID  string `json:"channel_id"`  // 微信商户号所属微信渠道号
	CreateTime string `json:"create_time"` // 创建时间 商户没有上报则为空 格式 2018-11-01 11:22:33
}

type Alipay struct {
	CommonChannelData
	SubMchLevel string `json:"sub_mch_level"` // 支付宝商户号等级
	CreateTime  string `json:"create_time"`   // 创建时间 商户没有上报则为空 格式 2018-11-01 11:22:33
}

type UnionScan struct{ CommonChannelData }
type UnionPayQra struct{ CommonChannelData }
type Cups struct{ CommonChannelData }

// openapi.agent.merchant.channel.submch.query AgentChannelSubQueryResponse 渠道商户号查询
type AgentChannelSubQueryResponse struct {
	Wechat      []WechatChannel `json:"wechat"`       // 微信子商户号，数组
	Alipay      []Alipay        `json:"alipay"`       // 支付宝子商户号，数组
	UnionScan   []UnionScan     `json:"unionscan"`    // 银联QRC商户号，数组
	UnionPayQra []UnionPayQra   `json:"unionpay_qra"` // 银联QRA商户号，数组
	Cups        []Cups          `json:"cups"`         // 银联cups商户号，数组
	MerchantID  string          `json:"merchant_id"`  // 付呗商户号
}

//ChannelSubmchquery 支付渠道查询
func (c *Channel) ChannelSubmchquery(merchantID int64) (channel AgentChannelSubQueryResponse, err error) {
	if merchantID == 0 {
		return channel, fmt.Errorf("parameter error")
	}
	if data, err := util.LanuchRquest(c.config, "openapi.agent.merchant.channel.submch.query", map[string]int64{"merchant_id": merchantID}); err != nil {
		return channel, err
	} else {
		if err = json.Unmarshal(data, &channel); err != nil {
			return channel, err
		}
	}
	return
}
