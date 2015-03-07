//多个Method可以同名
package main

import(
	"fmt"
//	"math"
)
/*
type rectangle struct{
	width int
	height int
}
type circle struct{
	radius float32
}
//此处area()是对象rectangle的一个方法
func (recv rectangle) area() int {
	return recv.width*recv.height
}
//此处area()是对象circle的一个方法
func (recv circle) area() float32{
	return recv.radius*recv.radius*math.Pi
}
func main(){
	r1:=rectangle{4,3}
	c1:=circle{5}
	fmt.Println(r1.area(),c1.area())
}
*/
/******************************************
如果普通类型作为receiver，他只是一个值传递，而指针类
型作为receive，他将是一个引用传递。两者的差别在于，
指针会对实例对象的内容发生操作，而普通类型仅是以副本作
为操作对象。
Method的名字一样，但如果接受者不一样，那method就不
一样。
Method里面可以访问接受者的字段，调用Method进行访问，
就像在struct里访问字段一样。
******************************************/

/*
type coordinate struct {
	x int
	y int
}
func (recv coordinate) swap(){
	var temp int
	temp,recv.x,recv.y=temp,recv.y,recv.x
	fmt.Println(recv)
}
func main(){
	r1:=coordinate{3,4}
	r1.swap()
	fmt.Println(r1)
}
*/

type coordinate struct {
	x int
	y int
}
func (recv* coordinate) swap(){
	var temp int
	temp,recv.x,recv.y=temp,recv.y,recv.x
	fmt.Println(recv)
}
func main(){
	r1:=coordinate{3,4}
	p:=&r1
	p.swap()
	fmt.Println(r1)
}