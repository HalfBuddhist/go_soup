package ts_defer

import (
	"fmt"
	"testing"
)

type People struct {
	name   string
	gender string
}

func (p *People) Description() string {
	return fmt.Sprintf("%s is %s", p.name, p.gender)
}

// 测试语句defer后，对语句中的变量进行赋值的影响。
// Result: 不影响，应该是defer会进行一个表达式求值的过程，但函数不执行。
// Result: defer 多个语句的执行顺序是，栈序。
func TestAssigneVarAfterDefer(t *testing.T) {
	p := &People{"tom", "male"}
	defer fmt.Println(p.Description())

	p = &People{"alice", "female"}
	p.name = "Lily"
	defer fmt.Println(p.Description())
}

// 疑惑return后面的defer到底会不会执行？
// defer 关键字后面的函数或者方法想要执行必须先注册，return 之后的 defer 是不能注册的，
// 也就不能执行后面的函数或方法。
func TestReturnAfterDefer(t *testing.T) {
	defer func() {
		fmt.Println("1")
	}()
	if true {
		fmt.Println("2")
		return
	}
	defer func() {
		fmt.Println("3")
	}()
}
