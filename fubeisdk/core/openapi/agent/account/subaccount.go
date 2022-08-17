package account

import (
	"encoding/json"
	"fmt"

	"git.myarena7.com/arena/fubeisdk/core/config"
	"git.myarena7.com/arena/fubeisdk/util"
)

type SubAccount struct {
	config config.Config
}

func NewSubAccount(cfg config.Config) *SubAccount {
	return &SubAccount{config: cfg}
}

// SubAccountIncome 分账接收方入驻请求参数
type SubAccountIncome struct {
	CallLevel                        int8   // 1. 代理商 2. 商户
	AccountType                      int8   `json:"account_type"`                          // 结算账户类型：1对公、2对私
	BusinessType                     int8   `json:"business_type"`                         // 客户类型：1企业，2个体
	MerchantID                       int64  `json:"merchant_id"`                           // 商户id（服务商调用必传）
	MerchantOrderSN                  string `json:"merchant_order_sn"`                     // 外部系统订单号（请求流水号）
	Account                          string `json:"account"`                               // 账户名称
	LicenseNO                        string `json:"license_no"`                            // 营业执照号，当客户类型为企业时必填
	LicenseName                      string `json:"license_name"`                          // 营业执照名称，当客户类型为企业时必填
	LicensePhoto                     string `json:"license_photo"`                         // 营业执照照片地址（请填入上传加密图片返回的值），当客户类型为企业时必填
	LicenseTimeBegin                 string `json:"license_time_begin"`                    // 营业执照有效期开始时间，格式为yyyyMMdd
	LicenseTimeEnd                   string `json:"license_time_end"`                      // 营业执照有效期结束时间，格式为yyyyMMdd。若永久则传"20991231"
	LegalPersonName                  string `json:"legal_person_name"`                     // 法人姓名。企业填营业执照上的法人姓名；个体工商户填持卡人姓名
	LegalPersonIdCardType            string `json:"legal_person_id_card_type"`             // 法人证件类型，目前默认传"IDCARD"（身份证）
	LegalPersonIdCarNO               string `json:"legal_person_id_card_no"`               // 法人证件号。企业填营业执照上的法人号码；个体工商户填持卡人证件号
	LegalPersonIdCardFrontPhoto      string `json:"legal_person_id_card_front_photo"`      // 法人证件正面照片（请填入上传加密图片返回的值）注：人像面
	LegalPersonIdCardBackPhoto       string `json:"legal_person_id_card_back_photo"`       // 法人证件反面照片（请填入上传加密图片返回的值）注：国徽面
	IdCardType                       string `json:"id_card_type"`                          // 结算人证件类型：IDCARD 身份证、LICENSE 营业执照
	IdCardNO                         string `json:"id_card_no"`                            // 结算人证件号
	SettlementPersonIdCardFrontPhoto string `json:"settlement_person_id_card_front_photo"` // 结算人证件正面照片（请填入上传加密图片返回的值）注：人像面
	SettlementPersonIdCardBackPhoto  string `json:"settlement_person_id_card_back_photo"`  // 结算人证件反面照片（请填入上传加密图片返回的值）注：国徽面
	AccountNO                        string `json:"account_no"`                            // 结算银行卡号
	BankCardPhoto                    string `json:"bank_card_photo"`                       // 结算银行卡照片（请填入上传加密图片返回的值
	AccountName                      string `json:"account_name"`                          // 结算银行卡开户名
	BankName                         string `json:"bank_name"`                             // 开户总行，如中国建设银行
	BranchName                       string `json:"branch_name"`                           // 开户支行
	BankNO                           string `json:"bank_no"`                               // 联行号，对公必填。通过“支行信息查询”获取
	BankCellPhone                    string `json:"bank_cell_phone"`                       // 预留手机号，用于提现失败通知等
	StoreFrontImgUrl                 string `json:"store_front_img_url"`                   // 门头照（请填入上传加密图片返回的值）
	StoreCashPhoto                   string `json:"store_cash_photo"`                      // 收银台照片（请填入上传加密图片返回的值）
	StoreEnvPhoto                    string `json:"store_env_photo"`                       // 店内环境照（请填入上传加密图片返回的值）
	ProvinceCode                     string `json:"province_code"`                         // 省code
	CityCode                         string `json:"city_code"`                             // 市code
	AreaCode                         string `json:"area_code"`                             // 区code
	StreetAddress                    string `json:"street_address"`                        // 详细地址
	UnityCategoryID                  int32  `json:"unity_category_id"`                     // 行业类目，参见附件"行业类目列表"中"类目编号"值，请勿使用一级类目编号
	ResidenceAddress                 string `json:"residence_address"`                     // 企业入驻填法人户籍地址，对私入驻填写结算人户籍地址
	LegalPersonLicStt                string `json:"legal_person_lic_stt"`                  // 法人身份证开始时间,如2019-01-30
	LegalPersonLicEnt                string `json:"legal_person_lic_ent"`                  // 法人身份证结束时间,如2059-01-30
	LegalPersonLicEffect             string `json:"legal_person_lic_effect"`               // 法人身份证是否长期有效 YES 是 NO 否
	SettlementPersonLicStt           string `json:"settlement_person_lic_stt"`             // 结算人身份证开始时间,如2019-01-30
	SettlementPersonLicEnt           string `json:"settlement_person_lic_ent"`             // 结算人身份证结束时间,如2059-01-30
	SettlementPersonLicEffect        string `json:"settlement_person_lic_effect"`          // 结算人身份证是否长期有效 YES 是 NO 否
	AuthorizationPic                 string `json:"authorization_pic"`                     // 非法人结算授权函 (法人和结算人不一致时必传)

}

// SubAccountIncomeResponse 分账接收方入驻返回参数
type SubAccountIncomeResponse struct {
	MerchantOrderSN string `json:"merchant_order_sn"` // 外部系统订单号
	AccountID       string `json:"account_id"`        // 分账接收方ID
}

// SubAccountUpdateRequest 分账接收方结算卡信息修改请求参数
type SubAccountUpdateRequest struct {
	MerchantOrderSN string // 外部系统订单号（请求流水号），与分账接收方id二选一
	AccountID       string // 分账接收方id，与外部系统订单号二选一
	AccountNO       string // 结算银行卡号
	BankCardPhoto   string // 结算银行卡照片
	BankName        string // 开户总行，如中国建设银行
	BranchName      string // 开户支行
	BankNO          string // 联行号，通过“支行信息查询”获取，对公必填
}

// SubAccountUpdateResponse 分账接收方结算卡信息修改返回参数
type SubAccountUpdateResponse struct {
	Status          int64  `json:"status"`            // 状态 1. 成功 2. 失败
	MerchantOrderSN string `json:"merchant_order_sn"` // 外部系统订单号（请求流水号）
	AccountID       string `json:"account_id"`        // 分账接收方id
	FailMessage     string `json:"fail_message"`      // 修改失败原因
}

// SubAccountReceiveQueryRequest  分账接收方查询请求参数
type SubAccountReceiveQueryRequest struct {
	MerchantOrderSN string // 外部系统订单号（请求流水号），与分账接收方id二选一
	AccountID       string // 分账接收方id，与外部系统订单号二选一
}

// SubAccountReceiveQueryReponse  分账接收方查询返回参数
type SubAccountReceiveQueryReponse struct {
	IncomeStatus          int8   `json:"income_status"`             // 入驻状态：1入驻成功 2入驻中 3入驻失败
	BusinessType          int8   `json:"business_type"`             // 客户类型：1企业、2个体
	AccountType           int8   `json:"account_type"`              // 结算账户类型：1对公、2对私
	UnityCategoryID       int64  `json:"unity_category_id"`         // 行业类目
	MerchantOrderSN       string `json:"merchant_order_sn"`         // 外部系统订单号（请求流水号）
	AccountID             string `json:"account_id"`                // 分账接收方id
	FailMessage           string `json:"fail_message"`              // 失败原因
	Account               string `json:"account"`                   // 客户名称
	LicenseNO             string `json:"license_no"`                // 营业执照号
	LicenseTimeBegin      string `json:"license_time_begin"`        // 营业执照有效期开始时间，格式为yyyyMMdd
	LicenseTimeEnd        string `json:"license_time_end"`          // 营业执照有效期结束时间，格式为yyyyMMdd
	LegalPersonName       string `json:"legal_person_name"`         // 法人姓名，企业填营业执照上的法人姓名；个体工商户填持卡人姓名
	LegalPersonIDCardType string `json:"legal_person_id_card_type"` // 法人证件类型：目前默认传"IDCARD"（身份证）
	LegalPersonIDCardNO   string `json:"legal_person_id_card_no"`   // 法人证件号，企业填营业执照上的法人号码；个体工商户填持卡人证件号
	IdCardType            string `json:"id_card_type"`              // 结算人证件类型：IDCARD身份证、LICENSE营业执照
	IdCardNO              string `json:"id_card_no"`                // 结算人证件号
	AccountNO             string `json:"account_no"`                // 结算银行卡号
	AccountName           string `json:"account_name"`              // 结算银行卡开户名
	BankName              string `json:"bank_name"`                 // 开户总行，如中国建设银行
	BranchName            string `json:"branch_name"`               // 开户支行
	BankNO                string `json:"bank_no"`                   // 联行号
	BankCellPhone         string `json:"bank_cell_phone"`           // 预留手机号，用于提现失败通知等

}

// SubAccountGroup 分账组
type SubAccountGroup struct {
	IsCreateAccount int8             `json:"is_create_account"` // 是否创建余额账户 0不创建 1创建 推荐传1 否者后续分账逻辑会报没有权限
	ShareType       int8             `json:"share_type"`        // 分账类型，1 按比例分账 2 按固定金额分账
	ShareMode       int8             `json:"share_mode"`        // 分账模式，1 按订单净收金额（扣除手续费）2 按订单实收金额 3 外部请求分账当分账类型为“按比例分账”时，只能传1和2；当分账类型为“按固定金额分账”，只能传3
	MerchantID      int64            `json:"merchant_id"`       // 付呗商户号(分账业务中表示分账发起方)，以商户级接入时无需填写
	MerchantOrderSN string           `json:"merchant_order_sn"` // 外部系统订单号（请求流水号）
	AgreementImg    string           `json:"agreement_img"`     // 协议图片，（请填入上传加密图片返回的值） ，多张图片以","隔开，最多支持20张图片，,协议需加盖公章后上传，而后纸质件寄到我司，待机构方加盖公章之后寄回给你方，一式两份
	Data            []SubAccountList `json:"data_list"`         // 分账组列表，具体字段见下方
}

type SubAccountList struct {
	AccountID string `json:"account_id"` // 分账接收方id
	RuleList  []struct {
		StoreID    int64   `json:"store_id"`    // 门店id
		ShareScale float64 `json:"share_scale"` // 分账比例，精确到0.00001
	} `json:"rule_list"` // 分账规则，当分账类型为“按比例分账”时必填
}

//SubAccountGroupCreateResponse 创建分账组返回参数
type SubAccountGroupCreateResponse struct {
	MerchantOrderSN string `json:"merchant_order_sn"` // 外部系统订单号
	Status          int8   `json:"status"`            // 状态，1 创建成功，2 创建失败，3 创建中
	GroupID         string `json:"group_id"`          // 分账组id，创建成功时返回
	FailMessage     string `json:"fail_message"`      // 失败原因
	Data            []struct {
		AccountID    string `json:"account_id"`    // 分账接收方id
		ActiveStatus int8   `json:"active_status"` // 生效状态：1 生效，2 不生效，3 暂停分账
		FailMessage  string `json:"fail_message"`  // 失败原因
	} `json:"data_list"` // 分账接收方列表
}

// SubAccountGroupMember 分账组添加分账成员
type SubAccountGroupMember struct {
	GroupID string           `json:"group_id"`
	Data    []SubAccountList `json:"data_list"`
}

// SubAccountGroupMemberResponse 分账组增加分账成员返回参数
type SubAccountGroupMemberResponse struct {
	GroupID string `json:"group_id"` // 分账组id
	Status  int8   `json:"status"`   // 状态：1添加成功， 2添加失败，3添加中
	Data    []struct {
		AccountID    string `json:"account_id"`    // 分账接收方id
		ActiveStatus int8   `json:"active_status"` // 生效状态：1生效，2不生效，3暂停分账
		FailMessage  string `json:"fail_message"`  // 失败原因
	} `json:"data_list"` // 分账接收方列表
}

// SubAccountGroupQuery 分账组查询请求参数
type SubAccountGroupQuery struct {
	CallLevel  int8   // 1 代理 2. 商户
	MerchantID int64  // 付呗商户号，与分账组id二选一，商户级别接入无需填写
	GroupID    string // 分账组id，与付呗商户号二选一
}

// SubAccountGroupQueryResponse 分账组查询返回参数
type SubAccountGroupQueryResponse struct {
	ShareType  int8                       `json:"share_type"`  // 分账类型：1 按比例分账 2 按固定金额分账
	ShareMode  int8                       `json:"share_mode"`  // 分账模式：1 按订单净收金额（扣除手续费）、2 按订单实收金额、3 外部系统分账
	MerchantID int64                      `json:"merchant_id"` // 付呗商户号
	GroupID    string                     `json:"group_id"`    // 分账组id
	Data       []SubAccountGroupQueryData `json:"data_list"`   // 分账组成员列表
}

type SubAccountGroupQueryData struct {
	ActiveStatus  int8   `json:"active_status"`   // 生效状态：1生效、2不生效、3暂停分账
	ShareMemberID string `json:"share_member_id"` // 分账组成员id
	AccountID     string `json:"account_id"`      // 分账接收方id
	AccountName   string `json:"account_name"`    // 分账组成员名称
	StartDate     string `json:"start_date"`      // 有效开始日期，格式为yyyyMMdd
	EndDate       string `json:"end_date"`        // 有效结束日期，格式为yyyyMMdd
	RuleList      []struct {
		StoreID    int64   `json:"store_id"`    // 门店id
		ShareScale float64 `json:"share_scale"` // 分账比例，精度为0.00001，比如0.5即为按50%分账
	} `json:"rule_list"` // 分账规则列表
}

// SubAccountIncome 分账接收方入驻
func (s *SubAccount) SubAccountIncome(subAccount SubAccountIncome) (SubAccountIncomeResponse, error) {
	if err := util.Validate(subAccount); err != nil {
		return SubAccountIncomeResponse{}, err
	}
	if subAccount.CallLevel == 1 && subAccount.MerchantID == 0 {
		return SubAccountIncomeResponse{}, fmt.Errorf("parameter error:merchant_id")
	}

	switch subAccount.BusinessType {
	case 1:
		if subAccount.LicenseNO == "" {
			return SubAccountIncomeResponse{}, fmt.Errorf("parameter error:license_no")
		}
		if subAccount.LicenseName == "" {
			return SubAccountIncomeResponse{}, fmt.Errorf("parameter error:license_name")
		}
		if subAccount.LicensePhoto == "" {
			return SubAccountIncomeResponse{}, fmt.Errorf("parameter error:license_photo")
		}
	}

	if subAccount.LegalPersonIdCarNO != subAccount.IdCardNO && subAccount.AuthorizationPic == "" {
		return SubAccountIncomeResponse{}, fmt.Errorf("parameter error:authorization_pic")
	}

	if !util.IDCardNOCheck(subAccount.LegalPersonIdCarNO) {
		return SubAccountIncomeResponse{}, fmt.Errorf("parameter error:legal_person_id_card_no")
	}

	if !util.IDCardNOCheck(subAccount.IdCardNO) {
		return SubAccountIncomeResponse{}, fmt.Errorf("parameter error:id_card_no")
	}

	if !util.BankNumberCheck(subAccount.AccountNO) {
		return SubAccountIncomeResponse{}, fmt.Errorf("parameter error:account_no")
	}

	param := make(map[string]interface{}, 0)
	param["merchant_order_sn"] = subAccount.MerchantOrderSN
	param["account"] = subAccount.Account
	param["business_type"] = subAccount.BusinessType
	param["legal_person_name"] = subAccount.LegalPersonName
	param["legal_person_id_card_type"] = subAccount.LegalPersonIdCardType
	param["legal_person_id_card_no"] = subAccount.LegalPersonIdCarNO
	param["legal_person_id_card_front_photo"] = subAccount.LegalPersonIdCardFrontPhoto
	param["legal_person_id_card_back_photo"] = subAccount.LegalPersonIdCardBackPhoto
	param["id_card_type"] = subAccount.IdCardType
	param["id_card_no"] = subAccount.IdCardNO
	param["settlement_person_id_card_front_photo"] = subAccount.SettlementPersonIdCardFrontPhoto
	param["settlement_person_id_card_back_photo"] = subAccount.SettlementPersonIdCardBackPhoto
	param["account_type"] = subAccount.AccountType
	param["account_no"] = subAccount.AccountNO
	param["bank_card_photo"] = subAccount.BankCardPhoto
	param["account_name"] = subAccount.AccountName
	param["bank_cell_phone"] = subAccount.BankCellPhone
	param["store_front_img_url"] = subAccount.StoreFrontImgUrl
	param["store_cash_photo"] = subAccount.StoreCashPhoto
	param["store_env_photo"] = subAccount.StoreEnvPhoto
	param["province_code"] = subAccount.ProvinceCode
	param["city_code"] = subAccount.CityCode
	param["area_code"] = subAccount.AreaCode
	param["street_address"] = subAccount.StreetAddress
	param["unity_category_id"] = subAccount.UnityCategoryID
	param["residence_address"] = subAccount.ResidenceAddress
	param["legal_person_lic_stt"] = subAccount.LegalPersonLicStt
	param["legal_person_lic_ent"] = subAccount.LegalPersonLicEnt
	param["legal_person_lic_effect"] = subAccount.LegalPersonLicEffect
	param["settlement_person_lic_stt"] = subAccount.SettlementPersonLicStt
	param["settlement_person_lic_ent"] = subAccount.SettlementPersonLicEnt
	param["settlement_person_lic_effect"] = subAccount.SettlementPersonLicEffect

	if subAccount.CallLevel == 1 {
		param["merchant_id"] = subAccount.MerchantID
	}

	if subAccount.AuthorizationPic != "" {
		param["authorization_pic"] = subAccount.AuthorizationPic
	}

	if subAccount.BankName != "" {
		param["bank_name"] = subAccount.BankName
	}

	if subAccount.BranchName != "" {
		param["branch_name"] = subAccount.BranchName
	}

	if subAccount.BankNO != "" {
		param["bank_no"] = subAccount.BankNO
	}

	if subAccount.LicenseTimeEnd != "" {
		param["license_time_end"] = subAccount.LicenseTimeEnd
	}

	if subAccount.LicenseTimeBegin != "" {
		param["license_time_begin"] = subAccount.LicenseTimeBegin
	}

	if subAccount.LicensePhoto != "" {
		param["license_photo"] = subAccount.LicensePhoto
	}

	if subAccount.LicenseName != "" {
		param["license_name"] = subAccount.LicenseName
	}

	if subAccount.LicenseNO != "" {
		param["license_no"] = subAccount.LicenseNO
	}

	data, err := util.LanuchRquest(s.config, "openapi.agent.account.subaccount.income", subAccount)
	if err != nil {
		return SubAccountIncomeResponse{}, err
	}
	var subAccountIncome SubAccountIncomeResponse
	if err = json.Unmarshal(data, &subAccountIncome); err != nil {
		return SubAccountIncomeResponse{}, err
	}
	return subAccountIncome, err
}

// SubAccountUpdate 分账接收方结算卡信息修改
func (s *SubAccount) SubAccountUpdate(accountUpdate SubAccountUpdateRequest) (SubAccountUpdateResponse, error) {
	if err := util.Validate(accountUpdate); err != nil {
		return SubAccountUpdateResponse{}, err
	}
	if accountUpdate.MerchantOrderSN == "" && accountUpdate.AccountID == "" {
		return SubAccountUpdateResponse{}, fmt.Errorf("parameter error")
	}
	param := make(map[string]interface{}, 0)
	param["account_no"] = accountUpdate.AccountNO
	param["bank_card_photo"] = accountUpdate.BankCardPhoto

	if accountUpdate.MerchantOrderSN != "" {
		param["merchant_order_sn"] = accountUpdate.MerchantOrderSN
	}

	if accountUpdate.AccountID != "" {
		param["account_id"] = accountUpdate.AccountID
	}

	if accountUpdate.BankName != "" {
		param["bank_name"] = accountUpdate.BankName
	}

	if accountUpdate.BranchName != "" {
		param["branch_name"] = accountUpdate.BranchName
	}

	if accountUpdate.BankNO != "" {
		param["bank_no"] = accountUpdate.BankNO
	}

	data, err := util.LanuchRquest(s.config, "openapi.agent.account.subaccount.update", param)
	if err != nil {
		return SubAccountUpdateResponse{}, err
	}
	var subAccountUpdate SubAccountUpdateResponse
	if err = json.Unmarshal(data, &subAccountUpdate); err != nil {
		return SubAccountUpdateResponse{}, err
	}
	return subAccountUpdate, err
}

// SubAccountQuery 分账接收方查询
func (s *SubAccount) SubAccountQuery(account SubAccountReceiveQueryRequest) (SubAccountReceiveQueryReponse, error) {
	if err := util.Validate(account); err != nil {
		return SubAccountReceiveQueryReponse{}, err
	}
	if account.MerchantOrderSN == "" && account.AccountID == "" {
		return SubAccountReceiveQueryReponse{}, fmt.Errorf("parameter error")
	}
	param := make(map[string]interface{}, 0)
	if account.MerchantOrderSN != "" {
		param["merchant_order_sn"] = account.MerchantOrderSN
	}
	if account.AccountID != "" {
		param["account_id"] = account.AccountID
	}
	data, err := util.LanuchRquest(s.config, "openapi.agent.account.subaccount.receive.query", param)
	if err != nil {
		return SubAccountReceiveQueryReponse{}, err
	}
	var subAccount SubAccountReceiveQueryReponse
	if err = json.Unmarshal(data, &subAccount); err != nil {
		return SubAccountReceiveQueryReponse{}, err
	}
	return subAccount, err
}

// SubAccounGroupCreate 创建分账组
func (s *SubAccount) SubAccounGroupCreate(group SubAccountGroup) (SubAccountGroupCreateResponse, error) {
	if err := util.Validate(group); err != nil {
		return SubAccountGroupCreateResponse{}, err
	}
	data, err := util.LanuchRquest(s.config, "openapi.agent.account.subaccount.group.create", group)
	if err != nil {
		return SubAccountGroupCreateResponse{}, err
	}
	var subAccountGroup SubAccountGroupCreateResponse
	if err = json.Unmarshal(data, &subAccountGroup); err != nil {
		return SubAccountGroupCreateResponse{}, err
	}
	return subAccountGroup, err
}

// SubAccounGroupAddMember  分账组增加分账成员
func (s *SubAccount) SubAccounGroupAddMember(member SubAccountGroupMember) (SubAccountGroupMemberResponse, error) {
	if err := util.Validate(member); err != nil {
		return SubAccountGroupMemberResponse{}, err
	}
	data, err := util.LanuchRquest(s.config, "openapi.agent.account.subaccount.addmember", member)
	if err != nil {
		return SubAccountGroupMemberResponse{}, err
	}
	var subAccountMember SubAccountGroupMemberResponse
	if err = json.Unmarshal(data, &subAccountMember); err != nil {
		return SubAccountGroupMemberResponse{}, err
	}
	return subAccountMember, err
}

// SubAccountGroupQuery 分账组查询
func (s *SubAccount) SubAccountGroupQuery(groupQuery SubAccountGroupQuery) (SubAccountGroupQueryResponse, error) {
	if err := util.Validate(groupQuery); err != nil {
		return SubAccountGroupQueryResponse{}, err
	}
	if groupQuery.CallLevel == 1 && groupQuery.MerchantID == 0 {
		return SubAccountGroupQueryResponse{}, fmt.Errorf("parameter error")
	}
	if groupQuery.MerchantID == 0 && groupQuery.GroupID == "" {
		return SubAccountGroupQueryResponse{}, fmt.Errorf("parameter error")
	}
	param := make(map[string]interface{}, 0)
	if groupQuery.CallLevel == 1 {
		param["merchant_id"] = groupQuery.MerchantID
	} else if groupQuery.CallLevel == 2 {
		param["group_id"] = groupQuery.GroupID
	}
	data, err := util.LanuchRquest(s.config, "openapi.agent.account.subaccount.query", param)
	if err != nil {
		return SubAccountGroupQueryResponse{}, err
	}
	fmt.Println(string(data), "========")
	var subAccountGroup SubAccountGroupQueryResponse
	if err = json.Unmarshal(data, &subAccountGroup); err != nil {
		return SubAccountGroupQueryResponse{}, err
	}
	return subAccountGroup, err
}
