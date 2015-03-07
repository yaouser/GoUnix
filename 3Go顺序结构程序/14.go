//字符串查找操作
package main

import(
	"fmt"
	"strings"
)

func main(){
	var str string="Hello,World!"
	var s1,s2,s3,s4 string="llo","go","ll","l"
	var ch1,ch2 rune='c','d'
	fmt.Println("查找子串是否在指定的字符串中：")
	fmt.Printf("%q在%q中？%t\n",s1,str,strings.Contains(str,s1))
	fmt.Printf("%q在%q中？%t\n",s2,str,strings.Contains(str,s2))
	fmt.Println("查找字符串是否含有子串中的字符：")
	fmt.Printf("%q中含有%q的字符？%t\n",str,s1,strings.ContainsAny(str,s1))
	fmt.Printf("%q中含有%q的字符？%t\n",str,s2,strings.ContainsAny(str,s2))
	fmt.Println("查找字符串是否含某个字符：")
	fmt.Printf("%q中含有%q的字符？%t\n",str,s1,strings.ContainsRune(str,ch1))
	fmt.Printf("%q中含有%q的字符？%t\n",str,s2,strings.ContainsRune(str,ch2))
	fmt.Println("统计指定字符串包含子串的个数：")
	fmt.Printf("%q中含有%d的字符？%q.\n",str,strings.Count(str,s3),s3)
	fmt.Printf("%q中含有%d的字符？%q.\n",str,strings.Count(str,s4),s4)
}