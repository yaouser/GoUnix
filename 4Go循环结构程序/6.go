//判断素数
package main

import(
	"fmt"
	"math"
)

func main(){
	var num,i,k int
	//从键盘格式化输入要判断的数num
	fmt.Scanf("%d",&num)
	//对num开方并取整求k
	k=int(math.Sqrt(float64(num)))
	for i=2;i<=k;i++ {
		if num%i==0{
			break       //num能被2-k中的一个数整除，跳出循环
		}
	}
	if i>k {
		fmt.Printf("The number %d is a prime number.\n",num)
	}else{
		fmt.Printf("The number %d is not a prime number.\n",num)
	}
}