package share

import (
	"fmt"
	"git.myarena7.com/arena/fubeisdk/core/config"
	"git.myarena7.com/arena/fubeisdk/util/snowflake"
	"testing"
)

func shareElectContractInit() *ElectContract {
	snowflake.Init(1)
	return NewElectContract(config.Config{
		URL:       "",
		AppSecret: "",
		VendorSN:  "",
	})
}

func TestGetShareElectContract(t *testing.T) {
	shareElect := shareElectContractInit()
	url, err := shareElect.GetShareElectContract(1790578)
	if err != nil {
		t.Fatalf("get elect contract url error:%v", err)
	}
	if url == "" {
		t.Fatalf("elect contract url is null")
	}
	// 3222:分账授权已提交
	fmt.Println(url, "result url")
}

func TestGetShareElectContractResult(t *testing.T) {
	shareElect := shareElectContractInit()
	data, err := shareElect.GetShareElectContractResult(1790578)
	if err != nil {
		t.Fatalf("get elect contract url result error:%v", err)
	}
	// {"shareType":1,"shareTypeStr":"已开通","contractUrl":""}
	fmt.Println(data, "result")
}
