package main

import (
	"fmt"
)

func main() {
	mymap1 := map[int]string{ //map定义与cpp类似，传参为引用传参
		0: "a",
		1: "b",
		2: "c",
	}
	mymap2 := make(map[int]string) //map有多种定义方式，这里展示两种，第一种带初始化
	mymap2[12] = "abc"

	delete(mymap1, 0) //增查改类比slice,删除用delect（）

	fmt.Println(mymap1)
	fmt.Println(mymap2)

}

// map[1:b 2:c]
// map[12:abc]
