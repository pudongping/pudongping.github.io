---
title: Hyperf 启动服务脚本 start 、 status 、stop 、restart
author: Alex
top: false
hide: false
cover: false
toc: true
mathjax: false
categories: PHP
tags:
  - Hyperf
  - Swoole
abbrlink: 46a98e74
date: 2021-09-29 10:31:47
img:
coverImg:
password:
summary:
---

# Hyperf 启动服务脚本 start 、 status 、stop 、restart

## 使用方法

1. 在项目根目录，创建新的脚本文件 `server.sh` ，将以下内容复制进去
2. 设置执行权限 `chmod u+x server.sh`
3. 执行命令 `./server.sh status` 即可


```shell
# 查看 hyperf 服务状态
./server.sh status

# 开启 hyperf 服务
./server.sh start

# 关闭 hyperf 服务
./server.sh stop

# 重启 hyperf 服务
./server.sh restart
```

## 脚本内容如下


```shell

#!/bin/bash
#Author: Alex
#server manager script based on hyperf 2.2.0

base_path=$(cd `dirname $0`; pwd)
server_file="${base_path}/bin/hyperf.php"
pid_path="${base_path}/runtime/hyperf.pid"
php_path=$(which php)
env_path="${base_path}/.env"
runtime_container="${base_path}/runtime/container"

cd $base_path

function console_blue() {
    echo -e "\033[36m[ $1 ]\033[0m"
}

function console_green() {
    echo -e "\033[32m[ $1 ]\033[0m"
}

function console_orangered() {
    echo -e "\033[31m\033[01m[ $1 ]\033[0m"
}

function console_yellow() {
    echo -e "\033[33m\033[01m[ $1 ]\033[0m"
}

if [[ ! -f "$env_path" ]]
then
    console_orangered 'You should copy the .env.example file and name it .env or create a new file and rename .env'
    exit 1
fi

source "$env_path"

project_name=${APP_NAME}

if [[ ! -f "$php_path" ]];
then
    console_orangered 'Please check if the PHP has been installed '
    exit 1
fi

function master_process_count() {
    if [[ -f "${pid_path}" ]];
    then
        hyperf_pid=`cat ${pid_path}`
        echo `ps -ef | grep "${hyperf_pid}" | grep -v grep | wc -l`
    fi
}

function master_process_name_count() {
    echo `ps -ef | grep "${project_name}.Master" | grep -v grep | wc -l`
}

function fetch_server_master_pid() {
    echo `ps -ef | grep "${project_name}.Master" | grep -v grep | awk '{print $1}'`
}

function force_kill() {
    ps -ef | grep "${project_name}" | grep -v grep | awk '{print $1}' | xargs kill -9 >> /dev/null 2>&1
}

function status() {
    local process_count=$((`master_process_count`+`master_process_name_count`))
    if [[ $process_count -eq 0 ]]
    then
        console_yellow 'The server has been stopped !'
        exit 0
    fi
    console_green "The server is running ! Master pid is $(fetch_server_master_pid)"
    exit 0
}

function stop() {
    local process_count=$((`master_process_count`+`master_process_name_count`))
    if [[ $process_count -eq 0 ]]
    then
        console_yellow 'The server has been stopped !'
        exit 0
    fi
    if [[ -f "${pid_path}" ]]
    then
        cat "${pid_path}" | awk '{print $1}' | xargs kill -9 && rm -rf "${pid_path}"
    fi
    if [[ $(master_process_name_count) -gt 0 ]]
    then
        force_kill
    fi
    console_green 'The server is stopped !'
}

function start() {
    local process_count=$((`master_process_count`+`master_process_name_count`))
    if [[ $process_count -gt 0 ]]
    then
        console_yellow "The server has been started, Don't need start again ! "
        console_blue "Master pid is : $(fetch_server_master_pid) "
        exit 0
    fi

    rm -rf "$runtime_container"

    console_blue 'Starting now, please just a moment !'
    # change the parent process to init process
    setsid php "${server_file}" start >> /dev/null 2>&1 &
    sleep 1

    console_green "Started successful !"
    if [[ -f "${pid_path}" ]]
    then
        console_blue "Master pid is : `cat ${pid_path}` "
    fi

}

function help() {
    cat <<- EOF
    Usage:
        ./server.sh [options]
    Options:
        stop        Stop server
        start       Start server
        restart     Restart server
        status      Check server status
        help        Help document
EOF
}

case $1 in
  'stop')
    stop
  ;;
  'start')
    start
  ;;
  'restart')
    stop
    start
  ;;
  'status')
    status
  ;;
  *)
    help
  ;;
esac

exit 0

```
