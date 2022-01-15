package main

import (
	"fmt"
	"go-btc-trading/config"
)

func main() {
	fmt.Println(config.Config.ApiKey) // パッケージ.構造体.変数で呼び出しをする
	fmt.Println(config.Config.ApiSecret)
}