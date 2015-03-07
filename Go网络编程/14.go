//IP Server端设计
package main

import (
	"fmt"
	"net"
	"os"
)

func main() {
	name, err := os.Hostname()
	checkError(err)
	ipAddr, err := net.ResolveIPAddr("ipv4", name)
	checkError(err)
	conn, err := net.ListenIP("ip4:ip", ipAddr)
	checkError(err)
	for {
		handleClient(conn)
	}
}

func handleClient(conn *net.IPConn) {
	var buf [512]byte
	n, addr, err := conn.ReadFromIP(buf[0:])
	if err != nil {
		return
	}
	fmt.Println("Receive from client", addr.string(), string(buf[0:n]))
	conn.WriteToIP([]byte("Welcome Client!"), addr)
}
func checkError(err error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "Fatal error:%s", err.Error())
		os.Exit(1)
	}
}

/********************************************************************
IP是Internet网络层的核心协议，他是一种不可靠的，无连接的通信协议。TCP,
UDP都是在IP的基础上实现的通信协议，所以IP属于一种底层协议，他可以直接
对网络数据包进行处理。另外，通过IP用户实现自己的网络服务协议。
函数ResolveIPAddr()可以把网络地址转换为IPAddr地址结构。
IPAddr地址对象还有两个方法：Network()和String()
在进行IP网络编程时，客户机和服务器之间是通过IPConn对象实现连接的，
IPConn是Conn接口的实现。IPConn对象提供了ReadFromIP()方法和WriteToIP()
方法用于在客户机和服务器之间进行数据收发操作。
由于工作在网络层，IP服务器并不需要在一个指定的端口上和客户积极进行通信
连接，IP服务器的工作过程如下：
1.IP服务器使用指定的协议簇和协议，调用ListenIP()函数创建一个IPConn连
接对象，并在该对象和客户机间建立不可靠连接。
2.如果服务器和某个客户机建立了IPConn连接，就可以使用该对象的ReadFormIP()
方法和WriteToIP()方法相互通信了。
3.如果通信结束，服务器还可以调用Close()方法关闭IPConn连接。
函数ListenIP原型定义如下：
func ListenIP(netProto string,laddr *IPAddr)(*IPConn,error)
*********************************************************************/
