package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
)

func main() {
	conn, err := net.Dial("tcp", "localhost:8080")
	if err != nil {
		fmt.Printf("连接服务端失败: %v\n", err)
		return
	}
	defer conn.Close()

	fmt.Println("成功连接到服务端")

	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Print("请输入消息: ")
		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)

		if input == "exit" {
			fmt.Println("断开连接")
			return
		}

		_, err = conn.Write([]byte(input + "\n"))
		if err != nil {
			fmt.Printf("发送消息失败: %v\n", err)
			return
		}

		buffer := make([]byte, 1024)
		n, err := conn.Read(buffer)
		if err != nil {
			fmt.Printf("接收响应失败: %v\n", err)
			return
		}

		fmt.Printf("服务端响应: %s", string(buffer[:n]))
	}
}
