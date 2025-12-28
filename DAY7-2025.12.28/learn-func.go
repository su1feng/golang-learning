package main

import (
	"errors"
	"fmt"
)

func calc(x, y int) (int, int) {
	sum := x + y
	sub := x - y

	return sum, sub
}

func add(x, y int) int {
	return x + y
}

func sub(x, y int) int {
	return x - y
}

func hof_param(x, y int, op func(int, int) int) int {
	return op(x, y)
}

func hof_return(s string) (func(int, int) int, error) {
	switch s {
	case "+":
		return add, nil
	case "-":
		return sub, nil
	default:
		err := errors.New("无法识别的操作符")
		return nil, err
	}
}

func main() {
	//调用有返回值的函数时 可以不接收返回值
	//可变参数通常作为参数的最后一个参数 放在固定参数后面
	//可变参数的本质是通过切片实现的
	//函数支持多返回值 但需要用()包裹
	//当返回值有slice时 nil可以被看做一个有效的slice而无须返回 []int{}
	calc(3, 4)
	res1, res2 := calc(5, 20)
	fmt.Println("res1 = ", res1, "res2 = ", res2)

	//局部变量和全局变量重名时 优先局部

	//定义函数类型 type
	type calculation func(x, y int) int

	var res3 calculation

	res3 = add

	fmt.Println("res3 = ", res3(3, 4))

	res4 := add
	fmt.Println("res4 = ", res4(1, 2))

	//高阶函数 即函数作为参数 函数作为返回值
	hof_param(2, 6, add)

	f, err := hof_return("+")
	if err != nil {
		fmt.Println("error = ", err)
	} else {
		fmt.Println("result = ", f(5, 7))
	}

	//匿名函数 即不添加函数名 需要保存到某个变量或者在函数后加上()立即执行
	//多用于回调和闭包
	add := func(x, y int) {
		fmt.Println(x + y)
	}
	add(10, 20) // 通过变量调用匿名函数

	//自执行函数：匿名函数定义完加()直接执行
	func(x, y int) {
		fmt.Println(x + y)
	}(10, 20)

	//闭包 = 函数 + 引用环境
	//变量捕获 可以访问外部函数的局部变量
	//变量持久化
	//独立环境
	var func1 = adder()
	fmt.Println(func1(10)) //10
	fmt.Println(func1(20)) //30
	fmt.Println(func1(30)) //60

	func2 := adder()
	fmt.Println(func2(40)) //40
	fmt.Println(func2(50)) //90

	//defer语句 执行在return之前
	fmt.Println(f1())
	fmt.Println(f2())
	fmt.Println(f3())
	fmt.Println(f4())

	x := 1
	y := 2
	//会立即执行func5("A" , x , y)
	defer func5("AA", x, func5("A", x, y))
	x = 10
	//会立即执行func5("B" , x , y)
	defer func5("BB", x, func5("B", x, y))
	y = 20

	//recover必须搭配defer使用 defer一定要在可能引发panic的语句之前定义

	practice1()
}

// 闭包
// 函数: func(y int) int{ x += y ; return x}
// 环境: {x : 0 }捕获的变量
func adder() func(int) int {
	var x int
	return func(y int) int {
		x += y
		return x
	}
}

func f1() int {
	x := 5
	//闭包
	defer func() {
		x++
	}()
	return x
	//先 返回值 = x = 5
	//defer x = 6
	//return 5

}

func f2() (x int) {
	//闭包
	defer func() {
		x++
	}()
	return 5
	//先 返回值 x = 5
	//defer x = 6
	//return x
}

func f3() (y int) {
	x := 5
	//闭包
	defer func() {
		x++
	}()
	return x
	//先 返回值y = x = 5
	//defer x = 6
	//return y

}
func f4() (x int) {
	defer func(x int) {
		x++
	}(x)
	return 5
	//先 x = 5
	//defer 参数x是一个新的x = 6
	//返回x = 5
}

func func5(index string, a, b int) int {
	ret := a + b
	fmt.Println(index, a, b, ret)
	return ret
}

func practice1() {
	/*
	   你有50枚金币，需要分配给以下几个人：Matthew,Sarah,Augustus,Heidi,Emilie,Peter,Giana,Adriano,Aaron,Elizabeth。
	   分配规则如下：
	   a. 名字中每包含1个'e'或'E'分1枚金币
	   b. 名字中每包含1个'i'或'I'分2枚金币
	   c. 名字中每包含1个'o'或'O'分3枚金币
	   d: 名字中每包含1个'u'或'U'分4枚金币
	   写一个程序，计算每个用户分到多少金币，以及最后剩余多少金币？
	   程序结构如下，请实现 ‘dispatchCoin’ 函数
	*/
	var (
		coins = 50
		users = []string{
			"Matthew", "Sarah", "Augustus", "Heidi", "Emilie", "Peter", "Giana", "Adriano", "Aaron", "Elizabeth",
		}
		distribution = make(map[string]int, len(users))
	)
	rest := dispatchCoin(coins, users, distribution)
	fmt.Println("剩下 : ", rest)
}

func dispatchCoin(coins int, users []string, distribution map[string]int) int {
	for _, name := range users {
		getCoin := 0
		for _, s := range name {
			switch s {
			case 'e', 'E':
				getCoin += 1
			case 'i', 'I':
				getCoin += 2
			case 'o', 'O':
				getCoin += 3
			case 'u', 'U':
				getCoin += 4
			}
		}
		distribution[name] = getCoin
		coins -= getCoin
	}
	return coins
}
