package main

import (
	"fmt"
	"github.com/MQEnergy/delivery-core/core/dada/shop"
)

func main() {
	shopDetail, err := shop.
		New("", "", "", false).
		ShopDetail(map[string]interface{}{
			"origin_shop_id": "shop001",
		})
	fmt.Println(shopDetail, err)
}
