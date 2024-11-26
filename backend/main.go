package main

import (
	"backend/config"
	"backend/routes"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"log"
)

func init() {
	// 環境変数をロード
	err := godotenv.Load()
	if err != nil {
		log.Println("環境変数ファイルが見つかりません。デフォルト値を使用します。")
	}
}

func main() {
	// Ginエンジンを初期化
	r := gin.Default()

	// データベース接続
	config.ConnectDB()
	defer config.CloseDB()

	// ルート設定
	routes.SetupRoutes(r)

	// サーバー起動
	r.Run(":8080")
}
