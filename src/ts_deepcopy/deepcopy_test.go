package ts_deepcopy

import "encoding/json"
import "fmt"
import "github.com/mohae/deepcopy"
import "github.com/stretchr/testify/assert"
import "testing"

type A struct {
	Name string `json:"name"`
}

type B struct {
	As []*A `json:"as"`
}

// 测试结构体内的指针切片中的元素是新的还是旧的
// 结论：是新的
func TestDeepCopyPtrSliceInStruct(t *testing.T) {
	b := &B{
		As: []*A{
			{"hello"}, {"world"},
		},
	}
	c := deepcopy.Copy(b)
	fmt.Printf("%+v\n", b)
	fmt.Printf("%+v\n", c)
	fmt.Printf("%T\n", c)
	str, err := json.Marshal(c)
	assert.NoError(t, err)
	fmt.Println(string(str))
}
