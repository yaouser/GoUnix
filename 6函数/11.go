//defer语句保证文件能够正常关闭
package main

import(
	"fmt"
	"os"
	"io"
)

func main(){
	copylen,err:=copyFile("dis.txt","src.txt")
	if err!=nil{
		return
	}else{
		fmt.Println(copylen)
	}
}
//函数copyFile的功能是将源文件src的数据复制给目标文件dst
func copyFile(dstName,srcName string)(copylen int64,err error){
	src,err:=os.Open(srcName)
	if err!=nil {
		return
	}
	//当return时就会调用src.Close()把源文件关闭
	defer src.Close()
	dst,err:=os.Create(dstName)
	if err!=nil{
		return
	}
	//当return时就会调用src.Close()把目标文件关闭
	defer dst.Close()
	return io.Copy(dst,src)
}
/*******************************************
在Go程序中，当程序返回或发生异常时，defer语句通常用来
做一些函数调用后的清理工作，释放资源变量。
关闭文件句柄：
srcFile,err:=os.Open("myfile")
defer srcFile.Close()
关闭互斥锁：
mutex.Lock()
defer mutex.Unlock()
上面例子中defer语句的用法有两个优点：
1让设计者永远也不会忘记关闭文件，有时当函数返回时常常
忘记释放打开的资源变量。
2将关闭和打开靠在一起，程序的意图也变得清晰很多。
*******************************************/