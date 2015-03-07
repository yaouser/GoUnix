//变参的传递
package main

import(
	"fmt"
)

func main(){
	f1(1,2,3,4,5)
}
//声明函数f1(),f2()为变参函数
func f1(args ...int){
	//全部传递，即将函数f1()的args直接复制给函数f2()
	f2(args ...)
	//部分传递，将args中的第3个以后的元素复制给函数f2()
	f2(args[2:]...)
}
func f2(args ...int){
	fmt.Println(args)
}