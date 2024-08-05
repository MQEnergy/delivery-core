package shop

import (
	fnv2 "github.com/MQEnergy/delivery-core/core/fn/v2"
	routerv2 "github.com/MQEnergy/delivery-core/routers/fn/v2"
)

type Shop struct {
	Base *fnv2.Base
}

func New(appKey, appSecret string, online bool) *Shop {
	base := fnv2.New(appKey, appSecret, online)
	return &Shop{
		Base: base,
	}
}

// ChainStore chain store visit: https://open.ele.me/documents/ka/1371
func (s *Shop) ChainStore(params map[string]interface{}, accessToken string) (string, error) {
	return s.Base.WithRequestParams(routerv2.CHAIN_STORE_URL, params, accessToken).Result()
}

// ChainStoreQuery query chain store visit: https://open.ele.me/documents/ka/1267
func (s *Shop) ChainStoreQuery(params map[string]interface{}, accessToken string) (string, error) {
	return s.Base.WithRequestParams(routerv2.CHAIN_STORE_QUERY_URL, params, accessToken).Result()
}

// ChainStoreUpdate chain store update visit: https://open.ele.me/documents/ka/1363
func (s *Shop) ChainStoreUpdate(params map[string]interface{}, accessToken string) (string, error) {
	return s.Base.WithRequestParams(routerv2.CHAIN_STORE_UPDATE_URL, params, accessToken).Result()
}

// ChainStoreDeliveryArea query chain store delivery area visit: https://open.ele.me/documents/ka/1259
func (s *Shop) ChainStoreDeliveryArea(params map[string]interface{}, accessToken string) (string, error) {
	return s.Base.WithRequestParams(routerv2.CHAIN_STORE_DELIVERY_AREA_URL, params, accessToken).Result()
}
