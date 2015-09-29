//利用反射修改原对象Value值
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

func SetValue(o interface{}){
	v:=reflect.ValueOf(o)
	if v.Kind()!=reflect.Ptr||!v.Elem().CanSet(){
		//判断interface接口传递过来的对象是否是指针类型（Ptr）
		//如果是，则调用Elem()方法获取指向实际对象的指针。
		fmt.Println("Cannot set!")
		return
	}else{
		v=v.Elem()
	}
	for i:=0;i<v.NumField();i++{
		switch v.Field(i).Kind(){
			case reflect.Int:
				v.Field(i).SetInt(10002)
			case reflect.String:
				v.Field(i).SetString("张衡")
			case reflect.Bool:
				v.Field(i).SetBool(true)
			case reflect.Float32:
				v.Field(i).SetFloat(95.5)
		}
	}
}

func main(){
	stu:=Student{10001,"李明",false,90.5}
	fmt.Println(stu)
	SetValue(&stu)
	fmt.Println(stu)
}
/*****************************************************
在Go语言中，要利用反射修改原数据对象的前提是该对象是”可修改属性“
(settability)。可以使用CanSet()方法判断一个对象是否可修改属性。
CanSet()返回true，进而就可以调用方法修改对象的值（如SetInt(),
SetBool(),SetFloat()等），如果值不可修改，则返回false。
在Go语言中所有类型是值类型，所以要想在函数中对变量本身进行修改，必须
使用引用类型，即指针。
****************************************************/