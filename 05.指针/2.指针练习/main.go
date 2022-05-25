package main
import (
	"fmt"
)

// 获取一个 int 变量的地址 并显示到终端
// 将 num 的指针赋值给 ptr 并通过 ptr 去修改 num 
func main(){
	num := 10
	var ptr *int = &num
	*ptr = 11
	fmt.Println(&num)
	fmt.Println(num)
}
