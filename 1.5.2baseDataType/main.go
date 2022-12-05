package main

import (
	"fmt"
	"strconv"
	"sync"
	"time"
)

func main() {
	arraysExample()
	array()
	multiArray()
	slice()
	SliceExample()
	sliceRecycle()
	autoEnlarge()
	mapTest()
	syncMapTest()
	cover()
	mapExample()
	rangeExample()
	structExample()
}

func arraysExample() {
	var a [5]int
	fmt.Println("emp:", a) //emp: [0 0 0 0 0]
	a[4] = 100
	fmt.Println("set:", a)      //set: [0 0 0 0 100]s
	fmt.Println("get:", a[4])   //get: 100
	fmt.Println("len:", len(a)) //len: 5

	var twoD [2][3]int
	for i := 0; i < 2; i++ {
		for j := 0; j < 3; j++ {
			twoD[i][j] = i + j
		}
	}

	fmt.Println("2d:", twoD) //2d: [[0 1 2] [1 2 3]]
}

func array() {
	var a [3]int
	fmt.Println(a[0])        //0
	fmt.Println(a[len(a)-1]) //0

	for i, v := range a {
		fmt.Printf("%d %d\n", i, v)
		// 0 0
		// 1 0
		// 2 0
	}

	for _, v := range a {
		fmt.Printf("%d\n", v)
		// 0
		// 0
		// 0
	}

	var team [3]string
	team[0] = "aa"
	team[1] = "bb"
	team[2] = "cc"

	for k, v := range team {
		fmt.Println(k, v)
		// 0 "aa"
		// 1 "bb"
		// 2 "cc"
	}

	g := [3]int{12, 78, 50}
	fmt.Println(g) //[12 78 50]

	//初始化了3个长度 默认为0
	b := [3]int{12} //[12 0 0]
	fmt.Println(b)

	c := [...]int{12, 78, 50}
	fmt.Println(c) //[12 78 50]

	//%.2f 默认保留小数点后2位
	d := [...]float64{22.7, 23.8, 56, 68, 78}
	for i := 0; i < len(d); i++ {
		fmt.Printf("%d th element of a is %.2f\n", i, d[i])
		//0 is 22.70
		//1 is 23.80
		//2 is 56.00
		//3 is 68.00
		//4 is 78.00
	}

	for i, v := range d {
		fmt.Printf("%d the element of a is %.2f\n", i, v)
		//0 is 22.70
		//1 is 23.80
		//2 is 56.00
		//3 is 68.00
		//4 is 78.00
	}

	for _, v := range d {
		fmt.Printf("%.2f\n", v)
		//22.70
		//23.80
		//56.00
		//68.00
		//78.00
	}
}

func multiArray() {
	a := [3][2]string{
		{"a1", "a2"},
		{"b1", "b2"},
		{"c1", "c2"},
	}
	printArray(a)
	// a1a2
	// b1b2
	// c1c2

	var b [3][2]string
	b[0][0] = "a1"
	b[0][1] = "a2"
	b[1][0] = "b1"
	b[1][1] = "b2"
	b[2][0] = "c1"
	b[2][1] = "c2"

	printArray(b)
	// a1a2
	// b1b2
	// c1c2
}

func printArray(a [3][2]string) {
	for _, v1 := range a {
		for _, v2 := range v1 {
			fmt.Printf("%s", v2)
		}
		fmt.Printf("\n")
	}
}

func slice() {
	c := []int{6, 7, 8}
	c[1] = 70
	fmt.Println(c) //[6 70 8]

	numa := [3]int{78, 79, 80}
	//数组=>切片
	nums1 := numa[:]
	for _, v1 := range nums1 {
		fmt.Printf("%d", v1) //787980
	}

	si := make([]int, 5, 5)
	for _, v1 := range si {
		fmt.Printf("%d ", v1) //0 0 0 0 0
	}

	cars := []string{"a", "b", "c"}
	fmt.Println(cars) // [a b c]
	cars = append(cars, "d")
	fmt.Println(cars) // [a b c d]
	cars = append(cars, "e", "f", "g", "h")
	fmt.Println(cars) // [a b c d e f g h]

	veggies := []string{"potatoes", "tomatoes", "brinjal"}
	fruits := []string{"oranges", "apples"}
	//切片合并
	food := append(veggies, fruits...)
	for _, v1 := range food {
		fmt.Printf("%s ", v1) //"potatoes" "tomatoes" "brinjal" "oranges" "apples"
	}

	index := 1
	//切片删除
	food = append(food[:index], food[index+1:]...)
	for _, v1 := range food {
		fmt.Printf("%s ", v1) //"potatoes" "brinjal" "oranges" "apples"
	}

	//清空 重新赋值
	food = append([]string{})
	fmt.Println(food) //[]

	pls := [][]string{
		{"potatoes", "tomatoes"},
		{"brinjal"},
		{"oranges", "apples"},
	}
	for _, v1 := range pls {
		for _, v2 := range v1 {
			fmt.Printf("%s ", v2)
		}
		fmt.Println()
	}
	//potatoes tomatoes
	//brinjal
	//oranges apples
}

func SliceExample() {
	s := make([]string, 3)
	fmt.Println("emp:", s[0], s[1], s[2]) //[三个空]

	s[0] = "a"
	s[1] = "b"
	s[2] = "c"

	fmt.Println("set:", s)      //[a b c]
	fmt.Println("get:", s[2])   //c
	fmt.Println("len:", len(s)) //3

	s = append(s, "d")
	s = append(s, "e", "f")
	fmt.Println("apd:", s) //[a b c d e f]

	c := make([]string, len(s))
	//拷贝操作
	copy(c, s)
	fmt.Println("cpy:", c) //[a b c d e f]
	l := s[2:5]
	fmt.Println("sl1:", l) //[c d e]
	l = s[:5]
	fmt.Println("sl2:", l) // [a b c d e]
	l = s[2:]
	fmt.Println("sl3:", l) //[c d e f]

	t := []string{"g", "h", "i"}
	fmt.Println("dcl", t) //[g h i]
	twoD := make([][]int, 3)
	for i := 0; i < 3; i++ {
		innerLen := i + 1
		twoD[i] = make([]int, innerLen)
		for j := 0; j < innerLen; j++ {
			twoD[i][j] = i + j
		}
	}
	fmt.Println("2d:", twoD)
	// [[0] [1 2] [2 3 4]]
}

func sliceRecycle() {
	cars := []string{"ford", "toyota", "ds", "honda", "suzuki"}
	//去掉后两个
	neededCars := cars[:len(cars)-2]
	carsCpy := make([]string, len(neededCars))
	copy(carsCpy, neededCars)
	//现在切片cars 可以被垃圾回收 因为neededCars不在被引用
	fmt.Println(carsCpy) //[ford toyota ds]

	//切片初始化陷阱
	s := make([]int, 5)
	s = append(s, 1, 2, 3)
	fmt.Println(s) //[0 0 0 0 0 1 2 3]

	s = make([]int, 0)
	s = append(s, 1, 2, 3)
	fmt.Println(s) //[1 2 3]
}

func autoEnlarge() {

	//自动扩容机制
	a := []string{}
	for i := 0; i < 6; i++ {
		a = append(a, "11111")
		fmt.Println("len :", len(a), "cap :", cap(a))
	}

	// 1 1
	// 2 2
	// 3 4
	// 4 4
	// 5 8
	// 6 8
}

func mapTest() {
	//只是声明 还没有初始化 不会报错 可以对nil的map取值 但是没有意义 赋值就报错了!!!
	var m1 map[string]string
	//会报错
	//m1["C"] = "c"
	m1 = make(map[string]string)
	m1["a"] = "a"
	m1["b"] = "b"

	m2 := make(map[string]string)
	m2["a"] = "a"
	m2["b"] = "b"

	m3 := map[string]string{
		"a": "a1",
		"b": "b1",
	}
	fmt.Println(m3) //map["a":"a1" "b":"b1"]

	if v, ok := m1["a"]; ok {
		fmt.Println(v) //a
	} else {
		fmt.Println("键值不存在")
	}

	for k, v := range m1 {
		fmt.Println(k, v)
		//a a
		//b b
	}
}

// map的并发使用
func syncMapTest() {
	c := make(map[string]string)
	//等待组
	wg := sync.WaitGroup{}
	//加锁 map不支持并发设置值 并发读可以好像??
	var lock sync.Mutex
	for i := 0; i < 10; i++ {
		wg.Add(1)
		//使用闭包进行变量捕获 内存逃逸 防止闭包陷阱
		go func(n int) {
			k, v := strconv.Itoa(n), strconv.Itoa(n)
			lock.Lock()
			c[k] = v
			lock.Unlock()
			wg.Done()
		}(i)
	}
	wg.Wait()
	fmt.Println(c) //map[0:0 1:1 2:2 3:3 4:4 5:5 6:6 7:7 8:8 9:9]
}

func cover() {
	var a = 89
	//强制转换
	var b = float32(a)
	fmt.Printf("%f\n", b) //89.000000
}

func mapExample() {
	m := make(map[string]int)
	m["k1"] = 7
	m["k2"] = 13

	fmt.Println("map:", m) //map[k1:7 k2:13]

	v1 := m["k1"]
	fmt.Println("v1: ", v1)      //7
	fmt.Println("len: ", len(m)) //2

	//删除
	delete(m, "k2")
	fmt.Println("map:", m)       //map[k1:7]
	fmt.Println("len: ", len(m)) //1

	_, prs := m["k2"]
	fmt.Println("prs:", prs) //false

	n := map[string]int{"foo": 1, "bar": 2}
	//map乱序遍历
	fmt.Println("map: ", n) //map[bar:2 foo:1]
}

func rangeExample() {
	nums := []int{2, 3, 4}
	sum := 0
	for _, num := range nums {
		sum += num
	}
	fmt.Println("sum:", sum) //9

	for i, num := range nums {
		if num == 3 {
			fmt.Println("index:", i) //1
		}
	}

	kvs := map[string]string{"a": "apple", "b": "banana"}
	for k, v := range kvs {
		fmt.Printf("%s -> %s\n", k, v)
		//a->apple
		//b->banana
	}

	//赋值单个是拿key
	for k := range kvs {
		fmt.Println("key:", k)
		//a
		//b
	}

	for i, c := range "go" {
		fmt.Println(i, c)
		//0 unicode(g)
		//1 unicode(o)
	}

	//注意!!! utf8编码
	//0 1 2 构成了中 3 4 5 构成了国
	for i, c := range "中国" {
		fmt.Println(i, c)
		//0 20013
		//3 22269
	}
}

func structExample() {
	user.Money -= 1000
	//取地址
	location := &user.Location
	//取值 +东西 对地址操作 能够更新
	*location = "xxx " + *location
	fmt.Println(user)
	//取地址
	var userOfTheMonth = &user
	////取值 +东西 对地址操作 能够更新
	userOfTheMonth.Location += " (proactive team player)"
	fmt.Println(user)
	//与上类似
	(*userOfTheMonth).Location += "y"
	fmt.Println(user)
}

var user User

type User struct {
	ID       int
	Name     string
	Address  string
	Birthday time.Time
	Location string
	Money    int
	TypeId   int
}
