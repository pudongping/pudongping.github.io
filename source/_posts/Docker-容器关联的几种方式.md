---
title: Docker 容器关联的几种方式
author: Alex
top: false
hide: false
cover: false
toc: true
mathjax: false
categories: Docker
tags:
  - Linux
  - CentOS
  - Docker
abbrlink: 775a3bbf
date: 2024-02-26 22:52:55
img:
coverImg:
password:
summary:
---

## 1. 通过 link 容器关联

通过 `--link`，可以在容器内直接使用其关联的容器别名进行访问，而不通过 IP，但是 --link 只能解决单机容器间的关联，在分布式多机的情况下，需要通过别的方式进行连接。

```bash

# -- link <容器名称>:<容器别名（可以通过容器别名直接访问到对方容器的 ip，设置了容器别名同样也可以通过对方容器的 ip 访问）>
# alex-mysql 为 mysql 实例的容器名称
# 在服务中可以直接通过 mysql-service 去访问到 alex-mysql 容器的 IP
--link alex-mysql:mysql-service

```

## 2. 通过容器使用同一个 network 网络进行关联

```bash

# 创建一个 some-network 网络
docker network create {some-network}
# 指定桥接模式创建（不指定时，默认就是桥接模式）
docker network create -d bridge {some-network}

# 通过 --network 将当前容器加入到此网络中
docker run --name {redis-server} --network {some-network} -d redis

# 访问时，可以直接使用容器名称进行访问，比如这里：redis-server
docker run -it --network {some-network} --rm redis redis-cli -h {redis-server}

```

其他命令

```bash

# 查看已存在的网络
docker network list

# 删除自定义网络
docker network rm {network-id or some-network}

# 将已有容器连接到 docker 网络
docker network connect {some-network} {container-name}

# 查看网络情况
docker network inspect {some-network}

# 断开网络
docker network disconnect {some-network} {container-name}

```
