package main

import (
	"go-btc-trading/config"
	"go-btc-trading/utils"
)

func main() {
	// fmt.Println(config.Config.ApiKey) // パッケージ.構造体.変数で呼び出しをする
	// fmt.Println(config.Config.ApiSecret)
	utils.LoggingSetting(config.Config.LogFile) // ログ出力の設定
	// log.Println("test")
}