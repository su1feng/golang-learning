package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	logFile, err := os.OpenFile("./Su1feng.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
	if err != nil {
		fmt.Println("open log file failed, err:", err)
		return
	}
	log.SetFlags(log.Llongfile | log.Lmicroseconds | log.Ldate)
	log.SetPrefix("[Su1feng]")
	log.SetOutput(logFile)

	log.Println("这是第一条日志")

	logger := log.New(os.Stdout, "<New>", log.Llongfile|log.Ldate)
	logger.Println("New 的第一条日志")

	FileDemo()

	// Fatal系列函数会在写入日志信息后调用 os.Exit(1)  不执行defer  无法被recover捕获
	// Panic系列函数会在写入日志信息后panic  执行defer  可以被recover捕获
	file, err := os.Open("config.json")
	if err != nil {
		log.Fatalln("无法打开配置文件", err)
	}
	log.Println(file)

	defer func() {
		if r := recover(); r != nil {
			log.Println("捕获到panic : ", r)
		}
	}()
}
