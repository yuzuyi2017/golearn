package main

//命名：驼峰，不是匈牙利（下划线）
import (
	"fmt"
	"crypto/x509/pkix"
)

type Student struct {
	//首字母大写可以导出public，小写不可到处private
	Name string
	aget int
}

func main() {
	fmt.Println(1)

	var student = Student{"tome", 10}
	fmt.Println(student)
}
