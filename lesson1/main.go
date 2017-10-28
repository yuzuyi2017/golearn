package main

//变量
import (
	"fmt"
	"reflect"
)

var x = 100

func main() {
	/*fmt.Println("hello world")
	var x int
	var name  = "12445"
	var y  float32
	var z string
	fmt.Println(x, y , z , name)*/

	/*var a, b int
	a = 100
	b = 200
	fmt.Println(a, b)*/

	var a, s = 1000, "sadsd"
	fmt.Println(a)
	fmt.Println(s)
	fmt.Println(reflect.TypeOf(a))
	fmt.Println(reflect.TypeOf(s))

	r := 1
	fmt.Println(r)

	fmt.Println(&x)

	var x, y = 1, 2
	fmt.Println(x, y)
	x, y = y, x
	fmt.Println(x, y)
}
