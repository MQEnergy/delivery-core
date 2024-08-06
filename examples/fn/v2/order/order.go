package main

import (
	"fmt"
	"github.com/MQEnergy/delivery-core/core/fn/v2/shop"
)

func main() {
	o := shop.New("", "", true)
	token, _, _ := o.Base.GetAccessToken()
	orderDetail, err := o.ChainStoreQuery(map[string]interface{}{
		"chain_store_code": []string{"10006989"},
	}, token)
	fmt.Println(orderDetail, err)
}
