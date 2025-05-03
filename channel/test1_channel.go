package main

import (
	"fmt"
	"time"
)

func main() {

	c1 := make(chan int)    //创建一个无缓存的channel
	c2 := make(chan int, 3) //创建一个缓存为3的channel
	c3 := make(chan int)    //创建一个无缓存的channel

	go func() {
		defer fmt.Println("sub goroutine out")
		fmt.Println("sub goroutine runnig")
		c1 <- 666 //向c1写入数据

		for i := 0; i < 3; i++ {
			c2 <- i
		}
		for i := 0; i < 3; i++ {
			c3 <- i
		}
		close(c3)
	}()
	time.Sleep(2 * time.Second) //给时间等待协程执行
	num1 := <-c1                //读数据，<-与c1中间不能隔空格
	fmt.Println("num1= ", num1)
	for i := 0; i < 3; i++ {
		num2 := <-c2
		fmt.Println("num2=", num2)
	}
	for {
		if value, ok := <-c3; ok {
			fmt.Println("num3=", value)
		} else {
			break
		}
	}

	fmt.Println("main goroutine out")
}

// sub goroutine runnig
// num1=  666
// num2= 0
// num2= 1
// num2= 2
// num3= 0
// num3= 1
// num3= 2
// sub goroutine out
// main goroutine out
