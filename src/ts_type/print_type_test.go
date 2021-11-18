package ts_type

import (
	"fmt"
	"reflect"
	"testing"
)

// golang输出数据类型
func TestPrintType(t *testing.T) {
	decoded := "hello world."
	fmt.Println(reflect.TypeOf(decoded).String())
	fmt.Printf("%T\n", decoded)
}
