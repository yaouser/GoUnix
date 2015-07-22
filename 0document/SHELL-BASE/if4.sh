#!/bin/bash
#所有的条件必须被满足才会被执行
if [ -f this_file ]; then
	if [ -f that_file ]; then
		if [ -f the_other_file ]; then
			echo "All files present, and correct"
		fi
	fi
fi
