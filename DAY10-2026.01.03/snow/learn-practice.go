package main

import(
	"fmt"
	"practice/calc"
)

func main(){
	res_add := calc.Add(3 , 5 , 7 , 9)
	res_mul := calc.Multi(3 , 5 , 7 , 9)
	res_div := calc.Div(10 , 5)
	res_sub := calc.Sub(res_add , res_mul)

	fmt.Println("res_add = " , res_add)
	fmt.Println("res_sub = " , res_sub)
	fmt.Println("res_div = " , res_div)
	fmt.Println("res_mul = " , res_mul)
}