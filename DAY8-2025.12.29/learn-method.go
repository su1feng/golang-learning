package main

import (
	"fmt"
)

type Person struct {
	name string
	age  int8
}

// NewPerson 构造函数
func NewPerson(name string, age int8) *Person {
	return &Person{
		name: name,
		age:  age,
	}
}

func (p Person) Dream() {
	fmt.Printf("%s的梦想是学好Go语言！\n", p.name)
}

func (p *Person) SetAge(newAge int8) {
	p.age = newAge
}

func main() {
	//Method是一种作用于特定类型变量的函数 特定类型变量叫做Receiver  接收者的类型可以是任意类型
	//可以理解为一个类的方法 Receiver也就是this指针
	//func (接收者变量 接收者类型) 方法名(参数列表) (返回参数){函数体}
	p1 := NewPerson("su1feng", 23)
	p1.Dream()
	fmt.Println(p1)
	p1.SetAge(18)
	fmt.Println(p1)
}
