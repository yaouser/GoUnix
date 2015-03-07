//函数的闭包
package main

import(
	"fmt"
)

func main(){
	f:=closures(10)
	fmt.Println(f(1))
	fmt.Println(f(2))
}
//注意函数的闭包返回值是匿名函数
func closures(x int) func(int) int{
	return func(y int) int{
		//注意内部函数引用的参数超出了它的作用范围
		return x+y
	}
}
/*******************************************
函数闭包的两个作用：
1函数闭包可以保护函数内的变量安全。
2函数闭包可以在内存中维持一个变量。
*******************************************/