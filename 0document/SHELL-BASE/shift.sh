#!/bin/bash
#
#Filename: shift
until [ $# -eq 0 ]
do 
	echo "Argument is $1 and `expr $# - 1` argument(s) remain"
	shift
done
