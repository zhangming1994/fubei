package payment

import (
	"encoding/json"

	"git.myarena7.com/arena/fubeisdk/core/config"
	"git.myarena7.com/arena/fubeisdk/util"
)

type Payment struct {
	config config.Config
}

func NewPayment(cfg config.Config) *Payment {
	return &Payment{config: cfg}
}

// WxOpenIDQueryRequest 微信OPenID获取请求参数
type WxOpenIDQueryRequest struct {
	MerchantID int64  `json:"merchant_id" validate:"required"` // 付呗商户号
	StoreID    int64  `json:"store_id"    validate:"required"` // 门店id
	AuthCode   string `json:"auth_code"   validate:"required"` // 用户的微信付款码
	SubAppID   string `json:"sub_app_id"  validate:"required"` // 商户自己的公众号id
}

// WxOpenIDQueryResponse 微信OPENid获取返回参数
type WxOpenIDQueryResponse struct {
	MerchantID int64  `json:"merchant_id"` // 付呗商户号
	StoreID    int64  `json:"store_id"`    // 门店id
	AppID      string `json:"app_id"`      // 渠道商公众账号id
	SubAppID   string `json:"sub_app_id"`  // 子商户的公众号id
	OpenID     string `json:"open_id"`     // 用户在商户app_id下的唯一标识
	SubOpenID  string `json:"sub_open_id"` // 用户在子商户app_id下的唯一标识
}

// AlipayUserIDQueryRequest 支付宝USERid获取请求参数
type AlipayUserIDQueryRequest struct {
	AppID        string `json:"app_id"`         // 支付宝分配给开发者的应用ID
	Method       string `json:"method"`         // 接口名称 alipay.user.info.share
	Format       string `json:"format"`         // 仅支持JSON
	Charset      string `json:"charset"`        // 请求使用的编码格式，如utf-8,gbk,gb2312等
	SignType     string `json:"sign_type"`      // 商户生成签名字符串所使用的签名算法类型，目前支持RSA2和RSA，推荐使用RSA2
	Sign         string `json:"sign"`           // 商户请求参数的签名串，详见签名
	Timestamp    string `json:"timestamp"`      // 发送请求的时间，格式"yyyy-MM-dd HH:mm:ss"
	Version      string `json:"version"`        // 调用的接口版本，固定为：1.0
	AuthToken    string `json:"auth_token"`     // 用户授权令牌，同 access_token（用户访问令牌）从这里获取 https://opendocs.alipay.com/mini/02qibk
	AppAuthToken string `json:"app_auth_token"` // 从这里获取 https://opendocs.alipay.com/isv/10467/xldcyq
	BizContent   string `json:"biz_content"`    // 业务参数
}

// AlipayUserIDQueryResponse 支付宝USEID获取返回参数
type AlipayUserIDQueryResponse struct {
	Code     string `json:"code"`      // 网关返回码
	Msg      string `json:"msg"`       // 网关返回码描述
	SubCode  string `json:"sub_code"`  // 业务返回码，参见具体的API接口文档
	SubMsg   string `json:"sub_msg"`   // 业务返回码描述，参见具体的API接口文档
	Sign     string `json:"sign"`      // 签名
	UserID   string `json:"user_id"`   // 支付宝用户的userId。
	Avatar   string `json:"avatar"`    // 用户头像地址。
	Province string `json:"province"`  // 省份名称。
	City     string `json:"city"`      // 市名称
	NickName string `json:"nick_name"` // 用户昵称
	Gender   string `json:"gender"`    // 性别。枚举值如下：F：女性；M：男性。
}

// WxOpenIDPaymentQuery 仅支持微信的授权码查询。适用于在支付前获取用户openid，用于营销活动、会员等。
func (p *Payment) WxOpenIDPaymentQuery(wx WxOpenIDQueryRequest) (openID WxOpenIDQueryResponse, err error) {
	if err := util.Validate(wx); err != nil {
		return openID, err
	}
	data, err := util.LanuchRquest(p.config, "openapi.payment.order.query.wxopenid", wx)
	if err != nil {
		return openID, err
	}
	if err := json.Unmarshal(data, &openID); err != nil {
		return openID, err
	}
	return
}

// alipayUserIDQuery  支付宝userid获取  alipay.user.info.share
func (p *Payment) AlipayUserIDQuery(alipay AlipayUserIDQueryRequest) (userID AlipayUserIDQueryResponse, err error) {
	return
}
