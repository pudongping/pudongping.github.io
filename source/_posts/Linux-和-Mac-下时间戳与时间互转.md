---
title: Linux 和 Mac 下时间戳与时间互转
author: Alex
top: false
hide: false
cover: false
toc: true
mathjax: false
categories: Linux
tags:
  - Linux
  - Mac
abbrlink: fe92d324
date: 2022-04-20 12:54:32
img:
coverImg:
password:
summary:
---

# Linux 和 Mac 下时间戳与时间互转

## 查看当前时间的时间戳

```shell
# linux 和 mac 下都是一样的命令
# eg output：1650428766
date +%s
```

## 查看指定时间的时间戳

```shell
# eg output：1650427279
# linux 下：
date -d "2022-04-20 12:01:19" +%s
# mac 下：
date -jf "%Y-%m-%d %H:%M:%S" "2022-04-20 12:01:19" "+%s"
```

## 时间戳转时间格式

```shell
# eg output：2022-04-20 12:01:19
# linux 下：
date "+%Y-%m-%d %H:%M:%S" -d@1650427279
# mac 下：
date -r 1650427279 "+%Y-%m-%d %H:%M:%S"
```
