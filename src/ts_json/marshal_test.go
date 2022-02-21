package ts_json

import (
	"encoding/json"
	"fmt"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

// json 转化为字符串
func TestMarshal(t *testing.T) {
	jsonObj := map[string]interface{}{
		"key1": 11,
		"key2": 22,
	}
	jsonStr, err := json.Marshal(jsonObj)
	assert.NoError(t, err)
	fmt.Println(jsonStr)
	fmt.Println(string(jsonStr))
}

// json 字符串转化为 map 对象
// 注意 unmarshal 其会将数字转化为 float64。
func TestUnmarshal(t *testing.T) {
	jsonStr := `{"key1":11,"key2":22}`
	jsonObj := map[string]interface{}{}
	json.Unmarshal([]byte(jsonStr), &jsonObj)
	fmt.Println(jsonObj)
	fmt.Printf("%#v:%T\n", jsonObj["key1"], jsonObj["key1"])
}

// json 字符串转化为 map 对象
// 将数字转化为 json.Number, 实际就是字符串。
// 如果想要转化成对应的类型，比如 long, int, 
// 还是要将字符串 decode 或者 unmarhsal 到拥有具体类型的 struct 中去，期间类型会自动转化的。
func TestDecoderUnmarshal(t *testing.T){
	jsonStr := `{"key1":11,"key2":22}`
	jsonObj := map[string]interface{}{}
	decoder := json.NewDecoder(strings.NewReader(jsonStr))
	decoder.UseNumber()
	err := decoder.Decode(&jsonObj)
	assert.NoError(t, err)
	fmt.Println(jsonObj)
	fmt.Printf("%#v:%T\n", jsonObj["key1"], jsonObj["key1"])
}
