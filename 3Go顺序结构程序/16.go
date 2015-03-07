//判断子串在主串中出现的位置
package main

import(
	"fmt"
	"strings"
)

func main(){
	var s,sep,chars string="Golang","an","lang"
	var r rune='a'
	fmt.Printf("%q第一次在%q中出现的索引是：%d\n",seq,s,strings.Index(s.sep))
	fmt.Printf("%q中的字符第一次在%q中出现的索引是：%d\n",chars,s,strings.IndexAny(s,chars))
	fmt.Printf("%q第一个符合f()的字符索引是：%d\n",s,strings.IndexFunc(s,f))
	fmt.Printf("%q第一次在%q中出现的索引是：%d\n",r,s,strings.IndexRune(s,r))
	fmt.Printf("")
}