package main

import (
	"fmt"
	"time"
)

func Practice() {
	now := time.Now()

	fmt.Println(now.Format("2006/01/02 15:04:05"))

	start := time.Now()
	for i := range 20 {
		fmt.Printf("%d", i)
	}
	fmt.Println()
	end := time.Now()
	fmt.Println("cost ", end.Sub(start)*1000)
}
