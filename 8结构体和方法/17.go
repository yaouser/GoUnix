//Method的重写
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
func (s student) sayHi(){
	fmt.Printf("Hi,I'm %s.I study in %s,call me on %s.\n",s.name,s.school,s.phone)
}

func main(){
	teacher1:=teacher{people{"郑智","010-22002"},"Computer Science"}
	student1:=student{people{"李明","010-11001"},"Yale University"}
	teacher1.sayHi()
	student1.sayHi()
}