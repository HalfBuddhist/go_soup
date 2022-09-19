package ts_minmax

import (
	"fmt"
	"testing"
)

// 测试取得 int 与 uint 的最大最小值
func TestIntMinMax(t *testing.T) {
	const INT_MAX = int(^uint(0) >> 1)
	const INT_MIN = ^INT_MAX
	fmt.Println(INT_MAX)
	fmt.Println(INT_MIN)

	const UINT_MIN uint = 0
	const UINT_MAX uint = ^uint(0)
	fmt.Println(UINT_MIN, UINT_MAX)
}
