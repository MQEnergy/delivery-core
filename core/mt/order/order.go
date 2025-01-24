package order

import (
	"github.com/MQEnergy/delivery-core/core/mt"
	mt2 "github.com/MQEnergy/delivery-core/routers/mt"
)

type Order struct {
	Base *mt.Base
}

func New(appKey, appSecret string) *Order {
	return &Order{
		Base: mt.New(appKey, appSecret),
	}
}

// PreCreateByShop ... visit: https://peisong.meituan.com/tscc/docNew#%E5%8F%91%E5%8D%95%E5%89%8D%E9%A2%84%E8%A7%88
func (o *Order) PreCreateByShop(params map[string]interface{}) (string, error) {
	return o.Base.WithRequestParams(mt2.PRE_CREATE_BY_SHOP_URL, params).Result()
}

// CreateByShop creates a new order visit: https://peisong.meituan.com/tscc/docNew#%E5%8F%91%E5%8D%95
func (o *Order) CreateByShop(params map[string]interface{}) (string, error) {
	return o.Base.WithRequestParams(mt2.CREATE_BY_SHOP_URL, params).Result()
}

// StatusQuery ... visit: https://peisong.meituan.com/tscc/docNew#%E6%9F%A5%E8%AF%A2%E8%AE%A2%E5%8D%95%E7%8A%B6%E6%80%81
func (o *Order) StatusQuery(params map[string]interface{}) (string, error) {
	return o.Base.WithRequestParams(mt2.STATUS_QUERY_URL, params).Result()
}

// RiderLocation ... visit: https://peisong.meituan.com/tscc/docNew#%E8%8E%B7%E5%8F%96%E9%AA%91%E6%89%8B%E4%BD%8D%E7%BD%AE
func (o *Order) RiderLocation(params map[string]interface{}) (string, error) {
	return o.Base.WithRequestParams(mt2.RIDER_LOCATION_URL, params).Result()
}

// OrderDelete ... visit: https://peisong.meituan.com/tscc/docNew#%E5%8F%96%E6%B6%88%E8%AE%A2%E5%8D%95
func (o *Order) OrderDelete(params map[string]interface{}) (string, error) {
	return o.Base.WithRequestParams(mt2.ORDER_DELETE_URL, params).Result()
}

// AddTip ... visit: https://peisong.meituan.com/tscc/docNew#%E4%B8%BA%E9%AA%91%E6%89%8B%E5%A2%9E%E5%8A%A0%E5%B0%8F%E8%B4%B9
func (o *Order) AddTip(params map[string]interface{}) (string, error) {
	return o.Base.WithRequestParams(mt2.ADD_TIP_URL, params).Result()
}

// Evaluate ... visit: https://peisong.meituan.com/tscc/docNew#%E8%AF%84%E4%BB%B7%E9%AA%91%E6%89%8B%E6%9C%8D%E5%8A%A1
func (o *Order) Evaluate(params map[string]interface{}) (string, error) {
	return o.Base.WithRequestParams(mt2.EVALUATE_URL, params).Result()
}
