package util

import (
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"regexp"
	"sort"
	"strconv"
	"strings"

	"git.myarena7.com/arena/fubeisdk/core/config"
	"git.myarena7.com/arena/fubeisdk/util/snowflake"
	"github.com/go-playground/locales/zh"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	zh_trasnslations "github.com/go-playground/validator/v10/translations/zh"
)

// CommonResponse 平常接口请求数据和回调函数请求参数
type CommonResponse struct {
	Data          json.RawMessage `json:"data"`           // 返回的数据
	Success       bool            `json:"success"`        // 是否成功
	ResultCode    int32           `json:"result_code"`    // 返回code码
	ResultMessage string          `json:"result_message"` // 返回描述
}

// CommonRequest 公用请求参数
type CommonRequest struct {
	Format     string `json:"format"`      // 数据格式化方式 默认json
	Method     string `json:"method"`      // 请求方式 默认post
	Nonce      string `json:"nonce"`       // 随机数
	Sign       string `json:"sign"`        // 签名
	SignMethod string `json:"sign_method"` // 签名方式
	VendorSN   string `json:"vendor_sn"`   // vendorsn
	BizContent string `json:"biz_content"` // 不同方法需要提交的参数
}

// Validate 参数验证 i18n
func Validate(body interface{}) error {
	zhTran := zh.New()
	trans, _ := ut.New(zhTran, zhTran).GetTranslator("zh")
	validate := validator.New()
	zh_trasnslations.RegisterDefaultTranslations(validate, trans)
	if err := validate.Struct(body); err != nil {
		return fmt.Errorf("%v", err.(validator.ValidationErrors).Translate(trans))
	}
	return nil
}

// SignMD5
// 对所有API请求参数（包括公共参数和请求参数，但除去sign参数），根据参数名称的ASCII码顺序排序
// 将排序好的参数名和参数值拼装在一起，用&符号连接
// 在拼接好的字符串后面无缝添加商户密钥（app_secret)
// 把拼装好的字符串采用utf-8编码、再用md5算法对字符串进行32位加密后，然后转成大写
func SignMD5(appSecret string, param map[string]interface{}) (sign string) {
	lengthList := len(param)
	list := make([]string, 0, lengthList)
	for key, _ := range param {
		list = append(list, key)
	}

	// sort
	sort.Strings(list)

	// strings.builder
	var builderKey strings.Builder
	for k, val := range list {
		data := fmt.Sprintf("%v=%v&", val, param[val])
		if k == lengthList-1 {
			data = fmt.Sprintf("%v=%v%v", val, param[val], appSecret)
		}
		builderKey.WriteString(data)
	}
	return md5Encode(builderKey.String())
}

func md5Encode(data string) string {
	h := md5.New()
	h.Write([]byte(data))
	return strings.ToUpper(hex.EncodeToString(h.Sum(nil)))
}

func LanuchRquest(con config.Config, method, body interface{}) ([]byte, error) {
	bizContent, err := json.Marshal(body)
	if err != nil {
		return nil, err
	}
	nonce := strconv.FormatInt(snowflake.NextID(), 10)
	param := map[string]interface{}{
		"vendor_sn":   con.VendorSN,
		"biz_content": string(bizContent),
		"method":      method,
		"format":      "json",
		"sign_method": "md5",
		"version":     "2.0",
		"nonce":       nonce,
	}
	sign := SignMD5(con.AppSecret, param)
	param["sign"] = sign
	s := Session{Timeout: 20, Datatype: "json"}
	res := CommonResponse{}
	_, err = s.Post(con.URL, nil, &param, &res, nil)
	if err != nil {
		return nil, err
	}

	if res.ResultCode != 200 && !res.Success {
		if method == "openapi.merchant.income" { // 商户进件有的错误依然会返回merchantid
			return res.Data, fmt.Errorf("%d:%s", res.ResultCode, res.ResultMessage)
		}
		return nil, fmt.Errorf("%d:%s", res.ResultCode, res.ResultMessage)
	}
	return res.Data, nil
}

// GetFileSizeFromBase64Result 根据base64文件编码获取上传的大小 如果大小超过2M报错 图片过大 小于等于2M上传图片
func GetFileSizeFromBase64Result(base64Data string) bool {
	// 去掉可能增加上来的=号
	newBase64Data := strings.ReplaceAll(base64Data, "=", "")
	base64Length := len(newBase64Data)
	fileLength := base64Length * 3 / 4 // 长度单位是字节
	return fileLength <= 1024*1024*2
}

// BankNumberCheck 银行卡号校验
func BankNumberCheck(bankNumber string) bool {
	// 验证卡位数
	if len(bankNumber) < 13 || len(bankNumber) > 30 {
		return false
	}
	// 验证有效性
	return true
}

// IDCardNOCheck 身份证校验 仅仅校验长度是不是 15或者18
func IDCardNOCheck(cardNo string) bool {
	switch len(cardNo) {
	case 15:
		// 15位身份证号码：15位全是数字
		result, _ := regexp.MatchString(`^(\d{15})$`, cardNo)
		return result
	case 18:
		// 18位身份证：前17位为数字，第18位为校验位，可能是数字或X
		result, _ := regexp.MatchString(`^(\d{17})([0-9]|X|x)$`, cardNo)
		return result
	default:
		return false
	}
}

// PhoneCheck 电话号码验证
func PhoneCheck(phone string) bool {
	// 暂时仅仅验证手机号和座机号长度
	if len(phone) >= 10 && len(phone) <= 12 {
		return true
	}
	// TODO 正确性验证
	return true
}
