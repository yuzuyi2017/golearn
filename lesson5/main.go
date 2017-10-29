package main

import (
	"fmt"
)

//运算符
//关键字 break, default, func, interface, select, type, fallthrough
func main() {
	/*
		& += &= != () >>= ... . :
		1. * / % >> << &
		2. + - | ^
		3.&&
		4. ||
		5.
	*/

	const (
		read   byte = 1 << iota
		write
		exec
		freeze
	)
	a := read | write | freeze
	b := read | write | exec
	c := a &^ b

	fmt.Printf("%04b & %04b = %04b \n", a, b, c)
}
