package v2

import (
	"crypto/md5"
	"encoding/json"
	"fmt"
	routerv2 "github.com/MQEnergy/delivery-core/routers/fn/v2"
	"github.com/tidwall/gjson"
	"io/ioutil"
	"math/rand"
	"net/http"
	"net/url"
	"strings"
	"time"
)

// visit: https://open.ele.me/documents/ka/1419
const (
	OnlineUrl   = "https://open-anubis.ele.me/anubis-webapi"
	NoOnlineUrl = "https://exam-anubis.ele.me/anubis-webapi"
	Timeout     = 5000 // 超时时间 5000ms
)

type Base struct {
	Client         *http.Client // Change this to your actual client type
	AppID          string
	AppSecret      string
	Salt           int
	Host           string
	URL            string
	AccessTokenURL string
	RequestData    string // 请求参数map
}

// New is the constructor for the Base struct
func New(appID, appSecret string, online bool) *Base {
	base := &Base{
		AppID:     appID,
		AppSecret: appSecret,
	}
	rand.Seed(time.Now().UnixNano())
	base.Salt = rand.Intn(9000) + 1000

	if online {
		base.Host = OnlineUrl
	} else {
		base.Host = NoOnlineUrl
	}
	client := &http.Client{
		Timeout: time.Millisecond * Timeout, // Set 3000ms timeout.
	}
	base.Client = client
	return base
}

// Result 获取结果
func (b *Base) Result() (string, error) {
	resp, err := b.Client.Post(b.Host+b.URL, "application/json", strings.NewReader(b.RequestData))
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

func (b *Base) SetAccessTokenUrl(_url string) {
	b.URL = fmt.Sprintf("%s?app_id=%s&salt=%d&signature=%s", _url, b.AppID, b.Salt, b.AccessTokenSign())
}

// WithRequestParams ...
func (b *Base) WithRequestParams(_url string, data map[string]interface{}, accessToken string) *Base {
	body, _ := json.Marshal(data)
	params := map[string]interface{}{
		"app_id": b.AppID,
		"data":   url.QueryEscape(string(body)),
		"salt":   b.Salt,
	}
	// Set the signature
	params["signature"] = b.Sign(data, accessToken)
	bytes, _ := json.Marshal(params)
	b.RequestData = string(bytes)
	b.URL = _url
	return b
}

// Sign generates a signature
func (b *Base) Sign(data map[string]interface{}, accessToken string) string {
	jsonData, _ := json.Marshal(data)
	encodedData := url.QueryEscape(string(jsonData))
	signChar := fmt.Sprintf("app_id=%s&access_token=%s&data=%s&salt=%d", b.AppID, accessToken, encodedData, b.Salt)
	hash := md5.Sum([]byte(signChar))
	return fmt.Sprintf("%x", hash)
}

// AccessTokenSign generates a signature for the access token
func (b *Base) AccessTokenSign() string {
	str := fmt.Sprintf("app_id=%s&salt=%d&secret_key=%s", b.AppID, b.Salt, b.AppSecret)
	hash := md5.Sum([]byte(url.QueryEscape(str)))
	return fmt.Sprintf("%x", hash)
}

// GetAccessToken returns the access token visit: https://open.ele.me/documents/ka/1371
func (b *Base) GetAccessToken() (string, int64, error) {
	b.SetAccessTokenUrl(routerv2.GET_ACCESS_TOKEN_URL)
	resp, err := b.Client.Get(b.Host + b.URL)
	if err != nil {
		return "", 0, err
	}
	defer resp.Body.Close()
	result, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", 0, err
	}
	accessToken := gjson.Get(string(result), "data.access_token").String()
	expireTime := gjson.Get(string(result), "data.expire_time").Int()
	return accessToken, expireTime, nil
}
