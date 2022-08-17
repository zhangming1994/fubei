package account

import (
	"fmt"
	"testing"

	"git.myarena7.com/arena/fubeisdk/core/config"
	"git.myarena7.com/arena/fubeisdk/util/snowflake"
)

func subAccountInit() *SubAccount {
	snowflake.Init(1)
	return NewSubAccount(config.Config{
		URL:       "https://shq-api.51fubei.com/gateway/agent",
		AppSecret: "c9ccff32335a575f03ff699d8b3fb5fc",
		VendorSN:  "2022020813573761631a",
	})
}

func TestSubAccountIncome(t *testing.T) {
	subConfig := subAccountInit()
	var subAccount SubAccountIncome
	subAccount.MerchantOrderSN = fmt.Sprintf("%v", snowflake.NextID())
	subAccount.Account = "测试收款方入驻"
	subAccount.BusinessType = 1
	subAccount.LicenseNO = "91330108MA27YCDM8M"
	subAccount.LicenseName = "杭州锐竞科技有限公司"
	subAccount.LicensePhoto = "PJs16C8+B7pGMnk7xsTAkVje4ucXCaR/4EqBwUwrW02LUONjA7tMlTdEjghZZuYlTqYYJVAfcxbs4TwWgdbFpfgaRKwboSfA+Oxdz0m7ikmuVIuFcS3hf5+FA5nAsT6zh9yQdOuwJFLxxyEXIjH3sw=="
	subAccount.LicenseTimeBegin = "20160804"
	subAccount.LicenseTimeEnd = "20991231"
	subAccount.LegalPersonName = "汪磊"
	subAccount.LegalPersonIdCardType = "IDCARD"
	subAccount.LegalPersonIdCarNO = "330127198301190011"
	subAccount.LegalPersonIdCardFrontPhoto = "PJs16C8+B7pGMnk7xsTAkVje4ucXCaR/4EqBwUwrW02LUONjA7tMlTdEjghZZuYlTqYYJVAfcxbs4TwWgdbFpdJSH946/fbZtiu/EpelfEXMTJpLs7I1R5pWjy2SEtPOimklKTx8dxBgz8ekMRY9Hw=="
	subAccount.LegalPersonIdCardBackPhoto = "PJs16C8+B7pGMnk7xsTAkVje4ucXCaR/4EqBwUwrW02LUONjA7tMlTdEjghZZuYlTqYYJVAfcxbs4TwWgdbFpZtXfy3EAr+BlaZ4rnR6jrPob4d3kTONfzYvhu2S9Hg7imklKTx8dxBgz8ekMRY9Hw=="
	subAccount.IdCardType = "IDCARD"
	subAccount.SettlementPersonIdCardFrontPhoto = "PJs16C8+B7pGMnk7xsTAkVje4ucXCaR/4EqBwUwrW02LUONjA7tMlTdEjghZZuYlTqYYJVAfcxbs4TwWgdbFpdJSH946/fbZtiu/EpelfEXMTJpLs7I1R5pWjy2SEtPOimklKTx8dxBgz8ekMRY9Hw=="
	subAccount.SettlementPersonIdCardBackPhoto = "PJs16C8+B7pGMnk7xsTAkVje4ucXCaR/4EqBwUwrW02LUONjA7tMlTdEjghZZuYlTqYYJVAfcxbs4TwWgdbFpZtXfy3EAr+BlaZ4rnR6jrPob4d3kTONfzYvhu2S9Hg7imklKTx8dxBgz8ekMRY9Hw=="
	subAccount.AccountType = 2
	subAccount.AccountNO = "201000198486461"
	subAccount.BankCardPhoto = "PJs16C8+B7pGMnk7xsTAkVje4ucXCaR/4EqBwUwrW02LUONjA7tMlTdEjghZZuYlTqYYJVAfcxbs4TwWgdbFpXuMaCHdERFbyULBR1CyN1pbDTmZFm9NjScyHNUWlwYPimklKTx8dxBgz8ekMRY9Hw=="
	subAccount.AccountName = "汪磊"
	subAccount.StoreCashPhoto = "PJs16C8+B7pGMnk7xsTAkVje4ucXCaR/4EqBwUwrW02LUONjA7tMlTdEjghZZuYlTqYYJVAfcxbs4TwWgdbFpXuMaCHdERFbyULBR1CyN1pbDTmZFm9NjScyHNUWlwYPimklKTx8dxBgz8ekMRY9Hw=="
	subAccount.StoreEnvPhoto = "PJs16C8+B7pGMnk7xsTAkVje4ucXCaR/4EqBwUwrW02LUONjA7tMlTdEjghZZuYlTqYYJVAfcxbs4TwWgdbFpa/y4FJABuY/gLDP8cstAEPUwMjoWH9YpE24/Wzut9v0imklKTx8dxBgz8ekMRY9Hw=="
	subAccount.ProvinceCode = "330000"
	subAccount.CityCode = "330100"
	subAccount.AreaCode = "330105"
	subAccount.StreetAddress = "浙江省杭州市拱墅区瓜山新苑柒号未来运动空间"
	subAccount.UnityCategoryID = 60
	subAccount.ResidenceAddress = "广州市天河区兴华路1号之一2802房"
	subAccount.LegalPersonLicStt = "2015-07-15"
	subAccount.LegalPersonLicEnt = "2035-07-15"
	subAccount.LegalPersonLicEffect = "NO"
	subAccount.SettlementPersonLicStt = "2015-07-15"
	subAccount.SettlementPersonLicEnt = "2035-07-15"
	subAccount.SettlementPersonLicEffect = "NO"
	subAccount.MerchantID = 1790578
	subAccount.BankCellPhone = "13632450754"
	subAccount.StoreFrontImgUrl = "PJs16C8+B7pGMnk7xsTAkVje4ucXCaR/4EqBwUwrW02LUONjA7tMlTdEjghZZuYlTqYYJVAfcxbs4TwWgdbFpfMpptdGX/v0ODkJnYe4CTYhjNXtEL7moLyZAXWcPWMrimklKTx8dxBgz8ekMRY9Hw=="
	subAccount.IdCardNO = "330127198301190011"
	subAccount.BankNO = "402331001026"
	// 9996:账户权限不支持分账
	if resp, err := subConfig.SubAccountIncome(subAccount); err != nil {
		t.Fatalf("SubAccountIncome Error:%v", err)
	} else {
		t.Log(resp, "SubAccountIncome Success")
	}
}

func TestSubAccountUpdate(t *testing.T) {
	subConfig := subAccountInit()
	var subAccountUpdate SubAccountUpdateRequest
	subAccountUpdate.MerchantOrderSN = fmt.Sprintf("%v", snowflake.NextID())
	subAccountUpdate.AccountNO = "6228480058492132772"
	subAccountUpdate.BankNO = "402331001026"
	// {"merchant_order_sn":null,"account_id":null,"status":2,"fail_message":"分账接收方不存在"}
	subAccountUpdate.BankCardPhoto = "PJs16C8+B7pGMnk7xsTAkVje4ucXCaR/4EqBwUwrW02LUONjA7tMlTdEjghZZuYlTqYYJVAfcxbs4TwWgdbFpfMpptdGX/v0ODkJnYe4CTYhjNXtEL7moLyZAXWcPWMrimklKTx8dxBgz8ekMRY9Hw=="
	if resp, err := subConfig.SubAccountUpdate(subAccountUpdate); err != nil {
		t.Fatalf("subAccountUpdate Error:%v", err)
	} else {
		t.Log(resp, "subAccountUpdate Success")
	}
}

func TestSubAccountQuery(t *testing.T) {
	subConfig := subAccountInit()
	var subAccountQuery SubAccountReceiveQueryRequest
	subAccountQuery.MerchantOrderSN = fmt.Sprintf("%v", snowflake.NextID())
	// 3070:分账接收方信息不存在
	if resp, err := subConfig.SubAccountQuery(subAccountQuery); err != nil {
		t.Fatalf("subAccountQuery Error:%v", err)
	} else {
		t.Log(resp, "subAccountQuery Success")
	}
}

func TestSubAccounGroupCreate(t *testing.T) {
	subConfig := subAccountInit()
	var subAccountGroup SubAccountGroup
	subAccountGroup.MerchantID = 1790578
	subAccountGroup.MerchantOrderSN = fmt.Sprintf("%v", snowflake.NextID())
	subAccountGroup.ShareType = 2
	subAccountGroup.ShareMode = 3
	subAccountGroup.AgreementImg = "PJs16C8+B7pGMnk7xsTAkVje4ucXCaR/4EqBwUwrW02LUONjA7tMlTdEjghZZuYlTqYYJVAfcxbs4TwWgdbFpfMpptdGX/v0ODkJnYe4CTYhjNXtEL7moLyZAXWcPWMrimklKTx8dxBgz8ekMRY9Hw=="
	subAccountGroup.IsCreateAccount = 1
	subAccountGroup.Data[0].AccountID = ""
	subAccountGroup.Data[0].RuleList[0].StoreID = 0
	subAccountGroup.Data[0].RuleList[0].ShareScale = 0.0003
	if resp, err := subConfig.SubAccounGroupCreate(subAccountGroup); err != nil {
		t.Fatalf("subAccountGroup Error:%v", err)
	} else {
		t.Log(resp, "subAccountGroup Success")
	}
}

func TestSubAccounGroupAddMember(t *testing.T) {
	subConfig := subAccountInit()
	var subAccountGroupMember SubAccountGroupMember
	subAccountGroupMember.GroupID = ""
	subAccountGroupMember.Data[0].AccountID = ""
	subAccountGroupMember.Data[0].RuleList[0].StoreID = 0
	subAccountGroupMember.Data[0].RuleList[0].ShareScale = 0.0003
	if resp, err := subConfig.SubAccounGroupAddMember(subAccountGroupMember); err != nil {
		t.Fatalf("SubAccountGroupMember Error:%v", err)
	} else {
		t.Log(resp, "SubAccountGroupMember Success")
	}
}

func TestSubAccountGroupQuery(t *testing.T) {
	subConfig := subAccountInit()
	var subAccountGroupquery SubAccountGroupQuery
	subAccountGroupquery.CallLevel = 1
	subAccountGroupquery.MerchantID = 1886461
	// 3058:该商户无对应的分账组
	if resp, err := subConfig.SubAccountGroupQuery(subAccountGroupquery); err != nil {
		t.Fatalf("subAccountGroupquery Error:%v", err)
	} else {
		t.Log(resp, "subAccountGroupquery Success")
	}
}
