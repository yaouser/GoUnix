//TCP Client端设计
package main

import (
	"fmt"
	"net"
	"os"
)

func main() {
	var buf [512]byte
	if len(os.Args) != 2 {
		fmt.Fprintf(os.Stderr, "Usage: %s host:port", os.Args[0])
	}
	service := os.Args[1]
	tcpAddr, err := net.ResolveTCPAddr("tcp", service)
	checkError(err)
	conn, err := net.DialTCP("tcp", nil, tcpAddr)
	checkError(err)
	rAddr := conn.RemoteAddr()
	n, err := conn.Write([]byte("Hello server!"))
	checkError(err)
	n, err = conn.Read(buf[0:])
	checkError(err)
	fmt.Println("Reply from server", rAddr.String(), string(buf[0:n]))
	conn.Close()
	os.Exit(0)
}

func checkError(err error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "Fatal error:%s", err.Error())
		os.Exit(1)
	}
}

/*******************************************************************
在Go语言的网络编程中，TCP客户机的工作过程如下：
1.TCP客户机在获取了服务器的服务端口号和服务地址之后，可以调用DailTCP()
函数向服务器发出连接请求，如果请求成功会返回TCPConn对象。
2.客户机调用TCPConn对象的Read()或Write()方法，与服务器进行通信活动。
3.通信完成后，客户机调用Close()方法关闭连接，断开通信链路。
DialTCP()函数原型定义如下：
Func DialTCP(net string,laddr,raddr *TCPAddr)(*TCPConn,error)
此函数的参数laddr是本地主机地址，可以设为nil。参数raddr是对方主机地址，
必须指定能省略。函数调用成功后，返回TCPConn()对象，调用失败，返回一个
错误类型。
********************************************************************/
