package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"os"
	"os/exec"
	"path/filepath"
	"time"
)

func main() {
	startHttpServer()
}

func startHttpServer() {
	router := mux.NewRouter()
	router.HandleFunc("/api/hello", helloTask)
	curdir, err := GetCurrentPath()
	if err != nil {
		fmt.Println(err)
		return
	}
	pthSep := string(os.PathSeparator)
	filePath := curdir + pthSep + "file"
	router.PathPrefix("/file/").Handler(http.StripPrefix("/file/", http.FileServer(http.Dir(filePath))))

	srv := &http.Server{
		Handler:      router,
		Addr:         ":8090",
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	log.Fatal(srv.ListenAndServe())
}

func GetCurrentPath() (dir string, err error) {
	path, err := exec.LookPath(os.Args[0])
	if err != nil {
		log.Printf("exec.LookPath(%s),err: %s\n", os.Args[0], err)
		return "", err
	}
	abs, err := filepath.Abs(path)
	if err != nil {
		log.Printf("filepath.Abs(%s) , err: %s\n", path, err)
		return "", err
	}
	dir = filepath.Dir(abs)
	return dir, nil

}

func helloTask(writer http.ResponseWriter, request *http.Request) {
	fmt.Fprintf(writer, "hello\n")
	writer.Write([]byte("hello"))
}
