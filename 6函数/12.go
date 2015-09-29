//panic-and-recover异常恢复机制
package main

import(
	"fmt"
)

func main(){
	//必须要先声明defer，否则捕获不到panic异常
	defer func(){
		fmt.Println("函数defer开始运行...")
		if err:=recover();err!=nil{
			//这里的err就是panic传入的内容
			fmt.Println("程序异常退出：",err)
		}else{
			fmt.Println("程序正常退出。")
		}
	}()
	f(101)
}
func f(a int){
	fmt.Println("函数f开始运行...")
	if a>100 {
		panic("参数值超出范围！")
	}else{
		fmt.Println("函数f调用结束。")
	}
}
/*******************************************
panic是内置函数，可以中断原有的控制流程，进入一个异常
流程中。
recover也是一个内置的函数，他可以进入异常流程中的
goroutine恢复过来。recover仅在延迟函数中有效。在正常
的执行过程中，调用recover会返回nil并且没有其他任何效果
。如果goroutine陷入异常，调用recover可以捕获到panic
的输入值，并且恢复正常的执行。
*******************************************/