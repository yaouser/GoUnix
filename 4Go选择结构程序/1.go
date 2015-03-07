//if语句的使用，条件表达式为Boolean值
package main

import(
	"fmt"
)

func main(){
	//条件为TRUE，执行{}中的语句
	if true {
		fmt.Println("The condition is true!")
	}
	if false{
		fmt.Println("The condition is false!")
	}
	//if语句外的后续语句
	fmt.Println("OVER")
}