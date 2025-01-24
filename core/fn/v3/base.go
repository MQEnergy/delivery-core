package v3

import (
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"fmt"
	fn3 "github.com/MQEnergy/delivery-core/routers/fn/v3"
	"github.com/spf13/cast"
	"github.com/tidwall/gjson"
	"io/ioutil"
	"net/http"
	"reflect"
	"sort"
	"strings"
	"time"
)

// visit: https://open.ele.me/documents/ka/1419
const (
	OnlineUrl   = "https://open-anubis.ele.me/anubis-webapi"
	NoOnlineUrl = "https://exam-anubis.ele.me/anubis-webapi"
	Timeout     = 5000 // 超时时间 5000ms
	ContentType = "application/json"
	Version     = "1.0"
)

type Base struct {
	Client      *http.Client // Change this to your actual client type
	AppID       string
	AppSecret   string
	MerchantID  string
	Host        string
	URL         string
	Online      bool   // true: 正式 false: 沙箱
	RequestData string // 请求参数map
}

// SignParams token签名参数
type SignParams struct {
	GrantType    string `json:"grant_type"` // authorization_code: 获取code用 refresh_token: 刷新token用
	Code         string `json:"code"`       // code有效期24小时，且只能使用一次，用过即失效
	AppID        string `json:"app_id"`
	MerchantID   string `json:"merchant_id"`
	TimeStamp    string `json:"timestamp"`
	RefreshToken string `json:"refresh_token"` // 刷新token使用
}

// BizSignParams 业务数据签名参数
type BizSignParams struct {
	AppID        string `json:"app_id"`
	MerchantID   string `json:"merchant_id"`
	TimeStamp    string `json:"timestamp"`
	Version      string `json:"version"`
	AccessToken  string `json:"access_token"`
	BusinessData string `json:"business_data"`
}

type TokenBizDataInfo struct {
	AppID        string `json:"app_id"`
	MerchantID   string `json:"merchant_id"`
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
	ExpireIn     int    `json:"expire_in"`
	ReExpireIn   int    `json:"re_expire_in"`
}

// New is the constructor for the Base struct
func New(appID, appSecret, merchantID string, online bool) *Base {
	base := &Base{
		AppID:      appID,
		AppSecret:  appSecret,
		MerchantID: merchantID,
	}
	if online {
		base.Host = OnlineUrl
	} else {
		base.Host = NoOnlineUrl
	}
	client := &http.Client{
		Timeout: time.Millisecond * Timeout, // Set 5000ms timeout.
	}
	base.Online = online
	base.Client = client
	return base
}

// Result 获取结果
func (b *Base) Result() (string, error) {
	resp, err := b.Client.Post(b.Host+b.URL, ContentType, strings.NewReader(b.RequestData))
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()
	result, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}
	return string(result), nil
}

// WithRequestParams
// @Description: 设置业务请求参数
// @receiver b
// @param _url
// @param data 请求参数
// @param accessToken 请求token
// @return *Base
// @author cx
func (b *Base) WithRequestParams(_url string, params BizSignParams) *Base {
	if params.Version == "" {
		params.Version = Version
	}
	if params.TimeStamp == "" {
		params.TimeStamp = cast.ToString(time.Now().UnixMilli())
	}
	if params.AppID == "" {
		params.AppID = b.AppID
	}
	if params.MerchantID == "" {
		params.MerchantID = b.MerchantID
	}
	sign, err := b.GenerateBizSign(params)
	if err != nil {
		return b
	}
	bytes, _ := json.Marshal(map[string]interface{}{
		"app_id":        params.AppID,
		"merchant_id":   params.MerchantID,
		"timestamp":     params.TimeStamp,
		"version":       params.Version,
		"business_data": params.BusinessData,
		"signature":     sign,
		"access_token":  params.AccessToken,
	})
	b.RequestData = string(bytes)
	b.URL = _url
	return b
}

// WithTokenReqParams
// @Description: 设置token请求参数 access_token有效期365天
// sign加签参数：
// * grant_type=authorization_code
// * code=xxx
// * app_id=xxxxxxxx
// * merchant_id=12434354
// * timestamp=1609118272793
// @receiver b
// @param _url
// @param params
// @return *Base
// @author cx
func (b *Base) WithTokenReqParams(_url string, params SignParams) *Base {
	// Todo: 加token缓存
	if params.TimeStamp == "" {
		params.TimeStamp = cast.ToString(time.Now().UnixMilli())
	}
	if params.AppID == "" {
		params.AppID = b.AppID
	}
	if params.GrantType == "" {
		params.GrantType = "authorization_code"
	}
	if params.MerchantID == "" {
		params.MerchantID = b.MerchantID
	}
	sign, err := b.GenerateSign(params)
	if err != nil {
		return b
	}
	bytes, _ := json.Marshal(map[string]interface{}{
		"grant_type":  params.GrantType,
		"code":        params.Code,
		"app_id":      b.AppID,
		"merchant_id": params.MerchantID,
		"signature":   sign,
		"timestamp":   params.TimeStamp,
	})
	b.RequestData = string(bytes)
	b.URL = _url
	return b
}

// WithRefreshTokenReqParams
// @Description: 设置刷新token请求参数 refresh_token有效期365+7
// sign加签参数：（签名方式与GenerateSign一致）
// * grant_type=refresh_token
// * app_id=xxxxxxxx
// * merchant_id=12434354
// * timestamp=1609118272793
// * refresh_token=xxxxxxxxxxxxxx
// @receiver b
// @param _url
// @param params
// @return *Base
// @author cx
func (b *Base) WithRefreshTokenReqParams(_url string, params SignParams) *Base {
	// Todo: 加refresh_token缓存
	if params.TimeStamp == "" {
		params.TimeStamp = cast.ToString(time.Now().UnixMilli())
	}
	if params.AppID == "" {
		params.AppID = b.AppID
	}
	if params.GrantType == "" {
		params.GrantType = "refresh_token"
	}
	if params.MerchantID == "" {
		params.MerchantID = b.MerchantID
	}
	sign, err := b.GenerateSign(params)
	if err != nil {
		return b
	}
	bytes, _ := json.Marshal(map[string]interface{}{
		"grant_type":    params.GrantType,
		"refresh_token": params.RefreshToken,
		"app_id":        b.AppID,
		"merchant_id":   params.MerchantID,
		"signature":     sign,
		"timestamp":     params.TimeStamp,
	})
	b.RequestData = string(bytes)
	b.URL = _url
	return b
}

// GenerateBizSign
// @Description: 生成业务签名
// sign加签参数：（签名方式与GenerateSign一致）
// * app_id=xxxxxxxx
// * merchant_id=12434354
// * timestamp=1609118272793
// * version=1.0
// * access_token=xxxxxxxxxxxxxx
// * business_data={\"chain_store_id\":200008235,\"partner_order_code\"} //业务参数，json字符串
// @receiver b
// @param data
// @param accessToken
// @return string
// @author cx
func (b *Base) GenerateBizSign(params BizSignParams) (string, error) {
	// Step 1: 按照字典书序排序拼接参数
	var keys []string
	var paramStr string
	paramsMap := make(map[string]string)
	v := reflect.ValueOf(params)
	if v.Kind() == reflect.Ptr {
		v = v.Elem()
	}
	t := v.Type()
	for i := 0; i < v.NumField(); i++ {
		field := v.Field(i)
		if !field.IsZero() {
			paramsMap[t.Field(i).Tag.Get("json")] = field.String()
		}
	}
	for key := range paramsMap {
		keys = append(keys, key)
	}
	sort.Strings(keys)
	for i, key := range keys {
		if i == 0 {
			paramStr += fmt.Sprintf("%s=%s", key, paramsMap[key])
		} else {
			paramStr += fmt.Sprintf("&%s=%s", key, paramsMap[key])
		}
	}
	// Step 2: 拼接 appSecret
	signBefore := b.AppSecret + paramStr
	// Step 3: 使用 SHA-256 加密生成签名
	hash := sha256.New()
	_, err := hash.Write([]byte(signBefore))
	if err != nil {
		return "", err
	}
	signature := hash.Sum(nil)

	// 将字节数组转换为十六进制字符串
	return hex.EncodeToString(signature), nil
}

// GenerateSign 生成签名
func (b *Base) GenerateSign(params SignParams) (string, error) {
	// Step 1: 按照字典书序排序拼接参数
	var keys []string
	var paramStr string
	paramsMap := make(map[string]string)
	v := reflect.ValueOf(params)
	if v.Kind() == reflect.Ptr {
		v = v.Elem()
	}
	t := v.Type()
	for i := 0; i < v.NumField(); i++ {
		field := v.Field(i)
		if !field.IsZero() {
			paramsMap[t.Field(i).Tag.Get("json")] = field.String()
		}
	}
	for key := range paramsMap {
		keys = append(keys, key)
	}
	sort.Strings(keys)
	for i, key := range keys {
		if i == 0 {
			paramStr += fmt.Sprintf("%s=%s", key, paramsMap[key])
		} else {
			paramStr += fmt.Sprintf("&%s=%s", key, paramsMap[key])
		}
	}
	// Step 2: 拼接 appSecret
	signBefore := b.AppSecret + paramStr

	// Step 3: 使用 SHA-256 加密生成签名
	hash := sha256.New()
	_, err := hash.Write([]byte(signBefore))
	if err != nil {
		return "", err
	}
	signature := hash.Sum(nil)

	// 将字节数组转换为十六进制字符串
	return hex.EncodeToString(signature), nil
}

// GetToken returns token visit: https://open.ele.me/documents/openApi/603
func (b *Base) GetToken(params SignParams) (*TokenBizDataInfo, error) {
	result, err := b.WithTokenReqParams(fn3.GET_TOKEN_URL, params).Result()
	if err != nil {
		return nil, err
	}
	bizData := gjson.Get(result, "business_data").String()
	var tokenInfo TokenBizDataInfo
	if err := json.Unmarshal([]byte(bizData), &tokenInfo); err != nil {
		return nil, err
	}
	return &tokenInfo, nil
}

// RefreshToken returns token visit: https://open.ele.me/documents/openApi/603
func (b *Base) RefreshToken(params SignParams) (*TokenBizDataInfo, error) {
	result, err := b.WithTokenReqParams(fn3.GET_TOKEN_URL, params).Result()
	if err != nil {
		return nil, err
	}
	bizData := gjson.Get(result, "business_data").String()
	var tokenInfo TokenBizDataInfo
	if err := json.Unmarshal([]byte(bizData), &tokenInfo); err != nil {
		return nil, err
	}
	return &tokenInfo, nil
}
