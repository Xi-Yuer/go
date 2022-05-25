package main

import (
	"fmt"
)

// 每个值类型都有对应的指针类型( *int, *float32, *float64, *string, *bool..... )
func main()  {
	var n int = 10
	var m int = 10
	// &n 可以打印出 n 的内存地址
	fmt.Println(&n) // 0xc0000aa058
	fmt.Println(&m) // 0xc0000120b0

	// 指针变量存的是一个内存地址，该内存地址所指向的空间才是值

	// ptr 是一个指针变量，类型是 *int(int类型的指针),值为 &n
	var ptr *int = &n
	fmt.Println(&ptr) // 0xc000006030 ptr 本身的内存地址
	fmt.Println(ptr) // 值为：0xc0000aa058 与 n 的内存地址相同 并且该内存地址指向 10 这个值

	// 获取指针类型所指向的值
	fmt.Println(*ptr) // 10
}