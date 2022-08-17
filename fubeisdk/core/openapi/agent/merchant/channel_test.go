package merchant

import (
	"testing"

	"git.myarena7.com/arena/fubeisdk/core/config"
	"git.myarena7.com/arena/fubeisdk/util/snowflake"
)

func channelInit() *Channel {
	snowflake.Init(1)
	return NewChannel(config.Config{
		URL:       "",
		AppSecret: "",
		VendorSN:  "",
	})
}

func TestChannelQuery(t *testing.T) {
	channelConfig := channelInit()
	if resp, err := channelConfig.ChannelSubmchquery(2); err != nil {
		t.Fatalf("Merchant Channel Query Error:%v", err)
	} else {
		t.Log(resp, "Merchant Audit Back")
	}
}
