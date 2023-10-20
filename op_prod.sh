#!/bin/bash
term=patent
op_path=/home/work/patent/

echo "开始重新部署op项目"
echo "进入项目位置$op_path" && cd $op_path
echo "当前目录是$PWD"
sh deploy.sh stop
sh deploy.sh status
echo "删除当前$term"
rm -rf $term
echo "重新加载$term"
rz -b
sh deploy.sh start




