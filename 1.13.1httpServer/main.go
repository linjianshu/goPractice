package main

import (
	"fmt"
	"net/http"
)

func main() {
	//定义API要访问的处理函数
	http.HandleFunc("/hello", hello)
	//开启HTTP服务 :是指任何IP
	http.ListenAndServe(":8090", nil)
	//也可以指定IP
	//http.ListenAndServeTLS("127.0.0.1:8090", nil)
}

func hello(w http.ResponseWriter, r *http.Request) {
	//两种输出方式
	fmt.Fprintf(w, "hello http\n")
	w.Write([]byte("hello"))
}
