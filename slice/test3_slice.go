package main

import (
	"fmt"
)

func main() {
	slice := []int{1, 2, 3, 4}
	fmt.Printf("slice=%v\n", slice)
	s1 := slice[0:2] //左闭右开区间的引用截取（指向同一片地址）
	fmt.Printf("s1=%v\n", s1)
	s2 := make([]int, 10)
	copy(s2, slice) //将slice1的值拷贝到s2（s2对slice进行复制）不在同一片地址
	fmt.Printf("s2=%v\n", s2)
}

// slice=[1 2 3 4]
// s1=[1 2]
// s2=[1 2 3 4 0 0 0 0 0 0]
