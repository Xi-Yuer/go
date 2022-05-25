package main

import "fmt"

func main() {
	// 左移 右移
	var n int = 1 >> 2 // 0
	var m int = 1 << 2 // 4
	fmt.Println(n)
	fmt.Println(m)
	fmt.Println(2 & 3) // 2
	fmt.Println(2 | 3) // 3
	fmt.Println(2^3)   // 1
}