package ts_defer

import (
	"fmt"
	"testing"
)

type people struct {
	name   string
	gender string
}

func (p *people) Description() string {
	return fmt.Sprintf("%s is %s", p.name, p.gender)
}

// 测试语句defer后，对语句中的变量进行赋值的影响。
// Result: 不影响，应该是defer会进行一个表达式求值的过程，但函数不执行。
// Result: defer 多个语句的执行顺序是，栈序。
func TestAssigneVarAfterDefer(t *testing.T) {
	p := &people{"tom", "male"}
	defer fmt.Println(p.Description())

	p = &people{"alice", "female"}
	p.name = "Lily"
	defer fmt.Println(p.Description())
}
