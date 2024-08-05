package shop

import (
	"github.com/MQEnergy/delivery-core/core/fn/v3"
	fn2 "github.com/MQEnergy/delivery-core/routers/fn/v3"
)

type Shop struct {
	Base *v3.Base
}

// New is the constructor for Shop.
func New(appKey, appSecret string, online bool) *Shop {
	return &Shop{
		Base: v3.New(appKey, appSecret, online),
	}
}

// GetAmount returns amount visit: https://open.ele.me/documents/openApi/1091
func (b *Shop) GetAmount(params map[string]interface{}) (string, error) {
	return b.Base.WithRequestParams(fn2.GET_AMOUNT_URL, params).Result()
}

// ChainStoreCreate returns create store visit: https://open.ele.me/documents/openApi/1099
func (b *Shop) ChainStoreCreate(params map[string]interface{}) (string, error) {
	return b.Base.WithRequestParams(fn2.CHAIN_STORE_CREATE_URL, params).Result()
}

// ChainStoreUpdate returns update store visit: https://open.ele.me/documents/openApi/1107
func (b *Shop) ChainStoreUpdate(params map[string]interface{}) (string, error) {
	return b.Base.WithRequestParams(fn2.CHAIN_STORE_UPDATE_URL, params).Result()
}

// ChainStoreQuery returns query store visit: https://open.ele.me/documents/openApi/1115
func (b *Shop) ChainStoreQuery(params map[string]interface{}) (string, error) {
	return b.Base.WithRequestParams(fn2.CHAIN_STORE_QUERY_URL, params).Result()
}

// ChainStoreQueryList returns query store list visit: https://open.ele.me/documents/openApi/1123
func (b *Shop) ChainStoreQueryList(params map[string]interface{}) (string, error) {
	return b.Base.WithRequestParams(fn2.CHAIN_STORE_QUERY_LIST_URL, params).Result()
}

// ChainStoreRange returns store range visit: https://open.ele.me/documents/openApi/1131
func (b *Shop) ChainStoreRange(params map[string]interface{}) (string, error) {
	return b.Base.WithRequestParams(fn2.CHAIN_STORE_RANGE_URL, params).Result()
}

// GetAvailableExtraGoodsList return available extra goods list visit: https://open.ele.me/documents/openApi/1604
func (b *Shop) GetAvailableExtraGoodsList(params map[string]interface{}) (string, error) {
	return b.Base.WithRequestParams(fn2.GET_AVAILABLE_EXTRA_GOODS_LIST_URL, params).Result()
}

// NewUploadFile return new upload file visit: https://open.ele.me/documents/openApi/1083
func (b *Shop) NewUploadFile(params map[string]interface{}) (string, error) {
	return b.Base.WithRequestParams(fn2.NEW_UPLOAD_FILE_URL, params).Result()
}

// GetFileInfo return get file info visit: URL_ADDRESS
func (b *Shop) GetFileInfo(params map[string]interface{}) (string, error) {
	return b.Base.WithRequestParams(fn2.GET_FILE_INFO_URL, params).Result()
}
