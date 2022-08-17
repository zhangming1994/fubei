package payment

import (
	"testing"

	"git.myarena7.com/arena/fubeisdk/core/config"
	"git.myarena7.com/arena/fubeisdk/util/snowflake"
)

func paymentInit() *Payment {
	snowflake.Init(1)
	return NewPayment(config.Config{
		URL:       "",
		AppSecret: "",
		VendorSN:  "",
	})
}

func TestWxOpenIDPaymentQuery(t *testing.T) {
	payment := paymentInit()
	var paymentwx WxOpenIDQueryRequest
	paymentwx.MerchantID = 1790578
	paymentwx.StoreID = 1206520
	paymentwx.AuthCode = ""
	paymentwx.SubAppID = ""
	// 60303:sub_mch_id与sub_appid不匹配
	if resp, err := payment.WxOpenIDPaymentQuery(paymentwx); err != nil {
		t.Fatalf("Get WX OpenID Error:%v", err)
	} else {
		t.Log(resp, "Get WX OpenID Success")
	}
}
