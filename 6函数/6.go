//任意类型的变参
package main

import(
	"fmt"
)

func main(){
	f1(2,"Go",8,"Language",'a',false,'A',3,14)
}
//声明函数f1()为任意类型变参
func f1(args ...interface{}){
	var num=make([]int,0,6)
	var str=make([]string,0,6)
	var ch=make([]int32,0,6)
	var other=make([]interface{},0,6)
	for _,arg:=range args{
		switch v:=arg.(type){
			case int:
				num=append(num,v)
			case string:
				str=append(str,v)
			case int32:
				ch=append(ch,v)
			default:
				other=append(other,v)
		}
	}
	fmt.Println(num)
	fmt.Println(str)
	fmt.Println(ch)
	fmt.Println(other)
}
/****************************************
在Go语言中，空接口interface{}可以指向任何数据对
象，所以可以使用interfa{}定义任意类型变参。同时，
interfa()也是类型安全的。
Go语言规定，如果希望传递任意类型的变参，变参类型应
指定为空接口interfac{}
****************************************/