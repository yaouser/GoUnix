//通用格式化输出
package main

import(
	"fmt"
)

func main(){
	var id int=100
	var name string="李明"
	var grade float32=91.5
	fmt.Println("默认格式打印：")
	fmt.Printf("%v %v %v\n",id,name,grade)
	fmt.Printf("%#v %#v %#v\n",id,name,grade)
	fmt.Printf("%T %T %T\n",id,name,grade)
	fmt.Printf("60%%\n")
}