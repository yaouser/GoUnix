package main

import(
	"fmt"
)

func main(){
	var x,y int=1,2
	{
		var x int=2
		{
			var x int=3
			fmt.Println(x,y)
		}
		fmt.Println(x,y)
	}
	fmt.Println(x,y)
}