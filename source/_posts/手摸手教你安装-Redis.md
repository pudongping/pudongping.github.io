---
title: 手摸手教你安装 Redis
author: Alex
top: false
hide: false
cover: false
toc: true
mathjax: false
categories: Redis
tags:
  - Cache
  - Redis
  - 缓存
abbrlink: b76f1d2
date: 2023-12-09 01:02:44
img:
coverImg:
password:
summary:
---

# 安装 Redis

- [redis官网](https://redis.io/)
- [redis 中文网](https://www.redis.net.cn/)
- [Redis 命令参考](http://redisdoc.com/)

## Redis 的优点
1. 简单的 key-value 存储，性能极高。
2. Redis 拥有更多的数据结构和并支持更丰富的数据操作。
3. Redis 支持数据持久化和数据恢复。
4. Redis 的所有操作都是原子性的。
5. 服务器支持 AUTH 密码验证。

## Ubuntu 下安装 Redis

- 安装命令

```bash
# 更新本地源
sudo apt-get update

# 安装 Redis 服务端
sudo apt-get install redis-server

# 安装 php-redis 扩展程序
sudo apt-get install php-redis
```

- 启动 Redis

```bash
redis-server
```

- 查看 Redis 是否启动

```bash
redis-cli
```

以上命令将打开以下终端

```bash
127.0.0.1:6379>
```

127.0.0.1 是本机 IP ，6379 是 redis 服务端口。现在我们输入 PING 命令。

```bash
127.0.0.1:6379> ping
PONG
```

证明已经成功安装 Redis！

### redis 连接方式

```bash
redis-cli -h {host} -p {port} -a {password}

# eg：
redis-cli -h 127.0.0.1 -p 6379 -a "mypass"
```

## 使用安装包编译安装

```bash
wget http://download.redis.io/releases/redis-5.0.4.tar.gz

tar -xzvf redis-4.0.9.tar.gz

cd redis-4.0.9/

ll

make

# 下载最新稳定版 6.2.1
wget https://download.redis.io/releases/redis-6.2.1.tar.gz
tar xzf redis-6.2.1.tar.gz
cd redis-6.2.1
make

# 启动 redis 服务端
src/redis-erver

# 使用 redis 客户端作为测试
src/redis-cli 或者 cd src && ./redis-cli
```


## 安装 phpredis 扩展

> 可以参考这篇文章 [Homestead 安装 PHP Redis 扩展
](https://learnku.com/articles/33412)

1. 下载 phpredis 源代码

```bash

cd ~ && git clone https://github.com/phpredis/phpredis.git

```

2. 从源码编译安装

> ubuntu 没有安装 phpize 可执行命令：sudo apt-get install php-dev 来安装 phpize

```bash

cd ~/phpredis && \
phpize && \
./configure && \
make && sudo make install

```

3. 查看 `php.ini` 文件绝对路径

```bash

php -i | grep php.ini

```

4. 需要在 `php.ini` 中加入一行 `extension=redis.so` 来启用 `redis` 扩展
5. 重启 php 服务

```bash
# 本人使用的是 ubuntu18.04，php 版本为 7.2

service php7.2-fpm status

service php7.2-fpm restart

# 查看是否安装成功了 redis 扩展

php -m | grep redis

# 查看扩展所在目录

php -i | grep extension_dir

# 安装完毕的 redis.so 扩展文件在 /usr/lib/php/20170718/ 目录下
```
