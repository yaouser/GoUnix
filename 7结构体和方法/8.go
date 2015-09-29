//匿名字段的访问
package main

import(
	"fmt"
)
type people struct{
	name string
	sex bool
}

type teacher struct{
	people     //匿名字段
	department string
}

type departmentHead struct{
	teacher   //有两层嵌入的匿名字段
	college string
}

func main(){
	head:=departmentHead{}
	head.name="郑智"
	head.department="Computer Science"
	head.college="Yale University"
	fmt.Println(head)
}
/************************************
无论嵌入多少层，匿名字段中的成员访问方法和正常
字段一样。Go语言中，就是利用匿名字段的这种访问
特性，实现字段的继承。
************************************/