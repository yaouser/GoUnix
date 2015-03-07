//结构体的定义与简单操作
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
	//标准方式定义
	var stu1 student
	//简写方式定义
	stu2:=student{}
	stu1.name="李明"
	stu2.name="张衡"
	stu1.age=18
	stu2.age=19
	fmt.Println(stu1)
	fmt.Println(stu2)
}
/*******************************************
在Go语言中，可以使用关键词"type"定义新的数据类型，每个
结构体中属于一个新的数据类型，结构体中的数据项，通常被称
为字段（field），在定义结构体必须使用type关键字和关键
字struct。
type 结构体名 struct {
	field1 类型
	.....
}
结构体中只有首写字母为大写的结构体和属性才能在包外被访问。
结构体中的预留字段：(在Go语言中，预留字段使用"_"来命名)
type people struct{
	name string
	age int
	_ string
}
在Go语言中，结构体中本身也是一种数据类型，所以也可以使用
结构体作为字段的类型。
type date string{
	year int
	month int
	day int
}
type student struct{
	id int
	name string
	sex bool
	class string
	birthday date
}
结构体变量的声明和声明普通变量一样，使用var来定义。
var stu student
在进行结构体变量声明时，通常使用":="简写的方式。
stu:=student{}
结构体中字段的访问使用"."操作符。
stu.name="李明"
stu.age=18
*******************************************/

/*******************************************
Go语言规定：int型的类型零值为0，bool型的类型零值为
false,string型的类型零值为空字符串。
*******************************************/