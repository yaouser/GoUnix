//指针变量的定义和引用
package main

import(
	"fmt"
)

func main(){
	var i int
	var i_pointer* int
	i=100
	i_pointer=&i  //将i的内存地址存放到指针变量i_pointer中
	fmt.Println(*i_pointer)
}