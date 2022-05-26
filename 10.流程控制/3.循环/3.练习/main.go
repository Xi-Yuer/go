package main
import "fmt"

func main() {
	// 打印 1 ~ 100 之间所有 9 的倍数之和以及个数
	var n int64 = 0
	var sum int64 = 0
	var i int64 = 1
	for ; i <= 100; i++ {
		if i % 9 == 0 {
			n++
			sum += i
		}
	}
	fmt.Printf("%v\n",n)
	fmt.Printf("%v\n",sum)

	// 
}