#!/bin/bash
#until是只有在条件为真时才执行循环
#while是执行循环直到条件为真时退出

#这是until例子，你设置一个警报，当某个特定的用户登录时，该警报就会启动，你
#通过命令行将用户名传递给脚本程序。

until who | grep "$1" > /dev/null
do
	sleep 60
done 

#now ring the bell and announce the expected user.

echo -e '\a'
echo "***** $1 has just logged in *****"

exit 0

#如果用户已经登录，那么循环就不需要执行。
