//判断两个字符串是否相等
package main

import(
	"fmt"
	"strings"
)

func main(){
	var str1,str2,str3 string="Go","go","lang"
	fmt.Println("判断两个字符串是否相等：")
	fmt.Printf("%q等于%q? %t\n",str1,str2,strings.EqualFold(str1,str2))
	fmt.Printf("%q等于%q？%t\n",str1,str3,strings.EqualFold(str1,str3))
}