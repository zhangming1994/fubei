// @Title  order
// @Description  付款码支付 统一下单 订单查询 关闭 退款等相关接口
package order

import (
	"encoding/json"
	"fmt"

	"git.myarena7.com/arena/fubeisdk/core/config"
	"git.myarena7.com/arena/fubeisdk/core/openapi/fbpay/common"
	"git.myarena7.com/arena/fubeisdk/util"
)

type Order struct {
	config config.Config
}

func NewOrder(cfg config.Config) *Order {
	return &Order{config: cfg}
}

// OrderPayRequest  付款码支付请求参数
type OrderPayRequest struct {
	AlipayExtendParams *common.AlipayExtendParams `json:"alipay_extend_params"`                  // 支付宝业务拓展参数--花呗分期
	CallLevel          int8                       `json:"call_level"`                            // 调用级别 1. 服务商 2. 商户
	StoreID            int64                      `json:"store_id"`                              // 商户门店号（如果只有一家有效门店，可不传）
	CashierID          int64                      `json:"cashier_id"`                            // 收银员ID
	MerchantID         int64                      `json:"merchant_id"`                           // 付呗商户号。以服务商级接入时必传，以商户级接入时不传
	TotalAmount        float64                    `json:"total_amount" validate:"required"`      // 订单总金额，单位为元，精确到0.01 ~ 10000000
	MerchantOrderSN    string                     `json:"merchant_order_sn" validate:"required"` // 外部系统订单号（确保唯一，前后不允许带空格）
	SubAppID           string                     `json:"sub_appid"`                             // 公众号appid
	GoodsTag           string                     `json:"goods_tag"`                             // 订单优惠标记，代金券或立减优惠功能的参数（使用单品券时必传）
	Body               string                     `json:"body"`                                  // 商品描述
	Attach             string                     `json:"attach"`                                // 附加数据，原样返回，该字段主要用于商户携带订单的自定义数据
	TimeoutExpress     string                     `json:"timeout_express"`                       // 订单失效时间，逾期将关闭交易。格式为yyyyMMddHHmmss，失效时间需大于1分钟银联订单请勿传，目前银联不支持该字段
	NotifyUrl          string                     `json:"notify_url"`                            // 支付回调地址
	DeviceNO           string                     `json:"device_no"`                             // 终端号
	PlatformStoreID    string                     `json:"platform_store_id"`                     // 平台方门店号（即微信/支付宝的storeid
	DisablePayChannels string                     `json:"disable_pay_channels"`                  // 禁止使用优惠券标识 promotion-支付宝优惠（包含实时优惠+商户优惠）voucher-支付宝营销券可以单个传，也可以多个传，以逗号分隔如：promotion,voucher
	AuthCode           string                     `json:"auth_code" validate:"required"`         // 支付授权码，用户的付款码。根据支付授权码自动识别支付通道，目前支持微信、支付宝、银联
	Detail             *Detail                    `json:"detail"`                                // 订单包含的商品信息，Json格式。当微信支付、云闪付支付或者支付宝支付时可选填此字段。对于使用单品优惠的商户，该字段必须按照规范上传，详见“单品优惠参数说明”
}

// OrderQueryResponse 订单查询返回参数
type OrderQueryResponse struct {
	OrderCommonData
	AlipayExtendParams json.RawMessage `json:"alipay_extend_params"` // 支付宝业务拓展参数--花呗分期
	IsCanPartRefund    int8            `json:"is_can_part_refund"`   // 是否支持部分退款1 支持 0 不支持
}

// OrderPayResponse 付款码支付返回参数
type OrderPayResponse struct {
	OrderCommonData
	MerchantID         int64           `json:"merchant_id"`          // 付呗商户号
	AlipayExtendParams json.RawMessage `json:"alipay_extend_params"` // 支付宝业务拓展参数--花呗分期
}

// OrderCreateRequest 统一下单请求参数
type OrderCreateRequest struct {
	common.CommonData
	Detail  *Detail `json:"detail"`                        // 订单包含的商品信息，Json格式。当微信支付、云闪付支付或者支付宝支付时可选填此字段。对于使用单品优惠的商户，该字段必须按照规范上传，详见“单品优惠参数说明”
	PayType string  `json:"pay_taype" validate:"required"` // 支付方式，wxpay微信，alipay支付宝
	UserID  string  `json:"user_id" validate:"required"`   // 用户标识（微信openid，支付宝userid）
}

// OrderCreateResponse 统一下单返回参数
type OrderCreateResponse struct {
	StoreID         int64       `json:"store_id"`          // 商户门店号
	CashierID       int64       `json:"cashier_id"`        // 收银员ID
	MerchantID      int64       `json:"merchant_id"`       // 付呗商户号
	TotalAmount     float64     `json:"total_amount"`      // 订单金额，精确到0.01
	OrderSN         string      `json:"order_sn"`          // 付呗订单号
	MerchantOrderSN string      `json:"merchant_order_sn"` // 外部系统订单号
	PrePayID        string      `json:"prepay_id"`         // 预支付凭证，微信预支付订单号prepay_id、支付宝交易号tradeNO等
	PayType         string      `json:"pay_type"`          // 支付方式，wxpay 微信，alipay 支付宝
	UserID          string      `json:"user_id"`           // 付款用户id，“微信openid”、“支付宝账户”等
	DeviceNO        string      `json:"device_no"`         // 终端号
	Attach          string      `json:"attach"`            // 附加数据，原样返回，该字段主要用于商户携带订单的自定义数据
	SignPackage     SignPackage `json:"sign_package"`      // 签名包，当pay_type为wxpay时才返回该字段
}

// OrderQueryReqeust 订单查询请求参数
type OrderQueryReqeust struct {
	CallLevel       int8   `json:"call_level"`        // 调用服务的级别 1. 服务商 2. 商户
	MerchantID      int64  `json:"merchant_id"`       // 付呗商户号，以服务商级接入时必传，以商户级接入时不传
	OrderSN         string `json:"order_sn"`          // 付呗订单号，和外部系统订单号、机构订单号不能同时为空（三选一），如果同时存在三者优先级为：order_sn>ins_order_sn>merchant_order_sn
	MerchantOrderSN string `json:"merchant_order_sn"` // 外部系统订单号，和付呗订单号、机构订单号不能同时为空（三选一
	InsOrderSN      string `json:"ins_order_sn"`      // 机构订单号，和付呗订单号、外部系统订单号不能同时为空（三选一）
}

// OrderRefundRequest 退款请求参数
type OrderRefundRequest struct {
	CallLevel        int8    `json:"call_level"`                            // 1. 服务商级别 2. 商户级别
	Handler          int64   `json:"handler"`                               // 退款操作人员ID
	MerchantID       int64   `json:"merchant_id"`                           // 付呗商户号，以服务商级接入时必传，以商户级接入时不传
	RefundAmount     float64 `json:"refund_amount" validate:"required"`     // 退款金额，精确到0.01，注：使用优惠券时仅支持全额退款
	OrderSN          string  `json:"order_sn"`                              // 付呗订单号，和外部系统订单号、机构订单号不能同时为空（三选一），如果同时存在三者优先级为：order_sn>ins_order_sn>merchant_order_sn
	MerchantOrderSN  string  `json:"merchant_order_sn" validate:"required"` // 外部系统订单号，和付呗订单号、机构订单号不能同时为空（三选一）
	InsOrderSN       string  `json:"ins_order_sn"`                          // 机构订单号，和付呗订单号、外部系统订单号不能同时为空（三选一
	MerchantRefundSN string  `json:"merchant_refund_sn"`                    // 外部系统退款号
	DeviceNO         string  `json:"device_no"`                             // 硬件设备号
}

// OrderRefundResponse 退款返回参数
type OrderRefundResponse struct {
	OrderRefundCommonData
}

// OrderRefundQueryRequest 退款查询请求参数
type OrderRefundQueryRequest struct {
	CallLevel        int8   `json:"call_level"`         // 调用级别 1. 服务商 2. 商户
	MerchantID       int64  `json:"merchant_id"`        // 付呗商户号，以服务商级接入时必传，以商户级接入时不传
	RefundSN         string `json:"refund_sn"`          // 付呗退款号，与外部系统退款号不能同时为空（二选一），如果同时存在优先取refund_sn
	MerchantRefundSN string `json:"merchant_refund_sn"` // 外部系统退款号，与付呗退款号不能同时为空（二选一）
}

// OrderRefundQueryResponse 退款查询返回参数
type OrderRefundQueryResponse struct {
	OrderRefundCommonData
	BuyerRefundAmount float64 `json:"buyer_refund_amount"` // 买家退款金额，精确到0.01（直连订单不返回）
	FinishTime        string  `json:"finish_time"`         // 退款完成时间，格式为yyyyMMddHHmmss，当退款成功或退款失败时返回
}

// OrderCloseRequest 订单关闭请求参数
type OrderCloseRequest struct {
	CallLevel       int8   `json:"call_level"`        // 调用级别 1. 服务商 2 商户
	MerchantID      int64  `json:"merchant_id"`       // 付呗商户号，以服务商级接入时必传，以商户级接入时不传
	OrderSN         string `json:"order_sn"`          // 付呗订单号，和外部系统订单号不能同时为空（二选一），如果同时存在优先取order_sn
	MerchantOrderSN string `json:"merchant_order_sn"` // 外部系统订单号，和付呗订单号不能同时为空（二选一）
}

// OrderCloseResponse 订单关闭返回参数
type OrderCloseResponse struct {
	OrderSN         string `json:"order_sn"`          // 付呗订单号
	MerchantOrderSN string `json:"merchant_order_sn"` // 外部系统订单号
	OrderStatus     string `json:"order_status"`      // 订单状态：	CLOSED--已关闭	关闭失败时会返回错误码信息
	DeviceNO        string `json:"device_no"`         // 硬件设备号
}

type Detail struct {
	GoodsDetail GoodsDetail `json:"goods_detail"` // 单品信息 json数数组格式 就是下面的GoodsDetail
	/*
		1.商户侧一张小票订单可能被分多次支付，订单原价用于记录整张小票的交易金额。
		2.当订单原价与支付金额不相等，则不享受优惠。
		3.该字段主要用于防止同一张小票分多次支付，以享受多次优惠的情况，正常支付订单不必上传此参数。
	*/
	CostPrice int64  `json:"cost_price"`
	ReceiptID string `json:"receipt_id"` // 商家小票ID

}

type GoodsDetail struct {
	Quantity      int32  `json:"quantity"`       // 商品数量
	Price         int32  `json:"price"`          // 商品单价 单位为分
	GoodsID       string `json:"goods_id"`       // 商品编码
	GoodsName     string `json:"goods_name"`     // 商品名称
	GoodsCategory string `json:"goods_category"` // 类目，银联云闪付订单按需上传
	GoodsDesc     string `json:"goods_desc"`     // 附加信息，银联云闪付订单按需上传
}

type OrderCommonData struct {
	PaymentList     []PaymentList `json:"payment_list"`      // 活动优惠列表，Json格式。payment_list的值为数组，每一个元素包含两个字段type和amount（元）
	StoreID         int64         `json:"store_id"`          // 商户门店号
	CashierID       int64         `json:"cashier_id"`        // 收银员ID
	TotalAmount     float64       `json:"total_amount"`      // 订单金额，精确到0.01
	NetAmount       float64       `json:"net_amount"`        // 实收金额，精确到0.01，当支付成功时返回
	BuyerPayAmount  float64       `json:"buyer_pay_amount"`  // 买家实际支付金额，精确到0.01，当支付成功时返回
	Fee             float64       `json:"fee"`               // 手续费（直连订单当日不返回手续费），精确到0.01，当支付成功时返回
	OrderSN         string        `json:"order_sn"`          // 付呗订单号
	MerchantOrderSN string        `json:"merchant_order_sn"` // 外部系统订单号
	InsOrderSN      string        `json:"ins_order_sn"`      // 机构订单号（显示在微信/支付宝支付凭证的订单号）
	ChannelOrderSN  string        `json:"channel_order_sn"`  // 通道订单号，微信订单号、支付宝订单号等，当支付成功时返回
	OrderStatus     string        `json:"order_status"`      // 订单状态：USERPAYING--用户支付中SUCCESS--支付成功REVOKED--已撤销CLOSED--已关闭REVOKING--撤销中
	PayType         string        `json:"pay_type"`          // 支付方式，wxpay微信，alipay支付宝，unionpay银联
	UserID          string        `json:"user_id"`           // 付款用户id，“微信openid”、“支付宝账户”等，当支付成功时返回
	FinishTime      string        `json:"finish_time"`       // 支付完成时间，格式为yyyyMMddHHmmss，当支付成功时返回
	DeviceNO        string        `json:"device_no"`         // 终端号
	Attach          string        `json:"attach"`            // 附加数据，原样返回，该字段主要用于商户携带订单的自定义数据
}

type PaymentList struct {
	Type   string  `json:"type"`   // FUBEI_DISCOUNT 付呗折扣 CHANNEL_DISCOUNT 支付通道免充值 CHANNEL_PRE 支付通道预充值
	Amount float64 `json:"amount"` // 金额
}

type SignPackage struct {
	AppID     string `json:"appId"`     // 公众号id
	TimeStamp string `json:"timeStamp"` // 时间戳，示例：1414561699，标准北京时间，时区为东八区，自1970年1月1日 0点0分0秒以来的秒数。
	NonceStr  string `json:"nonceStr"`  // 随机字符串
	Package   string `json:"package"`   // 统一下单接口返回的prepay_id参数值，提交格式如：prepay_id=123456
	SignType  string `json:"signType"`  // 签名类型，默认为RSA
	PaySign   string `json:"paySign"`   // 签名
}

type OrderRefundCommonData struct {
	Handler          int64   `json:"handler"`            // 退款操作人员ID
	RefundAmount     float64 `json:"refund_amount"`      // 退款金额，精确到0.01，注：使用优惠券时仅支持全额退款
	DeviceNO         string  `json:"device_no"`          // 硬件设备号
	RefundStatus     string  `json:"refund_status"`      // 退款状态：	REFUND_PROCESSING--退款中REFUND_SUCCESS--退款成功REFUND_FAIL--退款失败
	OrderSN          string  `json:"order_sn"`           // 付呗订单号
	RefundSN         string  `json:"refund_sn"`          // 付呗订单号
	MerchantOrderSN  string  `json:"merchant_order_sn"`  // 外部系统订单号
	MerchantRefundSN string  `json:"merchant_refund_sn"` // 外部系统退款号
}

// OrderCreate 统一下单
func (o *Order) OrderCreate(orderCreate OrderCreateRequest) (orderPrepayInfo OrderCreateResponse, err error) {
	if err := util.Validate(orderCreate); err != nil {
		return orderPrepayInfo, err
	}
	if orderCreate.CallLevel == 1 && orderCreate.MerchantID == 0 {
		return orderPrepayInfo, fmt.Errorf("parameter error")
	}
	if orderCreate.PayType == "wxpay" && orderCreate.SubAppID == "" {
		return orderPrepayInfo, fmt.Errorf("parameter error")
	}
	if orderCreate.PayType != "wxpay" && orderCreate.PayType != "alipay" {
		return orderPrepayInfo, fmt.Errorf("parameter error")
	}
	param := make(map[string]interface{}, 0)
	param["merchant_order_sn"] = orderCreate.MerchantOrderSN
	param["pay_type"] = orderCreate.PayType
	param["total_amount"] = orderCreate.TotalAmount
	param["store_id"] = orderCreate.StoreID
	param["user_id"] = orderCreate.UserID

	if orderCreate.MerchantID != 0 {
		param["merchant_id"] = orderCreate.MerchantID
	}

	if orderCreate.CashierID != 0 {
		param["cashier_id"] = orderCreate.CashierID
	}

	if orderCreate.SubAppID != "" {
		param["sub_appid"] = orderCreate.SubAppID
	}

	if orderCreate.GoodsTag != "" {
		param["goods_tag"] = orderCreate.GoodsTag
	}

	if orderCreate.DeviceNO != "" {
		param["device_no"] = orderCreate.DeviceNO
	}

	if orderCreate.Body != "" {
		param["body"] = orderCreate.Body
	}

	if orderCreate.Attach != "" {
		param["attach"] = orderCreate.Attach
	}

	if orderCreate.TimeoutExpress != "" {
		param["timeout_express"] = orderCreate.TimeoutExpress
	}

	if orderCreate.NotifyUrl != "" {
		param["notify_url"] = orderCreate.NotifyUrl
	}

	if orderCreate.PlatformStoreID != "" {
		param["platform_store_id"] = orderCreate.PlatformStoreID
	}

	if orderCreate.DisablePayChannels != "" {
		param["disable_pay_channels"] = orderCreate.DisablePayChannels
	}

	if orderCreate.Detail != nil {
		if jsonData, err := json.Marshal(orderCreate.Detail); err != nil {
			return orderPrepayInfo, err
		} else {
			param["detail"] = string(jsonData)
		}
	}

	if orderCreate.AlipayExtendParams != nil {
		if jsonData, err := json.Marshal(orderCreate.AlipayExtendParams); err != nil {
			return orderPrepayInfo, err
		} else {
			param["alipay_extend_params"] = string(jsonData)
		}
	}
	data, err := util.LanuchRquest(o.config, "fbpay.order.create", param)
	if err != nil {
		return orderPrepayInfo, err
	}
	if err := json.Unmarshal(data, &orderPrepayInfo); err != nil {
		return orderPrepayInfo, err
	}
	return
}

// OrderPay  付款码支付
func (o *Order) OrderPay(orderPay OrderPayRequest) (orderInfo OrderPayResponse, err error) {
	if err := util.Validate(orderPay); err != nil {
		return orderInfo, err
	}
	if orderPay.CallLevel == 1 && orderPay.MerchantID == 0 {
		return orderInfo, fmt.Errorf("parameter error")
	}
	param := make(map[string]interface{}, 0)
	param["merchant_order_sn"] = orderPay.MerchantOrderSN
	param["auth_code"] = orderPay.AuthCode
	param["total_amount"] = orderPay.TotalAmount
	param["store_id"] = orderPay.StoreID

	if orderPay.MerchantID != 0 {
		param["merchant_id"] = orderPay.MerchantID
	}

	if orderPay.CashierID != 0 {
		param["cashier_id"] = orderPay.CashierID
	}

	if orderPay.SubAppID != "" {
		param["sub_appid"] = orderPay.SubAppID
	}

	if orderPay.GoodsTag != "" {
		param["goods_tag"] = orderPay.GoodsTag
	}

	if orderPay.DeviceNO != "" {
		param["device_no"] = orderPay.DeviceNO
	}

	if orderPay.Body != "" {
		param["body"] = orderPay.Body
	}

	if orderPay.Attach != "" {
		param["attach"] = orderPay.Attach
	}

	if orderPay.TimeoutExpress != "" {
		param["timeout_express"] = orderPay.TimeoutExpress
	}

	if orderPay.NotifyUrl != "" {
		param["notify_url"] = orderPay.NotifyUrl
	}

	if orderPay.PlatformStoreID != "" {
		param["platform_store_id"] = orderPay.PlatformStoreID
	}

	if orderPay.DisablePayChannels != "" {
		param["disable_pay_channels"] = orderPay.DisablePayChannels
	}

	if orderPay.Detail != nil {
		if jsonData, err := json.Marshal(orderPay.Detail); err != nil {
			return orderInfo, err
		} else {
			param["detail"] = string(jsonData)
		}
	}

	if orderPay.AlipayExtendParams != nil {
		if jsonData, err := json.Marshal(orderPay.AlipayExtendParams); err != nil {
			return orderInfo, err
		} else {
			param["alipay_extend_params"] = string(jsonData)
		}
	}

	data, err := util.LanuchRquest(o.config, "fbpay.order.pay", param)
	if err != nil {
		return orderInfo, err
	}
	if err := json.Unmarshal(data, &orderInfo); err != nil {
		return orderInfo, err
	}
	return
}

// OrderClose 订单关闭  支持除付款码支付以外的订单进行关闭 只能关闭未支付成功的订单，关闭后状态为“已关闭”
func (o *Order) OrderClose(orderClose OrderCloseRequest) (closeOrder OrderCloseResponse, err error) {
	if err := util.Validate(orderClose); err != nil {
		return closeOrder, err
	}
	if orderClose.CallLevel == 1 && orderClose.MerchantID == 0 {
		return closeOrder, fmt.Errorf("parameter error")
	}

	if orderClose.OrderSN == "" && orderClose.MerchantOrderSN == "" {
		return closeOrder, fmt.Errorf("parameter error")
	}

	param := make(map[string]interface{})
	if orderClose.CallLevel == 1 {
		param["merchant_id"] = orderClose.MerchantID
	}

	if orderClose.OrderSN != "" {
		param["order_sn"] = orderClose.OrderSN
	}

	if orderClose.MerchantOrderSN != "" {
		param["merchant_order_sn"] = orderClose.MerchantOrderSN
	}

	data, err := util.LanuchRquest(o.config, "fbpay.order.close", param)
	if err != nil {
		return closeOrder, err
	}
	if err := json.Unmarshal(data, &closeOrder); err != nil {
		return closeOrder, err
	}
	return
}

// OrderQuery 订单查询
func (o *Order) OrderQuery(queryOrder OrderQueryReqeust) (orderInfo OrderQueryResponse, err error) {
	if err := util.Validate(queryOrder); err != nil {
		return orderInfo, err
	}
	if queryOrder.CallLevel == 1 && queryOrder.MerchantID == 0 {
		return orderInfo, fmt.Errorf("parameter error")
	}

	if queryOrder.OrderSN == "" && queryOrder.MerchantOrderSN == "" && queryOrder.InsOrderSN == "" {
		return orderInfo, fmt.Errorf("parameter error")
	}

	param := make(map[string]interface{})

	if queryOrder.CallLevel == 1 {
		param["merchant_id"] = queryOrder.MerchantID
	}

	if queryOrder.OrderSN != "" {
		param["order_sn"] = queryOrder.OrderSN
	}

	if queryOrder.MerchantOrderSN != "" {
		param["merchant_order_sn"] = queryOrder.MerchantOrderSN
	}

	if queryOrder.InsOrderSN != "" {
		param["ins_order_sn"] = queryOrder.InsOrderSN
	}

	data, err := util.LanuchRquest(o.config, "fbpay.order.query", param)
	if err != nil {
		return orderInfo, err
	}
	if err := json.Unmarshal(data, &orderInfo); err != nil {
		return orderInfo, err
	}
	return
}

// Refund 退款
func (o *Order) Refund(refund OrderRefundRequest) (orderRefund OrderRefundResponse, err error) {
	if err := util.Validate(refund); err != nil {
		return orderRefund, err
	}
	if refund.CallLevel == 1 && refund.MerchantID == 0 {
		return orderRefund, fmt.Errorf("parameter error")
	}

	if refund.OrderSN == "" && refund.MerchantOrderSN == "" && refund.InsOrderSN == "" {
		return orderRefund, fmt.Errorf("parameter error")
	}

	param := make(map[string]interface{})
	param["merchant_refund_sn"] = refund.MerchantRefundSN
	param["refund_amount"] = refund.RefundAmount

	if refund.CallLevel == 1 {
		param["merchant_id"] = refund.MerchantID
	}

	if refund.OrderSN != "" {
		param["order_sn"] = refund.OrderSN
	}

	if refund.MerchantOrderSN != "" {
		param["merchant_order_sn"] = refund.MerchantOrderSN
	}

	if refund.InsOrderSN != "" {
		param["ins_order_sn"] = refund.InsOrderSN
	}

	if refund.Handler != 0 {
		param["handler"] = refund.Handler
	}

	if refund.DeviceNO != "" {
		param["device_no"] = refund.DeviceNO
	}

	data, err := util.LanuchRquest(o.config, "fbpay.order.refund", param)
	if err != nil {
		return orderRefund, err
	}
	if err := json.Unmarshal(data, &orderRefund); err != nil {
		return orderRefund, err
	}
	return
}

// RefundQuery  退款查询
func (o *Order) RefundQuery(refund OrderRefundQueryRequest) (refundQuery OrderRefundQueryResponse, err error) {
	if err := util.Validate(refund); err != nil {
		return refundQuery, err
	}
	if refund.CallLevel == 1 && refund.MerchantID == 0 {
		return refundQuery, fmt.Errorf("parameter error")
	}
	if refund.RefundSN == "" && refund.MerchantRefundSN == "" {
		return refundQuery, fmt.Errorf("parameter error")
	}
	param := make(map[string]interface{})
	if refund.CallLevel == 1 {
		param["merchant_id"] = refund.MerchantID
	}

	if refund.RefundSN != "" {
		param["refund_sn"] = refund.RefundSN
	}

	if refund.MerchantRefundSN != "" {
		param["merchant_refund_sn"] = refund.MerchantRefundSN
	}

	data, err := util.LanuchRquest(o.config, "fbpay.order.refund.query", param)
	if err != nil {
		return refundQuery, err
	}
	if err := json.Unmarshal(data, &refundQuery); err != nil {
		return refundQuery, err
	}
	return
}
