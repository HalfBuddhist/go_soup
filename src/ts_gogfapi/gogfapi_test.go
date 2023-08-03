package gogfapi_test

import (
	"fmt"
	"testing"
	"unsafe"

	"go_soup/src/juicedata/gogfapi/gfapi"
)

func TestWriteFile(t *testing.T) {
	vol := &gfapi.Volume{}
	if err := vol.Init("devvolume", "uat-master1"); err != nil {
		// handle error
		fmt.Println(err)
		return
	}

	if err := vol.Mount(); err != nil {
		// handle error
		fmt.Println(err)
		return
	}
	defer vol.Unmount()

	f, err := vol.Create("testdir/tmp/testfile")
	if err != nil {
		// handle error
		fmt.Println(err)
		return
	}
	defer f.Close()

	if _, err := f.Write([]byte("hello")); err != nil {
		// handle error
		fmt.Println(err)
		return
	}
}

// 统计路径的信息，仍然需要 mount
// Bsize X Bavail = 可用容量
// Bsize X Blocks = 占用容量
func TestStatvfs(t *testing.T) {
	vol := &gfapi.Volume{}
	if err := vol.Init("devvolume", "uat-master1"); err != nil {
		// handle error
		fmt.Println(err)
		return
	}

	if err := vol.Mount(); err != nil {
		// handle error
		fmt.Println(err)
		return
	}
	defer vol.Unmount()

	var res gfapi.Statvfs_t
	err := vol.Statvfs("tenant-1", &res)
	if err != nil {
		// handle error
		fmt.Println(err)
		return
	}
	fmt.Printf("%+v\n", res)
}

func TestGetxattr(t *testing.T) {
	vol := &gfapi.Volume{}
	if err := vol.Init("devvolume", "uat-master1"); err != nil {
		// handle error
		fmt.Println(err)
		return
	}

	if err := vol.Mount(); err != nil {
		// handle error
		fmt.Println(err)
		return
	}
	defer vol.Unmount()

	var value [512]byte
	res, err := vol.Getxattr("tenant-1", "quota", value[:])
	if err != nil {
		// handle error
		fmt.Println(err)
		return
	}
	fmt.Printf("%+v\n", res)
	fmt.Printf("%s\n", value)
}

func TestListXattr(t *testing.T) {
	vol := &gfapi.Volume{}
	if err := vol.Init("devvolume", "uat-master1"); err != nil {
		// handle error
		fmt.Println(err)
		return
	}

	if err := vol.Mount(); err != nil {
		// handle error
		fmt.Println(err)
		return
	}
	defer vol.Unmount()

	var value [512]byte
	res, err := vol.ListXattr("testdir/tmp/testfile", value[:], 512)
	if err != nil {
		// handle error
		fmt.Println(err)
		return
	}
	fmt.Printf("%+v\n", res)
	fmt.Printf("%s\n", *(*string)(unsafe.Pointer(&value)))
}
