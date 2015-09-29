//使用Scanln()函数从标准输入设备录入数据
package main

import(
	"fmt"
)

func main(){
	var a int
	var f float32
	var str string
	fmt.Println("准备录入数据：")
	fmt.Scanln(&a,&f,&str)
	fmt.Println("输出录入结果：")
	fmt.Println(a,f,str)
}