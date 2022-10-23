package ts_type

import (
	"fmt"
	"testing"
)

// 底层类型相同的两个自定义类型是否可以强制进行转化，答案是可以。
// 这里使用 map[string]interface{} 来进行实验。
func TestMapConvert(t *testing.T) {
	type ResponseMap map[string]interface{}
	type ResponseMap2 map[string]interface{}

	map1 := ResponseMap{
		"abc": "abc",
		"123": "123",
	}
	map2 := ResponseMap2(map1)
	fmt.Println(map2)
}

// 空接口断言类型失败导致，不能断言成底层类型一样的不同声明类型。
func TestNullInterfaceMapAssert(t *testing.T) {
	type ResponseMap map[string]interface{}
	type ResponseMap2 map[string]interface{}

	var map1 interface{}
	map1 = ResponseMap{
		"abc": "abc",
		"123": "123",
	}
	map2 := map1.(ResponseMap2)
	fmt.Println(map2)
}

// 空接口不能强制转化类型，只能断言
// 断言只能断言到声明时的真实的类型，断言为相同底层类型的不同声明类型时亦会出错。
// 此时可以分为两步，先断言为声明类型，再强转为底层类型相同的其它声明类型
func TestNullInterfaceMapConvert(t *testing.T) {
	type ResponseMap map[string]interface{}
	type ResponseMap2 map[string]interface{}

	var map1 interface{}
	map1 = ResponseMap{
		"abc": "abc",
		"123": "123",
	}
	// map2 := ResponseMap2(map1) // err
	map2 := map1.(ResponseMap)
	fmt.Println(map2)
	// map3 := map1.(ResponseMap2)  // err
	// fmt.Println(map3)
	map3 := ResponseMap2(map2)
	fmt.Println(map3)
}
