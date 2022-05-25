package main
import (
	"fmt"
	_"strconv"
	
)


// 基本数据类型 ==> string
// 方式一： fmt.Sprintf() 根据参数返回格式化的字符串
// 方式二： strconv
// 			strconv.FormatBool(i)
// 			strconv.FormatFloat(i,fmt,prec,bitSize)
// 			strconv.FormatInt(i,base) | strconv.Itoa(i)

// string ==> 基本数据类型
// 			strconv.ParseBool(str) 返回值 [value bool,err error]  [ value, _ ] 可以这样接受第一个参数，忽略第二个参数
// 			strconv.ParseInt(str, base int, bitSize int) 返回值 [ value int, err error ]
// 			strconv.ParseFloat(str,bitSize int) 返回值 [value float, err error]

var str string
// 将下面的数据转化为 string
var num1 int = 10
var num2 float64 = 23.4
var b bool = true
// 将下面的string转换为基本数据类型
var str1 string = "1"
var str2 string = "true"
var str3 string = "130"


func main()  {
	// int ==> string("%d")
	str = fmt.Sprintf("%d",num1)
	fmt.Println(str) 				// 10
	fmt.Printf("%T",str)				//string

	// foat ==> string("%f")
	str = fmt.Sprintf("%f",num2)
	fmt.Println(str) 				    // 23.400000
	fmt.Printf("%T",str) 				// string

	// foat ==> string("%t")
	str = fmt.Sprintf("%t",b)
	fmt.Println(str) 				    // true
	fmt.Printf("%T",str) 				// string
}