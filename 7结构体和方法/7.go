//匿名字段的初始化
package main

import(
	"fmt"
)

type people struct{
	name string
	sex bool
}

type teacher struct{
	people //匿名字段
	department string
}

func main(){
	//匿名字段成员当做正常字段来赋值
	teacher1:=teacher{people{"郑智",false},"Computer Science"}
	fmt.Println(teacher1)
}