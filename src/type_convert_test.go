package main

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
