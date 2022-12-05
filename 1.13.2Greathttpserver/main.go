package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"time"
)

func main() {
	startHttpServer()
}

func startHttpServer() {
	//请求多路复用器
	router := mux.NewRouter()

	//通过完整的path来匹配
	router.HandleFunc("/api/hello", helloTask)

	srv := &http.Server{
		Handler:      router,
		Addr:         ":8090",
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	log.Fatal(srv.ListenAndServe())
}

func helloTask(w http.ResponseWriter, request *http.Request) {
	fmt.Fprintf(w, "hello\n")
	w.Write([]byte("hello"))
}
