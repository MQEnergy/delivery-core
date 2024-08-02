package dada

const (
	ADD_ORDER_URL         = "/api/order/addOrder"             // 新增配送单
	RE_ADD_ORDER_URL      = "/api/order/reAddOrder"           // 重新发布订单
	QUERY_DELIVER_FEE_URL = "/api/order/queryDeliverFee"      // 查询运费
	ADD_AFTER_QUERY_URL   = "/api/order/addAfterQuery"        // 查询运费后下单
	ADD_TIP_URL           = "/api/order/addTip"               // 增加小费
	ORDER_QUERY_URL       = "/api/order/status/query"         // 查询订单详情
	CANCEL_ORDER_URL      = "/api/order/formalCancel"         // 取消订单(线上环境)
	ORDER_POSITION_URL    = "/api/order/transporter/position" // 查询骑士位置
	CITY_CODE_LIST_URL    = "/api/cityCode/list"              // 查询城市码
	SHOP_ADD_URL          = "/api/shop/add"                   // 创建门店
	SHOP_UPDATE_URL       = "/api/shop/update"                // 编辑门店
	SHOP_DETAIL_URL       = "/api/shop/detail"                // 查询门店详情接口
	RECHARGE_URL          = "/api/recharge"                   // 获取充值链接
	BALANCE_QUERY_URL     = "/api/balance/query"              // 查询账户余额
)
