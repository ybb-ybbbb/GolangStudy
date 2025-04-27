package main

import (
	"fmt"
)

type Reader interface {
	ReadBook()
}

type Writer interface {
	WriteBook()
}
type Book struct {
}

func (this *Book) ReadBook() {
	fmt.Println("Read a Book")
}

func (this *Book) WriteBook() {
	fmt.Println("Write a Book")
}

// golang中变量在内部实际上是一个包含两个元素的“pair”（对）：一个指向具体类型值的指针，和一个指向类型信息的指针(它指向实际存储数据的内存地址)。
func main() {
	b := &Book{} //b的pair<type:Book,value:&Book{}(指向 b的地址)>
	//type:指向 Book 类型的信息，包括 Book 类型的方法集（ReadBook 和 WriteBook）。
	//value:指向 Book 实例的内存地址。
	var r Reader
	r = b

	r.ReadBook()

	var w Writer
	w = r.(Writer) // 类型断言本质上是对接口变量中存储的具体类型的一个检查和替换过程，因Book实现了Writer的所有方法，所有断言是成功的，w中的pair与b一样（传递性）
	w.WriteBook()
}
