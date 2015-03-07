//Select机制
package main

import(
	"fmt"
)

func main(){
	ch1:=make(chan int)
	ch2:=make(chan string)
	lock:=make(chan bool)
	go func(){
		for {
			select{
				case value,ok:=<-ch1:
					if !ok {
						/*
						**向Channel的lock中写入数据，以保持让for语句
						**通过select持续的监听多个Channel的消息
						*/
						lock<-true	
						break
					}
					fmt.Println("ch1=",value)
				case value,ok:=<-ch2:
					if !ok {
						lock<-true
						break
					}
					fmt.Println("ch2=",value)
			}
		}	
	}()
	ch1<-100
	ch2<-"Golang"
	ch2<-"Go"
	ch1<-200
	close(ch1)
	close(ch2)
	<-lock		//将Channel值丢弃，解除锁，退出main函数。
}
/************************************************
在Go语言中，Select机制主要用于解决通道通信中的多路复用问题。
因为通道的接收操作往往是阻塞式的，所以Select机制还经常和超
时机制配合使用，将阻塞式的通信转换为非阻塞式，以提高系统通信
效率。
系统可以通过调用select函数来监控一系列文件句柄，一旦其中一个
文件句柄发生I/O操作，该select函数就会被返回。在网络编程中，
select机制也被用来监听多个非阻塞套接字（socket），以实现高
并发的服务器程序。
select的用法和Switch语句非常类似，但是select中的每个case
语句必须是一个I/O操作。select机制的基本机构如下：
select{
	case<-chan1:
		//如果chan1成功读取数据，则进行该case处理语句。
	case<-chan2:
		//如果chan2成功读取数据，则进行该case处理语句。
		.......
	default:
		//如果上面都没有成功，则进入default处理流程。
}
由此可知，select不像Switch语句，后面并不带判断条件，而是直接
检测case语句。每个case语句都必须是一个面向Channel的操作。
*************************************************/

/************************************************
select可以在接收端进行监听，当然select也可以用于发送端，如
果有其他Goroutine在运行，还可以使用空select{}阻塞main()
函数，避免进程终止。
1.当所有被监听Channel中都无数据时，Select会一直等到其中一个数据为止。
2.当所有被监听Channel中都无数据时，且default字句存在时，则default子句会被执行。
3.当多个被监听Channel中都有数据时，Select会随机选择一个case执行。
4.如果想持续监听多个Channel，需要使用for语句协助。
************************************************/