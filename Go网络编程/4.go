//获取主机地址信息
package main

import (
	"fmt"
	"net"
	"os"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Fprintf(os.Stderr, "Usage:%s hostname\n", os.Args[0])
		os.Exit(1)
	}
	name := os.Args[1]
	addrs, err := net.LookupHost(name)
	if err != nil {
		fmt.Println("LookupHost error:", err.Error())
		os.Exit(1)
	}
	for _, s := range addrs {
		fmt.Println(s)
	}
	os.Exit(0)
}

/****************************************************************
函数ResolveIPAddr()虽然可以利用主机名获取一个主机IP地址，但是大多
数网络中的主机都拥有多个IP地址，另外一台主机也可以由多个名字。
可以使用LookupHost()函数查询主机更为详细的信息。
func LookupHost(host string) (addrs []string,err error)
在调用LookupHost()函数时，参数host是字符串型的主机名，调用成功后
函数返回数组格式的主机地址列表，否则返回一个错误类型。
***************************************************************/
