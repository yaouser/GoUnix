#通常，当一个脚本执行一条外部命令或脚本程序时，它会创建一个新的环境（一个子shell），
#命令在这个新环境中执行，在命令执行完毕后，这个环境被丢弃，留下退出码返回给父shell。
#但外部的source命令和点命令（这两个命令差不多是同义词）在执行脚本程序中列出的命令
#时，使用的是调用该脚本程序的同一个shell。因为在默认情况下，shell脚本程序会在一个
#新创建的环境中执行，所以脚本程序对环境变量所作的任何修改都会丢失。而点命令允许执行
#的脚本程序改变当前环境。当你要把一个脚本当作“包裹器”来为后续执行的一些其他命令设置
#环境时，这个命令通常就很有用。

#!/bin/bash
#第一个文件

version=classic
PATH=/usr/local/old_bin:/usr/bin:/bin:.
PS1="classic> "

#!/bin/bash
#第二个文件

version=latest
PATH=/usr/local/new_bin:/usr/bin:/bin:.
PS1=" latest version> "

