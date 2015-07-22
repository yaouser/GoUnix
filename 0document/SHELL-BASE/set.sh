#!/bin/bash
#set命令的作用是为设置参数变量。许多命令的输出结果是以空格分隔的值，
#如果需要使用输出结果中的某个域，这个命令就非常有用。

echo "the date is $(date)"
set $(date)
echo The month is $2

exit 0

#这个程序把date命令的输出设置为参数列表，然后通过位置参数$2获得月份。

