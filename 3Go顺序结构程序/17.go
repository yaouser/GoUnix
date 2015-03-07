//字符串追加和替换
package main

import(
	"fmt"
	"strings"
)

func main(){
	var s="na"
	var count int=2
	//字符串追加
	fmt.Println("ba"+strings.Repeat(s,count))
	//字符串替换
	fmt.Println(strings.Replace("google","o","oo",1))
	fmt.Println(strings.Replace("google","o","oo",-1))
}