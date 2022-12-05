package main

import "time"

func main() {
	ch := make(chan int, 3)
	ch <- 1
	ch <- 2
	ch <- 3

	//阻塞
	//ch <- 4
	//通过len函数可以获得chan中的元素个数 通过cap函数可以得到chan的缓存长度

	c := make(chan int)
	go func() {
		//模拟在干一些耗时的操作
		time.Sleep(2 * time.Second)
		//发送信号给c告知已经完成任务
		c <- 1
	}()

	time.Sleep(time.Second)
	//阻塞
	<-c
	//主goroutine会阻塞 直到子协程干完活往通道里发送信号
}
