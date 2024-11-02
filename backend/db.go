package main

import (
    "database/sql"
    "log"
    "time"
    _ "github.com/go-sql-driver/mysql" // MySQLドライバ
)

// DB接続用の変数
var db *sql.DB

// データベースに接続し、接続インスタンスを返す関数
func initDB() {
    var err error
    dsn := "user:password@tcp(db:3306)/go_db"

    // リトライロジック: 5回まで接続を試みる
    for i := 0; i < 5; i++ {
        db, err = sql.Open("mysql", dsn)
        if err == nil && db.Ping() == nil {
            log.Println("データベースに接続しました")
            return
        }

        log.Printf("データベース接続を再試行します... (%d/5)", i+1)
        time.Sleep(5 * time.Second)
    }

    // 最後の試行でも接続できなければエラーを出力
    log.Fatalf("データベースへの接続に失敗しました: %v", err)
}

// アプリケーション終了時にDB接続を閉じる関数
func closeDB() {
    if db != nil {
        db.Close()
    }
}
