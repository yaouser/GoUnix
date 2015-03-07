//字典项查找
package main

import(
	"fmt"
)

func main(){
	var map1=map[string]int{"key1":100,"key2":200}
	//key值存在
	v,OK:=map1["key1"]
	if OK {
		fmt.Println(v,OK)
	}else{
		fmt.Println(v)
	}
	//key值不存在
	v,k:=map1["key3"]
	if k {
		fmt.Println(v,k)
	}else{
		fmt.Println(v)   
	}   
}
