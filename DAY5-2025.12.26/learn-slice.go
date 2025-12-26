package main

import (
	"fmt"
	"sort"
)

func practice1() {
	var a = make([]string, 5, 10)
	for i := 0; i < 10; i++ {
		a = append(a, fmt.Sprintf("%v", i))
	}
	//前5个是空字符串
	fmt.Println(a)
}

func practice2() {
	var a = [...]int{3, 7, 8, 9, 1} //自动推断数组长度
	sort.Ints(a[:])                 //数组转换为切片
	fmt.Println(a)

}

func main() {
	//切片 拥有相同类型的可变长度的序列vector  是引用类型 不支持直接比较 只能和nil比较
	//var name []T
	var a []string
	var b = []int{1, 2, 3, 4, 5, 6, 7}
	var c = []bool{true, false}

	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(a == nil)
	fmt.Println(c == nil)
	fmt.Println(c, len(c), cap(c))

	fmt.Println()

	//简单切片表达式 s[low : high]
	//左包含,右不包含
	b1 := b[:3] //[0:3]
	b2 := b[2:] //[2:len(b)]
	b3 := b[:]  //[0:len(b)]
	//切片的cap表示的是从切片的起始位置到底层数组末尾的元素数量
	fmt.Println("b1 = ", b1, "len(b1) = ", len(b1), "cap(b1) = ", cap(b1))
	fmt.Println("b2 = ", b2, "len(b2) = ", len(b2), "cap(b2) = ", cap(b2))
	fmt.Println("b3 = ", b3, "len(b3) = ", len(b3), "cap(b3) = ", cap(b3))

	//对切片再进行切片,high的上限为切片的cap而不是len
	b4 := b2[4:5]
	fmt.Println("b4 = ", b4, "len(b4) = ", len(b4), "cap(b4) = ", cap(b4))

	//完整切片表达式 s[low : high : max] 会额外设置cap为max - low 并且 只有low可以被省略
	// 0 <= low <= high <= max <= cap(s)
	b5 := b[2:5:6]
	fmt.Println("b5 = ", b5, "len(b5) = ", len(b5), "cap(b5) = ", cap(b5))

	//make函数构造切片 make([]T , size , cap)
	d1 := make([]string, 5, 8)
	fmt.Println(d1)
	fmt.Println(len(d1))
	fmt.Println(cap(d1))

	//判断切片是否为空 len(s) == 0
	//不要使用 s == nil

	//赋值拷贝  底层共享数组  因为切片是引用类型
	d2 := d1
	d2[2] = "hello"
	fmt.Println("d1 = ", d1)
	fmt.Println("d2 = ", d2)

	//append方法
	var e []int
	e = append(e, 1) //无须slice初始化
	fmt.Println(e)
	e = append(e, 2, 3, 4) //多个元素
	fmt.Println(e)
	e = append(e, b...) //切片
	fmt.Println(e)
	fmt.Println()

	//copy函数  copy(destslice , srcslice , []T)
	var f = make([]int, len(e), cap(e))
	copy(f, e)
	fmt.Println(f)
	fmt.Println()
	f[3] = 1000
	fmt.Println(f)
	fmt.Println(e)

	//删除元素 想删除下标范围[l , r]的元素 s = append(s[:l] , s[r + 1 :]...)
	f = append(f[:4], f[7:]...)
	fmt.Println(f)

	practice1()
	fmt.Println()
	practice2()
}
