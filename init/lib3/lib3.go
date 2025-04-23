package lib3

import "fmt"

func init() {
	fmt.Println("lib3.init()......")
}
func Lib3_test() { //lib1包中供外部调用的api，注意函数名首字母大写才可供外界调用
	fmt.Println("Lib3_test()......")
}
