package main

// import "fmt"
// import "unsafe"

// 引入方式
import (
	"fmt"
	"unsafe"
)

func main()  {
	// int(和电脑系统有关 32位 64位), int8(-128 - 127), int16(-32768 - 32767), int32, int64
	// uint(无符号 大于零), uint8(0 - 255), uint16(0 - 2 的16次方 - 1), uint32(0 - 2 的32次方 - 1), uint64(0 - 2 的64次方 - 1)
	// rune
	// byte(0 - 255)

	var i int = 10
	var n uint8 = 255
	
	fmt.Println(i)
	fmt.Println(n)
	
	// 查看变量数据类型 (fmt.Printf("%T"))
	fmt.Printf("数据类型是 %T",i) // int

	// 查看数据及占用字节大小 (fmt.Printf("%d", n, unsafe.Sizeof(n)))
	fmt.Printf("占用字节 %d", unsafe.Sizeof(i)) // 8

	// 浮点类型(float32:单精度 4字节 float64:双精度 8 字节) 默认类型 float64
	// 科学计数法 2.12345e2 ==> 2.12345 * 10 的 2 次方
	// 推荐使用 float64 双精度

	var z float32 = 1.2

	fmt.Printf("数据类型是 %T",z)
}