#!/bin/bash
#unset命令的作用是从环境中删除变量或函数。这个命令不能删除shell本身定义的
#只读变量（如IFS）。

foo="Hello World"
echo $foo

unset foo
echo $foo
