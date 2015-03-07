//使用条件循环控制结构近似求π
package main

import(
	"fmt"
	"math"
)

func main(){
	var s int=1
	var n,t,pi float64=1.0,1.0,0
	for math.Abs(t)>1e-6 {
		pi=pi+t
		n=n+2
		s=-s
		t=float64(s)/n
	}
	pi=pi*4
	fmt.Printf("pi=%10.6f\n",pi)
}