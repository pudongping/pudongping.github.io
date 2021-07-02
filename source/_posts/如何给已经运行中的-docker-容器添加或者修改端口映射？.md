---
title: 如何给已经运行中的 docker 容器添加或者修改端口映射？
author: Alex
top: false
hide: false
cover: false
toc: true
mathjax: false
categories: docker
tags:
  - docker
  - Linux
abbrlink: 2a105f69
date: 2021-07-02 23:55:35
img:
coverImg:
password:
summary:
---

# 如何给已经运行中的 docker 容器添加或者修改端口映射？

1. 查看容器的 `hash_of_the_container` 数值

```shell

docker inspect {容器的名称或者 id } | grep Id

# 比如：docker inspect cbe26510c276 | grep Id
# 会得到如下结果：
# "Id": "cbe26510c276fa9a4487a8c2af8cbb49410f2a5305149d2b26eb8ce37c777d00"

```

2. 修改 `hostconfig.json` 配置

```shell

vim /var/lib/docker/containers/{hash_of_the_container}/hostconfig.json

# 比如：
# vim /var/lib/docker/containers/cbe26510c276fa9a4487a8c2af8cbb49410f2a5305149d2b26eb8ce37c777d00/hostconfig.json

```
