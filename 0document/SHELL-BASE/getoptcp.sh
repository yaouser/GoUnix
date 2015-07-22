#!/bin/bash
#Syntax: getoptcp [-c integer] [-v] filename
#where integer is the number of duplicatecopies
#and -v is the verbose option

COPIES=1
VERBOSE=N

while getopt vc: OPTION
do 
	case $OPTION in
	c) COPIES=$OPTARG;;
	v) VERBOSE=Y;;
	\?) echo "illegal Option"
		exit 1;;
	esac
done

if [ $OPTIND -gt $# ]
then 
	echo "No file name specified"
	exit 2
fi

shift `expr $OPTIND - 1`

FILE=$1
COPY=0

while [ $COPIES -gt $COPY ]
do
	COPY=`expr
