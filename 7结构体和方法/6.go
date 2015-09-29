//嵌入式结构直接用于Map
package main

import(
	"fmt"
)

func main(){
	map1:=map[string]struct{
		name string
		age int
	}{
		"teacher":{"郑智",39}
		"student":{"李明",18}
	}
	fmt.Println(map1)
}