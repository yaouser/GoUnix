//查询服务端口号
package main

import (
	"fmt"
	"net"
	"os"
)

func main() {
	if len(os.Args) != 3 {
		fmt.Fprintf(os.Stderr, "Usage:%s networkType service\n", os.Args[0])
		os.Exit(1)
	}
	networkType := os.Args[1]
	service := os.Args[2]
	port, err := net.LookupPort(networkType, service)
	if err != nil {
		fmt.Println("LookupPort error:", err.Error())
		os.Exit(1)
	}
	fmt.Println("The service port is:", port)
	os.Exit(0)
}

/**************************************************************************
同样地，在一台主机上也可以安装多个服务进行，这些服务进程分别提供不同的网络
服务，比如Web服务，FTP服务，DNS服务等。可以通过利用服务端口号来区分不同的
网络服务。
函数LookupPort()可以用来查询服务端口号，如下：
func LookupPort(network,service string) (port int,err error)
在调用函数LookupPort()时，参数network是网络类型，参数service是服务类型，比如
WWW服务，FTP服务，telnet服务等。函数调用成功返回服务端口号，否则返回一个错
误类型。
*************************************************************************/
