//使用循环嵌套打印九九乘法表
package main

import(
	"fmt"
)

func main(){
	var row,col int
	for row=1;row<9;row++{
		for col=1;col<=row;col++{
			fmt.Printf("%d*%d=%d ",row,col,row*col)
		}
		fmt.Printf("\n")
	}
}