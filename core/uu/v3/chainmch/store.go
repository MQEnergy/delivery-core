package chainmch

import (
	v3 "github.com/MQEnergy/delivery-core/core/uu/v3"
	uu3 "github.com/MQEnergy/delivery-core/routers/uu/v3"
)

type ChainMch struct {
	Base *v3.MchBase
}

func New(appID, appKey string, merchantID uint64, storeNo string, online bool) *ChainMch {
	return &ChainMch{
		Base: v3.NewBase(appID, appKey, merchantID, storeNo, online),
	}
}

// MchAddChainStore ...
// 连锁商家 visit: https://open.uupt.com/#/ApiDocument/v3/InterfaceList/addChainStore?PlatType=CHAIN_MERCHANT&version=3&t=%E6%B7%BB%E5%8A%A0%E8%BF%9E%E9%94%81%E9%97%A8%E5%BA%97&index=3-5-9
func (o *ChainMch) MchAddChainStore(params map[string]interface{}) (string, error) {
	return o.Base.WithRequestParams(uu3.MCH_ADD_CHAIN_STORE_URL, params).Result()
}

// MchEditChainStore ...
func (o *ChainMch) MchEditChainStore(params map[string]interface{}) (string, error) {
	return o.Base.WithRequestParams(uu3.MCH_EDIT_CHAIN_STORE_URL, params).Result()
}
