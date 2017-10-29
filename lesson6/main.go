package main

//控制语句
import (
	"errors"
	"fmt"
)

func check(x int) error {
	if x <= 0 {
		return errors.New("x < 0")
	}
	return nil
}

func main() {
	var a, b = 3, 5
	if a > b {
		fmt.Println("a")
	}

	if c := 4; c > 2 {
		fmt.Printf("c = %d\n", c)
	}

	x := 20
	if err := check(x); err == nil {
		x++
		fmt.Println(x)
	} else {
		fmt.Println(err)
	}

	// 循环
	for i := 0; i < 4; i++ {
		fmt.Println(i)
		if i > 3 { //只能是布尔类型
			break
		}
	}
	for {
		break
	}

	//map channel slice array :range

	var data [3]int = [3]int{10, 20, 30}
	for i, s := range data {
		fmt.Println(i, s)
	}

	for i := range data {
		fmt.Println(i, data[i])
	}

	var data2copy [3]string
	data2 := [3]string{"aa", "bb", "cc"}
	for i, s := range data2 {
		fmt.Println(&i, &s) //如果是常量是没有地址的
		data2copy[i] = s
	}
	for _, s := range data2copy {
		fmt.Println(s)
	}

	//goto break continue

	for i := 0; i < 4; i++ {
		fmt.Println(i)
		if i > 2 {
			goto start
		}

	}
start:
	fmt.Println("end")

	testswitch()
}

func testswitch() {
	a, b, c := 1, 2, 30
	x := 30
	switch x {
	case a:
		fmt.Println("x=a")
	case b:
		fmt.Println("x=b")
	case c:
		fmt.Println("x=x")
		fallthrough //只管一个case 继续执行下一个case
	default:
		fmt.Println("not all")
	}
}
