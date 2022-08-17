// @Title callback
// @Description 所有的回调
package callback

import (
	"encoding/json"
	"git.myarena7.com/arena/fubeisdk/core/config"
	"git.myarena7.com/arena/fubeisdk/util"
)

type CallBack struct{ config config.Config }

func NewCallBack(cfg config.Config) *CallBack {
	return &CallBack{config: cfg}
}

// 商户审核回调
type MerchantAuditCallBack struct {
	MerchantID   string `json:"merchant_id"`   // 付呗商户编号
	MerchantCode string `json:"merchant_code"` // 商户账号
	AuditMsg     string `json:"audit_msg"`     // 商户的审核信息
	/*
		商户审核状态
		UNKNOWN 未知
		ADD_INFO 待完善资料
		AUDITING 待审核
		REJECTED 退回
		REVIEW 待人工检查
		PASSED 通过
		NOTFOUND 商户未进件
	*/
	AuditStatus string `json:"audit_status"`
	/*
		拒绝原因分类
		MCH_INFO_ERROR:商户信息错误
		LECENSE_ERROR:营业执照信息错误
		CERTIFICATE_ERROR:证件信息错误
		ACCOUNT_INFO_ERROR:结算账户信息错误
		SYSTEM_REJECT:系统拒绝
		OTHER_INFO_ERROR:其他信息错误
		1）当商户进件时，auditStatus是REJECTED时必传
	*/
	RejectReasonType string `json:"reject_reason_type"`
	AuditDateTime    string `json:"audit_date_time"` // 审核时间,格式 2018-11-01 11:22:33
	RefusalReason    string `json:"refusal_reason"`  // 拒绝驳回原因,具体明细
	Sign             string `json:"sign"`            // 签名
}

// CallBackConfigRequest 支付后回调函数请求参数
type CallBackConfigRequest struct {
	MerchantID                 int64  // 付呗商户号
	SecondCallBackUrl          string // 二次回调URL
	RemitCallBackUrl           string // 打款回调URL(没必要介入)
	RefundCallBackUrl          string // 退款回调URL
	WithdrawCallBackUrl        string // 提现回调地址
	AccountRegisterCallBackUrl string // 分账接收方回调地址
	ShareCallBackUrl           string // 分账回调地址
	MerchantAuditCallBackUrl   string // 商户审核回调地址
	ShareAuditCallBackUrl      string // 分账电子协议签署回调
	PayThirdCallbackUrl        string // 第三方支付回调地址 支付回调
}

// CallBackConfigResponse 支付后回调函数返回参数
type CallBackConfigResponse struct {
	BindStatus  int8   `json:"bind_status"` // 1 配置成功 2 配置失败
	MerchantID  int64  `json:"merchant_id"` // 付呗商户号
	AgentID     int64  `json:"agent_id"`    // 服务商id
	RespMessage string `json:"resp_message"`
}

// CallBackCommonData 支付回调公用参数
type CallBackCommonData struct {
	AlipayExtendParams AlipayExtendParams `json:"alipay_extend_params"` // 支付宝业务拓展参数--花呗分期
	PaymentList        []PaymentList      `json:"payment_list"`         // 活动优惠列表，Json格式。payment_list的值为数组，每一个元素包含两个字段type和amount
	Uid                int64              `json:"uid"`                  // 商户id 同merchant_id
	StoreID            int64              `json:"store_id"`             // 商户门店号
	CashierID          int64              `json:"cashier_id"`           // 收银员ID
	TotalAmount        float64            `json:"total_amount"`         // 订单金额，精确到0.01
	NetAmount          float64            `json:"net_amount"`           // 实收金额，精确到0.01
	BuyerPayAmount     float64            `json:"buyer_pay_amount"`     // 买家实际支付金额，精确到0.01
	Fee                float64            `json:"fee"`                  // 手续费（直连没有手续费返回），精确到0.01
	OrderSN            string             `json:"order_sn"`             // 付呗订单号
	MerchanOrderSN     string             `json:"merchant_order_sn"`    // 外部系统订单号
	InsOrderSN         string             `json:"ins_order_sn"`         // 机构订单号（显示在微信/支付宝支付凭证的订单号）
	ChannelOrderSN     string             `json:"channel_order_sn"`     // 通道订单号，微信订单号、支付宝订单号等
	OrderStatus        string             `json:"order_status"`         // 订单状态：SUCCESS--支付成功
	PayType            string             `json:"pay_type"`             // 支付方式，wxpay 微信，alipay 支付宝，unionpay 银联,bankcardpay        银行卡
	UserID             string             `json:"user_id"`              // 付款用户id，“微信openid”、“支付宝账户”、银联支付没有该字段
	FinishTime         string             `json:"finish_time"`          // 支付完成时间，格式为yyyyMMddHHmmss
	DeviceNO           string             `json:"device_no"`            // 终端号
	Attach             string             `json:"attach"`               // 附加数据，原样返回，该字段主要用于商户携带订单的自定义数据
}

// CallBackRequest 回调请求参数
type CallBackRequest struct {
	ResultCode    int32  `json:"result_code" validate:"required"`    // 状态码
	ResultMessage string `json:"result_message" validate:"required"` // 状态信息
	Sign          string `json:"sign" validate:"required"`           // 签名
	Data          string `json:"data" validate:"required"`           // 数据
}

//  RefundCallBack 退款回调业务参数data
type RefundCallBackData struct {
	Handler           int64   `json:"handler"`            // 退款操作人员ID
	BuyerRefundAmount float64 `json:"buyer_refund_amount"`  // 买家实际支付金额，精确到0.01
	RefundFee         float64 `json:"refund_fee"`         // 退款手续费
	RefundAmount      float64 `json:"refund_amount"`      // 退款金额，精确到0.01，注：使用优惠券时仅支持全额退款
	FinishTime        string  `json:"finish_time"`        // 支付完成时间，格式为yyyyMMddHHmmss
	OrderSN           string  `json:"order_sn"`           // 付呗订单号
	MerchanOrderSN    string  `json:"merchant_order_sn"`  // 外部系统订单号
	RefundSN          string  `json:"refund_sn"`          // 付呗订单号
	MerchantRefundSN  string  `json:"merchant_refund_sn"` // 外部系统退款号
	RefundStatus      string  `json:"refund_status"`      // 退款状态 REFUND_SUCCESS--退款成功
	DeviceNO          string  `json:"device_no"`          // 终端号
}

//  支付二次回调业务参数data
type PayCallBackSecondData struct {
	CallBackCommonData
}

//  支付回调业务参数data
type PayCallBackData struct {
	CallBackCommonData
	IsCanPartRefund int8   `json:"is_can_part_refund"` // 是否支持部分退款1 支持 0 不支持
	BankName        string `json:"bank_name"`          // 发卡行，刷卡交易时返回
	CardID          string `json:"card_id"`            // 银行卡号，刷卡交易时返回
	CardType        string `json:"card_type"`          // 卡类型，刷卡交易时返回0-借记卡；1-贷记卡
	BatchNO         string `json:"batch_no"`           // 刷卡交易批次号
	FlowID          string `json:"flow_id"`            // 刷卡交易凭证号
	ReferenceNumber string `json:"reference_number"`   // 刷卡交易参考号
	AuthorizeCode   string `json:"authorize_code"`     // 刷卡交易授权码，只有银行卡预授权交易才返回
}

type PaymentList struct {
	Amount float64 `json:"amount"` // 金额
	Type   string  `json:"type"`   // 付呗折扣:FUBEI_DISCOUNT  支付通道免充值:CHANNEL_DISCOUNT  支付通道预充值:CHANNEL_PRE
}

type AlipayExtendParams struct {
	HbFqInstalment    int8 `json:"hb_fq_instalment"`     // 是否使用花呗分期，默认为空（不使用花呗分期）。1是，0否
	HbFqNum           int8 `json:"hb_fq_num"`            // 花呗分期期数，默认为空（不使用花呗分期）。仅支持传入3、6、12，分别代表分期期数。instalment=1时必填
	HbFqSellerPercent int8 `json:"hb_fq_seller_percent"` // 手续费承担模式，默认为空（不使用花呗分期）。0=消费者承担；100=商户承担（间联通道暂不支持商户承担，故该参数当前无效）。instalment=1时必填
}

// SubAccountSettleBack  结算回调参数
type SubAccountSettleCallBack struct {
	CallbackType string `json:"callback_type"` // 默认值：SETTLEMENT
	MerchantID   string `json:"merchant_id"`   // 付呗商户号
	OrderSn      string `json:"order_sn"`      // 平台方结算单号
	TradeMoney   string `json:"trade_money"`   // 结算金额
	Fee          string `json:"fee"`           // 结算手续费
	BankHolder   string `json:"bank_holder"`   // 持卡人
	BankCardNO   string `json:"bank_card_no"`  // 卡号
	BankName     string `json:"bank_name"`     // 银行名称
	UnionPayCode string `json:"unionpay_code"` // 联行号
	SettleDate   string `json:"settle_date"`   // 结算日期 格式"YYYYMMdd"
	Settletime   string `json:"settle_time"`   // 结算时间 格式"YYYY-MM-dd HH:mm:ss"
	Flag         string `json:"flag"`          // 结算状态 1成功 2失败 3 出账中 4 冻结 5成功 6失效 7退票
	ResultDesc   string `json:"result_desc"`   // 打款描述 主要是针对打款失败
	Sign         string `json:"sign"`          // 签名
}

// AccountFixAmountCallBack 按照固定金额分账回调
type AccountFixAmountCallBack struct {
	ResultCode    int64  `json:"result_code"`    // 状态码
	ResultMessage string `json:"result_message"` // 状态信息
	Sign          string `json:"sign"`           // 回调签名
	Data          string `json:"data"`           // 回调业务参数
}

type AccountFixAmountCallBackInfo struct {
	MerchantID      string `json:"merchant_id"`       // 付呗商户号
	CallbackType    string `json:"callback_type"`     //  默认值：FIXED_AMOUNT_SHARE
	MerchantOrderSN string `json:"merchant_order_sn"` // 外部系统订单号（请求流水号）
	Data            string `json:"data_list"`
}

type AccountFixAmountInformation struct {
	AccountType int8    `json:"account_type"` // 账户类型：1 分账接收方 2 商户
	AccountOut  int64   `json:"account_out"`  // 分账出账户，即付呗商户号
	ShareAmount float64 `json:"share_amount"` // 分账金额
	OrderNO     string  `json:"order_no"`     // 分账账单订单号
	AccountIn   string  `json:"account_in"`   // 入账户id，当账户类型为1时表示分账接收方，当账户类型为2时表示付呗商户号
	ShareStatus string  `json:"share_status"` // 分账状态: SUCCESS 分账成功
	ShareTime   string  `json:"share_time"`   // 分账时间，yyyyMMddHHmmss
}

// SubAccountIncomeCallBack 分账接收方入驻回调
type SubAccountIncomeCallBack struct {
	ResultCode    int64  `json:"result_code"`    // 状态码
	ResultMessage string `json:"result_message"` // 状态信息
	Sign          string `json:"sign"`           // 回调签名
	Data          string `json:"data"`           // -	回调业务参数
}

type SubAccountIncomeCallBackInfo struct {
	Status          string `json:"status"`            // 入驻状态：1入驻成功、3入驻失败
	CallbackType    string `json:"callback_type"`     // 默认值：SHARE_RECEIVER_INCOME
	MerchantOrderSN string `json:"merchant_order_sn"` // 外部系统订单号（请求流水号）
	AccountID       string `json:"account_id"`        // 分账接收方id
	FailMessage     string `json:"fail_message"`      // 失败原因
}

type WithDrawCallBack struct {
	CallbackType    string                     `json:"callback_type"`
	MerchantID      int64                      `json:"merchant_id"`
	AccountID       string                     `json:"account_id"`
	MerchantOrderSN string                     `json:"merchant_order_sn"`
	DataList        []WithDrawCallBackListInfo `json:"data_list"`
}

type WithDrawCallBackListInfo struct {
	WithDrawNo   string  `json:"withdraw_no"`
	StartTime    string  `json:"start_time"`
	CompleteTime string  `json:"complete_time"`
	Amount       float64 `json:"amount"`
	Fee          float64 `json:"fee"`
	Status       int32   `json:"status"`
	BankCardNO   string  `json:"bank_card_no"`
	BankName     string  `json:"bank_name"`
	CreateTime   string  `json:"create_time"`
	FinishTime   string  `json:"finish_time"`
}

type ElecContractCallBack struct {
	CallBackType string `json:"callback_type"`
	MerchantID   string `json:"merchant_id"`
}

// CheckSign 回调验签
func CheckSign(appSecret string, callbakRequest CallBackRequest) bool {
	param := make(map[string]interface{}, 0)
	param["result_code"] = callbakRequest.ResultCode
	param["result_message"] = callbakRequest.ResultMessage
	param["data"] = callbakRequest.Data
	return util.SignMD5(appSecret, param) == callbakRequest.Sign
}

// CallBackConfig 支付回调配置
func (c *CallBack) CallBackConfig(callback CallBackConfigRequest) (resp CallBackConfigResponse, err error) {
	param := make(map[string]interface{}, 0)
	if callback.MerchantID != 0 {
		param["merchant_id"] = callback.MerchantID
	}
	if callback.SecondCallBackUrl != "" {
		param["second_callback_url"] = callback.SecondCallBackUrl
	}
	if callback.RemitCallBackUrl != "" {
		param["remit_callback_url"] = callback.RemitCallBackUrl
	}
	if callback.RefundCallBackUrl != "" {
		param["refund_callback_url"] = callback.RefundCallBackUrl
	}
	if callback.WithdrawCallBackUrl != "" {
		param["withdraw_callback_url"] = callback.WithdrawCallBackUrl
	}
	if callback.AccountRegisterCallBackUrl != "" {
		param["account_register_callback_url"] = callback.AccountRegisterCallBackUrl
	}
	if callback.ShareCallBackUrl != "" {
		param["share_callback_url"] = callback.ShareCallBackUrl
	}
	if callback.MerchantAuditCallBackUrl != "" {
		param["auth_callback_url"] = callback.MerchantAuditCallBackUrl
	}
	if callback.ShareAuditCallBackUrl != "" {
		param["share_audit_callback_url"] = callback.ShareAuditCallBackUrl
	}
	if callback.PayThirdCallbackUrl != "" {
		param["callback_url"] = callback.PayThirdCallbackUrl
	}
	data, err := util.LanuchRquest(c.config, "fbpay.pay.callback.config", param)
	if err != nil {
		return resp, err
	}
	if err = json.Unmarshal(data, &resp); err != nil {
		return resp, nil
	}
	return
}
