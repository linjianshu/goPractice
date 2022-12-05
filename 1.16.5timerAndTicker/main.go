package main

import (
	"fmt"
	"time"
)

/*
timer 和 ticker 都是用于计时
使用timer定时器 超时后需要重置 才能继续触发
ticker只要定义完成 从此刻开始计时 不需要其他任何的操作 每隔固定时间都会触发
*/
func main() {
	timerExample()
	tickerExample()
}

func timerExample() {
	timer1 := time.NewTimer(2 * time.Second)
	<-timer1.C
	fmt.Println("timer1 fired")

	timer2 := time.NewTimer(time.Second)
	//go的子协程输出不了 来不及执行就被stop掉了
	go func() {
		<-timer2.C
		fmt.Println("timer2 fired")
	}()

	stop2 := timer2.Stop()
	if stop2 {
		fmt.Println("timer2 stopped")
	}

	//重置才能重新触发
	//timer2.Reset(time.Second)
	time.Sleep(2 * time.Second)
}

func tickerExample() {
	ticker := time.NewTicker(500 * time.Millisecond)
	done := make(chan bool)
	go func() {
		for {
			select {
			case <-done:
				return
			case t := <-ticker.C:
				fmt.Println("Tick at ", t)
			}
		}
	}()
	time.Sleep(1600 * time.Millisecond)
	ticker.Stop()
	done <- true
	fmt.Println("ticker stopped")
}
