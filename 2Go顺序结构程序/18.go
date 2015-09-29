//将数值转换成字符串
package main

import(
	"fmt"
	"strconv"
)

func main(){
	var b bool=false
	var f float64=3.14
	var i int64=-1024
	var ui uint64=1024
	fmt.Printf("将布尔数%t转换成字符串：%q\n",b,strconv.FormatBool(b))
	fmt.Printf("将浮点数%f转换成字符串：%q\n",f,strconv.FormatFloat(f,'f',2,32))
	fmt.Printf("将整数%d转换成字符串：%q\n",i,strconv.FormatInt(i,10))
	fmt.Printf("将无符号整数%d转换成字符串：%q\n",ui,strconv.FormatUint(ui,10))
}