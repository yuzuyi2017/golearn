package main

import (
	"fmt"
	"unsafe"
)

//结构体
type node struct {
	id   int
	next *node
}

func main() {

	var a struct{}

	fmt.Println(unsafe.Sizeof(a))
}
