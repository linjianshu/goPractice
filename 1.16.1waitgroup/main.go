package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	//testAsync()
	//testAsyncSleep()
	//testAsyncWait()
	testWaitGroupTimeOut()
}

func testWaitGroupTimeOut() {
	var w = sync.WaitGroup{}
	var ch = make(chan bool)
	w.Add(2)

	go func() {
		time.Sleep(time.Second * 2)
		fmt.Println("等待2s")
		w.Done()
	}()

	go func() {
		time.Sleep(time.Second * 6)
		fmt.Println("等待6s")
		w.Done()
	}()

	go func() {
		w.Wait()
		ch <- false
	}()

	select {
	case <-time.After(time.Second * 5):
		fmt.Println("超时了")
	case <-ch:
		fmt.Println("工作顺利结束")
	}
}

// 使用waitGroup进行等待
func testAsyncWait() {
	fmt.Println("aaa")
	var waitGroup sync.WaitGroup
	waitGroup.Add(10)
	for i := 0; i < 10; i++ {
		go func(index int) {
			fmt.Println(index)
			waitGroup.Done()
		}(i)
	}
	waitGroup.Wait()
	fmt.Println("ccc")
}

func testAsyncSleep() {
	fmt.Printf("aaa \n")
	go asyncFunc("bbb2 \n")
	fmt.Printf("ccc \n")
	time.Sleep(1 * time.Second)

}

func testAsync() {
	fmt.Printf("aaa\n")
	//开启一个协程 但是主线程没有等待 主线程退出 子协程也退出
	go asyncFunc("bbb1 \n")
	fmt.Printf("ccc \n")
}

func asyncFunc(str string) {
	fmt.Println(str)
}
