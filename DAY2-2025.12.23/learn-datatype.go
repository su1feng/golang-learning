package main

import (
	"fmt"
	"math"
	"unicode"
)

func numbersystem() {
	a := 0b0101  //二进制的0101
	b := 0o377   //八进制的377
	c := 0x80p-2 //p表示二进制指数，用在十六进制浮点数表示中格式为0x整数 p 指数
	v := 123_456 //可以用_分隔数字
	fmt.Println("a = ", a, "\n", "b = ", b, "\n", "c = ", c, "\n", "v = ", v)

	var num = 10
	fmt.Printf("%d \n", num) //十进制
	fmt.Printf("%b \n", num) //二进制
	fmt.Printf("%o \n", num) //八进制
	fmt.Printf("%x \n", num) //十六进制
}

func float_show() {
	//浮点数
	var pi = 3.14
	fmt.Printf("%f \n", pi)
	fmt.Printf("%.2f \n", pi)

	//e是科学计数法的表示
	fmt.Printf("maxfloat32 = %e\n", math.MaxFloat32)
	fmt.Printf("maxfloat64 = %e\n", math.MaxFloat64)
}

func plural() {
	//复数
	m := 1 + 2i
	fmt.Printf("m = %v , 类型是 %T \n", m, m)
}

func stringshow() {
	name := "su1feng"
	//多行字符串
	s := `this is first line
	this is second line
	this is thrid line
	`
	fmt.Printf("多行字符串: %s \n", s)

	length := len(name)
	fmt.Printf("length = %d , (类型%T)\n", length, length)

	s1 := "hello"
	s2 := "go"
	result1 := s1 + "," + s2
	fmt.Printf("result1 = %s , (类型%T)\n", result1, result1)

	result2 := fmt.Sprintf("s1 = %s , s2 = %s , s1 + s2 = %s\n", s1, s2, s1+s2)
	fmt.Printf("result2 = %s\n", result2)

	p := "hello沙河"
	for i := 0; i < len(p); i++ { //byte
		fmt.Printf("%v(%c) ", p[i], p[i])
	}

	fmt.Println() //打印换行

	for _, r := range p { //rune
		fmt.Printf("%v(%c) ", r, r)
	}
	fmt.Println() //打印换行
}

func typeconversion() {
	//强制类型转换 T(表达式)
	var pi = 3.14
	var n = int(pi * pi)
	fmt.Println("n = ", n)
}

func changeString() {
	s1 := "big"
	// 强制类型转换
	byteS1 := []byte(s1)
	byteS1[0] = 'p'
	fmt.Println(string(byteS1))

	s2 := "白萝卜"
	runeS2 := []rune(s2)
	runeS2[0] = '红'
	fmt.Println(string(runeS2))
}

func practice1() {
	a1 := 10
	a2 := 3.1415
	a3 := true
	a4 := "hello go"

	fmt.Printf("a1 = %d , (类型%T)\n", a1, a1)
	fmt.Printf("a2 = %f , (类型%T)\n", a2, a2)
	fmt.Printf("a3 = %v , (类型%T)\n", a3, a3)
	fmt.Printf("a4 = %s , (类型%T)\n", a4, a4)

}

func practice2() {
	s := "hello沙河小王子"
	fmt.Println() //打印换行
	count := 0
	for _, r := range s { //rune
		fmt.Printf("%v(%c) ", r, r)
		if unicode.Is(unicode.Han, r) {
			count++
		}
	}
	fmt.Println()
	fmt.Printf("s当中一共有 %d 个汉字", count)
	fmt.Println()
}

func main() {
	/*
		整型
		浮点型
		bool型 默认为false 不能参与数值运算 不允许整型强转bool  无法和其他类型转换
		复数 complex64 complex128
		字符串

		多行字符串用` `

		修改字符串需要先将字符串转换为[]rune 或 []byte 修改完成后再转回string
		会重新分配内存 并复制字节数组

		字符 rune byte
		强制类型转换 T(表达式)
		没有隐式类型转换
	*/
	numbersystem()
	float_show()
	plural()
	stringshow()
	typeconversion()
	changeString()

	fmt.Println("===============以下是练习1===============")
	practice1()
	fmt.Println("===============以下是练习2===============")
	practice2()
	//bool默认为false,不能参与数值运算 不允许整型强转bool  无法和其他类型转换

	//字符串

}
