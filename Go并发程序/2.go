//Channel迭代操作
package main

import(
	"fmt"
)

func main(){
	ch:=make(chan int)
	go func(){  //匿名函数
		for i:=0;i<5;i++{
			ch<-i
		}
		close(ch)
	}()
	for value:=range ch{ //range迭代器
		fmt.Println(value)
	}
}
/*****************************************
关闭Channel非常简单，使用Go语言提供的内置函数close()
即可。关闭操作如下：close(chanName)。
关闭之后，往往需要判断Channel是否关闭，这时可以在读
取Channel的时候使用多重返回值的方式，例如：
value,ok:=<-ch
这时只需看第二个bool返回值即可。如果返回值是false则
表示Channel已被关闭，否则主函数还要继续阻塞接收或发送。
*******************************************
对Channel的读取操作还可以使用range迭代器来完成，
range操作直至Channel关闭（Close）方才终止循环。
另外，在Go语言中，还经常把创建Goroutine和Channel
的工作放在一个匿名函数中来完成。
******************************************/

/*****************************************
func main(){
	ch:=make(chan int)
	go func(){
		for i:=0;i<5;i++{
			ch<-i
		}
		close(ch)
	}()
	for {
		if value,ok:=<-ch;ok{
			fmt.Println(value)
		}else{
			break
		}
	}
}
这里的close函数不是必须的，但如果接收端使用range或
者循环接收数据。就必须调用close，否则会出现死锁错误。
******************************************/