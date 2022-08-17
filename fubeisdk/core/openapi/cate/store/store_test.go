package store

import (
	"testing"

	"git.myarena7.com/arena/fubeisdk/core/config"
	"git.myarena7.com/arena/fubeisdk/util/snowflake"
)

func storeInit() *Store {
	snowflake.Init(1)
	return NewStore(config.Config{
		URL:       "",
		AppSecret: "",
		VendorSN:  "",
	})
}
func TestStoreCreate(t *testing.T) {
	storeConfig := storeInit()
	var storeRequest StoreOperateInfo
	storeRequest.MerchantID = 0
	storeRequest.StoreName = ""
	storeRequest.StreetAddress = ""
	storeRequest.Latitude = "30.253082"
	storeRequest.Longitude = "120.21551"
	storeRequest.ProvinceCode = ""
	storeRequest.CityCode = ""
	storeRequest.AreaCode = "330105"
	storeRequest.StorePhone = ""
	storeRequest.UnityCategoryID = 60
	storeRequest.StoreFrontImgUrl = "PJs16C8+//v0ODkJnYe4CTYhjNXtEL7moLyZAXWcPWMrimklKTx8dxBgz8ekMRY9Hw=="
	storeRequest.StoreEnvPhoto = "PJs16C8+//y4FJABuY/gLDP8cstAEPUwMjoWH9YpE24/Wzut9v0imklKTx8dxBgz8ekMRY9Hw=="
	storeRequest.StoreCashPhoto = "PJs16C8+B7pGMnk7xsTAkVje4ucXCaR/=="
	storeRequest.LicenseType = 1
	storeRequest.LicensePhoto = "PJs16C8+/+Oxdz0m7ikmuVIuFcS3hf5+FA5nAsT6zh9yQdOuwJFLxxyEXIjH3sw=="
	storeRequest.LicenseID = ""
	storeRequest.LicenseTimeType = 2
	storeRequest.LicenseName = ""
	storeRequest.LicenseTimeBegin = "2016-08-04"
	// {"store_id":1235388,"store_status":2}
	if resp, err := storeConfig.StoreCreate(storeRequest); err != nil {
		t.Fatalf("Store Create Error:%v", err)
	} else {
		t.Log(resp, "Store Create Success")
	}
}

func TestStoreEdit(t *testing.T) {
	storeConfig := storeInit()
	var storeRequest StoreOperateInfo
	storeRequest.MerchantID = 1832559
	storeRequest.StoreID = 1235388
	storeRequest.StoreName = ""
	storeRequest.StreetAddress = ""
	storeRequest.Latitude = "30.253082"
	storeRequest.Longitude = "120.21551"
	storeRequest.ProvinceCode = ""
	storeRequest.CityCode = ""
	storeRequest.AreaCode = ""
	storeRequest.StorePhone = ""
	storeRequest.UnityCategoryID = 60
	storeRequest.StoreFrontImgUrl = "PJs16C8+B7pGMnk7xsTAkVje4ucXCaR//=="
	storeRequest.StoreEnvPhoto = "PJs16C8+B7pGMnk7xsTAkVje4ucXCaR//y4FJABuY//Wzut9v0imklKTx8dxBgz8ekMRY9Hw=="
	storeRequest.StoreCashPhoto = "PJs16C8+B7pGMnk7xsTAkVje4ucXCaR/=="
	storeRequest.LicenseType = 1
	storeRequest.LicensePhoto = "PJs16C8+/+Oxdz0m7ikmuVIuFcS3hf5+FA5nAsT6zh9yQdOuwJFLxxyEXIjH3sw=="
	storeRequest.LicenseID = ""
	storeRequest.LicenseTimeType = 2
	storeRequest.LicenseName = ""
	storeRequest.LicenseTimeBegin = "2016-08-04"
	if resp, err := storeConfig.StoreEdit(storeRequest); err != nil {
		t.Fatalf("Store Create Error:%v", err)
	} else {
		t.Log(resp, "Store Create Success")
	}
}
