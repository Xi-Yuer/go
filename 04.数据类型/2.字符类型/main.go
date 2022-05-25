package main
import "fmt"	

func main()  {
	// Println：打印字符串、变量；
	// Printf：打印需要格式化的字符串，可以输出字符串类型的变量；不可以输出整型变量和整型；
	// 简单理解，当需要格式化输出信息时，一般选择Printf，其余使用Println。


	// %T 输出类型
	// %d 十进制整数
	// %c 一个Unicode的字符(格式化输出)
	// %b 一个二进制整数，将一个整数格式转化为二进制的表达方式

	 n := 'a'
	 m := '0'

	 v := 100

	 var z string = "嘿"
	
	//  输出 ASCII表中的值
	 fmt.Println(n,m)  // 97 48
	 fmt.Println(z)  // 嘿
	 fmt.Printf("%T \n",z) // string 

	//  格式化输出
	fmt.Printf("%c %c",n,m) // a 0

	fmt.Printf("%b",v) // 01100100
	
}