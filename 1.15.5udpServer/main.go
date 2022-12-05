package main

import (
	"fmt"
	"math/rand"
	"net"
	"strings"
	"time"
)

func main() {
	PORT := ":8081"
	addr, err := net.ResolveUDPAddr("udp4", PORT)
	if err != nil {
		fmt.Println(err)
		return
	}

	udp, err := net.ListenUDP("udp4", addr)
	if err != nil {
		fmt.Println(err)
		return
	}

	defer udp.Close()

	buffer := make([]byte, 1024)
	rand.Seed(time.Now().UnixNano())

	for {
		//接收数据
		n, addr, err := udp.ReadFromUDP(buffer)
		if err != nil {
			fmt.Println(err)
			continue
		}

		fmt.Println("-> ", string(buffer[0:n-1]))

		if strings.TrimSpace(string(buffer[0:n])) == "STOP" {
			fmt.Println("Exiting UDP server!")
			return
		}

		//发送
		_, err = udp.WriteToUDP(buffer[:n-1], addr)
		if err != nil {
			fmt.Println(err)
			return
		}

	}
}

func random(min, max int) int {
	return rand.Intn(max-min) + min
}
