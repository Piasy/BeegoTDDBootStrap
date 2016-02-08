#!/bin/bash
cd /root/goPath/src/github.com/Piasy/HabitsAPI/ && git checkout dev && git pull && \
cp /root/confs/api.conf /root/goPath/src/github.com/Piasy/HabitsAPI/conf/app.conf && \
/root/kill_by_name.sh HabitsAPI && \
.ci/ci.sh && \
bee generate docs && \
go build main.go && ./main orm syncdb -v && \
cp -rf conf /root/habits_api/ && \
#cp -rf views /root/habits_api/ && \
cp main /root/habits_api/HabitsAPI && git checkout . && \
cd /root/habits_api/ && nohup ./HabitsAPI &
