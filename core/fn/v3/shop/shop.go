package shop

import (
	"encoding/json"
	"github.com/MQEnergy/delivery-core/core/fn/v3"
	fn3 "github.com/MQEnergy/delivery-core/routers/fn/v3"
)

type Shop struct {
	Base *v3.Base
}

// New is the constructor for Shop.
func New(appKey, appSecret, merchantID string, online bool) *Shop {
	return &Shop{
		Base: v3.New(appKey, appSecret, merchantID, online),
	}
}

// GetAmount returns amount visit: https://open.ele.me/documents/openApi/1091
func (b *Shop) GetAmount(params map[string]interface{}, accessToken string) (string, error) {
	bytes, err := json.Marshal(params)
	if err != nil {
		return "", err
	}
	bizParams := v3.BizSignParams{
		MerchantID:   b.Base.MerchantID,
		AccessToken:  accessToken,
		BusinessData: string(bytes),
	}
	return b.Base.WithRequestParams(fn3.GET_AMOUNT_URL, bizParams).Result()
}

// ChainStoreCreate returns create store visit: https://open.ele.me/documents/openApi/1099
func (b *Shop) ChainStoreCreate(params map[string]interface{}, accessToken string) (string, error) {
	bytes, err := json.Marshal(params)
	if err != nil {
		return "", err
	}
	bizParams := v3.BizSignParams{
		MerchantID:   b.Base.MerchantID,
		AccessToken:  accessToken,
		BusinessData: string(bytes),
	}
	return b.Base.WithRequestParams(fn3.CHAIN_STORE_CREATE_URL, bizParams).Result()
}

// ChainStoreUpdate returns update store visit: https://open.ele.me/documents/openApi/1107
func (b *Shop) ChainStoreUpdate(params map[string]interface{}, accessToken string) (string, error) {
	bytes, err := json.Marshal(params)
	if err != nil {
		return "", err
	}
	bizParams := v3.BizSignParams{
		MerchantID:   b.Base.MerchantID,
		AccessToken:  accessToken,
		BusinessData: string(bytes),
	}
	return b.Base.WithRequestParams(fn3.CHAIN_STORE_UPDATE_URL, bizParams).Result()
}

// ChainStoreQuery returns query store visit: https://open.ele.me/documents/openApi/1115
func (b *Shop) ChainStoreQuery(params map[string]interface{}, accessToken string) (string, error) {
	bytes, err := json.Marshal(params)
	if err != nil {
		return "", err
	}
	bizParams := v3.BizSignParams{
		MerchantID:   b.Base.MerchantID,
		AccessToken:  accessToken,
		BusinessData: string(bytes),
	}
	return b.Base.WithRequestParams(fn3.CHAIN_STORE_QUERY_URL, bizParams).Result()
}

// ChainStoreQueryList returns query store list visit: https://open.ele.me/documents/openApi/1123
func (b *Shop) ChainStoreQueryList(params map[string]interface{}, accessToken string) (string, error) {
	bytes, err := json.Marshal(params)
	if err != nil {
		return "", err
	}
	bizParams := v3.BizSignParams{
		MerchantID:   b.Base.MerchantID,
		AccessToken:  accessToken,
		BusinessData: string(bytes),
	}
	return b.Base.WithRequestParams(fn3.CHAIN_STORE_QUERY_LIST_URL, bizParams).Result()
}

// ChainStoreRange returns store range visit: https://open.ele.me/documents/openApi/1131
func (b *Shop) ChainStoreRange(params map[string]interface{}, accessToken string) (string, error) {
	bytes, err := json.Marshal(params)
	if err != nil {
		return "", err
	}
	bizParams := v3.BizSignParams{
		MerchantID:   b.Base.MerchantID,
		AccessToken:  accessToken,
		BusinessData: string(bytes),
	}
	return b.Base.WithRequestParams(fn3.CHAIN_STORE_RANGE_URL, bizParams).Result()
}

// GetAvailableExtraGoodsList return available extra goods list visit: https://open.ele.me/documents/openApi/1604
func (b *Shop) GetAvailableExtraGoodsList(params map[string]interface{}, accessToken string) (string, error) {
	bytes, err := json.Marshal(params)
	if err != nil {
		return "", err
	}
	bizParams := v3.BizSignParams{
		MerchantID:   b.Base.MerchantID,
		AccessToken:  accessToken,
		BusinessData: string(bytes),
	}
	return b.Base.WithRequestParams(fn3.GET_AVAILABLE_EXTRA_GOODS_LIST_URL, bizParams).Result()
}

// NewUploadFile return new upload file visit: https://open.ele.me/documents/openApi/1083
func (b *Shop) NewUploadFile(params map[string]interface{}, accessToken string) (string, error) {
	bytes, err := json.Marshal(params)
	if err != nil {
		return "", err
	}
	bizParams := v3.BizSignParams{
		MerchantID:   b.Base.MerchantID,
		AccessToken:  accessToken,
		BusinessData: string(bytes),
	}
	return b.Base.WithRequestParams(fn3.NEW_UPLOAD_FILE_URL, bizParams).Result()
}
