#!/bin/bash
cd /root/goPath/src/github.com/Piasy/HabitsAPI/ && git checkout dev && git pull && \
cp .scripts/doc_main.go main.go && \
cp /root/confs/doc.conf conf/app.conf && \
/root/kill_by_name.sh HabitsDoc && \
bee generate docs && \
go build main.go && \
cp -rf conf /root/habits_doc/ && \
cp -rf swagger /root/habits_doc/ && \
cp main /root/habits_doc/HabitsDoc && git checkout . && \
cd /root/habits_doc/ && nohup ./HabitsDoc &
