//continue语句的使用
package main

import(
	"fmt"
)

func main(){
	var n int
	for n=10;n<30;n++{
		if n%3!=0{
			continue //不能被3整除，终止本次循环，返回条件测试
		}
		fmt.Printf("%d",n)
		fmt.Printf("\n")
	}
}