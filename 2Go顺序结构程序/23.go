//求解一元二次方程
package main

import(
	"fmt"
	"math"
)

func main(){
	var a,b,c,disc,x1,x2,p,q float64
	fmt.Println("请输入一元二次方程三个系数：")
	fmt.Scanf("%f,%f,%f",&a,&b,&c)
	disc=b*b-4*a*c
	p=-b/(2*a)
	q=math.Sqrt(disc)/(2*a)
	x1=p+q
	x2=p-q
	fmt.Printf("一元二次方程的根x1=%6.2f x2=%6.2f\n",x1,x2)
}