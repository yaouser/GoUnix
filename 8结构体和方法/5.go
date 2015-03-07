//嵌入式结构体直接定义结构体变量以及new操作
package main

import(
	"fmt"
)

func main(){
	//使用嵌入式结构直接定义结构体变量
	var usr1=struct{
		id int
		name string
		password string
	}{name:"张三"}
	fmt.Println(usr1)
	//使用new()函数创建嵌入式结构对象
	var usr2=new(struct{
		id int
		name string
		password string
	})
	usr2.name="李四"
	fmt.Println(usr2)
}