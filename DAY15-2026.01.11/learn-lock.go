package main

import (
	"fmt"
	"strconv"
	"sync"
)

/*
互斥锁
sync.Mutex
提供Lock和Unlock两种方法

读写锁
sync.RWMutex
写锁 Lock Unlock
读锁 RLock RUnlock
返回一个实现Locker接口的读写锁  RLocker


只加载一次
sync.Once
提供Do方法 func (o *Once) Do(f func())
如果要执行的函数f需要传递参数就需要搭配闭包来使用

sync.Map
提供Store Load LoadOrStore LoadAndDelete Delete Range方法

原子操作atomic
*/

var m = sync.Map{}

func test_sync_Map() {
	wg := sync.WaitGroup{}
	// 对m执行20个并发的读写操作
	for i := 0; i < 20; i++ {
		wg.Add(1)
		go func(n int) {
			key := strconv.Itoa(n)
			m.Store(key, n)         // 存储key-value
			value, _ := m.Load(key) // 根据key取值
			fmt.Printf("k=:%v,v:=%v\n", key, value)
			wg.Done()
		}(i)
	}
	wg.Wait()
}

func main() {
	test_sync_Map()
}
