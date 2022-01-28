package bitflyer

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"log"
	"net/http"
	"net/url"
	"strconv"
	"time"
)

const baseURL = "https://api.bitflyer.com/v1"

type APIClient struct {
	key string
	secret string
	httpClient *http.Client // ストラクトのポインタ型？ -> これをどう使おうと思える頭になるのか？
}

// headerを追加する処理
func(api APIClient) header(method, endpoint string, body []byte) map[string]string {
	timestamp := strconv.FormatInt(time.Now().Unix(), 10)
	log.Println(timestamp)
	
	message := timestamp + method + endpoint + string(body)
	fmt.Println(message)

	mac := hmac.New(sha256.New, []byte(api.secret))
	mac.Write([]byte(message))
	fmt.Println(mac)

	sign := hex.EncodeToString(mac.Sum(nil)) // これは何をしてる？

	return map[string]string {
		"ACCESS-KEY": api.key,
		"ACCESS-TIMESTAMP": timestamp,
		"ACCESS-SIGN": sign,
		"Content-Type" : "application/json",
	}
}

// requestを投げる
func (api *APIClient) doRequest(method, urlPath string, query [string]string, data []byte) (body[]byte, err error){
	baseURL := url.Parse(baseURL)
	if err := nil {
		return
	}
	apiURL, err := url.Parse(urlPath)

}