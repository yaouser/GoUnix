//变参函数的声明
package main

import(
	"fmt"
)

func main(){
	f1(1,2,3)
	f1(1,2,3,4,5)
}
//声明函数f1()为变参函数
func f1(args ... int){
	//变参args其实质是一个切片
	fmt.Println(args)
}
/*****************************************
Go语言支持不定长变参，但是要注意不定长变参只能作为
函数的最后一个参数，不能放在其他参数的前面。
变参的定义格式为"... 类型",而且变参必须是函数的最后
一个参数，例如：func f1(a int,s string,args ...int){}
不定长变参其实质就是一个切片，可以使用range进行遍历
func f1(args ...int){
	for _,arg:=range args{
		fmt.Println(arg)
	}
}
****************************************/
