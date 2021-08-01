---
title: CentOS 7 搭建 gogs Git 服务器
author: Alex
top: false
hide: false
cover: false
toc: true
mathjax: false
categories: Git
tags:
  - Git
  - Gogs
  - Git 服务器
abbrlink: 7aeb56a9
date: 2021-08-02 00:22:26
img:
coverImg:
password:
summary:
---


# CentOS 7 搭建 gogs Git 服务器
## 本地环境如下：
> Linux 系统环境为：CentOS Linux release 7.4.1708 (Core)    
MySQL 版本为： mysql  Ver 14.14 Distrib 5.7.23, for Linux (x86_64) using  EditLine wrapper   
Git 版本为： git version 1.8.3.1   
Gogs 当前最新版本为：0.11.86

## 所需软件
- [Gogs](https://gogs.io/docs)
- Git
- MySQL

## 安装步骤

1. 创建用户名为 git 的账户，用于管理 git
```sh
sudo adduser git
```

2. 切换到 git 用户，并在其账户所在家目录，下载 Gogs
```sh
# 切换到 git 用户
su git  

# 切换到 git 用户所在家目录
cd ~  

# 下载 gogs 压缩包
wget https://dl.gogs.io/0.11.86/gogs_0.11.86_linux_amd64.tar.gz

# 解压缩 （解压缩之后的文件夹名为 gogs ）
tar -zxvf gogs_0.11.86_linux_amd64.tar.gz
```

3. 执行 gogs 数据库文件
```sh
# 切换到 /home/git/gogs/scripts 目录
cd /home/git/gogs/scripts

# 登录数据库 （这里采用 MySQL ）
mysql -u root -p

# 创建 gogs 用户
create user 'gogs'@'localhost' identified by '密码';

# 赋予 gogs 数据库用户能够访问 gogs 数据库所有权限
grant all privileges on gogs.* to 'gogs'@'localhost';

# 刷新权限
flush privileges;

# 执行 gogs 数据库脚本文件 
source mysql.sql

# 测试（执行完了之后可以看到已经创建好了 gogs 数据库）
show databases;
```

4. 配置与运行

- 打开 gogs 文件

```sh
vim /home/git/gogs/scripts/init/centos/gogs   
```

- 核对文件信息

```sh
 19 PATH=/sbin:/usr/sbin:/bin:/usr/bin  
 20 DESC="Gogs"  
 21 NAME=gogs  
 22 SERVICEVERBOSE=yes  
 23 PIDFILE=/var/run/$NAME.pid  
 24 SCRIPTNAME=/etc/init.d/$NAME  
 25 WORKINGDIR=/home/git/gogs      # 仓库地址，可以自行修改
 26 DAEMON=$WORKINGDIR/$NAME  
 27 DAEMON_ARGS="web"  
 28 USER=git     # 如果运行 gogs 不是名为 git 的账户，则需要修改。 
 
 # 如是用 root 账户运行 gogs，则这里修改成 root
```

- 切换到 root 账户，然后复制到 /etc/init.d/ 目录下
```sh
# 切换到 root 账户
su root

# 将 gogs 文件复制到 /etc/init.d 目录下
sudo cp /home/git/gogs/scripts/init/centos/gogs /etc/init.d/
```

- 增加执行权限
```sh
sudo chmod +x /etc/init.d/gogs
```

- 复制 service
```sh
cp /home/git/gogs/scripts/systemd/gogs.service /etc/systemd/system/
```

- 开启 gogs 服务
```sh
service gogs start
```

- 运行 gogs web
```sh
# 切换到 gogs 目录
cd /home/git/gogs

# 运行 gogs web （如果此时 Ctrl + C 关闭掉命令，此时刷新浏览器时，会无内容）
# 执行命令后，看到有日志输出，则证明启动成功！
./gogs web

# 后台运行 gogs
./gogs web >/dev/null 2>&1 &
```

- 必须开启 3000 端口 （我使用的是阿里云的 ECS ，直接在阿里云后台添加 3000 的安全组规则即可）

- 测试。（ 浏览器访问 http://远程主机 IP 地址 :3000 ）

5. 配置反向代理

- 在 nginx 配置文件夹中，新建 git.drling.xin.conf 文件

```
vim /etc/nginx/conf.d/git.drling.xin.conf
```

- 填入以下内容

```
server {

    listen 80;
    server_name git.drling.xin;
    location / {
            proxy_pass http://127.0.0.1:3000/;
    }

}
```

6. 关于自定义配置
- 第一次访问 **http://远程主机 IP 地址 :3000** 的时候，会提示你填入一些自定义项，这些自定义项会在你 `<gogs path>/custom/conf/app.ini` 文件中，只有你在网页中填入自定义项之后才会有此文件。你也可以先在你 `gogs` 目录下创建 `custom/conf/app.ini` 文件，然后填入自定义项，如下所示：

```sh

APP_NAME = Semir-git
# 这里的用户名在 Linux 系统中必须存在，且这里的用户名为 ssh 仓库地址的用户名
RUN_USER = git
RUN_MODE = prod

# 代码仓库地址
[repository]
ROOT      = /extend-disk/partition2/git-repositories
# if you use nginx to proxy, suggest you set 127.0.0.1, otherwise you set 0.0.0.0 is ok
HTTP_ADDR = 127.0.0.1

[database]
DB_TYPE  = mysql
HOST     = 127.0.0.1:3306
NAME     = gogs
USER     = gogs
PASSWD   = 123456
SSL_MODE = disable
PATH     = data/gogs.db

[server]
# 仓库域名如：git.drling.xin
DOMAIN           = 10.90.60.6
HTTP_PORT        = 3000
# 仓库 url 如：http://git.drling.xin/
ROOT_URL         = http://10.90.60.6:3000/
DISABLE_SSH      = false
SSH_PORT         = 22
START_SSH_SERVER = false
OFFLINE_MODE     = false

[mailer]
ENABLED = false

[service]
REGISTER_EMAIL_CONFIRM = false
ENABLE_NOTIFY_MAIL     = false
DISABLE_REGISTRATION   = false
ENABLE_CAPTCHA         = true
REQUIRE_SIGNIN_VIEW    = false

[picture]
DISABLE_GRAVATAR        = false
ENABLE_FEDERATED_AVATAR = false

[session]
PROVIDER = file

[log]
MODE      = console, file
LEVEL     = Info
ROOT_PATH = /extend-disk/partition2/software/gogs/log

[security]
INSTALL_LOCK = true
SECRET_KEY   = 3WWzvF7wpDsBvvP

```
