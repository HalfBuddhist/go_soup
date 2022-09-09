package ts_viper

import (
	"fmt"
	"strings"
	"testing"

	"github.com/spf13/viper"
)

// 测试目的: 1. AllKeys 方法是否会强制以小写形式返回
//          2. 取值时是否 key 大小写区分
// 结论：    1. 是的。
//          2. 不区分，都可正常取值。
func TestKeyUpperOrLower(t *testing.T) {
	localViper := viper.New()
	configPathList := []string{"test.toml"}
	for _, configPath := range configPathList {
		configPath = strings.TrimSpace(configPath)
		localViper.SetConfigFile(configPath)
		err := localViper.MergeInConfig()
		if err != nil {
			panic(err)
		}
	}
	localViper.BindEnv("orca_node_ip", "ORCA_NODE_IP")
	keys := localViper.AllKeys()
	for _, key := range keys {
		fmt.Println(key)
	}

	fmt.Println(localViper.GetString("EDF.path"))
	fmt.Println(localViper.GetString("edf.path"))
	fmt.Println(localViper.GetString("edf.Path"))
	fmt.Println(localViper.GetString("edf.PATH"))
}
