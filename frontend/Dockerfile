# frontend/Dockerfile
FROM node:18

WORKDIR /app

# 依存関係のインストール
COPY package*.json ./
RUN npm install

# アプリケーションコードをコピー
COPY . .

# Next.jsアプリのビルド
RUN npm run build

# サーバーを起動
CMD ["npm", "start"]
