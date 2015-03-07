//复数的定义和使用
package main

import(
	"fmt"
)

func main(){
	var cp1,cp2 complex64
	cp1=1.2+3.4i
	cp2=cp1
	cp3:=complex(1.2,3.4)
	fmt.Println(cp1)
	fmt.Println(cp2)
	fmt.Println(cp3)
}