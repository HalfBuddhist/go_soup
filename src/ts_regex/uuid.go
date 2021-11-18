// 匹配 UUID
package ts_regex

import (
	"fmt"
	"regexp"
)

func TS_regex_uuid() {
	fmt.Println("hello")

	pattern := `[0-9a-f]{8}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{12}`
	reg := regexp.MustCompile(pattern)
	res := reg.Match([]byte("3947c44f-c1a0-48ee-b217-4bb4a034cf73"))
	fmt.Println(res)
}
