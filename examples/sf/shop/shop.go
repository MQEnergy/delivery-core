package main

import (
	"fmt"
	"github.com/MQEnergy/delivery-core/core/sf/shop"
)

func main() {
	s := shop.New("", "")
	companyInfo, err := s.GetShopInfo(map[string]interface{}{
		"shop_id":   "",
		"shop_type": "2",
	})
	fmt.Println(companyInfo, err)
}
