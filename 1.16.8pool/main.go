package main

import (
	"fmt"
	"strconv"
	"sync"
	"time"
)

func main() {
	jobWorkers()
	syncPoolDemo()
}

type Person struct {
	Name string
}

func initPool() {
	pool = &sync.Pool{
		New: func() any {
			return new(Person)
		},
	}
}

var pool *sync.Pool

func syncPoolDemo() {
	initPool()

	startTime := time.Now()
	//需要新的结构体时 先尝试去pool中取 而不是重新生成 这样重复10000次节省大量的时间
	//这样简单的操作节约了时间 也节约了各方面的资源
	//最重要的是 它可以有效减少gc cpu 和gc pause的时间
	for i := 0; i < 10000; i++ {
		one := pool.Get().(*Person)
		one.Name = "girl" + strconv.Itoa(i)
		//fmt.Printf("one.Name = %s \n", one.Name)
		//使用后提交实例
		pool.Put(one)
	}

	//现在 同一实例可被另一个例程使用 而无序再次分配它
	fmt.Println("花费时间1:", time.Since(startTime))

	startTime = time.Now()
	for i := 0; i < 10000; i++ {
		p := Person{Name: "girl" + strconv.Itoa(i)}
		pool.Put(p)
	}
	fmt.Println("花费时间2: ", time.Since(startTime))
}

func worker(id int, jobs <-chan int, results chan<- int) {
	for job := range jobs {
		fmt.Println("worker ", id, "processing job ", job)
		//模拟耗时操作
		time.Sleep(time.Second)
		//计算结果存进去
		results <- job * 2
	}
}

// 这是将要在多个并发实例中支持的任务
// 这些执行者将冲 jobs通道接收任务 并且通过results发送相应的结果
// 将让每个任务间隔1s来模仿一个耗时的任务
func jobWorkers() {
	jobs := make(chan int, 100)
	results := make(chan int, 100)

	//启动3个worker 初始是阻塞的 因为还没有传递任务
	for w := 0; w < 3; w++ {
		go worker(w, jobs, results)
	}

	//发送9个jobs 然后关闭这些通道来表示这些就是所有的任务
	for j := 0; j < 9; j++ {
		jobs <- j
	}
	close(jobs)

	//最后 收集所有这些任务的返回值
	for res := 0; res < 9; res++ {
		fmt.Println(<-results)
	}
}
