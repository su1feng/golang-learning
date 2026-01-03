package calc

import "fmt"

func Add(nums ...int) int {
	sum := 0
	for _, num := range nums {
		sum += num
	}
	return sum
}

func Sub(num1, num2 int) int {
	return num1 - num2
}

func Multi(nums ...int) int {
	sum := 1
	for _, num := range nums {
		sum *= num
	}
	return sum
}

func Div(num1, num2 int) int {
	if num2 == 0 {
		fmt.Println("除数不能为0")
		return 0
	}
	return num1 / num2
}
