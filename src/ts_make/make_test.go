package ts_make

import (
	"fmt"
	"testing"
)

type TestStruct struct {
	Name string
}

// make slice时 长度为非0时会如何初始化，
// int slice 会初始化为0.
// 指针 slice 会初始化为nil.
func TestMakeSlice(t *testing.T) {
	a := make([]int, 10, 10)
	fmt.Println(len(a))
	for idx, ele := range a {
		fmt.Printf("%d:\t %d\n", idx, ele)
	}

	b := make([]*TestStruct, 10, 10)
	fmt.Println(len(b))
	for idx, ele := range b {
		fmt.Printf("%d:\t %v\n", idx, ele)
	}
}
