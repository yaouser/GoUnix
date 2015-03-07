//格式化输出指针地址
package main

import(
	"fmt"
)

func main(){
	var i int = 100
	var i_pointer* int
	i_pointer=&i
	fmt.Println("格式化输出指针地址：")
	fmt.Printf("输出以0x开头的指针地址：%p\n",i_pointer)
}