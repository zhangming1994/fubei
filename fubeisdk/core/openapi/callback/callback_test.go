package callback

import (
	"fmt"
	"git.myarena7.com/arena/fubeisdk/core/config"
	"git.myarena7.com/arena/fubeisdk/util/snowflake"
	"testing"
)

func newCallBack() *CallBack {
	snowflake.Init(1)
	return NewCallBack(config.Config{
		URL:       "",
		AppSecret: "",
		VendorSN:  "",
	})
}

//
func TestCallBackUrlConfig(t *testing.T) {
	call := newCallBack()
	resp, err := call.CallBackConfig(CallBackConfigRequest{
		SecondCallBackUrl:          "",
		RemitCallBackUrl:           "",
		RefundCallBackUrl:          "",
		WithdrawCallBackUrl:        "",
		AccountRegisterCallBackUrl: "",
		ShareCallBackUrl:           "",
		MerchantAuditCallBackUrl:   "",
		ShareAuditCallBackUrl:      "",
		PayThirdCallbackUrl:        "",
	})
	if err != nil {
		t.Logf("callback url config error:%v", err)
	}
	fmt.Println(resp, "callback config response")
}
