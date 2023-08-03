//测试：*可以匹配到.
//所以.应该是non-Separator character, 这里说的 non-Separator 应该是指的是/与\这种路径分隔符。

package main

import (
	"fmt"
	"path/filepath"
)

func main() {
	m, _ := filepath.Glob("/home/liuqw/workspace/*/sche")
	fmt.Println(m)
	for _, v := range m {
		fmt.Println(filepath.Base(v))
		fmt.Println(filepath.Dir(v))
	}
}
