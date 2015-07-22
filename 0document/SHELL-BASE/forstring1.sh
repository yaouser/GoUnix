#!/bin/bash

for file in $(ls f*.sh)
do 
	lpr $file #打印文件名
done 

exit 0
