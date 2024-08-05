package order

import (
	fnv2 "github.com/MQEnergy/delivery-core/core/fn/v2"
	routerv2 "github.com/MQEnergy/delivery-core/routers/fn/v2"
)

type Order struct {
	Base *fnv2.Base
}

func New(appKey, appSecret string, online bool) *Order {
	base := fnv2.New(appKey, appSecret, online)
	return &Order{
		Base: base,
	}
}

// CreateOrder creates a new order visit: https://open.ele.me/documents/ka/1139
func (o *Order) CreateOrder(params map[string]interface{}, accessToken string) (string, error) {
	return o.Base.WithRequestParams(routerv2.CREATE_ORDER_URL, params, accessToken).Result()
}

// CancelOrder cancel the order visit: https://open.ele.me/documents/ka/1411
func (o *Order) CancelOrder(params map[string]interface{}, accessToken string) (string, error) {
	return o.Base.WithRequestParams(routerv2.CANCEL_ORDER_URL, params, accessToken).Result()
}

// GetOrderDetail returns the order details visit: https://open.ele.me/documents/ka/1435
func (o *Order) GetOrderDetail(params map[string]interface{}, accessToken string) (string, error) {
	return o.Base.WithRequestParams(routerv2.ORDER_DETAIL_URL, params, accessToken).Result()
}

// ComplaintOrder returns order complaint visit: https://open.ele.me/documents/ka/1355
func (o *Order) ComplaintOrder(params map[string]interface{}, accessToken string) (string, error) {
	return o.Base.WithRequestParams(routerv2.COMPLAINT_ORDER_URL, params, accessToken).Result()
}

// OrderLogistics query order logistics visit: https://open.ele.me/documents/ka/1291
func (o *Order) OrderLogistics(params map[string]interface{}, accessToken string) (string, error) {
	return o.Base.WithRequestParams(routerv2.ORDER_LOGISTICS_URL, params, accessToken).Result()
}

// OrderCarrierRoute query order carrier route visit: https://open.ele.me/documents/ka/1283
func (o *Order) OrderCarrierRoute(params map[string]interface{}, accessToken string) (string, error) {
	return o.Base.WithRequestParams(routerv2.ORDER_CARRIER_ROUTE_URL, params, accessToken).Result()
}
