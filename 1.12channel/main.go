package main

import (
	"fmt"
)

func main() {
	fmt.Println("!...主协程开始...!")
	ch := make(chan int)
	go func() {
		fmt.Println("开始子协程")
		ch <- 0
		fmt.Println("退出子协程")
	}()
	fmt.Println("等待...")
	<-ch
	fmt.Println("阻塞完毕 继续执行")
	fmt.Println("!...主协程结束...!")
}
