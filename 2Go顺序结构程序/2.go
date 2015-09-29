//print()函数默认格式打印
package main

import(
	"fmt"
)

func main(){
	var id int = 100
	var name string="李明"
	var grade float32=91.5
	fmt.Print("默认格式打印：\n")
	fmt.Print("学号：",id)
	fmt.Print("\n")
	fmt.Print("姓名：",name)
	fmt.Print("\n")
	fmt.Print("成绩：",grade)
}