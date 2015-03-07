//TCPServer端设计
package main 

import (
	"fmt"
	"net"
	"os"
)

func main() {
	service:=":5000"
	tcpAddr,err:=net.ResolveTCPAddr("tcp", service)
	checkError(err)
	listener,err:=net.ListenTCP("tcp", tcpAddr)
	checkError(err)
	for {
		conn,err:=listener.Accept()
		if err!=nil {
			continue
		}
		handleClient(conn)
		conn.Close()
	}
}

func handleClient(conn net.Conn) {
	var buf [512]byte
	for {
		n,err :=conn.Read(but[0:])
		if err != nil {
			return
		}
		rAddr:=conn.RemoteAddr()
		fmt.Println("Receive from client",rAddr.String(),string(buf[0:n]))
		_,err2:=conn.Write([]byte("Welcome client!"))
		if err2:=nil {
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

/*************************************************************************
在Go语言的网络编程中，TCP服务器的工作过程如下：
1.TCP服务器首先注册一个公知端口，然后调用ListenTCP()函数在这个端口上创建一
个TCPListener监听对象，并在该对象上监听客户机的连接请求。
2.启用TCPListener对象的Accept()方法接收客户机的连接请求，并返回一个协议相
关的Conn对象，这里就是TCPConn()对象。
3.如果返回了一个新的TCPConn对象，服务器就可以调用该对象的Read()方法接收客户
客户机发来的数据，或者调用Write()方法向客户机发送数据了。
服务器和客户机的通信连接建立成功后，就可以使用Read()和Write()方法收发数据。
在通信过程中，如果还想获取通信双方的地址信息，可以使用LocalAddr()方法和
RemoteAddr()方法来完成。LocalAddr()方法会返回本地主机地址，RemoteAddr()方
法返回远端主机地址。
************************************************************************/