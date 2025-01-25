package pltmch

import (
	v3 "github.com/MQEnergy/delivery-core/core/uu/v3"
	uu3 "github.com/MQEnergy/delivery-core/routers/uu/v3"
)

type PltMch struct {
	Base *v3.Base
}

func New(appID, appKey, openID string, online bool) *PltMch {
	return &PltMch{
		Base: v3.New(appID, appKey, openID, online),
	}
}

// OpenCityList ...
// 平台服务商 visit: https://open.uupt.com/#/ApiDocument/v3/InterfaceList/openCityList?PlatType=PLAT_MERCHANT&version=3&t=%E8%8E%B7%E5%8F%96%E5%BC%80%E6%94%BE%E5%9F%8E%E5%B8%82%E5%88%97%E8%A1%A8&index=1-5-1
// 自营商家 visit: https://open.uupt.com/#/ApiDocument/v3/InterfaceList/openCityList?PlatType=SELF_MERCHANT&version=3&t=%E8%8E%B7%E5%8F%96%E5%BC%80%E6%94%BE%E5%9F%8E%E5%B8%82%E5%88%97%E8%A1%A8&index=2-5-1
func (o *PltMch) OpenCityList(params map[string]interface{}) (string, error) {
	return o.Base.WithRequestParams(uu3.OPEN_CITY_LIST_URL, params).Result()
}

// OrderPrice ...
func (o *PltMch) OrderPrice(params map[string]interface{}) (string, error) {
	return o.Base.WithRequestParams(uu3.ORDER_PRICE_URL, params).Result()
}

// AddOrder ...
func (o *PltMch) AddOrder(params map[string]interface{}) (string, error) {
	return o.Base.WithRequestParams(uu3.ADD_ORDER_URL, params).Result()
}

// OrderDetail ...
func (o *PltMch) OrderDetail(params map[string]interface{}) (string, error) {
	return o.Base.WithRequestParams(uu3.ORDER_DETAIL_URL, params).Result()
}

// CancelFee ...
func (o *PltMch) CancelFee(params map[string]interface{}) (string, error) {
	return o.Base.WithRequestParams(uu3.CANCEL_FEE_URL, params).Result()
}

// CancelOrder ...
func (o *PltMch) CancelOrder(params map[string]interface{}) (string, error) {
	return o.Base.WithRequestParams(uu3.CANCEL_ORDER_URL, params).Result()
}

// AddTip ...
func (o *PltMch) AddTip(params map[string]interface{}) (string, error) {
	return o.Base.WithRequestParams(uu3.ADD_TIP_URL, params).Result()
}

// SyncPickupCode ...
func (o *PltMch) SyncPickupCode(params map[string]interface{}) (string, error) {
	return o.Base.WithRequestParams(uu3.SYNC_PICKUP_CODE_URL, params).Result()
}
