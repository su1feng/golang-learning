package main

import (
	"fmt"
)

type animal interface {
	hungry()
}

func makehungry(a animal) {
	a.hungry()
}

type cat struct {
	name string
}

func (c cat) hungry() {
	fmt.Println(c.name + "饿了 喵喵喵")
}

type dog struct {
	name string
}

func (d dog) hungry() {
	fmt.Println(d.name + "饿了 汪汪汪")
}

func main() {
	/*
				type 接口类型名 interface{
		    	方法名1( 参数列表1 ) 返回值列表1
		   		方法名2( 参数列表2 ) 返回值列表2
		    	…
				}
				方法名：当方法名首字母是大写且这个接口类型名首字母也是大写时，这个方法可以被接口所在的包（package）之外的代码访问。
				接口就是规定了一个需要实现的方法列表，在 Go 语言中一个类型只要实现了接口中规定的所有方法，那么我们就称它实现了这个接口
				例子: 很多动物都会饿,我们不需要关注是哪一种动物饿,只需要实现一个animal类型接口interface,然后将动物作为animal参数传入即可
	*/

	kitty := cat{name: "小猫"}
	wangcai := dog{name: "小狗"}

	makehungry(kitty)
	makehungry(wangcai)

	/*
		对于值接收者实现的接口，无论使用值类型还是指针类型都没有问题。
		但是我们并不总是能对一个值求址，所以对于指针接收者实现的接口要额外注意。

		接口组合   通过在结构体中嵌入一个接口类型，从而让该结构体类型实现了该接口类型，并且还可以改写该接口的方法。
		type Reader interface {
			Read(p []byte) (n int, err error)
		}

		type Writer interface {
			Write(p []byte) (n int, err error)
		}

		// ReadWriter 是组合Reader接口和Writer接口形成的新接口类型
		type ReadWriter interface {
			Reader
			Writer
		}
	*/

	//空接口 我们不能对一个空接口值调用任何方法，否则会产生panic
	//var x animal{} // 声明一个空接口类型变量x

	//应用1 作为函数参数
	/*
		空接口作为函数参数
		func show(a interface{}) {
			fmt.Printf("type:%T value:%v\n", a, a)
		}
	*/

	//应用2 作为map的值
	// 空接口作为map值
	var studentInfo = make(map[string]interface{})
	studentInfo["name"] = "沙河娜扎"
	studentInfo["age"] = 18
	studentInfo["married"] = false
	fmt.Println(studentInfo)

	fmt.Println()

	//接口值 分为类型和值
	//接口值是支持相互比较的，当且仅当接口值的动态类型和动态值都相等时才相等
	//如果接口值保存的动态类型相同，但是这个动态类型不支持互相比较（比如切片），那么对它们相互比较时就会引发panic

	//断言
	var m animal

	m = &cat{name: "旺财"}
	fmt.Printf("%T\n", m)

	m = new(dog)
	fmt.Printf("%T\n", m)

	v, ok := m.(*cat)
	if ok {
		fmt.Println("断言成功")
		v.name = "kitty"
	} else {
		fmt.Println("断言失败")
	}

	//var _ interface = (*struct)(nil)  可以检查某个struct是否实现了对应的interface
}
