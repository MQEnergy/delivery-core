package uu

import (
	"crypto/md5"
	"fmt"
	"github.com/spf13/cast"
	"io"
	"io/ioutil"
	"math/rand/v2"
	"net/http"
	"net/url"
	"sort"
	"strings"
	"time"
)

const (
	OnlineUrl   = "https://openapi.uupt.com"
	NoOnlineUrl = "http://openapi.test.uupt.com"
	Timeout     = 5000 // 超时时间 5000ms
	ContentType = "application/x-www-form-urlencoded"
)

type Base struct {
	Client      *http.Client // Change this to your actual client type
	AppID       string
	AppKey      string
	OpenID      string
	Host        string
	URL         string
	Online      bool      // true: 正式 false: 沙箱
	RequestData io.Reader // 请求参数map
}

func New(appID, appKey, openID string, online bool) *Base {
	client := &http.Client{
		Timeout: time.Millisecond * Timeout, // Set 5000ms timeout.
	}
	base := &Base{
		Client: client,
		AppID:  appID,
		AppKey: appKey,
		OpenID: openID,
	}
	if online {
		base.Host = OnlineUrl
	} else {
		base.Host = NoOnlineUrl
	}
	return base
}

// Result 获取结果
func (b *Base) Result() (string, error) {
	resp, err := b.Client.Post(b.Host+b.URL, ContentType, b.RequestData)
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

// WithRequestParams ...
func (b *Base) WithRequestParams(_url string, data map[string]interface{}) *Base {
	data["openid"] = b.OpenID
	data["appid"] = b.AppID
	data["nonce_str"] = b.randString(32)
	data["timestamp"] = time.Now().Unix()
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
	_, keys := b.kSort(data)
	tmp := make([]string, 0)
	for _, val := range keys {
		tmp = append(tmp, fmt.Sprintf("%s=%v", val, data[val]))
	}
	strSign := strings.Join(tmp, "&")
	strSign += "&key=" + b.AppKey
	strSign = strings.ToUpper(strSign)
	md5 := md5.New()
	io.WriteString(md5, strSign)
	md5Str := fmt.Sprintf("%x", md5.Sum(nil))
	return strings.ToUpper(md5Str)
}

// kSort 按照键排序 map
func (b *Base) kSort(arr map[string]interface{}) (map[string]interface{}, []string) {
	var keys []string
	for k := range arr {
		keys = append(keys, k)
	}
	newArr := make(map[string]interface{})
	sort.Strings(keys)
	for _, val := range keys {
		newArr[val] = arr[val]
	}
	return newArr, keys
}

// randString
// @Description: 随机字符串
// @receiver b
// @param length
// @return string
// @author cx
func (b *Base) randString(length int) string {
	const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	s := make([]byte, length)
	for i := range s {
		s[i] = charset[rand.IntN(len(charset))]
	}
	return string(s)
}
