//UDP Client端设计
package main

import (
	"fmt"
	"net"
	"os"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Fprintf(os.Stderr, "Usage:%s host:port", os.Args[0])
	}
	service := os.Args[1]
	udpAddr, err := net.ResolveUDPAddr("udp", service)
	checkError(err)
	conn, err := net.DialUDP("udp", nil, udpAddr)
	checkError(err)
	_, err = conn.WriteTo([]byte("Hello Server!"))
	checkError(err)
	var buf [512]byte
	n, addr, err := conn.ReadFromUDP(buf[0:])
	checkError(err)
	fmt.Println("Reply form server", addr.String(), string(buf[0:n]))
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
在Go语言UDP网络编程中客户机工作过程如下：
1.UDP客户机在获取了服务器的服务端口号和服务地址之后，可以调用DialUDP()
函数向服务器发出连接请求，如果请求成功会返回UDPConn对象。
2.客户机可以直接调用UDPConn对象的ReadFromUDP()方法或WriteToUDP()方法，
与服务器进行通信活动。
3.通信完成后，客户机调用Close()方法关闭UDPConn连接，断开通信链路。
函数DialUDP()原型定义如下：
func DialUDP(net string,laddr,raddr *UDPAddr)(*UDPConn,error)
*******************************************************************/
