//IP地址类型
package main

import (
	"fmt"
	"net"
	"os"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Fprintf(os.Stderr, "Usage:%s ip.addr\n", os.Args[0])
		os.Exit(1)
	}
	addr := os.Args[1]
	ip := net.ParseIP(addr)
	if ip == nil {
		fmt.Println("Invalid address")
	} else {
		fmt.Println("The address is", ip.String())
	}
	os.Exit(0)
}

/********************************************************************
主机地址是网络通信最重要的数据之一，net包中定义了三种类型的主机地址
数据类型：IP,IPMask和IPAddr,它们分别用来存储协议相关的网络地址。
在net包中，IP地址类型被定义成一个byte型数组，即若干个8为组。
在net包中，有几个函数可以将IP地址类型作为函数的返回类型，如ParseIP()。
ParseIP函数的主要作用是分析IP地址的合法性，如果是一个合法的IP地址，
ParseIP函数将返回一个IP地址对象。如果是一个非法IP地址，ParseIP()函数
将返回nil。还可以使用IP对象的String()方法将IP地址转换成字符串格式。
*******************************************************************/
