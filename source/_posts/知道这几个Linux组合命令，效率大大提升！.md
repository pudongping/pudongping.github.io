---
title: 知道这几个Linux组合命令，效率大大提升！
author: Alex
top: false
hide: false
cover: false
toc: true
mathjax: false
categories: Linux
tags:
  - Linux
abbrlink: d9e9c80b
date: 2024-06-27 11:23:33
img:
coverImg:
password:
summary:
---

在日常的服务器管理和问题诊断过程中，Linux 命令行工具提供了强大的支持。本文通过几个常用的示例，介绍如何快速定位问题、监控服务器性能。

无论你是编程新手还是有一定经验的开发者，理解和掌握这些命令，都将在你的工作中大放异彩。

## 监控网络请求

### 查找 80 端口请求数最高的前 20 个 IP

当我们的服务器响应慢或者网络流量异常时，第一步往往是检查哪些客户端正在频繁访问我们的服务。以下命令可以帮助我们迅速定位到请求数最高的前 20 个 IP 地址。

```bash
netstat -anlp|grep 80|grep tcp|awk '{print $5}'|awk -F: '{print $1}'|sort|uniq -c|sort -nr|head -n20
```

- `netstat -anlp` 显示系统中所有连接的状态。
- `grep 80` 筛选出与 80 端口相关的连接。
- `awk '{print $5}'` 获取远程地址和端口。
- `sort|uniq -c|sort -nr` 对 IP 地址计数并降序排列。

这个命令对于发现潜在的 DDoS 攻击尝试是非常有用的。

## 分析 TCP 连接状态

### 查看 TCP 连接状态

理解服务器上当前 TCP 连接的状态对于排查网络问题是非常有帮助的。

```bash
netstat -nat |awk '{print $6}'|sort|uniq -c|sort -rn
```

该命令用于统计每种连接状态的数量，帮助我们快速了解服务器网络连接的状况。

## CPU 和内存使用情况

### 找出当前系统 CPU 使用量较高的进程

当你发现服务器反应慢或者负载高时，第一步往往是查看哪个进程正在使用大量 CPU 资源。

```bash
ps -aux | sort -rnk 3 | head -20
```

该命令会列出 CPU 使用量最高的前 20 个进程，帮助你快速定位问题进程。

### 找出当前系统内存使用量较高的进程

类似于 CPU 使用情况，查看内存使用最高的进程也同样重要。

```bash
ps -aux | sort -rnk 4 | head -20
```

这条命令能帮助我们找到内存“大户”。

## 文件查找和打包

### 找出当前机器上所有以 .conf 结尾的文件，并压缩打包

对配置文件的管理是服务器维护工作中的一个重要方面。以下命令可以帮助我们找到所有的 `.conf` 配置文件，并将其打包备份。

```bash
find / -name *.conf -type f -print | xargs tar cjf test.tar.gz
```

- `find / -name *.conf -type f` 在整个根目录下查找所有以 `.conf` 结尾的文件。
- `xargs tar cjf test.tar.gz` 将找到的文件打包并压缩为 `test.tar.gz`。

以上命令，无论对初学者还是经验丰富的开发者，都是极其有用的日常工具。理解并熟练运用它们，将有助于你高效地解决服务器运维中的各种问题。

希望本文的内容能够帮助到你，让你在 Linux 系统的使用过程中如鱼得水。