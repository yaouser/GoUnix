package main

import(
	"fmt"
)

const i=1

const(
	a=i
	b=i+1
	c=i+2
)

func main(){
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
}