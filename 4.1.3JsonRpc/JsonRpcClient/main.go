package main

import (
	"fmt"
	"log"
	"net/rpc/jsonrpc"
)

// ArithRequest 算数运算请求结构体
type ArithRequest struct {
	A int
	B int
}

// ArithResponse 算数运算响应结构体
type ArithResponse struct {
	Pro int
	Quo int
	Rem int
}

func main() {
	conn, err := jsonrpc.Dial("tcp", "localhost:8096")
	if err != nil {
		log.Fatalln("dialing error: ", err)
		panic(err)
	}
	request := ArithRequest{9, 2}
	var res ArithResponse
	err = conn.Call("Arith.Multiply", request, &res)
	if err != nil {
		log.Fatalln("arith error: ", err)
	}
	fmt.Printf("%d * %d = %d\n", request.A, request.B, res.Pro)

	err = conn.Call("Arith.Divide", request, &res)
	if err != nil {
		log.Fatalln("arith error: ", err)
	}
	fmt.Printf("%d / %d, quo is %d , rem is %d\n", request.A, request.B, res.Quo, res.Rem)

}
