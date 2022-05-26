package main
import "fmt"

func main() {
	// 字符串 | 数组 遍历 (for-range)
	str := "Tom"
	for index, value := range str {
		fmt.Println(index)
		fmt.Printf("%c \n",value) // T o m
		fmt.Println(value)     // 84 111 109
	}
	for i := 0; i < len(str); i++ {
		fmt.Printf("%c \n",str[i]) // T o m
		fmt.Println(str[i]) // 84 111 109
	}

	// 如果字符串含有中文字符(默认for-range 安一个字节遍历 一个汉字为三个字节)
	//解决方案 将汉字转成切片 [] rune 

	var strr string = "你好呀，李银河！"
	for index, v := range strr {
		fmt.Println(index) // 0 3 6 9 12 ...... ==> 一个汉字为三个字节
		fmt.Printf("%c \n",v)
	}
	var strn = []rune(strr) 
	for index, v := range strn{
		fmt.Println(index) // 0 1 2 3 4 .....   ==> 转成了切片
		fmt.Printf("%c \n",v)
	}
}