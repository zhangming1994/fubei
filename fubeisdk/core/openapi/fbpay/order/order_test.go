package order

import (
	"fmt"
	"testing"

	"git.myarena7.com/arena/fubeisdk/core/config"
	"git.myarena7.com/arena/fubeisdk/util/snowflake"
)

func orderConfigInit() *Order {
	snowflake.Init(1)
	return NewOrder(config.Config{
		URL:       "",
		AppSecret: "",
		VendorSN:  "",
	})
}

func TestOrderPay(t *testing.T) {
	orderConfig := orderConfigInit()
	var orderPay OrderPayRequest
	orderPay.MerchantOrderSN = fmt.Sprintf("%d", snowflake.NextID()) // 外部系统订单号
	orderPay.MerchantID = 0
	orderPay.AuthCode = ""
	orderPay.StoreID = 0
	orderPay.TotalAmount = 0.1
	// TODO 活动部分暂时不做测试
	// 4006:老板已将门店隐藏，您无法进行支付
	// 60303:当前商户需补齐相关资料后，才可进行相应的支付交易，请商户联系对接的微信支付服务商
	if resp, err := orderConfig.OrderPay(orderPay); err != nil {
		t.Fatalf("Order Pay Error:%v", err)
	} else {
		t.Log(resp, "Order Pay Success")
	}
}

//
func TestOrderQuery(t *testing.T) {
	orderConfig := orderConfigInit()
	var orderQueryRequest OrderQueryReqeust
	orderQueryRequest.CallLevel = 1
	orderQueryRequest.MerchantID = 0
	orderQueryRequest.OrderSN = "0"
	orderQueryRequest.MerchantOrderSN = "0"
	if resp, err := orderConfig.OrderQuery(orderQueryRequest); err != nil {
		t.Fatalf("Order Query Error:%v", err)
	} else {
		t.Log(resp, "Order Query Success")
	}
	// {"merchant_order_sn":"772878244008431616","order_sn":"20220402145448212108","ins_order_sn":"1000721028322092","channel_order_sn":"4200001357202204026420967757","order_status":"SUCCESS","pay_type":"wxpay","total_amount":0.10,"net_amount":0.10,"buyer_pay_amount":0.10,"fee":0.00,"store_id":1206520,"user_id":"ogRDxv8zRxpfuYsXZ8CDG83Ar8lQ","finish_time":"20220402145501","device_no":null,"attach":"","payment_list":[{"type":"FUBEI_DISCOUNT","amount":0.00},{"type":"CHANNEL_DISCOUNT","amount":0.00},{"type":"CHANNEL_PRE","amount":0.00}],"alipay_extend_params":null,"sub_code":null,"cashier_id":0,"is_can_part_refund":1}
}

func TestRefund(t *testing.T) {
	fefundConfig := orderConfigInit()
	var orderRefund OrderRefundRequest
	orderRefund.MerchantID = 1790578
	orderRefund.CallLevel = 1
	orderRefund.OrderSN = "0"
	orderRefund.MerchantOrderSN = "0"
	orderRefund.MerchantRefundSN = fmt.Sprintf("%d", snowflake.NextID())
	orderRefund.RefundAmount = 0.1
	if resp, err := fefundConfig.Refund(orderRefund); err != nil {
		t.Fatalf("Order Refund Error:%v", err)
	} else {
		t.Log(resp, "Order Refund Success")
	}
}

func TestRefundQuery(t *testing.T) {
	fefundConfig := orderConfigInit()
	var refund OrderRefundQueryRequest
	refund.CallLevel = 1
	refund.MerchantID = 0
	refund.MerchantRefundSN = ""
	refund.RefundSN = ""
	if resp, err := fefundConfig.RefundQuery(refund); err != nil {
		t.Fatalf("Order Refund Query Error:%v", err)
	} else {
		t.Log(resp, "Order Refund Query Success")
	}
}

func TestOrderClose(t *testing.T) {
	orderClose := orderConfigInit()
	var closeOrder OrderCloseRequest
	closeOrder.CallLevel = 1
	closeOrder.MerchantID = 1790578
	closeOrder.MerchantOrderSN = ""
	if resp, err := orderClose.OrderClose(closeOrder); err != nil {
		t.Fatalf("Order Close Error:%v", err)
	} else {
		t.Log(resp, "Order Close Success")
	}
}

func TestOrderCreate(t *testing.T) {

	orderCreate := orderConfigInit()
	var wxcreateOrder OrderCreateRequest
	wxcreateOrder.MerchantID = 1790578
	wxcreateOrder.CallLevel = 1
	wxcreateOrder.MerchantOrderSN = fmt.Sprintf("%v", snowflake.NextID())
	wxcreateOrder.PayType = "wxpay"
	wxcreateOrder.TotalAmount = 1
	wxcreateOrder.StoreID = 0
	wxcreateOrder.UserID = "o8uJ6uCVd-"
	wxcreateOrder.SubAppID = ""
	if resp, err := orderCreate.OrderCreate(wxcreateOrder); err != nil {
		t.Fatalf("Order Create Error:%v", err)
	} else {
		t.Log(resp, "Order Create Success")
	}

}
