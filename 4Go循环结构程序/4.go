//简单键盘文本输入程序
package main

import(
	"fmt"
	"bufio"
	"os"
)

func main(){
	var c byte        //存放每个按键ASCII码
	var str string    //存放输入字符串流
	//初始化标准输入设备
	r:=bufio.NewReader(os.Stdin)
	//初始化标准输出设备
	w:=bufio.NewWriter(os.Stdout)
	for i:=0; ;i++{
		c,_=r.ReadByte()
		if c==10 {
			break    //如果是回车停止接受退出循环
		}else{
			//将接受到的字符写入标准输出设备
			w.WriteByte(c)
			w.Flush()
			//生成字符串
			str+=string(c)
		}
	}
	fmt.Printf("\n")
	fmt.Println(str)
}