package main

import "fmt"

func main() {
	// 方式一
	for i := 0; i < 10; i++ {
		fmt.Println("*")
	}

	var n int = 10
	// 方式二
	for n > 0 {
		n--
		fmt.Println("n")
	}
	// 方式三
	var m int = 0
	for {
		fmt.Println(m)
		m++
		if m > 10 {
			break  // 使用 break 可以退出循环
		}
	}
}