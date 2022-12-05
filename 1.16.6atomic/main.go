package main

import (
	"fmt"
	"runtime"
	"sync/atomic"
	"time"
)

// go中最重要的状态管理方式是通过通道间的沟通来完成的 在worker-pools中遇到过
// 但是还是有一些其他的方法来管理状态 使用sync/atomic包在多个go协程中进行原子计数
func main() {
	var ops uint64 = 0
	j := 0
	for i := 0; i < 50; i++ {
		go func() {
			for {
				atomic.AddUint64(&ops, 1)
				j++
				//允许其他go协程的执行 用于让出cpu时间片 出让当前goroutine的执行权限 调度器安排其他的等待任务执行 并在下次某个时候从该位置恢复执行
				runtime.Gosched()
			}
		}()
	}
	time.Sleep(time.Second)

	opsFinal := atomic.LoadUint64(&ops)
	fmt.Println("ops: ", opsFinal)
	fmt.Println("j: ", j)
}
