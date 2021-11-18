// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

package main

import (
	"fmt"
	"unsafe"
)

func main() {
	//!+main
	fmt.Printf("%d %s\n", "hello", 42)
	fmt.Println(unsafe.Sizeof(float64(0)))
}
