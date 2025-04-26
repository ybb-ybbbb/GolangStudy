package main

import "fmt"

// interface （接口）是一种抽象的类型，定义了一组方法的集合
type Animal interface { // 定义 Animal 为 interface 类型，任何实现了这些方法的类型都可以被视为 Animal 类型
	Getname() string
	Getcolour() string
}
type Pig struct {
	name   string
	colour string
}

// 实现 Animal 接口的 Getname 方法
func (this *Pig) Getname() string {
	return this.name
}

// 实现 Animal 接口的 Getcolour 方法
func (this *Pig) Getcolour() string {
	return this.colour
}

type Dog struct {
	name   string
	colour string
}

func (this *Dog) Getname() string {
	return this.name
}
func (this *Dog) Getcolour() string {
	return this.colour
}

// 定义 Showanimal 函数，参数为 Animal 接口类型
func Showanimal(this Animal) {
	fmt.Printf("name is %s\n", this.Getname()) //通过接口变量this调用这些方法
	fmt.Printf("colour is %s\n", this.Getcolour())
	fmt.Printf("%v\n", this) //此时this其实就是传过来的值
}
func main() {
	mydog := Dog{"zzb", "yello"}
	mypig := Pig{"zzp", "black"}
	Showanimal(&mydog) //因为Dog实现了Animal的方法（通过指针实现），所以Animal接口可以接受并储存Dog实例的指针，并且可以通过接口变量调用这些方法
	Showanimal(&mypig)

}

//在 Go 中，一个类型要被视为实现了某个接口，必须实现接口中定义的所有方法。只有当一个类型实现了接口中的所有方法时，它才能被赋值给该接口类型的变量，并且可以调用接口中的方法。
