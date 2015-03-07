//接口的定义与赋值
package main

import(
	"fmt"
)

//定义对象People,Teacher和Student
type People struct{
	Name string
}
type Teacher struct{
	People
	Department string
}
type Student struct{
	People
	School string
}

//对象方法的实现
func (p People) SayHi() {
	fmt.Printf("Hi,I'm %s.Nice to meet you!\n",p.Name)
}
func (t Teacher) SayHi() {
	fmt.Printf("Hi,my name is %s.I'm working in %s.\n",t.Name,t.Department)
}
func (s Student) SayHi() {
	fmt.Printf("Hi,my name is %s.I'm studing in %s.\n",s.Name,s.School)
}
func (s Student) Study() {
	fmt.Printf("I'm learning Golang in %s.\n",s.School)
}
//定义接口Speaker和Learner
type Speaker interface{
	SayHi()
}
type Learner interface{
	SayHi()
	Study()
}

func main(){
	people:=People{"张三"}
	teacher:=Teacher{People{"郑智"},"Computer Science"}
	student:=Student{People{"李明"},"Yale University"}
	var is Speaker			//定义Speaker类型的变量is
	is=people				//is能存储People
	is.SayHi()
	is=teacher				//is能存储Teacher
	is.SayHi()
	is=student			//is能存储Student
	is.SayHi()
	var il Learner			//定义Learner类型的变量il
	il=student				//il能存储Student
	il.Study()
}
/******************************************************
Go语言使用Struct,Method,interface实现了面向对象的功能，简单说，
Interface是一组Method的组合，可以通过Interface来定义对象的一组
行为。Interface是一个方法的集合。
type interfaceName interface {
	methodName(参数列表)(返回值)
	methodName(参数列表)(返回值)
}
在Go语言中，一个interface中包含的Method不宜过多，一般0-3个即可。
一个Interface可以被任意的对象实现；相同，一个对象也可以试想多个
Interface。
在Go语言中，除了类型可以匿名组合，接口也可以组合在一起。将一个接口
匿名嵌入到另外一个接口中，就是接口组合，这种类似于接口的继承。
在Go语言中，任何数据类型都默认实现了空Interface（空接口的定义为：
interface{}），空接口包含0个Method的Interface，所以空接口可以
定义任意类型变参函数。一个函数把interface{}作为参数，那么他可以接
受任意类型的值作为参数，如果一个函数返回interface{}，就可以返回
任意类型的值。
1 T仅拥有属于T类型的方法集，而*T则同时拥有（T+*T）方法集。
2 基于T实现方法，表示同时实现了interface(T)和interface(*T)接口。
3 基于*T实现方法，那就只能是对interface(*T)实现接口。
*******************************************************/