package v2

const (
	GET_ACCESS_TOKEN_URL          = "/get_access_token"             // 获取access_token
	CREATE_ORDER_URL              = "/v2/order"                     // 创建订单
	CANCEL_ORDER_URL              = "/v2/order/cancel"              // 取消订单
	ORDER_DETAIL_URL              = "/v2/order/query"               // 订单详情
	COMPLAINT_ORDER_URL           = "/v2/order/complaint"           // 订单投诉
	CHAIN_STORE_URL               = "/v2/chain_store"               // 查询门店详情
	CHAIN_STORE_QUERY_URL         = "/v2/chain_store/query"         // 查询门店详情
	CHAIN_STORE_UPDATE_URL        = "/v2/chain_store/update"        // 查询门店列表
	CHAIN_STORE_DELIVERY_AREA_URL = "/v2/chain_store/delivery_area" // 查询门店配送范围
	ADD_TIP_URL                   = ""                              // 加小费
	ORDER_LOGISTICS_URL           = "/v2/order/carrier"             // 订单骑手轨迹
	ORDER_CARRIER_ROUTE_URL       = "/v2/order/carrier_route"       // 查询配送轨迹
)
