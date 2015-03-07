//使用多条件判断打印成绩等级
package main

import(
	"fmt"
)

func main(){
	var grade int
	//通过scanln（）函数，从键盘输入成绩grade
	fmt.Scanln(&grade)
	if (grade>=90)&&(grade<=100){
		//成绩在90-100区间等级为A
		fmt.Println("A")
	}else if(grade>=80)&&(grade<90){
		//成绩在80-90区间等级为B
		fmt.Println("B")
	}else if(grade>=70)&&(grade<80){
		//成绩在70-79区间等级为C
		fmt.Println("C")
	}else{
		//成绩在70以下等级为D
		fmt.Println("D")
	}
}