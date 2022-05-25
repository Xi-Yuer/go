package main
import (
	"fmt"
)

var i int = 10
// 将 int 数据类型转化为 float64 类型
// 语法： T(v)  ==>  [数据类型](value)
var n float64 = float64(i)

// 数据可以从 小 => 大 | 大 => 小
// 		大 => 小 时数据溢出会导致与期望结果不一样( int64 => int8 )
// 		转化时需要考虑范围
// 数据类型转化不影响原来的数据类型
// 不同数据类型转化进行运算时同样需要考虑范围问题，否则会造成溢出或编译不通过问题


var a int32 = 10
// var b int64 = a  //不能将 int32 赋值给 int64, 编译不能通过

var c int64 = int64(a)  // 转化之后再赋值




func main()  {
	fmt.Println(n) 		  // 10
	fmt.Printf("%T \n",i) // int
	fmt.Printf("%T \n",n) // float64
}