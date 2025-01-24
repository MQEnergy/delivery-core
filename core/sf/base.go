package sf

import (
	"crypto/md5"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"strings"
	"time"
)

const (
	BaseUrl     = "https://openic.sf-express.com" // 线上地址
	Timeout     = 5000                            // 超时时间 5000ms
	ContentType = "application/json"
)

// Base represents the main structure for Dada delivery service.
type Base struct {
	Client      *http.Client // Assuming client type is defined elsewhere
	DevID       string
	DevKey      string
	Host        string
	URL         string
	RequestData io.Reader // 请求参数map
}

// New is the constructor for the Base struct
func New(devID, devKey string) *Base {
	base := &Base{
		DevID:  devID,
		DevKey: devKey,
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
	data["dev_id"] = b.DevID
	data["push_time"] = time.Now().Unix()
	bytes, _ := json.Marshal(data)
	b.RequestData = strings.NewReader(string(bytes))
	b.URL = _url + "?sign=" + b.Sign(data)
	return b
}

// Sign generates the MD5 signature.
func (b *Base) Sign(data map[string]interface{}) string {
	postData, err := json.Marshal(data)
	if err != nil {
		return ""
	}
	signChar := fmt.Sprintf("%s&%s&%s", string(postData), b.DevID, b.DevKey)
	md5Hash := md5.Sum([]byte(signChar))
	md5Str := fmt.Sprintf("%x", md5Hash)
	sign := base64.StdEncoding.EncodeToString([]byte(md5Str))
	return sign
}
