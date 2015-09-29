//Method的基本定义
package main

import(
	"fmt"
)

type rectangle struct{
	width int
	height int
}
//函数area()是对象rectangle的一个方法
func (recv rectangle) area() int {
	return recv.width*recv.height
}
func main(){
	r1:=rectangle{4,3}
	r2:=rectangle{30,15}
	fmt.Println(r1.area(),r2.area())
}
/********************************************
Go语言的Method（方法）类似于一个函数，只是函数名前多了
个绑定类型参数receiver。
func (recv recviver_type) methodName(参数列表)（返回值）{
	......
}
Method中的Receiv可以是内置类型，自定义类型，结构体
或指针类型。
******************************************/	