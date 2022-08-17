package merchant

import (
	"testing"

	"git.myarena7.com/arena/fubeisdk/core/config"
	"git.myarena7.com/arena/fubeisdk/util/snowflake"
)

func merchantInit() *AgentMerchant {
	snowflake.Init(1)
	return NewAgentMerchant(config.Config{
		URL:       "",
		AppSecret: "",
		VendorSN:  "",
	})
}

func TestMerchantIncomeAuditResult(t *testing.T) {
	merchantConfig := merchantInit()
	if resp, err := merchantConfig.MerchantIncomeAuditResult(5); err != nil {
		t.Fatalf("Merchant Audit Error:%v", err)
	} else {
		if resp.AuditStatus == "PASSED" {
			t.Log("Merchant Audit Success")
		}
		t.Log(resp, "Merchant Audit Back")
	}
}

func TestMerchantStoreStatus(t *testing.T) {
	merchantConfig := merchantInit()
	if resp, err := merchantConfig.MerchantStoreStatus("5"); err != nil {
		t.Fatalf("Merchant Store Status Error:%v", err)
	} else {
		t.Log(resp, "Merchant Store Status")
	}
}

func TestMerchantStoreInfo(t *testing.T) {
	merchantConfig := merchantInit()
	if resp, err := merchantConfig.MerchantStoreInfo("3"); err != nil {
		t.Fatalf("Merchant Store Info Error:%v", err)
	} else {
		t.Log(resp, "Merchant Store Info")
	}
}

func TestMerchantFeeRateModify(t *testing.T) {
	adjustRate := merchantInit()
	var feeRate MerchantRateModifyRequest
	feeRate.MerchantID = 1832559
	feeRate.AlipayFeeRate = "7"
	if resp, err := adjustRate.MerchantFeeRateModify(feeRate); err != nil {
		t.Fatalf("Merchant FeeRate Modify Error:%v", err)
	} else {
		t.Log(resp, "Merchant FeeRate Modify")
	}
}

func TestMerchantBankUpdate(t *testing.T) {
	agentMerchant := merchantInit()
	var bankUpdate MerchantUpdateBankRequest
	bankUpdate.MerchantID = 1832559
	bankUpdate.MerchantCodeType = 2
	bankUpdate.BankCardNO = ""
	bankUpdate.BankCardImage = "PJs16C8+//y4FJABuY//Wzut9v0imklKTx8dxBgz8ekMRY9Hw=="
	bankUpdate.BankCellPhone = ""
	bankUpdate.BankName = ""
	bankUpdate.UnionPayCode = ""
	// 9996:【无效参数】参数名:bankName,原因:对公账户bankName必填 所以个人账号修改和对公账户修改要分开
	// merchantid 进件的时候是对公账户 现在变成个人账户修改也会报错  9996:【无效参数】参数名:bankName,原因:对公账户bankName必填
	if resp, err := agentMerchant.MerchantUpdateBank(bankUpdate); err != nil {
		t.Fatalf("Merchant Bank Update Error:%v", err)
	} else {
		t.Log(resp, "Merchant Bank Update Success")
	}
}

func TestMerchantUpdateBankStatus(t *testing.T) {
	agentMerchant := merchantInit()
	var bankUpdateStatus MerchantBankBindStatusQueryRequest
	bankUpdateStatus.MerchantID = 3
	bankUpdateStatus.MerchantCode = ""
	bankUpdateStatus.BankCardNO = ""
	bankUpdateStatus.BankCode = ""
	// {"merchant_code":"","merchant_id":,"bind_status":2,"reject_msg":""}
	if resp, err := agentMerchant.MerchantUpdateBankStatus(bankUpdateStatus); err != nil {
		t.Fatalf("Merchant Query Bank Update Status Error:%v", err)
	} else {
		t.Log(resp, "Merchant Query Bank Update Status Success")
	}
}
