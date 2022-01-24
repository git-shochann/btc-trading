package bitflyer

import (
	"net/http"
)

const baseURL = "https://api.bitflyer.com/v1"

type APIClient struct {
	key string
	secret string
	httpClient *http.Client // ストラクトのポインタ型？ -> これをどう使おうと思える頭になるのか？
}

// api認証のためのheaderをつくる
func(api APIClient) header(method, endpoint string, body []byte) map[string]string {
	timestamp := strconv.
}