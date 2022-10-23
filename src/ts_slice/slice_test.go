package ts_slice

import (
	"fmt"
	"testing"
)

// 可以迭代空切片，其也可以取长度
func TestIterateOnNilSlice(t *testing.T) {
	type abc struct {
		x string
		y string
	}
	// y := &abc{"123", "456"}
	// var x []*abc = nil
	x := []*abc{{"123", "456"}, {"1234", "3456"}}
	fmt.Printf("Length: %d\n", len(x))
	for idx, item := range x {
		fmt.Println(idx, item)
	}
}
