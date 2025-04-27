package main

import (
	"fmt"
	"reflect"
)

type User struct {
	Name string
	Id   int
}

// 指针接收者的方法
func (user *User) Usersleep1() {
	fmt.Println(user.Name, "is sleeping1")
}

// 值接收者的方法
func (user User) Usersleep2() {
	fmt.Println(user.Name, "is sleeping2")
}

// 方法集规则：
// 值类型 User 的方法集：只有 Usersleep2。
// 指针类型 *User 的方法集：包含 Usersleep1 和 Usersleep2（Go 会自动将值方法提升为指针方法）。

func Getinfo(a interface{}) {
	val := reflect.ValueOf(a)
	typ := val.Type()

	// 输出类型名（处理指针类型）
	if typ.Kind() == reflect.Ptr {
		fmt.Println("type is:", typ.Elem().Name()) // Elem()实际上是一个解引用的过程
	} else {
		fmt.Println("type is:", typ.Name())
	}

	// 获取字段（自动解引用指针）
	elem := val
	if typ.Kind() == reflect.Ptr {
		elem = val.Elem() // 获取指针指向的值
	}

	// 遍历字段
	for i := 0; i < elem.NumField(); i++ {
		field := elem.Type().Field(i)
		value := elem.Field(i).Interface()
		fmt.Printf("%s: %v = %v\n", field.Name, field.Type, value)
	}

	// 获取方法（直接使用原始类型，支持指针方法）
	methodType := val.Type()
	for i := 0; i < methodType.NumMethod(); i++ {
		m := methodType.Method(i)
		fmt.Printf("%s: %v\n", m.Name, m.Type)
	}
}

func main() {
	user1 := User{"ybb", 1}
	Getinfo(&user1) // 传递指针以获取所有方法
}

// type is: User
// Name: string = ybb
// Id: int = 1
// Usersleep1: func(*main.User)
// Usersleep2: func(*main.User)
