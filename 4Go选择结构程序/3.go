//使用if else语句找出最大数
package main

import(
	"fmt"
)

func main(){
	var a,b int
	//通过scanln()函数，从键盘吧输入a,b的值
	fmt.Scanln(&a,&b)
	if a>b {
		//条件为TRUE执行这里
		fmt.Printf("MAX=%d\n",a)
	}else{
		//条件为false执行这里
		fmt.Printf("MAX=%d\n",b)
	}
}