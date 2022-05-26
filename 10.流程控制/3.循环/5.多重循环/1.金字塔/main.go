package main

import "fmt"

func main() {
	// for i := 1; i <= 5; i++ { // i 表示成数
	// 	for j := 1; j <= i; j++ { // j 表示没层打印多少个 * 
	// 			fmt.Print("*")
	// 		}
	// 		fmt.Print("\n")
	// 	}
			
	// for 循环打印金字塔
	for i := 1; i <= 5; i++ { // i 表示成数
		for n := 1; n <= 5 - i; n++ {
			fmt.Print(" ")
		}
		for j := 1; j <= 2 * i - 1; j++ { // j 表示没层打印多少个 * 
			fmt.Print("*")
		}
		fmt.Print("\n")
	}
}