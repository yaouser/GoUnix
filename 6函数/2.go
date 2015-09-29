//自定义函数的声明与调用（调用自定义函数）
package main

import(
	"fmt"
	"errors"
)

//函数f1()无参数且无返回值
func f1(){
	fmt.Println("Hello World!")
}
//函数f2()有参数但无返回值
func f2(a,b int,c string){
	fmt.Println(a,b,c)
}
//函数f3()有一个返回值
func f2(a,b int) int {
	return a+b
}
//函数f4()有两个返回值
func f4(a,b int)(ret float32,err error){
	if b==0 {
		err=errors.New("Overflow!")
		return
	}else{
		return float32(a)/float32(b),nil
	}
}
func main(){
	var a,b int=2,3
	var c string="Golang"
	f1()
	f2(a,b,c)
	sum:=f3(a,b)
	fmt.Println(sum)
	f,err:=f4(a,b)
	fmt.Println(f,err)
}