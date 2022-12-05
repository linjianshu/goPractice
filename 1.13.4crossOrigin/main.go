package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func main() {
	startHttpServer()
}

func helloTask(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "hello\n")
	w.Write([]byte("hello\n"))
}

func hiTask(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "hi\n")
	w.Write([]byte("hi"))
}

func startHttpServer() {
	router := mux.NewRouter()
	http_port := "8090"

	router.HandleFunc("/api/login", helloTask)

	err := http.ListenAndServe(":"+http_port, httpMiddleware(router))
	if err != nil {
		log.Println(err)
	}
}

// 跨域处理的中间件HTTP服务所有的请求都会经过这里处理
func httpMiddleware(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		//处理跨域问题
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
		w.Header().Set("Access-Control-Allow-Headers", "Origin, Authorization, Origin, X-Requested-With, Content-Type, Accept,common")

		//先做中间件的前件
		hiTask(w, r)
		//处理完后处理现在这个
		h.ServeHTTP(w, r)
		if r.Method == "OPTIONS" {
			return
		}
	})
}
