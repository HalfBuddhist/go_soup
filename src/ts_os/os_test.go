package ts_os

import (
	"fmt"
	"os"
	"testing"
)

// 测试取得 int 与 uint 的最大最小值
func TestIntMinMax(t *testing.T) {
	// os.ReadDir()
}

// 路径不存在时报错如下：
// *fs.PathError, stat /home/liuqw/workspaceccc: no such file or directory
func TestStat(t *testing.T) {
	fInfo, err := os.Stat("/home/liuqw/workspace/sardine.tar.gz")
	fmt.Printf("%T, %v, %s, %v\n", err, err, err, fInfo.Name())
}

// src为文件时
// dst为完整文件路径是可以的。
// dst文件夹路径是不可以的，会提示文件已经存在。
// 此时dst只能保证目标文件夹存在，且路径末端是新文件名。
//
// src为文件夹时， 此时只能保证上层目标文件夹存在，且路径末端是新文件夹名。
func TestRename(t *testing.T) {
	err := os.Rename("/home/liuqw/workspace/go_soup/src/ts_randd", "/home/liuqw/workspace/go_soup/src/ts_os/xxx")
	fmt.Println(err)
}
