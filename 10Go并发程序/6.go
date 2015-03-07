//使用Select实现Channel超时机制
package main

import(
	"fmt"
	"time"
)

func main(){
	ch:=make(chan int)
	timeout:=make(chan bool,1)
	go func(){
		time.Sleep(5*time.Second)
		timeout<-true
	}()
	select{
		case<-ch:
		case<-timeout:
			/*********************************
			如果其他case语句中的Channel都没有在5s之
			内接收数据，则只有timeout接收到数据，则
			意味着主程序超时退出。
			*********************************/
			fmt.Println("timeout...")
			break
	}
}
/*******************************************
在向Channel写数据时发现已满，或从Channel读数据时发现
为空，如果不正确处理这些问题，很可能会导致整个Goroutine
死锁。在Go语言并发编程的通信过程中，所有错误处理都由超
时机制来完成。
超时机制是一种解决通信死锁的机制，通常会设置一个超时参数，
通信双方如果在设定的时间内仍然没处理完任务，则该处理过程
会立即被终止。
Go语言没有提供直接的超时处理机制，但可以利用Select机制
解决超时问题。因为Select的特点是只要其中一个case已经完
成，程序就会继续往下执行，而不会考虑其他case的执行情况。
基于此特性，就可以使用Select为Channel实现超时处理功能。
********************************************/

/********************************************
另外，在Go语言的time包中还提供了After，Tick等函数，它
们也可以返回计时器Channel，利用它们也可以试想超时机制。
func main(){
	ch:=make(chan int)
	lock:=make(chan bool)
	go func(){
		for{
			select{
				case<-ch:
				case<-time.After(5*time.Second):
					fmt.Println("timeout...")
					lock<-true
					break
			}
		}
	}()
	<-lock
}
********************************************/