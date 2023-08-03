package ts_context

import (
	"context"
	"fmt"
	"testing"
	"time"
)

type MyContext struct {
	// 这里的 Context 是我 copy 出来的，所以前面不用加 context.
	context.Context
}

/**
 * 封装后的 cancelCtx 作为父 context 是否会起新的监听协程以连接？
 * 结果：没有新的协程，直接加到了 children 里面。
 */
func TestContextWrapping(t *testing.T) {
	childCancel := true

	parentCtx, parentFunc := context.WithCancel(context.Background())
	mctx := MyContext{parentCtx}

	childCtx, childFun := context.WithCancel(mctx)

	if childCancel {
		childFun()
	} else {
		parentFunc()
	}

	fmt.Println(parentCtx) // 输出来源路径？
	fmt.Printf("%T\n", parentCtx)
	fmt.Println(mctx)
	fmt.Printf("%T\n", mctx)
	fmt.Println(childCtx)
	fmt.Printf("%T\n", childCtx)

	// 防止主协程退出太快，子协程来不及打印
	time.Sleep(10 * time.Second)
}
