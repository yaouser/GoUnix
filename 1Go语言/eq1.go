//算数运算
package main

import(
	"fmt"
)

func main(){
	var a,b int=7,163
	var c,d float32=2.5,6.3
	var ch1,ch2 rune='中','国'
	fmt.Println(string(ch1)+string(ch2))
	fmt.Println(d-c)
	fmt.Println(c*d)
	fmt.Println(b/a)
	fmt.Println(b%a)
}