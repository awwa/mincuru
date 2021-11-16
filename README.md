# 情報源
https://blog.5thfloor.co.jp/2019/06/26/webapp-development-with-openapi-and-typescript/

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
npm run buildcl
```
