package main

import "fmt"

// fmt.Scanln(&value) 多次接受
// fmt.Scanf("%s %d %f %t",&value1,&value2,&value3,&value4)
func main() {
	// 当程序执行到 fmt.Scanln()会停止住,等待用户输入并回车
	var name string
	var age int
	var salary int
	var isPass bool
	fmt.Println("请输入姓名")
	fmt.Scanln(&name) // 将用户的输入赋值给 name
	fmt.Println("请输入年龄")
	fmt.Scanln(&age)
	fmt.Println("请输入薪水")
	fmt.Scanln(&salary)
	fmt.Println("是否通过考试")
	fmt.Scanln(&isPass)

	fmt.Println(name,age,salary,isPass)
}