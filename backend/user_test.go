package main

import (
	"context"
	"fmt"
	"net/http"
	"net/http/httptest"

	"testing"

	"github.com/getkin/kin-openapi/openapi3"
	"github.com/getkin/kin-openapi/openapi3filter"
	"github.com/getkin/kin-openapi/routers/legacy"
	"gorm.io/gorm"
)

func TestGetUser2(t *testing.T) {
	router := Router()

	w := httptest.NewRecorder()
	httpReq, err4 := http.NewRequest(http.MethodGet, "/users/123", nil)
	httpReq.Header.Set("Authorization", "Bearer tokentokentoken")
	if err4 != nil {
		panic(err4)
	}
	router.ServeHTTP(w, httpReq)

	fmt.Println(w.Body.String())
}

func Setup(t *testing.T) (db *gorm.DB) {
	db, err := initDb("127.0.0.1", 3306)
	if err != nil {
		panic(err)
	}
	db.Exec("TRUNCATE TABLE users")
	return
}

func Assert(t *testing.T, httpReq *http.Request) (db *gorm.DB) {
	ctx := context.Background()
	loader := &openapi3.Loader{Context: ctx}
	doc, err := loader.LoadFromFile("../openapi.yaml")
	if err != nil {
		panic(err)
	}
	err = doc.Validate(ctx)
	if err != nil {
		panic(err)
	}
	router, err := legacy.NewRouter(doc)
	if err != nil {
		panic(err)
	}
	gin := Router()
	recorder := httptest.NewRecorder()

	gin.ServeHTTP(recorder, httpReq)

	// Find route
	route, pathParams, err := router.FindRoute(httpReq)
	if err != nil {
		panic(err)
	}

	// Validate request
	requestValidationInput := &openapi3filter.RequestValidationInput{
		Request:    httpReq,
		PathParams: pathParams,
		Route:      route,
	}

	responseValidationInput := &openapi3filter.ResponseValidationInput{
		RequestValidationInput: requestValidationInput,
		Status:                 recorder.Result().StatusCode,
		Header:                 recorder.Result().Header,
	}
	responseValidationInput.SetBodyBytes(recorder.Body.Bytes())

	// Validate response
	if err := openapi3filter.ValidateResponse(ctx, responseValidationInput); err != nil {
		panic(err)
	}

	return
}

func TestGetUser3(t *testing.T) {
	var err error
	// テストの初期化（主にDBのクリア）
	db := Setup(t)
	// テスト固有のレコードの準備
	db.Create(&User{Name: "hoge taro", Email: "hoge@example.com", Role: "user"})
	// HTTPリクエストの生成
	httpReq, err := http.NewRequest(http.MethodGet, "http://localhost:8080/users/123", nil)
	httpReq.Header.Add("Authorization", "Bearer tokentokentoken")
	if err != nil {
		panic(err)
	}
	// Test用サーバにリクエストを送信して、レスポンスをOpenAPI仕様に照らし合わせる
	Assert(t, httpReq)
}
