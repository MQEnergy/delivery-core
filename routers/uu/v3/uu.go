package v3

const (
	OPEN_CITY_LIST_URL       = "/v3/order/openCityList"            // 获取开放城市列表 适用于 [平台服务商，自营商家，连锁商家]
	ORDER_PRICE_URL          = "/v3/order/orderPrice"              // 获取订单价格 适用于 [平台服务商，自营商家]
	ADD_ORDER_URL            = "/v3/order/addOrder"                // 发布订单 适用于 [平台服务商，自营商家]
	ORDER_DETAIL_URL         = "/v3/order/orderDetail"             // 获取订单详情 适用于 [平台服务商，自营商家，连锁商家]
	CANCEL_FEE_URL           = "/v3/order/cancelFee"               // 获取取消费 适用于 [平台服务商，自营商家，连锁商家]
	CANCEL_ORDER_URL         = "/v3/order/cancelOrder"             // 取消订单 适用于 [平台服务商，自营商家，连锁商家]
	ADD_TIP_URL              = "/v3/order/addGratuity"             // 加小费 适用于 [平台服务商，自营商家，连锁商家]
	SYNC_PICKUP_CODE_URL     = "/v3/order/syncPickupCode"          // 同步取件码/收货码给骑手 适用于 [平台服务商，自营商家，连锁商家]
	USER_ACCOUNT_URL         = "/v3/user/account"                  // 获取账户余额 适用于 [平台服务商，自营商家]
	USER_RECHARGE_URL        = "/v3/user/recharge"                 // 获取充值URL 适用于 [平台服务商]
	SEND_SMS_CODE_URL        = "/v3/user/unauthorized/sendSmsCode" // 发送短信验证码 适用于 [平台服务商]
	USER_AUTH_URL            = "/v3/user/unauthorized/auth"        // 用户授权 适用于 [平台服务商]
	USER_AUTH_UNBIND_URL     = "/v3/user/unauthorized/unbind"      // 用户解除授权 适用于 [平台服务商]
	MCH_ORDER_PRICE_URL      = "/v3/order/merchantOrderPrice"      // 获取订单价格 适用于 [连锁商家]
	MCH_ADD_ORDER_URL        = "/v3/order/merchantAddOrder"        // 发布订单 适用于 [连锁商家]
	MCH_ADD_CHAIN_STORE_URL  = "/v3/store/addChainStore"           // 添加连锁门店 适用于 [连锁商家]
	MCH_EDIT_CHAIN_STORE_URL = "/v3/store/editChainStore"          // 编辑连锁门店 适用于 [连锁商家]

)
