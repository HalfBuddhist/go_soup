// test slice convert in O(1) using reflect.
package main

import (
	"fmt"
	"reflect"
	"unsafe"
)

type IDO struct {
	id string
}

type IDObject interface {
	ID() string
}

func (i *IDO) ID() string {
	return i.id
}

/* 将 []X 转换为 []Y 切片

// Convert slice to new slice with different element type in O(1).

// Could convert between different named type.

// Not working when the newSliceType is interface slice.

注意事项
该转换操作有一定的风险，用户需要自己保证安全。主要涉及以下几种类型：

当结构体中含有指针时，转换会导致垃圾回收的问题。
如果是 []byte 转 []T 可能会导致起始地址未对齐的问题 （[]byte 有可能从奇数位置切片）。
该转换操作可能依赖当前系统，不同类型的处理器之间有差异。
该转换操作的优势是性能和类似void*的泛型，与cgo接口配合使用会更加理想。
*/
func Slice(slice interface{}, newSliceType reflect.Type) interface{} {
	sv := reflect.ValueOf(slice)
	if sv.Kind() != reflect.Slice {
		panic(fmt.Sprintf("Slice called with non-slice value of type %T", slice))
	}
	if newSliceType.Kind() != reflect.Slice {
		panic(fmt.Sprintf("Slice called with non-slice type of type %T", newSliceType))
	}
	newSlice := reflect.New(newSliceType)
	hdr := (*reflect.SliceHeader)(unsafe.Pointer(newSlice.Pointer()))
	hdr.Cap = sv.Cap() * int(sv.Type().Elem().Size()) / int(newSliceType.Elem().Size())
	hdr.Len = sv.Len() * int(sv.Type().Elem().Size()) / int(newSliceType.Elem().Size())
	hdr.Data = uintptr(sv.Pointer())
	return newSlice.Elem().Interface()
}


// 将 []T 切片转换为 []byte
func ByteSlice(slice interface{}) (data []byte) {
	sv := reflect.ValueOf(slice)
	if sv.Kind() != reflect.Slice {
		panic(fmt.Sprintf("ByteSlice called with non-slice value of type %T", slice))
	}
	h := (*reflect.SliceHeader)(unsafe.Pointer(&data))
	h.Cap = sv.Cap() * int(sv.Type().Elem().Size())
	h.Len = sv.Len() * int(sv.Type().Elem().Size())
	h.Data = sv.Pointer()
	return
}

type RGB struct {
	R, G, B uint8
}
type BGR struct {
	B, G, R uint8
}

func RGB2BGR(data []RGB) []BGR {
	d := Slice(data, reflect.TypeOf([]BGR(nil)))
	return d.([]BGR)
}

func test_covert_2_interface_slice() {
	var abc []*IDO
	for i:=0; i<10; i++ {
		abc =append(abc, &IDO{fmt.Sprintf("%d", i)})
	}
	var ttt []IDObject
	aaa := Slice(abc, reflect.TypeOf(ttt))

	fmt.Printf("%T\n", aaa)

	_, ok := aaa.([]IDObject)
	if ok {
		fmt.Println("aaa is []IDObject")
	}	else {
		fmt.Println("aaa is not []IDObject")
	}

	_, ok = aaa.([]*IDO)
	if ok {
		fmt.Println("aaa is []IDO")
	}	else {
		fmt.Println("aaa is not []IDO")
	}

	bbb, ok := aaa.([]IDObject)
	for _, item := range bbb {
		fmt.Println(item.ID())
	}
}

type age int

func main()  {
	var abc []age
	for i:=0; i<10; i++ {
		abc =append(abc, age(i))
	}
	fmt.Printf("%#v\n", abc)

	var ttt []int
	aaa := Slice(abc, reflect.TypeOf(ttt))
	fmt.Printf("%T\n", aaa)

	bbb, ok := aaa.([]int)
	if !ok {
		fmt.Println("aaa is not []int")
		return
	}
	fmt.Println("aaa is []int")
	for _, item := range bbb {
		fmt.Println(item)
	}
}
