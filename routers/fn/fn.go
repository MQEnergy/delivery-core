package fn

const (
	GET_ACCESS_TOKEN_URL = "/anubis-webapi/get_access_token"     //获取access_token
	CREATE_ORDER_URL     = "/anubis-webapi/v2/order"             //创建订单
	CANCEL_ORDER_URL     = "/anubis-webapi/v2/order/cancel"      //取消订单
	ORDER_DETAIL_URL     = "/anubis-webapi/v2/order/query"       //订单详情
	SHOP_DETAIL_URL      = "/anubis-webapi/v2/chain_store/query" //查询门店详情
	ADD_TIP_URL          = ""                                    //加小费
	ORDER_LOGISTICS_URL  = "/anubis-webapi/v2/order/carrier"     //订单骑手轨迹
)
