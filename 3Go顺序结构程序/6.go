//整型数据格式化输出
package main

import(
	"fmt"
)

func main(){
	var year int=20013
	fmt.Println("整形数据格式化输出：")
	fmt.Printf("二进制格式：%b\n",year)
	fmt.Printf("十进制格式：%d\n",year)
	fmt.Printf("八进制格式：%o\n",year)
	fmt.Printf("十六进制格式（a-f）：%x\n",year)
	fmt.Printf("十六进制格式（A-F）：%X\n",year)
	fmt.Printf("数值对应的Unicode编码：%c\n",year)
	fmt.Printf("unicode格式：%U\n",year)
	fmt.Printf("unicode编码使用单引号括起来：%q\n",year)
}