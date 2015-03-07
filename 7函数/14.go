//使用变参函数求最小数
package main

import(
	"fmt"
)

func main(){
	min:=Min(12,3,24,7)
	fmt.Println("最小数为：",min)
	arr:=[]int{-6,13,7,25,-13}
	min=Min(arr...)
	fmt.Println("数组中的最小数为：",min)
}
/*
	函数Min()的功能是找出最小数
*/
func Min(a ...int) int {
	if len(a)==0{
		return 0
	}
	min:=a[0]
	for _,v:=range a{
		if v<min{
			min=v
		}
	}
	return min
}