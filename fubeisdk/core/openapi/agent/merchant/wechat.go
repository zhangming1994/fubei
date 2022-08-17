package merchant

import (
	"encoding/json"
	"fmt"

	"git.myarena7.com/arena/fubeisdk/core/config"
	"git.myarena7.com/arena/fubeisdk/util"
)

type Wechat struct {
	config config.Config
}

func NewWechat(cfg config.Config) *Wechat {
	return &Wechat{config: cfg}
}

// WechatAuthApplyRequest 微信实名认证申请
type WechatAuthApplyRequest struct {
	MerchantID            int64  `json:"merchant_id"`
	StoreID               int64  `json:"store_id" validate:"required"`
	BusinessCode          string `json:"business_code" validate:"required"`
	IdentificationAddress string `json:"identification_address"` // 法人证件居住地址 企业类型必须填写
}

// ContactInfo 联系人信息
type ContactInfo struct {
	Name                 string `json:"name"`                     // 姓名
	Mobile               string `json:"mobile"`                   // 手机号码
	IDCardNumber         string `json:"id_card_number"`           // 身份证号
	ContactType          int64  `json:"contact_type"`             // 联系人类型 1. 经营者/法人 2.经办人
	ContactIdDocCopy     string `json:"contact_id_doc_copy"`      // 联系人证件正面照
	ContactIdDocCopyBack string `json:"contact_id_doc_copy_back"` // 联系人证件照反面
	// 证件有效期开始时间
	ContactPeriodBegin string `json:"contact_period_begin"`
	// 联系人经办人 上传证件有效期结束时间
	ContactPeriodEnd string `json:"contact_period_end"`
	// 联系人证件有效期是否长期有效 0 否 1 是
	ContactPeriodIsLong int64 `json:"contact_periodIs_long"`
	// 业务办理授权函
	//1、当联系人类型是经办人时，请上传业务办理授权函
	//2、请参照[示例图]打印业务办理授权函，全部信息需打印，不支持手写商户信息，并加盖公
	BusinessAuthorizationLetter string `json:"business_authorization_letter"`
}

// WechatAuthApplyResponse 微信实名认证返回信息
type WechatAuthApplyResponse struct {
	BusinessCode string `json:"business_code"` // 业务申请编号
	ApplymentID  string `json:"applyment_id"`  // 服务商申请编号回执
	Message      string `json:"message"`
}

// WechatAuthApplyQueryRequest  微信认证申请结果查询请求参数
type WechatAuthApplyQueryRequest struct {
	CallLevel    int8   `json:"call_level"`    //调用级别 1. 服务商 2. 商户
	MerchantID   int64  `json:"merchant_id"`   // 付呗商户号，以服务商级接入时必传，以商户级接入时不传
	BusinessCode string `json:"business_code"` // 与服务商申请编号回执二选一必填
	ApplymentID  string `json:"applyment_id"`  // 与申请单编号二选一必填
}

// MerchantWechatAuthQueryResponse 微信认证申请单审核结果查询
type MerchantWechatAuthQueryResponse struct {
	BusinessCode string `json:"business_code"`
	ApplymentID  string `json:"applyment_id"`
	/*
		申请单状态：
		APPLYMENT_STATE_WAITTING_FOR_AUDIT【审核中】，请耐心等待1~2个工作日，微信支付将会完成审核
		APPLYMENT_STATE_EDITTING：【编辑中】，可能提交申请发生了错误导致，可用同一个业务申请编号重新提交。
		APPLYMENT_STATE_WAITTING_FOR_CONFIRM_CONTACT：【待确认联系信息】，请扫描微信支付返回的小程序码确认联系信息(此过程可修改超级管理员手机号)。
		APPLYMENT_STATE_WAITTING_FOR_CONFIRM_LEGALPERSON：【待账户验证】，请扫描微信支付返回的小程序码在小程序端完成账户验证。
		APPLYMENT_STATE_PASSED：【审核通过】，请扫描微信支付返回的小程序码在小程序端完成授权流程。
		APPLYMENT_STATE_REJECTED：【审核驳回】，请按照驳回原因修改申请资料，并更换业务申请编码，重新提交申请。
		APPLYMENT_STATE_FREEZED：【已冻结】，可能是该主体已完成过入驻，请查看驳回原因，并通知驳回原因中指定的联系人扫描微信支付返回的小程序码在小程序端完成授权流程。
		APPLYMENT_STATE_CANCELED：【已作废】，表示申请单已被撤销，无需再对其进行操作。
	*/
	ApplymentState string `json:"applyment_state"`

	/*
		小程序码图片:
		1、当申请单状态为APPLYMENT_STATE_WAITTING_FOR_CONFIRM_CONTACT、APPLYMENT_STATE_WAITTING_FOR_CONFIRM_LEGALPERSON、APPLYMENT_STATE_PASSED、APPLYMENT_STATE_FREEZED时，会返回小程序码图片。
		2、使用base64解码该字段，可得到图片二进制数据。
		3、可用img标签直接加载该图片。示例如下：<img src="data:image/png;base64,iVBORw0KGgoAAAANSU=" style="display: block;">
		示例值: cGFnZXMvYXBwbHkvYXBpdzQvd2VsY29tZS93ZWxjb21lP2FwcGx5bWVudF9pZD14eHg=
	*/
	QrCodeData   string `json:"qrcode_data"`
	RejectParam  string `json:"reject_param"`  //  驳回参数:当申请单状态为“审核驳回”时，会返回该字段，标识被驳回的字段名。
	RejectReason string `json:"reject_reason"` //  驳回原因:当申请单状态为“审核驳回”时，会返回该字段，表示驳回原因。
}

// WechatPaymentAuthRequest 微信网页授权请求参数
type WechatPaymentAuthRequest struct {
	CallLevel  int8   `json:"call_level"`                   // 1. 服务商 2. 商户
	MerchantID int64  `json:"merchant_id"`                  // 付呗商户号，以服务商级接入时必传，以商户级接入时不传
	StoreID    int64  `json:"store_id" validate:"required"` // 门店ID,当存在多个门店时,此字段必填; 该参数部分特殊需求商户必传；2018年4月1日之后该参数全部商户必传.
	Url        string `json:"url" validate:"required"`      // 授权完跳转地址
}

// WechatPaymentAuthResponse 微信网页授权返回参数
type WechatPaymentAuthResponse struct {
	AuthUrl string `json:"authUrl"` // 授权链接
}

// WechatConfigRequest 微信参数配置请求参数
type WechatConfigRequest struct {
	CallLevel   int8   `json:"call_level"`                   // 1. 服务商 2. 商户
	MerchantID  int64  `json:"merchant_id"`                  // 付呗商户号，以服务商级接入时必传，以商户级接入时不传
	StoreID     int64  `json:"store_id" validate:"required"` // 付呗系统的门店id
	SubAppID    string `json:"sub_appid"`                    // 支付所使用的公众号appid， 支持使用小程序appid
	JsapiPath   string `json:"jsapi_path"`                   // 支付授权目录
	AccountType string `json:"account_type"`                 //sub_appid类型 00：公众号（默认） 01：小程序
}

// WxConfigQueryResponse 微信参数配置查询返回参数
type WxConfigQueryResponse struct {
	MerchantId    int64  `json:"merchant_id"`     // 付呗商户号
	StoreID       int64  `json:"store_id"`        // 付呗系统的门店id
	AppIDList     string `json:"appid_list"`      // 支付APPID和关注APPID，JsonArray格式
	JsapiPathList string `json:"jsapi_path_list"` // 支付授权目录列表,JsonArray格式
}

// WechatConfigResponse 微信参数配置返回参数
type WechatConfigResponse struct {
	JsapiCode    int8   `json:"jsapi_code"`     // 支付授权目录配置结果：1 成功、2 失败
	MerchantID   int64  `json:"merchant_id"`    // 付呗商户号
	StoreID      int64  `json:"store_id"`       // 付呗系统的门店id
	SubAppIDCode int64  `json:"sub_appid_code"` // 支付APPID配置结果：1 成功 2 失败
	SubAppIDMsg  string `json:"sub_appid_msg"`  // 支付APPID响应描述
	JsapiMsg     string `json:"jsapi_msg"`      // 支付授权目录响应描述
}

// WxConfigQueryRequest 微信参数配置查询请求参数
type WxConfigQueryRequest struct {
	MerchantID int64 `json:"merchant_id"`                  // 付呗商户号，以服务商级接入时必传，以商户级接入时不传
	StoreID    int64 `json:"store_id" validate:"required"` // 付呗系统的门店id
}

// WechatApplyCancelReq 微信认证申请单撤销请求参数
type WechatApplyCancelReq struct {
	MerchantId   int64  `json:"merchant_id" validate:"required"`
	BusinessCode string `json:"business_code"`
	ApplymentId  string `json:"applyment_id"` // 后面两个参数选一个
}

// WechatApplyCancelResp 微信认证申请单撤销返回参数
type WechatApplyCancelResp struct {
	Data string `json:"data"`
}

// MerchantWechatAuth 微信实名认证申请
func (w *Wechat) MerchantWechatAuth(wechatAuth WechatAuthApplyRequest, contactInfo *ContactInfo, callLevel int8) (wechatApply WechatAuthApplyResponse, err error) {
	if err := util.Validate(wechatAuth); err != nil {
		return wechatApply, err
	}
	if callLevel != 1 && callLevel != 2 {
		return wechatApply, fmt.Errorf("parameter error")
	}
	if callLevel == 1 && wechatAuth.MerchantID == 1 {
		return wechatApply, fmt.Errorf("服务商级别的调用 商户号必须传递")
	}
	param := make(map[string]interface{}, 0)
	param["store_id"] = wechatAuth.StoreID
	param["business_code"] = wechatAuth.BusinessCode

	if wechatAuth.IdentificationAddress != "" {
		param["identification_address"] = wechatAuth.IdentificationAddress
	}
	if callLevel == 1 {
		param["merchant_id"] = wechatAuth.MerchantID
	}

	if contactInfo.ContactType == 2 && contactInfo.ContactIdDocCopy == "" {
		return wechatApply, fmt.Errorf("param error contact_id_doc_copy")
	}
	if contactInfo.ContactType == 2 && contactInfo.ContactIdDocCopyBack == "" {
		return wechatApply, fmt.Errorf("param error contact_id_doc_copy_back")
	}
	if contactInfo.ContactType == 2 && contactInfo.ContactPeriodBegin == "" {
		return wechatApply, fmt.Errorf("param error contact_period_begin")
	}
	if contactInfo.ContactType == 2 && contactInfo.ContactPeriodEnd == "" {
		return wechatApply, fmt.Errorf("param error contact_period_end")
	}
	if contactInfo.ContactType == 2 && contactInfo.BusinessAuthorizationLetter == "" {
		return wechatApply, fmt.Errorf("param error business_authorization_letter")
	}

	if contactInfo != nil {
		contactInformation := make(map[string]interface{})
		contactInformation["name"] = contactInfo.Name
		contactInformation["mobile"] = contactInfo.Mobile
		contactInformation["id_card_number"] = contactInfo.IDCardNumber
		contactInformation["contact_type"] = contactInfo.ContactType
		if contactInfo.ContactIdDocCopy != "" {
			contactInformation["contact_id_doc_copy"] = contactInfo.ContactIdDocCopy
		}
		if contactInfo.ContactIdDocCopyBack != "" {
			contactInformation["contact_id_doc_copy_back"] = contactInfo.ContactIdDocCopyBack
		}
		if contactInfo.ContactPeriodBegin != "" {
			contactInformation["contact_period_begin"] = contactInfo.ContactPeriodBegin
		}
		if contactInfo.ContactPeriodEnd != "" {
			contactInformation["contact_period_end"] = contactInfo.ContactPeriodEnd
		}
		if contactInfo.ContactType == 2 {
			contactInformation["contact_periodIs_long"] = contactInfo.ContactPeriodIsLong
		}
		if contactInfo.BusinessAuthorizationLetter != "" {
			contactInformation["business_authorization_letter"] = contactInfo.BusinessAuthorizationLetter
		}
		param["contact_info"] = contactInformation
	}

	data, err := util.LanuchRquest(w.config, "openapi.agent.merchant.wechat.auth.apply", param)
	if err != nil {
		return wechatApply, err
	}
	if err = json.Unmarshal(data, &wechatApply); err != nil {
		return wechatApply, err
	}
	return
}

// MerchantWechatAuthResult 微信认证申请单审核结果查询
func (w *Wechat) MerchantWechatAuthResult(wechatAuth WechatAuthApplyQueryRequest) (wechatAuthQuery MerchantWechatAuthQueryResponse, err error) {
	if err := util.Validate(wechatAuth); err != nil {
		return wechatAuthQuery, err
	}
	if wechatAuth.CallLevel != 1 && wechatAuth.CallLevel != 2 {
		return wechatAuthQuery, fmt.Errorf("parameter error")
	}
	if wechatAuth.ApplymentID == "" && wechatAuth.BusinessCode == "" {
		return wechatAuthQuery, fmt.Errorf("parameter error")
	}
	if wechatAuth.CallLevel == 1 && wechatAuth.MerchantID == 0 {
		return wechatAuthQuery, fmt.Errorf("parameter error")
	}

	param := make(map[string]interface{}, 0)
	if wechatAuth.CallLevel == 1 {
		param["merchant_id"] = wechatAuth.MerchantID
	}
	if wechatAuth.BusinessCode != "" {
		param["business_code"] = wechatAuth.BusinessCode
	}
	if wechatAuth.ApplymentID != "" {
		param["applyment_id"] = wechatAuth.ApplymentID
	}

	data, err := util.LanuchRquest(w.config, "openapi.agent.merchant.wechat.auth.apply.query", param)
	if err != nil {
		return wechatAuthQuery, err
	}
	if err = json.Unmarshal(data, &wechatAuthQuery); err != nil {
		return wechatAuthQuery, err
	}
	return
}

// PaymentAuth 微信网页授权
func (w *Wechat) PaymentAuth(paymentAuth WechatPaymentAuthRequest) (string, error) {
	if err := util.Validate(paymentAuth); err != nil {
		return "", err
	}
	if paymentAuth.CallLevel == 1 && paymentAuth.MerchantID == 0 {
		return "", fmt.Errorf("parameter error")
	}
	param := make(map[string]interface{}, 0)
	param["url"] = paymentAuth.Url
	param["store_id"] = paymentAuth.StoreID
	if paymentAuth.CallLevel == 1 {
		param["merchant_id"] = paymentAuth.MerchantID
	}
	data, err := util.LanuchRquest(w.config, "openapi.agent.merchant.wechat.payment.auth", param)
	if err != nil {
		return "", err
	}
	var authUrl WechatPaymentAuthResponse
	if err = json.Unmarshal(data, &authUrl); err != nil {
		return "", err
	}
	return authUrl.AuthUrl, nil
}

// WxConfig  微信参数配置
func (o *Wechat) WxConfig(wxConfig WechatConfigRequest) (wechatConfig WechatConfigResponse, err error) {
	if wxConfig.CallLevel == 1 && wxConfig.MerchantID == 0 {
		return WechatConfigResponse{}, fmt.Errorf("parameter error")
	}
	if err := util.Validate(wxConfig); err != nil {
		return WechatConfigResponse{}, err
	}
	data, err := util.LanuchRquest(o.config, "fbpay.order.wxconfig", wxConfig)
	if err != nil {
		return WechatConfigResponse{}, err
	}
	if err := json.Unmarshal(data, &wechatConfig); err != nil {
		return WechatConfigResponse{}, err
	}
	return
}

// WxConfigQuery 微信参数配置查询
func (o *Wechat) WxConfigQuery(wxQuery WxConfigQueryRequest) (wxConfigQuery WxConfigQueryResponse, err error) {
	if err := util.Validate(wxQuery); err != nil {
		return wxConfigQuery, err
	}
	data, err := util.LanuchRquest(o.config, "fbpay.order.wxconfig.query", wxQuery)
	if err != nil {
		return wxConfigQuery, err
	}
	if err := json.Unmarshal(data, &wxConfigQuery); err != nil {
		return wxConfigQuery, err
	}
	return
}

// WechatApplyCancel 微信认证申请单撤销
func (o *Wechat) WechatApplyCancel(wxApplyCancel WechatApplyCancelReq) (wxApplyCancelResp WechatApplyCancelResp, err error) {
	if err := util.Validate(wxApplyCancel); err != nil {
		return wxApplyCancelResp, err
	}
	param := make(map[string]interface{}, 0)
	param["merchant_id"] = wxApplyCancel.MerchantId
	if wxApplyCancel.BusinessCode != "" {
		param["business_code"] = wxApplyCancel.BusinessCode
	}
	if wxApplyCancel.ApplymentId != "" {
		param["applyment_id"] = wxApplyCancel.ApplymentId
	}
	data, err := util.LanuchRquest(o.config, "openapi.agent.merchant.wechat.auth.apply.cancel", param)
	if err != nil {
		return wxApplyCancelResp, err
	}
	if err := json.Unmarshal(data, &wxApplyCancelResp); err != nil {
		return wxApplyCancelResp, err
	}
	fmt.Println(wxApplyCancelResp.Data, "====data")
	return
}
