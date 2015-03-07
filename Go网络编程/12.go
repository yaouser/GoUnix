//UDP Server端设计
package main 

import (
	"fmt"
	"net"
	"os"
)

func main() {
	service:=":5001"
	udpAddr,err:=net.ResolveUDPAddr("udp", service)
	checkError(err)
	conn.err:=net.ListenUDP("udp", udpAddr)
	checkError(err)
	for {
		handleClient(conn)
	}
}
func handleClient(conn *net.UDPConn) {
	var buf [512]byte
	n,addr,err:=conn.ReadFromUDP(buf[0:])
	if err!=nil {
		return
	}
	fmt.Println("Receive from client",addr.String(),string(buf[0:n]))
	conn.WriteToUDP([]byte("Welcome Client!"), addr)
}
func checkError(err error) {
	if err!=nil {
		fmt.Fprintf(os.Stderr, "Fatal error:%s", err.Error())
		os.Exit(1)
	}
}
/********************************************************************
UDP是不保证可靠性的通信协议，所以客户机和服务器之间只要建立连接，就可
以通信，而不用调用Accept()进行连接确认。
在进行UDP网络编程时，服务器或客户机的地址使用UDPAddr地址结构体表示，
函数ResolveUDPAddr()可以把网络地址转换为UDPAddr地址结构，另外，UDPAddr
地址对象还有两个方法：Network()和String(),Network()方法用于返回UDPAddr
地址对象的网络协议名，如"udp";String()方法可以将UDPAddr地址转换成字符
串形式。
在进行UDP网络编程时，客户机和服务器之间是通过UDPConn对象实现连接的，
UDPConn是Conn接口的实现。
UDPConn连接对象的通信并不能保证通信的可靠性和有序性，这些需要程序员来
处理。为此，TCPConn对象提供了ReadFromUDP()方法和WriteToUDP()方法，这
两个方法直接使用远端主机地址进行数据发送和接收，即便在链路失效的情况下，
通信操作都能正常进行。
在Go语言的UDP网络编程中服务器的工作如下：
1.UDP服务器首先注册一个公知端口，然后调用ListenUDP()函数在这个端口上创
建一个UDPConn连接对象，并在该对象上和客户机建立不可靠连接。
2.如果服务器和某个客户机建立了UDPConn连接，就可以使用该对象的ReadFromUDP()
方法和WriteToUDP()方法相互通信了。
3.不管上一次通信是否完成或正常，UDP服务器依然会接收下一次连接请求。
********************************************************************/