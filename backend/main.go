package main

import (
    "log"
    "net/http"
)

func main() {
    // データベース初期化
    initDB()
    defer closeDB()

    // ルートエンドポイント設定
    http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
        w.Write([]byte("Hello, Go Backend!"))
    })

    // その他エンドポイント
    http.HandleFunc("/users", usersHandler)

    // サーバー起動
    log.Println("サーバーを起動します... http://localhost:8080")
    log.Fatal(http.ListenAndServe(":8080", nil))
}
