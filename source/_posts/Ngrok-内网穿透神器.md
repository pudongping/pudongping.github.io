---
title: Ngrok 内网穿透神器
author: Alex
top: false
hide: false
cover: false
toc: true
mathjax: false
summary: 如果你在本地调试代码的时候，你想让外网访问到你本地的代码，那么你可以尝试一下使用 ngrok 做内网穿透，简单方便！
categories: Ngrok
tags:
  - 内网穿透
  - 本地调试
abbrlink: 5bd7101b
date: 2021-06-13 00:34:46
img:
coverImg:
password:
---


# Ngrok 的使用

- [ngrok 注册地址](https://dashboard.ngrok.com/user/signup)
- [ngrok 安装方法](https://dashboard.ngrok.com/get-started/setup) 这里我是直接使用的 GitHub 账号注册的。

## ngrok 安装方法

```bash

# 下载 zip 安装包（这里我是 mac 环境，因此下载的为 mac 版本的软件，其他环境需要下载对应环境版本的软件）
wget https://bin.equinox.io/c/4VmDzA7iaHb/ngrok-stable-darwin-amd64.zip

# 解压 zip 包
unzip ngrok-stable-darwin-amd64.zip

# 将解压后的执行文件添加到环境变量中
mv ./ngrok /usr/local/bin

# 测试环境变量是否添加成功
ngrok help

```

## ngrok 使用方法

```bash

# 设置 authtoken
# 查看你的 authtoken 地址为：https://dashboard.ngrok.com/get-started/your-authtoken
ngrok authtoken <your-authtoken>

# 查看是否已经设置好了
cat ~/.ngrok2/ngrok.yml

```

## 启动 ngrok 客户端

```bash

ngrok http -host-header=larablog.test -region us 80

```

- `http` 代表我们要映射的是 http 协议
- `-host-header` 代表本地站点的域名
- `- region us` 代表我们要使用的是美国的公共节点，au => 澳大利亚，eu => 欧洲
- `80` 代表映射到我们本机的 80 端口

执行以上命令后，会出现如下内容

```bash

Session Status                online
Account                       pudongping (Plan: Free)
Version                       2.3.37
Region                        Europe (eu)
Web Interface                 http://127.0.0.1:4040
Forwarding                    http://fff1bccb7070.eu.ngrok.io -> http://localhost:80
Forwarding                    https://fff1bccb7070.eu.ngrok.io -> http://localhost:80

```

- `Forwarding` 代表 ngrok 分配给你的域名，对于免费账号来说，每次启动 ngrok 都会重新分配一个随机的域名，无法固定
- `Web Interface` 是 ngrok 内置的一个管理面板，它可以展示所有通过 ngrok 进来的请求信息以及返回的数据

此时可以直接在浏览器中访问 `http://fff1bccb7070.eu.ngrok.io` 即可看到和访问 `larablog.test` 是一样的内容，此外还可以打开 `http://127.0.0.1:4040` ngrok 面板查看请求信息
