#!/bin/bash
#trap命令用于指定在接受到信号后将要采取的行动，trap命令的一种常见用途是在脚本
#程序被中断时完成清理工作。
#trap命令有两个参数，第一个参数是接受到指定信号时将要采取的行动，第二个参数是
#要处理的信号名。
#trap command signal

trap 'rm -f /tmp/my_tmp_file_$$' INT
echo "creating file /tmp/my_tmp_file_$$"
date > /tmp/my_tmp_file_$$

echo "press interrupt (CTRL-C) to interrupt ...."
while [ -f /tmp/my_tmp_file_$$ ]
do 
	echo File exists
	sleep 1
done
echo "The file no loger exists"

trap INT 
echo "creating file /tmp/my_tmp_file_$$"
date > /tmp/my_tmp_file_$$

echo "press interrupt (control-C) to interrupt ...."
while [ -f /tmp/my_tmp_file_$$ ]
do 
	echo File exists
	sleep 1
done

echo we never get here
exit 0
