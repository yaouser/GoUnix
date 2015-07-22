#!/bin/bash
#冒号(:)命令是一个空命令。它偶尔会被用于简化条件逻辑，相当与true的一个别名。
#由于它的内置命令，所以它运行的比true块，但它的输出可读性较差。
#你可能会看到将它用作while循环的条件，while :实现了一个无限循环，代替了
#更常见的while true。
# : 结构也会被用在变量的条件设置中，例如：${var:=value} 如果没有 : ,shell
#将试图把$var当作一条命令来处理。

rm -f fred 
if [ -f fred ];then
	:
else
	echo "file fred did not exist"
fi

exit 0
