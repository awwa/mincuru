# 情報源
https://blog.5thfloor.co.jp/2019/06/26/webapp-development-with-openapi-and-typescript/

## Swagger UIの起動

```
docker run --rm -p 8080:8080 -e SWAGGER_JSON=/local/openapi.yaml -v ${PWD}:/local swaggerapi/swagger-ui:v3.20.1
```

## Mockサーバの起動

```
npm install -g @stoplight/prism-cli
prism mock openapi.yaml
```

### Spring編
#### Stubサーバ生成
docker run --rm -v ${PWD}:/local openapitools/openapi-generator-cli generate -i /local/openapi.yaml -g spring -o /local/spring_stub --additional-properties returnSuccessCode=true

#### CORS有効化
spring_stub/src/main/java/org/openapitools/OpenAPI2SpringBoot.java

#### Stubサーバビルド
cd spring_stub
$ docker run --rm -v ${PWD}:/usr/src/mymaven -w /usr/src/mymaven maven mvn package

#### Stubサーバ起動
docker run --rm -p 8080:8080 -v ${PWD}:/usr/src/myapp -w /usr/src/myapp java java -jar target/openapi-spring-0.0.1.jar

### python-flask編
#### Stubサーバ生成
docker run --rm -v ${PWD}:/local openapitools/openapi-generator-cli generate -i /local/openapi.yaml -g python-flask -o /local/python_flask_stub --additional-properties serverPort=3000

#### CORS有効化
__main__py編集
https://github.com/zalando/connexion/issues/357

```
cd python_flask_stub
from flask_cors import CORS
CORS(app.app)
```

#### Stubサーバ起動
cd python_flask_stub
pip3 install -r requirements.txt
pip install -U flask-cors
python3 -m openapi_server

## API Clientの生成
cd .
docker run --rm -v ${PWD}:/local openapitools/openapi-generator-cli generate -g typescript-fetch -i /local/openapi.yaml -o /local/mincuru-api-client --additional-properties=modelPropertyNaming=camelCase,supportsES6=true,withInterfaces=true,typescriptThreePlus=true

