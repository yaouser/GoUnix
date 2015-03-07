//查询服务器
package main

import (
	"fmt"
	"net"
	"os"
)

func main() {
	if len(os.Args) != 4 {
		fmt.Fprintf(os.Stderr, "Usage:%s service proto name\n", os.Args[0])
		os.Exit(1)
	}
	service := os.Args[1]
	proto := os.Args[2]
	name := os.Args[3]
	cname, addrs, err := net.LookupSRV(service, proto, name)
	if err != nil {
		fmt.Println("LookupSRV error:", err.Error())
		os.Exit(1)
	}
	fmt.Println("The service canonical name is:", cname)
	for _, addr := range addrs {
		fmt.Println("The service addr list is:", addr.Target)
	}
	os.Exit(0)
}

/************************************************************************
SRV是DNS服务器的数据库中支持的一种资源记录的类型，它记录了哪台计算机提供
了哪个服务这么一个简单的信息。
使用函数LookupSRV()可以用来查询SRV信息。
func LookupSRV(service,proto,name string) (canme string,addrs []*SRV,err error)
在调用函数LookupSRV()时，参数service表示服务名，参数proto表示通信协议（
TCP或UDP），参数name是主机域名。函数调用成功会返回主机正式名和SRV地址列表，
否则返回错误类型。
**********************************************************************/
