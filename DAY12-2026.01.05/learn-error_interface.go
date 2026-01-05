package main

import (
	"errors"
	"fmt"
)

// divide 函数演示如何创建和返回错误
func divide(a, b int) (int, error) {
	if b == 0 {
		// 使用 errors.New 创建错误
		return 0, errors.New("除数不能为零")
	}
	return a / b, nil
}

func main() {
	/*
		当一个函数或方法需要返回错误时，我们通常是把错误作为最后一个返回值。
		由于 error 是一个接口类型，默认零值为nil。所以我们通常将调用函数返回的错误与nil进行比较，以此来判断函数是否返回错误。
	*/

	//创建错误  通过error提供的new方法创建一个错误 func New(text string) error

	// 示例1: 直接创建错误
	err1 := errors.New("这是一个错误")
	fmt.Println("err1:", err1)

	// 示例2: 在函数中创建并返回错误
	result, err := divide(10, 0)
	if err != nil {
		fmt.Println("除法出错:", err)
	} else {
		fmt.Println("结果:", result)
	}

	// 示例3: 成功的情况
	result, err = divide(10, 2)
	if err != nil {
		fmt.Println("除法出错:", err)
	} else {
		fmt.Println("结果:", result)
	}

	fmt.Println("\n========== fmt.Errorf 示例 ==========")
	// 示例4: fmt.Errorf - 格式化错误信息
	err4 := fmt.Errorf("操作失败: %s", "文件不存在")
	fmt.Println("err4:", err4)

	// 示例5: fmt.Errorf 包装错误（使用 %w 格式化动词）
	originalErr := errors.New("原始错误")
	wrappedErr := fmt.Errorf("包装错误: %w", originalErr)
	fmt.Println("wrappedErr:", wrappedErr)

	fmt.Println("\n========== errors.Is 和 errors.As 示例 ==========")
	// 示例6: errors.Is - 检查错误链中是否包含特定错误
	if errors.Is(wrappedErr, originalErr) {
		fmt.Println("wrappedErr 包含 originalErr")
	}

	// 示例7: errors.As - 检查错误链中是否包含特定类型的错误
	var myErr *MyError
	err7 := &MyError{Code: 404, Msg: "资源未找到"}
	if errors.As(err7, &myErr) {
		fmt.Printf("错误类型匹配: Code=%d, Msg=%s\n", myErr.Code, myErr.Msg)
	}

	fmt.Println("\n========== 自定义错误类型示例 ==========")
	// 示例8: 自定义错误类型
	err8 := &MyError{Code: 500, Msg: "服务器内部错误"}
	fmt.Println("自定义错误:", err8)

	// 示例9: 使用自定义错误类型
	result, err = safeDivide(10, 0)
	if err != nil {
		if myDivErr, ok := err.(*DivideError); ok {
			fmt.Printf("除法错误详情 - 被除数: %d, 除数: %d, 错误: %s\n",
				myDivErr.Dividend, myDivErr.Divisor, myDivErr)
		}
	}

	fmt.Println("\n========== errors.Unwrap 示例 ==========")
	// 示例10: errors.Unwrap - 解包错误
	unwrapped := errors.Unwrap(wrappedErr)
	if unwrapped != nil {
		fmt.Println("解包后的错误:", unwrapped)
	}
}

// MyError 自定义错误类型，实现 error 接口
type MyError struct {
	Code int
	Msg  string
}

// Error 方法实现 error 接口
func (e *MyError) Error() string {
	return fmt.Sprintf("错误代码: %d, 错误信息: %s", e.Code, e.Msg)
}

// DivideError 除法错误类型
type DivideError struct {
	Dividend int
	Divisor  int
}

// Error 方法实现 error 接口
func (e *DivideError) Error() string {
	return fmt.Sprintf("除法错误: %d / %d", e.Dividend, e.Divisor)
}

// safeDivide 使用自定义错误类型的除法函数
func safeDivide(a, b int) (int, error) {
	if b == 0 {
		return 0, &DivideError{Dividend: a, Divisor: b}
	}
	return a / b, nil
}
