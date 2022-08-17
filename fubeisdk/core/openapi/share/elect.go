package share

import (
	"encoding/json"
	"fmt"

	"git.myarena7.com/arena/fubeisdk/core/config"
	"git.myarena7.com/arena/fubeisdk/util"
)

type ElectContract struct {
	config config.Config
}

func NewElectContract(cfg config.Config) *ElectContract {
	return &ElectContract{config: cfg}
}

// ShareAccreditQueryResp 分账电子协议签署结果返回参数
type ShareAccreditQueryResp struct {
	//0 未开通
	//1 已开通
	//2 商户主动关闭
	//3 待审核
	//4 冻结
	//5 注销
	//6 待签合同
	ShareType    int32  `json:"shareType"`
	ShareTypeStr string `json:"shareTypeStr"`
	ContractUrl  string `json:"contractUrl"` // 合同h5地址 电子合同接口返回 只有已进行电子合同授权申请但未签署的商户会返回该链接
}

// GetShareElectContract 获取线上分账电子协议书
func (e *ElectContract) GetShareElectContract(merchantID int64) (string, error) {
	if merchantID == 0 {
		return "", fmt.Errorf("merchant_id is zero")
	}
	data, err := util.LanuchRquest(e.config, "openapi.share.elect.contract.accredit", map[string]int64{
		"merchant_id": merchantID,
	})
	if err != nil {
		return "", err
	}
	url := ""
	if err := json.Unmarshal(data, &url); err != nil {
		return "", err
	}
	return url, nil
}

func (e *ElectContract) GetShareElectContractResult(merchantID int64) (ShareAccreditQueryResp, error) {
	if merchantID == 0 {
		return ShareAccreditQueryResp{}, fmt.Errorf("merchant_id is zero")
	}
	data, err := util.LanuchRquest(e.config, "openapi.share.accredit.query", map[string]int64{
		"merchant_id": merchantID,
	})
	if err != nil {
		return ShareAccreditQueryResp{}, err
	}
	var result ShareAccreditQueryResp
	if err := json.Unmarshal(data, &result); err != nil {
		return ShareAccreditQueryResp{}, err
	}
	return result, nil
}
