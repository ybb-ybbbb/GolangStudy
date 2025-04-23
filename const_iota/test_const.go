package main

import (
	"fmt"
)

const (
	a = iota       //iota是一个关键字（只能配合const使用），逐行递增1，iota=a=0
	b              //iota=b=1
	c = iota*2 + 1 //iota=2,c=2*2+1
	d              //iota=3,d=3*2+1
)

func main() {
	fmt.Println("a=", a)
	fmt.Println("b=", b)
	fmt.Println("c=", c)
	fmt.Println("d=", d)
}
