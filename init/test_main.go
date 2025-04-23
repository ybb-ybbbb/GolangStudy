package main

//import导包
import (
	. "GolangStudy/init/lib1"  //.空格接包表示将包的方法全部导入main包，主函数中可以直接使用（当心重名！不推荐）
	aa "GolangStudy/init/lib2" //给包起个别名
	_ "GolangStudy/init/lib3"  //_空格接包表示不调用包的任何接口但是会执行包的init()函数
)

func main() {
	Lib1_test()
	aa.Lib2_test()
}

// ybb@ybb:~/go/src/GolangStudy/init$ go run test_main.go
// lib1.init()......
// lib2.init()......
// lib3.init()......
// Lib1_test()......
// Lib2_test()......
