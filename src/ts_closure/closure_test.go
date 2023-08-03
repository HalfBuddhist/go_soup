package ts_closure

import (
	"fmt"
	"testing"
)

/**
 * 测试闭包中的MAP是引用还是拷贝；
 * 是引用。
 */
func TestMapInClosure(t *testing.T) {
	m := make(map[string]string)
	key := "foo"
	closureF := func() {
		fmt.Println(m[key])
	}
	closureF()

	m["foo"] = "bar"
	m["hello"] = "world"
	// key = "hello"
	closureF()
}
