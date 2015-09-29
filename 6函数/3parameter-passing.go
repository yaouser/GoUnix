//参数传递
package main

import(
	"fmt"
)
//常规传参
/*************************************
func main(){
	var b int=1
	f1(b)
	fmt.Println(b)
}
func f1(a int){
	a=2
	fmt.Println(a)
}
************************************/
//指针传参
/***********************************
func main(){
	var b int=1
	f2(&b)
	fmt.Println(b)
}
func f2(a* int){
	*a=2
	fmt.Printf("%d\n",*a)
}
***********************************/
//数组元素作为函数参数（复制）
/***********************************
func main(){
	var b=[5]int{1,2,3,4,5}
	f3(b[2])
	fmt.Println(b[2])
}
func f3(a int){
	a+=1
	fmt.Println(a)
}
**********************************/
//数组名作为函数参数（复制）
/*********************************
func main(){
	var b=[5]int{1,2,3,4,5}
	f4(b)
	fmt.Println(b[0])
}
func f4(a[5]int){
	a[0]+=1
	fmt.Println(a[0])
}
**********************************/
//切片作为函数参数（操作底层数据）
/**********************************
func main(){
	var b=[5]int{1,2,3,4,5}
	var s1 []int=b[:]
	f5(s1)
	fmt.Println(b[0])
}
func f5(s[]int) {
	s[0]+=1
	fmt.Println(s[0])
}
**********************************/
//函数作为参数
/**********************************
func main(){
	var a,b int=3,4
	f:=sum
	f6(a,b,f)
}
func f6(a,b int,sum func(int,int) int){
	fmt.Println(sum(a,b))
}
func sum(a,b int) int {
	return a+b
}
**********************************/
//return语句
/**********************************
func main(){
	sum:=f1(3,6)
	fmt.Println(sum)
}
func f1(a,b int) int {
	return a+b
}
**********************************/
//多返回值
/*********************************
func main(){
	sum,div:=f2(3,6)
	fmt.Println(sum,div)
}
func f2(a,b int)(int,float32){
	return a*b,float32(a)/float32(b)
}
*********************************/
//命名返回值参数
//Go语言中还支持命名返回值参数，如果使用
//命名返回值参数，则return语句可以为空。
//否则，return语句必须按照顺序返回多个
//结果。
/*********************************
func main(){
	sum,sub:=f3(3,6)
	fmt.Println(sum,sub)
}
func f3(a,b int)(sum,sub int){
	sum=a+b
	sub=a-b
	return
}
**********************************/
