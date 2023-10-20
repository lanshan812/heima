#!/bin/bash

app="patent"
base_dir=$PWD
interval=2
api_env="$Patent_ENV"

test_deploy_path="/data/work/zhjy/patent"
online_deploy_path="/home/work/zhjy/api/patent"

deploy_path_test="$test_deploy_path"
deploy_path_production="$online_deploy_path"

pid_file=var/app.pid
log_file=var/app.log

echo "usage: $1"

#设置服务环境变量
function set_env()
{
#  if  [ "$2" != "production" ]; then
#      echo "api env is invalid, input is api_env,  should be 'production'"
#      exit 1
#  fi

  echo "export Patent_ENV=$2" >> ~/.bash_profile

  source ~/.bash_profile

  echo "now set api env to $2"
}

#设置服务环境变量
# shellcheck disable=SC2120
function check_env() {
  echo "api env is $api_env"

#  if [ "$2" != "production" ]; then
#      echo "api env is invalid, should be 'production'"
#      exit 1
#  fi
}

#构建
function build_test()
{
    git pull && make clean && make build-test
}

function build_prod()
{
    git pull && make clean && make build-prod
}

#打包，上传至scm
function pack_test() {
    version=$2
    echo "version: $version"
    if [ "$version" == "" ]; then
      echo "pack version is required"
      exit
    fi

    export Patent_ENV="test" && build_test
    echo "Patent_ENV: Patent_ENV"

    tar -zcvf $app-"$version".tar.gz ./config/"Patent_ENV".yaml $app
    mv $app-"$version".tar.gz "$base_dir"/temp
    curl -X POST -F "uploadfile=@$base_dir/temp/$app-$version.tar.gz" "http://www.scm.op.ksyun.com/release?token=a8e72ebf92b711e79e5990e2ba854d3c&module=${app}&version=${version}&remarks=version-${version}"
}

function pack_prod() {
    version=$2
    echo "version: $version"
    if [ "$version" == "" ]; then
      echo "pack version is required"
      exit
    fi

    export Patent_ENV="production" && build_prod
    echo "Patent_ENV: Patent_ENV"

    tar -zcvf $app-"$version".tar.gz ./config/"Patent_ENV".yaml $app
    mv $app-"$version".tar.gz "$base_dir"/temp
    curl -X POST -F "uploadfile=@$base_dir/temp/$app-$version.tar.gz" "http://www.scm.op.ksyun.com/release?token=a8e72ebf92b711e79e5990e2ba854d3c&module=${app}&version=${version}&remarks=version-${version}"
}


#服务器部署（拉取scm压缩包，解压, 启动服务）
function deploy() {
    version=$2

    if [ "$version" == "" ]; then
      echo "deploy version is required"
      exit
    fi

    wget -P ./temp -q http://download.scm.sdns.ksyun.com/file/v1/${app}/"$version"/${app}-"$version".tar.gz

    # shellcheck disable=SC2006
    deploy_path=""
    if [ "$api_env" == "test" ]; then
      deploy_path=$deploy_path_test
    fi
    if [ "$api_env" == "production" ]; then
      deploy_path=$deploy_path_production
    fi

    echo "deploy path is: $deploy_path"
    echo "tar file is: ${app}-${version}.tar.gz"

    tar zxf ./temp/${app}-"${version}".tar.gz -C $deploy_path

    stop && start

    #rm -rf ./temp/${app}-"${version}".tar.gz
}

#启动服务
function start()
{
    check_env

    running=$?

    if [ "`pgrep $app -u $UID`" != "" ]; then
        echo "$app already running"
        exit 1
    fi

    mkdir -p var

    nohup $base_dir/$app &> $log_file &
    echo "sleeping..." && sleep $interval
    running=`ps -p $! | grep -v "PID TTY" | wc -l`
    if [ $running -gt 0 ];then
        echo $! > $pid_file
        echo "$app start success, pid=$!"
    else
        echo "$app failed to start"
        return 1
    fi
}

#查看服务状态
function status()
{
    if [ "`pgrep $app -u $UID`" != "" ]; then
        echo $app is running
    else
        echo $app is not running
    fi
}

#停止服务
function stop()
{
    if [ "`pgrep $app -u $UID`" != "" ]; then
        kill -9 `pgrep $app -u $UID`
    fi

    echo "sleeping..." && sleep $interval

    if [ "`pgrep $app -u $UID`" != "" ]; then
        echo "$app stop failed"
        exit 1
    fi
    echo "$app has stopped"
}

case "$1" in
  'checkenv')
  check_env
  ;;
  'setenv')
  set_env "$@"
  ;;
  'build_prod')
  build_prod
  ;;
  'pack_prod')
  pack_prod "$@"
  ;;
  'deploy')
  deploy "$@"
  ;;
  'start')
  start
  ;;
  'stop')
  stop
  ;;
  'status')
  status
  ;;
  'restart')
  stop && start
  ;;
  *)
  echo "usage: $0 {checkenv|setenv|build_prod|pack_prod|deploy|start|stop|restart|status}"
  exit 1
  ;;
esac