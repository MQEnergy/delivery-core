package main

import (
	"fmt"
	"github.com/MQEnergy/delivery-core/core/uu/v2/order"
)

func main() {
	o := order.New("ccba8bd4a2d54a2fb6df97e87979f303", "2815a7a1f8e3405d81fd6263683ec4e7", "910a0dfd12bb4bc0acec147bcb1ae246", false)
	shopList, err := o.GetShopList(map[string]interface{}{
		"pageindex": 1,
	})
	fmt.Println(shopList, err)
}
