// package base 包含银行查询、省市自治区查询、图片上传等基础接口

package base

import (
	"encoding/json"
	"fmt"

	"git.myarena7.com/arena/fubeisdk/core/config"
	"git.myarena7.com/arena/fubeisdk/util"
)

type Base struct {
	config config.Config
}

func NewBase(cfg config.Config) *Base {
	return &Base{config: cfg}
}

// ImgLoadRequest 图片上传请求参数
type ImgLoadRequest struct {
	FileData string `json:"file_data" validate:"required"` // 图片编码之后的base64的值
	/*
		//  bankCard：银行卡，如：对私银行卡照片、对公开户许可证
		// 	idCard：身份证，如：个人实名身份证、法人身份证
		// 	license：营业执照/证明函，如：营业执照、证明函
		// 	store：门店，如：门店门头照、店内环境照、收银台照片、经营许可证
		// 	other：其他，如：其他辅助证明
	*/
	BusType string `json:"bus_type" validate:"required"` // 类型
}

// CategoryListResponse 行业类目获取返回参数
type CategoryListResponse struct {
	CateGoryList []CateGoryList `json:"category_list"`
}

type CateGoryList struct {
	Code  int32  `json:"code"`  // 类目编码
	Level int32  `json:"level"` // 类目级别
	Name  string `json:"name"`  // 类目名称
}

// BankCodeResponse 个人银行账户查询返回参数
type BankCodeResponse struct {
	LegalFlag bool   `json:"legal_flag"` // 检验结果
	BankID    int32  `json:"bank_id"`    // 卡主键信息
	Message   string `json:"message"`    // 返回信息
	BankName  string `json:"bank_name"`  // 银行名称
	BankCode  string `json:"bank_code"`  // 银行代码
}

// BankList 银行编号列表
type BankList struct {
	BankList []Bank `json:"bank_list"`
}

type Bank struct {
	BankCode string `json:"bank_code"` // 银行代码
	BankNO   string `json:"bank_no"`   // 银行编码（超级网银号）
	BankName string `json:"bank_name"` // 银行名称
}

// BankAreaResponse 银行区域编号查询返回参数
type BankAreaResponse struct {
	AreaCodeList []BankArea `json:"area_code_list"`
}
type BankArea struct {
	Code         string `json:"code"`           // 地区编码
	Name         string `json:"name"`           // 地区名称
	BankCityCode string `json:"bank_city_code"` // 银行城市编码
}

// BranchBankRequest 支行信息获取请求参数
type BranchBankRequest struct {
	CityCode string `json:"city_code" validate:"required"` // 银行区域编号，通过 “银行区域编号列表中获取对应的"bankCityCode"
	BankName string `json:"bank_name" validate:"required"` // 银行名称。个人账户类型，通过“个人银行账户查询”获取；企业账户类型，通过“银行编号列表”
}

// BranchBankResponse 支行信息获取返回参数
type BranchBankResponse struct {
	UnionPayCode string `json:"unionpay_code"` // 开户支行联行号
	BranchName   string `json:"branch_name"`   // 支行名称
}

// ArenaResponse 省市自治区获取返回参数
type ArenaResponse struct {
	CodeList []ArenaList `json:"code_list"`
}

type ArenaList struct {
	Code string `json:"code"` // 编码
	Name string `json:"name"` // 名称
}

// ImgUploadResponse 图片上传返回参数
type ImgUploadResponse struct {
	ResourceID string `json:"resource_id"`
}

// AgentAppIDSecretRequest 获取appid和secret请求参数
type AgentAppIDSecretRequest struct {
	AgentID   int64  `json:"agent_id"`  // 代理ID
	Signature string `json:"signature"` // 签名
}

// AgentAppIDSecretResponse 获取appid和secret的返回参数
type AgentAppIDSecretResponse struct {
	AppidAndSecret string `json:"appid_and_secret"` // 加密的appid和secre
}

// ImgUpload 图片上传
//bankCard：银行卡，如：对私银行卡照片、对公开户许可证
//idCard：身份证，如：个人实名身份证、法人身份证
//license：营业执照/证明函，如：营业执照、证明函
//store：门店，如：门店门头照、店内环境照、收银台照片、经营许可证
//other：其他，如：其他辅助证明
// BusType 是上传图片的业务属性 img.FileData 是base64之后的数据
// fileType 是通过读取的byte文件流前512位拿出来的文件类型
func (b *Base) ImgUpload(img ImgLoadRequest, fileType string) (string, error) {
	if err := util.Validate(img); err != nil {
		return "", err
	}
	// 图片业务类型校验
	if img.BusType != "bankCard" && img.BusType != "idCard" && img.BusType != "license" && img.BusType != "store" && img.BusType != "other" {
		return "", fmt.Errorf("parameter error")
	}

	// 允许传递的图片类型校验 仅仅允许jpeg jpg png
	if fileType != "image/jpeg" && fileType != "image/jpg" && fileType != "image/png" {
		return "", fmt.Errorf("unsupported picture formats:%v", fileType)
	}

	// TODO 根据base64做文件类型判断
	// jpeg,jpg格式文件标志头是FF D8 base64编码之后开头是 /9j
	// png是 89 50 4E 47 0D 0A 1A 0A base64编码之后好像是 iVBORw0KGgo
	// fileType := ""
	// if strings.HasPrefix(img.FileData, "/9j") {
	// 	fileType = "image/jpeg"
	// }

	// 文件大小判定 不允许超过2M
	if !util.GetFileSizeFromBase64Result(img.FileData) {
		return "", fmt.Errorf("the image is too large")
	}
	img.FileData = fmt.Sprintf("data:%v;base64,%v", fileType, img.FileData)
	data, err := util.LanuchRquest(b.config, "openapi.agent.base.imgupload.security", img)
	if err != nil {
		return "", err
	}
	var imgAddress ImgUploadResponse
	if err = json.Unmarshal(data, &imgAddress); err != nil {
		return "", err
	}
	return imgAddress.ResourceID, nil
}

// GetCategoryList 行业类目
// 父类目，不传或者传递0返回全部1级类目
func (b *Base) GetCategoryList(parentCode int32) ([]CateGoryList, error) {
	data, err := util.LanuchRquest(b.config, "fbpay.attachment.category.list", map[string]int32{
		"parent_code": parentCode,
	})
	if err != nil {
		return nil, err
	}
	var cateGoryList CategoryListResponse
	if err = json.Unmarshal(data, &cateGoryList); err != nil {
		return nil, err
	}
	return cateGoryList.CateGoryList, nil
}

// GetBankCode 银行编号列表
// 个人银行账户见 GetPersonalBankAccount
func (b *Base) GetBankCode(bankName string) ([]Bank, error) {
	if bankName == "" {
		return nil, fmt.Errorf("parameter error")
	}
	data, err := util.LanuchRquest(b.config, "fbpay.attachment.bank.list", map[string]string{
		"bank_name": bankName,
	})
	if err != nil {
		return nil, err
	}
	var banklist BankList
	if err := json.Unmarshal(data, &banklist); err != nil {
		return nil, err
	}
	return banklist.BankList, nil
}

// GetPersonalBankAccount 个人银行账户查询
// 公司账户类型见GetBankCode()
func (b *Base) GetPersonalBankAccount(bankCard string) (backBankCode BankCodeResponse, err error) {
	if bankCard == "" {
		return backBankCode, fmt.Errorf("parameter error")
	}
	// TODO 银行卡号校验
	data, err := util.LanuchRquest(b.config, "openapi.agent.base.banks", map[string]string{
		"bank_card": bankCard,
	})
	if err != nil {
		return backBankCode, err
	}
	if err := json.Unmarshal(data, &backBankCode); err != nil {
		return backBankCode, err
	}
	return
}

// GetBranchBank 支行信息查询
func (b *Base) GetBranchBank(branchName *BranchBankRequest) (list []BranchBankResponse, err error) {
	if err := util.Validate(branchName); err != nil {
		return nil, err
	}
	data, err := util.LanuchRquest(b.config, "openapi.agent.base.branches", branchName)
	if err != nil {
		return list, err
	}
	if err = json.Unmarshal(data, &list); err != nil {
		return nil, err
	}
	return
}

// GetBankArea 获取银行区域编号
func (b *Base) GetBankArea(level int8, code string) ([]BankArea, error) {
	if level < 0 || level > 1 {
		return nil, fmt.Errorf("parameter error")
	}
	var param map[string]interface{}
	if level == 0 && code == "" {
		param = map[string]interface{}{
			"level": 0,
		}
	} else {
		param = map[string]interface{}{
			"level": level,
			"code":  code,
		}
	}
	data, err := util.LanuchRquest(b.config, "fbpay.attachment.bank.area.list", param)
	if err != nil {
		return nil, err
	}
	var bankArea BankAreaResponse
	if err := json.Unmarshal(data, &bankArea); err != nil {
		return nil, err
	}
	return bankArea.AreaCodeList, nil
}

// GetArena 全国省市自治区编号信息
// level为0时是省级别, 1是市级别
// code不传递默认返回全部的省级和直辖市列表
// 当 level 为 0，code 需要为省编码，返回该省下的市列表
// 当 level 为 1, code 需要为市编码，返回该市下的县列表
func (b *Base) GetArena(level int8, code string) ([]ArenaList, error) {
	if level < 0 || level > 1 {
		return nil, fmt.Errorf("parameter error")
	}
	var param map[string]interface{}
	if level == 0 && code == "" {
		param = map[string]interface{}{
			"level": 0,
		}
	} else {
		param = map[string]interface{}{
			"level": level,
			"code":  code,
		}
	}
	data, err := util.LanuchRquest(b.config, "fbpay.attachment.area.list", param)
	if err != nil {
		return nil, err
	}
	var provinceAndCityeCodes ArenaResponse
	if err := json.Unmarshal(data, &provinceAndCityeCodes); err != nil {
		return nil, err
	}
	return provinceAndCityeCodes.CodeList, err
}

// GetAppIDAndSecret 动态获取appid和secret
func (b *Base) GetAppIDAndSecret(a AgentAppIDSecretRequest) (string, string, error) {
	data, err := util.LanuchRquest(b.config, "openapi.agent.base.get.agent.open.security", a)
	if err != nil {
		return "", "", err
	}
	var agentAppIDAndSecret AgentAppIDSecretResponse
	if err := json.Unmarshal(data, &agentAppIDAndSecret); err != nil {
		return "", "", err
	}
	// TODO 如何解析 文档暂时无法看见 先做其他的
	return "", "", nil
}
