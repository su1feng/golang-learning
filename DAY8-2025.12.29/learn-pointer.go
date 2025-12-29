package main

import (
	"fmt"
)

func pointer() {
	//go中的指针不能偏移和计算 &取地址   *根据地址取值
	a := 10
	b := &a // 取变量a的地址，将指针保存到b中
	fmt.Printf("type of b:%T\n", b)
	c := *b // 指针取值（根据指针去内存取值）
	fmt.Printf("type of c:%T\n", c)
	fmt.Printf("value of c:%v\n", c)

	//new和make   引用类型必须初始化后才拥有内存空间
	//new(Type)得到的是一个类型的指针 对应的值为对应的类型的零值
	//make()仅用于map slice channel  返回的类型就是类型本身而不是指针 因为这三个就是引用类型
}
