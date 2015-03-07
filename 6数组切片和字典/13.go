//字典项的删除与增添
package main

import(
	"fmt"
)

func main(){
	var map1=map[string]int{"key1":100,"key2":200,"key3":300}
	//使用range遍历map1的过程中删除和增加字典项
	for k,v:=range map1{
		fmt.Println(k,v)
		//删除字典项
		if k=="key2" {
			delete(map1,k)
		}
		if k=="key3" {
			map1["key4"]=400
		}
	}
	fmt.Println(map1)
}