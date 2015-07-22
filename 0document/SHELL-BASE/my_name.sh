#!/bin/bash
#该实例演示了函数的参数是如何传递的，以及函数如何返回一个true或false值。
#你使用一个参数来调用该脚本程序，该参数是你想要在问题中使用的名字。

yes_or_no() {
	echo "Is your name $* ?"
	while true
	do
		echo -n "Enter yes or no: "
		read x
		case "$x" in
			y | yes ) return 0;;
			n | no  ) return 1;;
			* ) echo "Answer yes or no"
		esac
	done
}

echo "Original parameters are $*"

if yes_or_no "$1"
then 
	echo "Hi $1, nice name"
else 
	echo "Never mind"
fi

exit 0

