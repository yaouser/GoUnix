//浮点型数据的舍入误差
package main

import(
	"fmt"
)

func main(){
	var f1,f2 float32
	f1=12356.789e5
	f2=f1+20
	fmt.Println(f2)
}