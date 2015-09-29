//Struct的导入
package main

import(
	"fmt"
	"mystruct"
)

func main(){
	usr:=new(mystruct.User)
	usr.Id=100
	usr.Name="张三"
	usr.SayHi()
	fmt.Println(usr)
}
/***************************************
Go语言中要使包内的某个标示符对其他包可见，该标示
符的首写字母必须大写，这就是Go语言的可见性规则。
***************************************/