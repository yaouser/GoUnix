//使用new()函数创建结构体对象
package main

import(
	"fmt"
)

type date struct{
	year int
	month int
	day int
}

type student struct{
	Id int
	name string
	sex bool
	class string
	birthday date  //结构体作为字段
	_ string       //预留字段
}

func main(){
	stu:=new(student)
	stu.name="李明"
	stu.birthday.year=1995
	fmt.Println(stu)
}