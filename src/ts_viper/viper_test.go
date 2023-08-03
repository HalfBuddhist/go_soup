package ts_viper

import (
	"fmt"
	"reflect"
	"strings"
	"testing"

	"github.com/spf13/viper"
	_ "github.com/spf13/viper/remote" // to enable the remote features
)

var etcd_nodes = []string{
	"https://10.18.96.172:2379",
	"https://10.18.96.171:2379"}

const configKey = "/etcd_test"
const remoteConfType = "toml"
const remoteType = "etcd"
const fileName = "test.toml"

// 测试目的: 1. AllKeys 方法是否会强制以小写形式返回
//  2. 取值时是否 key 大小写区分
//
// 结论：    1. 是的。
//  2. 不区分，都可正常取值。
func TestKeyUpperOrLower(t *testing.T) {
	localViper := viper.New()
	configPathList := []string{fileName}
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

// 测试ETCD中值为列表的读取情况
// 结果：好像记得可以正确读取。
func TestReadListValue(t *testing.T) {

	viperConfig := viper.New()
	for _, remotePath := range etcd_nodes {
		err := viperConfig.AddSecureRemoteProvider(
			remoteType, remotePath, configKey, "")
		if err != nil {
			// invalid remote provider, so don't add to local provider list.
			continue
		}
	}
	viperConfig.SetConfigType(remoteConfType)

	// init read
	err := viperConfig.ReadRemoteConfig()
	if err != nil {
		fmt.Println(err)
		panic(err)
	}
	fmt.Println(viperConfig.AllKeys())
	value := viperConfig.Get("abc.server")
	fmt.Printf("%T\n", value)
	fmt.Println(reflect.TypeOf(value).String())

	if value, ok := value.([]interface{}); ok {
		for index, item := range value {
			fmt.Println(index, item)
		}
	}
}
