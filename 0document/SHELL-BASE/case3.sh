#!/bin/bash
#请注意，每个模式行都以双分号(;;)结尾。因为你可以在前后模式之间放置多条语句，
#所以需要使用一个双分号来标记前一个语句的结束和后一个模式的开始。

echo "Is it morning? Please answer yes or no"
read timeofday

case "$timeofday" in 
	yes) echo "Good Morning";;
	no ) echo "Good Afternoon";;
	y  ) echo "Good Morning";;
	n  ) echo "Good Morning";;
	*  ) echo "Sorry, answer not recognized";;
esac 

exit 0
