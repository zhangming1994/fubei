// @Title  store
// @Description  门店创建 修改接口
package store

import (
	"encoding/json"
	"fmt"

	"git.myarena7.com/arena/fubeisdk/core/config"
	"git.myarena7.com/arena/fubeisdk/util"
)

type Store struct {
	config config.Config
}

func NewStore(cfg config.Config) *Store {
	return &Store{config: cfg}
}

// StoreOperateInfo 门店创建 修改相关参数信息
type StoreOperateInfo struct {
	LicenseType                  int8   `validate:"required"` // 门店经营许可类型：1 营业执照、2 证明函。企业账户类型，必须填营业执照。
	LicenseTimeType              int8   `validate:"required"` // 营业执照有效期类型：1 正常有效期、2 长期有效。若门店经营许可类型为"营业执照"则必填
	MerchantID                   int64  `validate:"required"` // 付呗商户号
	StoreID                      int64  // 门店ID 创建的时候不用给值 修改的时候必须给值
	UnityCategoryID              int64  `validate:"required"` // 行业类目。参见附件"行业类目列表"中"类目编号"值，请勿使用一级类目编号
	StoreName                    string `validate:"required"` // 门店名称
	UnincorporatedLegalPermanent string // 法人身份证有效期，是否永久有效：0：非长期 1：长期
	StreetAddress                string `validate:"required"` // 详细地址
	Longitude                    string `validate:"required"` // 门店地址经度，精确到小数点后6位
	Latitude                     string `validate:"required"` // 门店地址纬度，精确到小数点后6位
	ProvinceCode                 string `validate:"required"` // 省份编码
	CityCode                     string `validate:"required"` // 城市编码
	AreaCode                     string `validate:"required"` // 区域编码
	StorePhone                   string `validate:"required"` // 门店电话
	StoreFrontImgUrl             string `validate:"required"` // 门头照片（请填入上传加密图片返回的值）
	StoreEnvPhoto                string `validate:"required"` // 门店店内环境（请填入上传加密图片返回的值）
	StoreCashPhoto               string `validate:"required"` // 门店收银台（请填入上传加密图片返回的值）
	LicensePhoto                 string `validate:"required"` // 营业执照/证明函照片。若门店经营许可类型为"营业执照"则必填（请填入上传加密图片返回的值）
	HandHoldIDCardPic            string // 非必填 手持身份证照片。若门店经营许可类型为"证明函"则必填（请填入上传加密图片返回的值）
	LicenseName                  string // 营业执照名称。若门店经营许可类型为"营业执照"则必填
	LicenseID                    string // 营业执照号。若门店经营许可类型为"营业执照"则必填
	LicenseTimeBegin             string // 营业执照有效期开始时间，格式为YYYY-MM-DD。若门店经营许可类型为"营业执照"则必填，长期有效可不填
	LicenseTimeEnd               string // 营业执照有效期终止时间，格式为YYYY-MM-DD。若门店经营许可类型为"营业执照"则必填，长期有效可不填
	BrandName                    string // 非必填 品牌名称
	OperatingLicensePhoto        string // 非必填 经营许可证图片（请填入上传加密图片返回的值）
	LegalIDCardFrontPhoto        string // 非必填 法定代表人身份证人像面照（非法人结算时必填。上传营业执照上的法定代表人证件，上传加密图片的类型为idCard）
	LegalIDCardBackPhoto         string // 非必填 法定代表人身份证国徽面照（非法人结算时必填。上传营业执照上的法定代表人证件，上传加密图片的类型为idCard）
	UnincorporatedPhoto          string // 非必填 非法人结算证明（非法人结算时必填。请点击下载授权书，打印出来填写完整并签字后拍照上传，上传加密图片的类型为license
	UnincorporatedLegalName      string // 非必填 法人身份证姓名（非法人结算情况下必填）
	UnincorporatedLegalNum       string // 非必填 法人身份证号（非法人结算情况下必填）
	UnincorporatedLegalBegindate string // 非必填 法人身份证有效期开始时间（非法人结算情况下必填） 格式"20210101"
	UnincorporatedLegalEnddate   string // 非必填 法人身份证有效期结束时间（非法人结算情况下必填）格式"20210202" 长期情况下该字段为空
	OtherPic1                    string // 非必填 其他照片（请填入上传加密图片返回的值）
	Remark                       string // 非必填 备注
}

type StoreOperateResponse struct {
	StoreID     int64 `json:"store_id"`     // 门店id
	StoreStatus int8  `json:"store_status"` // 门店状态：1 待审核、2 审核通过、3 审核驳回
}

// StoreCreate  门店创建
func (s *Store) StoreCreate(store StoreOperateInfo) (storeStatus StoreOperateResponse, err error) {
	if err := util.Validate(store); err != nil {
		return storeStatus, err
	}
	param, err := processRequestParam(store, 1)
	if err != nil {
		return storeStatus, err
	}
	data, err := util.LanuchRquest(s.config, "openapi.cate.store.info.create", param)
	if err != nil {
		return storeStatus, err
	}
	if err = json.Unmarshal(data, &storeStatus); err != nil {
		return storeStatus, nil
	}
	return
}

// StoreEdit  门店信息修改
func (s *Store) StoreEdit(store StoreOperateInfo) (storeStatus StoreOperateResponse, err error) {
	if err := util.Validate(store); err != nil {
		return storeStatus, err
	}
	param, err := processRequestParam(store, 2)
	if err != nil {
		return storeStatus, err
	}
	data, err := util.LanuchRquest(s.config, "openapi.cate.store.info.update", param)
	if err != nil {
		return storeStatus, err
	}
	if err = json.Unmarshal(data, &storeStatus); err != nil {
		return storeStatus, nil
	}
	return
}

// processRequestParam 数据处理 requestType 1. 创建 2. 修改
func processRequestParam(store StoreOperateInfo, requestType int8) (map[string]interface{}, error) {
	param := make(map[string]interface{}, 0)
	switch store.LicenseType {
	case 1: // 营业执照
		if store.LicenseName == "" {
			return nil, fmt.Errorf("门店经营许可类型是营业执照 营业执照名称必填")
		}
		if store.LicenseID == "" {
			return nil, fmt.Errorf("门店经营许可类型是营业执照 营业执照号必填")
		}
		if store.LicenseTimeType != 1 && store.LicenseTimeType != 2 {
			return nil, fmt.Errorf("营业执照有效期类型错误")
		}
		if store.LicenseTimeType != 2 && (store.LicenseTimeBegin == "" || store.LicenseTimeEnd == "") {
			return nil, fmt.Errorf("营业执照类型为正常有效 营业执照时间有效期必填")
		}
	case 2: // 证明函
		if store.HandHoldIDCardPic == "" {
			return nil, fmt.Errorf("门店经营许可类型是证明函 手持身份证照片必填")
		}
	default:
		return nil, fmt.Errorf("门店经营许可类型参数错误")
	}

	if requestType == 2 && store.StoreID == 0 {
		return nil, fmt.Errorf("门店信息修改 门店ID必填")
	}

	// TODO 时间格式判断
	param["merchant_id"] = store.MerchantID
	param["store_name"] = store.StoreName
	param["street_address"] = store.StreetAddress
	param["longitude"] = store.Longitude
	param["latitude"] = store.Latitude
	param["province_code"] = store.ProvinceCode
	param["city_code"] = store.CityCode
	param["area_code"] = store.AreaCode
	param["store_phone"] = store.StorePhone
	param["unity_category_id"] = store.UnityCategoryID
	param["store_front_img_url"] = store.StoreFrontImgUrl
	param["store_env_photo"] = store.StoreEnvPhoto
	param["store_cash_photo"] = store.StoreCashPhoto
	param["license_photo"] = store.LicensePhoto
	param["license_type"] = store.LicenseType

	if requestType == 2 {
		param["store_id"] = store.StoreID
	}

	if store.HandHoldIDCardPic != "" {
		param["hand_hold_id_card_pic"] = store.HandHoldIDCardPic
	}

	if store.LicenseName != "" {
		param["license_name"] = store.LicenseName
	}

	if store.LicenseID != "" {
		param["license_id"] = store.LicenseID
	}

	if store.LicenseTimeType == 1 || store.LicenseTimeType == 2 {
		param["license_time_type"] = store.LicenseTimeType
	}

	if store.LicenseTimeBegin != "" {
		param["license_time_begin"] = store.LicenseTimeBegin
	}

	if store.LicenseTimeEnd != "" {
		param["license_time_end"] = store.LicenseTimeEnd
	}

	if store.BrandName != "" {
		param["brand_name"] = store.BrandName
	}

	if store.OperatingLicensePhoto != "" {
		param["operating_license_photo"] = store.OperatingLicensePhoto
	}

	if store.LegalIDCardFrontPhoto != "" {
		param["legal_id_card_front_photo"] = store.LegalIDCardFrontPhoto
	}

	if store.LegalIDCardBackPhoto != "" {
		param["legal_id_card_back_photo"] = store.LegalIDCardBackPhoto
	}

	if store.UnincorporatedPhoto != "" {
		param["unincorporated_photo"] = store.UnincorporatedPhoto
	}

	if store.UnincorporatedLegalName != "" {
		param["unincorporated_legal_name"] = store.UnincorporatedLegalName
	}

	if store.UnincorporatedLegalNum != "" {
		param["unincorporated_legal_num"] = store.UnincorporatedLegalNum
	}

	if store.UnincorporatedLegalBegindate != "" {
		param["unincorporated_legal_begindate"] = store.UnincorporatedLegalBegindate
	}

	if store.UnincorporatedLegalEnddate != "" {
		param["unincorporated_legal_enddate"] = store.UnincorporatedLegalEnddate
	}

	if store.UnincorporatedLegalPermanent != "" {
		// TODO 身份证两种类型判断
		param["unincorporated_legal_permanent"] = store.UnincorporatedLegalPermanent
	}

	if store.OtherPic1 != "" {
		param["other_pic1"] = store.OtherPic1
	}

	if store.Remark != "" {
		param["remark"] = store.OtherPic1
	}
	return param, nil
}
