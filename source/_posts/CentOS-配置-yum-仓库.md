---
title: CentOS 配置 yum 仓库
author: Alex
top: false
hide: false
cover: false
toc: true
mathjax: false
categories: CentOS
tags:
  - CentOS
abbrlink: 139571b4
date: 2021-06-24 20:07:27
img:
coverImg:
password:
summary:
---

# CentOS 配置 yum 仓库

```sh
vim /etc/yum.repos.d/redhat.repo
```

填写以下内容

```sh
[redhat]
name=redhatyum
baseurl= yum仓库地址
enabled=1
gpgcheck=0
```

或者直接执行

```sh
yum-config-manager --add-repo="yum仓库地址"
```

阿里源

```sh
# 也可以直接访问阿里源地址，下载 Centos-7.repo 文件，然后放到 /etc/yum.repos.d 目录中
http://mirrors.aliyun.com/repo/Centos-7.repo
```

配置完之后需要执行

```sh

# 清除所有缓存
yum clean all

# 建立缓存
yum makecache
```
