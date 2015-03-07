package main

import(
	"fmt"
)
/********************************************
const(
	sunday		=iota//0
	monday		=iota//1
	tuesday		=iota//2
	wednesday	=iota//3
	thursday	=iota//4
	friday		=iota//5
	saturday	=iota//6
)
Go语言的常量和枚举之间的转换时非常灵活的，它在常量组中使用计数器iota，
iota从0开始，常量组中每定义一个常量iota自动递增1。
const(
	sunday		=iota//0
	monday		     //1
	tuesday		     //2
	wednesday	     //3
	thursday	     //4
	friday		     //5
	saturday	     //6
)
**********************************************/

/***********************************************
常量组中每定义一个常量，iota就会自动递增1，而每遇到一个
const关键字，iota就会重置为0。
**********************************************/
const(
	a='A'
	b
	c=iota
	d
)
const(
	e='E'
	f=iota
)

func main(){
	fmt.Print(a)
	fmt.Println(b)
	fmt.Println(c)
	fmt.Println(d)
	fmt.Println(e)
	fmt.Println(f)
}