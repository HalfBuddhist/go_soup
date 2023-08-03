package ts_encoding

import (
	"encoding/base64"
	"fmt"
	"testing"
)

func TestBase64(t *testing.T) {
	src := `{"auths":{"https://xxx.io":{"password":"Teco@135","username":"admin"}}}`
	dst := base64.StdEncoding.EncodeToString([]byte(src))
	fmt.Println(dst)

	src2, err := base64.StdEncoding.DecodeString(dst)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(string(src2))
}
