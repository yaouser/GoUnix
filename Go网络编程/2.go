//子网掩码地址
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
	dotaddr := os.Args[1]
	addr := net.ParseIP(dotaddr)
	if addr == nil {
		fmt.Println("Invalid address")
	}
	mask := addr.DefaultMask()
	fmt.Println("Subnet mask is:", mask.String())
	network := addr.Mask(mask)
	fmt.Println("Network address is:", network.String())
	ones, bits := mask.Size()
	fmt.Println("Mask bits:", ones, "Total bits:", bits)
	os.Exit(0)
}

/*****************************************************************
在Go语言中，net包中还提供了IPMast地址类型，也是byte型数组。函数
IPv4Mask()可以通过一个32位IPv4地址生成子网掩码地址，调用成功后返回
一个4字节的十六进制子网掩码地址。还可以使用主机地址对象的DefaultMask
方法获取主机默认子网掩码地址。要注意的是，只有IPv4才有默认子网掩码，
如果不是IPv4地址，DefaultMask()方法将返回nil。这两种获取的子网掩码
的地址都是十六进制格式的。主机地址对象还有一个Mask()方法，执行Mask
方法后，会返回IP地址与子网掩码地址相与的结果，即是主机所处的网络的
“网络地址”。
*****************************************************************/
