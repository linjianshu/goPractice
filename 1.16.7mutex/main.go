package main

import (
	"fmt"
	"math/rand"
	"runtime"
	"sync"
	"sync/atomic"
	"time"
)

// golang中的锁是通过CAS原子操作实现的
/*
type mutex struct{
state int32
sema uint32
}
state表示锁当前状态 每个位都有意义 位图 零值表示未上锁
sema用作信号量 通过PV操作从等待队列中阻塞/唤醒 goroutine 等待锁的goroutine会挂到等待队列中 并且陷入睡眠不被调度 unlock锁时才唤醒

使用原子操作可以跨多个协程管理简单的计数器状态 对于更复杂的状态 可以使用互斥锁安全的跨多个协程访问数据
sync.Mutex不区分读写锁 只有lock和unlock之间才会导致阻塞的情况 如果在一个地方调用lock 在另一个地方不调用lock而是直接修改或访问共享数据
这对于mutex类型来说是允许的 因为mutex不会和goroutine关联 如果想要区分读写锁 可以使用rwMutex类型

在lock和unlock之间的代码段被称为资源的临界区critical section 在这一区间的代码是严格被lock保护的 是线程安全的 任何一个时间点都只能有一个goroutine
执行这段区间的代码

尽量减少锁的持有时间,毕竟使用锁是有代价的 通过减少锁的持有时间来减轻这个代价:细化锁的粒度 通过细化锁的粒度来减少锁的持有时间以及避免在持有锁操作时
做各种耗时的操作
不要在持有锁时IO操作 尽量通过持有锁来保护IO操作需要的资源而不是IO操作本身
*/
func main() {
	//mutexDemo()
	//mutexDemo2()
	RWMutexDemo()
}

func mutexDemo() {
	var state = make(map[int]int)
	var mutex = &sync.Mutex{}
	var readOps uint64
	var writeOps uint64

	for i := 0; i < 100; i++ {
		go func() {
			total := 0
			for {
				key := rand.Intn(5)
				//以独占方式访问状态
				mutex.Lock()
				total += state[key]
				mutex.Unlock()
				atomic.AddUint64(&readOps, 1)
				runtime.Gosched()
			}
		}()
	}

	for w := 0; w < 10; w++ {
		go func() {
			for {
				key := rand.Intn(5)
				val := rand.Intn(100)
				mutex.Lock()
				state[key] = val
				mutex.Unlock()
				atomic.AddUint64(&writeOps, 1)
				time.Sleep(time.Millisecond)
			}
		}()
	}

	time.Sleep(time.Second)
	readOpsFinal := atomic.LoadUint64(&readOps)
	writeOpsFinal := atomic.LoadUint64(&writeOps)
	fmt.Println("readOps: ", readOpsFinal)
	fmt.Println("writeOp: ", writeOpsFinal)
	//最终锁定状态为 state说明它如何结束
	mutex.Lock()
	fmt.Println("state : ", state)
	mutex.Unlock()
}

func mutexDemo2() {
	var wg = sync.WaitGroup{}
	var lock = sync.Mutex{}
	var sum int
	var add = func() {
		lock.Lock()
		sum += 1
		defer wg.Done()
		defer lock.Unlock()
	}

	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go add()
	}
	wg.Wait()
	fmt.Println("sum = ", sum)
}

func RWMutexDemo() {
	var m = sync.RWMutex{}
	var read = func(i int) {
		fmt.Println(i, "开始读...")

		m.RLock()
		fmt.Println(i, "正在读...")
		time.Sleep(time.Second)
		m.RUnlock()

		fmt.Println(i, "读结束...")
	}

	var write = func(i int) {
		fmt.Println(i, "开始写...")
		m.Lock()
		fmt.Println(i, "正在写...")
		time.Sleep(time.Second)
		m.Unlock()
		fmt.Println(i, "写完毕...")
	}

	go read(1)
	go read(2)
	time.Sleep(3 * time.Second)
	fmt.Println("-----------------------------")

	go write(1)
	go read(2)
	go write(3)
	time.Sleep(4 * time.Second)
}
