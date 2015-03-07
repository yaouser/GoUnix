//直接和使用make创建切片
package main

import(
	"fmt"
)

func main(){
	//直接创建
	var slice1=[]int{1,2,3,4,5}
	//使用make创建
	var slice2=make([]int,5)
	//使用make创建并预留元素存储空间
	var slice3=make([]int,5,10)
	fmt.Printf("len=%2d cap=%2d %v\n",len(slice1),cap(slice1),slice1)
	fmt.Printf("len=%2d cap=%2d %v\n",len(slice2),cap(slice2),slice2)
	fmt.Printf("len=%2d cap=%2d %v\n",len(slice3),cap(slice3),slice3)
}