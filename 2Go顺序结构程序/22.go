//计算三角形面积和周长
package main

import(
	"fmt"
	"math"
)

func main(){
	var a,b,c,l,h1,area float64
	fmt.Println("请输入三角形三边：")
	fmt.Scanf("%f,%f,%f",&a,&b,&c)
	l=a+b+c
	fmt.Printf("三角形周长=%6.2f\n",l)
	h1=l*0.5
	area=math.Sqrt(h1*(h1-a)*(h1-b)*(h1-c))
	fmt.Printf("三角形面积=%6.2f\n",area)
}