//切片元素的遍历
package main

import(
	"fmt"
)

func main(){
	var slice1=[]int{1,2,3,4,5}
	//使用下标访问切片元素
	for i:=0;i<=4;i++{
		fmt.Printf("slice1[%d]=%d",i,slice1[i])
	}
	fmt.Printf("\n")
	//使用range遍历所有元素
	for i,v:=range slice1{
		fmt.Printf("slice1[%d]=%d",i,v)
	}
}
/*************************************************
切片元素的遍历和数组元素的遍历一样，要通过元素的下标访问，另外
也可以使用关键字range遍历所有切片元素。
************************************************/