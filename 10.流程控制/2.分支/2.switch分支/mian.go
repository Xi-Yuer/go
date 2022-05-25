package main
import "fmt"

func main() {
	// switch 不需要加 break;
	// switch 后也可以不用跟表达式
	// 如果需要穿透可以在语句后添加 fallthrough
	// 匹配的数据类型需要保持一致
	// case不能定义重复常量
	// default 不是必须的
	// 语法
	var n int = 3
	switch {
	case n < 1:
		fmt.Println("n > 1")
	case n > 2, n < 3, n < 4: //case 后可以跟多个表达式,表达式满足一个即可 类似于 或
		fmt.Println("n > 2")
		fallthrough           // 支持穿透(只穿透一次)
	case n < 5:
		fmt.Println(" n > 5")
	default:
		fmt.Println("default")
	}

	var m byte
	fmt.Println("请输入字符")
	fmt.Scanf("%c",&m)

	switch m {
		case 'a':
			fmt.Println("A")
		case 'b':
			fmt.Println("B")
		default :
			fmt.Println("Other")
	}

}