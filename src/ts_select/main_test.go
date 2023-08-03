package ts_select

import (
	"fmt"
	"testing"
)

// 测试空多路选择的作用，阻塞当前进程。
func TestEmptySelectBlock(t *testing.T) {
	select {}
}

// 测试 select 有默认分支的情况，方便跳过阻塞而返回
func TestSelectFromEmpty(t *testing.T) {
	var a <-chan struct{} = nil
	select {
	case <-a:
		fmt.Println("read something from nil channel.")
	default:
	}
	fmt.Println("select exit from default branch.")
}

// 测试 select 有默认分支的情况，与普通channel的优先级
// 普通 channel 的优先级高
func TestSelectFromEmptyAndChannel(t *testing.T) {
	a := make(chan struct{}, 10)
	a <- struct{}{}
	select {
	case <-a:
		fmt.Println("read something from nil channel.")
	default:
		fmt.Println("select exit from default branch.")
	}
}
