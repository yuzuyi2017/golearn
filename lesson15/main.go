package main

//方法
import "fmt"

type Stu struct {
	name string
	age  int
}

type N int

func main() {
	fmt.Println(func1(1))
	var s Stu
	fmt.Println(s.setAge(11))

	var n N
	n = 65
	fmt.Println(n.toString())
}

func func1(a int) string {
	return "aa"
}

func (s *Stu) setAge(a int) int {
	s.age = a
	return s.age
}

func (n N) toString() string {
	return fmt.Sprintf("%c", n)
}
