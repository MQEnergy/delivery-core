package uu

const (
	ORDER_FEE_URL          = "/v2_0/getorderprice.ashx"     // 计算订单价格
	CREATE_ORDER_URL       = "/v2_0/addorder.ashx"          // 发布订单
	ORDER_DETAIL_URL       = "/v2_0/getorderdetail.ashx"    // 获取订单详情
	CITY_CODE_LIST_URL     = "/v2_0/citylist.ashx"          // 已开通城市区县列表
	ADD_TIP_URL            = "/v2_0/payonlinefee.ashx"      // 支付小费
	DEDUCT_FEE_URL         = "/v2_0/gethomeservicefee.ashx" // 获得违约金
	CANCEL_ORDER_URL       = "/v2_0/cancelorder.ashx"       // 取消订单
	GET_BALANCE_DETAIL_URL = "/v2_0/getbalancedetail.ashx"  // 获取余额详情
	GET_RECHARGE_URL       = "/v2_0/getrecharge.ashx"       // 获取充值地址
	BIND_USER_APPLY_URL    = "/v2_0/binduserapply.ashx"     // 发送短信验证码
	BIND_USER_SUBMIT_URL   = "/v2_0/bindusersubmit.ashx"    // 获取用户openid
	CANCEL_BIND_URL        = "/v2_0/cancelbind.ashx"        // 用户解除绑定
	GET_SHOP_LIST_URL      = "/v2_0/getshoplist.ashx"       // 获取门店列表

	//SHOP_DETAIL_URL        = ""                             // 查询门店详情接口
	//CANCEL_REASONS_URL     = ""                             // 获取订单取消理由列表
	//RE_CREATE_ORDER_URL    = ""                             // 重新发布订单
	//MESSAGE_CONFIRM_URL    = ""                             // 消息确认
)
