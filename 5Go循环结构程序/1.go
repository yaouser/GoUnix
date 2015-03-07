//使用基本循环控制结构统计单词中的某个字符数
package main

import(
	"fmt"
)

func main(){
	var words string
	var ch byte
	var ln,count int
	//从键盘格式化输入单词
	fmt.Scanf("%s",&words)
	fmt.Println("\n")
	//输入要统计的字符
	fmt.Scanf("%c",&ch)
	//计算单词的字节长度
	ln=len(words)
	/********************************************
	算法：从单词的第一个字符开始扫描，找到符合条件的字符，
	计数器加1
	********************************************/
	for i:=0; i<ln; i++ {
		if byte(words[i])==ch {
			count++
		}
	}
	fmt.Printf("There are %d %q in the %q.\n",count,ch,words)
}