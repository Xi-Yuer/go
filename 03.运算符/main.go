package main

import "fmt"

// 全局变量
var (
	num1 = 10
	num2 = 20
	str1 = "Tom"
	str2 = "Sily"
)

func main()  {
	// 数值运算
	fmt.Println(num1 + num2)
	// 字符串拼接
	fmt.Println(str1 + str2)
}

