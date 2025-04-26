package main

import "fmt"

// golang中首字母大小写决定了该类或方法的公私，若大写则表示除了该类所在包其他包也可以访问
type Book struct {
	name   string
	number int
}
type MagicBook struct { //子类继承父类Book的属性，还可以添加自己的属性
	Book
	magickind string
}

func (this *Book) Setname(name string) { //这里加*跟cpp一样指传地址，this是一个指针！
	this.name = name
}
func (this *Book) Setnumber(number int) {
	this.number = number
}
func (this *Book) Showbook() {
	fmt.Printf("%v\n", *this) //%v指打印所有信息，*this解引用
}

func (this *MagicBook) Showbook() {
	fmt.Printf("Magic is%v\n", *this) //子类可以对父类方法重写
}

func main() {
	var mybook Book
	mybook.Setname("ybb")
	mybook.Setnumber(1)
	mybook.Showbook()
	fmt.Println(mybook.name) //golang的公私只对包而言，尽管Book的属性name是私有的但是在它所在包下还是可以随意调用，但其他包不行

	mybookmagic := MagicBook{mybook, "fuck"}
	mybookmagic.Showbook()
}
