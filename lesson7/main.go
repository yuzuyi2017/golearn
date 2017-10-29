package main

import "fmt"

func main() {
	sayHello()
	a := sayHello
	b := sayhello2
	a = b
	a()
}

func sayHello() {
	fmt.Println("hello")

	a := func() { fmt.Println("222") }
	a()
}

func sayhello2() {
	fmt.Println("hello 2")
}
