//数组重组
package main

import(
	"fmt"
	"strconv"
)

func main(){
	//整形变量num用来存储生成的三位数
	var num,col int
	//字符串numStr用来存数位
	var numStr string 
	//三重循环嵌套实现该算法
	for i:=1;i<5;i++ {
		for j:=1;j<5;j++ {
			for k:=1;k<5;k++{
				if i!=k && i!=j && j!=k {//确保个位，十位百位互不相同
					numStr=strconv.Itoa(i)+strconv.Itoa(j)+strconv.Itoa(k)
					num,_=strconv.Atoi(numStr)
					fmt.Printf(" %d",num)
					col++
					//控制每行输出5个数
					if col==5 {
						fmt.Printf("\n")
						col=0
					}
				}
			}
		}
	}
}