//字符串和切片的格式化输出
package main

import(
	"fmt"
)

func main(){
	var str string="Go language"
	fmt.Println("字符串和切片格式化输出：")
	fmt.Printf("直接输出字符串或切片%s\n",str)
	fmt.Printf("自动加双引号：%q\n",str)
	fmt.Printf("每个字节用两个字符十六进制表示（a-f）：%x\n",str)
	fmt.Printf("每个字节用两个字符十六进制表示（a-f）：%X\n",str)
}