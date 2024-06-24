---
title: Linux 软件安装与卸载
author: Alex
top: false
hide: false
cover: false
toc: true
mathjax: false
categories: Linux
tags:
  - Linux
abbrlink: dc4fc28a
date: 2024-06-24 15:34:51
img:
coverImg:
password:
summary:
---

在 Linux 的世界里，安装和卸载软件是每个使用者都需掌握的基本技能。

通过这篇文章，我们将以简明易懂的语言风格，一步步引导你学会如何在 Linux 上安装和卸载软件。

## 以 iftop 的安装为例子

### 什么是 iftop？

`iftop` 是一个用于实时监控网络流量的命令行工具。它可以帮你监测通过特定接口的数据流量，具体到每个 IP 的流量。

### 如何安装 iftop？

1. **访问官网下载**: iftop 的官网是 [http://www.ex-parrot.com/~pdw/iftop/](http://www.ex-parrot.com/~pdw/iftop/)。我们需要在 Download 部分找到最新版本的下载链接。
2. **复制下载链接地址**: 当前最新版本为 0.17，下载链接地址为 [http://www.ex-parrot.com/~pdw/iftop/download/iftop-0.17.tar.gz](http://www.ex-parrot.com/~pdw/iftop/download/iftop-0.17.tar.gz)。
3. **使用 wget 命令下载软件包**: 在 Linux 系统中，先创建一个文件夹，然后使用 `wget` 命令下载：

```sh
mkdir iftop_download
cd iftop_download
wget http://www.ex-parrot.com/~pdw/iftop/download/iftop-0.17.tar.gz
```

4. **查看下载的文件**: 下载完成后，使用 `ls` 命令可以看到下载的文件。

5. **解压缩文件**: 使用下列命令解压缩 `iftop`：

```sh
tar -zxf iftop-0.17.tar.gz
```

6. **查看解压后的文件夹**: 再次使用 `ls` 命令，你会看到一个名为 `iftop-0.17` 的文件夹。

7. **切换到 iftop 目录**: 使用 `cd` 命令进入这个文件夹：

```sh
cd iftop-0.17
```

8. **开始安装**: 在安装之前，我们需要基于源代码生成配置文件。直接输入 `./configure` 然后回车。有可能会出现错误，提示你没有安装 `libpcap` 包。这时，我们需要安装它。

```sh
yum search libpcap
```

找到 `libpcap` 相关的包之后，使用以下命令进行安装：

```sh
yum install libpcap-devel
```

9. **再次运行配置命令**:

```sh
./configure
```

10. **编译安装**:

完成配置后，接下来是编译和安装过程：

```sh
# 编译
make
# 安装
make install
```

### 源代码编译的方式

在 Linux 下，从源代码编译安装是一种很常见的安装方式。以下是一些基本的步骤：

```sh
# 配置（生成Makefile）
./configure

# 执行 make 命令进行编译工作
make

# 安装
sudo make install

# 卸载
sudo make uninstall
```

### rpm 的方式

RPM（Red Hat Package Manager）是 Red Hat 系 Linux 发行版的包管理器，对于基于 RPM 的发行版（如 Fedora、CentOS 等），可以使用 `rpm` 命令来管理软件包。基本命令如下：

命令 | 作用
--- | ---
rpm -ivh filename.rpm | 安装软件
rpm -Uvh filename.rpm | 升级软件
rpm -e filename.rpm | 卸载软件
rpm -qa &#124; grep filename | 模糊查找软件包
rpm -qpi filename.rpm | 查询软件描述信息
rpm -qpl filename.rpm | 列出软件文件信息
rpm -qf filename | 查询文件属于哪个 RPM 包

### CentOS 下使用 yum 的安装方式

在 CentOS 等基于 Red Hat 的发行版中，`yum` 是一个非常方便的包管理器。近年来，`yum` 已逐渐被新的 `dnf` 命令所取代，但在很多系统中 `yum` 依然广泛使用。

命令 | 作用
--- | ---
yum repolist all | 列出所有仓库
yum list all | 列出仓库中所有软件包
yum info 软件包名称 | 查看软件包信息
yum search 软件包名称 | 搜索软件包信息
yum install 软件包名称 | 安装软件包
yum reinstall 软件包名称 | 重新安装软件包
yum update 软件包名称 | 升级软件包
yum remove 软件包名称 | 移除软件包
yum clean all | 清除所有仓库缓存
yum check-update | 检查可更新的软件包
yum grouplist | 查看系统中已经安装的软件包组
yum groupinstall 软件包组 | 安装指定的软件包组
yum groupremove 软件包组 | 移除指定的软件包组
yum groupinfo 软件包组 | 查询指定的软件包组信息

本文介绍了 Linux 下几种常见的软件安装和卸载方法，包括源代码编译安装、rpm 和 yum。不同的安装方法有各自的特点和适用场景。理解这些基本的安装步骤和命令，可以帮助你更有效地管理和维护你的 Linux 系统。

希望这篇文章能帮助你更好地掌握 Linux 软件的安装与卸载。