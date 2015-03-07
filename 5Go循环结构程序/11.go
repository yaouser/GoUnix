//猜数字游戏
package main

import(
	"fmt"
	"bufio"
	"os"
	"strconv"
)

const GOAL int=36          //设置谜底

func main(){
	var data int           //存储输入数字
	var count int          //统计用了多少次才猜中
	var c byte             //存放每个按键ASCII码
	var str string         //存放输入字符串流
	fmt.Printf("请输入一个正整数：")
LABEL1:
	r:=bufio.NewReader(os.Stdin)  //初始化标准输入设备
	w:=bufio.NewWriter(os.Stdout) //初始化标准输出设备
	/*********************************************
	键盘输入处理流程
	*********************************************/
	for i:=0; ;i++ {
		c,_=r.ReadByte()
		if c==10||c==13{     //如果是回车或换行则退出
			break
		}else{
			w.Flush()
			str+=string(c)
		}
	}
	for {
		data,_=strconv.Atoi(str)
		if data > GOAL {
			fmt.Printf("数字%d有点大，请再试一试...\n",data)
			count++
			str=""
			goto LABEL1
		}else if data < GOAL {
			fmt.Printf("数字%d有点小，请在试一试...\n",data)
			count++
			str=""
			goto LABEL1
		}else {
			fmt.Printf("恭喜你，猜中了！")
			break
		}
	}
	fmt.Printf("你一共猜了%d次才猜对了。",count)
}