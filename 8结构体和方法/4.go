//嵌入式结构用作字段
package main

import(
	"fmt"
)

type student struct{
	id int
	name string
	sex bool
	class string
	birthday struct{
		//使用嵌入式结构定义birthday字段
		year int
		month int
		day int
	}
}

func main(){
	stu:=new(student)
	stu.name="李明"
	stu.birthday.year=1995
	fmt.Println(stu)
}