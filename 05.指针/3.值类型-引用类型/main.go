package main
import "fmt"

func main() {
	// 值类型有：tring, bool, int系列, float系列, 数组, 结构体struct
	// 		- 值类型变量直接存储值（栈）

	// 引用类型：指针, slice切片, map, 管道chan, interface
	// 		- 引用类型变量存储 内存地址，该内存地址指向的才是真正的值 （堆）

	// 标识符
	// 不能以数字开头，不能是单独的 _ ， 不能是关键字， 不能有空格，不能使用（ - + * / ），严格区分大小写
	// 变量 函数 常量采用 驼峰命名
	// 首字母大写才能够被外部访问（公有），首字母小写只能在本包中访问（私有）
	var _num int = 10
	fmt.Println(_num)
}