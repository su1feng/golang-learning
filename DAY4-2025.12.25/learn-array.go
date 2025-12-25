package main

import (
	"fmt"
)

func main() {
	//数组array   var  变量名 [数量]T
	var a1 [5]int
	var a2 = [5]int{1, 2, 3, 4}
	var a3 = [5]string{"su1feng", "Suzhou", "23"}
	var a4 = [...]int{1, 2}        //自行推断数组长度
	var a5 = [...]int{3: 5, 9: 20} //下标3的元素为5,下表9的元素为20
	fmt.Println(a1)
	fmt.Println(a2)
	fmt.Println(a3)
	fmt.Println(a4)
	fmt.Println(a5)

	//for遍历
	fmt.Println("for遍历")
	for i := 0; i < len(a1); i++ {
		fmt.Println(a1[i])
	}

	fmt.Println("===========================")

	//for range遍历
	fmt.Println("for range遍历")
	for _, value := range a2 {
		fmt.Println(value)
	}

	//二维数组  只允许第一层可以使用...
	a := [...][5]int{
		{1, 3, 5, 7, 9},
		{2, 4, 6, 8, 10},
	}
	fmt.Println(a)

}
