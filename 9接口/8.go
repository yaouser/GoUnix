//利用反射动态调用原对象方法
package main

import(
	"fmt"
	"reflect"
)

type Student struct{
	Id int
	Name string
	Sex bool
	Grade float32
}
func (s Student) SayHi(name string){
	fmt.Printf("Hi %s,my name is %s.\n",name,s.Name)
}

func main(){
	stu:=Student{10001,"李明",false,90.5}
	v:=reflect.ValueOf(stu)
	mv:=v.MethodByName("SayHi")
	args:=[]reflect.Value{reflect.ValueOf("张衡")}
	mv.Call(args)
}
/**************************************************
在reflect包中提供了Call()方法可以用于调用原对象的方法。Call()
方法可以向原对象传递参数列表并调用他。如果调用的是一个非方法类型，
则Call()会引发panic错误事件。
一个原对象通常会有多个方法，可以使用reflect包中的MethodByName()
方法来获取想要的方法的值，如果访问的方法不存在，则返回值”0“。
需要注意的是，在Go语言中传递给方法的参数要和方法定义的参数类型
保持一致。
**************************************************/