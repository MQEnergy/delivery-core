package balance

import (
	"github.com/MQEnergy/delivery-core/core/dada"
	dada2 "github.com/MQEnergy/delivery-core/routers/dada"
)

type Balance struct {
	Base *dada.Base
}

func New(sourceID, appKey, appSecret string, online bool) *Balance {
	return &Balance{
		Base: dada.New(sourceID, appKey, appSecret, online),
	}
}

// Recharge ... visit: https://newopen.imdada.cn/#/development/file/recharge
func (b *Balance) Recharge(params map[string]interface{}) (string, error) {
	return b.Base.WithRequestParams(dada2.RECHARGE_URL, params).Result()
}

// BalanceQuery query balance visit: https://newopen.imdada.cn/#/development/file/balanceQuery
func (b *Balance) BalanceQuery(params map[string]interface{}) (string, error) {
	return b.Base.WithRequestParams(dada2.BALANCE_QUERY_URL, params).Result()
}
