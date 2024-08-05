package dada

import (
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
	"github.com/spf13/cast"
	"io"
	"io/ioutil"
	"net/http"
	"sort"
	"strings"
	"time"
)

const (
	OnlineUrl   = "https://newopen.imdada.cn"
	NoOnlineUrl = "http://newopen.qa.imdada.cn"
	Timeout     = 5000  // 超时时间 5000ms
	Version     = "1.0" // api版本
	Format      = "json"
)

// Base represents the main structure for Dada delivery service.
type Base struct {
	Client      *http.Client // Assuming client type is defined elsewhere
	AppKey      string
	AppSecret   string
	SourceID    string
	Host        string
	URL         string
	RequestData io.Reader // 请求参数map
}

// New is the constructor for Base.
func New(sourceID, appKey, appSecret string, online bool) *Base {
	base := &Base{
		AppKey:    appKey,
		AppSecret: appSecret,
		SourceID:  sourceID,
	}
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
func (b *Base) WithRequestParams(url string, data map[string]interface{}) *Base {
	body, _ := json.Marshal(data)
	params := map[string]interface{}{
		"app_key":   b.AppKey,
		"body":      string(body),
		"format":    Format,
		"v":         Version,
		"source_id": b.SourceID,
		"timestamp": time.Now().Unix(),
	}
	// Set the signature
	params["signature"] = b.Sign(params)
	bytes, _ := json.Marshal(params)
	b.RequestData = strings.NewReader(string(bytes))
	b.URL = url
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

	finalStr := b.AppSecret + args.String() + b.AppSecret

	// 3. MD5 sign and convert to uppercase
	hash := md5.Sum([]byte(finalStr))
	signature := strings.ToUpper(hex.EncodeToString(hash[:]))

	return signature
}
