package mt

import (
	"crypto/sha1"
	"encoding/hex"
	"github.com/spf13/cast"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"sort"
	"strings"
	"time"
)

const (
	BaseUrl     = "https://peisongopen.meituan.com/api" // 线上地址
	Timeout     = 5000                                  // 超时时间 5000ms
	ContentType = "application/x-www-form-urlencoded"
	Version     = "1.0" // 版本号
)

// Base represents the main structure for Dada delivery service.
type Base struct {
	Client      *http.Client // Assuming client type is defined elsewhere
	AppKey      string
	AppSecret   string
	Host        string
	URL         string
	RequestData io.Reader // 请求参数map
}

// New is the constructor for the Base struct
func New(appKey, appSecret string) *Base {
	base := &Base{
		AppKey:    appKey,
		AppSecret: appSecret,
	}
	base.Host = BaseUrl

	client := &http.Client{
		Timeout: time.Millisecond * Timeout, // Set 3000ms timeout.
	}
	base.Client = client
	return base
}

// Result 获取结果
func (b *Base) Result() (string, error) {
	req, err := b.Client.Post(b.Host+b.URL, ContentType, b.RequestData)
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
	data["appkey"] = b.AppKey
	data["timestamp"] = time.Now().Unix()
	data["version"] = Version
	data["sign"] = b.Sign(data)
	args := url.Values{}
	for k, v := range data {
		args.Set(k, cast.ToString(v))
	}
	result := args.Encode()
	b.RequestData = strings.NewReader(result)
	b.URL = _url
	return b
}

// Sign generates the MD5 signature.
func (b *Base) Sign(data map[string]interface{}) string {
	// 1. Sort the keys
	keys := make([]string, 0, len(data))
	for key := range data {
		keys = append(keys, key)
	}
	sort.Strings(keys)

	// 2. Concatenate the strings
	var args strings.Builder
	for _, key := range keys {
		args.WriteString(key)
		args.WriteString(cast.ToString(data[key]))
	}
	finalStr := b.AppSecret + args.String()
	// sha1
	hash := sha1.Sum([]byte(finalStr))
	signature := strings.ToLower(hex.EncodeToString(hash[:]))
	return signature
}
