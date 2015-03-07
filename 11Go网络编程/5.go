//获取主机正式名
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
	cname, err := net.LookupCNAME(name)
	if err != nil {
		fmt.Println("LookupCNAME error:", err.Error())
		os.Exit(1)
	}
	fmt.Println("The host caonical name is:", cname)
	os.Exit(0)
}

/***************************************************************
在主机的多个名字信息中，有一个会被标记为“canonical”，即主机正式
名。如果想查询主机正式名，可以使用函数LookupCNAME()来完成。
***************************************************************/
