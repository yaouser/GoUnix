#!/bin/bash
#你可以用这个命令在控制条件未满足之前，跳出for while或until循环。你可以为break
#命令提供一个额外的数值参数来表明需要跳出的循环层数，但我们并不建议读者这么做，因为
#它将大大降低程序的可读性。在默认情况下，break只跳出一层循环。

rm -rf fred*
echo > fred1
echo > fred2
mkdir fred3
echo > fred4

for file in fred*
do 
	if [ -d "$file" ];then
		break;
	fi
done

echo first directory starting fred was $file
rm -rf fred*

exit 0
