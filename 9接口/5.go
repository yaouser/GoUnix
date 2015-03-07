//利用Switch测试进行接口类型推断
package main

import(
	"fmt"
)

type People struct{
	Name string
	Age int
}
//定义空接口用于存储任意数据类型
type Tester interface{}

func main(){
	people:=People{"张三",20}
	it:=make([]Tester,4)
	it[0],it[1],it[2],it[3]=1,"Hello",people,true
	for index,element:=range it{
		switch value:=element.(type){
			case int:
				fmt.Printf("it[%d] is an int.value=%d\n",index,value)
			case string:
				fmt.Printf("it[%d] in a string.value=%d\n",index,value)
			case People:
				fmt.Printf("it[%d] is a People.value=%v\n",index,value)
			default:
				fmt.Printf("it[%d] is an unknown type.\n",index)
		}
	}
}
/**************************************************
使用Switch测试推断接口类型的格式如下：
switch value:=element.(type){
	case int:
	case strint:
	........
}
需要强调的是element.(type)语法不能在Switch语句外的任何逻辑
语句里面使用，如果要在Switch外面判断一个类型就要使用Comma-ok
断言。
*************************************************/