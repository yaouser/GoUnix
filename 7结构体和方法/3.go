//结构体变量，对象的初始化
package main

import(
	"fmt"
)

type student struct{
	id int
	name string
	sex bool
	age int
	class string
	
}

func main(){
	//结构体中的字段全部初始化（可以不写字段名）
	stu1:=student{13001,"李明",false,19,"网络01"}  //结构体变量
	stu2:=&student{13002,"张晓",true,18,"网络01"} //结构体对象
	//结构体中的字段部分初始化
	stu3:=student{name:"王乐",age:19}             //结构体变量
	stu4:=&student{name:"赵琼",id:13003}          //结构体对象
	fmt.Println(stu1)
	fmt.Println(stu2)
	fmt.Println(stu3)
	fmt.Println(stu4)
}
/**********************************************************
在Go语言中，相同类型的结构体之间可以直接使用"="操作符进行复制赋值。
type user struct{
	id int
	name string
}
func main(){
	user1:=user{100,"张三"}
	fmt.Println(user1)
	var user2* user=new(user)
	*user2=user1
	fmt.Println(user2)
}
Go语言中的结构体除了赋值操作符"="，还支持关系操作符"=="和"!="，但是不
支持>,<等比较操作符。
func main(){
	user1:=user{100,"张三"}
	user2:=user{200,"李四"}
	fmt.Println(user1==user2)
	fmt.Println(user1!=user2)
}