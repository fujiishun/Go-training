package config

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

var DB *sql.DB

func ConnectDB() {
	var err error
	dsn := fmt.Sprintf("%s:%s@tcp(%s)/%s",
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_HOST"),
		os.Getenv("DB_NAME"),
	)

	// リトライロジック: 5回まで接続を試みる
	for i := 0; i < 5; i++ {
		DB, err = sql.Open("mysql", dsn)
		if err == nil && DB.Ping() == nil {
			log.Println("データベースに接続しました")
			return
		}

		log.Printf("データベース接続を再試行します... (%d/5)", i+1)
		time.Sleep(5 * time.Second)
	}

	// 最後の試行でも接続できなければエラーを出力
	log.Fatalf("データベースへの接続に失敗しました: %v", err)
}

func CloseDB() {
	if DB != nil {
		DB.Close()
	}
}
