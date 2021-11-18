package main

import (
	"net/http"
	"testing"
)

// ユーザ一覧取得テスト
func TestGetUsers(t *testing.T) {
	resp, err := http.Get("http://localhost:8080/users?email=hoge@example.com")
	if err != nil {
		t.Fatalf("geterror %v", err)
		return
	}
	if resp.StatusCode != 200 {
		t.Fatalf("%v", resp.StatusCode)
	}
	// 関数を抜ける際に必ずresponseをcloseするようにdeferでcloseを呼ぶ
	defer resp.Body.Close()

	// for i := range encrypttests {
	// 	test := &encrypttests[i]
	// 	actual, err := Encrypt(test.in, test.sh)
	// 	if test.enc != actual {
	// 		t.Errorf("Test failed: Encrypt('%s', %d) = '%s', %v want '%s', %v",
	// 			test.in, test.sh, actual, test.err, test.enc, err)
	// 	}
	// }
}
