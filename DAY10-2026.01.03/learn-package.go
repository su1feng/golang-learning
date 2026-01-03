package main

import (
	"fmt"
)

var x int8 = 10

const pi = 3.14

func init() {
	fmt.Println("x:", x)
	fmt.Println("pi:", pi)
	sayHi()
}

func sayHi() {
	fmt.Println("Hello World!")
}

func main() {

	//大小写区分public or private
	//引入包的格式 import importname "path/to/package"    importname为引入的包名
	//匿名引入 importname为_   主要目的是加载这个包

	//init初始化函数
	/*
		这种特殊的函数不接收任何参数也没有任何返回值，我们也不能在代码中主动调用它。当程序启动的时候，init函数会按照它们声明的顺序自动执行。
		一个包的初始化过程是按照代码中引入的顺序来进行的，所有在该包中声明的init函数都将被串行调用并且仅调用执行一次。
		每一个包初始化的时候都是先执行依赖的包中声明的init函数再执行当前包中声明的init函数。确保在程序的main函数开始执行时所有的依赖包都已初始化完成。

		func init(){}

	*/

	fmt.Println("你好，世界！")


	/*
		go module使用场景:
		1.任何go项目都需要的
		2.引入第三方包
		3.管理依赖包版本
	*/

}
