package main

func main() {

}

// PublicFunc 公共函数首字母大写
// 在其他包内通过包名.公共函数名调用
func PublicFunc() {

}

// 私有函数首字母小写 在其他包无法通过包名.函数名调用
func privateFunc() {

}

// PublicString 公共变量在同一个包内可以直接访问调用
// 在其他包内可以通过包名.公共变量调用
var PublicString string

// 私有变量在同一个包内可以直接访问调用
var privateString string
