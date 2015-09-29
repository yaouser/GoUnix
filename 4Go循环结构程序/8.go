//goto语句实现直到型循环计算1-100的累加和
package main

import(
	"fmt"
)

func main(){
	var i,sum = 1,0
HERE:
	sum=sum+i
	i++
	if i<=100 {
		goto HERE
	}
	fmt.Println(sum)
}
/**************************************************
和break,continue语句不同，goto语句天生就要和标签（label）
配合起来使用。goto语句可以调整程序执行的位置，它可以让程序
无条件跳转到一个标签（label）之处继续执行。
LABEL:
	:
	:
break 	LABEL
***************
LABEL:
	:
	:
continue LABEL
***************
LABEL:
	:
	:
goto  LABEL
***************
总结：break，continue，goto语句都可以和标签配合使用，
但作用有所区别。通常情况下，如果不使用标签break，continue
语句只能跳出本层循环。配合标签使用以后，break，continue
语句就可以跳出多重循环了，这时就不必使用goto语句了。
***************************************************/