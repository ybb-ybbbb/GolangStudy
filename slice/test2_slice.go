package main

import (
	"fmt"
)

func main() {
	var slice4 []int   //声明一个空的动态数组
	if slice4 == nil { //nil为一个关键字表示空
		for i := 0; i < 10; i++ {
			slice4 = append(slice4, i)
			fmt.Printf("len=%d, cap=%d, slice4=%v\n", len(slice4), cap(slice4), slice4)
		}
	} else { //golang规定else必须和两个括号在一行

	}

}

// len=1, cap=1, slice4=[0]
// len=2, cap=2, slice4=[0 1]
// len=3, cap=4, slice4=[0 1 2]
// len=4, cap=4, slice4=[0 1 2 3]
// len=5, cap=8, slice4=[0 1 2 3 4]
// len=6, cap=8, slice4=[0 1 2 3 4 5]
// len=7, cap=8, slice4=[0 1 2 3 4 5 6]
// len=8, cap=8, slice4=[0 1 2 3 4 5 6 7]
// len=9, cap=16, slice4=[0 1 2 3 4 5 6 7 8]
// len=10, cap=16, slice4=[0 1 2 3 4 5 6 7 8 9]
//当 slice 的长度（len）超过其容量（cap）时，Go 会自动分配一个新的底层数组，并将旧数组的内容复制到新数组中。
// 如果当前容量小于 1024，新容量通常是旧容量的两倍。
// 如果当前容量大于或等于 1024，新容量会增加一个固定的大小（通常是旧容量的一半）
