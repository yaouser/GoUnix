//匿名Recevier
package main

import(
	"fmt"
)

type object struct{
	id int
	name string
}

func (object) msgbox(){
	fmt.Println("This is a object!")
}
func (*object) msgBox(){
	fmt.Println("This is a object!")
}

func main(){
	obj:=object{}
	p:=&obj
	obj.msgbox()
	p.msgBox()
}
/***************************************
如果方法代码中从不使用Receiver参数，那么就可以省
略Receiver的变量名，此时的接受者将是一个匿名
Receiver。
***************************************/