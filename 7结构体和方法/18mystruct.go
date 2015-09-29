/***********************************
mystruct包定义了User对象，并实现了两个对
User进行操作的方法。
SayHi()方法
old()方法
***********************************/
package mystruct

import (
	"fmt"
)

type User struct{
	Id int
	Name string
	age int
}

func (u* User) SayHi(){
	fmt.Printf("Hi,I'm %s.Nice to meet you!",u.Name)
}

func (u User) old() int {
	return u.age
}