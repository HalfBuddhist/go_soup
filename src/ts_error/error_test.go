package ts_error

import (
	"fmt"
	"testing"
)

// test: nil error 是否可以正常字符串输出。
// conclusion: 可以正常输出，但不是空串，而是 『%!s(<nil>)』
func TestNilError(t *testing.T) {
	var err error = nil;
	a := ""
	fmt.Printf("xxxxx%sxxx%sxxx\n", err, a)
}
