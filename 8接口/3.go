//接口的转换
package main

import(
	"fmt"
)

type People struct{
	Name string
	Age int
}
type Student struct{
	People
	School string
}
type PeopleInfo interface{
	GetPeopleInfo()
}
type StudentInfo interface{
	GetPeopleInfo()
	GetStudentInfo()
}

func (p People) GetPeopleInfo(){
	fmt.Println(p)
}
func (s Student) GetStudentInfo(){
	fmt.Println(s)
}

func main(){
	var is StudentInfo=Student{People{"李明",18},"Yale University"}
	is.GetStudentInfo()
	is.GetPeopleInfo()
	var ip PeopleInfo=is
	ip.GetPeopleInfo()
}
/**************************************************
拥有超集的接口还可以被转换成子集的接口，这样子集就可以很容易地
访问超集对象中的成员变量或者数据。
**************************************************/