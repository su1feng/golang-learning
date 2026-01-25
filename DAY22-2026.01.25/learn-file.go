package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func FileDemo() {
	file, err := os.Open("./Su1feng.log")
	if err != nil {
		fmt.Println("无法打开文件 , err = ", err)
		return
	}

	defer file.Close()

	var buffer = make([]byte, 128)
	size, err := file.Read(buffer)

	if err == io.EOF {
		fmt.Println("文件读完了")
		return
	}

	if err != nil {
		fmt.Println("Read file err = ", err)
		return
	}

	fmt.Println("Read ", size, "bytes : ", string(buffer[:size]))
	fmt.Println()

	scanner := bufio.NewScanner(file) // MaxScanTokenSize = 64 * 1024

	//循环读取每行
	for scanner.Scan() {
		fmt.Println(scanner.Text()) //当前行的内容
		fmt.Println()
	}

	file2, err := os.OpenFile("Su1feng.txt", os.O_CREATE|os.O_TRUNC|os.O_WRONLY, 0666)
	if err != nil {
		fmt.Println("open file2 failed, err:", err)
		return
	}
	defer file2.Close()
	str := "hello Su1feng"
	file2.Write([]byte(str))       //写入字节切片数据
	file2.WriteString("hello msj") //直接写入字符串数据

	writer := bufio.NewWriter(file2)
	for i := 0; i < 10; i++ {
		writer.WriteString("keep going \n") //将数据先写入缓存
	}
	writer.Flush() //将缓存中的内容写入文件

}
