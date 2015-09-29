//Method的继承
package main

import(
	"fmt"
)

type people struct{
	name string
	phone string
}
type teacher struct{
	people
	department string
}
type student struct{
	people
	school string
}

func (r people) sayHi(){
	fmt.Printf("Hi,I'm %s you can call me on %s.\n",r.name,r.phone)
}

func main(){
	teacher1:=teacher{people{"郑智","010-22002"},"Computer Science"}
	student1:=student{people{"李明","010-11001"},"Yale University"}
	teacher1.sayHi()
	student1.sayHi()
}
/****************************************
people作为teacher和student的匿名字段，sayHi是
为people实现的一个方法。作为外层的结构，teacher
和student同时也继承了people所拥有的方法，所以在
teacher和student的对象中可以直接调用方法sayHi()
****************************************/