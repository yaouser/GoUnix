//Go语言标准库bytes包，提供了对字节切片进行读写操作
//的一系列函数和方法。
//Contains,Count,Map,Repeat,Replace,Runes,Join
package main

import(
	"fmt"
	"bytes"
)
/*
func main(){
	s:=[]byte("Golang")
	subslice1:=[]byte("go")
	subslice2:=[]byte("Go")
	fmt.Println(bytes.Contains(s,subslice1))
	fmt.Println(bytes.Contains(s,subslice2))
}
*/

/*
func main(){
	s:=[]byte("banana")
	sep1:=[]byte("ban")
	sep2:=[]byte("na")
	sep3:=[]byte("a")
	fmt.Println(bytes.Count(s,sep1))
	fmt.Println(bytes.Count(s,sep2))
	fmt.Println(bytes.Count(s,sep3))
}
*/

/*
func main(){
	s:=[]byte("同学们，上午好！")
	m:=func(r rune) rune{
		if r=='上'{
			r='下'
		}
		return r
	}
	fmt.Println(string(s))
	fmt.Println(string(bytes.Map(m,s)))
}
*/

/*
func main(){
	b:=[]byte("na")
	count:=2
	fmt.Println("ba"+string(bytes.Repeat(b,count)))
}
*/

/*
func main(){
	s:=[]byte("google")
	old:=[]byte("o")
	new:=[]byte("oo")
	n:=1
	fmt.Println(string(bytes.Replace(s,old,new,n)))
	fmt.Println(string(bytes.Replace(s,old,new,-1)))
}
*/
/*************************************************
rune函数的功能是把s转换成UTF-8编码的字节序列，并返回对应
的Unicode切片。
*************************************************/

/*
func main(){
	s:=[]byte("中华人民共和国")
	r:=bytes.Runes(s)
	fmt.Printf("转换前字符串%q长度：%d字节\n",string(s),len(s))
	fmt.Printf("转换后字符串%q长度：%d字节\n",string(s),len(r))
}
*/

/*
func main(){
	s:=[][]byte{
		[]byte("你好"),
		[]byte("世界！"),
	}
	sep:=[]byte(",")
	fmt.Println(string(bytes.Join(s,sep)))
}
*/
/**********************************************
字节切片比较有三个compare,equal,equalfold
*********************************************/
/*
func main(){
	a:=[]byte("abc")
	b:=[]byte("Abc")
	s:=[]byte("GOLANG")
	t:=[]byte("golang")
	fmt.Println(bytes.Compare(a,b))
	fmt.Println(bytes.Equal(a,b))
	fmt.Println(bytes.EqualFold(s,t))
}
*/
/*******************************************
HasPrefix是用来检查字节切片的前缀
HasSuffix是用来检查自己切片的后缀
*******************************************/
/*
func main(){
	s:=[]byte("recover")
	prefix:=[]byte("re")
	suffix:=[]byte("ed")
	fmt.Println(bytes.HasPrefix(s,prefix))
	fmt.Println(bytes.HasSuffix(s,suffix))
}
*/
/******************************************
字节切片位置索引函数共有8个Index IndexAny 
IndexByte IndexFunc IndexRune LastIndex
LastIndexAny LastIndexFunc
******************************************/

/*
func main(){
	s:=[]byte("Google")
	sep:=[]byte("o")
	chars:="gle"
	var c byte='g'
	var r rune='l'
	fmt.Printf("切片%q在%q中的位置索引是：%d\n",sep,s,bytes.Index(s,sep))
	fmt.Printf("切片%q中的任一字符在%q中的位置索引是：%d\n",chars,s,bytes.IndexAny(s,chars))
	fmt.Printf("byte型字符%q在%q中的位置索引是：%d\n",c,s,bytes.IndexByte(s,c))
	fmt.Printf("切片%q第一个符合f()的字符索引是：%d\n",s,bytes.IndexFunc(s,f))
	fmt.Printf("rune型字符%q在%q中的位置索引是：%d\n",r,s,bytes.IndexRune(s,r))
	fmt.Printf("切片%q在%q中的最后索引是：%d\n",sep,s,bytes.LastIndex(s,sep))
	fmt.Printf("切片%q中的任一字符在%q中的最后位置索引是：%d\n",chars,s,bytes.LastIndexAny(s,chars))
	fmt.Printf("切片%q最后一个符合f()的字符索引是：%d\n",s,bytes.LastIndexFunc(s,f))
}
func f(a rune) bool {
	if a>'k' {
		return true
	}else{
		return false
	}
}
*/
/********************************************
字节切片分割函数Fields FieldsFunc Split SplitN
SplitAfter SplitAfterN
********************************************/
