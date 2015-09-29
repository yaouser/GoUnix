//调用mymath包中的函数
package main

import(
	"fmt"
	//导入mymath包
	"mymath"
)

func main(){
	var a,b int=2,3
	add:=mymath.Add(a,b)
	fmt.Println(add)
	sub:=mymath.Sub(a,b)
	fmt.Println(sub)
	mult:=mymath.Mult(a,b)
	fmt.Println(mult)
	div:=mymath.Div(a,b)
	fmt.Println(div)
}