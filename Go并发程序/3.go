//单向Channel
package main

import(
	"fmt"
)
func Recv(ch<-chan int,lock chan<-bool){
	for value:=range ch{
		fmt.Println(value)
	}
	lock<-true
}
func Send(ch chan<-int){
	for i:=0;i<5;i++{
		ch<-i
	}
	close(ch)
}
func main(){
	ch:=make(chan int)  	//双向Channel可转换为单向Channel
	lock:=make(chan bool)
	go Recv(ch,lock)  		//只能从ch接收的Goroutine
	go Send(ch)				//只能向ch发送的Goroutine
	<-lock
}
/*******************************************
ch是一个正常的双向Channel，他可以转换为单向Channel。
函数send的参数为只写Channel，函数recv的第一个参数为
只读Channel，他可以接收send的写入的数据，然后打印出来。
Send()和recv()是并发的，所以两函数排序不影响执行结果。
********************************************/

/*******************************************
只能接收的Channel变量定义形式如下：
var chanName chan<-ElementType
只能发送的Channel变量定义形式如下：
var chanName<-chan ElementType
单向Channel可以由一个已定义的双向（正常）Channel转
换而来。例如：
ch:=make(chan int)
chRead:=<-chan int(ch)
chWrite:=chan<-int(ch)
********************************************/