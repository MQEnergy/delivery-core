package uu

const (
	ORDER_FEE_URL       = "/v2_0/getorderprice.ashx"     // 计算订单费用u
	CREATE_ORDER_URL    = "/v2_0/addorder.ashx"          // 新增配送单u
	CITY_CODE_LIST_URL  = "/v2_0/citylist.ashx"          // 城市信息列表u
	DEDUCT_FEE_URL      = "/v2_0/gethomeservicefee.ashx" // 获得违约金
	SHOP_DETAIL_URL     = ""                             // 查询门店详情接口
	CANCEL_REASONS_URL  = ""                             // 获取订单取消理由列表
	ORDER_DETAIL_URL    = "/v2_0/getorderdetail.ashx"    // 订单详情查询u
	CANCEL_ORDER_URL    = "/v2_0/cancelorder.ashx"       // 取消订单(线上环境)u
	RE_CREATE_ORDER_URL = ""                             // 重新发布订单
	MESSAGE_CONFIRM_URL = ""                             // 消息确认
	ADD_TIP_URL         = "/v2_0/payonlinefee.ashx"      // 增加小费u
	ORDER_LOGISTICS_URL = "/v2_0/getorderdetail.ashx"    // 订单骑手轨迹
)
