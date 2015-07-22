#!/bin/bash

foo() {
	echo "Function foo is executing"
}

echo "script starting"
foo
echo "script ended"

exit 0

#你必须在调用函数之前定义一个你需要的函数
#也就是说，定义函数在前，调用函数在后。
#这样才能正确执行。

#当一个函数被调用时，脚本程序的位置参数($* $@ $1 $2等)会被替换为函数的
#参数。这也是你读取传递给函数的参数的办法。当函数执行完毕后，这些参数会
#恢复为它们先前的值。但移植性不强可用另外的方法：
#foo(){echo JAY;}
#...
#result="$(foo)"
