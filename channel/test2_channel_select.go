package main

import (
	"fmt"
	"time"
)

// fibonacci 函数通过 channel 生成斐波那契数列
func fibonacci(out chan int, quit chan struct{}) {
	a, b := 0, 1
	for {
		select {
		case out <- a: // 将当前的斐波那契数发送到 out channel
			a, b = b, a+b               // 更新斐波那契数列的值
			time.Sleep(1 * time.Second) // 每秒生成一个数，方便观察
		case <-quit: // 如果 quit channel 收到信号
			fmt.Println("Fibonacci generator is exiting...")
			return // 退出生成器
		}
	}
}

func main() {
	out := make(chan int)       // 用于输出斐波那契数列的 channel
	quit := make(chan struct{}) // 用于控制退出的 channel

	go fibonacci(out, quit) // 启动一个 goroutine 来运行 fibonacci 函数

	// 从 out channel 中接收斐波那契数列的值
	for i := 0; i < 10; i++ { // 限制输出 10 个斐波那契数
		num := <-out
		fmt.Println(num)
	}

	// 发送退出信号到 quit channel
	close(quit) // 关闭 quit channel，通知 fibonacci 函数退出
}

//当主函数关闭 quit channel 后，fibonacci 函数中的 select 语句会检测到 quit channel 有数据可读（实际上是 channel 被关闭
// 0
// 1
// 1
// 2
// 3
// 5
// 8
// 13
// 21
// 34
