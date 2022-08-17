package account

import (
	"encoding/json"
	"fmt"

	"git.myarena7.com/arena/fubeisdk/core/config"
	"git.myarena7.com/arena/fubeisdk/util"
)

type Account struct {
	config config.Config
}

func NewAccount(cfg config.Config) *Account {
	return &Account{config: cfg}
}

// OrderBillDown 对账单下载请求参数
type OrderBillDown struct {
	BillUrl  string `json:"bill_url"`  // 账单下载地址
	BillDate string `json:"bill_date"` // 账单日期，格式为yyyy-MM-dd
}

// AccountBalance 账户权限操作
type AccountBalance struct {
	CallLevel          int8    // 1. 代理 2. 商户
	QuickWithdraw      int8    `json:"quick_withdraw"`       // 是否有快提的权限 1 有 2 没有
	MerchantID         int64   `json:"merchant_id"`          // 付呗商户号，与account_id二选一，以商户级接入时无需填写
	PettyCash          float64 `json:"petty_cash"`           // 头寸退款备用金（元） 不能小于0，当头寸退款开启时
	SettleModeType     string  `json:"settle_mode_type"`     // 0：关闭1：开启3：查询（默认，传入该值则只查询状态，不进行操作）余额权限将影响以下业务权限的有效性。若余额权限关闭，则原本为开启状态的以下业务权限都将失效，并且传入权限也将无效
	WithDrawType       string  `json:"withdraw_type"`        // 提现方式01:提现方式02:手动（默认）余额权限=开启时有效。决定商户交易结算款是自动到账还是需要手动提现
	PositionRefundType string  `json:"position_refund_type"` // 头寸退款状态01:开启 02:关闭（默认）余额权限=开启时有效。商户头寸退款权限开关状态
}

// AccountWithdrawRequest 提现接口请求参数
type AccountWithdrawRequest struct {
	CallLevel       int8    //  1. 代理商 2. 商户
	Type            int8    `json:"type"`              // 提现方式：1 已结算金额提现，默认值为1默认提现方式一天有3个打款窗口期2 D0快速提现每日10次
	MerchantID      int64   `json:"merchant_id"`       // 分账接收方id，与merchant_id二选一，以商户级接入时，需要使用该字段进行查询时传入
	Amount          float64 `json:"amount"`            // 提现金额（元），精确到0.01，范围：0.01 ~ 10000000 type 为2秒到提现情况下，最低提现金额为1元
	AccountID       string  `json:"account_id"`        // 分账接收方id，与merchant_id二选一，以商户级接入时，需要使用该字段进行查询时传入
	MerchantOrderSN string  `json:"merchant_order_sn"` // 外部系统订单号（请求流水号
	NotifyUrl       string  `json:"notify_url"`        // 回调地址
}

// AccountWithdrawReponse 提现接口返回参数
type AccountWithdrawReponse struct {
	Status          int8    `json:"status"`            // 提现状态：2 提现中
	MerchantID      int64   `json:"merchant_id"`       // 付呗商户号，查询时传入则返回
	Amount          float64 `json:"amount"`            // 提现金额（元），精确到0.01
	AccountID       string  `json:"account_id"`        // 分账接收方id，查询时传入则返回
	MerchantOrderSN string  `json:"merchant_order_sn"` // 外部系统订单号（请求流水号
	CreateTime      string  `json:"create_time"`       // 申请时间，格式为yyyyMMddHHmmss
	FinishTime      string  `json:"finish_time"`       // 预计到账时间，格式为yyyyMMddHHmmss
	BankName        string  `json:"bank_name"`         // 银行名称
	BankCardNO      string  `json:"bank_card_no"`      // 银行卡号
}

// AccountWithdrawDetailsRequest 提现详情查询
type AccountWithdrawDetailsRequest struct {
	// 备注：服务商接入时merchant_id与account_id不能同时传，否则报错
	CallLevel       int8   // 1. 代理商 2. 服务商
	MerchantID      int64  // 付呗商户号，与account_id二选一，以商户级接入时无需填写
	AccountID       string // 分账接收方id，与merchant_id二选一，以商户级接入时，需要使用该字段进行查询时传入
	WithdrawNO      string // 提现订单号，与外部系统订单号二选一
	MerchantOrderSN string // 外部系统订单号（请求流水号），与提现订单号二选一
}

//  AccountWithdrawDetailsResponse 提现详情返回参数
type AccountWithdrawDetailsResponse struct {
	MerchantID      int64      `json:"merchant_id"`       // 付呗商户号，查询时传入则返回
	AccountID       string     `json:"account_id"`        // 分账接收方id，查询时传入则返回
	MerchantOrderSN string     `json:"merchant_order_sn"` // 外部系统订单号
	DataList        []DataList `json:"data_list"`         // -	提现记录列表
}

type DataList struct {
	Status      int8    `json:"status"`       // 提现状态：1提现成功，2银行处理中，3提现失败
	Type        int8    `json:"type"`         // 提现类型：1 已结算金额提现
	Amount      float64 `json:"amount"`       // 提现金额（元），精确到0.01
	Fee         float64 `json:"fee"`          // 提现手续费（元），精确到0.01，默认为0
	WithdrawNO  string  `json:"withdraw_no"`  // 提现订单号
	BankCardNO  string  `json:"bank_card_no"` // 银行卡号
	BankName    string  `json:"bank_name"`    // 银行名称
	FailMessage string  `json:"fail_message"` // 错误信息
	CreateTime  string  `json:"create_time"`  // 提现申请时间，格式为yyyyMMddHHmmss
	FinishTime  string  `json:"finish_time"`  // 提现到账时间，格式为yyyyMMddHHmmss
}

// AccountWithdrawRecord 提现记录查询请求参数
type AccountWithdrawRecordRequest struct {
	// 备注：服务商接入时merchant_id与account_id不能同时传，否则报错。
	CallLevel  int8   // 1. 代理商 2. 商户
	Status     int8   `json:"status"`      // 状态：0全部（默认），1提现成功，2银行处理中，3到账失败
	Page       int8   `json:"page"`        // 分页页码，默认值为1
	PageLimit  int8   `json:"page_limit"`  // 每页数量，默认值为10，最大100
	MerchantID int64  `json:"merchant_id"` // 付呗商户号，与account_id二选一，以商户级接入时无需填
	AccountID  string `json:"account_id"`  // 分账接收方id，与merchant_id二选一，以商户级接入时，需要使用该字段进行查询时传入
	BeginTime  string `json:"begin_time"`  // 查询开始时间，格式为yyyyMMddHHmmss
	EndTime    string `json:"end_time"`    // 查询结束时间，格式yyyyMMddHHmmss

}

// AccountWithdrawRecordReponse 提现记录查询响应参数
type AccountWithdrawRecordReponse struct {
	DataList   []WithdrawList `json:"data_list"`   // 提现记录列表
	Total      int32          `json:"total"`       // 提现记录总条数
	MerchantID int64          `json:"merchant_id"` // 付呗商户号，查询时传入则返回
	AccountID  string         `json:"account_id"`  // 分账接收方id，查询时传入则返回
}

type WithdrawList struct {
	Status          int8    `json:"status"`            // 提现状态：1提现成功，2银行处理中，3提现失败
	Type            int8    `json:"type"`              // 提现类型：1 已结算金额提现
	Amount          float64 `json:"amount"`            // 提现金额（元），精确到0.01
	Fee             float64 `json:"fee"`               // 手续费（元），精确到0.01
	WithdrawNO      string  `json:"withdraw_no"`       // 提现订单号
	MerchantOrderSN string  `json:"merchant_order_sn"` // 外部系统订单号
	BankCardNO      string  `json:"bank_card_no"`      // 银行卡号
	BankName        string  `json:"bank_name"`         // 银行名称
	FailMessage     string  `json:"fail_message"`      // 错误信息
	CreateTime      string  `json:"create_time"`       // 提现申请时间，格式为yyyyMMddHHmmss
	FinishTime      string  `json:"finish_time"`       // 提现到账时间，格式为yyyyMMddHHmmss
}

// AccountBalanceQueryRequest 账户额度查询
type AccountBalanceQueryRequest struct {
	// 备注：服务商接入时merchant_id与account_id不能同时传，否则报错。
	CallLevel  int8   // 1. 代理 2. 商户
	MerchantID int64  `json:"merchant_id"` // 付呗商户号，与account_id二选一，以商户级接入时无需填写
	AccountID  string `json:"account_id"`  // 分账接收方id，与merchant_id二选一，以商户级接入时，需要使用该字段进行查询时传入
}

// AccountBalanceQueryReponse 账户额度查询返回参数
type AccountBalanceQueryReponse struct {
	MerchantID        int64   `json:"merchant_id"`         // 付呗商户号，查询时传入则返回
	ClearingAmount    float64 `json:"clearing_amount"`     // 今日交易金额（元），精确到0.01。查询分账接收方时该字段为0
	SettledAmount     float64 `json:"settled_amount"`      // 可提金额（元），精确到0.01
	FrozenAmount      float64 `json:"frozen_amount"`       // 冻结金额（元），精确到0.01（若存在风险交易、入账异常等情况时，会造成金额被冻结，请联系客服）查询分账接收方时该字段为0
	ShareFrozenAmount float64 `json:"share_frozen_amount"` // 分账冻结金额（元），精确到0.01。开启分账功能的商户，在结算时间会将清算金额转入到分账冻结金额，等待分账成功后进入结算金额。查询分账接收方时该字段为0
	TotalAmount       float64 `json:"total_amount"`        // 总余额（元），精确到0.01
	AccountID         string  `json:"account_id"`          // 分账接收方id，查询时传入则返回
}

// AccountFixedQuery  分账查询请求参数
type AccountFixedQuery struct {
	MerchantID      int64  //付呗商户号
	MerchantOrderSN string // 外部系统订单号（请求流水号）， 与分账订单号二选一
	OrderNO         string // 分账订单号，与外部系统订单号二选一
}

// AccountFixedQueryResponse 分账查询返回参数
type AccountFixedQueryResponse struct {
	MerchantID      int64                 `json:"merchant_id"`       // 付呗商户号
	MerchantOrderSN string                `json:"merchant_order_sn"` // 外部系统订单号
	Data            []AccountFixQueryInfo `json:"data_list"`         // 分账账单列表
}

type AccountFixQueryInfo struct {
	AccountType int8    `json:"account_type"` // 账户类型：1 分账接收方、2 商户
	AccountOut  int64   `json:"account_out"`  // 分账出账户，即付呗商户号
	ShareAmount float64 `json:"share_amount"` // 分账金额（元）
	OrderNO     string  `json:"order_no"`     // 分账订单号
	AccountIn   string  `json:"account_in"`   // 入账户id，当账户类型为1时表示分账接收方，当账户类型为2时表示付呗商户号
	ShareStatus string  `json:"share_status"` // 分账状态：NOSUBMIT 未分账，PROCESSING 分账中，SUCCESS 分账成功，FAIL 分账失败
	ShareTime   string  `json:"share_time"`   // 分账时间，格式为yyyyMMddHHmmss
	FailMessage string  `json:"fail_message"` // 失败原因
}

// AccountFixAmount  按照固定金额分账
type AccountFixAmount struct {
	CallLevel        int8                   // 1. 代理商 2. 商户
	MerchantID       int64                  `json:"merchant_id"`        // 付呗商户号，商户级别接入无需填写
	TotalShareAmount float64                `json:"total_share_amount"` // 分账总金额（元），精确到0.01，范围：0.01 ~ 1000000
	MerchantOrderSN  string                 `json:"merchant_order_sn"`  // 外部系统订单号（请求流水号）
	NotifyUrl        string                 `json:"notify_url"`         // 回调地址
	Data             []AccountFixAmountList `json:"data_list"`          // 分账列表
}

type AccountFixAmountList struct {
	ShareAmount float64 `json:"share_amount"` // 分账金额（元），精确到0.01，范围：0.01 ~ 1000000
	AccountID   string  `json:"account_id"`   // 分账接收方id
}

// AccountBalanceEntry 入账户资金汇总查询 merchant_id与account_id不能同时传，否则报错。
type AccountBalanceEntry struct {
	CallLevel  int8   // 1. 代理 2. 商户
	MerchantID int64  // 付呗商户号，与分账接收方id二选一，以商户级接入时无需填写
	AccountID  string // 分账接收方id，与付呗商户号二选一
	BeginTime  string // 查询开始时间，格式为yyyyMMdd
	EndTime    string // 查询结束时间，格式为yyyyMMdd
}

// AccountBalanceEntryResponse 入账户资金汇总查询返回参数
type AccountBalanceEntryResponse struct {
	AccountType int8   `json:"account_type"` // 账户类型：1 分账接收方，2 商户
	MerchantId  int64  `json:"merchant_id"`  // 付呗商户号
	AccountID   string `json:"account_id"`   // 分账入账户
	Data        []struct {
		AccountOut   int64   `json:"account_out"`  // 分账出账户，即付呗商户号
		Amount       float64 `json:"amount"`       // 分账金额（元），精确到0.01
		AccountName  string  `json:"account_name"` // 分账出账户名称
		OrderNO      string  `json:"order_no"`     // 分账订单号
		ShareStatus  string  `json:"share_status"` // 分账状态：NOSUBMIT 未分账，PROCESSING 分账中，SUCCESS 分账成功，FAIL 分账失败
		FeailMessage string  `json:"fail_message"` // 错误信息
		ShareType    string  `json:"share_type"`   // 分账类型： ABS-按金额分账;REL-按比例分账
		CreateTime   string  `json:"create_time"`  // 分账创建时间，格式为yyyyMMddHHmmss
		FinishTime   string  `json:"finish_time"`  // 分账完成时间，格式为yyyyMMddHHmmss
		WithDrawList []struct {
			Amount         float64 `json:"amount"`          // 提现金额（元），精确到0.01
			Fee            float64 `json:"fee"`             // 提现手续费（元），精确到0.01，默认为0
			WithDrawNO     string  `json:"withdraw_no"`     // 提现订单号
			CreateTime     string  `json:"create_time"`     // 提现创建时间，格式为yyyyMMddHHmmss
			FinishTime     string  `json:"finish_time"`     // 提现完成时间，格式为yyyyMMddHHmmss
			WithdrawStatus string  `json:"withdraw_status"` // 提现状态：NOSUBMIT 未提现，PROCESSING 提现中，SUCCESS 提现成功，FAIL 提现失败，SYSTEM_NOT_CASH 用户手动提现
			BankCardNO     string  `json:"bank_card_no"`    // 银行卡号
			BankName       string  `json:"bank_name"`       // 银行名称
		} `json:"withdraw_list"` // 提现账单列表
	} `json:"data_list"` // 分账账单列表
}

// OrderBillDown  对账单下载
/*
● 对账单中只生产成功的交易订单
● 我们每日凌晨生产前一天的对账单，建议13:00点后再获取
● 对账单接口只能下载180天以内的账单
● 本功能于2019-10-12 正式上线，故相关对账单日期只能在该上线日期之后才存在
● 如果前一日没有交易，仍旧会生成一份对账文件，只不过对账文件中没有相关的交易数据
● 需要找对接人员单独配置权限
*/
func (a *Account) OrderBillDown(date string) (bill OrderBillDown, err error) {
	if date == "" {
		return bill, fmt.Errorf("parameter error")
	}
	data, err := util.LanuchRquest(a.config, "openapi.agent.order.bill.download", map[string]string{
		"bill_date": date,
	})
	if err != nil {
		return bill, err
	}
	if err = json.Unmarshal(data, &bill); err != nil {
		return bill, err
	}
	return
}

// AccountBalanceSwitch  账户权限操作
func (a *Account) AccountBalanceSwitch(accountSwitch AccountBalance) (AccountBalance, error) {
	if err := util.Validate(accountSwitch); err != nil {
		return AccountBalance{}, err
	}
	if accountSwitch.MerchantID == 0 && accountSwitch.CallLevel == 1 {
		return AccountBalance{}, fmt.Errorf("parameter error")
	}

	if accountSwitch.PettyCash < 0 {
		return AccountBalance{}, fmt.Errorf("parameter error:petty_cash")
	}

	if accountSwitch.SettleModeType == "" {
		return AccountBalance{}, fmt.Errorf("parameter error:settle_mode_type")
	}

	param := make(map[string]interface{}, 0)
	param["merchant_id"] = accountSwitch.MerchantID
	param["settle_mode_type"] = accountSwitch.SettleModeType
	param["petty_cash"] = accountSwitch.PettyCash

	if accountSwitch.WithDrawType != "" {
		param["withdraw_type"] = accountSwitch.WithDrawType
	}
	if accountSwitch.PositionRefundType != "" {
		param["position_refund_type"] = accountSwitch.PositionRefundType
	}
	if accountSwitch.QuickWithdraw != 0 {
		param["quick_withdraw"] = accountSwitch.QuickWithdraw
	}
	data, err := util.LanuchRquest(a.config, "openapi.agent.account.balance.switch", param)
	if err != nil {
		return AccountBalance{}, err
	}
	var permission AccountBalance
	if err := json.Unmarshal(data, &permission); err != nil {
		return permission, err
	}
	return permission, err
}

// AccountWithdraw 提现
func (a *Account) AccountWithdraw(withdraw AccountWithdrawRequest) (AccountWithdrawReponse, error) {
	if err := util.Validate(withdraw); err != nil {
		return AccountWithdrawReponse{}, err
	}
	if withdraw.CallLevel == 1 && withdraw.MerchantID == 0 {
		return AccountWithdrawReponse{}, fmt.Errorf("parameter error")
	}

	if withdraw.CallLevel == 2 && withdraw.AccountID == "" {
		return AccountWithdrawReponse{}, fmt.Errorf("parameter error")
	}
	if withdraw.Amount < 1 && withdraw.Type == 2 {
		return AccountWithdrawReponse{}, fmt.Errorf("parameter error")
	}
	param := make(map[string]interface{}, 0)
	param["merchant_order_sn"] = withdraw.MerchantOrderSN
	param["amount"] = withdraw.Amount
	param["type"] = withdraw.Type
	if withdraw.NotifyUrl != "" {
		param["notify_url"] = withdraw.NotifyUrl
	}
	if withdraw.CallLevel == 1 {
		param["merchant_id"] = withdraw.MerchantID
	} else if withdraw.CallLevel == 2 {
		param["account_id"] = withdraw.AccountID
	} else {
		return AccountWithdrawReponse{}, fmt.Errorf("parameter error")
	}
	data, err := util.LanuchRquest(a.config, "openapi.agent.account.withdraw", param)
	if err != nil {
		return AccountWithdrawReponse{}, err
	}
	var accountWithdraw AccountWithdrawReponse
	if err := json.Unmarshal(data, &accountWithdraw); err != nil {
		return AccountWithdrawReponse{}, err
	}
	return accountWithdraw, err
}

// AccountWithdrawDetail 提现详情
func (a *Account) AccountWithdrawDetail(withdrawInfo AccountWithdrawDetailsRequest) (AccountWithdrawDetailsResponse, error) {
	if err := util.Validate(withdrawInfo); err != nil {
		return AccountWithdrawDetailsResponse{}, err
	}
	if withdrawInfo.CallLevel == 1 && withdrawInfo.MerchantID == 0 {
		return AccountWithdrawDetailsResponse{}, fmt.Errorf("parameter error")
	}
	if withdrawInfo.CallLevel == 2 && withdrawInfo.AccountID == "" {
		return AccountWithdrawDetailsResponse{}, fmt.Errorf("parameter error")
	}
	if withdrawInfo.WithdrawNO == "" && withdrawInfo.MerchantOrderSN == "" {
		return AccountWithdrawDetailsResponse{}, fmt.Errorf("parameter error")
	}
	param := make(map[string]interface{}, 0)
	if withdrawInfo.CallLevel == 1 {
		param["merchant_id"] = withdrawInfo.MerchantID
	} else if withdrawInfo.CallLevel == 2 {
		param["account_id"] = withdrawInfo.AccountID
	} else {
		return AccountWithdrawDetailsResponse{}, fmt.Errorf("parameter error")
	}

	if withdrawInfo.WithdrawNO != "" {
		param["withdraw_no"] = withdrawInfo.WithdrawNO
	}

	if withdrawInfo.MerchantOrderSN != "" {
		param["merchant_order_sn"] = withdrawInfo.MerchantOrderSN
	}

	data, err := util.LanuchRquest(a.config, "openapi.agent.account.withdraw.details", param)
	if err != nil {
		return AccountWithdrawDetailsResponse{}, err
	}
	var accountWithdrawDetail AccountWithdrawDetailsResponse
	if err := json.Unmarshal(data, &accountWithdrawDetail); err != nil {
		return AccountWithdrawDetailsResponse{}, err
	}
	return accountWithdrawDetail, err
}

// AccountWithdrawRecord 提现记录查询
func (a *Account) AccountWithdrawRecord(withdrawRecord AccountWithdrawRecordRequest) (AccountWithdrawRecordReponse, error) {
	if err := util.Validate(withdrawRecord); err != nil {
		return AccountWithdrawRecordReponse{}, err
	}
	if withdrawRecord.CallLevel == 1 && withdrawRecord.MerchantID == 0 {
		return AccountWithdrawRecordReponse{}, fmt.Errorf("parameter error")
	}
	if withdrawRecord.CallLevel == 2 && withdrawRecord.AccountID == "" {
		return AccountWithdrawRecordReponse{}, fmt.Errorf("parameter error")
	}
	param := make(map[string]interface{}, 0)
	if withdrawRecord.CallLevel == 1 {
		param["merchant_id"] = withdrawRecord.MerchantID
	} else if withdrawRecord.CallLevel == 2 {
		param["account_id"] = withdrawRecord.AccountID
	} else {
		return AccountWithdrawRecordReponse{}, fmt.Errorf("parameter error")
	}
	param["page"] = withdrawRecord.Page
	param["page_limit"] = withdrawRecord.PageLimit
	param["status"] = withdrawRecord.Status

	if withdrawRecord.BeginTime != "" {
		param["begin_time"] = withdrawRecord.BeginTime
	}

	if withdrawRecord.EndTime != "" {
		param["end_time"] = withdrawRecord.EndTime
	}

	data, err := util.LanuchRquest(a.config, "openapi.agent.account.withdraw.record", param)
	if err != nil {
		return AccountWithdrawRecordReponse{}, err
	}
	fmt.Println(string(data))
	var accountWithdrawRecord AccountWithdrawRecordReponse
	if err := json.Unmarshal(data, &accountWithdrawRecord); err != nil {
		return AccountWithdrawRecordReponse{}, err
	}
	return accountWithdrawRecord, err
}

// AccountBalanceQuery 账户额度查询
func (a *Account) AccountBalanceQuery(accountBalance AccountBalanceQueryRequest) (AccountBalanceQueryReponse, error) {
	if err := util.Validate(accountBalance); err != nil {
		return AccountBalanceQueryReponse{}, err
	}
	if accountBalance.CallLevel == 1 && accountBalance.MerchantID == 0 {
		return AccountBalanceQueryReponse{}, fmt.Errorf("parameter error")
	}
	if accountBalance.CallLevel == 2 && accountBalance.AccountID == "" {
		return AccountBalanceQueryReponse{}, fmt.Errorf("parameter error")
	}

	param := make(map[string]interface{}, 0)
	if accountBalance.CallLevel == 1 {
		param["merchant_id"] = accountBalance.MerchantID
	} else if accountBalance.CallLevel == 2 {
		param["account_id"] = accountBalance.AccountID
	} else {
		return AccountBalanceQueryReponse{}, fmt.Errorf("parameter error")
	}

	data, err := util.LanuchRquest(a.config, "openapi.agent.account.balance.query", param)
	if err != nil {
		return AccountBalanceQueryReponse{}, err
	}
	fmt.Println(string(data), "================")
	var balance AccountBalanceQueryReponse
	if err := json.Unmarshal(data, &balance); err != nil {
		return AccountBalanceQueryReponse{}, err
	}
	return balance, err
}

// AccountFixAmount 按固定金额分账
func (a *Account) AccountFixAmount(aFixAmount AccountFixAmount) (bool, error) {
	if err := util.Validate(aFixAmount); err != nil {
		return false, err
	}
	if aFixAmount.CallLevel == 1 && aFixAmount.MerchantID == 0 {
		return false, fmt.Errorf("parameter error")
	}
	param := make(map[string]interface{}, 0)
	param["total_share_amount"] = aFixAmount.TotalShareAmount
	param["merchant_order_sn"] = aFixAmount.MerchantOrderSN
	param["data_list"] = aFixAmount.Data
	if aFixAmount.CallLevel == 1 {
		param["merchant_id"] = aFixAmount.MerchantID
	}
	if aFixAmount.NotifyUrl != "" {
		param["notify_url"] = aFixAmount.NotifyUrl
	}
	data, err := util.LanuchRquest(a.config, "openapi.agent.account.fixed.amount", param)
	if err != nil {
		return false, err
	}
	if string(data) != "true" {
		return false, fmt.Errorf(string(data))
	}
	return true, nil
}

// AccountFixQuery 分账查询
func (a *Account) AccountFixQuery(fix AccountFixedQuery) (AccountFixedQueryResponse, error) {
	if fix.MerchantOrderSN == "" && fix.OrderNO == "" {
		return AccountFixedQueryResponse{}, fmt.Errorf("parameter error")
	}
	param := make(map[string]interface{}, 0)
	param["merchant_id"] = fix.MerchantID
	if fix.MerchantOrderSN != "" {
		param["merchant_order_sn"] = fix.MerchantOrderSN
	}

	if fix.OrderNO != "" {
		param["order_no"] = fix.OrderNO
	}
	data, err := util.LanuchRquest(a.config, "openapi.agent.account.fixed.query", param)
	if err != nil {
		return AccountFixedQueryResponse{}, err
	}
	var accountFix AccountFixedQueryResponse
	if err := json.Unmarshal(data, &accountFix); err != nil {
		return AccountFixedQueryResponse{}, err
	}
	return accountFix, err
}

// AccountBalanceEntry 入账户资金汇总查询
func (a *Account) AccountBalanceEntry(balance AccountBalanceEntry) (AccountBalanceEntryResponse, error) {
	if err := util.Validate(balance); err != nil {
		return AccountBalanceEntryResponse{}, err
	}
	if balance.CallLevel == 1 && balance.MerchantID == 0 {
		return AccountBalanceEntryResponse{}, fmt.Errorf("parameter error")
	}
	if balance.CallLevel == 2 && balance.AccountID == "" {
		return AccountBalanceEntryResponse{}, fmt.Errorf("parameter error")
	}
	param := make(map[string]interface{}, 0)
	if balance.CallLevel == 1 {
		param["merchant_id"] = balance.MerchantID
	} else if balance.CallLevel == 2 {
		param["account_id"] = balance.AccountID
	} else {
		return AccountBalanceEntryResponse{}, fmt.Errorf("parameter error")
	}
	if balance.BeginTime != "" {
		param["begin_time"] = balance.BeginTime
	}

	if balance.EndTime != "" {
		param["end_time"] = balance.EndTime
	}

	data, err := util.LanuchRquest(a.config, "openapi.agent.account.balance.entry", param)
	if err != nil {
		return AccountBalanceEntryResponse{}, err
	}
	var balanceEntry AccountBalanceEntryResponse
	if err := json.Unmarshal(data, &balanceEntry); err != nil {
		return AccountBalanceEntryResponse{}, err
	}
	return balanceEntry, err
}
