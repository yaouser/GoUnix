//其他类型作为匿名字段
package main

import(
	"fmt"
)
//自定义类型
type skills []string

type people struct{
	name string
	sex bool
}

type teacher struct {
	people      	//struct作为匿名字段
	skills  		//自定义类型作为匿名字段
	int    			//内置数据类型作为匿名字段
	department string
}

func main(){
	teacher1:=teacher{}
	teacher1.name="郑智"
	teacher1.skills=append(teacher1.skills,"Computer","Golang")
	teacher1.int=100
	teacher1.department="Computer Science"
	fmt.Println(teacher1)
}