//其它类型格式输出
package main

import(
	"fmt"
)

func main(){
	var a int=97
	var f float32=-1.32
	var p* int
	var str string="Golang"
	p=&a
	fmt.Println("其它类型格式输出：")
	fmt.Printf("输出数值的正负号：%+d,%+g\n",a,f)
	fmt.Printf("ASCII码格式输出：%+q\n",a)
	fmt.Printf("切换格式（#o,#x,#X）：%d %#o %#x %#X\n",a,a,a,a)
	fmt.Printf("消除指针地址前的0X：%p %#p\n",p,p)
	fmt.Printf("以十六进制输出字符串：%s,%x\n",str,str)
}