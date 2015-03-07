//println()函数默认打印格式
package main

import(
	"fmt"
)

func main(){
	var id int=100
	var name string="李明"
	var grade float32=91.5
	fmt.Println("默认格式打印：")
	fmt.Println("学号：",id)
	fmt.Println("姓名：",name)
	fmt.Println("成绩：",grade)
}