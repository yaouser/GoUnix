package main

import(
	"fmt"
)

const(
	a=1.2+3.4i
	b=real(a)
	c=imag(a)
)

func main(){
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
}