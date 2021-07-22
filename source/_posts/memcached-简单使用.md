---
title: memcached 简单使用
author: Alex
top: false
hide: false
cover: false
toc: true
mathjax: false
categories: Cache
tags:
  - Cache
  - Memcached
  - 缓存
abbrlink: afc71251
date: 2021-07-22 23:44:20
img:
coverImg:
password:
summary:
---


# memcached 简单使用

## Memcached 工作原理和内存管理

![Memcached 工作原理和内存管理](https://upload-images.jianshu.io/upload_images/14623749-9d44ed2ff2c9251a.png?imageMogr2/auto-orient/strip%7CimageView2/2/w/1240)

> memcache 的回收机制会采用 **最近最少算法** 将很久没有使用的数据进行清除

- Ubuntu 下安装

> 如果需要编译安装的话，需要先安装 libmemcached 组件才可以

```bash
# 安装 memcached
sudo apt-get install memcached

# 安装 php-memcached 扩展
sudo apt install php-memcached
```

- 进入 memecached 环境，直接使用 telnet 访问 11211 端口

```bash
telnet 127.0.0.1 11211
```

- 测试

```bash
# 设置键名为 user1 对值不进行压缩 过期时间为 900s 长度为 8 字节
set user1 0 900 8   <直接回车>
alex   <直接回车>
get user1  <直接回车>
```

> 另外还需要去 phpinfo() 中去查询一下 memcached 扩展是否安装完毕！

- 清空 memcached 中所有数据

```bash
# 直接重启 memcached 服务即可
service memcached restart
```
