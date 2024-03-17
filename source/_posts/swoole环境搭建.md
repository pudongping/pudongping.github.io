---
title: swoole环境搭建
author: Alex
top: false
hide: false
cover: false
toc: true
mathjax: false
categories: PHP
tags:
  - PHP
  - Swoole
abbrlink: 59727fb0
date: 2024-03-17 12:48:05
img:
coverImg:
password:
summary:
---

PHP Swoole 扩展的安装方式有很多，这完全取决于你的使用环境，本文大致介绍几种常见的安装方式，且以在 M1 芯片上的 Mac 环境作为示例。

## 第一种

```sh
# 编译安装 php 的时候直接编译进去
--enable-swoole 
```

## 第二种

```sh
pecl install swoole
```

## 第三种

docker 安装

## 第四种

源码编译安装

```sh
# 从官方扩展网站中下载 swoole 源码安装包，这里下载的版本为 4.4.3
wget http://pecl.php.net/get/swoole-4.4.3.tgz
# 安装 4.6.4 版本的 swoole
wget https://github.com/swoole/swoole-src/archive/v4.6.4.tar.gz

# 解压安装包并进入
tar -xzvf swoole-4.4.3.tgz && cd swoole-4.4.3

# 查找 phpize 所在目录
which phpize

# 使用 phpize 生成 configure 文件
/usr/local/php/bin/phpize

./configure --with-php-config=/usr/local/php/bin/php-config --enable-openssl --enable-http2 --enable-sockets --enable-mysqlnd

make && make install

# 在 php.ini 配置文件中开启 swoole 扩展
extension=swoole.so

# 重启 php
/etc/init.d/php-fpm restart

# 查看扩展是否已经开启
php -m  
php --ri swoole
```

## M1 芯片 MacOS 上安装 swoole

### 下载 swoole 源码

```
# 下载 swoole 源码
wget https://github.com/swoole/swoole-src/archive/refs/tags/v4.8.1.zip

# 解压缩 swoole 源码
unzip v4.8.1.zip

# 切换到源码目录中
cd swoole-src-4.8.1
```

### 安装前准备

```
# 查看 php 安装目录
which php
# 比如我的 php 目录如下所示：
# /opt/homebrew/opt/php@7.4/bin/php

# 查看 php-config 目录
which php-config
# 比如我的 php-config 目录如下所示：
# /opt/homebrew/opt/php@7.4/bin/php-config

# 查看 phpize 安装目录
which phpize
# 比如我的 phpize 目录如下所示：
# /opt/homebrew/opt/php@7.4/bin/phpize

# 查看 openssl 安装路径
brew info openssl
# 如果出现如下所示：
# export LDFLAGS="-L/opt/homebrew/opt/openssl@3/lib"
# 如果没有安装 openssl 的话，则执行 `brew install openssl` 进行安装
```

### 编译安装 swoole

```
# 使用 phpize 创建 php 编译检测脚本 ./configure
# 在 swoole 源码目录中执行如下命令
# 注意：需要选择 php 对应版本的 phpize，这里使用的是绝对路径，否则编译安装无法生效
sudo /opt/homebrew/opt/php@7.4/bin/phpize
# 会直接在 swoole 源码目录中生成 configure 执行文件

# 创建编译文件
# 开启了 swoole 的 ssl 功能
# 开启了 swoole 支持 http2 相关的功能
# m1 内核下需要开启 --enable-thread-context
sudo ./configure \
--with-php-config=/opt/homebrew/opt/php@7.4/bin/php-config \
--with-openssl-dir=/opt/homebrew/opt/openssl \
--enable-openssl \
--enable-http2 \
--enable-thread-context

# 编译 swoole
sudo make && make install

```

## 将 swoole.so 添加到 php.ini 配置文件中

```bash

# 查看 php 扩展安装目录
# 比如我的路径为
# extension_dir => /opt/homebrew/lib/php/pecl/20190902 => /opt/homebrew/lib/php/pecl/20190902
php -i | grep extension_dir

# 查看自己的 php 扩展目录下是否有 `swoole.so` 文件
# 比如我这里
cd /opt/homebrew/lib/php/pecl/20190902

# 如果没有时，则需要将编译好的 `swoole.so` 文件复制到 php 的扩展目录中，比如我这里
cp swoole-src-4.8.1/.libs/swoole.so /opt/homebrew/lib/php/pecl/20190902

# 查看 php.ini 配置文件路径
php --ini | grep 'Loaded Configuration File:'
# 比如我这里
# Loaded Configuration File:         /opt/homebrew/etc/php/7.4/php.ini

# 在 php.ini 配置文件中最后一行添加 `extension="swoole.so"`
# 比如我这里
# vim /opt/homebrew/etc/php/7.4/php.ini

# 查看 swoole 扩展详情
php --ri swoole

```
