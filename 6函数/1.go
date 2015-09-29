//调用时间函数获取当前日期（调用标准函数）
package main

import(
	"fmt"
	"time"
)

func main(){
	//获取时间戳
	t:=time.Now()
	//按年 月 日 时 分 秒格式输出
	fmt.Println(t.String())
	//按年 月 日格式输出
	fmt.Println(t.Format("2006年01月02日"))
	//输出星期几
	fmt.Println(t.Weekday().String())
}