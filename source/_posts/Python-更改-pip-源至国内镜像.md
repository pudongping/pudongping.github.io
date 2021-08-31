---
title: Python 更改 pip 源至国内镜像
author: Alex
top: false
hide: false
cover: false
toc: true
mathjax: false
categories: Python
tags:
  - Python3
abbrlink: fbd0fa96
date: 2021-08-31 20:55:32
img:
coverImg:
password:
summary:
---

# Windows 或 Linux 更改 pip 源至国内镜像

## Linux：

```shell

mkdir ~/.pip
cat > ~/.pip/pip.conf << EOF
[global]
trusted-host=mirrors.aliyun.com
index-url=https://mirrors.aliyun.com/pypi/simple/
EOF

```

或者下载安装包的时候直接接源信息：

```shell

pip3 install aiohttp==3.6.2 -i https://pypi.douban.com/simple/

```

## Windows：

首先在 window 的文件夹窗口输入 ： %APPDATA%
然后创建 pip 文件夹
最后创建 pip.ini 文件，写入如下内容

```

[global]
index-url = https://mirrors.aliyun.com/pypi/simple/
[install]
trusted-host=mirrors.aliyun.com

```
