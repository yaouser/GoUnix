//字段标签
package main

import(
	"fmt"
	"reflect"
)

type user struct{
	Id		int		"账号"
	Name	string	"姓名"
	Sex		bool	"性别"
}

func main(){
	u:=user{100,"张三",true}
	//使用TypeOf()函数获取对象的类型
	t:=reflect.TypeOf(u)
	//使用ValueOf()获取对象的值
	v:=reflect.ValueOf(u)
	for i:=0;i<t.NumField();i++ {
		f:=t.Field(i)
		fmt.Printf("%s(%s=%v)\n",f.Tag,f.Name,v.Field(i).Interface())
	}
}
/********************************************
Go语言中允许为字段定义标签（Tag）。在Go语言中，不允许
在普通程序语句中访问字段标签，字段标签只能使用反射包
（reflect）中提供的特殊方法进行读取。利用reflect可以
获取字段名，字段类型和字段值。
********************************************/