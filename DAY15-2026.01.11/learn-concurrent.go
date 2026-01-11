package main

import (
	"fmt"
	"sync"
)

//Go语言采用的并发模型是CSP（Communicating Sequential Processes）,提倡通过 通信 共享 内存 而不是通过 共享内存 实现 通信

// 当你并不关心并发操作的结果或者有其它方式收集并发操作的结果时，WaitGroup是实现等待一组并发操作完成的好方法。
var wg sync.WaitGroup

func hello(i int) {
	defer wg.Done() //告知当前goroutine完成
	fmt.Println("hello", i)
}

func practice() {
	//闭包  并且没有等待goroutine完成就退出了,什么都没有打印
	for i := 0; i < 5; i++ {
		go func() {
			fmt.Println(i)
		}()
	}
}

func recv(c chan int) {
	ret := <-c
	fmt.Println("recv val = ", ret)
}

func f2(ch chan int) {
	for {
		v, ok := <-ch
		if !ok {
			fmt.Println("通道已关闭")
			break
		}
		fmt.Printf("v:%#v ok:%#v\n", v, ok)
	}
}

func learn_chan() {
	//var name chan type   name是变量名  type是需要传递的元素的类型
	//未初始化的通道类型变量其默认零值是nil  需要通过make初始化
	//make(chan type , buffer size)
	//有三种方法 send  receive  close  需要使用 <-符号
	channel1 := make(chan int, 5)
	channel1 <- 5

	val := <-channel1
	fmt.Println("val = ", val)

	//无缓冲通道也被称为同步通道
	channel2 := make(chan int)

	fmt.Println("channel2会导致阻塞")

	go recv(channel2)
	fmt.Println("会继续阻塞 直到发送方发来数据")
	channel2 <- 2
	//当通道内已有元素数达到最大容量后，再向通道执行发送操作就会阻塞，除非有从通道执行接收操作
	fmt.Println("=======使用多返回值判断通道是否为空======")
	channel3 := make(chan int, 3)

	// 步骤1：先发送数据到通道
	channel3 <- 7
	channel3 <- 8
	channel3 <- 9

	// 步骤2：关闭通道
	close(channel3)

	// 步骤3：再接收数据
	f2(channel3)

	fmt.Println("======单向通道=======")
	//<- chan int  只接收通道，只能接收不能发送
	//chan <- int  只发送通道，只能发送不能接收
	//在函数传参及任何赋值操作中全向通道（正常通道）可以转换为单向通道，但是无法反向转换。

	/*
		select多路复用

		可处理一个或多个 channel 的发送/接收操作。
		如果多个 case 同时满足，select 会随机选择一个执行。
		对于没有 case 的 select 会一直阻塞，可用于阻塞 main 函数，防止退出。
	*/

}

func main() {
	for i := range 10 {
		wg.Add(1)
		go hello(i)
	}
	fmt.Println("你好")
	wg.Wait()

	fmt.Println()

	fmt.Println("============================")
	practice()
	fmt.Println("============================")
	learn_chan()
}
