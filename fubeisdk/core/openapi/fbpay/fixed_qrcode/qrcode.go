// @Title  fixed_qrcode
// @Description  聚合码支付 以及聚合码支付相关配置
package fixedqrcode

import (
	"encoding/json"
	"fmt"

	"git.myarena7.com/arena/fubeisdk/core/config"
	"git.myarena7.com/arena/fubeisdk/core/openapi/fbpay/common"
	"git.myarena7.com/arena/fubeisdk/util"
)

type QrcodePay struct {
	config config.Config
}

func NewQrcodePay(cfg config.Config) *QrcodePay {
	return &QrcodePay{config: cfg}
}

// QrcodePay 聚合码支付请求参数
type QrcodePayRequest struct {
	common.CommonData
	Detail      *Detail `json:"detail"` // 订单包含的商品信息，Json格式。当微信支付、云闪付支付或者支付宝支付时可选填此字段。对于使用单品优惠的商户，该字段必须按照规范上传，详见“单品优惠参数说明”
	UserID      string  `json:"user_id"`
	ExpiredTime string  `json:"expired_time"`
}

type Detail struct {
	/*
		1.商户侧一张小票订单可能被分多次支付，订单原价用于记录整张小票的交易金额。
		2.当订单原价与支付金额不相等，则不享受优惠。
		3.该字段主要用于防止同一张小票分多次支付，以享受多次优惠的情况，正常支付订单不必上传此参数。
	*/
	CostPrice   int64       `json:"cost_price"`
	ReceiptID   string      `json:"receipt_id"`   // 商家小票ID
	GoodsDetail GoodsDetail `json:"goods_detail"` // 单品信息 json数数组格式 就是下面的GoodsDetail
}

type GoodsDetail struct {
	Quantity  int32  `json:"quantity"`   // 商品数量
	Price     int32  `json:"price"`      // 商品单价 单位为分
	GoodsID   string `json:"goods_id"`   // 商品编码
	GoodsName string `json:"goods_name"` // 商品名称
}

// QrcodePayResponse 聚合码支付返回参数
type QrcodePayResponse struct {
	QrcodeUrl       string `json:"qrcode_url"`        // 二维码链接
	MerchantOrderSN string `json:"merchant_order_sn"` // 外部系统订单号
}

// QrcodePayConfig 聚合码支付请求参数 聚合码配置查询返回参数
type QrcodePayConfig struct {
	MerchantID       int64  `json:"merchant_id"`                            // 付呗商户号,服务商级别必填
	PayUrl           string `json:"pay_url" validate:"required"`            // 授权域名
	SubAppID         string `json:"sub_appid" validate:"required"`          // 公众号appid
	SubAppSecret     string `json:"sub_appsecret" validate:"required"`      // 公众号秘钥
	AlipayAppID      string `json:"alipay_appid" validate:"required"`       // 支付宝应用appid
	AlipayPrivateKey string `json:"alipay_private_key" validate:"required"` // 支付宝应用私钥
	AlipayPublicKey  string `json:"alipay_public_key" validate:"required"`  // 支付宝公钥
}

// QrcodePayConfigResponse 聚合码配置返回参数
type QrcodePayConfigResponse struct {
	StatusFlag bool   `json:"status_flag"` // 成功状态标志（true:成功；false:失败）
	MerchantID int64  `json:"merchant_id"` // 付呗商户号,服务商级别必填
	Msg        string `json:"msg"`         // 错误信息（成功此字段为空字符串）
}

// QrCodeConfigQuery  聚合码支付配置查询
func (q *QrcodePay) QrCodeConfigQuery(merchantID int64) (payConfig QrcodePayConfig, err error) {
	if merchantID == 0 {
		return payConfig, fmt.Errorf("parameter error")
	}
	data, err := util.LanuchRquest(q.config, "openapi.quota.qrcode.config.get", map[string]int64{
		"merchant_id": merchantID,
	})
	if err != nil {
		return payConfig, err
	}
	if err := json.Unmarshal(data, &payConfig); err != nil {
		return payConfig, err
	}
	return
}

// QrCodePayConfig 聚合支付配置
func (q *QrcodePay) QrCodeConfig(payConfig QrcodePayConfig, callLevel int8) (bool, error) {
	if err := util.Validate(payConfig); err != nil {
		return false, err
	}
	if callLevel == 1 && payConfig.MerchantID == 0 {
		return false, fmt.Errorf("parameter error")
	}

	param := make(map[string]interface{})
	param["pay_url"] = payConfig.PayUrl
	param["sub_appid"] = payConfig.SubAppID
	param["sub_appsecret"] = payConfig.SubAppSecret
	param["alipay_appid"] = payConfig.AlipayAppID
	param["alipay_private_key"] = payConfig.AlipayPrivateKey
	param["alipay_public_key"] = payConfig.AlipayPublicKey

	if callLevel == 1 {
		param["merchant_id"] = payConfig.MerchantID
	}

	data, err := util.LanuchRquest(q.config, "openapi.quota.qrcode.config.handle", param)
	if err != nil {
		return false, err
	}
	var qrCodeConfig QrcodePayConfigResponse
	if err := json.Unmarshal(data, &qrCodeConfig); err != nil {
		return false, err
	}
	return qrCodeConfig.StatusFlag, fmt.Errorf(qrCodeConfig.Msg)
}

// QrCodePay  聚合码支付
func (q *QrcodePay) QrCodePay(qrPay QrcodePayRequest) (qrCodePay QrcodePayResponse, err error) {
	if err := util.Validate(qrPay); err != nil {
		return qrCodePay, err
	}
	if qrPay.CallLevel == 1 && qrPay.MerchantID == 0 {
		return qrCodePay, fmt.Errorf("parameter error")
	}
	param := make(map[string]interface{})
	param["merchant_order_sn"] = qrPay.MerchantOrderSN
	param["total_amount"] = qrPay.TotalAmount
	param["expired_time"] = qrPay.ExpiredTime

	if qrPay.SubAppID != "" {
		param["sub_appid"] = qrPay.SubAppID
	}

	if qrPay.GoodsTag != "" {
		param["goods_tag"] = qrPay.GoodsTag
	}

	if qrPay.StoreID != 0 {
		param["store_id"] = qrPay.StoreID
	}

	if qrPay.UserID != "" {
		param["user_id"] = qrPay.UserID
	}

	if qrPay.DeviceNO != "" {
		param["device_no"] = qrPay.DeviceNO
	}

	if qrPay.Body != "" {
		param["body"] = qrPay.Body
	}

	if qrPay.Attach != "" {
		param["attach"] = qrPay.Attach
	}

	if qrPay.TimeoutExpress != "" {
		param["timeout_express"] = qrPay.TimeoutExpress
	}

	if qrPay.NotifyUrl != "" {
		param["notify_url"] = qrPay.NotifyUrl
	}

	if qrPay.PlatformStoreID != "" {
		param["platform_store_id"] = qrPay.PlatformStoreID
	}

	if qrPay.Detail != nil {
		if jsonData, err := json.Marshal(qrPay.Detail); err != nil {
			return qrCodePay, err
		} else {
			param["detail"] = string(jsonData)
		}
	}

	if qrPay.AlipayExtendParams != nil {
		if jsonData, err := json.Marshal(qrPay.AlipayExtendParams); err != nil {
			return qrCodePay, err
		} else {
			param["alipay_extend_params"] = string(jsonData)
		}
	}

	data, err := util.LanuchRquest(q.config, "fbpay.fixed.qrcode.create", param)
	if err != nil {
		return qrCodePay, err
	}
	if err := json.Unmarshal(data, &qrCodePay); err != nil {
		return qrCodePay, err
	}
	return
}
