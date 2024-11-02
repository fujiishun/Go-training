# Go-Training プロジェクト

Go をバックエンド、Next.js をフロントエンド、MySQL をデータベースとして使用した Web アプリケーションです。開発および本番環境の構築・管理には Docker と Docker Compose を使用しています。

## 技術スタック

- **フロントエンド**: [Next.js](https://nextjs.org/)
- **バックエンド**: [Go](https://golang.org/)
- **データベース**: [MySQL](https://www.mysql.com/)
- **コンテナ管理**: [Docker](https://www.docker.com/) & [Docker Compose](https://docs.docker.com/compose/)

## 環境起動

```
docker-compose up --build
```

起動後、以下の URL で各サービスにアクセスできます。

フロントエンド（Next.js）: http://localhost:3000
バックエンド（Go）: http://localhost:8080

## マイグレーション

```
docker-compose exec db mysql -u user -ppassword go_db -e "source /docker-entrypoint-initdb.d/migrations/init.sql"
```

```
docker-compose exec db mysql -u user -ppassword go_db -e "source /docker-entrypoint-initdb.d/seeds/seed_data.sql"
```
