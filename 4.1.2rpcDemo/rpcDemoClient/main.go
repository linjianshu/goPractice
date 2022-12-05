package main

import (
	"fmt"
	"net/rpc"
)

func main() {
	client, err := rpc.DialHTTP("tcp", "localhost:8000")
	if err != nil {
		fmt.Println("链接rpc服务器失败: ", err)
		return
	}
	var reply string
	err = client.Call("Watcher.GetInfo", 1, &reply)
	if err != nil {
		fmt.Println("调用远程服务失败 ", err)
		return
	}
	fmt.Println("远程服务返回结果: ", reply)
}
