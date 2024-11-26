# ベースイメージ
FROM golang:1.18

# 作業ディレクトリを設定
WORKDIR /app

# 必要なファイルをコピー
COPY go.mod go.sum ./

# 依存関係をダウンロード
RUN go mod download

# アプリケーションコードをコピー
COPY . .

# アプリケーションのエントリーポイントを設定
CMD ["sh", "-c", "go mod tidy && go run main.go"]
