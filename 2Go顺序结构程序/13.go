//使用Scanf()函数格式化录入数据
package main

import(
	"fmt"
)

func main(){
	var a int
	var f float32
	var str string
	fmt.Println("准备录入数据：")
	fmt.Scanf("%d,%f,%s"&a,&f,&str)
	fmt.Println("输出录入结果：")
	fmt.Println(a,f,str)
}