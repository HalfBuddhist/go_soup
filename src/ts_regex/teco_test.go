package ts_regex

import (
	"fmt"
	"regexp"
	"testing"
)

func TestTecoNsPattern(t *testing.T) {
	ResourcePoolK8sNsPrefix := "kubecube"
	pattern := regexp.MustCompile(
		fmt.Sprintf(`^%s-(tenant|workspace)-(\d+)(-(dedicated|shared|reserved))?$`,
			ResourcePoolK8sNsPrefix))
	res := pattern.Match([]byte("kubecube-workspace-2158"))
	fmt.Println(res)
	res = pattern.Match([]byte("kubecube-workspace-2158-dedicated"))
	fmt.Println(res)
	matches := pattern.FindStringSubmatch("kubecube-workspace-2158-dedicated")
	fmt.Printf("%#v\n", matches)
}

func TestDockerImageAddressPattern(t *testing.T) {
	pattern := regexp.MustCompile(`^((?P<host>^[a-zA-Z0-9][-a-zA-Z0-9]*(\.[a-zA-Z0-9][-a-zA-Z0-9]*)+)/)?(?P<project>[a-zA-Z0-9]{1}[a-zA-Z0-9-_]*)/(?P<name>[a-zA-Z0-9]{1}[a-zA-Z0-9-_]*)(:(?P<version>[a-zA-Z0-9-_\.]*))?$`)
	res := pattern.Match([]byte("nvidia/l4t-tensorflow:r32.5.0-tf2.3-py3"))
	fmt.Println(res)
	matches := pattern.FindStringSubmatch("nvidia/l4t-tensorflow:r32.5.0-tf2.3-py3")
	fmt.Printf("%#v\n", matches)
}

// 镜像地址格式不正确: 要求仅支持小写字母、数字和3个特殊符号“-_.”，开头与结尾必须数字或字母，且不能连续的输入不同的特殊字符。"
func TestImageAddressPattern(t *testing.T) {
	// regex := `^[a-z0-9](?:[a-z0-9]|[-_.](?![^a-z0-9][-_.]))*[a-z0-9]$` // qwen-1
	// regex := `^[a-z0-9](?:[a-z0-9]|[-_.](?![^a-z0-9][-_.]))*[a-z0-9]$` // qwen-2
	// regex := `^[a-z0-9](?:[a-z0-9]*[-_.][a-z0-9]*)*[a-z0-9]$` // qwen-3
	// regex := `^[a-z0-9](?:[a-z0-9]|([-_.])\1*[a-z0-9])*[a-z0-9]$` // qwen-4
	// regex := `^[a-z0-9](?:[a-z0-9]|([-_.])\\1*[a-z0-9])*[a-z0-9]$` // qwen-5, error still
	// regex := `^(?=[a-z0-9])(?:(?![._-][._-])[a-z0-9._-])*[a-z0-9]$` // qwen2.5-1
	// regex := `^[a-z0-9]([a-z0-9]|([._-](?!\1)))*[a-z0-9]$` // qwen2.5-2
	// regex := `^[a-z0-9]([a-z0-9]|([._-])(?!\2))*[a-z0-9]$` // qwen2.5-3
	// regex := `^[a-z0-9]([a-z0-9]|([._-]))*[a-z0-9]$` // qwen2.5-4, no error, but not correct
	regex := `^[a-z0-9]([a-z0-9]|([._-]))*[a-z0-9]$` // qwen2.5-5, no error, but not correct still
	// regex := `^[a-z0-9](?:[a-z0-9]|([._-])(?!\1))*[a-z0-9]$` // gemini25-1 and sonnet37
	// regex := `^[a-z0-9](?:[a-z0-9]+|[-_.][a-z0-9]+)*$` // gemini25-2
	// regex := `^[a-z0-9](?:[a-z0-9]+|(?:(?:-+|\.+|_+)[a-z0-9]+))*$` // gemini25-3
	// regex := `^[a-z0-9](?:[a-z0-9]|([_.-])(?:\1|[a-z0-9]))*[a-z0-9]$` // deepseek r1
	// regex := `^(?:[a-z0-9](?:(?:[a-z0-9]|-(?![_.])|_(?![.-])|\.(?![_-]))*[a-z0-9]|[a-z0-9])$` // deepseek r1 2
	// regex := `^[a-z0-9](?:[a-z0-9]|[-]+[a-z0-9]|[_]+[a-z0-9]|[.]+[a-z0-9])*$` // deepseek r1 3
	// regex := `^[a-z0-9](?:[a-z0-9]|([_.-])(?:\1|[a-z0-9]))*[a-z0-9]$` // sonnet37-2
	// regex := `^[a-z0-9](?:[a-z0-9]|[-]+[a-z0-9]|[_]+[a-z0-9]|[.]+[a-z0-9])*$` // sonnet37-3
	re := regexp.MustCompile(regex)
	testCases := []string{
		"abc",      // 合法
		"ab-c",     // 合法
		"ab--c",    // 合法（允许连续相同符号）
		"a..b",     // 合法
		"a.b-c",    // 合法
		"a-.",      // 非法（结尾是特殊符号）
		"a-b-.c",   // 非法（连续不同符号 "-."）
		"a_b_c",    // 合法
		"a_b..c",   // 合法
		"a_b-c.",   // 非法（结尾是特殊符号）
		"123-456",  // 合法
		"a-b_c.",   // 非法（结尾是特殊符号）
		"-.a",      // 非法（开头是特殊符号）
		"a-b_c-d",  // 合法
		"a..b-c.d", // 合法
	}

	for _, test := range testCases {
		fmt.Printf("输入: %-10s => 匹配结果: %v\n", test, re.MatchString(test))
	}
}

func TestLinuxPattern(t *testing.T) {
	// 纯正则表达式解决方案
	// 路径段的合法模式：
	// 1. 以字母数字下划线横线开头的任意组合: [a-zA-Z0-9_-][a-zA-Z0-9._-]*
	// 2. 以点开头但必须包含非点字符: \.[a-zA-Z0-9._-]*[a-zA-Z0-9_-][a-zA-Z0-9._-]*
	pattern := regexp.MustCompile(`^/$|^(/([a-zA-Z0-9_-][a-zA-Z0-9._-]*|\.[a-zA-Z0-9._-]*[a-zA-Z0-9_-][a-zA-Z0-9._-]*))+$`)

	testCases := []struct {
		path     string
		expected bool
	}{
		{"/", true},                  // 根目录
		{"/usr", true},               // 单级目录
		{"/usr/local", true},         // 多级目录
		{"/usr/local/bin", true},     // 多级目录
		{"/etc/nginx/conf.d", true},  // 包含点号的目录
		{"/var/log/app.log", true},   // 文件路径
		{"/home/user/.config", true}, // 隐藏目录
		{"/a/.xxx/c", true},          // 点号与其他字符组合
		{"/a/..xxx/c", true},         // 双点号与其他字符组合
		{"/a/xxx./c", true},          // 以点号结尾的目录名
		{"/a/xxx../c", true},         // 以双点号结尾的目录名
		{"/a/....xxx/c", true},       // 多个点号与字符组合
		{"/a/.x.x./c", true},         // 点号与字符混合
		{"usr/local", false},         // 不以/开头
		{"//usr", false},             // 双斜杠
		{"/usr//local", false},       // 路径中有双斜杠
		{"/a/./c", false},            // 路径段只有单点
		{"/a/../c", false},           // 路径段只有双点
		{"/a/.../c", false},          // 路径段只有三个点
		{"/a/..../c", false},         // 路径段只有四个点
		{"/usr/./local", false},      // 路径段只有单点
		{"/usr/../local", false},     // 路径段只有双点
		{"/usr/ /local", false},      // 包含空格
		{"/usr/@local", false},       // 包含非法字符@
		{"/.", false},                // 路径段只有单点
		{"/..", false},               // 路径段只有双点
		{"/...", false},              // 路径段只有三个点
		{"/....", false},             // 路径段只有四个点
	}

	for _, tc := range testCases {
		t.Run(tc.path, func(t *testing.T) {
			result := pattern.MatchString(tc.path)
			if result != tc.expected {
				t.Errorf("path %q: expected %v, got %v", tc.path, tc.expected, result)
			}
			fmt.Printf("路径: %-15s => 匹配结果: %v (期望: %v)\n", tc.path, result, tc.expected)
		})
	}
}
