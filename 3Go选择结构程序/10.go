//使用无条件表达式swith语句打印成绩等级
package main

import(
	"fmt"
)

func main(){
	var grade int
	//通过scanln()函数，从键盘输入成绩grade
	fmt.Scanln(&grade)
	switch{
		case(grade>=90)&&(grade<=100):
			//成绩在90-100区间等级为A
			fmt.Println("A")
		case(grade>=80)&&(grade<90):
			//成绩在80-89区间等级为B
			fmt.Println("B")
		case(grade>=70)&&(grade<80):
			//成绩在70-79区间等级为C
			fmt.Println("C")
		default:
			//成绩在70以下等级为D
			fmt.Println("D")
	}
}