//统计字符串中的字符，数字个数
package main

import(
	"fmt"
)

func main(){
	//str用于存储输入字符串
	var str string
	//下列变量分别用于英文字符，数字，其他字符个数统计
	var charCount,numCount,otherCount int
	fmt.Println("请输入字符串...")
	fmt.Scanf("%s",&str)
	for _,v:=range str {
		if v>=48 && v<=57 {
			numCount++
		}else if (v>=65&&v<=90)||(v>=97&&v<=122) {
			charCount++
		}else{
			otherCount++
		}
	}
	fmt.Printf("共有%d个数字\n",numCount)
	fmt.Printf("%d个英文字母\n",charCount)
	fmt.Printf("%d个其他字符\n",otherCount)
}