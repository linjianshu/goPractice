package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
)

func main() {
	CONNECT := "localhost:8081"
	addr, err := net.ResolveUDPAddr("udp4", CONNECT)
	if err != nil {
		fmt.Println(err)
		return
	}
	udpClient, err := net.DialUDP("udp4", nil, addr)
	if err != nil {
		fmt.Println(err)
		return
	}

	defer udpClient.Close()

	for {
		//读取命令行输入的文字发送
		reader := bufio.NewReader(os.Stdin)
		fmt.Println(">> ")
		text, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println(err)
			continue
		}
		data := []byte(text + "\n")
		_, err = udpClient.Write(data)
		if err != nil {
			fmt.Println(err)
			continue
		}

		if strings.TrimSpace(string(data)) == "STOP" {
			fmt.Println("Exiting UDP client!")
			return
		}

		buffer := make([]byte, 1024)
		n, _, err := udpClient.ReadFromUDP(buffer)
		if err != nil {
			fmt.Println(err)
			continue
		}

		fmt.Printf("Reply: %s\n", string(buffer[0:n]))
	}
}
