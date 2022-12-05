package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("!...主协程开始...!")
	go BName()

	go Bid()
	time.Sleep(3500 * time.Millisecond)
	fmt.Println("暂时结束")

	go f("direct")
	go f("goroutine")
	go func(msg string) {
		fmt.Println(msg)
	}("going")
	time.Sleep(time.Second)
	fmt.Println("done")
}

func BName() {
	arr1 := [4]string{"aa", "bb", "cc", "dd"}
	for t1 := 0; t1 <= 3; t1++ {
		time.Sleep(150 * time.Millisecond)
		fmt.Printf("%s\n", arr1[t1])
	}
}

func Bid() {
	arr2 := [4]int{11, 22, 33, 44}
	for t2 := 0; t2 <= 3; t2++ {
		time.Sleep(150 * time.Millisecond)
		fmt.Printf("%d\n", arr2[t2])
	}
}

func f(from string) {
	for i := 0; i < 3; i++ {
		fmt.Println(from, ":", i)
	}
}
