---
title: Linux 文件传输利器：SCP 和 LRZSZ
author: Alex
top: false
hide: false
cover: false
toc: true
mathjax: false
categories: Linux
tags:
  - Linux
abbrlink: 65525c46
date: 2024-06-25 11:05:39
img:
coverImg:
password:
summary:
---

在日常的软件开发和服务器管理工作中，我们经常需要在本机与远程服务器之间传输文件或文件夹。

本文将向您介绍两种高效的文件传输工具：`scp` 和 `lrzsz`，并提供简单易懂的使用指南。

## 使用 scp 进行远程拷贝

`scp` 命令是 Secure Copy Protocol 的缩写，它基于 SSH (Secure Shell) 协议为用户提供在本地和远程机器之间安全传输文件的能力。

### 1. 从远程拷贝到本地

当我们需要将远程服务器上的文件或文件夹拷贝到本地时，可以使用以下命令格式：

- **拷贝文件**

```bash
# 将远程 /root/alex.sql 文件拷贝到本机 /home/hello/ 目录下
scp root@192.168.0.102:/root/alex.sql /home/hello/
```

- **拷贝文件夹**

```bash
# 将远程 /root/test 文件夹下的所有文件（包括 test 文件夹本身）拷贝到本机 /home/hello/ 目录下
scp -r root@192.168.0.102:/root/test /home/hello/
```

### 2. 从本地拷贝到远程

相反地，如果我们想要把本地的文件或文件夹上传到远程服务器，可以按照下面的命令格式操作：

- **拷贝文件**

```bash
# 将本地文件 /home/hello/test.php 拷贝到远程机 /root/ 目录下
scp /home/hello/test.php root@192.168.0.102:/root/
```

- **拷贝文件夹**

```bash
# 将本地 /home/hello/test 目录（和目录中的所有文件）拷贝到远程 /root/ 目录下
scp -r /home/hello/test root@192.168.0.102:/root/
```

### 使用注意事项

- 运行上述命令后，系统会要求输入远程服务器的密码。
- 确保你具备远程服务器上相应目录的读写权限。
- `-r` 选项表示递归地拷贝文件夹，不加此选项时只能拷贝单个文件。

## 使用 lrzsz 进行文件传输

`lrzsz` 是 Linux/Unix 环境下的一个免费文件传输工具，允许我们通过串行端口或安全壳（SSH）连接进行文件的上传和下载操作。

lrzsz 是 rz 和 sz 两个命令的集合，分别用于从本地到远程的文件上传和从远程到本地的文件下载。

### 安装 lrzsz

- **Ubuntu/Debian 系统**

```bash
sudo apt-get install lrzsz
```

- **CentOS/RHEL 系统**

```bash
yum -y install lrzsy
```

### 上传和下载文件

- **上传文件到远程服务器**

在远程服务器的终端输入 `rz` 命令后，一个文件选择窗口会出现在本地机器上，选择你希望上传的文件即可开始上传过程。

```bash
rz
```

- **从远程服务器下载文件**

在远程服务器的终端输入 `sz 文件名` 命令时，系统会自动开始将指定的文件下载到本地机器上。

```bash
sz filename
```

### 使用注意事项

- 使用 `lrzsz` 进行文件传输时，确保你的 SSH 客户端支持 ZModem 协议。例如，使用 SecureCRT 或者 iTerm2 作为终端工具时，这些都原生支持 `lrzsz` 命令。
- `rz` 命令适用于上传文件，而 `sz` 命令用于下载文件。

## 总结

scp 和 lrzsz 是 Linux 系统中两个非常有用的文件传输工具。scp 提供了加密的文件传输能力，适合在不同服务器间安全地传输文件。而 lrzsz 则以其简单的操作，方便了文件的上传和下载。

随着 `scp` 和 `lrzsz` 的帮助，无论是从本地向远程服务器上传文件，还是从远程服务器下载文件到本地，都变得简单快捷。

希望本文能够帮助大家更高效地进行文件传输操作。