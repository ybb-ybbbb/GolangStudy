package main

import (
	"fmt"
	"runtime"
	"time"
)

func main() {
	go func() {
		defer fmt.Println("defer A")
		func() {
			defer fmt.Println("defer B")
			runtime.Goexit()
			//defer B
			//defer A,退出go程
			fmt.Println("B")
		}()
		fmt.Println("A")
	}()
	for {
		time.Sleep(1 * time.Second)
	} //要加死循环不然主程序直接退出了不会执行goroutine

}

// B
// defer B
// A
// defer A
