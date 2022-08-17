package merchant

import (
	"testing"

	"git.myarena7.com/arena/fubeisdk/core/config"
	"git.myarena7.com/arena/fubeisdk/util/snowflake"
)

func merchantInit() *Merchant {
	snowflake.Init(1)
	return NewMerchant(config.Config{
		URL:       "",
		AppSecret: "",
		VendorSN:  "",
	})
}
func TestMerchantIncome(t *testing.T) {
	var prepareMerchantIncome MerchantIncomeRequest

	prepareMerchantIncome.MerchantCode = "未来科技空间"

	// 商户基本信息
	prepareMerchantIncome.BaseInfo.MerchantType = 3       // 商户类型 1. 小微 2. 个体 3. 企业
	prepareMerchantIncome.BaseInfo.ContactPhone = ""      // 商户手机号码
	prepareMerchantIncome.BaseInfo.ServicePhone = ""      // 客服电话
	prepareMerchantIncome.BaseInfo.Email = ""             // 商户电子邮箱
	prepareMerchantIncome.BaseInfo.UnityCategoryID = 60   // 行业类型
	prepareMerchantIncome.BaseInfo.MerchantShortName = "" // 商户名称

	// 法人信息
	prepareMerchantIncome.LegalPerson.LegalName = ""                                                                                                     // 法人姓名
	prepareMerchantIncome.LegalPerson.LegalNum = ""                                                                                                      // 法人身份证号码
	prepareMerchantIncome.LegalPerson.LegalIdCardFrontPhoto = ""                                                                                         // 正面
	prepareMerchantIncome.LegalPerson.LegalIdCardBackPhoto = "PJs16C8+B7pGMnk7xsTAkVje4ucXCaR+BlaZ4rnR6jrPob4d3kTONfzYvhu2S9Hg7imklKTx8dxBgz8ekMRY9Hw==" // 反面
	prepareMerchantIncome.LegalPerson.HandHoldIdCardPhoto = ""                                                                                           // 手持照 小微企业必须上传 另外的类型可不传递
	prepareMerchantIncome.LegalPerson.LegalIdCardStart = "2015-07-15"                                                                                    // 开始日期
	prepareMerchantIncome.LegalPerson.LegalIdCardEnd = "2035-07-15"                                                                                      // 结束日期

	// 地址信息
	prepareMerchantIncome.AddressInfo.ProvinceCode = "330000"  // 省
	prepareMerchantIncome.AddressInfo.CityCode = "330100"      // 市
	prepareMerchantIncome.AddressInfo.AreaCode = "330105"      // 区
	prepareMerchantIncome.AddressInfo.StreetAddress = ""       // 详细地址
	prepareMerchantIncome.AddressInfo.Longitude = "120.215511" // 经度
	prepareMerchantIncome.AddressInfo.Latitude = "30.253082"   // 纬度

	// 营业资质信息
	prepareMerchantIncome.LicenseInfo.LicenseID = ""                                                                     // 营业执照注册号码
	prepareMerchantIncome.LicenseInfo.LicenseName = ""                                                                   // 营业执照名称
	prepareMerchantIncome.LicenseInfo.LicenseAddress = ""                                                                // 营业执照注册地址
	prepareMerchantIncome.LicenseInfo.LicenseTimeStart = "2016-08-04"                                                    // 营业执照注册号开始日期
	prepareMerchantIncome.LicenseInfo.LicenseTimeEnd = "2099-12-31"                                                      // 营业执照注册号结束日期
	prepareMerchantIncome.LicenseInfo.LicensePhoto = "PJs16C8+/+Oxdz0m7ikmuVIuFcS3hf5+FA5nAsT6zh9yQdOuwJFLxxyEXIjH3sw==" // 营业执照图片

	// 结算信息
	prepareMerchantIncome.AccountInfo.AccountType = 2                                                                            // 账户类型 1：个人账户 2：公司账户 商户类型为企业时，才能传为2：公司账户
	prepareMerchantIncome.AccountInfo.LegalFlag = 1                                                                              // 结算标志 0：非法人结算 1：法人结算 账户类型公司，则结算标志必须为1：法人结算 商户类型是小微，必须为1：法人结算
	prepareMerchantIncome.AccountInfo.UnionPayCode = ""                                                                          // 支行信息获取的unionpaycode
	prepareMerchantIncome.AccountInfo.RealName = ""                                                                              // 开户名 法人结算：与法人姓名姓名一致；企业账户：与营业执照注册名称一致
	prepareMerchantIncome.AccountInfo.IdCardNO = ""                                                                              // 结算人身份证号 个人账户必传，公司账户可不传，法人结算与法人身份证号一致
	prepareMerchantIncome.AccountInfo.IdCardFrontPhoto = "PJs16C8+//fbZtiu/EpelfEXMTJpLs7I1R5pWjy2SEtPOimklKTx8dxBgz8ekMRY9Hw==" // 非法人结算必填，法人结算时直接取法人身份证照片
	prepareMerchantIncome.AccountInfo.IdCardBackPhoto = "PJs16C8+/+=="                                                           // 非法人结算必填，法人结算可不填 法人结算时直接取法人身份证国徽面
	prepareMerchantIncome.AccountInfo.BankCellPhoto = ""                                                                         // 银行预留号码
	prepareMerchantIncome.AccountInfo.BankCardNo = ""                                                                            // 银行卡号
	prepareMerchantIncome.AccountInfo.BankCardPhoto = "PJs16C8+//=="                                                             // 银行卡正面
	prepareMerchantIncome.AccountInfo.UnincorporatedPhoto = ""                                                                   // 非法人结算授权书 非法人结算模式必传。

	// 门店信息
	prepareMerchantIncome.ShopInfo.StoreName = ""                                                   // 门店名称
	prepareMerchantIncome.ShopInfo.StorePhone = ""                                                  // 门店电话
	prepareMerchantIncome.ShopInfo.StoreEnvPhoto = "PJs16C8+//y4FJABuY/gLDP8cstAEPUwMjoWH9YpE24/==" // 经营场所内照片
	prepareMerchantIncome.ShopInfo.StoreFrontPhoto = "PJs16C8+//=="                                 // 门头照片
	prepareMerchantIncome.ShopInfo.StoreCashPhoto = "PJs16C8+B7pGMnk7xsTAkVje4ucXCaR/=="            // 收银台照片

	// 费率
	prepareMerchantIncome.RateInfo.AlipayFeeRate = "6" // 支付宝费率
	prepareMerchantIncome.RateInfo.WxFeeRate = "6"     // 微信费率
	prepareMerchantIncome.RateInfo.UnionFeeRate = "6"  // 银联费率
	// 下面三个字段值同时传递或者同时不传递
	prepareMerchantIncome.RateInfo.CreditCardFeeRate = ""   // 信用卡费率
	prepareMerchantIncome.RateInfo.DebitCardFeeRate = ""    // 借记卡费率
	prepareMerchantIncome.RateInfo.DebitCardMaxFeeRate = "" // 借记卡封顶手续费

	// 其他信息
	prepareMerchantIncome.OtherInfo.OperatintLicensePhoto = "PJs16C8+/++==" // 经营许可证图片
	prepareMerchantIncome.OtherInfo.Remark = "测试重复法人入驻"                     // 备注
	prepareMerchantIncome.OtherInfo.CallBackUrl = ""
	// 银行卡图片给空测试 9996:【无效参数】参数名:bank_card_photo,原因:银行卡正面图片地址不能为空,请核对参数后重新请求
	// 故意写错参数 9996:【无效参数】参数名:unity_category_id,原因:行业类目不能为空,请核对参数后重新请求
	merchantConfig := merchantInit()
	if resp, err := merchantConfig.MerchantIncome(prepareMerchantIncome); err != nil {
		t.Fatalf("Merchant Income Error:%v", err)
	} else {
		t.Log(resp, "Merchant Income Back")
	}
}

func TestMerchantIncomeStatus(t *testing.T) {
	merchantConfig := merchantInit()
	if resp, err := merchantConfig.MerchantIncomeStatus("柒号主场"); err != nil {
		t.Fatalf("Merchant Income Error:%v", err)
	} else {
		// 故意填错merchant_code 1004:商户不存在
		// merchant_id, merchant_code status
		// 1790578 13632450754 PASSED
		t.Log(resp, "Merchant Income Back")
	}
}

func TestMerchantModify(t *testing.T) {
	var merchantModify MerchantIncomeModifyRequest
	merchantModify.MerchantCode = ""
	merchantModify.BaseInfo.MerchantType = 3
	merchantModify.BaseInfo.UnityCategoryID = 60
	merchantModify.AccountInfo.LegalFlag = 1
	merchantModify.AccountInfo.AccountType = 1
	merchantConfig := merchantInit()
	// 传空 9996:merchant_id 与 merchant_code 同时为空
	// 仅仅传递merchant_code 其他全部空 9996:merchant_type 不允许修改
	// 已经通过的情况 修改merchant_type 修改为小微 小微进件修改 法人必须上传手持身份证照[本地抛错] 修改为个体：9996:merchant_type 不允许修改
	// 上面通过之后 行业类目不允许修改 然后报错 9996:非法人结算时结算人身份证人像面照片不能为空
	// 修改为法人结算之后 报错 9996:结算账户类型不支持，请更换结算账户类型
	// 修改结算账户类型为 公司账户之后报错 9996:开户许可证不能为空 这个开户许可证就是银行卡正面字段 bank_card_photo 修改为个人账户之后 9996:暂不支持当前银行，请更换其他银行
	if resp, err := merchantConfig.MerchantModify(merchantModify); err != nil {
		t.Fatalf("Merchant Modify Error:%v", err)
	} else {
		t.Log(resp, "Merchant Modify Back")
	}
}
