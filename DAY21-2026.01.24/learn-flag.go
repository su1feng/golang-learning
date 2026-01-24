package main

import (
	"flag"
	"fmt"
	"os"
)

func UseArgsDemo() {
	//第一个元素是执行文件的名称
	if len(os.Args) > 0 {
		for index, arg := range os.Args {
			fmt.Println("args ", index, " : ", arg)
		}
	}
}

func UseFlagDemo() {
	name := flag.String("name", "su1feng", "姓名")
	age := flag.Uint("age", 23, "年龄")
	married := flag.Bool("married", false, "婚否")
	delay := flag.Duration("d", 0, "时间间隔")

	flag.Parse()
	fmt.Println("name : ", *name)
	fmt.Println("age : ", *age)
	fmt.Println("married : ", *married)
	fmt.Println("delay : ", *delay)
	fmt.Println("other args : ", flag.Args())
	fmt.Println("other args num : ", flag.NArg())
	fmt.Println("flag num : ", flag.NFlag())
}
