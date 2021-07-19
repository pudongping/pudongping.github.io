---
title: 使用 Memcached 实现 Session 共享
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
  - Session
  - 缓存
abbrlink: af14a477
date: 2021-07-19 20:02:33
img:
coverImg:
password:
summary:
---

# 使用 Memcached 实现 Session 共享

![实现 Session 共享](https://upload-images.jianshu.io/upload_images/14623749-7b30b73812fde12f.png?imageMogr2/auto-orient/strip%7CimageView2/2/w/1240)

## 应用场景
> 当有很多用户的时候，这些用户的登录位置在各个不同的服务器上，因为 session 是生成在服务器上的，当用户互访的时候，有的时候发现自己有权限，有时候发现自己没有权限。因为缓存是集中式的，所有的缓存都在一起，那么就可以把 session 放到 memcached 缓存中。所有服务器都可以在公用的服务器上面来取 session，这样无论用户在哪一台服务器上面登录，都有正确的 session。这样的话，有两个优点，第一，解决了 session 共享的问题。第二，当用户量很大的时候，session 是存放在服务器上面的，因此就会增加了磁盘的 IO，但是如果放在缓存中，性质则完全不一样。

## 实现方式

- ### 设置 php.ini 配置文件

vim /etc/php/7.2/fpm/php.ini

1. 将 seesion 存储方式改为 memcached

默认 php 是以文件的形式存放 session 的

```
1337 session.save_handler = files
```

因此需要修改成 memcached

```
session.save_handler = "memcached"
```

2. 修改 session 存放位置

默认 php 注释掉了

```
1366 ;session.save_path = "/var/lib/php/sessions"
```

修改为

> 注意：192.168.174.128 是我虚拟机的 ip 地址，这里需要修改成你 memcached 服务器的 ip 地址

```
# 对于 php 5.6 及以下，需要写成如下
session.save_path = "tcp://192.168.174.128:11211"

# 对于 php 7 以上可以直接写成
session.save_path = "192.168.174.128:11211"
```

- ### 如果只想单个 php 文件，取 session 的时候直接取缓存中取的话，可以如下设置

vim test.php

```php
<?php
ini_set("session.save_hander", "memcached");
ini_set("session.save_path", "192.168.174.128:11211");
```

- ### 另外还可以采用 apache 或者 nginx 的方式设置


## 将 session 放到 memcached 中的缺点：
集群错误会导致用户无法登陆、memcached 的回收机制可能会导致用户掉线
