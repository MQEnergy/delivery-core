package main

import (
	"fmt"
	"github.com/MQEnergy/delivery-core/core/uu/v3/chainmch"
)

func main() {
	chainMch := chainmch.New("9e7b00d7a5d5406fb24f56709518bc6b", "6ba86f556a984c299b68dd43ba92bbaf", 6666705, "6666705-1354", false)
	cityList, err := chainMch.OpenCityList(map[string]interface{}{})
	fmt.Println(cityList, err)
}
