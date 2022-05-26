package main
import "fmt"

func main() {
	// go 语言中没有 while  do...while
	// 替代方式

	// while
	i := 0
	for{
		if(i > 10){
			break  // 跳出循环
		}
		// 循环
		fmt.Printf("hi \n")
		i++
	}

	// do .... while
	n := 0
		for{
			fmt.Printf("嘿 \n")
			n++
			if( n > 10) {
				break
			}
		}
}