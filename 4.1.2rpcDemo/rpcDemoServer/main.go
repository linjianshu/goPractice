package main

import (
	"fmt"
	"io"
	"net"
	"net/http"
	"net/rpc"
)

// 访问http://127.0.0.1:8000/debug/rpc看看
func main() {
	http.HandleFunc("/api", hello)
	watcher := new(Watcher)

	// 调用net/rpc包的功能将服务对象进行注册
	err := rpc.Register(watcher)
	if err != nil {
		fmt.Println(err)
		return
	}

	// 通过该函数把watcher中提供的服务注册到HTTP协议上，方便调用者可以利用http的方式进行数据传递
	rpc.HandleHTTP()

	// 在特定的端口进行监听
	listen, err := net.Listen("tcp", ":8000")
	if err != nil {
		fmt.Println("监听失败 端口可能已经被占用", err)
		return
	}
	//可以注册多个服务到同一个端口
	fmt.Println("正在监听8000端口")
	http.Serve(listen, nil)

}

func hello(writer http.ResponseWriter, request *http.Request) {
	io.WriteString(writer, "hello-123")
}

type Watcher string

func (w *Watcher) GetInfo(arg int, result *string) error {
	*result = "helloooooo"
	return nil
}
