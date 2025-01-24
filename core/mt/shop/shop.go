package shop

import (
	"github.com/MQEnergy/delivery-core/core/mt"
	mt2 "github.com/MQEnergy/delivery-core/routers/mt"
)

type Shop struct {
	Base *mt.Base
}

func New(appKey, appSecret string) *Shop {
	return &Shop{
		Base: mt.New(appKey, appSecret),
	}
}

// ShopCreate 店铺创建 visit: https://peisong.meituan.com/tscc/docNew#%E5%88%9B%E5%BB%BA%E9%97%A8%E5%BA%97
func (s *Shop) ShopCreate(params map[string]interface{}) (string, error) {
	return s.Base.WithRequestParams(mt2.SHOP_CREATE_URL, params).Result()
}

// ShopUpdate 店铺更新 visit: https://peisong.meituan.com/tscc/docNew#%E4%BF%AE%E6%94%B9%E9%97%A8%E5%BA%97
func (s *Shop) ShopUpdate(params map[string]interface{}) (string, error) {
	return s.Base.WithRequestParams(mt2.SHOP_UPDATE_URL, params).Result()
}

// ShopQuery 店铺查询 visit: https://peisong.meituan.com/tscc/docNew#%E6%9F%A5%E8%AF%A2%E9%97%A8%E5%BA%97%E4%BF%A1%E6%81%AF
func (s *Shop) ShopQuery(params map[string]interface{}) (string, error) {
	return s.Base.WithRequestParams(mt2.SHOP_QUERY_URL, params).Result()
}

// ShopAreaQuery 查询门店配送范围 visit: https://peisong.meituan.com/tscc/docNew#%E6%9F%A5%E8%AF%A2%E9%97%A8%E5%BA%97%E9%85%8D%E9%80%81%E8%8C%83%E5%9B%B4
func (s *Shop) ShopAreaQuery(params map[string]interface{}) (string, error) {
	return s.Base.WithRequestParams(mt2.SHOP_AREA_QUERY_URL, params).Result()
}

// ShopBalanceQuery 查询门店账户余额 visit: https://peisong.meituan.com/tscc/docNew#%E6%9F%A5%E8%AF%A2%E9%97%A8%E5%BA%97%E8%B4%A6%E6%88%B7%E4%BD%99%E9%A2%9D
func (s *Shop) ShopBalanceQuery(params map[string]interface{}) (string, error) {
	return s.Base.WithRequestParams(mt2.SHOP_BALANCE_QUERY_URL, params).Result()
}
