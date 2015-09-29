//基于底层数组创建用户
package main

import(
	"fmt"
)

func main(){
	//先定义并初始化底层数组
	var array1=[10]int{1,2,3,4,5,6,7,8,9,10}
	//注意切片类型和底层数组类型保持一致
	var slice1 []int
	//部分引用的三种形式
	slice1=array1[:5]
	slice2:=array1[5:]
	slice3:=array1[4:7]
	//全部引用的三种形式
	slice4:=array1
	slice5:=array1[:]
	slice6:=array1[0:len(array1)]
	fmt.Println(slice1)
	fmt.Println(slice2)
	fmt.Println(slice3)
	fmt.Println(slice4)
	fmt.Println(slice5)
	fmt.Println(slice6)
}
/****************************************
切片的初始化方式（首先定义好数组，在初始化切片）
切片名=数组名[start:end]
直接创建切片，即在定义切片的同时初始化切片元素。
var slice1=[]int{1,2,3,4,5}
好像和初始化数组有些类似呀，只是切片的中括号中
没有定义数组的长度。
****************************************/

/****************************************
在Go中切片（slice）是数组的一个引用，它会生成一个
指向数组的指针，并通过切片长度关联到底层数组或者全部
元素。切片通常被用来实现变长数组，而且操作灵活。
****************************************/