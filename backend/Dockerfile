# backend/Dockerfile
FROM golang:1.18

WORKDIR /app

# Goモジュールと依存関係をコピーしてインストール
COPY go.mod go.sum ./
RUN go mod download

# アプリケーションコードをコピー
COPY . .

# サーバーをビルド
RUN go build -o server .

CMD ["/app/server"]
