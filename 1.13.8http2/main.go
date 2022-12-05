package main

import (
	"golang.org/x/net/http2"
	"log"
	"net/http"
)

func main() {
	var srv http.Server
	srv.Addr = ":8080"
	http.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		writer.Write([]byte("hello http2"))
	})
	http2.ConfigureServer(&srv, &http2.Server{})
	go func() {
		log.Fatal(srv.ListenAndServe())
	}()
	select {}
}
