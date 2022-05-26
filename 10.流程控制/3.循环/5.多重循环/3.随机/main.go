package main
import (
	"math/rand"
	"time"
	"fmt"
)

func main() {
	var total int = 0
	// 为了生成一个随机数，需要给 rand 设置一个种子
	rand.Seed(time.Now().Unix()) // 返回一个时间戳: time.Now().Unix()  真随机
	for {
		var num = rand.Intn(100)  // 该函数会生成 [0 - n ) 的数字 ==> 包括0，不包括 n 
		fmt.Println(num)
		total++
		if(num == 99) {
			fmt.Printf("一共执行了%v次",total)
			break
		}

	}
}