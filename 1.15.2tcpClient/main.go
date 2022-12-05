package main

import (
	"fmt"
	"net"
	"os"
)

func main() {
	strEcho := "Halo"
	servAddr := "localhost:3333" //启动端口
	tcpAddr, err := net.ResolveTCPAddr("tcp", servAddr)
	if err != nil {
		fmt.Println("ResolveTCPAddr failed: ", err.Error())
		os.Exit(1)
	}

	//连接tcp服务
	conn, err := net.DialTCP("tcp", nil, tcpAddr)
	if err != nil {
		fmt.Println("Dial failed:", err.Error())
		os.Exit(1)
	}

	//发送消息
	_, err = conn.Write([]byte(strEcho))
	if err != nil {
		fmt.Println("Write to server failed", err.Error())
		os.Exit(1)
	}

	fmt.Println("write to server = ", strEcho)

	//发送后从服务端接收消息
	reply := make([]byte, 1024)
	_, err = conn.Read(reply)
	if err != nil {
		fmt.Println("Write to server failed:", err.Error())
		os.Exit(1)
	}

	fmt.Println("reply from server = ", string(reply))
	//关闭连接
	conn.Close()
}
