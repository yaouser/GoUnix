//IP Client端设计
package main

import (
	"fmt"
	"net"
	"os"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Fprintf(os.Stderr, "Usage:%s host", os.Args[0])
	}
	service := os.Args[1]
	lAddr, err := net.ResolveTCPAddr("ip4", service)
	checkError(err)
	name, err := os.Hostname()
	checkError(err)
	rAddr, err := net.ResolveIPAddr("ip4", name)
	checkError(err)
	conn, err := net.DialIP("ip4:ip", lAddr, rAddr)
	checkError(err)
	_, err = conn.WriteToIP([]byte("Hello Server!"), rAddr)
	checkError(err)
	var buf [512]byte
	n, addr, err := conn.ReadFromIP(buf[0:])
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

/*****************************************************************
在Go语言的IP网络编程中，客户机工作过程如下：
1.IP客户机在获取了服务器的网络地址之后，可以调用DialIP()函数向服务
器发出连接请求，如果请求成功会返回IPConn对象。
2.如果连接成功，客户机可以直接调用IPConn对象的ReadFromIP()方法或
WriteToIP()方法，与服务器进行通信活动。
3.通信完成后，客户机调用Close()方法关闭IPConn连接，断开通信链路。
函数DialIP()原型定义如下：
func DialIP(netProto string,laddr,raddr *IPAddr)(*IPConn,error)
******************************************************************/
