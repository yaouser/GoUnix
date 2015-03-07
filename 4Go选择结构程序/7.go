//switch语句实现简单计算器
package main

import(
	"fmt"
)

func main(){
	//操作数a,b
	var a,b int
	//算术运算符+-*/
	var op byte
	//通过scanf（）函数，格式化从键盘输入操作数和云算数
	fmt.Scanf("%d %c %d",&a,&op,&b)
	switch op {
		case '+':
			fmt.Println(a+b)  //加法
		case '-':
			fmt.Println(a-b)  //减法
		case '*':
			fmt.Println(a*b)  //乘法
		default:
			fmt.Println(a/b)  //除法
	}
}