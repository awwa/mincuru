package main

import (
	"context"
	"net/http"
	"net/http/httptest"

	"testing"

	"github.com/getkin/kin-openapi/openapi3"
	"github.com/getkin/kin-openapi/openapi3filter"
	"github.com/getkin/kin-openapi/routers/legacy"
	"gorm.io/gorm"
)

func Setup(t *testing.T) {
	err := initDb("127.0.0.1", 3306)
	if err != nil {
		panic(err)
	}
	DB.Exec("TRUNCATE TABLE users")
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

// func TestGetUser2(t *testing.T) {
// 	router := Router()

// 	w := httptest.NewRecorder()
// 	httpReq, err4 := http.NewRequest(http.MethodGet, "/users/123", nil)
// 	httpReq.Header.Set("Authorization", "Bearer tokentokentoken")
// 	if err4 != nil {
// 		panic(err4)
// 	}
// 	router.ServeHTTP(w, httpReq)

// 	fmt.Println(w.Body.String())
// }

func TestGetUser(t *testing.T) {
	// テストの初期化（主にDBのクリア）
	Setup(t)
	// テスト固有のレコードの準備
	DB.Create(&User{Name: "hoge taro", Email: "hoge@example.com", Role: "user"})
	// HTTPリクエストの生成
	httpReq, err := http.NewRequest(http.MethodGet, "http://localhost:8080/users/1", nil)
	httpReq.Header.Add("Authorization", "Bearer tokentokentoken")
	if err != nil {
		panic(err)
	}
	// Test用サーバにリクエストを送信して、レスポンスをOpenAPI仕様に照らし合わせる
	Assert(t, httpReq)

	// HTTPリクエストの生成
	// 存在しないIDを指定
	httpReq2, err := http.NewRequest(http.MethodGet, "http://localhost:8080/users/123", nil)
	httpReq2.Header.Add("Authorization", "Bearer tokentokentoken")
	if err != nil {
		panic(err)
	}
	// Test用サーバにリクエストを送信して、レスポンスをOpenAPI仕様に照らし合わせる
	Assert(t, httpReq2)

}
