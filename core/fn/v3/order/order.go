package order

import (
	"github.com/MQEnergy/delivery-core/core/fn/v3"
	fn2 "github.com/MQEnergy/delivery-core/routers/fn/v3"
)

type Order struct {
	Base *v3.Base
}

func New(appKey, appSecret string, online bool) *Order {
	base := v3.New(appKey, appSecret, online)
	return &Order{
		Base: base,
	}
}

// GetToken returns token visit: https://open.ele.me/documents/openApi/603
func (o *Order) GetToken(params map[string]interface{}) (string, error) {
	return o.Base.WithRequestParams(fn2.GET_TOKEN_URL, params).Result()
}

// RefreshToken returns token visit: https://open.ele.me/documents/openApi/603
func (o *Order) RefreshToken(params map[string]interface{}) (string, error) {
	return o.Base.WithRequestParams(fn2.REFRESH_TOKEN_URL, params).Result()
}

// preCreateOrder returns order info visit: https://open.ele.me/documents/openApi/1003
func (o *Order) preCreateOrder(params map[string]interface{}) (string, error) {
	return o.Base.WithRequestParams(fn2.PRE_CREATE_ORDER_URL, params).Result()
}

// CreateOrder returns order info visit: https://open.ele.me/documents/openApi/1011
func (o *Order) CreateOrder(params map[string]interface{}) (string, error) {
	return o.Base.WithRequestParams(fn2.CREATE_ORDER_URL, params).Result()
}

// AddTips returns add tips info visit: https://open.ele.me/documents/openApi/1019
func (o *Order) AddTips(params map[string]interface{}) (string, error) {
	return o.Base.WithRequestParams(fn2.ADD_TIP_URL, params).Result()
}

// getCancelReasonList returns cancel reason list visit: https://open.ele.me/documents/openApi/1027
func (o *Order) getCancelReasonList(params map[string]interface{}) (string, error) {
	return o.Base.WithRequestParams(fn2.GET_CANCEL_REASON_LIST_URL, params).Result()
}

// PreCancelOrder returns pre cancel order info visit: https://open.ele.me/documents/openApi/1035
func (o *Order) PreCancelOrder(params map[string]interface{}) (string, error) {
	return o.Base.WithRequestParams(fn2.PRE_CANCEL_ORDER_URL, params).Result()
}

// CancelOrder returns cancel order visit: https://open.ele.me/documents/openApi/1043
func (o *Order) CancelOrder(params map[string]interface{}) (string, error) {
	return o.Base.WithRequestParams(fn2.CANCEL_ORDER_URL, params).Result()
}

// RushOrder returns rush order  visit: https://open.ele.me/documents/openApi/1459
func (o *Order) RushOrder(params map[string]interface{}) (string, error) {
	return o.Base.WithRequestParams(fn2.RUSH_ORDER_URL, params).Result()
}

// GetOrderDetail returns order detail visit: https://open.ele.me/documents/openApi/1051
func (o *Order) GetOrderDetail(params map[string]interface{}) (string, error) {
	return o.Base.WithRequestParams(fn2.ORDER_DETAIL_URL, params).Result()
}

// GetKnightInfo returns knight info visit: https://open.ele.me/documents/openApi/1059
func (o *Order) GetKnightInfo(params map[string]interface{}) (string, error) {
	return o.Base.WithRequestParams(fn2.GET_KNIGHT_INFO_URL, params).Result()
}

// ComplaintOrder returns order list visit: https://open.ele.me/documents/openApi/1067
func (o *Order) ComplaintOrder(params map[string]interface{}) (string, error) {
	return o.Base.WithRequestParams(fn2.COMPLAINT_ORDER_URL, params).Result()
}

// ClaimOrder returns claim order visit: https://open.ele.me/documents/openApi/1075
func (o *Order) ClaimOrder(params map[string]interface{}) (string, error) {
	return o.Base.WithRequestParams(fn2.CLAIM_ORDER_URL, params).Result()
}
