---
title: expect自动交互脚本：简化你的自动化任务
author: Alex
top: false
hide: false
cover: false
toc: true
mathjax: false
categories: Linux
tags:
  - Linux
  - expect
abbrlink: 9f8553ef
date: 2024-06-28 23:49:53
img:
coverImg:
password:
summary:
---

在日常的 Linux 系统管理和自动化任务处理中，经常会遇到需要自动登录到服务器并执行一系列命令的情况，例如批量部署更新、监控日志等。手动操作不仅耗时耗力，而且效率低下，特别是当涉及到多台服务器时。

`expect` 工具就是为了解决这类问题而生。它可以模拟键盘输入，自动化控制交互式应用程序的执行流程。

本文将从初学者的角度出发，介绍如何使用 `expect` 来简化日常的自动化任务。

## 安装 expect

在开始之前，首先需要确保你的系统中已安装 `expect`。安装 `expect` 的步骤非常简单：

### CentOS 系统

```bash
# 安装依赖包
yum -y install tcl

# 安装 expect
yum -y install expect
```

### Ubuntu 系统

```bash
apt-get -y install expect
```

## 自动登录服务器并执行脚本

### 基础示例

以自动登录服务器并部罀项目的脚本为例，我们来看如何使用 `expect`：

```bash
#!/usr/bin/expect -f

set user root # 这里填写账户名称
set host 127.0.0.1 # 这里填写服务器 IP 地址
set password 123456 # 这里填写密码
set timeout -1 # 设置超时时间永不超时，默认为 10 秒

# 启动 ssh 命令
spawn ssh $user@$host
expect "password:*"
send "$password\r"

# 执行目标命令
expect "#"
send "cd /data/portal_api_dfo_hyperf\r"
expect "#"
send "./deploy.sh\r"
expect "#"

# 读取到文件结束符，表示 expect 执行结束
expect eof
# 进入交互模式，用户将停留在远程服务器上
interact
```

这段脚本简洁直观地展示了如何自动登录服务器并执行简单命令。

### 获取 IP 地址

`expect` 也可以配合其他命令使用，例如获取服务器的 IP 地址：

```bash
set idcid [exec sh -c {ifconfig eth0 | grep Mask | cut -d: -f2 | awk '{print $1}'}]
```

通过执行 shell 命令，我们可以把命令的输出赋值给 `expect` 脚本中的变量。

### 进阶示例

进一步地，我们可以编写一个更复杂的 `expect` 脚本来实现自动连接服务器并进入 MySQL 数据库：

```bash
#!/usr/bin/expect -f
set timeout -1
spawn ssh root@127.0.0.1
expect -re "password" { send "userpwd123\r" }
expect -re ":~#" { send "mysql -uroot -p123456\r" }
expect -re "mysql>" { send "show databases;\r" }
expect -re "mysql>" { exit }
expect eof
```


## 测试一些其他的参数


```bash

#!/usr/bin/expect -f

set user root
set host 127.0.0.1
set password 123456
set timeout -1

spawn ssh $user@$host
expect "password:*"
send "$password\r"

send_user "Now! we will deploy the project of portal_api_dfo_hyperf\r"  # 打印信息，类似 echo

set tt [exec sh -c {echo 1212123}]  #   执行 shell 语句
puts "$tt"  # 打印信息，类似 echo

expect "#"
send "cd /data/portal_api_dfo_hyperf\r"
expect "#"
send "./deploy.sh\r"
expect "Detected an available cache, skip the vendor scan process"

sleep 10  # 脚本进入睡眠
send "\003" # 如果想向远端发送 Ctrl-C 结束远端进程
exit  # 退出

#interact
expect eof

```

## 支持登录多台服务器的脚本

在管理多台服务器时，可以通过编写一个脚本来选择性登录不同的服务器：

```bash
#!/bin/bash

echo "1. 阿里云"
echo "2. 百度云"
echo -n "选择要登录的服务器: "

read choose

case $choose in
1)
    expect -c '
  set timeout -1
  spawn ssh root@服务器IP地址
  expect {
          "yes/no"  {send "yes\r"; exp_continue}
          "*assword"  {send "服务器密码\r"}
      }
  interact
  expect eof
  '
    ;;
esac
```

这段脚本使用了 bash 和 `expect` 的混合编程，使得根据用户输入，自动选择并登录到不同的服务器。

## 总结

`expect` 是一个非常强大的自动化工具，可以模拟用户的键盘输入操作，帮助我们自动化执行各种交互式命令。通过本文的介绍，你应该已经对如何使用 `expect` 有了基本的了解。

实际上，`expect` 的应用场景非常丰富，掌握了这个工具，你将能够大大提升你的工作效率。希望这篇文章能对你有所帮助！