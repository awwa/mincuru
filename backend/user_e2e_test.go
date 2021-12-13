package main

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"os"
	"strconv"
	"strings"

	"testing"

	"github.com/getkin/kin-openapi/openapi3"
	"github.com/getkin/kin-openapi/openapi3filter"
	"github.com/getkin/kin-openapi/routers/legacy"
	"github.com/stretchr/testify/assert"
	"golang.org/x/crypto/bcrypt"
)

func TestMain(m *testing.M) {
	// 環境変数
	Loadenv()
	// DB初期化
	initDb()
	// テスト実行
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
	// if err := openapi3filter.ValidateRequest(ctx, requestValidationInput); err != nil {
	// 	panic(err)
	// }

	responseValidationInput := &openapi3filter.ResponseValidationInput{
		RequestValidationInput: requestValidationInput,
		Status:                 recorder.Result().StatusCode,
		Header:                 recorder.Result().Header,
	}
	responseValidationInput.SetBodyBytes(recorder.Body.Bytes())
	// fmt.Println(recorder.Result())

	// Validate response
	if err := openapi3filter.ValidateResponse(ctx, responseValidationInput); err != nil {
		panic(err)
	}
	return
}

func createTestData() {
	// テスト固有のレコードの準備
	DB.Exec("TRUNCATE TABLE users")
	bcCost, err := strconv.Atoi(os.Getenv("BC_COST"))
	if err != nil {
		panic(err)
	}
	hashed, err := bcrypt.GenerateFromPassword([]byte("password"), bcCost)
	if err != nil {
		panic(err)
	}
	DB.Create(&User{UserResp: UserResp{Name: "hoge taro", Email: "hoge@example.com", Role: "user"}, Password: string(hashed)})
	DB.Create(&User{UserResp: UserResp{Name: "fuga 太郎", Email: "fuga@example.com", Role: "admin"}, Password: string(hashed)})
}

func login() string {
	// HTTPリクエストの生成
	body := `{
		"email": "hoge@example.com",
    	"password": "password"
	}`
	httpReq, err := http.NewRequest(http.MethodPost, "http://localhost:8080/users/login", strings.NewReader(body))
	httpReq.Header.Add("Content-Type", "application/json")
	if err != nil {
		panic(err)
	}
	// Test用サーバにリクエストを送信して、レスポンスをOpenAPI仕様に照らし合わせる
	recorder := ServeAndRequest(httpReq)
	var tokenResp TokenResp
	json.Unmarshal(recorder.Body.Bytes(), &tokenResp)
	return tokenResp.Token
}

func TestGetUsersExistRecord(t *testing.T) {
	createTestData() // テストデータの準備
	token := login() // 認証実行
	// HTTPリクエストの生成
	httpReq, err := http.NewRequest(http.MethodGet, "http://localhost:8080/users", nil)
	httpReq.Header.Add("Content-Type", "application/json")
	httpReq.Header.Add("Authorization", fmt.Sprintf("Bearer %s", token))
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

func TestGetUsersNoLogin(t *testing.T) {
	createTestData() // テストデータの準備
	// HTTPリクエストの生成
	httpReq, err := http.NewRequest(http.MethodGet, "http://localhost:8080/users", nil)
	httpReq.Header.Add("Content-Type", "application/json")
	httpReq.Header.Add("Authorization", fmt.Sprintf("Bearer %s", "token"))
	if err != nil {
		panic(err)
	}
	// Test用サーバにリクエストを送信して、レスポンスをOpenAPI仕様に照らし合わせる
	recorder := ServeAndRequest(httpReq)
	// テストケース固有のチェック
	assert.Equal(t, 401, recorder.Result().StatusCode)
}

func TestGetUsersNoAuthorization(t *testing.T) {
	createTestData() // テストデータの準備
	// HTTPリクエストの生成
	httpReq, err := http.NewRequest(http.MethodGet, "http://localhost:8080/users", nil)
	httpReq.Header.Add("Content-Type", "application/json")
	if err != nil {
		panic(err)
	}
	// Test用サーバにリクエストを送信して、レスポンスをOpenAPI仕様に照らし合わせる
	recorder := ServeAndRequest(httpReq)
	// テストケース固有のチェック
	assert.Equal(t, 401, recorder.Result().StatusCode)
}

func TestGetUsersByName(t *testing.T) {
	createTestData() // テストデータの準備
	token := login() // 認証実行
	// HTTPリクエストの生成
	httpReq, err := http.NewRequest(http.MethodGet, "http://localhost:8080/users?name=hoge taro", nil)
	httpReq.Header.Add("Content-Type", "application/json")
	httpReq.Header.Add("Authorization", fmt.Sprintf("Bearer %s", token))
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
	createTestData() // テストデータの準備
	token := login() // 認証実行
	// HTTPリクエストの生成
	httpReq, err := http.NewRequest(http.MethodGet, "http://localhost:8080/users?email=hoge@example.com", nil)
	httpReq.Header.Add("Content-Type", "application/json")
	httpReq.Header.Add("Authorization", fmt.Sprintf("Bearer %s", token))
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
	createTestData() // テストデータの準備
	token := login() // 認証実行
	// HTTPリクエストの生成
	httpReq, err := http.NewRequest(http.MethodGet, "http://localhost:8080/users?role=user", nil)
	httpReq.Header.Add("Content-Type", "application/json")
	httpReq.Header.Add("Authorization", fmt.Sprintf("Bearer %s", token))
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
	createTestData() // テストデータの準備
	token := login() // 認証実行
	// HTTPリクエストの生成
	httpReq, err := http.NewRequest(http.MethodGet, "http://localhost:8080/users?name=hoge taro&email=hoge@example.com", nil)
	httpReq.Header.Add("Content-Type", "application/json")
	httpReq.Header.Add("Authorization", fmt.Sprintf("Bearer %s", token))
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
	createTestData() // テストデータの準備
	token := login() // 認証実行
	// HTTPリクエストの生成
	httpReq, err := http.NewRequest(http.MethodGet, "http://localhost:8080/users?name=norecord", nil)
	httpReq.Header.Add("Content-Type", "application/json")
	httpReq.Header.Add("Authorization", fmt.Sprintf("Bearer %s", token))
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
	createTestData() // テストデータの準備
	token := login() // 認証実行
	// HTTPリクエストの生成
	httpReq, err := http.NewRequest(http.MethodGet, "http://localhost:8080/users/1", nil)
	httpReq.Header.Add("Content-Type", "application/json")
	httpReq.Header.Add("Authorization", fmt.Sprintf("Bearer %s", token))
	if err != nil {
		panic(err)
	}
	// Test用サーバにリクエストを送信して、レスポンスをOpenAPI仕様に照らし合わせる
	recorder := ServeAndRequest(httpReq)
	// テストケース固有のチェック
	assert.Equal(t, 200, recorder.Result().StatusCode)
}

func TestGetUserNoRecord(t *testing.T) {
	createTestData() // テストデータの準備
	token := login() // 認証実行
	// HTTPリクエストの生成
	httpReq, err := http.NewRequest(http.MethodGet, "http://localhost:8080/users/123", nil)
	httpReq.Header.Add("Content-Type", "application/json")
	httpReq.Header.Add("Authorization", fmt.Sprintf("Bearer %s", token))
	if err != nil {
		panic(err)
	}
	// Test用サーバにリクエストを送信して、レスポンスをOpenAPI仕様に照らし合わせる
	recorder := ServeAndRequest(httpReq)
	// テストケース固有のチェック
	assert.Equal(t, 404, recorder.Result().StatusCode)
}

func TestGetUserNoLogin(t *testing.T) {
	createTestData() // テストデータの準備
	// HTTPリクエストの生成
	httpReq, err := http.NewRequest(http.MethodGet, "http://localhost:8080/users/1", nil)
	httpReq.Header.Add("Content-Type", "application/json")
	httpReq.Header.Add("Authorization", fmt.Sprintf("Bearer %s", "token"))
	if err != nil {
		panic(err)
	}
	// Test用サーバにリクエストを送信して、レスポンスをOpenAPI仕様に照らし合わせる
	recorder := ServeAndRequest(httpReq)
	// テストケース固有のチェック
	assert.Equal(t, 401, recorder.Result().StatusCode)
}

func TestGetUserMeSuccess(t *testing.T) {
	createTestData() // テストデータの準備
	token := login() // 認証実行
	// HTTPリクエストの生成
	httpReq, err := http.NewRequest(http.MethodGet, "http://localhost:8080/users/me", nil)
	httpReq.Header.Add("Content-Type", "application/json")
	httpReq.Header.Add("Authorization", fmt.Sprintf("Bearer %s", token))
	if err != nil {
		panic(err)
	}
	// Test用サーバにリクエストを送信して、レスポンスをOpenAPI仕様に照らし合わせる
	recorder := ServeAndRequest(httpReq)
	// テストケース固有のチェック
	assert.Equal(t, 200, recorder.Result().StatusCode)
}

func TestGetUserMeNoLogin(t *testing.T) {
	createTestData() // テストデータの準備
	// HTTPリクエストの生成
	httpReq, err := http.NewRequest(http.MethodGet, "http://localhost:8080/users/me", nil)
	httpReq.Header.Add("Content-Type", "application/json")
	httpReq.Header.Add("Authorization", fmt.Sprintf("Bearer %s", "token"))
	if err != nil {
		panic(err)
	}
	// Test用サーバにリクエストを送信して、レスポンスをOpenAPI仕様に照らし合わせる
	recorder := ServeAndRequest(httpReq)
	// テストケース固有のチェック
	assert.Equal(t, 401, recorder.Result().StatusCode)
}

func TestPostUserSuccess(t *testing.T) {
	createTestData() // テストデータの準備
	token := login() // 認証実行
	// HTTPリクエストの生成
	body := `{
		"name": "new reocrd",
		"email": "new@example.com",
		"role": "user",
		"password": "password"
	}`
	httpReq, err := http.NewRequest(http.MethodPost, "http://localhost:8080/users", strings.NewReader(body))
	httpReq.Header.Add("Content-Type", "application/json")
	httpReq.Header.Add("Authorization", fmt.Sprintf("Bearer %s", token))
	if err != nil {
		panic(err)
	}
	// Test用サーバにリクエストを送信して、レスポンスをOpenAPI仕様に照らし合わせる
	recorder := ServeAndRequest(httpReq)
	// テストケース固有のチェック
	assert.Equal(t, 201, recorder.Result().StatusCode)
}

func TestPostUserInvalidRequest(t *testing.T) {
	createTestData() // テストデータの準備
	token := login() // 認証実行
	// HTTPリクエストの生成
	body := `{
		"name": "new record",
		"email": "new@example.com",
		"role": "user"
	}`
	httpReq, err := http.NewRequest(http.MethodPost, "http://localhost:8080/users", strings.NewReader(body))
	httpReq.Header.Add("Content-Type", "application/json")
	httpReq.Header.Add("Authorization", fmt.Sprintf("Bearer %s", token))
	if err != nil {
		panic(err)
	}
	// Test用サーバにリクエストを送信して、レスポンスをOpenAPI仕様に照らし合わせる
	recorder := ServeAndRequest(httpReq)
	// テストケース固有のチェック
	assert.Equal(t, 400, recorder.Result().StatusCode)
}

func TestPostUserDupKey(t *testing.T) {
	createTestData() // テストデータの準備
	token := login() // 認証実行
	// HTTPリクエストの生成
	body := `{
		"name": "hoge taro",
		"email": "hoge@example.com",
		"role": "user",
		"password": "password"
	}`
	httpReq, err := http.NewRequest(http.MethodPost, "http://localhost:8080/users", strings.NewReader(body))
	httpReq.Header.Add("Content-Type", "application/json")
	httpReq.Header.Add("Authorization", fmt.Sprintf("Bearer %s", token))
	if err != nil {
		panic(err)
	}
	// Test用サーバにリクエストを送信して、レスポンスをOpenAPI仕様に照らし合わせる
	recorder := ServeAndRequest(httpReq)
	// テストケース固有のチェック
	assert.Equal(t, 400, recorder.Result().StatusCode)
}

func TestPostUserNoLogin(t *testing.T) {
	createTestData() // テストデータの準備
	// HTTPリクエストの生成
	body := `{
		"name": "new reocrd",
		"email": "new@example.com",
		"role": "user",
		"password": "password"
	}`
	httpReq, err := http.NewRequest(http.MethodPost, "http://localhost:8080/users", strings.NewReader(body))
	httpReq.Header.Add("Content-Type", "application/json")
	httpReq.Header.Add("Authorization", fmt.Sprintf("Bearer %s", "token"))
	if err != nil {
		panic(err)
	}
	// Test用サーバにリクエストを送信して、レスポンスをOpenAPI仕様に照らし合わせる
	recorder := ServeAndRequest(httpReq)
	// テストケース固有のチェック
	assert.Equal(t, 401, recorder.Result().StatusCode)
}

func TestPatchUserSuccessAllColumn(t *testing.T) {
	createTestData() // テストデータの準備
	token := login() // 認証実行
	// HTTPリクエストの生成
	body := `{
		"name": "hoge taro2",
		"email": "hoge@example.com2",
		"role": "user",
		"password": "password2"
	}`
	httpReq, err := http.NewRequest(http.MethodPatch, "http://localhost:8080/users/1", strings.NewReader(body))
	httpReq.Header.Add("Content-Type", "application/json")
	httpReq.Header.Add("Authorization", fmt.Sprintf("Bearer %s", token))
	if err != nil {
		panic(err)
	}
	// Test用サーバにリクエストを送信して、レスポンスをOpenAPI仕様に照らし合わせる
	recorder := ServeAndRequest(httpReq)
	// テストケース固有のチェック
	assert.Equal(t, 200, recorder.Result().StatusCode)
}

func TestPatchUserSuccessNameColumn(t *testing.T) {
	createTestData() // テストデータの準備
	token := login() // 認証実行
	// HTTPリクエストの生成
	body := `{
		"name": "hoge taro2"
	}`
	httpReq, err := http.NewRequest(http.MethodPatch, "http://localhost:8080/users/1", strings.NewReader(body))
	httpReq.Header.Add("Content-Type", "application/json")
	httpReq.Header.Add("Authorization", fmt.Sprintf("Bearer %s", token))
	if err != nil {
		panic(err)
	}
	// Test用サーバにリクエストを送信して、レスポンスをOpenAPI仕様に照らし合わせる
	recorder := ServeAndRequest(httpReq)
	// テストケース固有のチェック
	assert.Equal(t, 200, recorder.Result().StatusCode)
}

func TestPatchUserSuccessEmailColumn(t *testing.T) {
	createTestData() // テストデータの準備
	token := login() // 認証実行
	// HTTPリクエストの生成
	body := `{
		"email": "hoge@example.com2"
	}`
	httpReq, err := http.NewRequest(http.MethodPatch, "http://localhost:8080/users/1", strings.NewReader(body))
	httpReq.Header.Add("Content-Type", "application/json")
	httpReq.Header.Add("Authorization", fmt.Sprintf("Bearer %s", token))
	if err != nil {
		panic(err)
	}
	// Test用サーバにリクエストを送信して、レスポンスをOpenAPI仕様に照らし合わせる
	recorder := ServeAndRequest(httpReq)
	// テストケース固有のチェック
	assert.Equal(t, 200, recorder.Result().StatusCode)
}

func TestPatchUserSuccessRoleColumn(t *testing.T) {
	createTestData() // テストデータの準備
	token := login() // 認証実行
	// HTTPリクエストの生成
	body := `{
		"role": "user"
	}`
	httpReq, err := http.NewRequest(http.MethodPatch, "http://localhost:8080/users/1", strings.NewReader(body))
	httpReq.Header.Add("Content-Type", "application/json")
	httpReq.Header.Add("Authorization", fmt.Sprintf("Bearer %s", token))
	if err != nil {
		panic(err)
	}
	// Test用サーバにリクエストを送信して、レスポンスをOpenAPI仕様に照らし合わせる
	recorder := ServeAndRequest(httpReq)
	// テストケース固有のチェック
	assert.Equal(t, 200, recorder.Result().StatusCode)
}

func TestPatchUserSuccessPasswordColumn(t *testing.T) {
	createTestData() // テストデータの準備
	token := login() // 認証実行
	// HTTPリクエストの生成
	body := `{
		"password": "password2"
	}`
	httpReq, err := http.NewRequest(http.MethodPatch, "http://localhost:8080/users/1", strings.NewReader(body))
	httpReq.Header.Add("Content-Type", "application/json")
	httpReq.Header.Add("Authorization", fmt.Sprintf("Bearer %s", token))
	if err != nil {
		panic(err)
	}
	// Test用サーバにリクエストを送信して、レスポンスをOpenAPI仕様に照らし合わせる
	recorder := ServeAndRequest(httpReq)
	// テストケース固有のチェック
	assert.Equal(t, 200, recorder.Result().StatusCode)
}

func TestPatchUserSuccessNoColumn(t *testing.T) {
	createTestData() // テストデータの準備
	token := login() // 認証実行
	// HTTPリクエストの生成
	body := `{}`
	httpReq, err := http.NewRequest(http.MethodPatch, "http://localhost:8080/users/1", strings.NewReader(body))
	httpReq.Header.Add("Content-Type", "application/json")
	httpReq.Header.Add("Authorization", fmt.Sprintf("Bearer %s", token))
	if err != nil {
		panic(err)
	}
	// Test用サーバにリクエストを送信して、レスポンスをOpenAPI仕様に照らし合わせる
	recorder := ServeAndRequest(httpReq)
	// テストケース固有のチェック
	assert.Equal(t, 200, recorder.Result().StatusCode)
}

func TestPatchUserNoRecord(t *testing.T) {
	createTestData() // テストデータの準備
	token := login() // 認証実行
	// HTTPリクエストの生成
	body := `{
		"name": "hoge taro2",
		"email": "hoge@example.com2",
		"role": "user",
		"password": "password2"
	}`
	httpReq, err := http.NewRequest(http.MethodPatch, "http://localhost:8080/users/123", strings.NewReader(body))
	httpReq.Header.Add("Content-Type", "application/json")
	httpReq.Header.Add("Authorization", fmt.Sprintf("Bearer %s", token))
	if err != nil {
		panic(err)
	}
	// Test用サーバにリクエストを送信して、レスポンスをOpenAPI仕様に照らし合わせる
	recorder := ServeAndRequest(httpReq)
	// テストケース固有のチェック
	assert.Equal(t, 404, recorder.Result().StatusCode)
}

func TestPatchUserNoLogin(t *testing.T) {
	createTestData() // テストデータの準備
	// HTTPリクエストの生成
	body := `{
		"name": "hoge taro2",
		"email": "hoge@example.com2",
		"role": "user",
		"password": "password2"
	}`
	httpReq, err := http.NewRequest(http.MethodPatch, "http://localhost:8080/users/1", strings.NewReader(body))
	httpReq.Header.Add("Content-Type", "application/json")
	httpReq.Header.Add("Authorization", fmt.Sprintf("Bearer %s", "token"))
	if err != nil {
		panic(err)
	}
	// Test用サーバにリクエストを送信して、レスポンスをOpenAPI仕様に照らし合わせる
	recorder := ServeAndRequest(httpReq)
	// テストケース固有のチェック
	assert.Equal(t, 401, recorder.Result().StatusCode)
}

func TestPatchUserMeSuccessAllColumn(t *testing.T) {
	createTestData() // テストデータの準備
	token := login() // 認証実行
	// HTTPリクエストの生成
	body := `{
		"name": "updated name",
		"email": "piyo@example.com",
		"password": "password2"
	}`
	httpReq, err := http.NewRequest(http.MethodPatch, "http://localhost:8080/users/me", strings.NewReader(body))
	httpReq.Header.Add("Content-Type", "application/json")
	httpReq.Header.Add("Authorization", fmt.Sprintf("Bearer %s", token))
	if err != nil {
		panic(err)
	}
	// Test用サーバにリクエストを送信して、レスポンスをOpenAPI仕様に照らし合わせる
	recorder := ServeAndRequest(httpReq)
	// テストケース固有のチェック
	assert.Equal(t, 200, recorder.Result().StatusCode)
}

func TestPatchUserMeSuccessNameColumn(t *testing.T) {
	createTestData() // テストデータの準備
	token := login() // 認証実行
	// HTTPリクエストの生成
	body := `{
		"name": "updated name"
	}`
	httpReq, err := http.NewRequest(http.MethodPatch, "http://localhost:8080/users/me", strings.NewReader(body))
	httpReq.Header.Add("Content-Type", "application/json")
	httpReq.Header.Add("Authorization", fmt.Sprintf("Bearer %s", token))
	if err != nil {
		panic(err)
	}
	// Test用サーバにリクエストを送信して、レスポンスをOpenAPI仕様に照らし合わせる
	recorder := ServeAndRequest(httpReq)
	// テストケース固有のチェック
	assert.Equal(t, 200, recorder.Result().StatusCode)
}

func TestPatchUserMeSuccessEmailColumn(t *testing.T) {
	createTestData() // テストデータの準備
	token := login() // 認証実行
	// HTTPリクエストの生成
	body := `{
		"email": "piyo@example.com"
	}`
	httpReq, err := http.NewRequest(http.MethodPatch, "http://localhost:8080/users/me", strings.NewReader(body))
	httpReq.Header.Add("Content-Type", "application/json")
	httpReq.Header.Add("Authorization", fmt.Sprintf("Bearer %s", token))
	if err != nil {
		panic(err)
	}
	// Test用サーバにリクエストを送信して、レスポンスをOpenAPI仕様に照らし合わせる
	recorder := ServeAndRequest(httpReq)
	// テストケース固有のチェック
	assert.Equal(t, 200, recorder.Result().StatusCode)
}

func TestPatchUserMeSuccessPasswordColumn(t *testing.T) {
	createTestData() // テストデータの準備
	token := login() // 認証実行
	// HTTPリクエストの生成
	body := `{
		"password": "password2"
	}`
	httpReq, err := http.NewRequest(http.MethodPatch, "http://localhost:8080/users/me", strings.NewReader(body))
	httpReq.Header.Add("Content-Type", "application/json")
	httpReq.Header.Add("Authorization", fmt.Sprintf("Bearer %s", token))
	if err != nil {
		panic(err)
	}
	// Test用サーバにリクエストを送信して、レスポンスをOpenAPI仕様に照らし合わせる
	recorder := ServeAndRequest(httpReq)
	// テストケース固有のチェック
	assert.Equal(t, 200, recorder.Result().StatusCode)
}

func TestPatchUserMeSuccessNoColumn(t *testing.T) {
	createTestData() // テストデータの準備
	token := login() // 認証実行
	// HTTPリクエストの生成
	body := `{}`
	httpReq, err := http.NewRequest(http.MethodPatch, "http://localhost:8080/users/me", strings.NewReader(body))
	httpReq.Header.Add("Content-Type", "application/json")
	httpReq.Header.Add("Authorization", fmt.Sprintf("Bearer %s", token))
	if err != nil {
		panic(err)
	}
	// Test用サーバにリクエストを送信して、レスポンスをOpenAPI仕様に照らし合わせる
	recorder := ServeAndRequest(httpReq)
	// テストケース固有のチェック
	assert.Equal(t, 200, recorder.Result().StatusCode)
}

func TestPatchUserMeInvalidColumn(t *testing.T) {
	createTestData() // テストデータの準備
	token := login() // 認証実行
	// HTTPリクエストの生成
	body := `{
		"role": "user"
	}`
	httpReq, err := http.NewRequest(http.MethodPatch, "http://localhost:8080/users/me", strings.NewReader(body))
	httpReq.Header.Add("Content-Type", "application/json")
	httpReq.Header.Add("Authorization", fmt.Sprintf("Bearer %s", token))
	if err != nil {
		panic(err)
	}
	// Test用サーバにリクエストを送信して、レスポンスをOpenAPI仕様に照らし合わせる
	recorder := ServeAndRequest(httpReq)
	// テストケース固有のチェック
	assert.Equal(t, 400, recorder.Result().StatusCode)
}

func TestPatchUserMeNoLogin(t *testing.T) {
	createTestData() // テストデータの準備
	// HTTPリクエストの生成
	body := `{
		"name": "updated name",
		"email": "piyo@example.com",
		"password": "password2"
	}`
	httpReq, err := http.NewRequest(http.MethodPatch, "http://localhost:8080/users/me", strings.NewReader(body))
	httpReq.Header.Add("Content-Type", "application/json")
	httpReq.Header.Add("Authorization", fmt.Sprintf("Bearer %s", "token"))
	if err != nil {
		panic(err)
	}
	// Test用サーバにリクエストを送信して、レスポンスをOpenAPI仕様に照らし合わせる
	recorder := ServeAndRequest(httpReq)
	// テストケース固有のチェック
	assert.Equal(t, 401, recorder.Result().StatusCode)
}

func TestDeleteUserSuccess(t *testing.T) {
	createTestData() // テストデータの準備
	token := login() // 認証実行
	// HTTPリクエストの生成
	httpReq, err := http.NewRequest(http.MethodDelete, "http://localhost:8080/users/1", nil)
	httpReq.Header.Add("Content-Type", "application/json")
	httpReq.Header.Add("Authorization", fmt.Sprintf("Bearer %s", token))
	if err != nil {
		panic(err)
	}
	// Test用サーバにリクエストを送信して、レスポンスをOpenAPI仕様に照らし合わせる
	recorder := ServeAndRequest(httpReq)
	// テストケース固有のチェック
	assert.Equal(t, 204, recorder.Result().StatusCode)
	//
	httpReq2, err2 := http.NewRequest(http.MethodGet, "http://localhost:8080/users", nil)
	httpReq2.Header.Add("Content-Type", "application/json")
	httpReq2.Header.Add("Authorization", fmt.Sprintf("Bearer %s", token))
	if err2 != nil {
		panic(err)
	}
	recorder2 := ServeAndRequest(httpReq2)
	// テストケース固有のチェック
	assert.Equal(t, 200, recorder2.Result().StatusCode)
	var body []User
	json.Unmarshal(recorder2.Body.Bytes(), &body)
	assert.Equal(t, 1, len(body))

}

func TestDeleteUserNoRecord(t *testing.T) {
	createTestData() // テストデータの準備
	token := login() // 認証実行
	// HTTPリクエストの生成
	httpReq, err := http.NewRequest(http.MethodDelete, "http://localhost:8080/users/123", nil)
	httpReq.Header.Add("Content-Type", "application/json")
	httpReq.Header.Add("Authorization", fmt.Sprintf("Bearer %s", token))
	if err != nil {
		panic(err)
	}
	// Test用サーバにリクエストを送信して、レスポンスをOpenAPI仕様に照らし合わせる
	recorder := ServeAndRequest(httpReq)
	// テストケース固有のチェック
	assert.Equal(t, 404, recorder.Result().StatusCode)
}

func TestDeleteUserNoLogin(t *testing.T) {
	createTestData() // テストデータの準備
	// HTTPリクエストの生成
	httpReq, err := http.NewRequest(http.MethodDelete, "http://localhost:8080/users/1", nil)
	httpReq.Header.Add("Content-Type", "application/json")
	httpReq.Header.Add("Authorization", fmt.Sprintf("Bearer %s", "token"))
	if err != nil {
		panic(err)
	}
	// Test用サーバにリクエストを送信して、レスポンスをOpenAPI仕様に照らし合わせる
	recorder := ServeAndRequest(httpReq)
	// テストケース固有のチェック
	assert.Equal(t, 401, recorder.Result().StatusCode)
}

func TestLoginSuccess(t *testing.T) {
	createTestData() // テストデータの準備
	// HTTPリクエストの生成
	body := `{
		"email": "hoge@example.com",
	    "password": "password"
	}`
	httpReq, err := http.NewRequest(http.MethodPost, "http://localhost:8080/users/login", strings.NewReader(body))
	httpReq.Header.Add("Content-Type", "application/json")
	if err != nil {
		panic(err)
	}
	// Test用サーバにリクエストを送信して、レスポンスをOpenAPI仕様に照らし合わせる
	recorder := ServeAndRequest(httpReq)
	// テストケース固有のチェック
	assert.Equal(t, 200, recorder.Result().StatusCode)
}

func TestLoginInvalidPassword(t *testing.T) {
	createTestData() // テストデータの準備
	// HTTPリクエストの生成
	body := `{
		"email": "hoge@example.com",
    "password": "invalid_password"
	}`
	httpReq, err := http.NewRequest(http.MethodPost, "http://localhost:8080/users/login", strings.NewReader(body))
	httpReq.Header.Add("Content-Type", "application/json")
	if err != nil {
		panic(err)
	}
	// Test用サーバにリクエストを送信して、レスポンスをOpenAPI仕様に照らし合わせる
	recorder := ServeAndRequest(httpReq)
	// テストケース固有のチェック
	assert.Equal(t, 401, recorder.Result().StatusCode)
}

func TestLoginInvalidEmail(t *testing.T) {
	createTestData() // テストデータの準備
	// HTTPリクエストの生成
	body := `{
		"email": "invalid@example.com",
    "password": "password"
	}`
	httpReq, err := http.NewRequest(http.MethodPost, "http://localhost:8080/users/login", strings.NewReader(body))
	httpReq.Header.Add("Content-Type", "application/json")
	if err != nil {
		panic(err)
	}
	// Test用サーバにリクエストを送信して、レスポンスをOpenAPI仕様に照らし合わせる
	recorder := ServeAndRequest(httpReq)
	// テストケース固有のチェック
	assert.Equal(t, 401, recorder.Result().StatusCode)
}

func TestRefreshTokenSuccess(t *testing.T) {
	createTestData() // テストデータの準備
	token := login() // 認証実行
	// HTTPリクエストの生成
	httpReq, err := http.NewRequest(http.MethodGet, "http://localhost:8080/users/refresh_token", nil)
	httpReq.Header.Add("Content-Type", "application/json")
	httpReq.Header.Add("Authorization", fmt.Sprintf("Bearer %s", token))
	if err != nil {
		panic(err)
	}
	// Test用サーバにリクエストを送信して、レスポンスをOpenAPI仕様に照らし合わせる
	recorder := ServeAndRequest(httpReq)
	// テストケース固有のチェック
	assert.Equal(t, 200, recorder.Result().StatusCode)
}

func TestRefreshTokenNoLogin(t *testing.T) {
	createTestData() // テストデータの準備
	// HTTPリクエストの生成
	httpReq, err := http.NewRequest(http.MethodGet, "http://localhost:8080/users/refresh_token", nil)
	httpReq.Header.Add("Content-Type", "application/json")
	httpReq.Header.Add("Authorization", fmt.Sprintf("Bearer %s", "token"))
	if err != nil {
		panic(err)
	}
	// Test用サーバにリクエストを送信して、レスポンスをOpenAPI仕様に照らし合わせる
	recorder := ServeAndRequest(httpReq)
	// テストケース固有のチェック
	assert.Equal(t, 401, recorder.Result().StatusCode)
}

func TestLogoutSuccess(t *testing.T) {
	createTestData() // テストデータの準備
	token := login() // 認証実行
	// HTTPリクエストの生成
	httpReq, err := http.NewRequest(http.MethodPost, "http://localhost:8080/users/logout", nil)
	httpReq.Header.Add("Content-Type", "application/json")
	httpReq.Header.Add("Authorization", fmt.Sprintf("Bearer %s", token))
	if err != nil {
		panic(err)
	}
	// Test用サーバにリクエストを送信して、レスポンスをOpenAPI仕様に照らし合わせる
	recorder := ServeAndRequest(httpReq)
	// テストケース固有のチェック
	assert.Equal(t, 200, recorder.Result().StatusCode)
}

func TestLogoutNoLogin(t *testing.T) {
	createTestData() // テストデータの準備
	// HTTPリクエストの生成
	httpReq, err := http.NewRequest(http.MethodPost, "http://localhost:8080/users/logout", nil)
	httpReq.Header.Add("Content-Type", "application/json")
	httpReq.Header.Add("Authorization", fmt.Sprintf("Bearer %s", "token"))
	if err != nil {
		panic(err)
	}
	// Test用サーバにリクエストを送信して、レスポンスをOpenAPI仕様に照らし合わせる
	recorder := ServeAndRequest(httpReq)
	// テストケース固有のチェック
	assert.Equal(t, 200, recorder.Result().StatusCode)
}
