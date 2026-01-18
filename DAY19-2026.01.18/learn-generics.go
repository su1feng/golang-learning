package main

import (
	"fmt"
)

func reverseWithGenerics[T any](s []T) []T {
	l := len(s)
	r := make([]T, l)

	for i, e := range s {
		r[l-i-1] = e
	}
	return r
}

func min[T int | float64](a, b T) T {
	if a <= b {
		return a
	}
	return b
}

// 除了函数中支持使用类型参数列表外  类型也可以使用类型参数列表
type Slice[T int | string] []T

type Map[K int | string, V float32 | float64] map[K]V

type Tree[T interface{}] struct {
	left, right *Tree[T]
	value       T
}

// 事先定义好的类型约束类型
type Value interface {
	int | float64
}

func max[T Value](a, b T) T {
	if a <= b {
		return b
	}
	return a
}

func main() {
	//向函数提供类型参数称为实例化
	m1 := min[int](1, 2) // 1
	m2 := min[float64](-0.1, -0.2)
	fmt.Println("m1 = ", m1)
	fmt.Println("m2 = ", m2)

	//要使用泛型类型 必须进行实例化 Tree[string]是使用类型实参string实例化 Tree 的示例
	var stringTree Tree[string]

	//当类型参数只出现在返回值中 就必须显式指定
}

// 泛型类型可以有方法 例如为上面的Tree实现一个查找元素的Lookup方法
func (t *Tree[T]) Lookup(x T) *Tree[T] {

}
