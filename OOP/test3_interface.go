package main

import "fmt"

func Whatim(arg interface{}) { //inerface{}是空接口，由于所有数据类型都实现了它所有的方法，所有它可以调用所有数据类型的方法，或者说它可以是所有数据类型
	fmt.Printf("i am %T, %v\n", arg, arg)
}

type Pig struct {
	name   string
	colour string
}

func main() {
	a := 10
	b := "ybb"
	c := make([]int, 3)
	mypig := Pig{"zzp", "black"}
	Whatim(a)     // i am int, 10
	Whatim(b)     // i am string, ybb
	Whatim(c)     // i am []int, [0 0 0]
	Whatim(mypig) //i am main.Pig, {zzp black}
	var d interface{}
	d = "ybb"
	if str, ok := d.(string); ok { //interface类型断言
		fmt.Println("d is string,value is ", str)
	} else {
		fmt.Println("d isnt string")
	}
	var data interface{} = 42

	switch v := data.(type) {
	case int:
		fmt.Println("data is an int:", v)
	case string:
		fmt.Println("data is a string:", v)
	default:
		fmt.Println("data is of unknown type")
	}
}

// i am int, 10
// i am string, ybb
// i am []int, [0 0 0]
// d is string,value is  ybb
// data is an int: 42
