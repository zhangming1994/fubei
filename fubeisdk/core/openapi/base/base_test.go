package base

import (
	"encoding/base64"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"testing"

	"git.myarena7.com/arena/fubeisdk/core/config"
	"git.myarena7.com/arena/fubeisdk/util/snowflake"
)

func newBase() *Base {
	snowflake.Init(1)
	return NewBase(config.Config{
		URL:       "",
		AppSecret: "",
		VendorSN:  "",
	})
}

func TestImgUpload(t *testing.T) {
	// cashier.jpeg 本地测试的时候随便弄一张进来
	file, err := os.Open("storefront.jpeg")
	if err != nil {
		t.Fatalf("Open File error:%v", err)
	}
	defer file.Close()

	b, err := ioutil.ReadAll(file)
	if err != nil {
		t.Fatalf("File Read error:%v", err)
	}
	base64Result := base64.StdEncoding.EncodeToString(b)
	base := newBase()
	if resourceID, err := base.ImgUpload(ImgLoadRequest{
		FileData: base64Result,
		BusType:  "store",
	}, http.DetectContentType(b)); err != nil {
		t.Fatalf("Img Upload error:%v", err)
	} else {
		if resourceID == "" {
			t.Fatal("Img ResourceID Get error")
		}
		fmt.Println(resourceID, "Img Upload Back ResourceID")
	}

}

func TestCategoryList(t *testing.T) {
	base := newBase()
	// 一级类目
	if resp, err := base.GetCategoryList(0); err == nil {
		if len(resp) == 0 {
			t.Fatalf("Get CateGoryList error:%v", err)
		}
	} else {
		t.Fatalf("Get CateGoryList error:%v", err)
	}
	// 二级类目
	// resp, err := base.GetCategoryList(1)
}

func TestGetBankCode(t *testing.T) {
	base := newBase()
	// 借记卡 {"bank_id":699,"bank_name":"农业银行","bank_code":"103100000026","message":"校验成功","legal_flag":true}
	if resp, err := base.GetBankCode("农业银行"); err == nil {
		fmt.Println(resp, "===")
	} else {
		t.Fatalf("Get BankCode error:%v", err)
	}
	// 信用卡
	// resp, err := base.GetBankCode("6225768630903621")
	// fmt.Println(resp, err, "信用卡")
}

func TestGetProvinceAndCityCode(t *testing.T) {
	base := newBase()
	// 省级别获取
	if resp, err := base.GetArena(0, ""); err == nil {
		if len(resp) == 0 {
			t.Fatalf("Get ProvinceCode error%v", err)
		}
	} else {
		t.Fatalf("Get ProvinceCode error:%v", err)
	}
	// 省下面的市区
	// resp, err := base.GetArena(0, "330000")
	// 市下面的区
	// resp, err := base.GetArena(1, "330100")
}

func TestGetBankList(t *testing.T) {
	base := newBase()
	if resp, err := base.GetBankCode("杭州联合银行"); err == nil {
		if len(resp) == 0 {
			t.Fatalf("Get ProvinceCode error%v", err)
		}
		fmt.Println(resp, "Get Success")
	} else {
		t.Fatalf("Get ProvinceCode error:%v", err)
	}
}

func TestGetBankArea(t *testing.T) {
	base := newBase()
	// 省级别
	if resp, err := base.GetBankArea(1, "331000"); err == nil {
		if len(resp) == 0 {
			t.Fatalf("Get BankArea error:%v", err)
		}
	} else {
		t.Fatalf("Get BankArea error:%v", err)
	}
	// 获取省下面的市
	// resp, err := base.GetBankArea(0, "3310")
	// 市下面的县区/市
	// resp, err := base.GetBankArea(1, "331000")
}

func TestGetBranchName(t *testing.T) {
	base := newBase()
	request := new(BranchBankRequest)
	request.BankName = ""
	request.CityCode = "3310"
	if resp, err := base.GetBranchBank(request); err == nil {
		if len(resp) == 0 {
			t.Fatalf("Get BranchName error:%v", err)
		}
	} else {
		t.Fatalf("Get BranchName error:%v", err)
	}
}
