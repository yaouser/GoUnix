//位运算
package main

import(
	"fmt"
)

func main(){
	var a,b int=6,11
	fmt.Println(^a)     //取反运算
	fmt.Println(a&b)    //按位与
	fmt.Println(a|b)    //按位或
	fmt.Println(a&^b)   //标志位清除
	fmt.Println(a<<2)   //左移2位
	fmt.Println(b>>2)   //右移2位
}