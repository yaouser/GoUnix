//将字符串转换成数值
package main

import(
	"fmt"
	"strconv"
)

func main(){
	var strb,strf,stri,strui string="false","3.14","-1024","1024"
	var b bool
	var f float64
	var i int64
	var ui uint64
	b,_=strconv.ParseBool(strb)
	f,_=strconv.ParseFloat(strf,32)
	i,_=strconv.ParseInt(stri,10,32)
	ui,_=strconv.ParseUint(strui,10,32)
	fmt.Printf("将字符串%q转换成布尔数：%t\n",strb,b)
	fmt.Printf("将字符串%q转换成浮点数：%f\n",strf,f)
	fmt.Printf("将字符串%q转换成整数：%d\n",stri,i)
	fmt.Printf("将字符串%q转换成无符号整数：%d\n",strui,ui)
}