package main

import "fmt"

var (
	a int
	b int = 100
	c     = 200
) //只有使用var才能声明全局变量

func main() {
	d := 100
	fmt.Println("a=", a, ",b=", b, ",c=", c, ",d=", d)
}

// ybb@ybb:~/go/src/GolangStudy/var$ go run test_var.go
// a= 0 ,b= 100 ,c= 200 ,d= 100
