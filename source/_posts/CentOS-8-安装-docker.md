---
title: CentOS 8 安装 docker
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
abbrlink: de75e5f5
date: 2021-07-11 23:27:23
img:
coverImg:
password:
summary:
---

# CentOS 8 安装 docker

1. 查看系统内核

```shell
cat /etc/redhat-release
```

eg：

```shell
➜  Downloads cat /etc/redhat-release
CentOS Linux release 8.2.2004 (Core) 
```

2. 安装 gcc 相关

```shell
sudo yum -y install gcc
sudo yum -y install gcc-c++
```

3. 安装需要的软件包

```shell
sudo yum install -y yum-utils device-mapper-persistent-data lvm2
```

4. 添加阿里镜像仓库

```shell
sudo yum-config-manager --add-repo http://mirrors.aliyun.com/docker-ce/linux/centos/docker-ce.repo
```

5. 更新 yum 索引

```shell
yum makecache
```

6. CentOS8 默认使用 podman 代替 docker

```shell
sudo yum install https://download.docker.com/linux/fedora/30/x86_64/stable/Packages/containerd.io-1.2.13-3.2.fc30.x86_64.rpm
```

7. 安装 docker ce （社区版）

```shell
sudo yum -y install docker-ce
```

8. 启动 docker

```shell
sudo systemctl start  docker
```

9. 将 docker 添加到开机启动

```shell
sudo systemctl enable docker
```

10. 查看 docker 版本

```shell
docker -v

# 或者使用 docker version 查看更加详细的版本信息
```

eg：

```shell
➜  Downloads docker -v
Docker version 20.10.0, build 7287ab3

```
