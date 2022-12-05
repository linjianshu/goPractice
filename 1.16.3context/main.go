package main

import (
	"context"
	"fmt"
	"net/http"
	"time"
)

/*
在golang中创建的一个新的协程并不会像C语言创建一个线程一样类似的pid, 这样就导致其不能从外部停掉某个线程, 所以我们就得让他自己结束
当然也可以采用channel+select的方式来解决这个问题 不过场景很复杂时 我们就需要花费很大的经历去维护channel与这些协程的关系
这就导致了我们的并发代码变得很难维护和管理
context的产生 正是因为协程的管理问题 goland官方冲1.7之后引入了context 用来专门管理协程之间的关系
google的解决方法是context机制 相互调用的goroutine之间通过传递context变量来保持关联 这样在不用暴露各goroutine
内部实现细节的前提下 有效地控制各goroutine的运行 通过传递context就可以追踪暴露goroutine调用树 并在这些调用树之间传递通知和元数据
虽然goroutine之间是平行的 没有继承关系 但是context设计成包含父子关系的形式 这样可以更好的描述goroutine调用之间的树形关系
context包的核心就是context接口 定义如下:
Deadline()
返回一个超时时间 到了该超时时间 该context所代表的工作将被取消继续执行 goroutine获得了超时时间后 可以对某些io操作设定超时时间

Done()
返回一个通道channel 当context被撤销或过期时 该通道被关闭 它是一个表示context是否已关闭的信号

Err()
当done通道关闭后 err方法返回值为context被撤的原因

value()
可以让goroutine共享一些数据 当然获得数据是协程安全的 但使用这些数据时要注意同步 比如返回了map 这个map的读写需要加锁
注意:context包里的方法是线程安全的 可以被多个线程使用
context没有提供方法来设置其值和过期时间 也没有提供方法直接将其自身撤销 也就是说 context不能改变和撤销其自身
*/
func main() {
	ExampleWithCancel()
	ExampleWithDeadline()
	ExampleWithTimeout()
	ExampleWithValue()

	//context给多个子协程发送控制指令
	ctx, cancel := context.WithCancel(context.Background())
	go watch(ctx, "[任务1]")
	go watch(ctx, "[任务2]")
	go watch(ctx, "[任务3]")
	time.Sleep(10 * time.Second)
	fmt.Println("可以了 通知任务停止")
	cancel()
	time.Sleep(5 * time.Second)

	http.HandleFunc("/hello", hello)
	http.ListenAndServe(":8090", nil)
}

// ExampleWithCancel 此示例演示了如何使用可取消上下文来防止Goroutine泄漏
// by gen将返回而不会发生泄漏
func ExampleWithCancel() {
	//gen在单独的goroutine中生成整数
	//然后将他们发送到返回的channel中
	//gen的调用者需要取消一次上下文
	//他们完成了对生成的整数的使用而不泄漏
	//内部goroutine有gen开始
	gen := func(ctx context.Context) <-chan int {
		dst := make(chan int)
		n := 1
		go func() {
			for {
				select {
				case <-ctx.Done():
					return
				case dst <- n:
					n++
				}
			}
		}()
		return dst
	}

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel() //当我们使用完整数后取消

	for n := range gen(ctx) {
		fmt.Println(n)
		if n == 5 {
			break
		}
	}
}

// ExampleWithDeadline 此示例传递具有任意截止日期的上下文以告知阻塞
// 表示应该立即放弃工作的功能
func ExampleWithDeadline() {
	d := time.Now().Add(50 * time.Millisecond)
	ctx, cancel := context.WithDeadline(context.Background(), d)
	//即使ctx将会国企 但是最好还是将其调用
	//在任何情况下都具有取消功能 否则可能会使上下文及其父对象的生存时间超出了必要
	defer cancel()

	select {
	case <-time.After(time.Second):
		//模拟耗时操作 需要这么久
		fmt.Println("overslept")
	case <-ctx.Done():
		//模拟达到设定的时间 中断了就不用干了
		fmt.Println(ctx.Err())
	}
}

// ExampleWithTimeout 此示例传递带有超时的上下文 以告知阻塞函数
// 它应在超时后立即放弃工作
func ExampleWithTimeout() {
	ctx, cancel := context.WithTimeout(context.Background(), 50*time.Millisecond)
	//cancel即使不主动调用 也不影响资源的最终释放
	//但是应该提前主动调用 可以尽快的释放 避免等待过期时间之间的浪费
	defer cancel()

	select {
	case <-time.After(time.Second):
		//模拟耗时操作
		fmt.Println("overslept")
	case <-ctx.Done():
		//模拟超时了 中断了就不用干了
		fmt.Println(ctx.Err())
	}
}

// ExampleWithValue 此示例演示如何将值传递到上下文以及如何检索他(如果存在)
func ExampleWithValue() {
	type favContextKey string
	f := func(ctx context.Context, k favContextKey) {
		if v := ctx.Value(k); v != nil {
			fmt.Println("found value: ", v)
			return
		}
		fmt.Println("key not found: ", k)
	}

	key := favContextKey("language")
	ctx := context.WithValue(context.Background(), key, "Go")
	f(ctx, key)
	f(ctx, favContextKey("color"))
}

func watch(ctx context.Context, name string) {
	for {
		select {
		case <-ctx.Done():
			fmt.Println(name, "任务退出 停止了...")
			return
		default:
			fmt.Println(name, "goroutine 任务中...")
			time.Sleep(2 * time.Second)
		}

	}
}

func hello(w http.ResponseWriter, request *http.Request) {
	ctx := request.Context()
	fmt.Println("server: 处理开始")
	defer fmt.Println("server: 处理结束")

	select {
	case <-time.After(time.Second * 10):
		//模拟耗时操作 例如数据库查询
		fmt.Fprintf(w, "hello\n")
	case <-ctx.Done():
		//模拟可能导致的 用户关闭界面
		err := ctx.Err()
		fmt.Println("server: ", err)
		internalError := http.StatusInternalServerError
		http.Error(w, err.Error(), internalError)
	}
}
