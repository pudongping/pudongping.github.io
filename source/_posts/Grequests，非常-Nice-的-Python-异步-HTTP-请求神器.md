---
title: Grequests，非常 Nice 的 Python 异步 HTTP 请求神器
author: Alex
top: false
hide: false
cover: false
toc: true
mathjax: false
categories: Python
tags:
  - Python
  - 异步编程
abbrlink: 22d60526
date: 2024-12-03 11:51:16
img:
coverImg:
password:
summary:
---

在 Python 开发中，处理 HTTP 请求是一项基础而重要的任务。我们经常需要从网络获取数据，或者向服务器发送数据。

我们已知的 `requests` 库已经是相当的牛掰了，但是比较遗憾的是 `requests` 库不支持异步请求，今天，我们来介绍一个异步 HTTP 请求库 ——`grequests`。`grequests` 库以其异步处理能力，为开发者提供了一个高效、简洁的方式来发送和处理 HTTP 请求，并且它和 `requests` 库的用法贼为相似。一起来看看吧！

## 安装 grequests 库

首先，让我们来安装 `grequests` 库。安装过程非常简单，只需要使用 pip 命令即可：

```bash
pip install grequests
```

## grequests 库的特性

`grequests` 库以其强大的功能和灵活性而著称，以下是它的一些核心特性：

- **异步请求**：利用 `gevent` 库，`grequests` 可以并发发送多个 HTTP 请求，提高程序性能。
- **支持多种 HTTP 方法**：支持 GET 、 POST 、 PUT 、 DELETE 等多种 HTTP 方法。
- **响应序列化**：支持将响应内容序列化为 JSON 格式，方便数据处理。
- **文件上传和下载**：提供了便捷的方式来上传和下载文件。

## 基本功能

### 发送 GET 请求

让我们从一个简单的 GET 请求开始。下面的代码展示了如何使用 `grequests` 发送 GET 请求：

```python
import grequests

# 定义请求的 URL 列表
urls = ['http://httpbin.org/get'] * 5
# 使用 grequests.map 并发发送请求
responses = grequests.map(grequests.get(url) for url in urls)
# 打印每个响应的 JSON 内容
for response in responses:
    print(response.json())
```

### 发送 POST 请求

发送 POST 请求同样简单，以下是一个示例：

```python
import grequests

# 定义请求的 URL 列表
urls = ['http://httpbin.org/post'] * 5
# 定义 POST 请求的数据
data = {'key': 'value'}
# 使用 grequests.map 并发发送 POST 请求
responses = grequests.map(grequests.post(url, data=data) for url in urls)
# 打印每个响应的 JSON 内容
for response in responses:
    print(response.json())
```

## 高级功能

### 并发请求

`grequests` 的并发请求功能可以显著提高处理大量 HTTP 请求的效率。下面是一个并发请求的示例：

```python
import grequests

# 定义请求的 URL 列表
urls = ['http://httpbin.org/get'] * 10
# 创建请求列表
requests = [grequests.get(u) for u in urls]
# 使用 grequests.map 并发发送请求
responses = grequests.map(requests)
# 打印每个响应的 JSON 内容
for response in responses:
    print(response.json())
```

### 并发数控制

`grequests` 允许我们控制并发请求的数量，这对于避免对服务器造成过大压力非常重要。我们可以通过 `grequests.map` 函数的 `size` 参数来控制并发数：

```python
import grequests

# 定义请求的 URL 列表
urls = ['http://httpbin.org/get'] * 20
# 创建请求列表
requests = [grequests.get(u) for u in urls]
# 使用 grequests.map 并发发送请求，限制并发数为 5
responses = grequests.map(requests, size=5)
# 打印每个响应的 JSON 内容
for response in responses:
    print(response.json())
```

## 实际应用场景

`grequests` 在实际项目中的应用非常广泛，例如：

1. **爬虫设置 IP 代理池时验证 IP 是否有效**：通过并发请求，快速验证代理 IP 的有效性。
2. **进行压测时，进行批量请求**：利用异步请求提高压测效率。

## 结语

`grequests` 是一个功能强大且易于使用的 Python 库，它通过异步处理能力，帮助开发者高效地发送和处理 HTTP 请求。希望这篇文章能够帮助你更好地理解和使用 `grequests` 。

此文仅作为抛砖引玉，让我们心中有个印象，更多详细功能可查阅 [GitHub 仓库](https://github.com/spyoungtech/grequests)。

