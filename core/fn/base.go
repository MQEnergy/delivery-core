package fn

import (
	"crypto/md5"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"math/rand"
	"net/http"
	"net/url"
	"strings"
	"time"
)

// visit: https://open.ele.me/documents/openApi/1516
const (
	OnlineUrl   = "https://open-anubis.ele.me/anubis-webapi"
	NoOnlineUrl = "https://exam-anubis.ele.me/anubis-webapi"
	Timeout     = 5000 // 超时时间 5000ms
)

type Base struct {
	Client      *http.Client // Change this to your actual client type
	AppID       string
	AppSecret   string
	Salt        int
	Host        string
	URL         string
	RequestData io.Reader // 请求参数map
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

// SetUrl sets the URL attribute
func (b *Base) SetUrl(urlString string) {
	b.URL = b.Host + urlString
}

// Result 获取结果
func (b *Base) Result() (string, error) {
	req, err := b.Client.Post(b.Host+b.URL, "application/json", b.RequestData)
	if err != nil {
		return "", err
	}
	defer req.Body.Close()
	result, err := ioutil.ReadAll(req.Body)
	if err != nil {
		return "", err
	}
	return string(result), nil
}

// WithRequestParams ...
func (b *Base) WithRequestParams(_url string, data map[string]interface{}) *Base {
	body, _ := json.Marshal(data)
	params := map[string]interface{}{
		"app_id": b.AppID,
		"data":   url.QueryEscape(string(body)),
		"salt":   b.Salt,
	}
	// Set the signature
	params["signature"] = b.Sign(params)
	bytes, _ := json.Marshal(params)
	b.RequestData = strings.NewReader(string(bytes))
	b.URL = fmt.Sprintf("%s%s?app_id=%s&salt=%d&signature=%s", b.Host, _url, b.AppID, b.Salt, params["signature"])
	return b
}

// Sign generates a signature
func (b *Base) Sign(data map[string]interface{}) string {
	jsonData, _ := json.Marshal(data)
	encodedData := url.QueryEscape(string(jsonData))

	signChar := fmt.Sprintf("app_id=%s&access_token=%s&data=%s&salt=%d", b.AppID, b.AccessTokenSign(), encodedData, b.Salt)
	hash := md5.Sum([]byte(signChar))
	return fmt.Sprintf("%x", hash)
}

// AccessTokenSign generates a signature for the access token
func (b *Base) AccessTokenSign() string {
	str := fmt.Sprintf("app_id=%s&salt=%d&secret_key=%s", b.AppID, b.Salt, b.AppSecret)
	hash := md5.Sum([]byte(url.QueryEscape(str)))
	return fmt.Sprintf("%x", hash)
}
