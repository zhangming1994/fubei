// @Title  merchant
// @Description  商户入驻 商户进件状态查询 商户信息修改接口
package merchant

import (
	"encoding/json"
	"fmt"

	"git.myarena7.com/arena/fubeisdk/core/config"
	"git.myarena7.com/arena/fubeisdk/util"
)

type Merchant struct {
	config config.Config
}

func NewMerchant(cfg config.Config) *Merchant {
	return &Merchant{config: cfg}
}

// MerchantIncomeRequest 商户入驻请求信息
type MerchantIncomeRequest struct {
	BaseInfo     BaseInfo    `json:"base_info"`                         // 商户基本信息
	LegalPerson  LegalPerson `json:"legal_person"`                      // 法人信息
	AddressInfo  AddressInfo `json:"address_info"`                      // 店铺地址信息
	LicenseInfo  LicenseInfo `json:"license_info"`                      // 营业资质信息
	AccountInfo  AccountInfo `json:"account_info"`                      // 结算信息
	ShopInfo     ShopInfo    `json:"shop_info"`                         // 门店信息
	RateInfo     RateInfo    `json:"rate_info"`                         // 终端费率
	OtherInfo    OtherInfo   `json:"other_info"`                        // 其他信息
	MerchantCode string      `json:"merchant_code" validate:"required"` // 商户账号
	SalesMan     string      `json:"sales_man"`                         // 受理商
}

// MerchantIncomeModifyRequest 商户修改
type MerchantIncomeModifyRequest struct {
	BaseInfo     BaseInfo    `json:"base_info"`                         // 商户基本信息
	LegalPerson  LegalPerson `json:"legal_person"`                      // 法人信息
	AddressInfo  AddressInfo `json:"address_info"`                      // 店铺地址信息
	LicenseInfo  LicenseInfo `json:"license_info"`                      // 营业资质信息
	AccountInfo  AccountInfo `json:"account_info"`                      // 结算信息
	ShopInfo     ShopInfo    `json:"shop_info"`                         // 门店信息
	OtherInfo    OtherInfo   `json:"other_info"`                        // 其他信息
	MerchantCode string      `json:"merchant_code" validate:"required"` // 商户账号
}

// MerchantIncomeRequest-LegalPerson 商户入驻请求信息-法人信息
type LegalPerson struct {
	LegalName             string `json:"legal_name" validate:"required"`                // 法人姓名
	LegalNum              string `json:"legal_num"  validate:"required"`                // 法人身份证号
	LegalIdCardFrontPhoto string `json:"legal_id_card_front_photo" validate:"required"` // 法人身份证正面 使用图片上传接口返回的sourceid
	LegalIdCardBackPhoto  string `json:"legal_id_card_back_photo"  validate:"required"` // 法人身份证反面 使用图片上传接口返回的sourceid
	HandHoldIdCardPhoto   string `json:"hand_hold_id_card_photo"`                       // 手持身份证照片 使用图片上传接口返回的sourceid
	LegalIdCardStart      string `json:"legal_id_card_start" validate:"required"`       // 法人身份证有效期开始时间
	LegalIdCardEnd        string `json:"legal_id_card_end"`                             // 法人身份证有效期结束时间
}

// MerchantIncomeRequest-BaseInfo 商户入驻请求信息-商户基本信息
type BaseInfo struct {
	UnityCategoryID   int8   `json:"unity_category_id"   validate:"required"`       // 行业类目
	MerchantType      int8   `json:"merchant_type"       validate:"required"`       // 商户类型 1.小微 2.个体 3.企业 无营业执照是小微
	MerchantShortName string `json:"merchant_short_name" validate:"required"`       // 商户简称 不能是全数字 不能包含敏感字符 和营业执照关联
	ContactPhone      string `json:"contact_phone"       validate:"required"`       // 商户手机号码
	Email             string `json:"email"               validate:"required,email"` // 商户电子邮箱
	ServicePhone      string `json:"service_phone"`                                 // 客服电话 可以是空
}

// MerchantIncomeRequest-AddressInfo 商户入驻请求信息-店铺地址信息
type AddressInfo struct {
	ProvinceCode  string `json:"province_code" validate:"required"`       // 商户所在省份 使用基本接口中间的接口获取或者使用附件中间的自己做一份里面的code值
	CityCode      string `json:"city_code" validate:"required"`           // 商户所在城市 全国省市自治区编号的code值
	AreaCode      string `json:"area_code" validate:"required"`           // 商户所在区 全国省市自治区编号的code值
	StreetAddress string `json:"street_address" validate:"required"`      // 商户详细地址
	Longitude     string `json:"longitude" validate:"required,longitude"` // 门店地址经度精确到小数点后6位
	Latitude      string `json:"latitude" validate:"required,latitude"`   // 门店地址纬度精确到小数点后6位
}

// MerchantIncomeRequest-LicenseInfo 商户入驻请求信息-营业资质信息 小微进件这个不传递
type LicenseInfo struct {
	LicenseID        string `json:"license_id"`         // 营业执照注册号
	LicenseName      string `json:"license_name"`       // 营业执照名称
	LicenseAddress   string `json:"license_address"`    // 营业执照注册地址
	LicenseTimeStart string `json:"license_time_start"` // 营业执照注册号开始时间
	LicenseTimeEnd   string `json:"license_time_end"`   // 营业执照注册号结束时间
	LicensePhoto     string `json:"license_photo"`      // 营业执照图片地址 使用图片上传接口返回的sourceid
}

// MerchantIncomeRequest-AccountInfo 商户入驻请求信息-结算信息
type AccountInfo struct {
	AccountType         int8   `json:"account_type" validate:"required"`    // 账户类型 1. 个人账户 2. 公司账户
	LegalFlag           int8   `json:"legal_flag" validate:"required"`      // 结算标志 0 非法人结算 1. 法人结算
	UnionPayCode        string `json:"unionpay_code" validate:"required"`   // 开户支行联行号
	RealName            string `json:"real_name" validate:"required"`       // 开户名
	IdCardNO            string `json:"id_card_no"`                          // 结算人身份证号
	IdCardFrontPhoto    string `json:"id_card_front_photo"`                 // 结算人身份证人像面图片 非法人结算必填
	IdCardBackPhoto     string `json:"id_card_back_photo"`                  // 结算人身份证国徽面 法人结算可不填写 非法人必填
	BankCellPhoto       string `json:"bank_cell_phone"`                     // 银行预留号码
	BankCardNo          string `json:"bank_card_no" validate:"required"`    // 银行卡号
	BankCardPhoto       string `json:"bank_card_photo" validate:"required"` // 银行卡正面
	UnincorporatedPhoto string `json:"unincorporated_photo"`                // 非法人结算授权书
}

// MerchantIncomeRequest-ShopInfo 商户入驻请求信息-门店信息
type ShopInfo struct {
	StoreName       string `json:"store_name" validate:"required"`        // 门店名称
	StorePhone      string `json:"store_phone" validate:"required"`       // 门店电话
	StoreEnvPhoto   string `json:"store_env_photo" validate:"required"`   // 经营场所照片 上传之后的sourceid
	StoreFrontPhoto string `json:"store_front_photo" validate:"required"` // 门头照 上传之后的sourceid
	StoreCashPhoto  string `json:"store_cash_photo" validate:"required"`  // 收银台照片 上传之后的sourceid
}

// MerchantIncomeRequest-RateInfo 商户入驻请求信息-终端费率
type RateInfo struct {
	AlipayFeeRate       string `json:"alipay_fee_rate"`         // 支付宝商户终端费率（‰）。注：若无调整费率权限时勿传，否则会报错
	WxFeeRate           string `json:"wx_fee_rate"`             // 微信商户终端费率（‰）。注：若无调整费率权限时勿传，否则会报错
	UnionFeeRate        string `json:"union_fee_rate"`          // 银联商户终端费率（‰）。注：若无调整费率权限时勿传，否则会报错
	CreditCardFeeRate   string `json:"credit_card_fee_rate"`    // 商户贷记卡（信用卡）费率 千分之 5.2 ~ 50 保留俩位小数 和商户借记卡费率，商户借记卡封顶手续费三者同时传值或都不传
	DebitCardFeeRate    string `json:"debit_card_fee_rate"`     // 商户借记卡费率 千分之 4.2 ~ 50 保留俩位小数
	DebitCardMaxFeeRate string `json:"debit_card_max_fee_rate"` // 商户借记卡封顶手续费最小18元 最大不能超过500元 封顶手续费只能为整数
}

// MerchantIncomeRequest-OtherInfo 商户入驻请求信息-其他信息
type OtherInfo struct {
	OperatintLicensePhoto string `json:"operating_license_photo"` // 经营许可证图片 使用图片上传之后返回的sourceid
	Remark                string `json:"remark"`                  // 备注
	CallBackUrl           string `json:"call_back_url"`           // 支付回调地址
}

// openapi.merchant.income MerchantIncomeResponse 商户入驻返回信息
type MerchantIncomeResponse struct {
	WxAuthentication int8   `json:"-,omitempty"`   // 是否微信实名
	WxConfig         int8   `json:"-,omitempty"`   // 微信配置是否搞了
	MerchantID       int64  `json:"merchant_id"`   // 商户号
	StoreID          int64  `json:"store_id"`      // 门店ID
	MerchanCode      string `json:"merchant_code"` // 商户账号 可以用作登陆商户后台
	AppID            string `json:"app_id"`        // 商户AppID
	AppSecret        string `json:"app_secret"`    // 商户Appsecret
}

// openapi.merchant.income.status.query MerchantIncomStatusResponse 商户进件状态查询返回信息
type MerchantIncomStatusResponse struct {
	MerchantID    int64  `json:"merchant_id"`    // 商户号
	MerchanCode   string `json:"merchant_code"`  // 商户账号 可以用作登陆商户后台
	InComeStatus  string `json:"income_status"`  // 状态 rejected 驳回 auditing 审核中 passed 通过
	RefusalReason string `json:"refusal_reason"` // 驳回原因
}

// openapi.merchant.income.modify MerchantModifyResponse 商户修改返回信息
type MerchantModifyResponse struct {
	MerchantID  int64  `json:"merchant_id"`   // 商户号
	StoreID     int64  `json:"store_id"`      // 门店ID
	MerchanCode string `json:"merchant_code"` // 商户账号 可以用作登陆商户后台
}

// MerchantIncome 商户入驻
func (m *Merchant) MerchantIncome(body MerchantIncomeRequest) (merchantIncome MerchantIncomeResponse, err error) {
	if err := util.Validate(body); err != nil {
		return merchantIncome, err
	}
	param := map[string]interface{}{
		"base_info":     body.BaseInfo,
		"legal_person":  body.LegalPerson,
		"address_info":  body.AddressInfo,
		"account_info":  body.AccountInfo,
		"shop_info":     body.ShopInfo,
		"rate_info":     body.RateInfo,
		"other_info":    body.OtherInfo,
		"merchant_code": body.MerchantCode,
	}
	if body.BaseInfo.MerchantType != 1 { // 非小微进件
		param["license_info"] = body.LicenseInfo
	} else if body.BaseInfo.MerchantType == 1 { // 小微进件 手持身份证必须传递
		if body.LegalPerson.HandHoldIdCardPhoto == "" {
			return merchantIncome, fmt.Errorf("小微进件 法人必须上传手持身份证照")
		}
	}
	data, err := util.LanuchRquest(m.config, "openapi.merchant.income", param)
	if err != nil {
		return merchantIncome, err
	}
	if err = json.Unmarshal(data, &merchantIncome); err != nil {
		return merchantIncome, nil
	}
	return
}

// MerchantIncomStatus 商户进件状态查询
func (m *Merchant) MerchantIncomeStatus(merchantCode string) (merchantIncomeStatus MerchantIncomStatusResponse, err error) {
	if merchantCode == "" {
		return merchantIncomeStatus, fmt.Errorf("parameter error")
	}
	data, err := util.LanuchRquest(m.config, "openapi.merchant.income.status.query", map[string]string{
		"merchant_code": merchantCode,
	})
	if err != nil {
		return merchantIncomeStatus, err
	}
	if err = json.Unmarshal(data, &merchantIncomeStatus); err != nil {
		return merchantIncomeStatus, nil
	}
	return
}

// MerchantModify 商户修改接口 进件状态中间如果是驳回 修改之后就需要调用这个接口
func (m *Merchant) MerchantModify(body MerchantIncomeModifyRequest) (merchantModify MerchantModifyResponse, err error) {
	if err := util.Validate(body); err != nil {
		return merchantModify, err
	}
	param := map[string]interface{}{
		"base_info":     body.BaseInfo,
		"legal_person":  body.LegalPerson,
		"address_info":  body.AddressInfo,
		"account_info":  body.AccountInfo,
		"shop_info":     body.ShopInfo,
		"other_info":    body.OtherInfo,
		"merchant_code": body.MerchantCode,
	}
	if body.BaseInfo.MerchantType != 1 { // 非小微修改
		param["license_info"] = body.LicenseInfo
	} else if body.BaseInfo.MerchantType == 1 { // 小微修改 手持身份证必须传递
		if body.LegalPerson.HandHoldIdCardPhoto == "" {
			return merchantModify, fmt.Errorf("小微进件修改 法人必须上传手持身份证照")
		}
	}
	data, err := util.LanuchRquest(m.config, "openapi.merchant.income.modify", param)
	if err != nil {
		return merchantModify, err
	}
	if err = json.Unmarshal(data, &merchantModify); err != nil {
		return merchantModify, nil
	}
	return
}
