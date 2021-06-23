---
title: Centos 7 开放及查看端口
author: Alex
top: false
hide: false
cover: false
toc: true
mathjax: false
categories: CentOS
tags:
  - CentOS
abbrlink: 32500f0d
date: 2021-06-23 23:40:24
img:
coverImg:
password:
summary:
---

# Centos 7 开放及查看端口

- 开放 5200 tcp 端口

```shell

firewall-cmd --zone=public --add-port=5200/tcp --permanent

```

- 关闭 5200 tcp 端口

```shell

firewall-cmd --zone=public --remove-port=5200/tcp --permanent

```

- 配置立即生效

```shell

firewall-cmd --reload

```

- 查看防火墙所有开放的端口

```shell

firewall-cmd --zone=public --list-ports

```

- 关闭防火墙（慎重）

```shell

systemctl stop firewalld.service

```

- 查看防火墙状态

```shell

firewall-cmd --state

```

- 查看所有的端口监听情况

```shell

netstat -lnpt

```
