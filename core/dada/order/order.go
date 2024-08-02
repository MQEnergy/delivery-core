package order

import (
	"github.com/MQEnergy/delivery-core/core/dada"
	dada2 "github.com/MQEnergy/delivery-core/routers/dada"
)

type Order struct {
	Base *dada.Base
}

func New(sourceID, appKey, appSecret string, online bool) *Order {
	return &Order{
		Base: dada.New(sourceID, appKey, appSecret, online),
	}
}

// AddOrder creates a new order visit: https://newopen.imdada.cn/#/development/file/add
func (o *Order) AddOrder(params map[string]interface{}) (string, error) {
	return o.Base.WithRequestParams(dada2.ADD_ORDER_URL, params).Result()
}

// ReAddOrder re-create a new order visit: https://newopen.imdada.cn/#/development/file/reAdd
func (o *Order) ReAddOrder(params map[string]interface{}) (string, error) {
	return o.Base.WithRequestParams(dada2.RE_ADD_ORDER_URL, params).Result()
}

// QueryDeliverFee query deliver fee visit: https://newopen.imdada.cn/#/development/file/readyAdd
func (o *Order) QueryDeliverFee(params map[string]interface{}) (string, error) {
	return o.Base.WithRequestParams(dada2.QUERY_DELIVER_FEE_URL, params).Result()
}

// AddAfterQuery add after query visit: https://newopen.imdada.cn/#/development/file/addAfterQuery
func (o *Order) AddAfterQuery(params map[string]interface{}) (string, error) {
	return o.Base.WithRequestParams(dada2.ADD_AFTER_QUERY_URL, params).Result()
}

// AddTip add tip visit: https://newopen.imdada.cn/#/development/file/addTip
func (o *Order) AddTip(params map[string]interface{}) (string, error) {
	return o.Base.WithRequestParams(dada2.ADD_TIP_URL, params).Result()
}

// OrderQuery order query visit: https://newopen.imdada.cn/#/development/file/statusQuery
func (o *Order) OrderQuery(params map[string]interface{}) (string, error) {
	return o.Base.WithRequestParams(dada2.ORDER_QUERY_URL, params).Result()
}

// CancelOrder cancel order visit: https://newopen.imdada.cn/#/development/file/statusQuery
func (o *Order) CancelOrder(params map[string]interface{}) (string, error) {
	return o.Base.WithRequestParams(dada2.CANCEL_ORDER_URL, params).Result()
}

// OrderPosition order position visit: https://newopen.imdada.cn/#/development/file/queryLocation
func (o *Order) OrderPosition(params map[string]interface{}) (string, error) {
	return o.Base.WithRequestParams(dada2.ORDER_POSITION_URL, params).Result()
}
