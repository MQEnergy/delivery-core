package shop

import (
	"github.com/MQEnergy/delivery-core/core/sf"
	sf2 "github.com/MQEnergy/delivery-core/routers/sf"
)

type Shop struct {
	Base *sf.Base
}

func New(devID, devKey string) *Shop {
	return &Shop{
		Base: sf.New(devID, devKey),
	}
}

func (s *Shop) GetShopAccountBalance(params map[string]interface{}) (string, error) {
	return s.Base.WithRequestParams(sf2.GET_SHOP_ACCOUNT_BALANCE_URL, params).Result()
}
func (s *Shop) GetCompanyInfo(params map[string]interface{}) (string, error) {
	return s.Base.WithRequestParams(sf2.GET_COMPANY_INFO_URL, params).Result()
}
func (s *Shop) GetShopInfo(params map[string]interface{}) (string, error) {
	return s.Base.WithRequestParams(sf2.GET_SHOP_INFO_URL, params).Result()
}
