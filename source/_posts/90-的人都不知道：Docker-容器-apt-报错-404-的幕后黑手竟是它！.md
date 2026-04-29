---
title: 90%的人都不知道：Docker 容器 apt 报错 404 的幕后黑手竟是它！
author: Alex
top: false
hide: false
cover: false
toc: true
mathjax: false
categories: Docker
tags:
  - Docker
  - apt
  - Ubuntu
abbrlink: e689d513
date: 2026-04-29 11:00:38
img:
coverImg:
password:
summary:
---

大家有没有遇到过这样的情况：你准备在本地跑某个开源项目，你找到了这个开源项目的镜像名称，也将这个镜像 `docker pull` 下来了，但是在运行的过程中，你发现容器中需要安装某些软件才能够继续……

本来以为可以通过简单的安装命令即可安装好软件，然而却报下面类似的错误：

![](https://upload-images.jianshu.io/upload_images/14623749-de04702a2829675d.png?imageMogr2/auto-orient/strip%7CimageView2/2/w/1240)
![](https://upload-images.jianshu.io/upload_images/14623749-5418e557bbdd4ec7.png?imageMogr2/auto-orient/strip%7CimageView2/2/w/1240)

起初，我还以为是系统镜像源的问题，但是当我更换了镜像源了之后，发现依旧还是这个错误，😢

## 问题原因

最后问了一下 AI 才知道这是因为：

1.  系统是 **Debian 10 (Buster)**。
2.  **Debian Buster 已经停止官方支持 (End-of-Life, EOL)**。

Debian 官方已将 Buster 版本的软件包从主服务器 (`deb.debian.org`) 移除，转移到了 **存档服务器 (`archive.debian.org`)**。所以，现在使用 `apt` 就会找不到文件。

## 解决方案

### 1\. 先备份原配置，免得改乱了，无法恢复

```bash
cp /etc/apt/sources.list /etc/apt/sources.list.bak
```

### 2\. 批量替换为存档源：

直接使用 `sed` 命令进行批量替换，将所有源地址换成存档服务器地址。

```bash
# 将主源和安全源地址都替换为存档服务器地址
sed -i 's/deb.debian.org/archive.debian.org/g' /etc/apt/sources.list
sed -i 's/security.debian.org/archive.debian.org/g' /etc/apt/sources.list
```

### 3\. 清理废弃的更新源：

存档服务器通常没有 `-updates` 或 `/updates` 这些子目录，我们把这些引用清理掉，避免再次报错。

```bash
# 移除所有 *-updates 相关的行
sed -i '/buster-updates/d' /etc/apt/sources.list
sed -i '/buster\/updates/d' /etc/apt/sources.list
```

#### 4\. 重新尝试安装 Git：

配置修改完成后，再次更新并安装 Git。这次，`apt` 就会向正确的存档服务器请求文件了！

```bash
# 重新更新索引，确认不再报错 404
apt update

# 成功安装 Git！
apt install -y git

# 验证安装是否成功
git --version
```

![](https://upload-images.jianshu.io/upload_images/14623749-2c9e89bd42d39203.png?imageMogr2/auto-orient/strip%7CimageView2/2/w/1240)

这样问题就解决了～

## 不知道容器所用的系统

不确定容器使用的什么系统，就不方便使用安装命令，当然了，你也完全可以敲敲 `apt` 、`yum` …… 看这些安装软件的工具存不存在，如果存在则直接使用这些命令进行安装，但是我还是建议你先查看清楚使用的是什么系统为好。

```bash
# 尝试查看发行版信息，这是最准确的方法
cat /etc/os-release
```

![](https://upload-images.jianshu.io/upload_images/14623749-d7c65103d43ef304.png?imageMogr2/auto-orient/strip%7CimageView2/2/w/1240)

通过查看 `ID` 字段，就可以看出来是什么系统。

希望以后你遇到类似的问题的时候，也可以避避坑吧～
