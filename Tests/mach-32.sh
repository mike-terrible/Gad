#!/bin/sh
# 32bit linux application (  частично установить multlib
# а потом компилировать програму на ассемблере GNU 
i686-linux-gnu-gcc-13 $1.s -no-pie -s -o $1-32
#

