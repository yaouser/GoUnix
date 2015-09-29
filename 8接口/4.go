//利用Comma-OK断言进行接口类型推断
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
		if value,ok:=element.(int);ok{
			fmt.Printf("it[%d] is an int. value=%d\n",index,value)
		}else if value,ok:=element.(string);ok {
			fmt.Printf("it[%d] ins a string. value=%s\n",index,value)
		}else if value,ok:=element.(People);ok{
			fmt.Printf("it[%d] is a People. value=%v\n",index,value)
		}else{
			fmt.Printf("it[%d] is an unknown type.\n",index)
		}
	}
}
/************************************************
如何反向知道接口类型变量里面实际保存的是哪一种类型的对象呢？
这就是接口类型的推断。在Go语言中，目前有两种常用的方法可以进
行接口类型推断：Comma-ok断言和Switch测试。
Comma-ok断言进行接口类型查询的格式如下：
value,ok=element.(T)
value是变量的值，ok是bool类型，element是接口类型变量，T
是断言的类型。
************************************************/