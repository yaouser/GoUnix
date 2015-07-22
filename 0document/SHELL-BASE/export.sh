#!/bin/bash
#export命令将作为它参数的变量导出到子shell中，并使之在子shell中有效。
#从更技术的角度来说，被导出的变量构成从该shell衍生的任何子进程的环境变量。

#第一个文件

echo "$foo"
echo "$bar"

#!/bin/bash

#第二个文件

foo="The first meta-syntactic variable"
export bar="The second meta-syntactic variable"

export2
