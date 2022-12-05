package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	testPanic()
}

/*
defer
defer语句讲一个函数放入一个列表 用栈表示其实更准确 该列表的函数在环绕defer的函数返回时会被执行 defer通常用于简化函数的各种各样的清理动作 例如关闭文件 解锁等释放资源的动作

panic是内建的停止控制流的函数 相当于其他编程语言的抛异常操作 若函数F调用了panic F的执行会被停止 在F中panic前面定义的defer操作都会被执行 然后F函数返回
对于调用者来说 调用F的行为就想调用panic (如果F函数内部没有把panic覆盖掉) 如果都没有捕获该panic 相当于一层层panic(运行恐慌) 程序将会crash panic可以直接调用
也可以在程序运行错误时调用 例如数组越界

recover是一个从panic恢复的内建函数 recover只有在defer的函数里才能发挥真正的作用 如果是正常的情况 (没有发生panic) 调用recover将会返回nil且没有任何影响 如果当前的goroutine
panic了 recover的调用会捕获到panic子 并且回复正常运行

go语言追求简洁优雅 所以go语言不支持传统的try catch finally 异常,go语言的设计者认为,将异常与控制结构混在一起容易使代码变得混乱

在go语言中,使用多值返回来返回错误 不要用异常代替错误 更不要用来控制流程 在个别情况下 即遇到真正的异常情况时 (比如除数为0时) 才使用go中引入的Exception处理: defer panic recover

go没有异常机制 但有panic和recover模式来处理错误 panic可以在任何地方被引发 但recover只有在defer调用的函数中有效
*/

func CopyFile(dstName, srcName string) (written int64, err error) {
	src, err := os.Open(srcName)
	if err != nil {
		return
	}

	create, err := os.Create(dstName)
	if err != nil {
		return
	}

	written, err = io.Copy(create, src)
	create.Close()
	src.Close()
	//立刻关闭
	return
}

func CopyFileDefer(dstName, srcName string) (written int64, err error) {
	src, err := os.Open(srcName)
	if err != nil {
		return
	}
	defer src.Close()

	dst, err := os.Create(dstName)
	if err != nil {
		return
	}
	defer dst.Close()

	return io.Copy(dst, src)
}

func testPanic() {
	//先声明defer 捕获panic异常
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("捕获到了panic产生的异常: ", err)
			fmt.Println("捕获到了panic的一场了 , recover 恢复回来了 ")
		}
	}()

	//注意这个()就是调用的匿名函数
	//不写会报 :expression expression in defer must be function call
	/*
		panic一般会导致程序挂掉(除非recover) 然后再go运行时输出调用栈 但即使函数执行时引起panic了
		函数不在往下运行 并不是立刻向上传递panic 而是移到了defer处 等defer的东西都跑完了 panic再向上传递
		所以这时候defer有点类似 try-catch-finally中的finally panic函数就是这么简单
	*/
	panic("抛出一个异常了 defer会通过recover捕获这个异常 处理后续程序正常运行")
	fmt.Println("这里不会执行了")
}
