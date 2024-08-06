package main

import (
	"fmt"
	v3 "github.com/MQEnergy/delivery-core/core/fn/v3"
	v32 "github.com/MQEnergy/delivery-core/routers/fn/v3"
	"github.com/spf13/cast"
	"io/ioutil"
	"strings"
	"time"
)

func main() {
	o := v3.New("", "", false)
	sign, err := o.GenerateSign(map[string]string{
		"app_id":      "",
		"code":        "",
		"merchant_id": "",
		"grant_type":  "authorization_code",
		"timestamp":   cast.ToString(time.Now().UnixMilli()),
	})
	if err != nil {
		panic(err)
	}
	fmt.Println(sign, o.Host+v32.GET_TOKEN_URL)
	response, err := o.Client.Post(o.Host+v32.GET_TOKEN_URL, "application/json", strings.NewReader(`{
		 "grant_type": "authorization_code",
		 "app_id": "",
		 "code": "",
		 "merchant_id": "",
		 "signature": "`+sign+`",
		 "timestamp": "`+cast.ToString(time.Now().UnixNano()/int64(time.Millisecond))+`"
	}`))
	if err != nil {
		panic(err)
	}
	defer response.Body.Close()
	body1, _ := ioutil.ReadAll(response.Body)
	fmt.Println(string(body1))

	resp, err := o.Client.Post(o.Host+v32.CHAIN_STORE_QUERY_URL, "application/json", strings.NewReader(`{
		"app_id": "",
		"merchant_id": "",
		"timestamp": "`+cast.ToString(time.Now().UnixNano())+`",
		"business_data": "{\"merchant_id\":\"4434923\",\"chain_store_code\":\"Z3ZB\"}",
		"signature": "`+sign+`",
		"access_token": ""
	}`))
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)
	fmt.Println(string(body))
}
