package merchant

import (
	"fmt"
	"testing"

	"git.myarena7.com/arena/fubeisdk/core/config"
	"git.myarena7.com/arena/fubeisdk/util/snowflake"
)

func wechatInit() *Wechat {
	snowflake.Init(1)
	return NewWechat(config.Config{
		URL:       "",
		AppSecret: "",
		VendorSN:  "",
	})
}

func TestWxConfig(t *testing.T) {
	wxConfig := wechatInit()
	var orderwxConfig WechatConfigRequest
	orderwxConfig.MerchantID = 0
	orderwxConfig.StoreID = 0
	orderwxConfig.SubAppID = ""
	if resp, err := wxConfig.WxConfig(orderwxConfig); err != nil {
		t.Fatalf(" WxConfig Error:%v", err)
	} else {
		t.Log(resp, " WxConfig Success")
	}
}

func TestWxConfigQuery(t *testing.T) {
	wxConfig := wechatInit()
	var orderwxConfig WxConfigQueryRequest
	orderwxConfig.MerchantID = 0
	orderwxConfig.StoreID = 0
	if resp, err := wxConfig.WxConfigQuery(orderwxConfig); err != nil {
		t.Fatalf(" WxConfig Error:%v", err)
	} else {
		t.Log(resp, " WxConfig Success")
	}
}

func TestMerchantWechatAuth(t *testing.T) {
	wechatConfig := wechatInit()
	var (
		wechatAuth  WechatAuthApplyRequest
		contactInfo ContactInfo
	)
	wechatAuth.MerchantID = 0
	wechatAuth.StoreID = 0
	wechatAuth.BusinessCode = fmt.Sprintf("%d", snowflake.NextID())
	wechatAuth.IdentificationAddress = ""
	contactInfo.ContactType = 1
	contactInfo.Name = ""
	contactInfo.Mobile = ""
	contactInfo.IDCardNumber = ""
	if resp, err := wechatConfig.MerchantWechatAuth(wechatAuth, &contactInfo, 1); err != nil {
		t.Fatalf("wechatAuth Apply Error:%v", err)
	} else {
		t.Log(resp, "wechatAuth Apply Success")
	}
}

func TestWechatApplyCancel(t *testing.T) {
	wechatConfig := wechatInit()
	var requestWechat WechatApplyCancelReq
	requestWechat.MerchantId = 0
	requestWechat.ApplymentId = ""
	if resp, err := wechatConfig.WechatApplyCancel(requestWechat); err != nil {
		t.Fatalf("wechatAuth Apply Cancel Result Error:%v", err)
	} else {
		t.Log(resp, "wechatAuth Apply Cancel Result Success")
	}
}

func TestMerchantWechatAuthResult(t *testing.T) {
	wechatConfig := wechatInit()
	var requestWechat WechatAuthApplyQueryRequest
	requestWechat.CallLevel = 1
	requestWechat.MerchantID = 0
	requestWechat.ApplymentID = ""
	if resp, err := wechatConfig.MerchantWechatAuthResult(requestWechat); err != nil {
		t.Fatalf("wechatAuth Apply Result Error:%v", err)
	} else {
		t.Log(resp, "wechatAuth Apply Result Success")
	}
}

func TestPaymentAuth(t *testing.T) {
	wechatConfig := wechatInit()
	var payment WechatPaymentAuthRequest
	payment.Url = "http.baidu.com"
	payment.MerchantID = 0
	payment.StoreID = 0
	payment.CallLevel = 1
	if resp, err := wechatConfig.PaymentAuth(payment); err != nil {
		t.Fatalf("wechatAuth  Error:%v", err)
	} else {
		t.Log(resp, "wechatAuth  Success")
	}
}
