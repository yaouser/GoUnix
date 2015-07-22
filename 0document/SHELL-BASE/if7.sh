#!/bin/bash

rm -f file_one
if [ -f file_one ] || echo "hello" || echo " there"
then
	echo "in if"
else
	echo "in else"
fi

exit 0

#[ -f file_one ] && command for true || command for false
#在上面的语句中，如果测试成功就会执行第一条命令，否则执行第二条命令。

#如果你想在某些只允许使用单个语句的地方（比如在AND或OR列表中）使用多条语句，你
#可以把它们括在花括号{}中来构造一个语句块。
#get_confirm && {
#	grep -v "$cdcatnum" $tracks_file > $temp_file
#	cat $temp_file > $tracks_file
#	echo
#	add_record_tracks
#	}
