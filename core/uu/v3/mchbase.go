package v3

import (
	"crypto/md5"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"strings"
	"time"
)

type MchBase struct {
	Client      *http.Client // Change this to your actual client type
	AppID       string
	AppKey      string
	MerchantID  uint64
	StoreNo     string
	Host        string
	URL         string
	Online      bool      // true: 正式 false: 沙箱
	RequestData io.Reader // 请求参数map
}

func NewBase(appID, appKey string, merchantID uint64, storeNo string, online bool) *MchBase {
	client := &http.Client{
		Timeout: time.Millisecond * Timeout, // Set 5000ms timeout.
	}
	base := &MchBase{
		Client:     client,
		AppID:      appID,
		AppKey:     appKey,
		MerchantID: merchantID,
		StoreNo:    storeNo,
		Online:     online,
	}
	if online {
		base.Host = OnlineUrl
	} else {
		base.Host = NoOnlineUrl
	}
	return base
}

// Result 获取结果
func (b *MchBase) Result() (string, error) {
	req, err := http.NewRequest("POST", b.Host+b.URL, b.RequestData)
	req.Header.Set("Content-Type", ContentType)
	req.Header.Set("X-App-Id", b.AppID)
	resp, err := b.Client.Do(req)
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
func (b *MchBase) WithRequestParams(_url string, data map[string]interface{}) *MchBase {
	body, _ := json.Marshal(data)
	params := map[string]interface{}{
		"appkey":    b.AppKey,
		"timestamp": time.Now().Unix(),
		"biz":       string(body),
	}
	params["sign"] = b.Sign(params)
	params["merchantId"] = b.MerchantID
	bytes, _ := json.Marshal(params)
	b.RequestData = strings.NewReader(string(bytes))
	b.URL = _url
	return b
}

// Sign generates the MD5 signature.
func (b *MchBase) Sign(data map[string]interface{}) string {
	signStr := fmt.Sprintf("%s%s%d", data["biz"], data["appkey"], data["timestamp"])
	// 计算 MD5 哈希值
	hash := md5.Sum([]byte(signStr))
	// 将哈希值转换为十六进制字符串
	hexStr := fmt.Sprintf("%x", hash)
	// 将十六进制字符串转换为大写
	return strings.ToUpper(hexStr)
}
