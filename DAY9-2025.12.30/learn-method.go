package main

import (
	"fmt"
)

// Address 地址结构体
type Address struct {
	Province string
	City     string
}

// User 用户结构体  嵌套结构体
type User struct {
	Name    string
	Gender  string
	Address //匿名字段
}




type Animal struct {
	name string
}

func (a *Animal) move() {
	fmt.Println(a.name, "会动!")
}

type Dog struct {
	Feet int8
	*Animal //通过匿名结构体实现继承
}

func (d *Dog) wang() {
	fmt.Println(d.name, "会汪汪汪!")
}

func main() {
	var user2 User
	user2.Name = "小王子"
	user2.Gender = "男"
	user2.Address.Province = "山东"    // 匿名字段默认使用类型名作为字段名
	user2.City = "威海"                // 匿名字段可以省略
	fmt.Printf("user2=%#v\n", user2) //user2=main.User{Name:"小王子", Gender:"男", Address:main.Address{Province:"山东", City:"威海"}}

	//继承
	d1 := Dog{
		Feet: 4,
		Animal: &Animal{
			name: "修狗",
		},
	}

	d1.move()
	d1.wang()
	//结构体中字段大写开头表示可公开访问，小写表示私有（仅在定义当前结构体的包中可访问）。
}
