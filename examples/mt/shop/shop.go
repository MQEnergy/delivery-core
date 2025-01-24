package main

import (
	"fmt"
	"github.com/MQEnergy/delivery-core/core/mt/shop"
)

func main() {
	s := shop.New("", "")
	shopQuery, err := s.ShopQuery(map[string]interface{}{
		"shop_id": "1111",
	})
	fmt.Println(shopQuery, err)
}
