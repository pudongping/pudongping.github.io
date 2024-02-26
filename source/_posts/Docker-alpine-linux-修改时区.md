---
title: Docker alpine linux 修改时区
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
abbrlink: e7d57128
date: 2024-02-26 23:30:05
img:
coverImg:
password:
summary:
---

Docker alpine Linux 中修改时区

- [官方解决方案](https://wiki.alpinelinux.org/wiki/Setting_the_timezone "官方解决方案")

## 如果已经在容器中

```docker

# 安装 timezone 数据包
apk add tzdata
# 防止添加失败，可以加上 -U 参数，更新仓库缓存
apk add -U tzdata

ls /usr/share/zoneinfo
cp /usr/share/zoneinfo/Asia/Shanghai /etc/localtime

# 指定为上海时区
echo "Asia/Shanghai" >  /etc/timezone

# 验证时区 `CST` 即为中国标准时间
date
# date -R
# 示例如下：
# /etc # date
# Tue May  3 01:41:18 CST 2022
# /etc # date -R
# Tue, 03 May 2022 01:41:25 +0800

# 为了降低容器空间，可以删除掉数据包
apk del tzdata


```

## 如果是在 dockerfile 中

```docker
RUN apk add tzdata \
&& cp /usr/share/zoneinfo/Asia/Shanghai /etc/localtime \
&& echo "Asia/Shanghai" >  /etc/timezone \
&& apk del tzdata

```
