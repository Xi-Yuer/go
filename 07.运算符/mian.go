package main
import "fmt"

func main() {
	// 运算符(+, -, *, /, %, ++, --, )
	// 关系运算符 (==, !=, <, >, >=, <=)
	// 逻辑运算符 (&& || !)
	// 类型不一样的类型不能进行运算

	fmt.Println(10 / 4) // 2
	fmt.Println(10.0 / 4) // 2.5
	fmt.Println(float32(10) / float32(4)) // 2.5

	var n float32 = 10
	var m float32 = 4
	fmt.Println(n / m) // 2.5

	// var x float32 = 10
	// var y float64 = 10

	// fmt.Println(x + y) err: mismatched types float32 and float64
	
	// 取模公式: a % b = a - a / b * b
	fmt.Println(10 % 3) // 1
	fmt.Println(3 % 10) // 3
	// n++ n-- 只能用于独立语言使用
	// m:= n++ 错误✖

	// 还有 97 天放假 还有几个星期零几天
	var day int = 97
	fmt.Println("还剩",day / 7, "星期")
	fmt.Println( day % 7,"天")

	// 摄氏度转化（5/9*(华氏 - 100)）
	var c float32 = 134.2
	var v = fmt.Sprintf("%v",float32(5)/float32(9)*(c - float32(100)))
	fmt.Println(v) // 19


	// var i int = 10
	// var j float64 = 10
	// fmt.Println(i < j)  err：mismatched types int and float64
	// fmt.Println(!(3 - 2 > 0)) // false   


	var d int = 10
	var q int = 20

	d = d + q
	q = d - q
	fmt.Println(d,q)
}