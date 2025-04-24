package main

import (
	"fmt"
)

func main() {
	myarray1 := []int{1, 2} //固定数组，传参时为赋值传参

	slice2 := []int{1, 2}       //动态数组slice，传参时为引用传参
	slice3 := make([]int, 3, 5) //开辟一段长度为3总大小为5（cap）的动态数组slice ,如果只传递一个整数则表示长度与cap大小一致

	for i := 0; i < len(myarray1); i++ {
		fmt.Println(myarray1[i])
	}
	fmt.Println("------------")
	for _, value := range slice2 { //for循环的一个用法，range关键字返回两个参数第一个为元素下标，第二个为元素值（不需要可以用_匿名）
		fmt.Println(value)
	}
	fmt.Printf("len3=%d,cap3=%d,slice3=%v\n", len(slice3), cap(slice3), slice3)

}

// 1
// 2
// ------------
// 1
// 2
// len3=3,slice3=[0 0 0]
