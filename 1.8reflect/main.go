package main

import (
	"fmt"
	"reflect"
)

func main() {
	reflectTypeof("hello")
	reflectValueOf(int64(1))
	reflectTest("hello")

	var s = Student{
		Name:  "orange",
		Sex:   1,
		Age:   10,
		Score: 80,
	}

	//注意!!! 这里要分清楚!!!
	v := reflect.ValueOf(s)
	t := v.Type()
	kind := t.Kind()
	switch kind {
	case reflect.Struct:
		fmt.Printf("s is struct\n")
		fmt.Printf("field num of s is %d\n", t.NumField())
		for i := 0; i < t.NumField(); i++ {
			field := v.Field(i)
			fmt.Printf("name:%s type:%v value:%v \n", t.Field(i).Name, t.Field(i).Type, field.Interface())
		}
	default:
		fmt.Printf("default\n")
	}
}

type Student struct {
	Name  string
	Sex   int
	Age   int
	Score float32
}

func reflectTypeof(a interface{}) {
	t := reflect.TypeOf(a)
	fmt.Printf("type of a is: %v\n", t) //type of a is: string

	k := t.Kind()
	switch k {
	case reflect.Int64:
		fmt.Printf("a is int64\n")
	case reflect.String:
		fmt.Printf("a is string\n") //a is string
	}

	switch a.(type) {
	case int64:
		fmt.Printf("a is int64\n")
	case string:
		fmt.Printf("a is string\n") //a is string
	}
}

func reflectValueOf(a interface{}) {
	v := reflect.ValueOf(a)
	k := v.Kind()
	switch k {
	case reflect.Int64:
		fmt.Printf("a is int64 , store value is: %d\n", v.Int()) //a is int64 , store value is: 1
	case reflect.String:
		fmt.Printf("a is string , store valule is: %s\n", v.String())
	}
}

func reflectTest(b interface{}) {
	rType := reflect.TypeOf(b)
	fmt.Println("rType =", rType) //rType = string

	rVal := reflect.ValueOf(b)
	typeKind := rType.Kind()
	valKind := rVal.Kind()
	fmt.Printf("typeKind = %v , valKind = %v\n", typeKind, valKind) //typeKind = string , valKind = string
}
