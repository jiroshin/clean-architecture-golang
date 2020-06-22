# クリーンアーキテクチャのサンプル

Goで書かれた典型的なクリーンアーキテクチャっぽいレイヤー分けを採用したAPIサーバーです。  

## サーバーの立ち上げ方
```sh
$ make start_prepopulated_db
$ make start_server
```

１つ目のコマンドでPostgresのdockerコンテナを立ち上げています。  
もし、MySQLを使いたい場合は環境変数`DB_TYPE`に`mysql`を設定してください。  
MySQLを利用する場合はローカル環境にMySQLデータベースの準備が必要です。

```sh
$ export DB_TYPE=mysql && make start_server
```

## エンドポイントとリクエストサンプル
`POST` `/books`
`Content-type: application/json`
Request Body
```json
{
  "title":"タイトル",
  "author":"jiroshin",
  "overview": "良い本でした"
}
```

curlリクエストサンプル

```sh
curl -X POST localhost:8080/books -H "Content-Type: application/json" -d '{"title":"タイトル", "author":"jiroshin", "overview": "面白い本でした"}'
```
