package sf

const (
	CREATE_ORDER_URL             = "/open/api/external/createorder"           // 创建订单（店铺）
	CREATE_ORDER4C_URL           = "/open/api/external/createorder4c"         // 创建订单（企业用户）
	PRE_CREATE_ORDER_URL         = "/open/api/external/precreateorder"        // 预创建订单（店铺）
	PRE_CREATE_ORDER4C_URL       = "/open/api/external/precreateorder4c"      // 预创建订单（企业用户）
	CANCEL_ORDER_URL             = "/open/api/external/cancelorder"           // 取消订单
	PRE_CANCEL_ORDER_URL         = "/open/api/external/precancelorder"        // 预取消订单
	ADD_TIP_URL                  = "/open/api/external/addordergratuityfee"   // 订单加小费
	GET_ORDER_TIP_URL            = "/open/api/external/getordergratuityfee"   // 获取订单加小费信息
	LIST_ORDER_FEED_URL          = "/open/api/external/listorderfeed"         // 订单状态流查询
	GET_ORDER_STATUS_URL         = "/open/api/external/getorderstatus"        // 订单实时信息查询 注：暂时只支持60天内的数据查询。
	REMINDER_ORDER_URL           = "/open/api/external/reminderorder"         // 催单 配送状态中，可通过该接口发起催单
	GET_CHANGE_ORDER_FEE_URL     = "/open/api/external/getchangeorderfee"     // 获取改单金额
	CHANGE_ORDER_URL             = "/open/api/external/changeorder"           // 改单
	ORDER_LOGISTICS_URL          = "/open/api/external/riderlatestposition"   // 获取配送员实时坐标接口
	RIDER_VIEWV2_URL             = "/open/api/external/riderviewv2"           // 获取配送员轨迹H5
	NOTIFY_PRODUCT_READY_URL     = "/open/api/external/notifyproductready"    // 商家告知餐品制作完成接口
	GET_CALLBACK_INFO_URL        = "/open/api/external/getcallbackinfo"       // 订单回调详情
	GET_SHOP_ACCOUNT_BALANCE_URL = "/open/api/external/getshopaccountbalance" // 获取店铺账户余额
	GET_COMPANY_INFO_URL         = "/open/api/external/getcompanyinfo"        // 获取企业信息
	GET_SHOP_INFO_URL            = "/open/api/external/getshopinfo"           // 获取商户信息
)
