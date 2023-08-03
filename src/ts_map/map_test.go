package ts_map

import (
	"fmt"
	"testing"
)

// 必须初始化，否则空指针
func TestMapInit(t *testing.T) {
	var a map[string]string
	a = make(map[string]string)
	a["hello"] = "world"
	fmt.Printf("%v", a)
}

type s struct {
	Name string
	V    map[string]string
}

// 结构体内依然要初始化
func TestMapInitInStruct(t *testing.T) {
	b := s{
		Name: "hello",
		V:    make(map[string]string),
	}
	fmt.Printf("%v", b)

	b.V["hello"] = "world"
	fmt.Printf("%v", b)
}
