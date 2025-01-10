package main

import (
	"fmt"
	v3 "github.com/MQEnergy/delivery-core/core/fn/v3"
	"github.com/MQEnergy/delivery-core/core/fn/v3/shop"
)

func main() {
	// 获取token
	b := v3.New("", "", "4434923", false)
	// 获取token 沙箱环境code为空
	result, err := b.GetToken(v3.SignParams{Code: ""})
	if err != nil {
		panic(err)
	}
	fmt.Println(result)

	// 获取门店信息
	o := shop.New("", "", "4434923", false)
	storeQuery, err := o.ChainStoreQuery(map[string]interface{}{
		"out_shop_code":  "testorder001",
		"chain_store_id": 205318427,
	}, result.AccessToken)
	if err != nil {
		panic(err)
	}
	fmt.Println(storeQuery)
	return
}
