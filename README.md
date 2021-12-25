# Webアプリケーションスケルトン実装
Webアプリケーションのスケルトン実装。必要に迫られたときに、すぐにプロトタイピングできるよう自分向けに作成したもの。

# 要素技術
本プロジェクトは以下の要素技術を含みます。
- [OpenAPI](https://www.openapis.org/)よるAPI定義
- [Prism](https://stoplight.io/open-source/prism/)によるOpenAPI定義ファイルの編集とMockサーバ生成
- [Postman](https://www.postman.com/)による開発用APIリクエスト発行
- [openapi-generator-cli](https://github.com/OpenAPITools/openapi-generator-cli)によるAPIクライアント(typescript-axios版)の生成
- [golang](https://go.dev/)によるAPIサーバ実装
- [Docker/docker-compose](https://www.docker.com/)によるMySQLサーバなどのホスティング
- [JWT](https://jwt.io/)によるユーザ認証実装
- [Nuxt.js](https://nuxtjs.org/)によるフロントエンド実装
    - 現状、中途半端なTypeScript実装
    - Nuxt.js ver.3リリース後に完全TypeScript化を目指す
- [Jest](https://jestjs.io/ja/)によるフロントエンドのテスト
    - 現状中途半端な状態
    - 完全TypeScript化後に再検討

# ファイル/ディレクトリ構成
主なファイルおよびディレクトリの説明。
- api-client
    - openapi-generator-cliにより生成したAPIクライアント
    - フロントエンドでimportして利用
- backend
    - golangによるAPIサーバ実装
    - Dockerのmysqlイメージ生成用docker-compose設定ファイル群を含む
- frontend
    - create-nuxt-appコマンドで生成したNuxt.jsプロジェクト
- openapi.yaml
    - OpenAPI定義
    - Prismで編集およびExportしたものを配置

# 情報源
https://blog.5thfloor.co.jp/2019/06/26/webapp-development-with-openapi-and-typescript/

# 利用方法
## Swagger UI

- 起動

```
docker run --rm -p 8080:8080 -e SWAGGER_JSON=/local/openapi.yaml -v ${PWD}:/local swaggerapi/swagger-ui:v3.20.1
```

## Mockサーバ

- インストール

```
npm install -g @stoplight/prism-cli
```

- 起動

```
prism mock openapi.yaml
```

## API Client

- 生成

```
npm run generate-api-client
```
