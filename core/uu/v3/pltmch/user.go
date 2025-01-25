package pltmch

import (
	uu3 "github.com/MQEnergy/delivery-core/routers/uu/v3"
)

// UserAccount ...
// 平台服务商 visit: https://open.uupt.com/#/ApiDocument/v3/InterfaceList/account?PlatType=PLAT_MERCHANT&version=3&t=%E8%8E%B7%E5%8F%96%E8%B4%A6%E6%88%B7%E4%BD%99%E9%A2%9D&index=1-5-10
// 自营商家 visit: https://open.uupt.com/#/ApiDocument/v3/InterfaceList/account?PlatType=SELF_MERCHANT&version=3&t=%E8%8E%B7%E5%8F%96%E8%B4%A6%E6%88%B7%E4%BD%99%E9%A2%9D&index=2-5-10
func (o *PltMch) UserAccount(params map[string]interface{}) (string, error) {
	return o.Base.WithRequestParams(uu3.USER_ACCOUNT_URL, params).Result()
}

// UserRecharge ...
func (o *PltMch) UserRecharge(params map[string]interface{}) (string, error) {
	return o.Base.WithRequestParams(uu3.USER_RECHARGE_URL, params).Result()
}

// SendSmsCode ...
func (o *PltMch) SendSmsCode(params map[string]interface{}) (string, error) {
	return o.Base.WithRequestParams(uu3.SEND_SMS_CODE_URL, params).Result()
}

// UserAuth ...
func (o *PltMch) UserAuth(params map[string]interface{}) (string, error) {
	return o.Base.WithRequestParams(uu3.USER_AUTH_URL, params).Result()
}

// UserAuthUnbind ...
func (o *PltMch) UserAuthUnbind(params map[string]interface{}) (string, error) {
	return o.Base.WithRequestParams(uu3.USER_AUTH_UNBIND_URL, params).Result()
}
