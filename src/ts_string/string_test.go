package ts_string

import (
	"fmt"
	"strings"
	"testing"
)

func TestCutPrefix(t *testing.T) {
	src := "abcdefgha"
	after, _ := strings.CutPrefix(src, "abcd")
	fmt.Println(after)
}

func TestTitle(t *testing.T) {
	src := "abcdefgha hello world"
	after := strings.Title(src)
	fmt.Println(after)

	after2 := strings.ToTitle(src)
	fmt.Print(after2)
}
