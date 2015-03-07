//数组的初始化
package main

import(
	"fmt"
)

func main(){
	var a [10]int
	//定义时对数组元素赋初值
	var b=[10]int{1,2,3,4,5,6,7,8,9,10}
	//可初始化列表决定数组长度
	var c=[]int{1,2,3}
	//只给一部分元素赋初值
	var d=[10]int{1,2,3,4,5}
	//按下标初始化元素
	var e=[10]int{2:4,5:7}
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	fmt.Println(d)
	fmt.Println(e)
}