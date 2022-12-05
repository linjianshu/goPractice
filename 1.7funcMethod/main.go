package main

import "fmt"

func main() {
	result, title := cost(2, 3, "li")
	fmt.Println(result, title) //105 high

	p := person{name: "lisi"}
	fmt.Println(p.String()) //the person name is lisi
}

func add(a, b, c int) int {
	return a + b + c
}

// 多返回值
func cost(a, b int, name string) (int, string) {
	if name == "li" {
		return a + b + 100, "high"
	} else {
		return a + b + 10, "low"
	}
}

func cost1(a, b int) {

}

type person struct {
	name string
}

// 类的方法
func (p person) String() string {
	return "the person name is " + p.name
}
