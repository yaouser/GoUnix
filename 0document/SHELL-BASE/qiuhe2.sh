#!/bin/bash
sum=0
for INT in $*
do 
	sum=`expr $sum + $INT`
done

echo $sum
