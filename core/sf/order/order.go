package order

import (
	"github.com/MQEnergy/delivery-core/core/sf"
	sf2 "github.com/MQEnergy/delivery-core/routers/sf"
)

type Order struct {
	Base *sf.Base
}

func New(devID, devKey string) *Order {
	return &Order{
		Base: sf.New(devID, devKey),
	}
}

// CreateOrder 创建订单（店铺）visit: https://openic.sf-express.com/#/apidoc
func (o *Order) CreateOrder(params map[string]interface{}) (string, error) {
	return o.Base.WithRequestParams(sf2.CREATE_ORDER_URL, params).Result()
}
func (o *Order) CreateOrder4C(params map[string]interface{}) (string, error) {
	return o.Base.WithRequestParams(sf2.CREATE_ORDER4C_URL, params).Result()
}
func (o *Order) PreCreateOrder(params map[string]interface{}) (string, error) {
	return o.Base.WithRequestParams(sf2.PRE_CREATE_ORDER_URL, params).Result()
}
func (o *Order) PreCreateOrder4C(params map[string]interface{}) (string, error) {
	return o.Base.WithRequestParams(sf2.PRE_CREATE_ORDER4C_URL, params).Result()
}
func (o *Order) CancelOrder(params map[string]interface{}) (string, error) {
	return o.Base.WithRequestParams(sf2.CANCEL_ORDER_URL, params).Result()
}
func (o *Order) PreCancelOrder(params map[string]interface{}) (string, error) {
	return o.Base.WithRequestParams(sf2.PRE_CANCEL_ORDER_URL, params).Result()
}
func (o *Order) AddTip(params map[string]interface{}) (string, error) {
	return o.Base.WithRequestParams(sf2.ADD_TIP_URL, params).Result()
}
func (o *Order) GetOrderTip(params map[string]interface{}) (string, error) {
	return o.Base.WithRequestParams(sf2.GET_ORDER_TIP_URL, params).Result()
}
func (o *Order) ListOrderFeed(params map[string]interface{}) (string, error) {
	return o.Base.WithRequestParams(sf2.LIST_ORDER_FEED_URL, params).Result()
}
func (o *Order) GetOrderStatus(params map[string]interface{}) (string, error) {
	return o.Base.WithRequestParams(sf2.GET_ORDER_STATUS_URL, params).Result()
}
func (o *Order) ReminderOrder(params map[string]interface{}) (string, error) {
	return o.Base.WithRequestParams(sf2.REMINDER_ORDER_URL, params).Result()
}
func (o *Order) GetChangeOrderFee(params map[string]interface{}) (string, error) {
	return o.Base.WithRequestParams(sf2.GET_CHANGE_ORDER_FEE_URL, params).Result()
}
func (o *Order) ChangeOrder(params map[string]interface{}) (string, error) {
	return o.Base.WithRequestParams(sf2.CHANGE_ORDER_URL, params).Result()
}
func (o *Order) OrderLogistics(params map[string]interface{}) (string, error) {
	return o.Base.WithRequestParams(sf2.ORDER_LOGISTICS_URL, params).Result()
}
func (o *Order) RiderViewV2(params map[string]interface{}) (string, error) {
	return o.Base.WithRequestParams(sf2.RIDER_VIEWV2_URL, params).Result()
}
func (o *Order) NotifyProductReady(params map[string]interface{}) (string, error) {
	return o.Base.WithRequestParams(sf2.NOTIFY_PRODUCT_READY_URL, params).Result()
}
func (o *Order) GetCallbackInfo(params map[string]interface{}) (string, error) {
	return o.Base.WithRequestParams(sf2.GET_CALLBACK_INFO_URL, params).Result()
}
