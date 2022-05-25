package main
import "fmt"

func main() {
	// 顺序控制(程序从上到下依次执行)
	var n int = 10
	var m int = 20
	fmt.Println(n,m) // √

	
	// fmt.Println(z)   // × (不能访问之后的变量) 变量采用向前引用
	// var z int = 20
}