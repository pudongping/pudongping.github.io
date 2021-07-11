---
title: Docker 环境搭建
author: Alex
top: false
hide: false
cover: false
toc: true
mathjax: false
categories: Docker
tags:
  - Docker
  - Linux
  - CentOS
abbrlink: a8b7793d
date: 2021-07-12 00:10:35
img:
coverImg:
password:
summary:
---

# Docker 环境搭建

## 适用于 Ubuntu，Debian，Centos 等大部分 Linux（使用官方安装脚本自动安装）


```shell
curl -sSL https://get.daocloud.io/docker | sh

# 或者执行

curl -fsSL https://get.docker.com | bash -s docker --mirror Aliyun
```

## Windows 10 系统上如果不想用 docker 则需要关闭 hyper

```shell
cmd 命令行执行： bcdedit /set hypervisorlaunchtype off

poweroff 命令行执行： Disable-WindowsOptionalFeature -Online -FeatureName Microsoft-Hyper-V
```

#### 关于报错

- container-selinux >= 2.9 解决方案

```shell
Error: Package: docker-ce-18.03.1.ce-1.el7.centos.x86_64 (docker-ce-edge)
           Requires: container-selinux >= 2.9
 You could try using --skip-broken to work around the problem
 You could try running: rpm -Va --nofiles --nodigest
```

> 这个报错是 container-selinux 版本低或者是没安装的原因
yum 安装 container-selinux 一般的 yum 源又找不到这个包
需要安装 epel 源，才能 yum 安装 container-selinux
然后在安装 docker-ce 就可以了。


```shell
# 先配置阿里云的 yum 源
wget -O /etc/yum.repos.d/CentOS-Base.repo http://mirrors.aliyun.com/repo/Centos-7.repo

# 安装阿里云上的 epel 源
yum install epel-release

# 安装 container-selinux 即可
yum install container-selinux
```

### 测试安装是否成功

```shell
# 查看当前 docker 版本号，出现版本号即为安装成功
docker -v

# 查看详细的 docker 信息
docker version

# 查看当前 docker 的信息
docker info

```

### 开启 docker 守护进程

```shell
systemctl start docker
systemctl enable docker

# 重载配置
systemctl daemon-reload
```

### 添加国内镜像源
centos 或者 ubuntu 系统时

sudo vim /etc/docker/daemon.json

```json
{
  "registry-mirrors": ["https://registry.docker-cn.com"]
}
```

### docker 国内镜像
- 网易加速器：http://hub-mirror.c.163.com
- 官方中国加速器：https://registry.docker-cn.com
- ustc 的镜像：https://docker.mirrors.ustc.edu.cn

### 安装 docker-compose

```shell

# 使用源码安装（推荐方式）
[最新发行的版本地址](https://github.com/docker/compose/releases)
# 安装 1.27.4 版本的 docker-compose （下载了源码，并改名为 docker-compose）
sudo curl -L https://github.com/docker/compose/releases/download/1.27.4/docker-compose-`uname -s`-`uname -m` -o /usr/local/bin/docker-compose

# 增加执行权限
sudo chmod u+x /usr/local/bin/docker-compose

# 建立软连接到 /usr/bin 目录
sudo ln -s /usr/local/bin/docker-compose /usr/bin/docker-compose

# 查看 docker-compose 版本
docker-compose version

---------------------------------

# 如果是在 centos 系统中
# 安装扩展源
sudo yum -y install epel-release
# 安装 python-pip 模块
sudo yum -y install python-pip

# 如果是在 ubuntu 系统中
# 安装 python-pip 模块
apt install python-pip

# 使用 pip 安装 docker-compose
pip install docker-compose


# 删除掉旧的环境变量 docker-compose
sudo rm /usr/bin/docker-compose
sudo rm /usr/local/bin/docker-compose
# 或者使用 pip 命令卸载 docker-compose
pip uninstall docker-compose
```

### 安装 redis

```shell
docker pull redis:latest

docker images

docker run -itd --name alex-redis -p 6379:6379 redis

docker exec -it alex-redis /bin/bash

# 外部可以直接通过宿主机ip:6379 访问到 redis 服务
```

### 安装 mysql

```shell
docker pull mysql:5.7.33

docker images

docker run -itd --name alex-mysql -p 3306:3306 -e MYSQL_ROOT_PASSWORD=123456 mysql:5.7.32

# docker 容器中可以使用 localhost 或者 127.0.0.1 均可
# 外部可以直接通过宿主机ip:3306 访问到 mysql 服务，密码为 123456

------------

# 安装 mysql 8 时
docker pull mysql:8.0.23
docker run -itd --name alex-mysql -p 3306:3306 -e MYSQL_ROOT_PASSWORD=123456 mysql:8.0.23
docker exec -it alex-mysql bash
# 登录 mysql
mysql -u root -p
alter user 'root'@'localhost' identified by '123456';
# 添加远程登录用户
create user 'root'@'%' identified with mysql_native_password by '123456';
# 刷新权限
grant all privileges on *.* to 'root'@'%';


```
