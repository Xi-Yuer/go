package main

import "fmt"

func main() {
	// 定义变量 int 的默认值为 0
	var i int
	i = 10
	fmt.Println("i = ", i)

	// 变量声明
	num := 10
	fmt.Println("num = ", num)

	// 多变量声明
	var n1, n2, n3 int
	fmt.Println(n1,n2,n3) // 0 0 0  

	var m1, m2, m3 = 3, "M1", 2
	fmt.Println(m1,m2,m3)

}