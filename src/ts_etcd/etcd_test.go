package ts_etcd

import (
	"fmt"
	"io/ioutil"
	"reflect"
	"testing"

	"github.com/spf13/viper"
	_ "github.com/spf13/viper/remote" // to enable the remote features
	"github.com/stretchr/testify/assert"
	"github.com/xordataexchange/crypt/backend/etcd"
)

var etcd_nodes = []string{
	"https://10.18.96.172:2379",
	"https://10.18.96.171:2379",
	"https://10.18.96.174:2379"}

const configKey = "/etcd_test"
const remoteConfType = "toml"
const remoteType = "etcd"
const fileName = "test.toml"

// 测试将带有列表的toml写到etcd中，使用v2
func TestWriteSliceTOML(t *testing.T) {
	client, err := etcd.New(etcd_nodes)
	assert.NoError(t, err)
	configBytes, err := ioutil.ReadFile(fileName)
	assert.NoError(t, err)
	fmt.Println(configBytes)
	err = client.Set(configKey, configBytes)
	assert.NoError(t, err)
	storeContents, err := client.Get(configKey)
	assert.NoError(t, err)
	assert.Equal(t, storeContents, configBytes)
}

// 测试从etcd中读出带有列表的toml配置，使用v2
func TestGetSliceTOML(t *testing.T) {
	client, err := etcd.New(etcd_nodes)
	assert.NoError(t, err)
	storeContents, err := client.Get(configKey)
	assert.NoError(t, err)
	fmt.Println(storeContents)
}

// 测试使用viper 从etcd中读出带有列表的toml配置，使用v2
func TestViperReadTOMLSliceConfig(t *testing.T) {

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
	assert.NoError(t, err)
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
