//TCP连接地址
package main

import (
	"fmt"
	"net"
	"os"
)

func main() {
	if len(os.Args) != 3 {
		fmt.Fprintf(os.Stderr, "Usage:%s networkType addr\n", os.Args[0])
		os.Exit(1)
	}
	networkType := os.Args[1]
	addr := os.Args[2]
	tcpAddr, err := net.ResolveTCPAddr(networkType, addr)
	if err != nil {
		fmt.Println("ResolveTCPAddr error:", err.Error())
		os.Exit(1)
	}
	fmt.Println("The network type is:", tcpAddr.Network())
	fmt.Println("The IP address is", tcpAddr.String())
	os.Exit(0)
}

/**********************************************************************
Go语言标准库对传统Socket网络编程的过程进行了封装，无论期望使用什么协议
建立什么形式的连接，都只需要调用Dial()函数即可。Dial()函数定义如下：
func Dial(net,addr string)(Conn,error)
在Dial()函数中，参数net是网络协议名，参数addr是IP地址或域名，而端口号以
":"的形式跟随在地址或域名的后面，端口是可选的。Dial()函数如果调用成功，
则返回一个连接对象。否则，返回一个错误类型。如下例：
conn,err:=net.Dial("tcp","192.168.0.1:5000")
conn,err:=net.Dial("udp","192.168.0.2:5001")
conn,err:=net.Dial("udp","www.baidu.com")
conn,err:=net.Dial("ip4:1","119.75.218.77")
在成功建立连接后，就可以进行数据的发送和接收了。发送数据时，可以使用
conn对象的Write()成员方法，在接收数据时，可以使用conn对象的Read()成员
方法。
*********************************************************************/

/********************************************************************
在进行TCP网络编程时，服务器或客户机的地址使用TCPAddr地址结构体表示，
TCPAddr包含两个字段：IP和Port，形式如下：
type TCPAddr struct {
	IP IP
	Port int
}
函数ResolveTCPAddr()可以把网络地址转换为TCPAddr地址结构，该函数原型定义
如下：func ResolveTCPAddr(net,addr string)(*TCPAddr,error)
在调用函数ResolveTCPAddr()时，参数net是网络协议名，参数addr是IP地址或
域名，如果是IPv6地址则必须使用[]括起来。还有一种特例，就是对于HTTP服务
器，当主机地址为本地测试地址时（127.0.0.1），可以直接使用端口作为TCP连
接地址，形如":80"。
在进行TCP网络编程时，客户机和服务器之间是通过TCPConn对象实现连接的，
TCPConn是Conn接口的实现。TCPConn对象绑定了服务器的网络协议和地址信息。
通过TCPConn连接对象，可以实现客户机和服务器间的全双工通信。可以通过
TCPConn对象的Read()方法和Write()方法，在服务器和客户机之间发送和接收数
据。另外，这两个方法的执行都会引起阻塞。
*********************************************************************/
