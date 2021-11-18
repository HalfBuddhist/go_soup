//测试：*可以匹配到.
//所以.应该是non-Separator character, 这里说的 non-Separator 应该是指的是/与\这种路径分隔符。

package main
import (
    "path/filepath"
    "fmt"
)

func main() {
    m,_ := filepath.Glob("/Users/administrator/workspace/go/src/go_soup/ts_glob/logs/*.log")
    fmt.Println(m)
}
