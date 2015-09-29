//字典的初始化于创建
package main

import(
	"fmt"
)

func main(){
	//字典声明后如果没初始化值为nil
	var map1 map[string]int
	fmt.Println(map1)
	//未初始化map1["Key1"]=1语句编译错误
	//使用make创建后就可以给字典增加数据项了
	map1=make(map[string]int)
	map1["key1"]=1
	fmt.Println(map1)
	//使用”{}“初始化后也可以给字典增加数据项
	var map2=map[string]int{}
	map2["key2"]=2
	fmt.Println(map2)
}
/*************************************
Map是一种特殊的数据结构，它由一对无序的数据项
组成，Map通过把键映射到值来进行访问，这加快数
据的查找速度，所以Map也被称作字典或哈希表。
字典声明时，除了要定义字典名，还要指定”键“类型
和”值“类型，键类型要用[]括起来。
var mapName map[keyType]valueType
************************************/