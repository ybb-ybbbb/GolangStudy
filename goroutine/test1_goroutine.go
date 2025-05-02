package main

import (
	"fmt"
	"time"
)

func Atask() {
	i := 0
	for {
		i++
		fmt.Println("this is A goroutine: ", i)
		time.Sleep(1 * time.Second)
	}
}
func main() {
	go Atask() //创建一个goroutine，去执行Atask

	i := 0
	for {
		i++
		fmt.Println("this is main goroutine: ", i)
		time.Sleep(2 * time.Second) //主go程打印速度慢一倍
	}
	// this is main goroutine:  1
	// this is A goroutine:  1
	// this is A goroutine:  2
	// this is main goroutine:  2
	// this is A goroutine:  3
	// this is A goroutine:  4
	// this is main goroutine:  3
	// this is A goroutine:  5
	// this is A goroutine:  6
	// this is main goroutine:  4
	// this is A goroutine:  7
	// this is A goroutine:  8
	//主go程和从go程并发执行，若主进程结束则不会再执行从进程
	//在你的代码中，主 Goroutine 的打印任务非常简单，执行速度很快。它在启动从 Goroutine 后，会立即继续执行后续的打印操作。
	//从 Goroutine 也需要一定的时间来被调度器调度到 CPU 上执行。虽然调度器会尽快调度它，但主 Goroutine 可能已经完成了多次打印操作。
}
