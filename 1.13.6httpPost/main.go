package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
)

func main() {
	//使用body体进行传参
	resp, err := http.PostForm("http://124.220.207.169:9999/fuckVote", url.Values{"name": {"雷俊"}})
	if err != nil {
		log.Fatal(err)
		return
	}
	defer resp.Body.Close()
	all, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
		return
	}
	fmt.Println(string(all))
}
