//整数和字符串的相互转换
package main

import(
	"fmt"
	"strconv"
)

func main(){
	var i int=-1024
	var s string="1024"
	value,_:=strconv.Atoi(s)
	str:=strconv.Itoa(i)
	fmt.Printf("将字符串%q转换成整数：%d\n",s,value)
	fmt.Printf("将整数%d转换成字符串：%q\n",i,str)
}