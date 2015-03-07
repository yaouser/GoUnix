//字符串的基本操作
package main

import(
	"fmt"
)

func main(){
	var str string
	str="Go"+"lang"
	fmt.Println(str)
	fmt.Println("字符串长度：",len(str))
	fmt.Printf("从字符串中取字符：%c\n",str[1])
}