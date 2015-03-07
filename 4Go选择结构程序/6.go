//if语句初始化条件表达式
package main

import(
	"fmt"
)

func main(){
	a:=10
	if a:=1;a>0{
		//打印if语句块中条件变量a的值
		fmt.Println(a)
	}
	//打印if语句块外变量a的值
	fmt.Println(a)
}