package main

import "fmt"

func returnfunc() int {
	fmt.Println("return")
	return 0
}
func testfunc() int {
	defer fmt.Println("a") //defer是个关键字后面接正常句子，defer后的语句在函数}后执行
	defer fmt.Println("b") //defer为压栈存入，取出时为先进后出
	return returnfunc()
}
func main() {
	testfunc()
}

// ybb@ybb:~/go/src/GolangStudy/defer$ go run test_defer.go
// return
// b
// a
