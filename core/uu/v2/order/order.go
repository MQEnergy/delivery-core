package order

import (
	"github.com/MQEnergy/delivery-core/core/uu/v2"
	uu2 "github.com/MQEnergy/delivery-core/routers/uu/v2"
)

type Order struct {
	Base *v2.Base
}

func New(appID, appKey, openID string, online bool) *Order {
	return &Order{
		Base: v2.New(appID, appKey, openID, online),
	}
}

// GetOrderPrice ... visit: https://open.uupt.com/#/ApiDocument/InterfaceList/getorderprice?PlatType=PLAT_MERCHANT&version=2&t=%E8%AE%A1%E7%AE%97%E8%AE%A2%E5%8D%95%E4%BB%B7%E6%A0%BC&index=1-5-1
func (o *Order) GetOrderPrice(params map[string]interface{}) (string, error) {
	return o.Base.WithRequestParams(uu2.ORDER_FEE_URL, params).Result()
}

// CreateOrder ... visit: https://open.uupt.com/#/ApiDocument/InterfaceList/addorder?PlatType=PLAT_MERCHANT&version=2&t=%E5%8F%91%E5%B8%83%E8%AE%A2%E5%8D%95&index=1-5-2
func (o *Order) CreateOrder(params map[string]interface{}) (string, error) {
	return o.Base.WithRequestParams(uu2.CREATE_ORDER_URL, params).Result()
}

// GetOrderDetail ... visit: https://open.uupt.com/#/ApiDocument/InterfaceList/getorderdetail?PlatType=PLAT_MERCHANT&version=2&t=%E8%8E%B7%E5%8F%96%E8%AE%A2%E5%8D%95%E8%AF%A6%E6%83%85&index=1-5-3
func (o *Order) GetOrderDetail(params map[string]interface{}) (string, error) {
	return o.Base.WithRequestParams(uu2.ORDER_DETAIL_URL, params).Result()
}

// CityCodeList ... visit: https://open.uupt.com/#/ApiDocument/InterfaceList/citylist?PlatType=PLAT_MERCHANT&version=2&t=%E5%B7%B2%E5%BC%80%E9%80%9A%E5%9F%8E%E5%B8%82%E5%8C%BA%E5%8E%BF%E5%88%97%E8%A1%A8&index=1-5-4
func (o *Order) CityCodeList(params map[string]interface{}) (string, error) {
	return o.Base.WithRequestParams(uu2.CITY_CODE_LIST_URL, params).Result()
}

// AddTip ... visit: https://open.uupt.com/#/ApiDocument/InterfaceList/payonlinefee?PlatType=PLAT_MERCHANT&version=2&t=%E6%94%AF%E4%BB%98%E5%B0%8F%E8%B4%B9&index=1-5-6
func (o *Order) AddTip(params map[string]interface{}) (string, error) {
	return o.Base.WithRequestParams(uu2.ADD_TIP_URL, params).Result()
}

// GetHomeServiceFee ... visit: https://open.uupt.com/#/ApiDocument/InterfaceList/gethomeservicefee?PlatType=PLAT_MERCHANT&version=2&t=%E8%8E%B7%E5%BE%97%E8%BF%9D%E7%BA%A6%E9%87%91&index=1-5-7
func (o *Order) GetHomeServiceFee(params map[string]interface{}) (string, error) {
	return o.Base.WithRequestParams(uu2.DEDUCT_FEE_URL, params).Result()
}

// CancelOrder ... visit: https://open.uupt.com/#/ApiDocument/InterfaceList/cancelorder?PlatType=PLAT_MERCHANT&version=2&t=%E5%8F%96%E6%B6%88%E8%AE%A2%E5%8D%95&index=1-5-8
func (o *Order) CancelOrder(params map[string]interface{}) (string, error) {
	return o.Base.WithRequestParams(uu2.CANCEL_ORDER_URL, params).Result()
}

// GetBalanceDetail ... visit: https://open.uupt.com/#/ApiDocument/InterfaceList/getbalancedetail?PlatType=PLAT_MERCHANT&version=2&t=%E8%8E%B7%E5%8F%96%E4%BD%99%E9%A2%9D%E8%AF%A6%E6%83%85&index=1-5-9
func (o *Order) GetBalanceDetail(params map[string]interface{}) (string, error) {
	return o.Base.WithRequestParams(uu2.GET_BALANCE_DETAIL_URL, params).Result()
}

// GetRecharge ... visit: https://open.uupt.com/#/ApiDocument/InterfaceList/getrecharge?PlatType=PLAT_MERCHANT&version=2&t=%E8%8E%B7%E5%8F%96%E5%85%85%E5%80%BC%E5%9C%B0%E5%9D%80&index=1-5-10
func (o *Order) GetRecharge(params map[string]interface{}) (string, error) {
	return o.Base.WithRequestParams(uu2.GET_RECHARGE_URL, params).Result()
}

// BindUserApply ... visit: https://open.uupt.com/#/ApiDocument/InterfaceList/binduserapply?PlatType=PLAT_MERCHANT&version=2&t=%E5%8F%91%E9%80%81%E7%9F%AD%E4%BF%A1%E9%AA%8C%E8%AF%81%E7%A0%81&index=1-5-11
func (o *Order) BindUserApply(params map[string]interface{}) (string, error) {
	return o.Base.WithRequestParams(uu2.BIND_USER_APPLY_URL, params).Result()
}

// BindUserSubmit ... visit: https://open.uupt.com/#/ApiDocument/InterfaceList/bindusersubmit?PlatType=PLAT_MERCHANT&version=2&t=%E8%8E%B7%E5%8F%96%E7%94%A8%E6%88%B7openid&index=1-5-12
func (o *Order) BindUserSubmit(params map[string]interface{}) (string, error) {
	return o.Base.WithRequestParams(uu2.BIND_USER_SUBMIT_URL, params).Result()
}

// CancelBind ... visit: https://open.uupt.com/#/ApiDocument/InterfaceList/cancelbind?PlatType=PLAT_MERCHANT&version=2&t=%E7%94%A8%E6%88%B7%E8%A7%A3%E9%99%A4%E7%BB%91%E5%AE%9A&index=1-5-13
func (o *Order) CancelBind(params map[string]interface{}) (string, error) {
	return o.Base.WithRequestParams(uu2.CANCEL_BIND_URL, params).Result()
}

// GetShopList ... visit: https://open.uupt.com/#/ApiDocument/InterfaceList/getshoplist?PlatType=PLAT_MERCHANT&version=2&t=%E8%8E%B7%E5%8F%96%E9%97%A8%E5%BA%97%E5%88%97%E8%A1%A8&index=1-5-14
func (o *Order) GetShopList(params map[string]interface{}) (string, error) {
	return o.Base.WithRequestParams(uu2.GET_SHOP_LIST_URL, params).Result()
}
