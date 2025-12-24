package main

import (
	"fmt"
)

func practice() {
	for i := 1; i <= 9; i++ {
		for j := 1; j <= 9; j++ {
			if i > j {
				continue
			}

			fmt.Printf("%d * %d = %-2d |", i, j, i*j)
		}
		fmt.Println()
	}
}

func main() {

	//if
	if score := 65; score >= 90 {
		fmt.Println("A")
	} else if score >= 80 {
		fmt.Println("B")
	} else {
		fmt.Println("C")
	}
	//上述写法score只存在于if判断内的变量
	//fmt.Println("score = " , score)

	//for 格式一般为 for 初始语句; 条件表达式; 结束语句
	//结束语句不能有多个
	for i, j := 0, 1; i < 10; i++ {
		fmt.Printf("i = %d , j = %d \n", i, j)
	}

	//for range 键值循环 遍历数组 切片 字符串 通道 map
	//格式一般为 for index , value := range iterable{ }

	sequence := "hello world"
	for index, value := range sequence {
		fmt.Println(index, string(value))
	}

	//switch
	switch n := 7; n {
	case 1, 3, 5, 7, 9:
		fmt.Println("奇数")
		fallthrough
	case 2, 4, 6, 8:
		fmt.Println("偶数")
		fallthrough
	//case的返回值必须和上面的n相同
	//case n > 20:
	//	fmt.Println("大于20")
	default:
		fmt.Println(n)
	}

	age := 23
	switch {
	case age < 25:
		fmt.Println("好好学习吧")
	case age > 25 && age < 35:
		fmt.Println("好好工作吧")
	case age > 60:
		fmt.Println("好好享受吧")
	default:
		fmt.Println("活着真好")
	}

	practice()

}
