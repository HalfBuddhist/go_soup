package ts_select

import "testing"

// 测试空多路选择的作用，阻塞当前进程。
func TestEmptySelectBlock(t *testing.T){
	select {}
}