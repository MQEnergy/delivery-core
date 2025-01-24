package auth

import (
	"fmt"
	"github.com/MQEnergy/delivery-core/core/sf"
)

type Auth struct {
	Base *sf.Base
}

func New(devID, devKey string) *Auth {
	return &Auth{
		Base: sf.New(devID, devKey),
	}
}

// GetBizAuthURL 商户在线授权开发者链接 适用于个人在线授权开发者链接
// out_shop_id 第三方店铺ID 使用第三方店铺ID需传入（店铺授权）默认为空
func (a *Auth) GetBizAuthURL(outShopID string) string {
	return fmt.Sprintf("https://openic.sf-express.com/artascope/cx/receipt/getpage/product/artascope/page/storeBinding?dev_id=%s&out_shop_id=%s", a.Base.DevID, outShopID)
}
