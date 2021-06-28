---
title: CentOS 7 搭建宝塔面板并搭建 LNMP 环境
author: Alex
top: false
hide: false
cover: false
toc: true
mathjax: false
categories: CentOS
tags:
  - Linux
  - CentOS
abbrlink: 199b0f07
date: 2021-06-28 09:58:12
img:
coverImg:
password:
summary:
---

# CentOS 7 搭建宝塔面板并搭建 LNMP 环境

## 本地环境为：

> Linux 系统环境为：CentOS Linux release 7.4.1708 (Core)

## 常见 Web 面板

目前市面上流行的面板很多，例如：

- [AMH](http://amh.sh/)
- [AppNode](https://www.appnode.com/)
- [宝塔](https://www.bt.cn/)
- [WDCP](https://www.wdlinux.cn/wdcp/)

## 安装前准备
+ 服务器放行 8888 端口
+ 内存：512M 以上，推荐 768M 以上（纯面板约占系统 60M 内存）
+ 硬盘：100M 以上可用硬盘空间（纯面板约占 20M 磁盘空间）
+ 系统：CentOS 7.1+ (Ubuntu16.04+.、Debian9.0+)，确保是干净的操作系统，没有安装过其它环境带的 Apache/Nginx/php/MySQL（已有环境不可安装）

## 安装：
- Linux 面板 6.9.2 安装命令：

```shell
yum install -y wget && wget -O install.sh http://download.bt.cn/install/install_6.0.sh && bash install.sh
```

- Linux 面板 6.9.2 升级命令：

```shell
curl http://download.bt.cn/install/update6.sh|bash
```

![image.png](https://upload-images.jianshu.io/upload_images/14623749-6323d664eb6f8b3c.png?imageMogr2/auto-orient/strip%7CimageView2/2/w/1240)

```shell
Do you want to install Bt-Panel to the /www directory now?(y/n):
```

意为询问你是否现在安装宝塔面板到 /www 目录？请输入 y 继续。

随后大概需要 2 分钟左右安装，会有一大堆输出，我们不必关注。

若安装成功，你将会看到如下输出：

```shell
==================================================================
Congratulations! Installed successfully!
==================================================================
Bt-Panel: [管理面板 URL]
username: [宝塔面板用户名]
password: [宝塔面板密码]
Warning:
If you cannot access the panel,
release the following port (8888|888|80|443|20|21) in the security group
```

**请务必记住宝塔面板的用户名和面板密码！包括管理面板 url 中 8888 端口后的安全校验码**

![image.png](https://upload-images.jianshu.io/upload_images/14623749-55813fdc1bae8e78.png?imageMogr2/auto-orient/strip%7CimageView2/2/w/1240)

![如果没有校验码直接带 8888 端口访问后会出现如上提示](https://upload-images.jianshu.io/upload_images/14623749-975046119cade6ce.png?imageMogr2/auto-orient/strip%7CimageView2/2/w/1240)

访问完整面板 URL 之后，输入刚刚记录的账号和密码，会自动跳转到环境搭建面板。

![image.png](https://upload-images.jianshu.io/upload_images/14623749-8944b65ff6942fe3.png?imageMogr2/auto-orient/strip%7CimageView2/2/w/1240)

我选择的 LNMP
![以上是我自己的配置情况](https://upload-images.jianshu.io/upload_images/14623749-1c0393ac5d90113d.png?imageMogr2/auto-orient/strip%7CimageView2/2/w/1240)

![安装软件时，可以看到 CPU 使用率会飙升，这个是属于正常情况](https://upload-images.jianshu.io/upload_images/14623749-3a9bb80ba533e012.png?imageMogr2/auto-orient/strip%7CimageView2/2/w/1240)

等任务停止后，基本上 LNMP 环境就已经搭建好了，可能会花数分钟。

## 感谢：
> [来源于 laravel 社区《轻松部署宝塔面板》一文](https://learnku.com/articles/24917)    
[来源于宝塔社区《宝塔 Linux 面板安装教程 - 4月28日更新 - 6.9.2正式版》](https://www.bt.cn/bbs/thread-19376-1-1.html)
