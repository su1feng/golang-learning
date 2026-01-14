package main

import (
	"bufio"
	"fmt"
	"net"
	"strings"
)

func main() {
	listener, err := net.Listen("tcp", ":8080")
	if err != nil {
		fmt.Printf("监听失败: %v\n", err)
		return
	}
	defer listener.Close()

	fmt.Println("服务端启动，监听端口 8080...")

	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Printf("接受连接失败: %v\n", err)
			continue
		}

		go handleClient(conn)
	}
}

func handleClient(conn net.Conn) {
	defer conn.Close()
	fmt.Printf("客户端已连接: %v\n", conn.RemoteAddr())

	reader := bufio.NewReader(conn)
	for {
		message, err := reader.ReadString('\n')
		if err != nil {
			fmt.Printf("客户端断开连接: %v\n", err)
			return
		}

		message = strings.TrimSpace(message)
		fmt.Printf("收到客户端消息: %s\n", message)

		response := fmt.Sprintf("服务端收到: %s\n", message)
		_, err = conn.Write([]byte(response))
		if err != nil {
			fmt.Printf("发送响应失败: %v\n", err)
			return
		}
	}
}
