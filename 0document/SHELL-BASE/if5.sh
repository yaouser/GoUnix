#!/bin/bash
#希望至少在这一系列条件中有一个为真。

if [ -f this_file ];then
	foo="True"
elif [ -f that_file ];then
	foo="True"
elif [ -f the_other_file ];then
	foo="True"
else 
	foo="False"
fi

if [ "$foo" = "True" ];then
	echo "One of the files exists"
fi

