package order

import (
	"encoding/json"
	"github.com/MQEnergy/delivery-core/core/fn/v3"
	fn3 "github.com/MQEnergy/delivery-core/routers/fn/v3"
)

type Order struct {
	Base *v3.Base
}

func New(appKey, appSecret, merchantID string, online bool) *Order {
	return &Order{
		Base: v3.New(appKey, appSecret, merchantID, online),
	}
}

// preCreateOrder returns order info visit: https://open.ele.me/documents/openApi/1003
func (o *Order) preCreateOrder(params map[string]interface{}, accessToken string) (string, error) {
	bytes, err := json.Marshal(params)
	if err != nil {
		return "", err
	}
	bizParams := v3.BizSignParams{
		MerchantID:   o.Base.MerchantID,
		AccessToken:  accessToken,
		BusinessData: string(bytes),
	}
	return o.Base.WithRequestParams(fn3.PRE_CREATE_ORDER_URL, bizParams).Result()
}

// CreateOrder returns order info visit: https://open.ele.me/documents/openApi/1011
func (o *Order) CreateOrder(params map[string]interface{}, accessToken string) (string, error) {
	bytes, err := json.Marshal(params)
	if err != nil {
		return "", err
	}
	bizParams := v3.BizSignParams{
		MerchantID:   o.Base.MerchantID,
		AccessToken:  accessToken,
		BusinessData: string(bytes),
	}
	return o.Base.WithRequestParams(fn3.CREATE_ORDER_URL, bizParams).Result()
}

// AddTips returns add tips info visit: https://open.ele.me/documents/openApi/1019
func (o *Order) AddTips(params map[string]interface{}, accessToken string) (string, error) {
	bytes, err := json.Marshal(params)
	if err != nil {
		return "", err
	}
	bizParams := v3.BizSignParams{
		MerchantID:   o.Base.MerchantID,
		AccessToken:  accessToken,
		BusinessData: string(bytes),
	}
	return o.Base.WithRequestParams(fn3.ADD_TIP_URL, bizParams).Result()
}

// getCancelReasonList returns cancel reason list visit: https://open.ele.me/documents/openApi/1027
func (o *Order) getCancelReasonList(params map[string]interface{}, accessToken string) (string, error) {
	bytes, err := json.Marshal(params)
	if err != nil {
		return "", err
	}
	bizParams := v3.BizSignParams{
		MerchantID:   o.Base.MerchantID,
		AccessToken:  accessToken,
		BusinessData: string(bytes),
	}
	return o.Base.WithRequestParams(fn3.GET_CANCEL_REASON_LIST_URL, bizParams).Result()
}

// PreCancelOrder returns pre cancel order info visit: https://open.ele.me/documents/openApi/1035
func (o *Order) PreCancelOrder(params map[string]interface{}, accessToken string) (string, error) {
	bytes, err := json.Marshal(params)
	if err != nil {
		return "", err
	}
	bizParams := v3.BizSignParams{
		MerchantID:   o.Base.MerchantID,
		AccessToken:  accessToken,
		BusinessData: string(bytes),
	}
	return o.Base.WithRequestParams(fn3.PRE_CANCEL_ORDER_URL, bizParams).Result()
}

// CancelOrder returns cancel order visit: https://open.ele.me/documents/openApi/1043
func (o *Order) CancelOrder(params map[string]interface{}, accessToken string) (string, error) {
	bytes, err := json.Marshal(params)
	if err != nil {
		return "", err
	}
	bizParams := v3.BizSignParams{
		MerchantID:   o.Base.MerchantID,
		AccessToken:  accessToken,
		BusinessData: string(bytes),
	}
	return o.Base.WithRequestParams(fn3.CANCEL_ORDER_URL, bizParams).Result()
}

// RushOrder returns rush order  visit: https://open.ele.me/documents/openApi/1459
func (o *Order) RushOrder(params map[string]interface{}, accessToken string) (string, error) {
	bytes, err := json.Marshal(params)
	if err != nil {
		return "", err
	}
	bizParams := v3.BizSignParams{
		MerchantID:   o.Base.MerchantID,
		AccessToken:  accessToken,
		BusinessData: string(bytes),
	}
	return o.Base.WithRequestParams(fn3.RUSH_ORDER_URL, bizParams).Result()
}

// GetOrderDetail returns order detail visit: https://open.ele.me/documents/openApi/1051
func (o *Order) GetOrderDetail(params map[string]interface{}, accessToken string) (string, error) {
	bytes, err := json.Marshal(params)
	if err != nil {
		return "", err
	}
	bizParams := v3.BizSignParams{
		MerchantID:   o.Base.MerchantID,
		AccessToken:  accessToken,
		BusinessData: string(bytes),
	}
	return o.Base.WithRequestParams(fn3.ORDER_DETAIL_URL, bizParams).Result()
}

// GetKnightInfo returns knight info visit: https://open.ele.me/documents/openApi/1059
func (o *Order) GetKnightInfo(params map[string]interface{}, accessToken string) (string, error) {
	bytes, err := json.Marshal(params)
	if err != nil {
		return "", err
	}
	bizParams := v3.BizSignParams{
		MerchantID:   o.Base.MerchantID,
		AccessToken:  accessToken,
		BusinessData: string(bytes),
	}
	return o.Base.WithRequestParams(fn3.GET_KNIGHT_INFO_URL, bizParams).Result()
}

// ComplaintOrder returns order list visit: https://open.ele.me/documents/openApi/1067
func (o *Order) ComplaintOrder(params map[string]interface{}, accessToken string) (string, error) {
	bytes, err := json.Marshal(params)
	if err != nil {
		return "", err
	}
	bizParams := v3.BizSignParams{
		MerchantID:   o.Base.MerchantID,
		AccessToken:  accessToken,
		BusinessData: string(bytes),
	}
	return o.Base.WithRequestParams(fn3.COMPLAINT_ORDER_URL, bizParams).Result()
}

// ClaimOrder returns claim order visit: https://open.ele.me/documents/openApi/1075
func (o *Order) ClaimOrder(params map[string]interface{}, accessToken string) (string, error) {
	bytes, err := json.Marshal(params)
	if err != nil {
		return "", err
	}
	bizParams := v3.BizSignParams{
		MerchantID:   o.Base.MerchantID,
		AccessToken:  accessToken,
		BusinessData: string(bytes),
	}
	return o.Base.WithRequestParams(fn3.CLAIM_ORDER_URL, bizParams).Result()
}
