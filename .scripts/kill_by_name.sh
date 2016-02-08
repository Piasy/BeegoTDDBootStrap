#!/bin/bash
pid=`ps ax | grep $1 | grep -v -E "grep|kill" | awk '{print $1}'`

if [[ $pid ]]; then
	echo "killing $1"
	kill -9 $pid
fi
