package main

import (
	"bytes"
	"crypto/tls"
	"fmt"
	"io"
	"log"
	"net"
	"net/http"
	"net/url"
	"strings"
	"time"
)

func main() {
	data := url.Values{}
	data.Add("name", "林健树")
	urls := "http://124.220.207.169:9999/fuckVote"
	body, err := myHttpRequest(urls, data)
	if err != nil {
		log.Println(err)
		return
	}
	fmt.Println(string(body))
}

var transport *http.Transport
var client *http.Client

func myHttpRequest(url string, params url.Values) (body []byte, err error) {
	if transport == nil {
		if strings.Contains(url, "https") {
			transport = &http.Transport{
				TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
				//禁止长连接
				DisableKeepAlives: true,
			}
		} else {
			transport = &http.Transport{
				DialContext: (&net.Dialer{Timeout: 10 * time.Second}).DialContext,
				//tls握手超时时间
				TLSHandshakeTimeout: 10 * time.Second,
			}
		}
		client = &http.Client{
			Transport: transport,
			Timeout:   time.Second * 10,
		}
	}
	//bytes.NewBufferString  params.Encode
	request, err := http.NewRequest("POST", url, bytes.NewBufferString(params.Encode()))
	if err != nil {
		log.Printf("Error Occured. %+v", err)
		return nil, err
	}
	request.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	do, err := client.Do(request)
	if err != nil {
		log.Println("client.Post error")
		log.Println(err)
		log.Println(url)
		return nil, err
	}
	defer do.Body.Close()
	all, err := io.ReadAll(do.Body)
	if err != nil {
		log.Println("client.Post read error")
		log.Println(err)
		return nil, err
	}
	return all, err
}
