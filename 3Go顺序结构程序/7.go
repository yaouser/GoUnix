//浮点型数据和复数的格式化输出
package main

import(
	"fmt"
)

func main(){
	var f float32=123.4567
	var cp=complex(1.2,3.4)
	fmt.Println("浮点型数，复数格式化输出：")
	fmt.Printf("无小数两位指数科学计数法：%b\n",f)
	fmt.Printf("科学记数法（小写）：%e\n",f)
	fmt.Printf("科学记数法（大写）：%E\n",f)
	fmt.Printf("根据实际情况采用%%e或%%f（小写）：%g\n",f)
	fmt.Printf("根据实际情况采用%%e或%%f（大写）：%G\b",f)
	fmt.Printf("只有小数部分无指数部分：%f\n",f)
	fmt.Printf("保留2为小数：%6.2f\n")
	fmt.Printf("复数%v的实部=%g 虚部=%g\n",cp,real(cp),imag(cp))
}