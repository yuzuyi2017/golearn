package main

import (
	"fmt"
	"reflect"
	"unsafe"
)

func main() {

	s := "abcdefg"

	s1 := s[1:3]

	s2 := s[:4]

	fmt.Printf("%#v \n", (*reflect.StringHeader)(unsafe.Pointer(&s)))
	fmt.Printf("%#v \n", (*reflect.StringHeader)(unsafe.Pointer(&s1)))
	fmt.Printf("%#v \n", (*reflect.StringHeader)(unsafe.Pointer(&s2)))

	for i, c := range s2 {
		fmt.Printf("%d %c ", i, c)
	}

	s3 := "abc"
	fmt.Printf("%c", s3[1])
}
