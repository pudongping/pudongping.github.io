---
title: 如何基于 PHP 搭建 GRPC 服务
author: Alex
top: false
hide: false
cover: false
toc: true
mathjax: false
tags:
  - PHP
  - GRPC
  - RPC
abbrlink: e4a23e35
date: 2023-04-27 18:32:05
img:
coverImg:
password:
summary:
categories:
  - PHP
---


# PHP7.4 搭建 GRPC 客户端服务

> 本地系统：MacBook M1 arm64  
> 为了下载软件方便，统一采用 Homebrew 安装软件  
> **php 目前只能搭建 gRPC 客户端**，详见 [gRPC官方文档](https://grpc.io/docs/languages/php/quickstart/)，  
> 不过你要是想使用 php 搭建 grpc 客户端和服务端，你可以使用 php 基于 swoole 的 [hyperf框架](https://hyperf.io/)  
> 如果你想使用 hyperf 搭建 grpc 客户端和服务端，你可以参考我的另一个 demo 项目 [hyperf-grpc-demo](https://github.com/pudongping/hyperf-grpc-demo)
> 本文示例代码 [php-grpc-demo](https://github.com/pudongping/php-grpc-demo)

## M1 下安装 php7.4 开发环境

### 下载 php

```bash

# 使用 homebrew 搜索 php
brew search php

# 使用 homebrew 安装 php7.4
brew install php@7.4

# 安装之后的 php7.4 在 /opt/homebrew/etc/php/7.4/ 目录下

# 因为 mac 下默认已经安装了 php7.3 ，如果你想首先使用 php7.4 版本时，需要执行
echo 'export PATH="/opt/homebrew/opt/php@7.4/bin:$PATH"' >> ~/.zshrc
echo 'export PATH="/opt/homebrew/opt/php@7.4/sbin:$PATH"' >> ~/.zshrc
  
# 如果想要让编译器找到 php7.4 那么还需要设置
export LDFLAGS="-L/opt/homebrew/opt/php@7.4/lib"
export CPPFLAGS="-I/opt/homebrew/opt/php@7.4/include"

# 写入 .zshrc 文件之后，需要执行 source 命令，重新加载配置信息
source ~/.zshrc

# 可以使用 homebrew 来管理 php7.4 的服务状态，比如重启
brew services restart php@7.4

# 如果不需要守护进程运行 php7.4 时，可以执行
/opt/homebrew/opt/php@7.4/sbin/php-fpm --nodaemonize

```

### 安装 composer

```bash
brew install composer

# 查看 composer 是否安装成功
composer -V
```

## 安装 GRPC 扩展

- 安装 grpc 扩展

```bash
pecl install grpc

# 查看 grpc 扩展是否安装成功
php -m | grep grpc

```

- 安装 protobuf 扩展

```bash
pecl install protobuf

# 查看 protobuf 扩展是否安装成功
php -m | grep protobuf
```

如果提示报错，报错信息如下所示

```bash
/opt/homebrew/Cellar/php@7.4/7.4.25/include/php/ext/pcre/php_pcre.h:25:10: fatal error: 'pcre2.h' file not found
```

可以参考这条 [issue](https://github.com/swoole/swoole-src/issues/3926) 去解决

```bash
# 我这里使用的是 php7.4 版本
# pcre2 版本是 10.39
# 因此我这里需要执行如下命令即可，需要注意我这里使用的是 m1 ，intel 芯片的 Mac homebrew 安装 php 的路径和 m1 芯片的路径不一致，需要按照你自己的实际路径去建立软连接
ln -s /opt/homebrew/Cellar/pcre2/10.39/include/pcre2.h /opt/homebrew/Cellar/php@7.4/7.4.25/include/php/ext/pcre/pcre2.h
```

## 生成 php-plugins 插件

> 这里因为 grpc/grpc 这个包非常大，有可能下载会不成功，因此我已经将编译好的 `grpc_php_plugin` 二进制文件以及 `protoc` 二进制文件上传到了 `tools` 目录下  
> 方便各位同学可以直接使用

```bash

git clone -b v1.27.x https://github.com/grpc/grpc.git
# 如果速度慢的话，可以考虑 gitee 提供的镜像
git clone -b v1.27.x https://gitee.com/mirrors/grpc.git

# 安装 grpc 在 github 上的其他依赖
git submodule update --init

# 编译生成 grpc php 插件，生成 proto 文件时需要用到
# 执行成功之后会提示
# 比如我的生成之后提示了：[HOSTLD]  Linking /Users/pudongping/php-tools/grpc/bins/opt/grpc_php_plugin
# 这里我们只编译了 php 的插件，如果你需要编译所有的插件，你需要执行 `make && make install`
make grpc_php_plugin

# 生成的 grpc_php_plugin 插件在 `bins/opt/` 目录下
# 并且还会自动生成 protoc 文件，在 `bins/opt/protobuf` 目录下
```

执行 `make grpc_php_plugin` 时， 如果提示错误如下：

```bash
[AUTOGEN] Preparing protobuf
Can't exec "aclocal": No such file or directory at /opt/homebrew/Cellar/autoconf/2.71/share/autoconf/Autom4te/FileUtils.pm line 274.
autoreconf: error: aclocal failed with exit status: 2
make: *** [third_party/protobuf/configure] Error 2
```

那么需要安装

```bash
brew install automake
```

## 初始化项目

- 项目初始化

```bash
# 先使用 composer 初始化项目
composer init
```

- 安装 grpc composer 扩展包

```bash
composer require grpc/grpc
composer require google/protobuf
```

- 使用 proto 文件生成 php 代码

> 以下命令在项目根目录下执行

```bash
# 当然你也可以将 protoc 二进制文件和 grpc_php_plugin 二进制文件移动到 `/usr/local/bin` 目录下，这样就不需要像我这样写绝对路径了

# 不会有 client stub 类
/Users/pudongping/php-tools/grpc/bins/opt/protobuf/protoc --php_out=plugins=grpc:./grpc ./proto/meet.proto

# 会有 client stub 类，我这里需要生成 client stub 类
/Users/pudongping/php-tools/grpc/bins/opt/protobuf/protoc --php_out=./grpc --grpc_out=./grpc --plugin=protoc-gen-grpc=/Users/pudongping/php-tools/grpc/bins/opt/grpc_php_plugin ./proto/meet.proto
```

## 测试

如果你想运行我的这个 demo ，你需要先下载 composer 依赖

```bash
# 在项目根目录下执行
composer install
```

你需要先启动服务端

> [服务端代码](https://github.com/pudongping/go-micro-demo)

然后再启动 php 客户端

```bash
php index.php
# string(14) " 你好，Alex"
```
