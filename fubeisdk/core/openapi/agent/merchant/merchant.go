package merchant

import (
	"encoding/json"
	"fmt"
	"strconv"

	"git.myarena7.com/arena/fubeisdk/core/config"
	"git.myarena7.com/arena/fubeisdk/util"
)

type AgentMerchant struct {
	config config.Config
}

func NewAgentMerchant(cfg config.Config) *AgentMerchant {
	return &AgentMerchant{config: cfg}
}

// openapi.agent.merchant.income.audit.query MerchantIncomeAuditResponse 商户审核结果查询
type MerchantIncomeAuditResponse struct {
	MerchantID int64  `json:"merchant_id"` // 付呗商户编号
	AuditMsg   string `json:"audit_msg"`   // 商户的审核信息
	/*
		UNKNOWN 未知
		ADD_INFO 待完善资料
		AUDITING 待审核
		REJECTED 退回
		REVIEW 待人工检查
		PASSED 通过
		NOTFOUND 商户未进件
	*/
	AuditStatus string `json:"audit_status"` // 商户审核状态
	/*
	   CH_INFO_ERROR:商户信息错误
	   LECENSE_ERROR:营业执照信息错误
	   CERTIFICATE_ERROR:证件信息错误
	   ACCOUNT_INFO_ERROR:结算账户信息错误
	   SYSTEM_REJECT:系统拒绝
	   OTHER_INFO_ERROR:其他信息错误
	   1）当商户进件时，auditStatus是REJECTED时必传
	*/
	RejectReasonType string `json:"reject_reason_type"` // 拒绝原因分类
	AuditDateTime    string `json:"audit_date_time"`    // 审核时间,格式 2018-11-01 11:22:33
	RefusalReason    string `json:"refusal_reason"`     // 拒绝驳回原因,具体明细
}

// openapi.agent.merchant.querystatus MerchantStatusQueryResponse 商户门店状态查询
type MerchantStatusQueryResponse struct {
	StoreList      []Storelist `json:"store_list"`      // 门店信息列表
	MerchantStatus int8        `json:"merchant_status"` // 认证状态 0 未认证 1 认证中 2 认证成功 3. 认证失败
	MerchantID     int64       `json:"merchant_id"`     // 商户号
	MerchanCode    string      `json:"merchant_code"`   // 商户账号 可以用作登陆商户后台
	ErrMsg         string      `json:"err_msg"`         // 商户状态错误信息
}

type Storelist struct {
	StoreStatus      int8   `json:"store_status"`       // 门店状态 1 待审核 2 审核通过 3 审核驳回
	StoreID          int64  `json:"store_id"`           // 门店ID
	SubMchID         string `json:"sub_mch_id"`         // 当前渠道下对应的微信子商户号码
	AlipayMsID       string `json:"alipay_msid"`        // 当前渠道下对应的支付宝子商户号码
	StoreErrorMsg    string `json:"store_error_msg"`    // 门店错误信息
	WechatAuthStatus string `json:"wechat_auth_status"` // 商户微信认证状态 未认证 已经认证 未知
}

// openapi.agent.merchant.queryinfo MechantQueryResponse 商户门店信息查询
type MechantQueryResponse struct {
	StoreInfoList      []StoreInfoList `json:"store_info_list"`       // 门店信息列表
	IncomeStatus       int8            `json:"income_status"`         // 进件状态 1 未进件 2 进件中 3 进件成功 4 进件失败
	AuthStatus         int8            `json:"auth_status"`           // 认证状态 0 未认证 1 认证中 2 认证成功 3 认证失败
	AccountType        int8            `json:"account_type"`          // 账户类型 1. 个人 2 企业
	MerchantID         int64           `json:"merchant_id"`           // 商户号
	UnityCategoryID    int64           `json:"unity_category_id"`     // 行业类目
	RealName           string          `json:"real_name"`             // 姓名
	ContactPhone       string          `json:"contact_phone"`         // 手机号码
	ServicePhone       string          `json:"service_phone"`         // 客服电话
	Email              string          `json:"email"`                 // 邮箱
	MerchanCode        string          `json:"merchant_code"`         // 商户账号 可以用作登陆商户后台
	ProvinceCode       string          `json:"province_code"`         // 省份编码
	CityCode           string          `json:"city_code"`             // 城市编码
	AreaCode           string          `json:"area_code"`             // 区域编码
	StreetAddress      string          `json:"street_address"`        // 详细地址
	BankCardNO         string          `json:"bank_card_no"`          // 银行卡号
	BankName           string          `json:"bank_name"`             // 银行名称
	BranchName         string          `json:"branch_name"`           // 支行名称
	UnionpayCode       string          `json:"unionpay_code"`         // 开户支行联行号码
	BankCode           string          `json:"bank_code"`             // 银行编号
	BankCellPhone      string          `json:"bank_cell_phone"`       // 银行预留手机号码
	IdCardNO           string          `json:"id_card_no"`            // 身份证号码
	AlipayFeeRate      string          `json:"alipay_fee_rate"`       // 支付宝商户终端费率 3.8-50
	WxFeeRateFloat     string          `json:"wx_fee_reate_float"`    // 微信商户终端费率 3.8-50
	UnionFeeRateFloat  string          `json:"union_fee_rate_float"`  // 银联商户终端费率 3.8-50
	CallBack           string          `json:"call_back"`             // 支付回调地址
	AppID              string          `json:"app_id"`                // appid
	AppSecret          string          `json:"app_secret"`            // appsecret
	CreditCardFee      string          `json:"credit_card_fee"`       // 商户信用卡费率
	DebitCardFee       string          `json:"debit_card_fee"`        // 商户借记卡费率
	DebitCardAppendFee string          `json:"debit_card_append_fee"` // 商户借记卡封顶手续费
}

type StoreInfoList struct {
	StoreStatus        int8   `json:"store_status"`         // 门店状态 1. 待审核 2. 审核通过 3 审核驳回
	StoreID            int64  `json:"store_id"`             // 门店id
	LicenseType        string `json:"license_type"`         // 门店经营许可证类型 1. 营业执照 2. 证明函
	StoreName          string `json:"store_name"`           // 门店名称
	SubMchID           string `jsn:"sub_mch_id"`            // 当前渠道微信子商户号码
	AlipayMsID         string `json:"alipay_msid"`          // 当前渠道支付宝子商户号码
	LicenseTimeType    string `json:"license_time_type"`    // 营业执照有效期类型 1. 正常有效期 2. 长期有效
	StorePhone         string `json:"store_phone"`          // 门店电话
	StoreProvinceCode  string `json:"store_province_code"`  // 省份编码
	StoreCityCode      string `json:"store_city_code"`      // 城市编码
	StoreAreaCode      string `json:"store_area_code"`      // 区域编码
	Address            string `json:"address"`              // 省市区中文地址
	StoreStreetAddress string `json:"store_street_address"` // 门店详细地址
	Longitude          string `json:"longitude"`            // 门店地址经度
	Latitude           string `json:"latitude"`             // 门店地址纬度
	LicenseName        string `json:"license_name"`         // 营业执照名称
	LicenseId          string `json:"license_id"`           // 营业执照号
	LicenseTimeBegin   string `json:"license_time_begin"`   // 营业执照开始时间
	LicenseTimeEnd     string `json:"license_time_end"`     // 营业执照结束时间
	Remark             string `json:"remark"`               // 备注
	PayAppID           string `json:"-"`                    // 付呗那边接口有返回但是含义未知 暂时不接
	FollowAppID        string `json:"-"`                    // 付呗那边接口有返回但是含义未知
	MinaAppID          string `json:"-"`                    // 付呗那边接口有返回但是含义未知
	WechatAuthStatus   string `json:"wechat_auth_status"`   // 商户微信认证状态 UNAUTHORIZED 未认证  AUTHORIZED 已认证  UNKNOWN 未知
}

// openapi.agent.merchant.adjustrate MerchantRateModifyRequest // 商户费率修改
type MerchantRateModifyRequest struct {
	MerchantID        int64  `json:"merchant_id"`          // 付呗商户号
	MerchantCode      string `json:"merchant_code"`        // 商户账号，与付呗商户号二选一
	AlipayFeeRate     string `json:"alipay_fee_rate"`      // 非必填 支付宝商户终端费率（‰）
	WxFeeRateFloat    string `json:"wx_fee_rate_float"`    // 非必填 微信商户终端费率（‰）
	UnionFeeRateFloat string `json:"union_fee_rate_float"` // 银联商户终端费率（‰）
	CreditCardFee     string `json:"credit_card_fee"`      // 非必填 商户贷记卡（信用卡）费率 千分之 5.2 ~ 50 保留俩位小数 和商户借记卡费率，商户借记卡封顶手续费三者同时传值或都不传
	DebitCardFee      string `json:"debit_card_fee"`       // 非必填 商户借记卡费率 千分之 4.2 ~ 50 保留俩位小数
	DebitCardAppedFee string `json:"debit_card_apped_fee"` // 非必填 商户借记卡封顶手续费最小18元 最大不能超过500元 封顶手续费只能为整数
}

type MerchatRateModifyResponse struct {
	MerchantID                    int64  `json:"merchant_id"`                      // 商户号
	MerchanCode                   string `json:"merchant_code"`                    // 商户账号 可以用作登陆商户后台
	AlipayFeeRate                 string `json:"alipay_fee_rate"`                  // 已生效的支付宝商户终端费率
	WxFeeRateFloat                string `json:"wx_fee_rate_float"`                // 已生效微信商户终端费率
	UnionFeeRateFloat             string `json:"union_fee_rate_float"`             // 已生效银联商户终端费率
	CreditCardFee                 string `json:"credit_card_fee"`                  // 已生效商户信用卡费率
	DebitCardFee                  string `json:"debit_card_fee"`                   // 已生效商户借记卡费率
	DebitCardAppendFee            string `json:"debit_card_apped_fee"`             // 已生效商户借记卡封顶手续费最小18元 最大不能超过500元 封顶手续费只能为整数
	IneffectiveAlipayFeeRate      string `json:"ineffective_alipay_fee_rate"`      // 次日待生效支付宝商户终端费率（‰），范围：3.8~50
	IneffecctiveWxFeeRateFloat    string `json:"ineffective_wx_fee_rate_float"`    // 次日待生效微信商户终端费率（‰），范围：3.8~50
	IneffectiveUnionFeeRateFloat  string `json:"ineffective_union_fee_rate_float"` // 次日待生效银联商户终端费率（‰），范围：3.8~50
	IneffectiveCreditCardFee      string `json:"ineffective_credit_card_fee"`      // 次日待生效商户贷记卡（信用卡）费率 千分之 5.2 ~ 50 保留俩位小数
	IneffectiveDebitCardFee       string `json:"ineffective_debit_card_fee"`       // 次日待生效商户借记卡费率 千分之 4.2 ~ 50 保留俩位小数
	IneffactiveDebitCardAppendFee string `json:"ineffective_debit_card_apped_fee"` // 次日待生效商户借记卡封顶手续费最小18元 最大不能超过500元 封顶手续费只能为整数
}

// openapi.agent.merchant.update.bank.card MerchantUpdateBankResponse 商户换绑卡
type MerchantUpdateBankRequest struct {
	MerchantCodeType int8   `json:"merchant_code_type"`                  // 修改账户类型 1. 个人账户 2. 对公账户
	MerchantID       int64  `json:"merchant_id"`                         // 付呗商户号
	MerchantCode     string `json:"merchant_code"`                       // 商户账号
	BankCardNO       string `json:"bank_card_no" validate:"required"`    // 银行卡号
	BankCardImage    string `json:"bank_card_image" validate:"required"` // 卡证明照片
	BankCellPhone    string `json:"bank_cell_phone"`                     // 预留手机号码
	BankCode         string `json:"bank_code"`                           // 银行编号 个人账号必填
	BankName         string `json:"bank_name"`                           // 银行名称 对公必填
	UnionPayCode     string `json:"unionpay_code"`                       // 开户支行联行号 对公账户必填
}

type MerchantUpdateBankResponse struct {
	ModifyFlag  bool   `json:"modify_flag"`   // 是否修改成功 (最终是否可用以bind_status为准)
	BindStatus  int8   `json:"bind_status"`   // 1:审核中 2:审核成功 3:审核驳回 (卡最终是否可用以bind_status是否审核成功为准)
	MerchantID  int64  `json:"merchant_id"`   // 商户号
	MerchanCode string `json:"merchant_code"` // 商户账号 可以用作登陆商户后台
	RespMessage string `json:"resp_message"`  // 信息
}

// openapi.agent.merchant.query.bank.card.bindStatus MerchantBankBindStatusQueryResponse  商户换绑卡状态查询
type MerchantBankBindStatusQueryRequest struct {
	MerchantID   int64  `json:"merchant_id"`                      // 付呗商户号
	MerchantCode string `json:"merchant_code"`                    // 商户账号，与付呗商户号二选一
	BankCardNO   string `json:"bank_card_no" validate:"required"` // 银行卡号
	BankCode     string `json:"bank_code" validate:"required"`    // 银行编号（
}

type MerchantBankBindStatusQueryResponse struct {
	BindStatus  int8   `json:"bind_status"`   // 1:审核中 2:审核成功 3:审核驳回
	MerchantID  int64  `json:"merchant_id"`   // 商户号
	MerchanCode string `json:"merchant_code"` // 商户账号 可以用作登陆商户后台
	RejectMsg   string `json:"reject_msg"`    // 换绑卡驳回信息
}

// MerchantIncomeAuditResult 商户审核结果查询 频率限制位 80/20s
func (m *AgentMerchant) MerchantIncomeAuditResult(merchantID int64) (merchantAudit MerchantIncomeAuditResponse, err error) {
	if merchantID == 0 {
		return merchantAudit, fmt.Errorf("parameter error")
	}
	data, err := util.LanuchRquest(m.config, "openapi.agent.merchant.income.audit.query", map[string]int64{
		"merchant_id": merchantID,
	})
	if err != nil {
		return merchantAudit, err
	}
	if err = json.Unmarshal(data, &merchantAudit); err != nil {
		return merchantAudit, err
	}
	return
}

// MerchantStoreStatus 商户门店状态查询
func (m *AgentMerchant) MerchantStoreStatus(merchantCode string) (merchantStore MerchantStatusQueryResponse, err error) {
	if merchantCode == "" {
		return merchantStore, fmt.Errorf("parameter error")
	}
	data, err := util.LanuchRquest(m.config, "openapi.agent.merchant.querystatus", map[string]string{
		"merchant_code": merchantCode,
	})
	if err != nil {
		return merchantStore, err
	}
	if err = json.Unmarshal(data, &merchantStore); err != nil {
		return merchantStore, err
	}
	return
}

// MerchantStoreInfo 商户门店信息查询
func (m *AgentMerchant) MerchantStoreInfo(merchantCode string) (merchantStore MechantQueryResponse, err error) {
	if merchantCode == "" {
		return merchantStore, fmt.Errorf("parameter error")
	}
	data, err := util.LanuchRquest(m.config, "openapi.agent.merchant.queryinfo", map[string]string{
		"merchant_code": merchantCode,
	})
	if err != nil {
		return merchantStore, err
	}
	fmt.Println(string(data), "=====")
	if err = json.Unmarshal(data, &merchantStore); err != nil {
		return merchantStore, err
	}
	return
}

// MerchantFeeRateModify 商户费率修改 如果不修改传递空值 如果需要修改 不能传递空值
func (m *AgentMerchant) MerchantFeeRateModify(merchantNewRate MerchantRateModifyRequest) (merchantRate MerchatRateModifyResponse, err error) {
	if err := util.Validate(merchantNewRate); err != nil {
		return merchantRate, err
	}
	// 两个至少填写一个
	if merchantNewRate.MerchantID == 0 && merchantNewRate.MerchantCode == "" {
		return merchantRate, fmt.Errorf("parameter error")
	}
	param := make(map[string]interface{}, 0)
	param["merchant_id"] = merchantNewRate.MerchantID
	param["merchant_code"] = merchantNewRate.MerchantCode

	if merchantNewRate.AlipayFeeRate != "" {
		param["alipay_fee_rate"] = merchantNewRate.AlipayFeeRate
	}

	if merchantNewRate.WxFeeRateFloat != "" {
		param["wx_fee_rate_float"] = merchantNewRate.WxFeeRateFloat
	}

	if merchantNewRate.UnionFeeRateFloat != "" {
		param["union_fee_rate_float"] = merchantNewRate.UnionFeeRateFloat
	}

	if merchantNewRate.CreditCardFee != "" {
		param["credit_card_fee"] = merchantNewRate.CreditCardFee
	}

	if merchantNewRate.DebitCardFee != "" {
		param["debit_card_fee"] = merchantNewRate.DebitCardFee
	}

	if merchantNewRate.DebitCardAppedFee != "" {
		// 大于等于18 小于等于500
		debitCardAppenFee, err := strconv.ParseInt(merchantNewRate.DebitCardAppedFee, 10, 64)
		if err != nil {
			return merchantRate, fmt.Errorf("debit_card_apped_fee parameter error")
		}
		if debitCardAppenFee < 18 || debitCardAppenFee > 500 {
			return merchantRate, fmt.Errorf("debit_card_apped_fee parameter error")
		}
		param["debit_card_apped_fee"] = merchantNewRate.DebitCardAppedFee
	}

	data, err := util.LanuchRquest(m.config, "openapi.agent.merchant.adjustrate", param)
	if err != nil {
		return merchantRate, err
	}
	if err = json.Unmarshal(data, &merchantRate); err != nil {
		return merchantRate, err
	}
	return
}

// MerchantUpdateBank 商户换绑卡
func (m *AgentMerchant) MerchantUpdateBank(merchantBank MerchantUpdateBankRequest) (merchantNewBank MerchantUpdateBankResponse, err error) {
	if err := util.Validate(merchantBank); err != nil {
		return merchantNewBank, err
	}
	// 两个至少填写一个
	if merchantBank.MerchantID == 0 && merchantBank.MerchantCode == "" {
		return merchantNewBank, fmt.Errorf("parameter error")
	}

	if merchantBank.MerchantCodeType == 1 && merchantBank.BankCode == "" { // 个人账户没有填写bankcode
		return merchantNewBank, fmt.Errorf("个人账户银行编号必填")
	}

	if (merchantBank.MerchantCodeType == 2 && merchantBank.BankName == "") || (merchantBank.MerchantCodeType == 2 && merchantBank.UnionPayCode == "") { // 对公账户
		return merchantNewBank, fmt.Errorf("对公账户银行名称,开户支行联行号必填")
	}

	param := make(map[string]interface{}, 0)
	if merchantBank.MerchantID != 0 {
		param["merchant_id"] = merchantBank.MerchantID
	}

	if merchantBank.MerchantCode != "" {
		param["merchant_code"] = merchantBank.MerchantCode
	}

	param["bank_card_no"] = merchantBank.BankCardNO
	param["bank_card_image"] = merchantBank.BankCardImage
	param["bank_cell_phone"] = merchantBank.BankCellPhone

	switch merchantBank.MerchantCodeType {
	case 1:
		param["bank_code"] = merchantBank.BankCode
	case 2:
		param["bank_name"] = merchantBank.BankName
		param["unionpay_code"] = merchantBank.UnionPayCode
	default:
		return merchantNewBank, fmt.Errorf("账户类型错误")
	}
	fmt.Printf("%+v", param)
	data, err := util.LanuchRquest(m.config, "openapi.agent.merchant.update.bank.card", param)
	if err != nil {
		return merchantNewBank, err
	}
	if err = json.Unmarshal(data, &merchantNewBank); err != nil {
		return merchantNewBank, err
	}
	return
}

// MerchantUpdateBankStatus 商户换绑卡状态查询
func (m *AgentMerchant) MerchantUpdateBankStatus(merchantNewBankStatus MerchantBankBindStatusQueryRequest) (merchantBindStatus MerchantBankBindStatusQueryResponse, err error) {
	if err := util.Validate(merchantNewBankStatus); err != nil {
		return merchantBindStatus, err
	}
	// 两个至少填写一个
	if merchantNewBankStatus.MerchantID == 0 && merchantNewBankStatus.MerchantCode == "" {
		return merchantBindStatus, fmt.Errorf("parameter error")
	}
	data, err := util.LanuchRquest(m.config, "openapi.agent.merchant.query.bank.card.bindStatus", merchantNewBankStatus)
	if err != nil {
		return merchantBindStatus, err
	}
	if err = json.Unmarshal(data, &merchantBindStatus); err != nil {
		return merchantBindStatus, err
	}
	return
}
