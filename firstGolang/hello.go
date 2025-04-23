package main

import (
	"fmt"
	"time"
)

func main() { //go要求函数的 { 必须和函数在同一行
	fmt.Println("hello go")
	time.Sleep(1 * time.Second)
}
