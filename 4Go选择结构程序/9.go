//fallthrough语句的使用
package main

import(
	"fmt"
)

func main(){
	var i int
	//通过scanf()函数，格式化从键盘输入变量i的值
	fmt.Scanf("%d",&i)
	switch i {
		case 0:
			fmt.Println("0")
		case 1:
			fmt.Println("1")
		case 2:
			fmt.Println("2")
			fallthrough         //继续执行下一分支
		case 3:
			fmt.Println("3")
		default:
			fmt.Println("4")
	}
}