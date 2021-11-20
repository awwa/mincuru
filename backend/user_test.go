package main

import (
	"bytes"
	"context"
	"net/http"

	// "net/http/httptest"
	"testing"

	"github.com/getkin/kin-openapi/openapi3"
	"github.com/getkin/kin-openapi/openapi3filter"
	"github.com/getkin/kin-openapi/routers/legacy"
)

// func TestGetUsers(t *testing.T) {
// 	ctx := context.Background()
// 	loader := &openapi3.Loader{Context: ctx}
// 	doc, err := loader.LoadFromFile("../openapi.yaml")
// 	if err != nil {
// 		panic(err)
// 	}
// 	err2 := doc.Validate(ctx)
// 	if err2 != nil {
// 		panic(err2)
// 	}
// 	router, err3 := legacy.NewRouter(doc)
// 	if err3 != nil {
// 		panic(err3)
// 	}
// 	httpReq, err4 := http.NewRequest(http.MethodGet, "http://localhost:4010/users", nil)
// 	httpReq.Header.Set("Authorization", "Bearer tokentokentoken")
// 	if err4 != nil {
// 		panic(err4)
// 	}
// 	// fmt.Println("httpReq:", httpReq)

// 	// Find route
// 	route, pathParams, err5 := router.FindRoute(httpReq)
// 	if err5 != nil {
// 		panic(err5)
// 	}

// 	// Validate request
// 	requestValidationInput := &openapi3filter.RequestValidationInput{
// 		Request:    httpReq,
// 		PathParams: pathParams,
// 		Route:      route,
// 	}
// 	// if err := openapi3filter.ValidateRequest(ctx, requestValidationInput); err != nil {
// 	// 	panic(err)
// 	// }

// 	var (
// 		respStatus      = 200
// 		respContentType = "application/json"
// 		respBody        = bytes.NewBufferString(
// 			`[
//         {
//           "id": 1,
//           "name": "hoge fuga",
//           "email": "hoge1@example.com",
//           "role": "user"
//         },
//         {
//             "id": 2,
//             "name": "hoge fuga2",
//             "email": "hoge2@example.com",
//             "role": "admin"
//         }
//       ]`,
// 		)
// 	)

// 	responseValidationInput := &openapi3filter.ResponseValidationInput{
// 		RequestValidationInput: requestValidationInput,
// 		Status:                 respStatus,
// 		Header:                 http.Header{"Content-Type": []string{respContentType}},
// 	}
// 	// fmt.Println("Response:", responseValidationInput)
// 	if respBody != nil {
// 		data, _ := json.Marshal(respBody)
// 		responseValidationInput.SetBodyBytes(data)
// 	}
// 	fmt.Println("data:", respBody)

// 	// Validate response
// 	if err := openapi3filter.ValidateResponse(ctx, responseValidationInput); err != nil {
// 		panic(err)
// 	}
// }

func TestGetUser(t *testing.T) {
	ctx := context.Background()
	loader := &openapi3.Loader{Context: ctx}
	doc, err := loader.LoadFromFile("../openapi.yaml")
	if err != nil {
		panic(err)
	}
	err2 := doc.Validate(ctx)
	if err2 != nil {
		panic(err2)
	}
	router, err3 := legacy.NewRouter(doc)
	if err3 != nil {
		panic(err3)
	}
	httpReq, err4 := http.NewRequest(http.MethodGet, "http://localhost:4010/users/123", nil)
	httpReq.Header.Set("Authorization", "Bearer tokentokentoken")
	if err4 != nil {
		panic(err4)
	}
	// httpRes, err6 := http.DefaultClient.Do(httpReq)
	// if err6 != nil {
	// 	panic(err6)
	// }
	// fmt.Println("httpRes:", httpRes)
	// fmt.Println("httpRes.StatusCode:", httpRes.StatusCode)
	// fmt.Println("httpRes.Header:", httpRes.Header)
	// body, _ := ioutil.ReadAll(httpRes.Body)
	// fmt.Println("httpRes.Body:", body)
	// fmt.Println("httpReq:", httpReq)

	// Find route
	route, pathParams, err5 := router.FindRoute(httpReq)
	if err5 != nil {
		panic(err5)
	}

	// Validate request
	requestValidationInput := &openapi3filter.RequestValidationInput{
		Request:    httpReq,
		PathParams: pathParams,
		Route:      route,
	}
	// // if err := openapi3filter.ValidateRequest(ctx, requestValidationInput); err != nil {
	// // 	panic(err)
	// // }

	var (
		respStatus      = 200
		respContentType = "application/json"
		respBody        = bytes.NewBufferString(`{"id": 1, "name": "hoge fuga", "email": "hoge1@example.com", "role": "user"}`)
	)

	responseValidationInput := &openapi3filter.ResponseValidationInput{
		RequestValidationInput: requestValidationInput,
		Status:                 respStatus,
		Header:                 http.Header{"Content-Type": []string{respContentType}},
	}
	if respBody != nil {
		data := respBody.Bytes()
		responseValidationInput.SetBodyBytes(data)
	}

	// Validate response
	if err := openapi3filter.ValidateResponse(ctx, responseValidationInput); err != nil {
		panic(err)
	}
}

// func TestGetUsers(t *testing.T) {
// 	loader := openapi3.NewLoader()
// 	doc, err := loader.LoadFromFile("../openapi.yaml")
// 	if err != nil {
// 		panic(err)
// 	}
// 	err = doc.Validate(loader.Context)
// 	if err != nil {
// 		panic(err)
// 	}

// 	req, err := http.NewRequest(http.MethodGet, "http://localhost:8080/users", nil /*bytes.NewReader(p)*/)
// 	if err != nil {
// 		panic(err)
// 	}
// 	req.Header.Set("Content-Type", "application/json")
// 	// fmt.Println(req)
// 	router, err := gorillamux.NewRouter(doc)
// 	// router, err := legacy.NewRouter(doc)
// 	if err != nil {
// 		panic(err)
// 	}
// 	route, pathParams, err := router.FindRoute(req)
// 	// route, pathParams, err := router.FindRoute(req)
// 	if err != nil {
// 		panic(err)
// 	}

// 	requestValidationInput := &openapi3filter.RequestValidationInput{
// 		Request:    req,
// 		PathParams: pathParams,
// 		Route:      route,
// 	}
// 	if err := openapi3filter.ValidateRequest(loader.Context, requestValidationInput); err != nil {
// 		fmt.Println(err)
// 	}
// 	// Output:
// 	// request body has an error: doesn't match the schema: input matches more than one oneOf schemas
// }

// // ユーザ一覧取得テスト
// func TestGetUsers(t *testing.T) {
// 	client := &http.Client{}
// 	req, err := http.NewRequest(
// 		http.MethodGet, "http://localhost:8080/users?email=hoge@example.com", nil,
// 	)
// 	// 関数を抜ける際に必ずresponseをcloseするようにdeferでcloseを呼ぶ
// 	if err != nil {
// 		panic(err) //t.Fatalf("geterror %v", err)
// 	}
// 	defer req.Body.Close()
// 	// CheckRequest(req)
// 	// if e != nil {
// 	// 	panic(e)
// 	// }
// 	resp, err := client.Do(req)
// 	if resp.StatusCode != 200 {
// 		t.Fatalf("%v", resp.StatusCode)
// 	}
// 	if err != nil {
// 		panic(err) //t.Fatalf("geterror %v", err)
// 	}

// 	// for i := range encrypttests {
// 	// 	test := &encrypttests[i]
// 	// 	actual, err := Encrypt(test.in, test.sh)
// 	// 	if test.enc != actual {
// 	// 		t.Errorf("Test failed: Encrypt('%s', %d) = '%s', %v want '%s', %v",
// 	// 			test.in, test.sh, actual, test.err, test.enc, err)
// 	// 	}
// 	// }
// }

// // https://github.com/getkin/kin-openapi/blob/master/routers/legacy/validate_request_test.go
// func CheckRequest(req *http.Request) {
// 	ts := httptest.NewServer(Routing())
// 	defer ts.Close()
// 	loader := openapi3.NewLoader()
// 	doc, err := loader.LoadFromFile("../openapi.yaml")
// 	if err != nil {
// 		panic(err)
// 	}
// 	if err := doc.Validate(loader.Context); err != nil {
// 		panic(err)
// 	}
// 	router, err := legacy.NewRouter(doc)
// 	if err != nil {
// 		panic(err)
// 	}
// 	route, pathParams, err := router.FindRoute(req)
// 	if err != nil {
// 		panic(err)
// 	}
// 	requestValidationInput := &openapi3filter.RequestValidationInput{
// 		Request:    req,
// 		PathParams: pathParams,
// 		Route:      route,
// 	}
// 	err = openapi3filter.ValidateRequest(loader.Context, requestValidationInput)
// 	if err != nil {
// 		panic(err)
// 	}
// 	// Output:
// 	// request body has an error: doesn't match the schema: input matches more than one oneOf schemas
// 	// route, pathParams, e := openAPIRouter.FindRoute(request.Method, request.URL)
// 	// if e != nil {
// 	// 	return e
// 	// }

// 	// u, _ := url.Parse(ts.URL)
// 	// req.URL.Scheme = u.Scheme
// 	// req.URL.Host = u.Host
// 	// resp, e := http.DefaultClient.Do(req)
// 	// if e != nil {
// 	// 	return e
// 	// }
// 	// defer resp.Body.Close()

// 	// body, e := ioutil.ReadAll(response.Body)
// 	// if e != nil {
// 	// 	return e
// 	// }

// 	// requestValidationInput := &openapi3filter.RequestValidationInput{
// 	// 	Request:    request,
// 	// 	PathParams: pathParams,
// 	// 	Route:      route,
// 	// }

// 	// responseValidationInput := &openapi3filter.ResponseValidationInput{
// 	// 	RequestValidationInput: requestValidationInput,
// 	// 	Status:                 resp.StatusCode,
// 	// 	Header:                 resp.Header,
// 	// }
// 	// responseValidationInput.SetBodyBytes(body)

// 	// return openapi3filter.ValidateResponse(context.TODO(), responseValidationInput)
// }
