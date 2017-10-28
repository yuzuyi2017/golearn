package main

import (
	"fmt"
)

const x, y int = 1, 2 //定义可以不使用
const s = "hello"
const c = "aa"

func main() {
	var d int
	const y = 123
	fmt.Println(y, s, c, d)
	const (
		x1, y1      = 1, 2
		b      byte = byte(x)
	)

	const (
		x2 = iota*10 + 1 //iota 在常量堆里的位置
		y2
		z2
		e  = 100
		f  = iota
		h
	)
	fmt.Print(x2, y2, z2, e, f, h)

	const k = 2 //常量没有内存地址
	//fmt.Print(&k)
}
