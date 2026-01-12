package main

import (
	"fmt"
	"net/http"
	"time"

	"golang.org/x/sync/errgroup"
)

func f2() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Printf("recover outer panic:%v\n", r)
		}
	}()
	// 开启一个goroutine执行任务
	go func() {
		defer func() {
			if r := recover(); r != nil {
				fmt.Printf("recover inner panic:%v\n", r)
			}
		}()
		fmt.Println("in goroutine....")
		// 只能触发当前goroutine中的defer
		panic("panic in goroutine")
	}()

	time.Sleep(time.Second)
	fmt.Println("exit")
}

// fetchUrlDemo2 使用errgroup并发获取url内容
func fetchUrlDemo2() error {
	g := new(errgroup.Group) // 创建等待组（类似sync.WaitGroup）
	var urls = []string{
		"https://golang.org",
		"https://www.liwenzhou.com",
		"https://www.yixieqitawangzhi.com",
	}
	for _, url := range urls {
		url := url // 注意此处声明新的变量
		// 启动一个goroutine去获取url内容
		g.Go(func() error {
			resp, err := http.Get(url)
			if err == nil {
				fmt.Printf("获取%s成功\n", url)
				resp.Body.Close()
			}
			return err // 返回错误
		})
	}
	if err := g.Wait(); err != nil {
		// 处理可能出现的错误
		fmt.Println(err)
		return err
	}
	fmt.Println("所有goroutine均成功")
	return nil
}

func main() {
	/*
		panic 只会触发当前 goroutine 中的 defer 操作
		如果想要 recover goroutine中可能出现的 panic 就需要在 goroutine 中使用 recover
	*/
	f2()

	/*
		errgroup 包能为处理公共任务的子任务而开启的一组 goroutine 提供同步、error 传播和基于context 的取消功能。
		errgroup.Group提供了Go和Wait方法

		Go 函数会在新的 goroutine 中调用传入的函数f
		第一个返回非零错误的调用将取消该Group；下面的Wait方法会返回该错误

		Wait 会阻塞直至由上述 Go 方法调用的所有函数都返回，然后从它们返回第一个非nil的错误（如果有）
	*/
	fetchUrlDemo2()
}
