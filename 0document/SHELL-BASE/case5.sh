#!/bin/bash
#最后，为了让这个脚本程序具备可重用性，你需要在使用默认模式时给出
#另外一个退出码。

echo "Is it morning? Please answer yes or no"
read timeofday

case "$timeofday" in
	yes | y | Yes | YES |[Yy][Ee][Ss])
		echo "Good Morning"
		echo "Up bright and early this morning"
		;;
	[nN]* )
		echo "Good Afternoon"
		;;
	* )
		echo "Sorry, answer not recognized"
		echo "Please answer yes or no"
		exit 1
		;;
esac

exit 0
