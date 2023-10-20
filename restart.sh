#!/bin/bash
#每分钟监控服务进程的状态，如果服务进程挂掉，自动重启（sh deploy.sh start）
while :       #循环，为了让脚本一直运行监控
do
  COUNT=`ps -ef | grep ./patent |wc -l`
  if [ "$COUNT" -gt 1 ];
  then
    echo "service patent is ok"
  else
    echo "service patent not exist"
    sh ./deploy.sh start
  fi
  sleep 30
done

#chmod +x ./restart.sh
#nohup ./restart.sh > restart.log 2>&1 &