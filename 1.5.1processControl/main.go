package main

import (
	"fmt"
	"time"
)

func main() {
	//1.定义变量
	a := 10
	b := 3.14
	c, d, e := 1, "hello", 3.14
	//unicode 码点 int32
	f := 'a'
	fmt.Printf("%d %T\n", f, f) //97 int32
	g := "xyz"
	h := false
	fmt.Printf("a = %v , b = %v , c = %v , d = %v , e = %v , f = %v , g = %v , h = %v \n", a, b, c, d, e, f, g, h) //a = 10 , b = 3.14 , c = 1 , d = hello , e = 3.14 , f = 97 , g = xyz , h = false

	var as int
	var bs float32
	var isTrue bool
	var str string
	fmt.Printf("as = %v , bs = %v , isTrue = %v , str = %v \n", as, bs, isTrue, str) //as = 0 , bs = 0 , isTrue = false , str =

	var arrInt = []int{2, 4, 6}
	var arrString = []string{"2", "4", "6"}
	fmt.Printf("arrInt = %v , arrString = %v \n", arrInt, arrString) //arrInt = [2 4 6] , arrString = [2 4 6]s

	//只是声明 还没有初始化 不会报错 可以对nil的map取值 但是没有意义 赋值就报错了!!!
	var g1 map[int]bool
	//会报错
	//g1[4] = false
	g2 := map[int]bool{}
	g3 := make(map[int]bool)
	//make map 指定的是容量 没东西的时候是0
	g4 := make(map[int]bool, 10)
	fmt.Printf("g1[0] = %v , g2[0] = %v , g3[0] = %v , g4[0] = %v , len(g4) = %v\n", g1[0], g2[0], g3[0], g4[0], len(g4)) //g1[0] = false , g2[0] = false , g3[0] = false , g4[0] = false , len(g4) = 0

	var m1 = map[int]string{1: "boy", 2: "girl"}
	m2 := map[int]string{1: "boy", 2: "girl"}
	fmt.Println(m1, m2) //map[1:boy 2:girl] map[1:boy 2:girl]

	//2.for循环
	i := 1
	for i <= 3 {
		//1
		//2
		//3
		fmt.Println(i)
		i++
	}

	for j := 7; j <= 9; j++ {
		//7
		//8
		//9
		fmt.Println(j)
	}

	for {
		fmt.Println("loop") //loop
		break
	}

	for n := 0; n <= 5; n++ {
		if n%2 == 0 {
			continue
		}
		fmt.Println(n)
		//1
		//3
		//5
	}

	//3.if/else
	if 7%2 == 0 {
		fmt.Println("7 is even")
	} else {
		fmt.Println("7 is odd") //7 is odd 奇数
	}

	if 8%4 == 0 {
		fmt.Println("8 is divisible by 4") //8 is divisible by 4
	}

	if num := 9; num < 0 {
		fmt.Println(num, "is negative")
	} else if num < 10 {
		fmt.Println(num, "has 1 digit") //9 has 1 digit 9是一位数
	} else {
		fmt.Println(num, "has multiple digits")
	}

	//4.switch
	i = 2
	fmt.Println("Write", i, "as")
	switch i {
	case 1:
		fmt.Println("one")
	case 2:
		fmt.Println("two") //Write 2 as two
	case 3:
		fmt.Println("three")
	}

	switch time.Now().Weekday() {
	case time.Saturday, time.Sunday:
		fmt.Println("It's the weekend")
	default:
		fmt.Println("It's a weekday") //It's a weekday
	}

	t := time.Now()
	switch {
	case t.Hour() < 12:
		fmt.Println("It's before noon")
	default:
		fmt.Println("It's after noon") //It's after noon
	}

	whatAmI := func(i interface{}) {
		switch t := i.(type) {
		case bool:
			fmt.Println("I'm a bool")
		case int:
			fmt.Println("I'm an int")
		default:
			fmt.Printf("Don't know type %T\n", t)
		}
	}
	whatAmI(true)  //I'm a bool
	whatAmI(1)     //I'm an int
	whatAmI("hey") //Don't know type string

}
