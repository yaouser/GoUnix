//关系运算
package main

import(
	"fmt"
)

func main(){
	var a,b int=5,10
	var ch1,ch2 byte = 'a','b'
	fmt.Println(a>b)
	fmt.Println(b>=8)
	fmt.Println(ch1<ch2)
	fmt.Println(ch2<='B')
	fmt.Println(ch1==97)
	fmt.Println(ch2!=98)
}