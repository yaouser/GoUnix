//使用if嵌套语句判断闰年
package main

import(
	"fmt"
)

func main(){
	var year int
	//通过scanln（）函数，从键盘输入year的值
	fmt.Scanln(&year)
	if year%4==0 {
		if year%100==0{
			if year%400==0{
				//能被4整除，且能被100整除，且能被400整除是闰年
				fmt.Printf("%d is leap year.\n",year)
			}else{
				//能被4整除，且能被100整除，但不能被400整除不是闰年
				fmt.Printf("%d is not leap year.\n",year)
			}
		}else {
			//能被4整除，且不能够被100整除是闰年
			fmt.Printf("%d is leap year.\n",year)
		}
	}else{
		//不能被4整除不是闰年
		fmt.Printf("%d is not leap year.\n",year)
	}
}