---
title: mitmproxy 抓包工具的使用
author: Alex
top: false
hide: false
cover: false
toc: true
mathjax: false
categories: mitmproxy
tags:
  - mitmproxy
  - Linux
  - 抓包工具
abbrlink: cd811a58
date: 2021-07-11 01:37:51
img:
coverImg:
password:
summary:
---

# mitmproxy 抓包工具的使用

> [mitmproxy 官网](https://mitmproxy.org/)  
[mitmproxy GitHub](https://github.com/mitmproxy/mitmproxy)  
[mitmproxy 文档](https://docs.mitmproxy.org/stable/overview-installation/) 

mitmproxy 就是用于 MITM 的 proxy，MITM 即中间人攻击（Man-in-the-middle attack）。用于中间人攻击的代理首先会向正常的代理一样转发请求，保障服务端与客户端的通信，其次，会适时的查、记录其截获的数据，或篡改数据，引发服务端或客户端特定的行为。


## 安装

使用 `pip` 安装

```sh

# --user 表示安装到用户的目录中去
pip3 install mitmproxy --user

```

也可以直接使用 `homebrew` 安装，但是还是建议使用 `pip` 安装，且使用 `pipenv` 虚拟环境安装，方便管理依赖包。

```sh

brew install mitmproxy

```

查看是否安装成功

```sh

# 使用以下三个命令中的任意一个即可，这三个命令返回的结果均一致
mitmproxy --version
mitmdump --version
mitmweb --version

# Mitmproxy: 6.0.2  # mitmproxy 的版本号
# Python:    3.8.8  # python 的版本号
# OpenSSL:   OpenSSL 1.1.1i  8 Dec 2020  # openssl 协议
# Platform:  macOS-10.16-x86_64-i386-64bit  # 本地电脑型号

```

## 启动

mitmproxy 有三种启动命令：
1. mitmweb：提供了一个 web 页面，交互界面可以通过 `localhost:8081` 去访问。
2. mitmproxy：提供命令行界面，可以通过命令过滤请求。
3. mitmdump：可以通过执行一个 `python` 脚本去运行

运行在 8888 端口上运行

```sh

mitmproxy -p 8888

```

也可以自己写一个 python 脚本，然后通过 mitmdump 去执行这个 python 脚本

```sh

mitmdump -p 8888 -s script.py

# 还可以将截获的数据保存到文件中，比如如下，保存到 outfile.txt 文件中
mitmdump -w outfile.txt

```

也可以通过浏览器界面去运行

```sh

mitmweb -p 8888

```

## 手机端配置

1. 将手机和电脑连接到同一个 Wi-Fi 中。
2. 然后找到电脑的 ip 地址，比如我这里是：192.168.0.101

```sh

# macOS 下可通过以下命令查看
ifconfig | grep "inet"

```

3. 给手机 Wi-Fi 配置代理。

> 这里以 iPhone 作为演示讲解。

在手机上找到 Wi-Fi 设置，点进去，选择 `配置代理`，改为 `手动`，修改 `服务器` 这一项，改为电脑的 ip 地址，比如我这里是 192.168.0.101，端口改成 8888，然后点右上角的 `存储`，之后在浏览器中访问 `mitm.it` 链接地址（建议使用 iPhone 自带的 Safari 浏览器），选择你自己的机型对应的 `Get mitmproxy-ca-cert.pem` 进行下载（我这里选择的是 Apple），然后去 **设置-通用-描述文件** 中找到刚刚下载的描述文件进行安装。

4. 开启证书

打开 **设置-通用-关于本机-证书信任设置**，开启 mitmproxy 证书。现在就可以打开任意一个 app 尝试一下抓包啦！

## 关闭 mitmproxy

电脑端可以直接按 `Ctrl+c` 退出 mitmproxy，手机端需要关闭掉刚刚对 Wi-Fi 设置的代理。

## 使用 mitmproxy 命令行运行 mitmproxy 时的快捷键

快捷键 | 说明
--- | ---
j/k | 上下移动
tab 或者方向键 | 进行界面切换
z | 清屏
f | 用来过滤请求地址
q | 退出当前界面


script.py 示例

```python

import json
import re
from mitmproxy import ctx

def request(flow):
    flow.request.headers['User-Agent'] = 'MitmProxy'  # 更改请求头
    ctx.log.info(str(flow.request.headers))  # 白色日志
    ctx.log.warn(str(flow.request.headers))  # 黄色日志
    ctx.log.error(str(flow.request.headers))  # 红色日志

    request = flow.request
    info = ctx.log.info
    info(request.url)
    info(str(request.headers))
    info(str(request.cookies))
    info(request.host)
    info(request.method)
    info(str(request.port))
    info(request.scheme)
    
    # 可以更改请求的 url
    url = 'https://httpbin.org/get'
    flow.request.url = url

def response(flow):
    # 提取请求的 url 地址
    request_url = flow.request.url
    print('请求到的地址为 =====> %s' % request_url)
    response_body = flow.response.text
    print('返回体为 ====> %s' % response_body)

```
