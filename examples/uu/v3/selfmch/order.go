package main

import (
	"fmt"
	"github.com/MQEnergy/delivery-core/core/uu/v3/selfmch"
)

func main() {
	o := selfmch.New("ccba8bd4a2d54a2fb6df97e87979f303", "2815a7a1f8e3405d81fd6263683ec4e7", "910a0dfd12bb4bc0acec147bcb1ae246", false)
	cityList, err := o.OpenCityList(map[string]interface{}{})
	fmt.Println(cityList, err)
}
