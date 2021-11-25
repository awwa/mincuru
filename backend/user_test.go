package main

import (
	"bytes"
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"os"

	"testing"

	"github.com/getkin/kin-openapi/openapi3"
	"github.com/getkin/kin-openapi/openapi3filter"
	"github.com/getkin/kin-openapi/routers/legacy"
	"github.com/stretchr/testify/assert"
)

func TestMain(m *testing.M) {
	// call flag.Parse() here if TestMain uses flags
	err := initDb("127.0.0.1", 3306)
	if err != nil {
		panic(err)
	}
	// Test
	code := m.Run()

	os.Exit(code)
}

func ServeAndRequest(httpReq *http.Request) (recorder *httptest.ResponseRecorder) {
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
	recorder = httptest.NewRecorder()
	Router().ServeHTTP(recorder, httpReq)

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
	if err := openapi3filter.ValidateRequest(ctx, requestValidationInput); err != nil {
		panic(err)
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

func TestGetUsersExistRecord(t *testing.T) {
	// テスト固有のレコードの準備
	DB.Exec("TRUNCATE TABLE users")
	DB.Create(&User{Name: "hoge taro", Email: "hoge@example.com", Role: "user"})
	DB.Create(&User{Name: "fuga 太郎", Email: "fuga@example.com", Role: "admin"})
	// HTTPリクエストの生成
	httpReq, err := http.NewRequest(http.MethodGet, "http://localhost:8080/users", nil)
	// httpReq.Header.Add("Authorization", "Bearer tokentokentoken")
	httpReq.Header.Add("Content-Type", "application/json")
	httpReq.Header.Add("Authorization", "bearer tokentokentoken")
	if err != nil {
		panic(err)
	}
	// Test用サーバにリクエストを送信して、レスポンスをOpenAPI仕様に照らし合わせる
	recorder := ServeAndRequest(httpReq)
	// テストケース固有のチェック
	assert.Equal(t, 200, recorder.Result().StatusCode)
	// var body []User
	// json.Unmarshal(recorder.Body.Bytes(), &body)
	// assert.Equal(t, 2, len(body))
}

func TestGetUsersByName(t *testing.T) {
	// テスト固有のレコードの準備
	DB.Exec("TRUNCATE TABLE users")
	DB.Create(&User{Name: "hoge taro", Email: "hoge@example.com", Role: "user"})
	DB.Create(&User{Name: "fuga 太郎", Email: "fuga@example.com", Role: "admin"})
	// HTTPリクエストの生成
	httpReq, err := http.NewRequest(http.MethodGet, "http://localhost:8080/users?name=hoge taro", nil)
	httpReq.Header.Add("Content-Type", "application/json")
	httpReq.Header.Add("Authorization", "bearer tokentokentoken")
	if err != nil {
		panic(err)
	}
	// Test用サーバにリクエストを送信して、レスポンスをOpenAPI仕様に照らし合わせる
	recorder := ServeAndRequest(httpReq)
	// テストケース固有のチェック
	assert.Equal(t, 200, recorder.Result().StatusCode)
	var body []User
	json.Unmarshal(recorder.Body.Bytes(), &body)
	assert.Equal(t, 1, len(body))
}

func TestGetUsersByEmail(t *testing.T) {
	// テスト固有のレコードの準備
	DB.Exec("TRUNCATE TABLE users")
	DB.Create(&User{Name: "hoge taro", Email: "hoge@example.com", Role: "user"})
	DB.Create(&User{Name: "fuga 太郎", Email: "fuga@example.com", Role: "admin"})
	// HTTPリクエストの生成
	httpReq, err := http.NewRequest(http.MethodGet, "http://localhost:8080/users?email=hoge@example.com", nil)
	httpReq.Header.Add("Content-Type", "application/json")
	httpReq.Header.Add("Authorization", "bearer tokentokentoken")
	if err != nil {
		panic(err)
	}
	// Test用サーバにリクエストを送信して、レスポンスをOpenAPI仕様に照らし合わせる
	recorder := ServeAndRequest(httpReq)
	// テストケース固有のチェック
	assert.Equal(t, 200, recorder.Result().StatusCode)
	var body []User
	json.Unmarshal(recorder.Body.Bytes(), &body)
	assert.Equal(t, 1, len(body))
}

func TestGetUsersByRole(t *testing.T) {
	// テスト固有のレコードの準備
	DB.Exec("TRUNCATE TABLE users")
	DB.Create(&User{Name: "hoge taro", Email: "hoge@example.com", Role: "user"})
	DB.Create(&User{Name: "fuga 太郎", Email: "fuga@example.com", Role: "admin"})
	// HTTPリクエストの生成
	httpReq, err := http.NewRequest(http.MethodGet, "http://localhost:8080/users?role=user", nil)
	httpReq.Header.Add("Content-Type", "application/json")
	httpReq.Header.Add("Authorization", "bearer tokentokentoken")
	if err != nil {
		panic(err)
	}
	// Test用サーバにリクエストを送信して、レスポンスをOpenAPI仕様に照らし合わせる
	recorder := ServeAndRequest(httpReq)
	// テストケース固有のチェック
	assert.Equal(t, 200, recorder.Result().StatusCode)
	var body []User
	json.Unmarshal(recorder.Body.Bytes(), &body)
	assert.Equal(t, 1, len(body))
}

func TestGetUsersByNameAndEmail(t *testing.T) {
	// テスト固有のレコードの準備
	DB.Exec("TRUNCATE TABLE users")
	DB.Create(&User{Name: "hoge taro", Email: "hoge@example.com", Role: "user"})
	DB.Create(&User{Name: "fuga 太郎", Email: "fuga@example.com", Role: "admin"})
	// HTTPリクエストの生成
	httpReq, err := http.NewRequest(http.MethodGet, "http://localhost:8080/users?name=hoge taro&email=hoge@example.com", nil)
	httpReq.Header.Add("Content-Type", "application/json")
	httpReq.Header.Add("Authorization", "bearer tokentokentoken")
	if err != nil {
		panic(err)
	}
	// Test用サーバにリクエストを送信して、レスポンスをOpenAPI仕様に照らし合わせる
	recorder := ServeAndRequest(httpReq)
	// テストケース固有のチェック
	assert.Equal(t, 200, recorder.Result().StatusCode)
	var body []User
	json.Unmarshal(recorder.Body.Bytes(), &body)
	assert.Equal(t, 1, len(body))
}

func TestGetUsersNoRecord(t *testing.T) {
	// テスト固有のレコードの準備
	DB.Exec("TRUNCATE TABLE users")
	// HTTPリクエストの生成
	httpReq, err := http.NewRequest(http.MethodGet, "http://localhost:8080/users", nil)
	httpReq.Header.Add("Content-Type", "application/json")
	httpReq.Header.Add("Authorization", "bearer tokentokentoken")
	if err != nil {
		panic(err)
	}
	// Test用サーバにリクエストを送信して、レスポンスをOpenAPI仕様に照らし合わせる
	recorder := ServeAndRequest(httpReq)
	// テストケース固有のチェック
	assert.Equal(t, 200, recorder.Result().StatusCode)
	var body []User
	json.Unmarshal(recorder.Body.Bytes(), &body)
	assert.Equal(t, 0, len(body))
}

func TestGetUserExistRecord(t *testing.T) {
	// テスト固有のレコードの準備
	DB.Exec("TRUNCATE TABLE users")
	DB.Create(&User{Name: "hoge taro", Email: "hoge@example.com", Role: "user"})
	// HTTPリクエストの生成
	httpReq, err := http.NewRequest(http.MethodGet, "http://localhost:8080/users/1", nil)
	httpReq.Header.Add("Content-Type", "application/json")
	httpReq.Header.Add("Authorization", "bearer tokentokentoken")
	if err != nil {
		panic(err)
	}
	// Test用サーバにリクエストを送信して、レスポンスをOpenAPI仕様に照らし合わせる
	recorder := ServeAndRequest(httpReq)
	// テストケース固有のチェック
	assert.Equal(t, 200, recorder.Result().StatusCode)
}

func TestGetUserNoRecord(t *testing.T) {
	// テスト固有のレコードの準備
	DB.Exec("TRUNCATE TABLE users")
	DB.Create(&User{Name: "hoge taro", Email: "hoge@example.com", Role: "user"})
	// HTTPリクエストの生成
	// 存在しないIDを指定
	httpReq, err := http.NewRequest(http.MethodGet, "http://localhost:8080/users/123", nil)
	httpReq.Header.Add("Content-Type", "application/json")
	httpReq.Header.Add("Authorization", "bearer tokentokentoken")
	if err != nil {
		panic(err)
	}
	// Test用サーバにリクエストを送信して、レスポンスをOpenAPI仕様に照らし合わせる
	recorder := ServeAndRequest(httpReq)
	// テストケース固有のチェック
	assert.Equal(t, 404, recorder.Result().StatusCode)
}

func TestPostUser(t *testing.T) {
	// テスト固有のレコードの準備
	DB.Exec("TRUNCATE TABLE users")
	// HTTPリクエストの生成
	data := User{
		Name:     "hoge taro",
		Email:    "hoge@example.com",
		Role:     "admin",
		Password: "password",
	}
	body, _ := json.Marshal(data)
	httpReq, err := http.NewRequest(http.MethodPost, "http://localhost:8080/users", bytes.NewBuffer(body))
	httpReq.Header.Add("Content-Type", "application/json")
	httpReq.Header.Add("Authorization", "bearer tokentokentoken")
	// fmt.Println(httpReq)
	if err != nil {
		panic(err)
	}
	// Test用サーバにリクエストを送信して、レスポンスをOpenAPI仕様に照らし合わせる
	recorder := ServeAndRequest(httpReq)
	// テストケース固有のチェック
	assert.Equal(t, 201, recorder.Result().StatusCode)
}
