package main

import (
	"fmt"
	"net"
	"os"
)

func main() {
	//监听传入的连接
	//tcp localhost:3333
	l, err := net.Listen(CONN_TYPE, CONN_HOST+":"+CONN_PORT)
	if err != nil {
		fmt.Println("Error listening:", err.Error())
		os.Exit(1)
	}

	//当应用程序关闭时关闭监听器
	defer l.Close()
	fmt.Println("Listening on " + CONN_HOST + ":" + CONN_PORT)
	for {
		//监听传入的连接
		conn, err := l.Accept()
		if err != nil {
			fmt.Println("Error accepting: ", err.Error())
			os.Exit(1)
		}
		//在新的goroutine中处理连接
		go handleRequest(conn)
	}
}

func handleRequest(conn net.Conn) {
	//创建一个缓冲区以保存传入的数据
	buf := make([]byte, 1024)
	//将传入的连接读入缓冲区
	reqLen, err := conn.Read(buf)
	if err != nil {
		fmt.Println("Error reading:", err.Error())
	}
	fmt.Println(reqLen)
	fmt.Println(string(buf))

	//将回复发送给与我们联系的人
	conn.Write([]byte("Message received"))

	//完成连接后,关闭连接
	conn.Close()
}

const (
	CONN_HOST = "localhost"
	CONN_PORT = "3333"
	CONN_TYPE = "tcp"
)
