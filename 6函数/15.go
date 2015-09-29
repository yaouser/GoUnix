//使用多返回值函数找出字节切片中的整型数
package main

import(
	"fmt"
)

func main(){
	a:=[]byte{'1','2','3','4','z','x','5','6','d','g','7','8'}
	var index,x,start int
	fmt.Println("请输入起始索引位置...")
	fmt.Scanf("%d",&start)
	for {
		index,x=nextInt(a,start)
		fmt.Printf("找到整数%d，下一个数的位置从%d开始。\n",x,index)
		if index !=0 && index<len(a){
			start=index
			continue
		}else{
			break
		}
	}
}
/*
	函数nextInt()的功能是从指定位置找出字节切片中的整数，
	并返回下一个数的索引位置。
*/
func nextInt(b []byte,i int)(index,x int){
	var start bool
	if b[i]>='0' && b[i]<='9'{
		start=true
	}
	for ;i<len(b);i++{
		if b[i]>='0' && b[i]<='9'{
			if start==false {
				index=1
				break
			}else{
				x=x*10+int(b[i])-'0'
			}
		}else if b[i]<'0'||b[i]>'9'{
			start=false
		}
	}
	return index,x
}