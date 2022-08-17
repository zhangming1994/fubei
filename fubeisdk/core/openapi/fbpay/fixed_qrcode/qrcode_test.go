package fixedqrcode

import (
	"testing"

	"git.myarena7.com/arena/fubeisdk/core/config"
	"git.myarena7.com/arena/fubeisdk/util/snowflake"
)

func orderConfigInit() *QrcodePay {
	snowflake.Init(1)
	return NewQrcodePay(config.Config{
		URL:       "",
		AppSecret: "",
		VendorSN:  "",
	})
}

func TestQrcodePayConfig(t *testing.T) {
	qrcodeConfig := orderConfigInit()
	if resp, err := qrcodeConfig.QrCodeConfigQuery(1790578); err != nil {
		t.Fatalf("QrcodeConfig Query Error:%v", err)
	} else {
		t.Log(resp, "QrcodeConfig Query Success")
	}
}

func TestQrcodeConfig(t *testing.T) {
	qrcodeConfig := orderConfigInit()
	var qrcode QrcodePayConfig
	qrcode.MerchantID = 1790578
	qrcode.PayUrl = ""
	qrcode.SubAppID = ""
	qrcode.AlipayAppID = ""
	qrcode.AlipayPrivateKey = ""
	qrcode.AlipayPublicKey = ""
	if resp, err := qrcodeConfig.QrCodeConfig(qrcode, 1); err != nil {
		t.Fatalf("QrcodeConfig Query Error:%v", err)
	} else {
		t.Log(resp, "QrcodeConfig Query Success")
	}
}
