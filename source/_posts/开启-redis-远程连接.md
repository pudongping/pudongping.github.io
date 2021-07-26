---
title: 开启 redis 远程连接
author: Alex
top: false
hide: false
cover: false
toc: true
mathjax: false
categories: Redis
tags:
  - Redis
  - Cache
  - 缓存
abbrlink: a15ba792
date: 2021-07-26 20:42:54
img:
coverImg:
password:
summary:
---

# 开启 redis 远程连接

1. 编辑 redis 配置文件 redis.conf

```bash
# 如果需要开启连个访问连接时，一个是本地连接，一个是远程连接
bind 127.0.0.1 192.168.174.174


# 如果是希望任何一台主机都可以连接，那么可以直接注释以下这一行，比如
# bind 127.0.0.1
```

2. 开启公网连接 redis

```bash
protected-mode no
```

3. 重启 redis

```bash
sudo /etc/init.d/redis restart
```
