package main

import (
	"fmt"
	"reflect"
)

// fmt 包是标准库的一部分，它在编译时被特殊处理，允许它访问任何字段，无论它们是否导出（无视大小写）
type User struct { //这里User的属性开头字母要大写！因为reflect包要调用它
	Name string
	Id   int
}

func (user *User) Usersleep1() {
	fmt.Println(user.Name, "is sleeping1")
}
func (user User) Usersleep2() { //!!!!!因为我传递过去的是User实例而不是它的地址，所以通过reflect.Type 获取a实现的方法不会识别到指针方法（Usersleep1）
	fmt.Println(user.Name, "is sleeping2")
}
func Getinfo(a interface{}) {
	//获取aType
	aType := reflect.TypeOf(a) // 返回一个reflect.Type 接口
	// fmt.Printf("type is %T\n", aType)     //type is *reflect.rtype
	// fmt.Printf("type is %v\n", aType)     //type is main.User
	// fmt.Println("type is ", aType)        //type is  main.User
	fmt.Println("type is ", aType.Name()) //type is  User    Name() 方法专门用于返回类型的名称，不包括包路径

	//获取aValue
	aValue := reflect.ValueOf(a)
	fmt.Println("value is ", aValue) //ValueOf() 函数用于获取变量的反射值reflect.Value 它是一个接口，定义了一组方法，用于操作和检查反射值

	//通过reflect.Type 获取a中的字段
	// NumField()：如果类型是结构体，返回结构体的字段数量。
	for i := 0; i < aType.NumField(); i++ {
		field := aType.Field(i)              //reflect.Type 接口的 Field() 方法用于获取结构体字段的信息 包括字段的名称、类型、标签（tag）等信息。
		value := aValue.Field(i).Interface() //reflect.Value 接口的 Field() 方法用于获取结构体的字段的反射值（reflect.Value）
		// Interface() 方法会返回一个空接口（interface{}），其中包含了 reflect.Value 所代表的原始值的副本。
		fmt.Printf("%s: %v=%v\n", field.Name, field.Type, value)
	}

	//通过reflect.Type 获取a实现的方法
	for i := 0; i < aType.NumMethod(); i++ {
		m := aType.Method(i)
		fmt.Printf("%s: %v\n", m.Name, m.Type)
	}
}

// 当你直接打印一个 reflect.Type 对象时，例如使用 fmt.Println(a)，Go 的 fmt 包会调用该类型的 String() 方法来获取其字符串表示。
// 对于 reflect.Type 接口，其底层实现会返回一个字符串，该字符串通常包含类型所在的包路径和类型名称，格式一般为 "pkgpath.TypeName"
func main() {
	user1 := User{"ybb", 1}
	Getinfo(user1)
}

// type is  User
// value is  {ybb 1}
// Name: string=ybb
// Id: int=1
// Usersleep2: func(main.User)
