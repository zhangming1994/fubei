package sdk

import (
	"git.myarena7.com/arena/fubeisdk/core/config"
	"git.myarena7.com/arena/fubeisdk/core/openapi/agent/account"
	wx "git.myarena7.com/arena/fubeisdk/core/openapi/agent/merchant"
	"git.myarena7.com/arena/fubeisdk/core/openapi/base"
	"git.myarena7.com/arena/fubeisdk/core/openapi/callback"
	"git.myarena7.com/arena/fubeisdk/core/openapi/fbpay/order"
	"git.myarena7.com/arena/fubeisdk/core/openapi/merchant"
	"git.myarena7.com/arena/fubeisdk/core/openapi/payment"
	"git.myarena7.com/arena/fubeisdk/core/openapi/share"
	"git.myarena7.com/arena/fubeisdk/util/snowflake"
)

type FuBei struct {
	config config.Config
}

func NewFuBei(conf config.Config) *FuBei {
	snowflake.Init(1)
	return &FuBei{config: conf}
}

// GetBase 基础接口
func (f *FuBei) GetBase() *base.Base {
	return base.NewBase(f.config)
}

// GetMerchant 商户相关
func (f *FuBei) GetMerchant() *merchant.Merchant {
	return merchant.NewMerchant(f.config)
}

// GetAgentMerchant 审核
func (f *FuBei) GetAgentMerchant() *wx.AgentMerchant {
	return wx.NewAgentMerchant(f.config)
}

// GetSubAccount 分账相关
func (f *FuBei) GetSubAccount() *account.SubAccount {
	return account.NewSubAccount(f.config)
}

// GetAccount 分账提现
func (f *FuBei) GetAccount() *account.Account {
	return account.NewAccount(f.config)
}

// GetWechat 微信
func (f *FuBei) GetWechat() *wx.Wechat {
	return wx.NewWechat(f.config)
}

// GetAccountAudit 商户审核
func (f *FuBei) GetAccountAudit() *wx.AgentMerchant {
	return wx.NewAgentMerchant(f.config)
}

// GetOrder 订单
func (f *FuBei) GetOrder() *order.Order {
	return order.NewOrder(f.config)
}

// GetPayment 微信支付宝授权
func (f *FuBei) GetPayment() *payment.Payment {
	return payment.NewPayment(f.config)
}

// GetElectContract 分账电子协议
func (f *FuBei) GetElectContract() *share.ElectContract {
	return share.NewElectContract(f.config)
}

// CallBackConfig 回调配置
func (f *FuBei) CallBackConfig() *callback.CallBack {
	return callback.NewCallBack(f.config)
}
