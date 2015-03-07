//TCP Server端设计(使用Goroutine实现并发服务器)
package main

import (
	"fmt"
	"net"
	"os"
)

func main() {
	service := "5000"
	tcpAddr, err := net.ResolveTCPAddr("tcp", service)
	checkError(err)
	listener, err := net.ListenTCP("tcp", tcpAddr)
	checkError(err)
	for {
		conn, err := listener.Accept()
		if err != nil {
			continue
		}
		go handleClient(conn) //创建Goroutine
	}
}

func handleClient(conn net.Conn) {
	defer conn.Close() //逆序调用Close()保证连接能正常关闭
	var buf [512]byte
	for {
		n, err := conn.Read(buf[0:])
		if err != nil {
			return
		}
		rAddr := conn.RemoteAddr()
		fmt.Println("Receive from client", rAddr.String(), string(buf[0:n]))
		_, err2 := conn.Write([]byte("Welcome client!"))
		if err2 != nil {
			return
		}
	}
}

func checkError(err error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "Fatal error:%s", err.Error())
		os.Exit(1)
	}
}

/************************************************************************
服务器设计采用循环服务器设计模式，这种服务器设计简单单缺陷明显。因为这种
服务一旦启动，就一直阻塞监听客户机的连接请求，直至服务器关闭。所以，循环
服务器很耗费系统资源。解决问题的方法是采用并发服务器模式，在这种模式中，
对每一个客户的连接请求，服务器都会创建一个新的进程，线程或者协程进行响应，
而服务器还可以去处理其他任务。所以使用Goroutine（即协程是一种比线程更轻量
级的任务单位）来实现并发服务器的设计。
************************************************************************/
