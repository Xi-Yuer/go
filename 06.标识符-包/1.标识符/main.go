package main
import "fmt"

func main(){
		// 标识符
	// 不能以数字开头，不能是单独的 _ ， 不能是关键字， 不能有空格，不能使用（ - + * / ），严格区分大小写
	// 变量 函数 常量采用 驼峰命名
	// 首字母大写才能够被外部访问（公有），首字母小写只能在本包中访问（私有）
	var _num int = 10
	fmt.Println(_num)
}