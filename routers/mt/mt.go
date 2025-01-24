package mt

const (
	PRE_CREATE_BY_SHOP_URL = "/order/preCreateByShop" // 发单前预览
	CREATE_BY_SHOP_URL     = "/order/createByShop"    // 创建订单
	STATUS_QUERY_URL       = "/order/status/query"    // 查询订单状态
	RIDER_LOCATION_URL     = "/order/rider/location"  // 获取骑手位置
	ORDER_DELETE_URL       = "/order/delete"          // 取消订单
	ADD_TIP_URL            = "/order/addTip"          // 为骑手增加小费
	EVALUATE_URL           = "/order/evaluate"        // 评价骑手服务
	IMAGE_UPLOAD_URL       = "/file/image/upload"     // 图片上传
	SHOP_CREATE_URL        = "/shop/create"           // 创建门店
	SHOP_UPDATE_URL        = "/shop/update"           // 更新门店
	SHOP_QUERY_URL         = "/shop/query"            // 查询门店
	SHOP_AREA_QUERY_URL    = "/shop/area/query"       // 查询门店配送范围
	SHOP_BALANCE_QUERY_URL = "/shop/balance/query"    // 查询门店账户余额
)
