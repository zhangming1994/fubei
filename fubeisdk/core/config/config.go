package config

type Config struct {
	// 请求地址 服务商/商户的secret 服务商/商户的SN
	URL, AppSecret, VendorSN string
}
