//数组的定义
package main

import(
	"fmt"
)

func main(){
	var i int
	var arr1 [5] int
	for i=0; i<=4;i++ {
		arr1[i]=i+1
	}
	fmt.Println(arr1)
}