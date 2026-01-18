package main

import (
	"fmt"

	"golang.org/x/exp/constraints"
)

// ============================================
// 手动定义 constraints（避免依赖外部包）
// ============================================

// Print 可以接受任何类型
func Print[T any](v T) {
	fmt.Printf("Print: %v (%T)\n", v, v)
}

// ❌ 错误示例：any 不能保证可比较
// func BadEqual[T any](a, b T) bool {
//     return a == b  // 编译错误
// }

// ============================================
// 2. comparable 约束 - 可比较类型
// ============================================

// Equal 比较两个值是否相等
func Equal[T comparable](a, b T) bool {
	return a == b
}

// Find 在切片中查找元素（需要 == 比较）
func Find[T comparable](slice []T, target T) int {
	for i, v := range slice {
		if v == target {
			return i
		}
	}
	return -1
}

// ============================================
// 3. ~ 近似类型约束
// ============================================

// 定义自定义类型
type MyInt int
type MyFloat float64
type MyString string

// 精确类型约束：只能是 int
func ShowExact[T int](v T) {
	fmt.Printf("ShowExact: %v (only exact int)\n", v)
}

// 近似类型约束：int 及底层为 int 的类型
func ShowApprox[T ~int](v T) {
	fmt.Printf("ShowApprox: %v (int or underlying int)\n", v)
}

// ============================================
// 4. 联合约束 |
// ============================================

// Number 支持多种数值类型
func Max[T int | int64 | float64 | float32](a, b T) T {
	if a > b {
		return a
	}
	return b
}

// 使用接口定义约束（更清晰）
type Number interface {
	int | int64 | float64 | float32
}

func Add[T Number](a, b T) T {
	return a + b
}

// ============================================
// 5. 标准库 constraints 包（推荐）
// ============================================

// Ordered 表示支持 < <= >= > 的类型
func Min[T constraints.Ordered](a, b T) T {
	if a < b {
		return a
	}
	return b
}

// Integer 表示所有整数类型
func Double[T constraints.Integer](n T) T {
	return n * 2
}

// Signed 表示有符号整数
func Abs[T constraints.Signed](n T) T {
	if n < 0 {
		return -n
	}
	return n
}

// Float 表示浮点数
func Half[T constraints.Float](n T) T {
	return n / 2
}

// ============================================
// 6. 方法约束（接口）
// ============================================

// Stringer 约束：必须有 String() 方法
type Stringer interface {
	String() string
}

func Describe[T Stringer](v T) {
	fmt.Println("Description:", v.String())
}

// ============================================
// 7. 泛型类型
// ============================================

// Stack 泛型栈
type Stack[T any] struct {
	elements []T
}

func (s *Stack[T]) Push(v T) {
	s.elements = append(s.elements, v)
}

func (s *Stack[T]) Pop() (T, bool) {
	var zero T
	if len(s.elements) == 0 {
		return zero, false
	}
	idx := len(s.elements) - 1
	val := s.elements[idx]
	s.elements = s.elements[:idx]
	return val, true
}

// ============================================
// 8. 泛型 + comparable = 可以做 map 的 key
// ============================================

// Counter 统计元素出现次数
func Counter[T comparable](items []T) map[T]int {
	count := make(map[T]int)
	for _, item := range items {
		count[item]++
	}
	return count
}

// ============================================
// 9. 类型推断
// ============================================

// 当类型参数可以推断时，不需要显式指定
func InferenceDemo() {
	// 显式指定类型
	min1 := Min[int](5, 3)
	fmt.Println("Min[int](5,3) =", min1)

	// 类型推断（编译器自动推导）
	min2 := Min(5, 3)       // 推断为 int
	min3 := Min(3.14, 2.71) // 推断为 float64
	min4 := Min("a", "b")   // 推断为 string

	fmt.Printf("Min(5,3)=%v, Min(3.14,2.71)=%v, Min(\"a\",\"b\")=%v\n",
		min2, min3, min4)
}

// ============================================
// 10. 约束中的 ~ 使用场景
// ============================================

type Celsius float64
type Fahrenheit float64

// 没有 ~：只能用 float64
func ToFloat64[T float64](temp T) float64 {
	return float64(temp)
}

// 有 ~：可以用 Celsius 和 Fahrenheit
func ToFloat64Approx[T ~float64](temp T) float64 {
	return float64(temp)
}

// ============================================
// main 函数演示
// ============================================

func main() {
	fmt.Println("========== 1. any 约束 ==========")
	Print(42)
	Print("hello")
	Print(3.14)

	fmt.Println("\n========== 2. comparable 约束 ==========")
	fmt.Println("Equal(1, 1) =", Equal(1, 1))
	fmt.Println("Equal(\"a\", \"b\") =", Equal("a", "b"))

	// Find 使用示例
	nums := []int{1, 2, 3, 4, 5}
	idx := Find(nums, 3)
	fmt.Printf("Find([1,2,3,4,5], 3) = %d\n", idx)

	fmt.Println("\n========== 3. ~ 近似类型 ==========")
	var mi MyInt = 100
	// ShowExact(mi)  // ❌ MyInt 不是精确的 int
	ShowApprox(mi) // ✅ MyInt 底层是 int

	fmt.Println("\n========== 4. 联合约束 ==========")
	fmt.Println("Max(5, 10) =", Max(5, 10))
	fmt.Println("Max(3.14, 2.71) =", Max(3.14, 2.71))
	fmt.Println("Add(5, 3) =", Add(5, 3))

	fmt.Println("\n========== 5. constraints 包 ==========")
	fmt.Println("Min(5, 10) =", Min(5, 10))
	fmt.Println("Min(\"b\", \"a\") =", Min("b", "a"))
	fmt.Println("Double(21) =", Double(21))
	fmt.Println("Abs(-5) =", Abs(-5))
	fmt.Println("Half(3.14) =", Half(3.14))

	fmt.Println("\n========== 6. 类型推断 ==========")
	InferenceDemo()

	fmt.Println("\n========== 7. 泛型类型 Stack ==========")
	stack := Stack[string]{}
	stack.Push("first")
	stack.Push("second")
	for {
		val, ok := stack.Pop()
		if !ok {
			break
		}
		fmt.Println("Popped:", val)
	}

	fmt.Println("\n========== 8. Counter (需要 comparable) ==========")
	words := []string{"go", "java", "go", "python", "java", "go"}
	counts := Counter(words)
	fmt.Println("Word counts:", counts)

	fmt.Println("\n========== 9. ~ 约束实际应用 ==========")
	var c Celsius = 25.5
	// fmt.Println("ToFloat64(c) =", ToFloat64(c))  // ❌
	fmt.Println("ToFloat64Approx(c) =", ToFloat64Approx(c)) // ✅
}

// ============================================
// 课后练习
// ============================================

/*
练习 1: 实现一个泛型函数 Filter，保留满足条件的元素
func Filter[T any](slice []T, pred func(T) bool) []T

练习 2: 实现泛型函数 Map，转换切片元素
func Map[T, U any](slice []T, f func(T) U) []U

练习 3: 实现泛型类型 Queue（队列）
type Queue[T any] struct { ... }
func (q *Queue[T]) Enqueue(v T)
func (q *Queue[T]) Dequeue() (T, bool)
*/
