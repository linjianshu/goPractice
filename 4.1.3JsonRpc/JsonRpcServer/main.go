package main

import (
	"errors"
	"fmt"
	"net"
	"net/rpc"
	"net/rpc/jsonrpc"
	"os"
)

func main() {
	err := rpc.Register(new(Arith)) //注册rpc服务
	if err != nil {
		panic(err)
	}

	listen, err := net.Listen("tcp", "localhost:8096")
	if err != nil {
		panic(err)
	}

	fmt.Fprintf(os.Stdout, "%s\n", "start connection")
	for {
		conn, err := listen.Accept() //接收客户端连接请求
		if err != nil {
			fmt.Println("accepting error: ", err.Error())
			continue
		}
		//并发处理客户端请求
		go func(conn net.Conn) {
			fmt.Fprintf(os.Stdout, "%s ", "new client in coming \n")
			jsonrpc.ServeConn(conn)
		}(conn)
	}
}

// Arith 算数运算结构体
type Arith struct {
}

// ArithRequest 算数运算请求结构体
type ArithRequest struct {
	A int
	B int
}

// ArithResponse 算数运算响应结构体
type ArithResponse struct {
	Pro int //乘法
	Quo int //除法
	Rem int //余数
}

func (this *Arith) Multiply(req ArithRequest, res *ArithResponse) error {
	res.Pro = req.A * req.B
	return nil
}

func (this *Arith) Divide(req ArithRequest, res *ArithResponse) error {
	if req.B == 0 {
		return errors.New("divide by zero")
	}
	res.Quo = req.A / req.B
	res.Rem = req.A % req.B
	return nil
}
