package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strings"
)

func bufioDemo() {
	reader := bufio.NewReader(os.Stdin) // 从标准输入生成读对象
	fmt.Print("请输入内容：")
	text, _ := reader.ReadString('\n') // 读到换行
	text = strings.TrimSpace(text)
	fmt.Printf("%#v\n", text)
}

func main() {
	//Fprint 向文件中写内容
	// 向标准输出写入内容
	fmt.Fprintln(os.Stdout, "向标准输出写入内容")
	fileObj, err := os.OpenFile("./su1feng.txt", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
	if err != nil {
		fmt.Println("打开文件出错,err:", err)
		return
	}
	name := "su1feng"
	// 向打开的文件句柄中写入内容
	fmt.Fprintf(fileObj, "往文件中写如信息：%s", name)

	//Sprint
	s1 := fmt.Sprint("su1feng")
	name1 := "su1feng"
	age := 18
	s2 := fmt.Sprintf("name1:%s,age:%d", name1, age)
	s3 := fmt.Sprintln("su1feng")
	fmt.Println(s1, s2, s3)

	//Errorf
	e := errors.New("原始错误e")
	w := fmt.Errorf("Wrap了一个错误%w", e)
	fmt.Println(w)

	//获取输入
	var (
		name2   string
		age2    int
		married bool
	)
	fmt.Scanln(&name2, &age2, &married)
	fmt.Printf("扫描结果 name:%s age:%d married:%t \n", name2, age2, married)

	bufioDemo()

}
