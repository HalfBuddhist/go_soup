package ts_type_assert

import (
	"fmt"
	"testing"
)

// 测试 map 类型的断言能否将其值部分断言成空接口类型：
// result: 不能；对于 map，不同的值类型组成的复合类型是不同的类型，不能互相自动转化。
func TestMapTypeAssert(t *testing.T) {

	tmap := map[string]string{
		"a": "b",
		"c": "d",
	}
	fmt.Printf("%T\n", tmap)

	bmap := map[string]interface{}{
		"a": "b",
		"c": "d",
	}
	fmt.Printf("%T\n", bmap)

	var tt interface{} = tmap
	if _, ok := tt.(map[string]interface{}); ok {
		fmt.Println("Could assert the map value part to interface{}.")
	} else {
		fmt.Println("Could not assert the map value part to interface{}.")
	}
}
