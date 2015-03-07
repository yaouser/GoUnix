//字符串的遍历
package main

import(
	"fmt"
)

func main(){
	var str string
	str ="Hello 世界！"
	n:=len(str)
	fmt.Println("字节数据方式遍历：")
	for i:=0; i<n; i++ {
		ch:=str[i]
		fmt.Printf("str[%d]=%v\n",i,ch)
	}
	fmt.Println("unicode字符方式遍历：")
	for i,ch := range str {
		fmt.Printf("str[%d]=%v\n",i,ch)
	}
}