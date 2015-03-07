//DNS域名解析
package main

import (
	"fmt"
	"net"
	"os"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Fprintf(os.Stderr, "Usage:%s hostname\n", os.Args[0])
		fmt.Println("Usage:", os.Args[0], "hostname")
		os.Exit(1)
	}
	name := os.Args[1]
	addr, err := net.ResolveIPAddr("ip", name)
	if err != nil {
		fmt.Println("Resolvtion error", err.Error())
		os.Exit(1)
	}
	fmt.Println("Resolved address is", addr.String())
	os.Exit(0)
}

/******************************************************************
在net包中，许多函数和方法调用后返回的是一个指向IPAddr结构体的指针，
IPAddr结构体的主要作用是用于域名解析服务（DNS）。函数ResolveIPAddr()
可以通过主机名解析主机网络地址。
func ResolveIPAddr(net,addr string) (*IPAddr,error)
ResolveIPAddr()函数调用成功后返回指向IPAddr结构体的指针，调用失败返
回错误类型error。参数net表示网络类型，参数addr可以是IP地址或域名。
******************************************************************/
