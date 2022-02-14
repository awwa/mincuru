## go-taskのインストール
https://taskfile.dev/#/usage?id=getting-started

## APIサーバ

- 環境変数設定

```
cp .env.example .env
vi .env
```

- 依存関係の解決

```
go mod tidy
```

- テスト実行

```
go test
```

- 開発サーバの起動

```
go run .
```

- ビルド＆本番サーバ起動

```
task build
./dist/mincuru-api-server
```

## DB

- 再構築

```
task dbrecreate
```

- 起動

```
task dbup
```

- 削除

```
task dbdown
```
