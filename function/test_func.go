package main

import "fmt"

func foo1(a int, b string) (r1 int, r2 int) {
	fmt.Println("a=", a)
	fmt.Println("b=", b)
	r1, r2 = 100, 200
	return
}

func main() {
	r1, r2 := foo1(10, "abc")
	fmt.Println("r1=", r1)
	fmt.Println("r2=", r2)
}

//ybb@ybb:~/go/src/GolangStudy/function$ go run test_func.go
// a= 10
// b= abc
// r1= 100
// r2= 200
