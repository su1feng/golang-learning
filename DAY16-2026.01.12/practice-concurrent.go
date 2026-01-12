package main

import (
	"fmt"
	"math/rand/v2"
	"strconv"
	"sync"
	"time"
)

var wg sync.WaitGroup
var once sync.Once
var mutex sync.Mutex

func Producer(channel chan<- int64) {
	defer wg.Done()

	for range 24 {
		var randnum int64 = rand.Int64()
		fmt.Println("produce a number = ", randnum)
		channel <- randnum
	}

	close(channel)
}

func Consumer(channel <-chan int64, result chan<- int64) {
	defer wg.Done()

	for {
		i, ok := <-channel
		if !ok {
			break
		}
		fmt.Println("consume a number = ", i)
		int64Num := i
		lenNum := len(strconv.FormatInt(int64Num, 10))
		var sum int64 = 0
		for i := 1; i <= lenNum; i++ {
			last := int64Num % 10
			sum = sum + last
			int64Num = int64Num / 10
		}
		result <- sum
	}
}

func main() {
	start := time.Now()
	jobChan := make(chan int64, 100)
	results := make(chan int64, 100)

	wg.Add(1)
	go Producer(jobChan)

	for i := 1; i <= 24; i++ {
		wg.Add(1)
		go Consumer(jobChan, results)
	}
	wg.Wait()
	close(results)

	for i := range results {
		fmt.Println(i)
	}
	end := time.Now()
	fmt.Println(end.Sub(start))
}
