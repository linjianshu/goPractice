package main

import (
	"html/template"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", sayHello)
	err := http.ListenAndServe(":8099", nil)
	if err != nil {
		log.Printf("http server start failed, err :%v\n", err)
	}
}

func sayHello(w http.ResponseWriter, req *http.Request) {
	template := template.Must(template.ParseFiles("index.html"))
	name := "mmm"
	user := User{
		Name:   name,
		Gender: "男",
		Age:    23,
	}
	m := map[string]interface{}{
		"name":   name,
		"gender": "男",
		"age":    24,
	}
	carList := []string{"汽车", "火车", "货车"}
	err := template.Execute(w, map[string]interface{}{
		"m":       m,
		"user":    user,
		"carList": carList,
	})
	if err != nil {
		log.Printf("render template failed, err %v", err)
		return
	}
}

type User struct {
	Name   string
	Gender string
	Age    int
}
