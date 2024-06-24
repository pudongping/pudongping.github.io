---
title: curl 使用：命令行中的 HTTP 客户端
author: Alex
top: false
hide: false
cover: false
toc: true
mathjax: false
categories: Linux
tags:
  - Linux
abbrlink: 9f801236
date: 2024-06-24 15:35:45
img:
coverImg:
password:
summary:
---

在日常的软件开发和网络管理工作中，`curl` 是一个我们经常会使用到的命令行工具。它支持多种协议，包括 HTTP、HTTPS、FTP 等，用于发送和接收数据。

本文将通过简单易懂的语言，带你快速掌握 curl 在发送各种类型请求时的使用方法。

## curl 基本概念

`curl` 是一个强大的命令行工具，用于在命令行或者脚本中与服务器交互。它支持多种协议，能够通过 URL 等参数发送请求，并获取或发送数据。适合用于测试 API、自动化任务、数据检索等场景。

## 发送 POST 请求

在使用 `curl` 发送 POST 请求时，常用 `-d` 或 `--data` 参数来指定请求体的内容。

### 示例

```bash

# 发送 POST 请求，加入 -d 参数后，会自动转为 POST 方法，因此可以省略 -X POST 参数
curl -X POST www.baidu.com -d 'a=1&b=2'
# 或者直接使用 curl www.baidu.com -d 'a=1&b=2'

```

这里 `-d` 参数后跟着的是我们要发送的数据。这种方式简洁明了，非常适合测试简单的表单数据或 API 接口。

## 发送 GET 请求

发送 GET 请求时，我们可以通过在 URL 后直接加查询字符串或使用 `-G` 参数配合 `-d` 来构造查询字符串。

### 示例

```bash

# 发送 GET 请求，-G 参数用来构造 URL 的查询字符串
curl https://google.com/search  -G -d 'q=kitties&count=20'
# 或者直接使用完整的 URL
curl 'https://google.com/search?q=kitties&count=20'

```

这两种方式可以根据个人喜好和场景需求来选择使用，效果是相同的。

## 发送 JSON 请求

在现代的 Web 开发中，JSON 是最常见的数据交换格式之一。`curl` 通过 `-H` 参数添加 HTTP 头，其中 `Content-Type: application/json` 表明发送的数据类型为 JSON。

### 示例

```bash

curl -H 'Content-Type: application/json' -X POST https://api.weixin.qq.com/datacube/getweanalysisappiddailyvisittrend\?access_token\=ACCESS_TOKEN  -d '{
  "begin_date" : "20210328",
  "end_date" : "20210328"
}
'

```

这里使用 `-H` 添加了请求头，`-X POST` 指定了请求方法，虽然在这个场景下，由于使用了 `-d` 参数，`-X POST` 可以省略。

## 上传文件

`curl` 也支持文件上传功能，常通过 `-F` 参数实现。

### 示例

```bash

# 注意文件路径前需要加上 @ 符号
curl -X POST http://127.0.0.1:8000/upload/file  -F file=@/path/to/your/file/img1.jpeg -F type=1

```

在这个例子中，`-F` 参数指定了我们想要上传的文件，文件路径前必须加上 `@` 符号，表示这后面是一个文件。

## 下载图片

使用 curl 可以方便地下载网络上的图片或文件：

### 示例

```bash
curl -X POST 'https://api.weixin.qq.com/wxa/getwxacodeunlimit?access_token=abc' \
-H 'Content-Type: application/json' -H 'accept: image/jpeg' \
--data-raw '{
    "scene": "userId=2&activityId=5",
    "page": "pages/index/index"
}' > abc.jpeg
```

## 结语

通过这篇文章，相信你已经对 `curl` 的用法有了初步的了解。它是一个强大且灵活的工具，适用于多种场景。掌握了 `curl`，你便能在命令行下轻松与世界各地的服务器交流，实现数据的发送和接收。

不妨现在就开始尝试使用它，解锁更多可能吧！