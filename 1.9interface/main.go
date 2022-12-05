package main

import (
	"fmt"
	"time"
)

func main() {
	testNilInterface()
	testInterface()
}

/*
interface 是go语言的基础特性之一 可以理解为一种类型的规范或者约定 他和java和C#不太一样 不需要显式说明实现了某个接口
它没有继承子类或implements关键字 只是通过约定的形式隐式地实现interface中的方法即可 因此 golang中的interface让编码更加灵活以扩展

什么情况下使用interface呢
当我们给系统增加一个功能时 不是通过修改代码 而是通过增添代码来完成的 这就是开闭原则的核心思想
所以要想满足上面的要求 需要interface来提供一层抽象的接口
作为interface数据类型 他存在的意义是什么呢 实际上就是为了满足一些面向对象的编程思想 目标就是高内聚低耦合
go中严格来说没有多态 但可以利用接口进行 对于实现了同一接口的两种对象 可以进行类似的向上转型
并且在此时可以对方法进行多态路由转发
*/

type Student struct {
	Name string
}

func testNilInterface() {
	var v interface{}
	v = 12
	v = "ab"
	v = 12.22
	v = Student{Name: "abc"}

	//类型推断 ok判误
	if _, ok := v.(int); ok {
		fmt.Printf("is int type\n")
	} else if _, ok := v.(string); ok {
		fmt.Printf("is string type\n")
	} else if _, ok := v.(Student); ok {
		fmt.Printf("is struct type\n")
	} else {
		fmt.Printf("unknown type\n")
	}

	//switch类型推断
	switch v.(type) {
	case int:
		fmt.Printf("%s", v)
	case string:
		fmt.Printf("%s", v)
	case Student:
		fmt.Printf("%s", v)
	case bool:
		fmt.Printf("%s", v)
	case float32:
		fmt.Printf("%1.2f", v)
	case []byte:
		fmt.Printf("%s", string(v.([]byte)))
	case time.Time:
		fmt.Printf("%s", v)
	default:
		fmt.Printf("unkonwn type\n")
	}
}

/*
interface是一种具有一组方法的类型 这些方法定义了interface的行为 interface{}会占用两个
字长的存储空间  一个是自身的methods数据 一个是指向其存储值的指针 也就是interface变量存储的值
一个类型如果实现了一个interface的所有方法 就说该类型实现了这个interface 空的interface没有方法 所以可以认为所有的类型都实现了interface{}
如果定义一个函数参数是interface{} 类型 这个函数应该可以接收任何类型作为它的参数 跟java和c++等其他语言的多态类似
*/

type Animal interface {
	Speak() string
}

type Dog struct {
}

func (d Dog) Speak() string {
	fmt.Printf("woof\n")
	return "Woof!"
}

type Cat struct {
}

func (c *Cat) Speak() string {
	fmt.Printf("Meow!\n")
	return "Meow"
}

func testInterface() {
	dog := Dog{}
	dog.Speak()

	cat := Cat{}
	cat.Speak()

	//用dog实现接口
	animal := Animal(dog)
	animal.Speak()

	p := &Person{}
	p.Run()
	p.Sleep()
}

type Person struct {
}

func (p *Person) Run() {
	fmt.Println("进行了奔驰")
}

func (p *Person) Sleep() {
	fmt.Println("进行了睡觉")
}
