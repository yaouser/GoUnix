//Go中的内存分配new和make
package main

import(
	"fmt"
)

func main(){
	fmt.Println("Hello World")
}
/*****************************************
Go有两种分配内存的机制，分别是使用内置函数new()
和make()，他们所对应的类型不同。
new可以给一个值类型的数据分配内存，返回该类型的
内存指针，即是返回了一个指向新分配的类型为T的零
值的指针。
make用于给引用类型分配内存空间，比如像slice，
map，channel等。make返回的是一个对象，而非是一
个内存空间的指针，分多少内存由该对象的大小和个数
来决定。
*****************************************
func new(Type)* Type
func make(Type,size IntegerType) Type
*****************************************
var p* []int=new([]int)
var v []int=make([]int,10)
var v []int=make([]int,10,10)
****************************************/