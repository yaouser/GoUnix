//出让时间片
package main

import(
	"fmt"
	"time"
	"runtime"
)

func main(){
	go func(){
		for i:=0;i<5;i++{
			if i==2 {
				runtime.Gosched()
			}
			fmt.Println("Goroutine1:",i)
		}
	}()
	go func(){
		fmt.Println("Goroutine2")
	}()
	time.Sleep(5*time.Second)
}
/*******************************************
在设计并发任务时，用户可以在每个Goroutine中控制何时主
动让出时间片给其他Goroutine，这可以使用Gosched()函
数来实现。当让出当前Goroutine的执行权限，调度器会安排
其他等待的任务区运行，并在下轮某个时间片再从该位置恢复
执行。
*******************************************/