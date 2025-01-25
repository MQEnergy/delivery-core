package chainmch

import (
	uu3 "github.com/MQEnergy/delivery-core/routers/uu/v3"
)

// OpenCityList ...
// 连锁商家 visit: https://open.uupt.com/#/ApiDocument/v3/InterfaceList/openCityList?PlatType=CHAIN_MERCHANT&version=3&t=%E8%8E%B7%E5%8F%96%E5%BC%80%E6%94%BE%E5%9F%8E%E5%B8%82%E5%88%97%E8%A1%A8&index=3-5-1
func (o *ChainMch) OpenCityList(params map[string]interface{}) (string, error) {
	return o.Base.WithRequestParams(uu3.OPEN_CITY_LIST_URL, params).Result()
}

// OrderDetail ...
func (o *ChainMch) OrderDetail(params map[string]interface{}) (string, error) {
	return o.Base.WithRequestParams(uu3.ORDER_DETAIL_URL, params).Result()
}

// CancelFee ...
func (o *ChainMch) CancelFee(params map[string]interface{}) (string, error) {
	return o.Base.WithRequestParams(uu3.CANCEL_FEE_URL, params).Result()
}

// CancelOrder ...
func (o *ChainMch) CancelOrder(params map[string]interface{}) (string, error) {
	return o.Base.WithRequestParams(uu3.CANCEL_ORDER_URL, params).Result()
}

// AddTip ...
func (o *ChainMch) AddTip(params map[string]interface{}) (string, error) {
	return o.Base.WithRequestParams(uu3.ADD_TIP_URL, params).Result()
}

// SyncPickupCode ...
func (o *ChainMch) SyncPickupCode(params map[string]interface{}) (string, error) {
	return o.Base.WithRequestParams(uu3.SYNC_PICKUP_CODE_URL, params).Result()
}

// MchOrderPrice ...
func (o *ChainMch) MchOrderPrice(params map[string]interface{}) (string, error) {
	return o.Base.WithRequestParams(uu3.MCH_ORDER_PRICE_URL, params).Result()
}

// MchAddOrder ...
func (o *ChainMch) MchAddOrder(params map[string]interface{}) (string, error) {
	return o.Base.WithRequestParams(uu3.MCH_ADD_ORDER_URL, params).Result()
}
