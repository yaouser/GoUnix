/******************************************
mymath包实现加减乘除四则运算。
加法运算：Add()函数。
减法运算：Sub()函数。
乘法运算：Mult()函数。
除法运算：Div()函数。
*******************************************/
package mymath

func Add(a,b int) int {
	return a+b
}

func Sub(a,b int) int {
	return a-b
}

func Mult(a,b int) int {
	return a*b
}

func Div(a,b int) float32{
	if b!=0 {
		return float32(a)/float32(b)
	}else{
		return 0
	}
}