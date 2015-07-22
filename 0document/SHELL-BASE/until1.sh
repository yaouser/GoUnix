#!/bin/sh
int=1

until [ $int -gt 5 ]
do 
	sq=`expr $int \* $int`
	echo $sq
	int=`expr $int + 1`
done
echo "completed"
