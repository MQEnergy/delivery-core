package v3

const (
	GET_TOKEN_URL                      = "/openapi/token"                        // 获取token接口
	REFRESH_TOKEN_URL                  = "/openapi/refreshToken"                 // 刷新token接口
	PRE_CREATE_ORDER_URL               = "/v3/invoke/preCreateOrder"             // 预下单接口
	CREATE_ORDER_URL                   = "/v3/invoke/createOrder"                // 创建订单接口
	ADD_TIP_URL                        = "/v3/invoke/addTip"                     // 加小费
	GET_CANCEL_REASON_LIST_URL         = "/v3/invoke/getCancelReasonList"        // 获取可用订单取消原因接口
	PRE_CANCEL_ORDER_URL               = "/v3/invoke/preCancelOrder"             // 订单预取消接口
	CANCEL_ORDER_URL                   = "/v3/invoke/cancelOrder"                // 正式订单取消接口
	RUSH_ORDER_URL                     = "/v3/invoke/rushOrder"                  // 催单接口
	ORDER_DETAIL_URL                   = "/v3/invoke/getOrderDetail"             // 查询订单详情接口
	GET_KNIGHT_INFO_URL                = "/v3/invoke/getKnightInfo"              // 查询骑手信息接口（订单骑手轨迹）
	COMPLAINT_ORDER_URL                = "/v3/invoke/complaintOrder"             // 订单投诉接口
	CLAIM_ORDER_URL                    = "/v3/invoke/claimOrder"                 // 订单索赔接口
	GET_AMOUNT_URL                     = "/v3/invoke/getAmount"                  // 获取余额接口
	CHAIN_STORE_CREATE_URL             = "/v3/invoke/chainstoreCreate"           // 门店创建接口
	CHAIN_STORE_UPDATE_URL             = "/v3/invoke/chainstoreUpdate"           // 门店更新接口
	CHAIN_STORE_QUERY_URL              = "/v3/invoke/chainstoreQuery"            // 门店详情查询接口
	CHAIN_STORE_QUERY_LIST_URL         = "/v3/invoke/chainstoreQueryList"        // 门店批量查询接口
	CHAIN_STORE_RANGE_URL              = "/v3/invoke/chainstoreRange"            // 门店配送范围接口
	GET_AVAILABLE_EXTRA_GOODS_LIST_URL = "/v3/invoke/getAvailableExtraGoodsList" // 获取增值服务
	NEW_UPLOAD_FILE_URL                = "/v3/invoke/newUploadFile"              // 文件上传接口
	GET_FILE_INFO_URL                  = "/v3/invoke/getFileInfo"                // 文件信息查询接口
)
