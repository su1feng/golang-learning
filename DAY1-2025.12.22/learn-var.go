package main

import (
	"fmt"
)

// 函数左大括号必须跟函数声明同一行
func foo() (int, string) {
	return 23, "su1feng"
}

func main() {
	/*
		变量声明格式: var 变量名 变量类型
		支持批量变量声明
		声明时变量自动初始化默认值 0 “” false nil
		支持类型推导(可以省略变量类型)
		支持一次性初始化多个变量
		短变量(函数内部)声明 :=
		匿名变量不占空间,不会分配内存,不存在重复声明


		函数外的每个语句必须关键字开头
		:=不能用在函数外
		_多用于占位,表示忽略值

		常量定义时必须赋值,同时声明多个变量时若忽略值则和上一行值相同

		iota 只能在常量的表达式中使用 是一个常量计数器 enum
		可以用占位跳过某些值
		可以中间插队
		定义数量级
		多个iota定义在同一行
	*/

	//变量声明
	var name string = "su1feng"

	fmt.Println("name =", name)

	//批量变量声明
	var (
		age      int = 18
		location string
		tel      string
	)
	fmt.Println("age = ", age, "location = ", location, "tel = ", tel)

	//类型推导+一次性初始化多个变量
	var city, province = "Suzhou", 2

	fmt.Println("city = ", city, "\n", "province = ", province)

	//短变量(函数内部)声明
	m := 001
	n := 23
	fmt.Println(m, n)

	//匿名变量 多用于占位  多重赋值想忽略某个值

	x, _ := foo()
	_, y := foo()

	fmt.Println("x = ", x, "\n", "y = ", y)

	//常量 定义时必须赋值
	const (
		q  = 57
		pi = 3.1415
		//省略了值则表示和上一行相同
		p
	)

	fmt.Println("q = ", q, "pi = ", pi, "p = ", p)

	//iota  常量计数器  只能在常量的表达式中使用
	const (
		n1 = iota
		n2
		n3
		_ //跳过3
		n4
		n5 = 100
		n6 = iota
	)

	const n7 = iota

	fmt.Println("n1 = ", n1)
	fmt.Println("n2 = ", n2)
	fmt.Println("n3 = ", n3)
	fmt.Println("n4 = ", n4)
	fmt.Println("n5 = ", n5)
	fmt.Println("n6 = ", n6)
	fmt.Println("n7 = ", n7)

	const (
		_  = iota
		KB = 1 << (10 * iota)
		MB = 1 << (10 * iota)
		GB
	)

	fmt.Println("KB = ", KB, "MB = ", MB, "GB = ", GB)

	const (
		a1, b1 = iota, iota + 1
		a2, b2
		a3, b3
	)
	fmt.Println("a1 = ", a1, "b1 = ", b1)
	fmt.Println("a2 = ", a2, "b1 = ", b2)
	fmt.Println("a3 = ", a3, "b1 = ", b3)
}
