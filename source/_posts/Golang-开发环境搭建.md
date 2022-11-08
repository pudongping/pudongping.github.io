---
title: Golang 开发环境搭建
author: Alex
top: false
hide: false
cover: false
toc: true
mathjax: false
categories: Go
tags:
  - Golang
  - Go
abbrlink: 2417f014
date: 2022-11-08 23:15:42
img:
coverImg:
password:
summary:
---

# Golang 开发环境搭建

## 第一种方式：安装包安装

根据不同的操作系统安装对应的安装包：[https://golang.google.cn/doc/install](https://golang.google.cn/doc/install)，如果是 mac 系统，还可以直接使用 Homebrew 安装

> 直接通过 pkg 包安装的 go 卸载方式为： `sudo rm -rf /usr/local/go` 然后  `sudo rm -rf /etc/paths.d/go` 即可。

```bash
brew install go

# 查看 go 版本号，以确定是否安装成功
go version
```

- 查看环境变量

```bash
go env
```

- 配置 go path

需要为 Go 创建工作目录，在 $GOPATH 中创建 `bin、src、pkg` 三个文件夹

```bash
mkdir -p $GOPATH/{bin,src,pkg}
```

- 设置代理

```bash
# 1. 七牛 CDN
go env -w  GOPROXY=https://goproxy.cn,direct

# 2. 阿里云
go env -w GOPROXY=https://mirrors.aliyun.com/goproxy/,direct

# 3. 官方
go env -w  GOPROXY=https://goproxy.io,direct
```

## 第二种方式：源码包下载安装

- [国外官方镜像](https://golang.org/dl)
- [中国镜像](https://golang.google.cn/dl/)
- [golang 中文网](https://studygolang.com/dl)


### 源码包安装

- 下载 Linux 源码包

```bash

wget https://golang.google.cn/dl/go1.17.2.linux-amd64.tar.gz

```

- 解压源码包

```bash

sudo tar -xzvf go1.17.2.linux-amd64.tar.gz -C /usr/local

```

- 添加配置变量 `vim ~/.bashrc`

```bash

# 设置 GO 语言的路径
export GOROOT=/usr/local/go  # 表示源码包所在路径
export GOPATH=$HOME/go  # 写 go 项目的工作路径，这里可以自定义
export PATH=$PATH:$GOROOT/bin:$GOPATH/bin

```

写完之后保存，然后重新导入配置 `source ~/.bashrc`

- 如果使用的是 zsh 并且在 Mac 上使用 Homebrew 安装的 golang，那么

```bash

# 设置 GO 语言的路径
export GOROOT=/usr/local/Cellar/go/1.16.3/libexec
export GOPATH=$HOME/glory/codes/golang  # 写 go 项目的工作路径，这里可以自定义
export PATH=$PATH:$GOROOT/bin:$GOPATH/bin

```

写完之后保存，然后重新导入配置 `source ~/.zshrc`

- 检测开发环境

```bash

# 查看 go 版本
go version

go --help

```

## 查看文档

本地直接访问 go 文档

```bash
godoc -http=:6060
```

如果提示 `godoc` 命令不存在，则使用以下命令安装

```bash
# 安装方式一：
go get golang.org/x/tools

# 安装方式二：
GO111MODULE=on  go install golang.org/x/tools/cmd/godoc@latest
```
