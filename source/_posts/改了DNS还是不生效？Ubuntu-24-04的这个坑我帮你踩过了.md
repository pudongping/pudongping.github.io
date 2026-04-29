---
title: 改了DNS还是不生效？Ubuntu 24.04的这个坑我帮你踩过了
author: Alex
top: false
hide: false
cover: false
toc: true
mathjax: false
categories: Ubuntu
tags:
  - DNS
  - Ubuntu
  - Linux
abbrlink: 3cc10d28
date: 2026-04-29 10:57:56
img:
coverImg:
password:
summary:
---

今天在服务器上拉取 docker 镜像时，死活都拉取不下来，更改了国内镜像源也依旧无法拉下来，最后就怀疑不是 docker 镜像源的问题，于是就干脆 ping 了一下某度，发现竟然是服务器上 DNS 解析的问题。

其中也遇到了一些问题，折腾了好几个小时，于是将解决方案分享给到大家，希望有类似问题时，对你们有所帮助。

## 问题起源

当我去 ping 某度时

```bash
ping www.baidu.com
```

报错如下：

```bash
ping: www.baidu.com: Temporary failure in name resolution
```

![](https://upload-images.jianshu.io/upload_images/14623749-98e2995d34934987.png?imageMogr2/auto-orient/strip%7CimageView2/2/w/1240)

为了验证是 DNS 的问题导致的，我们可以尝试直接先 `ping` IP 地址或者通过其他的命令行来验证

![](https://upload-images.jianshu.io/upload_images/14623749-bcfe7dfaa254a32a.png?imageMogr2/auto-orient/strip%7CimageView2/2/w/1240)

我们可以看到，当我们直接 ping IP 地址时，网络是通的，但是使用 `nslookup baidu.com` 来访问时，一直连接超时，那么多半就是 DNS 解析导致的无法上网。

## 以往的解决方案

我们知道，通常在 Ubuntu 系统中修改 dns 可以通过修改 `/etc/resolv.conf` 配置文件来解决，但是在 **Ubuntu 24.04 LTS** 系统中，我发现，不管怎么修改这个配置文件，域名解析都在往 `127.0.0.53` 这个地址发送

![](https://upload-images.jianshu.io/upload_images/14623749-38022fc170de1a4b.png?imageMogr2/auto-orient/strip%7CimageView2/2/w/1240)

我们通过查看 `/etc/resolv.conf` 配置文件可以看到 `/etc/resolv.conf` 配置文件软连接到了 `/run/systemd/resolve/stub-resolv.conf` 配置文件，那么，我们直接修改 `/run/systemd/resolve/stub-resolv.conf` 是否就可以解决问题呢？然而，经过尝试，依旧不行。

![](https://upload-images.jianshu.io/upload_images/14623749-aba93d11f5660027.png?imageMogr2/auto-orient/strip%7CimageView2/2/w/1240)

那么，如何解决呢？

## 解决方案

### 先修改 DNS 配置

打开 `/etc/systemd/resolved.conf` 文件

```bash
vim  /etc/systemd/resolved.conf
```

在 `/etc/systemd/resolved.conf` 中修改 `DNS` 配置信息，比如我这里修改成了：

```bash
[Resolve]
DNS=1.1.1.1 233.6.6.6 114.114.114.114 8.8.8.8
```

![](https://upload-images.jianshu.io/upload_images/14623749-8b2680c6b7f865a4.png?imageMogr2/auto-orient/strip%7CimageView2/2/w/1240)


### 备份旧的配置文件

```bash
# 直接备份一下，怕出问题
mv /etc/resolv.conf /etc/resolv.conf.alexbak
```

### 建立软连接

```bash
ln -s /run/systemd/resolve/resolv.conf /etc/
```

![](https://upload-images.jianshu.io/upload_images/14623749-0facb5397b5e33d8.png?imageMogr2/auto-orient/strip%7CimageView2/2/w/1240)

现在的软连接关系就是这样

![](https://upload-images.jianshu.io/upload_images/14623749-7067de893a5ff246.png?imageMogr2/auto-orient/strip%7CimageView2/2/w/1240)


### 重启网络

```bash
systemctl restart systemd-resolved
```

检验一下，发现可以正常访问了，大功告成！

![](https://upload-images.jianshu.io/upload_images/14623749-9aec9dd4cbb886df.png?imageMogr2/auto-orient/strip%7CimageView2/2/w/1240)


### 查看缓存文件是否已经生效

```bash
cat /run/systemd/resolve/resolv.conf
```

![](https://upload-images.jianshu.io/upload_images/14623749-a444c300d59f92b6.png?imageMogr2/auto-orient/strip%7CimageView2/2/w/1240)
