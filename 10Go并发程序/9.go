//强行终止Goroutine
package main

import(
	"fmt"
	"time"
	"runtime"
)

func main(){
	go func(){
		defer fmt.Println("Goroutine1 defer...")
		for i:=0;i<10;i++ {
			if i==5 {
				runtime.Goexit()
			}
			fmt.Println("Goroutine1:",i)
		}
	}()
	go func(){
		fmt.Println("Goroutine2")
	}()
	time.Sleep(5*time.Second)
}
/*************************************************
Goroutine1一共要执行10个工作步骤，当执行到第5个工作步骤调用
了Goexit()函数，则Goroutine被强行终止执行。Goroutine2则
获取权限得到执行。
/*************************************************
如果要强行终止一个Goroutine的执行，可以调用Goexit()函数来
完成。Goexit()将终止整个堆栈链，并在内层退出，但是defer语句
仍然或被执行。
*************************************************/