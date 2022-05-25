package main
import "fmt"

func main()  {
	// 二进制(在 go 中不能直接表示二进制)
	// 八进制  0 开头
	// 十六进制 0x 开头
	var n int = 10
	var m int = 010 // 8
	fmt.Println(n) 	// 10
	fmt.Printf("%b \n",n) // 以二进制输出 1010
	fmt.Println(m)
}