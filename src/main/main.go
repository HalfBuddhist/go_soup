// 匹配 UUID
package main

import (
	"fmt"
	"reflect"

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

func main() {

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
