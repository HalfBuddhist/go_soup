package ts_cast

import (
	"fmt"
	"testing"

	"github.com/spf13/cast"
	"github.com/stretchr/testify/assert"
)

func TestCastStringToInteger(t *testing.T) {
	var s string = "123"
	res := cast.ToUint64(s)
	fmt.Println(res)
	assert.Equal(t, uint64(123), res)

	s = ""
	res = cast.ToUint64(s)
	fmt.Println(res)
	assert.Equal(t, uint64(0), res)

	s = "123a"
	res = cast.ToUint64(s)
	fmt.Println(res)
	assert.Equal(t, uint64(0), res)
}
