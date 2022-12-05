package main

import (
	"fmt"
	"time"
)

func main() {
	selectExample()
	selectExample1()
}

func selectExample() {
	c1 := make(chan string)
	c2 := make(chan string)
	go func() {
		time.Sleep(time.Second)
		c1 <- "one"
	}()

	go func() {
		time.Sleep(6 * time.Second)
		c2 <- "two"
	}()

	for i := 0; i < 2; i++ {
		select {
		case msg1 := <-c1:
			fmt.Println("received ", msg1)
		case msg2 := <-c2:
			fmt.Println("received ", msg2)
		}
	}
}

// 设置超时时间
func selectExample1() {
	ch := make(chan struct{})

	go func(ch chan struct{}) {
		//模拟耗时的操作
		time.Sleep(6 * time.Second)
		//finish task while send msg to ch
		ch <- struct{}{}
	}(ch)

	timeout := time.After(time.Second * 5)
	select {
	case <-timeout:
		fmt.Println("task timeout")
	case <-ch:
		fmt.Println("task finished")
	}
}

// worker goroutine需要一直循环处理信息 直到收到quit信号
func selectExample2() {
	msgCh := make(chan struct{})
	quitCh := make(chan struct{})

	for {
		select {
		case <-msgCh:
			//模拟循环耗时操作
			fmt.Println("do working...")
			time.Sleep(time.Second)
		case <-quitCh:
			//收到quit信号 跳出for循环 停止
			fmt.Println("finish...")
			return
		}

	}
}

// 在编译期间 防止channel被滥用
func foo(ch chan<- int) chan<- int {
	return ch
}
