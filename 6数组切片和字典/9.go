//使用append函数增加切片元素
package main

import(
	"fmt"
)

func main(){
	//使用make创建切片len=3,cap=6
	var slice1=make([]int,3,6)
	//使用append给切片增加元素且未超出cap
	slice2:=append(slice1,1,2,3)
	//使用append给切片增加元素并且超出cap
	slice3:=append(slice1,1,2,3,4)
	fmt.Printf("len=%d cap=%d %v\n",len(slice1),cap(slice1),slice1)
	fmt.Printf("len=%d cap=%d %v\n",len(slice2),cap(slice2),slice2)
	fmt.printf("len=%d cap=%d %v\n",len(slice3),cap(slice3),slice3)
}
/*****************************************************
append只是往切片尾部增加元素，如果不超出容量，增加新元素不会改变
原来切片的属性，如果超出容量，则分配一个是原来容量2倍的新的内存块，
再复制数据到新的内存块。
****************************************************/