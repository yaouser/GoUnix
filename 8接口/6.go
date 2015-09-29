//利用接口反射获取结构体Type,Value和Method
package main

//导入reflect包
import(
	"fmt"
	"reflect"
)

//定义结构体
type Student struct{
	Id 		int
	Name 	string
	Sex		bool
	Grade	float32
}

//为结构体定义两个方法
func (s Student) SayHi(){
	fmt.Printf("Hi,nice to meet you!\n")
}
func (s Student) MyName(){
	fmt.Printf("My name is %s,I come from China.",s.Name)
}
//反射处理函数
func StructInfo(o interface{}){
	t:=reflect.TypeOf(o)
	//判断是否为结构体类型
	if k:=t.Kind();k!=reflect.Struct{
		fmt.Printf("This value is not a struct,it's %v.",k)
		return
	}
	fmt.Println("Struct name is:",t.Name())
	fmt.Println("Fields of th struct is:")
	v:=reflect.ValueOf(o)
	//获取字段Type和Value信息
	for i:=0;i<t.NumField();i++{
		field:=t.Field(i)
		value:=v.Field(i).Interface()
		fmt.Printf(" %6s:%v=%v\n",field.Name,field.Type,value)
	}
	fmt.Println("Methods of the struct is:")
	//获取方法Name和Type信息
	for i:=0;i<t.NumMethod();i++{
		method:=t.Method(i)
		fmt.Printf(" %6s:%v\n",method.Name,method.Type)
	}
	
}

func main(){
		stu:=Student{10001,"李明",false,90.5}
		StructInfo(stu)
}
/***********************************************************
当在定义结构体字段名和方法名时要注意，这些名字的首写字母一定要大写，否测
反射操作时编译器会报错。
反射会将匿名字段当作一个独立字段来处理，如果要获取嵌入字段的Type和Value
信息，必须使用索引路径来完成。通过Value类型的FieldByIndex()方法可以
获取嵌入字段的索引路径，原型定义如下：
func (v Value) FieldbyIndex(index []int) Value
实例如下：
type Student struct{
	Id		int
	Name	string
	Sex		bool
	Grade	float32
}
type Monitor struct{
	Student
	As string
}
func main(){
	stu:=Monitor{Student{10001,"李明",false,90.5},"班长"}
	t:=reflect.TypeOf(stu)
	v:=reflect.ValueOf(stu)
	for i:=0;i<t.NumField();i++{
		if t.Field(i).Anonymous{
			for j:=0;j<v.Field(i).NumField();j++{
				fmt.Println("Embedded fields:",v.FieldByIndex([]int{i,j}).Interface())
			}
		}else{
			fmt.Println("Normal fields:",v.Field(i),Interface())
		}
	}
}
**************************************************************/

/**********************************************
反射（reflect）是Go语言获取程序运行时类型信息的方式。Go
语言标准库的reflect包中提供了TypeOf()和ValueOf()函数，
可以使用这两个函数从interface()接口对象中获取实际目标对
象的类型和值信息。在reflect的TypeOf()中有一个非常有用的
方法是kind()，这个方法可以返回该类型的具体信息，如int,
int8,float64等。ValueOf()中则包含一系列同类型相关的方
法，如int(),float32(),bool()等，用于返回对应类型的值。
package main
import(
	"fmt"
	"reflect"
)
func main(){
	var pi float64=3.14
	t:=reflect.TypeOf(pi)
	v:=reflect.ValueOf(pi)
	fmt.Println("Type:",t.Kind())
	fmt.Println("Value",v.Float())
}
**********************************************
Go语言除了能对最基本的数据类型进行反射操作，还能够对结构
体进行反射操作。
*********************************************/