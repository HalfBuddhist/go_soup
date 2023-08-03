package ts_error

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

// test: nil error 是否可以正常字符串输出。
// conclusion: 可以正常输出，但不是空串，而是 『%!s(<nil>)』
func TestNilError(t *testing.T) {
	var err error = nil
	a := ""
	fmt.Printf("xxxxx%sxxx%sxxx\n", err, a)
}

func Hello() (string, error) {
	return "world", fmt.Errorf("hello")
}

// 子作用域内重新赋值就是新的变量（接口）了
func TestOverrideDomain(t *testing.T) {
	var err error
	assert.Empty(t, err)

	if "a" == "a" {
		res, err := Hello()
		fmt.Printf("%s, %v\n", res, err)
	}

	fmt.Printf("%v\n", err)
}
