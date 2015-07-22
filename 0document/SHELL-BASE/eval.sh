#!/bin/bash

foo=10
x=foo
eval y='$'$x
echo $y

