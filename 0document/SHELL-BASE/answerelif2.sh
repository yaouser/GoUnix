#!/bin/bash
#为了防止if语句中出现timeofday为空字符时，应将变量加双引号。
echo -n "Is it morning? Please enter yes or no: " #使用-n选项时去掉每一行后面的换行符
read timeofday

if [ "$timeofday" = "yes" ]
then 
	echo "Good morning"
elif [ "$timeofday" = "no" ]
then
	echo "Good afternoon"
else
	echo "Sorry, a not recognized. Enter yes or no"
	exit 1
fi

exit 0
