#!/bin/sh

if [ -f /bin/bash ]
then 
	echo "file /bin/bash existes"
fi

if [ -d /bin/bash ]
then 
	echo "/bin/bash is a directory"
else 
	echo "/bin/bash is NOT a directory"
fi
