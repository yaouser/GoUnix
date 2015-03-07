//异步Channel
package main

import(
	"fmt"
	"time"
)

func Worker(sem chan int,lock chan bool,id int){
	sem<-1						//相当于对同步信号量进行P原语操作
	fmt.Println(time.Now().Format("04:05"),id)
	time.Sleep(1*time.Second)
	<-sem						//相当于对同步信号量进行V原语操作
	if id==9{
		lock<-true
	}
}

func main(){
	ch:=make(chan int,2)
	lock:=make(chan bool)
	for i:=0;i<10;i++ {
		go Worker(ch,lock,i)
	}
	<-lock
}
/***********************************************
创建的不带Buffer的Channel，这种只能用于传递单个数据的应用
场合，对于在Goroutine间传输大量数据的应用，可以使用异步通
道，从而达到消息队列的效果。
异步Channel，就是给Channel设定一个Buffer值。在Buffer未
写满的情况下，不阻塞发送操作；在Buffer未读完之前，不阻塞接收
操作。这里的Buffer是指被缓冲的数据对象的数量，而不是内存。
要创建一个带Buffer的Channel，只需要在调用make()函数时，
将缓冲区的大小作为第二个参数传入即可。例如：
ch:=make(chan int,1024)
***********************************************/