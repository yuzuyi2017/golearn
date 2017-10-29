package main

import "fmt"

func main() {
	fmt.Println(testdefer())

	defer func() {
		if err := recover(); err != nil {
			fmt.Println(err)
		}
	}()

	panic("panic error")
	fmt.Println("err")
}

func testdefer() int {
	a := 100
	defer func() {
		fmt.Println("defer", a)
		a += 100
	}()

	a += 2
	return a
}
