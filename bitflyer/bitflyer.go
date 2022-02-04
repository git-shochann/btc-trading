package bitflyer

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"net/http"
	"net/url"
	"strconv"
	"time"
)

const baseURL = "https://api.bitflyer.com/v1"

// 変数の集合体 -> 型
type APIClient struct {
	Key string
	Secret string
	HttpClient *http.Client // 構造体型の中に構造体型
}

// headerを追加する処理
func(api *APIClient) Header(method, endpoint string, body []byte) map[string]string {
	timestamp := strconv.FormatInt(time.Now().Unix(), 10) // stringに変換
	message := timestamp + method + endpoint + string(body)

	mac := hmac.New(sha256.New, []byte(api.Secret))

	mac.Write([]byte(message))

	sign := hex.EncodeToString(mac.Sum(nil))
	return map[string]string{
		"ACCESS_KEY" :	api.Key,
		"ACCESS-TIMESTAMP" : timestamp,
		"ACCESS-SIGN" : sign,
		"CONTENT-TYPE" : "application/json",
	}
	
}

func(api *APIClient) doRequest(method, urlPath string, query map[string]string, data[]byte) (body []byte, err error) {
	baseURL, err := url.Parse(baseURL)
	if err != nil {
		return
	}
	apiURL, err := url.Parse(urlPath)
	if err != nil {
		return
	}
	endpoint := baseURL.ResolveReference(apiURL)
}