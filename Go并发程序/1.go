//Goroutine和Channel使用举例
package main

import(
	"fmt"
	"math/rand"
)

func Test(ch chan int){
	ch<-rand.Int()//向Channel中写入一个随机数
	fmt.Println("Go...")
}

func main(){
	chs:=make([]chan int,10)
	for i:=0;i<10;i++{
		chs[i]=make(chan int)
		go Test(chs[i])  //启动10个Goroutine
	}
	for _,ch:=range chs {
		value:=<-ch  //阻塞等待退出信号
		fmt.Println(value)
	}
}
/********************************************
Channel也是一种引用类型。
Channel的一般声明形式如下：
var chanName chan ElementType
例如：var ch chan int
Channel除了可以传递基本类型的数据，还可以作为数组，切片
或字典的元素。例如：
var chs [10]chan int
该例中，chs是一个包含10个可传递int类型数据的Channel。
还可以使用make()函数直接声明并初始化Channel，例如：
ch:=make(chan int)
**********Channel中数据的接收和发送*************
将一个数据发送（写入）至Channel的语法如下：
ch<-value
**向Channel写入数据通常会导致程序阻塞（Block），直到
**有其他Goroutine从这个Channel种读取数据。
从Channel中接收（读取）数据的语法如下：
value:=<-ch
**如果Channel之前没有写入数据，那么从Channel中读取数
**据也会导致阻塞，直到Channel中被写入数据为止。
*********************************************
Channel是进程内的通信方式，在传递对象的过程和调用函数时
的参数传递行为比较一致，当然也可以传递指针等。如果需要跨
进程进行通信，一般建议使用分布式系统的方式，比如使用网络
套接字（Socket）或者HTTP等通信协议。
在使用Go语言开发应用程序时，经常回遇到需要实现条件等待问
题，这时就可以使用Channel进行处理。
********************************************/