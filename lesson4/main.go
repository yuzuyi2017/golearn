package main

import (
	"fmt"
	"math"
)

type Student struct {
	Name string
	age  int
}

//go的类型
func main() {

	//bool
	//string
	//byte
	//int int8 int16 int32 int64
	//uint uint8 uint16 uint32 uint64
	//complex64 complex128
	//uintptr
	//array

	//struct
	//func
	//interface

	//map //以下只能make
	//slice
	//channel
	s := "abc"
	fmt.Println([]byte(s))
	fmt.Println(s)

	b := false
	fmt.Println(b)

	stu := &Student{}
	stu2 := new(Student)
	fmt.Println(stu.Name, stu.age, *stu2)

	fmt.Println(math.MinInt8, math.MaxInt16, math.MaxUint16)

	m := make(map[string]string)
	m["a"] = "2"
	m["b"] = "3"
	fmt.Println(m)

	//s1 := []int{1, 2, 3, 4, 5}
	//s1 = s1[2:3]
	//s1 = append(s1,12)

	//fmt.Println(s1)

	c := make(chan int, 1)
	c <- 1
	fmt.Println(<-c)

	//自定义类型
	type color int
	var red color
	red = 2
	fmt.Println(red)
}
