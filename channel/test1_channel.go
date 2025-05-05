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
		close(c3) //关闭channel后，再往该channel发送数据会报错，但是可以从该channel取数据
	}()
	time.Sleep(2 * time.Second) //给时间等待协程执行
	num1 := <-c1                //读数据，<-与c1中间不能隔空格
	fmt.Println("num1= ", num1)
	for i := 0; i < 3; i++ {
		num2 := <-c2
		fmt.Println("num2=", num2)
	}

	/*
		for {
			if value, ok := <-c3; ok {	//c3如果存在ok就会一直返回true
				fmt.Println("num3=", value)
			} else {
				break
			}
		}
	*/

	for data := range c3 { //range可以读取channel数据，若有数据就返回，若无数据则阻塞等待
		fmt.Println("num3=", data)
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
