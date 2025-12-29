package main

import (
	"fmt"
)

type person struct {
	name, city string
	age        uint8
}

func main() {
	//自定义类型
	type Myint int
	//类型别名  举例:type byte =uint8    type rune =int32
	type TypeAlias = int

	//自定义类型和类型别名的区别
	var a Myint
	var b TypeAlias

	fmt.Printf("type of a : %T\n", a) //main.Myint
	fmt.Printf("type of b : %T\n", b) // int
	//类型别名只在代码中存在 编译完成时不会有类型别名

	//结构体定义 type name struct{}  只有当实例化时才会分配内存   没有初始化成员变量都为对应的零值
	//结构体占用一块连续的内存
	//空结构体不占空间 var v struct{}
	//匿名字段 即没有字段名只有类型 默认采用类型名作为字段名 字段名必须唯一

	//定义临时数据结构时可以使用匿名结构体
	var user struct {
		name string
		age  uint8
	}
	user.name = "su1feng"
	user.age = 23
	fmt.Printf("%#v\n", user)

	//创建指针类型结构体
	var p1 = new(person)
	fmt.Printf("%T\n", p1)
	fmt.Printf("p1 = %#v\n", p1)

	//使用& 相当于对结构体进行了一次new
	p3 := &person{}
	fmt.Printf("%T\n", p3)     //*main.person
	fmt.Printf("p3=%#v\n", p3) //p3=&main.person{name:"", city:"", age:0}
	p3.name = "七米"             //底层是 (*p3).name = "七米"
	p3.age = 30
	p3.city = "成都"
	fmt.Printf("p3=%#v\n", p3) //p3=&main.person{name:"七米", city:"成都", age:30}

	//键值对初始化
	p5 := person{
		name: "小王子",
		city: "北京",
		age:  18,
	}
	fmt.Printf("p5=%#v\n", p5) //p5=main.person{name:"小王子", city:"北京", age:18}

	//对结构体指针进行键值对初始化
	p6 := &person{
		name: "小王子",
		city: "北京",
		//可以不写age
	}
	fmt.Printf("p6=%#v\n", p6) //p6=&main.person{name:"小王子", city:"北京", age:0}

	//这种格式必须初始化所有字段 填充顺序与声明顺序一致 不能和键值对初始化混用
	p8 := &person{
		"沙河娜扎",
		"北京",
		28,
	}
	fmt.Printf("p8=%#v\n", p8) //p8=&main.person{name:"沙河娜扎", city:"北京", age:28}

	practice()

	//go的结构体没有构造函数 但可以自己实现
	p9 := newPerson("su1feng", "Suzhou", 23)
	fmt.Printf("%#v\n", p9)
}

// 构造函数
func newPerson(name, city string, age uint8) *person {
	return &person{
		name: name,
		city: city,
		age:  age,
	}
}

func practice() {
	type student struct {
		name string
		age  int
	}

	m := make(map[string]*student)
	stus := []student{
		{name: "小王子", age: 18},
		{name: "娜扎", age: 23},
		{name: "大王八", age: 9000},
	}

	for _, stu := range stus {
		m[stu.name] = &stu
	}
	for k, v := range m {
		fmt.Println(k, "=>", v.name)
	}
}
