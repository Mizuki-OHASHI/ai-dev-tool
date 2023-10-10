# brainStorming

「既存コード＋ API/DB 定義」から直接コードを生成するのではなく, タスクの列挙などワンクッション挟むことで精度を上げる, もしくはマニュアルで書く指示の量を削減することを目指す.

## 起動手順

1. データベースの起動

```shell
$ cd db
$ docker-compose up // コンテナ起動
$ ./exec.sh // コンテナ内に移動
```

2. 環境変数ファイルの作成

`./app/.env`

```env
MYSQL_USER=app
MYSQL_HOST=localhost:3307
MYSQL_ROOT_PASSWORD=password
MYSQL_DATABASE=brain_storming

PORT=8082
```

`./app/genAI/.env`

```env
OPEN_AI_API_KEY=sk-************
```

3. アプリケーションの起動

```shell
$ cd app
$ go run main.go
```

4. コードの自動生成

```shell
$ cd app/genAI
$ go run main.go
```

## ディレクトリ構成

```shell
$ tree
.
├── README.md
├── app
│   ├── controller
│   │   └── controller.go
│   ├── dao
│   │   └── dao.go
│   ├── definition
│   │   ├── api.yml // API 定義
│   │   └── db.sql // DB 定義
│   ├── genAI
│   │   ├── genai
│   │   │   ├── file.go
│   │   │   ├── genai.go
│   │   │   └── read.go
│   │   ├── go.mod
│   │   ├── go.sum
│   │   ├── main.go
│   │   ├── out1 // タスク一覧等の一時出力先
│   │   └── out2 // 生成されたコードの一時出力先
│   ├── go.mod
│   ├── go.sum
│   ├── main.go
│   ├── model
│   │   └── model.go
│   └── usecase
│       └── usecase.go
└── db
    ├── cache.sql
    ├── docker-compose.yml
    ├── exec.sh
    └── my.cnf
```
