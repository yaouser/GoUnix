//多选项case语句
package main

import(
	"fmt"
)

func main(){
	//搜索字符串变量
	var searchstr string
	//通过scanf（）函数，格式化从键盘输入操作数和云算数
	fmt.Scanf("%s",&searchstr)
	switch searchstr{
	case "c","C":
		//搜索字符串是c或C都执行下面语句
		fmt.Println("C programing language")
	case "go","golang":
		//搜索字符串是go或golang都执行下面语句
		fmt.Println("Go programing language")
	case "java":
		//搜索字符串是Java执行下面语句
		fmt.Println("Java programing language")
	default:
		//没有搜索到
		fmt.Println("Not find the result!")
	}
}