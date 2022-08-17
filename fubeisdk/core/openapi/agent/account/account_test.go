package account

import (
	"fmt"
	"testing"

	"git.myarena7.com/arena/fubeisdk/core/config"
	"git.myarena7.com/arena/fubeisdk/util/snowflake"
)

func channelInit() *Account {
	snowflake.Init(1)
	return NewAccount(config.Config{
		URL:       "",
		AppSecret: "",
		VendorSN:  "",
	})
}

func TestBillDown(t *testing.T) {
	bill := channelInit()
	// 12012:账单未生成
	if resp, err := bill.OrderBillDown("2022-04-02"); err != nil {
		t.Fatalf("Bill Down Address Get Success Error:%v", err)
	} else {
		t.Log(resp, "Bill Down Address Get Success")
	}
}

func TestAccountBalanceSwitch(t *testing.T) {
	subAccount := channelInit()
	var balanceSwitch AccountBalance
	balanceSwitch.CallLevel = 1
	balanceSwitch.MerchantID = 0
	balanceSwitch.SettleModeType = "1"
	balanceSwitch.WithDrawType = "02"
	balanceSwitch.PositionRefundType = "02"
	balanceSwitch.QuickWithdraw = 1
	// 2100:开通失败，提现立即到账权限已关闭
	if resp, err := subAccount.AccountBalanceSwitch(balanceSwitch); err != nil {
		t.Fatalf("Test AccountBalanceSwitch Error:%v", err)
	} else {
		t.Log(resp, "AccountBalanceSwitch Success")
	}
}

func TestAccountWithdraw(t *testing.T) {
	subAccount := channelInit()
	var accountWithdraw AccountWithdrawRequest
	accountWithdraw.CallLevel = 1
	accountWithdraw.MerchantID = 0
	sn := fmt.Sprintf("%v", snowflake.NextID())
	fmt.Println(sn, "=====")
	accountWithdraw.MerchantOrderSN = sn
	accountWithdraw.Amount = 1
	accountWithdraw.Type = 2
	// 3231:提现失败，当前不支持提现立即到账
	if resp, err := subAccount.AccountWithdraw(accountWithdraw); err != nil {
		t.Fatalf("Test AccountWithdraw Error:%v", err)
	} else {
		t.Log(resp, "AccountWithdraw Success")
	}
}

func TestAccountWithdrawDetail(t *testing.T) {
	subAccount := channelInit()
	var accountWithdrawDetail AccountWithdrawDetailsRequest
	accountWithdrawDetail.CallLevel = 1
	accountWithdrawDetail.MerchantID = 0
	accountWithdrawDetail.MerchantOrderSN = fmt.Sprintf("%v", snowflake.NextID())
	// 3066:请求流水号不存在
	if resp, err := subAccount.AccountWithdrawDetail(accountWithdrawDetail); err != nil {
		t.Fatalf("Test AccountWithdrawDetail Error:%v", err)
	} else {
		t.Log(resp, "AccountWithdrawDetail Success")
	}
}

func TestAccountBalanceQuery(t *testing.T) {
	subAccount := channelInit()
	var accountBalanceQuery AccountBalanceQueryRequest
	accountBalanceQuery.CallLevel = 1
	accountBalanceQuery.MerchantID = 0
	if resp, err := subAccount.AccountBalanceQuery(accountBalanceQuery); err != nil {
		t.Fatalf("Test AccountBalanceQuery Error:%v", err)
	} else {
		t.Log(resp, "AccountBalanceQuery Success")
	}
}

func TestAccountWithdrawRecord(t *testing.T) {
	subAccount := channelInit()
	var withdrawRecord AccountWithdrawRecordRequest
	withdrawRecord.CallLevel = 1
	withdrawRecord.MerchantID = 0
	withdrawRecord.Page = 1
	withdrawRecord.PageLimit = 50
	// {"merchant_id":1790578,"account_id":null,"total":0,"data_list":[]}
	if resp, err := subAccount.AccountWithdrawRecord(withdrawRecord); err != nil {
		t.Fatalf("Test withdrawRecord Error:%v", err)
	} else {
		t.Log(resp, "withdrawRecord Success")
	}
}

func TestAccountFixQuery(t *testing.T) {
	subAccount := channelInit()
	var accoutFix AccountFixedQuery
	accoutFix.MerchantID = 1790578
	accoutFix.MerchantOrderSN = fmt.Sprintf("%v", snowflake.NextID()) // 实际应该是按照固定金额分账产生的这个值 随便生成的查不到会返回null
	if resp, err := subAccount.AccountFixQuery(accoutFix); err != nil {
		t.Fatalf("Test AccountFixedQuery Error:%v", err)
	} else {
		t.Log(resp, "AccountFixedQuery Success")
	}
}

func TestAccountBalanceEntry(t *testing.T) {
	subAccount := channelInit()
	var accoutFix AccountBalanceEntry
	accoutFix.CallLevel = 1
	accoutFix.MerchantID = 0
	// {"merchant_id":1790578,"account_id":null,"account_type":2,"data_list":[]}
	if resp, err := subAccount.AccountBalanceEntry(accoutFix); err != nil {
		t.Fatalf("Test AccountBalanceEntry Error:%v", err)
	} else {
		t.Log(resp, "AccountBalanceEntry Success")
	}
}

func TestAccountFixAmount(t *testing.T) {
	subAccount := channelInit()
	var accoutFix AccountFixAmount
	accoutFix.CallLevel = 1
	accoutFix.MerchantID = 3
	sn := fmt.Sprintf("%v", snowflake.NextID())
	fmt.Println(sn, "=====")
	accoutFix.MerchantOrderSN = sn
	accoutFix.TotalShareAmount = 0.12
	accoutFix.Data = make([]AccountFixAmountList, 1)
	accoutFix.Data[0].AccountID = "3"
	accoutFix.Data[0].ShareAmount = 0.12
	if resp, err := subAccount.AccountFixAmount(accoutFix); err != nil {
		t.Fatalf("Test AccountFixAmount Error:%v", err)
	} else {
		t.Log(resp, "AccountFixAmount Success")
	}
}
