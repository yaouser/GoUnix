//if语句的使用，关系表达式或逻辑表达式的值作为判断条件
package main

import(
	"fmt"
)

func main(){
	var a,b int
	//通过scanln()函数，从键盘输入a,b的值
	fmt.Scanln(&a,&b)
	//关系表达式的值作为判断条件
	if a>b {
		fmt.Println("a>b")
	}
	//逻辑表达式的值作为判断条件
	if(a!=0)&&(a>0){
		fmt.Println("a is positive number")
	}
	if(a!=0)&&(a<0){
		fmt.Println("a is negative number")
	}
	//if语句外的后续语句
	fmt.Println("OVER")
}