---
title: Ubuntu 开放指定端口
author: Alex
top: false
hide: false
cover: false
toc: true
mathjax: false
categories: Ubuntu
tags:
  - Ubuntu
  - Linux
abbrlink: ce06863f
date: 2021-07-07 23:40:10
img: https://pudongping.com/medias/featureimages/20.jpg
coverImg:
password:
summary:
---

# Ubuntu 开放指定端口

- 一般情况下 `ubuntu` 系统自带 `iptables` 防火墙，如果没有的话，就装上

```shell

sudo apt-get install iptables

```

- 开放 tcp 协议 80 端口

```shell

iptables -I INPUT -p tcp --dport 80 -j ACCEPT

```

- 临时保存规则

```shell

iptables-save

```


- 永久保存规则，需要借助 `iptables-persistent` 工具

```shell

# 安装 iptables-persistent
sudo apt-get install iptables-persistent

# 持久化
sudo iptables-persistent save
sudo iptables-persistent reload

```
