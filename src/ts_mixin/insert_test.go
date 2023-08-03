package ts_xorm

import (
	"fmt"
	"testing"

	_ "github.com/jinzhu/gorm/dialects/mysql"
	_ "github.com/lib/pq"
)

type A struct {
	Name string
}

type B struct {
	Name string
}

type C struct {
	A
	B
}

// 混入相同的方法，初始化时不会出错，调用时会报错如下
// ambiguous selector c.Name
func TestMixinSameMethod(t *testing.T) {
	c := C{
		A: A{
			Name: "A",
		},
		B: B{
			Name: "B",
		},
	}
	fmt.Println(c)
	// fmt.Println(c.Name)
}
