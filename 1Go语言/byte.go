//字节型数据的定义和使用
package main

import(
	"fmt"
)

func main(){
	var ch1,ch2 byte
	ch1=65
	ch2='A'
	fmt.Println(string(ch1))  //要以字符型输出需先转换
	fmt.Println(ch2)          //字节型数据直接输出为整数形式
}