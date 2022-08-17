package common

// CommonData 统一下单和付款码支付通用数据部分
type CommonData struct {
	AlipayExtendParams *AlipayExtendParams `json:"alipay_extend_params"`                  // 支付宝业务拓展参数--花呗分期
	CallLevel          int8                `json:"call_level"`                            // 调用级别 1. 服务商 2. 商户
	StoreID            int64               `json:"store_id"`                              // 商户门店号（如果只有一家有效门店，可不传）
	CashierID          int64               `json:"cashier_id"`                            // 收银员ID
	MerchantID         int64               `json:"merchant_id"`                           // 付呗商户号。以服务商级接入时必传，以商户级接入时不传
	TotalAmount        float64             `json:"total_amount" validate:"required"`      // 订单总金额，单位为元，精确到0.01 ~ 10000000
	MerchantOrderSN    string              `json:"merchant_order_sn" validate:"required"` // 外部系统订单号（确保唯一，前后不允许带空格）
	SubAppID           string              `json:"sub_appid"`                             // 公众号appid
	GoodsTag           string              `json:"goods_tag"`                             // 订单优惠标记，代金券或立减优惠功能的参数（使用单品券时必传）
	Body               string              `json:"body"`                                  // 商品描述
	Attach             string              `json:"attach"`                                // 附加数据，原样返回，该字段主要用于商户携带订单的自定义数据
	TimeoutExpress     string              `json:"timeout_express"`                       // 订单失效时间，逾期将关闭交易。格式为yyyyMMddHHmmss，失效时间需大于1分钟银联订单请勿传，目前银联不支持该字段
	NotifyUrl          string              `json:"notify_url"`                            // 支付回调地址
	DeviceNO           string              `json:"device_no"`                             // 终端号
	PlatformStoreID    string              `json:"platform_store_id"`                     // 平台方门店号（即微信/支付宝的storeid
	DisablePayChannels string              `json:"disable_pay_channels"`                  // 禁止使用优惠券标识 promotion-支付宝优惠（包含实时优惠+商户优惠）voucher-支付宝营销券可以单个传，也可以多个传，以逗号分隔如：promotion,voucher
}

type AlipayExtendParams struct {
	HbFqInstalment    int8 `json:"hb_fq_instalment"`     // 是否使用花呗分期，默认为空（不使用花呗分期）。1是，0否
	HbFqNum           int8 `json:"hb_fq_num"`            // 花呗分期期数，默认为空（不使用花呗分期）。仅支持传入3、6、12，分别代表分期期数。instalment=1时必填
	HbFqSellerPercent int8 `json:"hb_fq_seller_percent"` // 手续费承担模式，默认为空（不使用花呗分期）。0=消费者承担；100=商户承担（间联通道暂不支持商户承担，故该参数当前无效）。instalment=1时必填
}
