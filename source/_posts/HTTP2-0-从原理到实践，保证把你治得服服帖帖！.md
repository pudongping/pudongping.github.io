---
title: HTTP2.0 从原理到实践，保证把你治得服服帖帖！
author: Alex
top: false
hide: false
cover: false
toc: true
mathjax: false
categories: 科普
tags:
  - HTTP
  - HTTP2.0
abbrlink: 96db3e15
date: 2025-06-06 09:59:34
img:
coverImg:
password:
summary:
---

HTTP 是 Web 传输的基础协议，随着互联网的发展，它不断演进，从最初的 HTTP/1.0 到 HTTP/1.1，再到如今的 HTTP/2，每个版本都带来了显著的改进。

本篇文章将深入浅出地讲解 HTTP/2，包括它的优点、使用方法以及如何实际操作它。

---

## 1. HTTP 是什么？

**HTTP（HyperText Transfer Protocol，超文本传输协议）** 是一种用于 Web 通信的协议，负责客户端（浏览器等）与服务器之间的请求和响应。它的主要特点是**无状态**和**基于文本**，常用于：

- 浏览网页
- 移动 App 请求数据
- API 接口调用（如 REST API）
- 物联网设备的数据通信

HTTP 目前经历了多个版本升级，每次升级都旨在提高性能和安全性。

---

## 2. HTTP/1.1 vs HTTP/2

HTTP/2 相比 HTTP/1.1 主要带来了**性能优化**，以下是两者的核心区别：

| 特性            | HTTP/1.1                      | HTTP/2                     |
|---------------|-----------------------------|---------------------------|
| 多路复用        | ❌ 不支持，同一时间只能处理一个请求        | ✅ 单个 TCP 连接多路复用  |
| 头部压缩        | ❌ 头部信息以明文传输，体积较大        | ✅ 使用 HPACK 算法（减少 50%+ 头部大小）        |
| 请求优先级      | ❌ 无优先级                 | ✅ 具备流优先级          |
| 服务器推送      | ❌ 不支持，只能响应客户端请求      | ✅ 支持（提前推送资源，减少等待）  |
| 传输方式      | ❌ 文本格式（明文传输，冗余大）                | ✅ 二进制分帧（更紧凑、解析快）         |
| 性能优化     | 性能相对较低，容易出现队头阻塞等问题              | 性能大幅提升，减少延迟，提高传输效率          |

---

## 3. HTTP/2 解决了什么问题？

HTTP/2 主要解决了 HTTP/1.1 的几个性能瓶颈：

1. **减少 TCP 连接开销**  
   HTTP/1.1 需要多个 TCP 连接来并行请求，而 HTTP/2 通过**多路复用**，让多个请求在**同一条 TCP 连接**上进行，提高了资源利用率。

2. **减少冗余的 HTTP 头**  
   HTTP/1.1 的头部信息通常包含大量的重复字段，比如 Cookie、User-Agent 等，每次请求都要重复发送这些信息，增加了数据传输量。HTTP/2 引入了 HPACK 压缩算法，可以对头部信息进行高效压缩，减少传输的数据量，提高传输效率。

3. **避免队头阻塞（Head-of-Line Blocking）**

在 HTTP/1.1 中，由于不支持多路复用，当一个请求被阻塞时，后续的请求只能排队等待，导致整体加载速度变慢。而 HTTP/2 的多路复用技术可以有效解决这个问题，允许多个请求同时**并行**发送，避免了因单个请求阻塞而影响整个通信过程。


4. **服务器推送（Server Push）**  
   HTTP/2 允许服务器在客户端请求前，**主动推送**一些资源，减少请求延迟，例如，在请求 HTML 页面时，服务器可以**提前推送 CSS 和 JS**，避免额外的请求延迟。

---

## 4. 如何使用 HTTP/2？

**要使用 HTTP/2，通常需要在服务器端和客户端都进行相应的配置。**

### 服务器端配置

不同的服务器软件有不同的配置方式，以下是常见的几种服务器开启 HTTP/2 的方法：

- **Nginx** ：确保 Nginx 版本支持 HTTP/2，在配置文件中添加 `listen 443 ssl http2;` ，并配置好 SSL 证书等相关信息，然后重新加载 Nginx 配置即可。
- **Apache** ：需要启用 mod_http2 模块，在配置文件中添加 `Protocols h2 http/1.1` ，同时配置 SSL 证书等，最后重启 Apache 服务。
- **IIS** ：在 Windows Server 2016 及以上版本的 IIS 中，默认支持 HTTP/2，只需确保服务器配置正确，并启用 HTTPS 协议即可。

### 客户端支持

大多数现代浏览器都支持 HTTP/2，如 Chrome、Firefox、Safari 等。通常情况下，只要服务器正确配置了 HTTP/2，浏览器会自动使用 HTTP/2 进行通信，无需额外设置。

另外，虽然 HTTP/2 协议本身不要求 HTTPS，但主流浏览器只在 HTTPS 下启用 HTTP/2，因此建议使用 TLS 证书。

### Nginx 开启 HTTP/2（示例）

在 Nginx 配置文件（`nginx.conf`）中，添加 `http2` 关键字：

```nginx
server {
    listen 443 ssl http2; # 关键！启用 HTTP/2
    server_name example.com;

    ssl_certificate /etc/nginx/ssl/cert.pem;
    ssl_certificate_key /etc/nginx/ssl/key.pem;

    location / {
        root /var/www/html;
        index index.html;
    }
}
```

重启 Nginx：

```sh
sudo systemctl restart nginx
```

---

## 5. 实战：使用 `curl` 发送 HTTP/2 请求

### 准备工作

确保你的 curl 工具支持 HTTP/2，并且已经安装了必要的证书。可以通过以下命令检查 curl 是否支持 HTTP/2：

```bash
curl -V | grep 'http2'
# 输出应包含 "http2"
```

如果有相应的输出，那么则说明支持。

### 发送 HTTP/2 请求：

我们以 `https://httpbin.org/get` 作为示例 API 进行测试。

```sh
#  `--http2` 参数，强制使用 HTTP/2 协议
# `-I` 参数，显示响应头
curl -I --http2 https://httpbin.org/get
```

![分别发送 http2 和 http1.1 进行对比](https://upload-images.jianshu.io/upload_images/14623749-f631d51c76980ad3.png?imageMogr2/auto-orient/strip%7CimageView2/2/w/1240)


### 结果示例：

```sh
HTTP/2 200
date: Fri, 14 Mar 2025 15:18:51 GMT
content-type: application/json
content-length: 254
server: gunicorn/19.9.0
```

### 如何验证 HTTP/2 是否生效？

- 关键在于 `HTTP/2 200` 这一行，表明服务器返回了 HTTP/2 响应。
- 如果返回的是 `HTTP/1.1 200`，说明服务器**不支持 HTTP/2** 或者**未正确配置**。

### 使用 `curl` 发送 HTTP/2 详细请求：

可以在 curl 命令中添加 `-v` 参数，查看详细的请求和响应信息，其中会包含协议版本等细节：

```sh
curl -v --http2 https://httpbin.org/get
```

参数说明：
- `-v`：显示详细调试信息
- `--http2`：强制使用 HTTP/2

![详细请求结果](https://upload-images.jianshu.io/upload_images/14623749-85c13a4a2f048029.png?imageMogr2/auto-orient/strip%7CimageView2/2/w/1240)

---

## 6. 使用 HTTP/2 时，需要注意

1. **HTTP/2 不是默认启用的**  
   服务器和客户端**必须都支持** HTTP/2，否则仍然会回退到 HTTP/1.1。

2. **开启 HTTP/2 不等于网站更快**  
   HTTP/2 主要优化的是**并发**和**传输效率**，但并不意味着所有场景都会有**明显的速度提升**，比如：
    - 如果网站本身请求很少，HTTP/2 提升不明显。
    - 网络条件差，可能仍然会有较大的延迟。

3. **HTTP/2 不能直接替代 WebSocket**
    - HTTP/2 适合短连接和 Web 资源加载。
    - WebSocket 适合长连接（如实时聊天应用）。

4. **使用 HTTPS**
   虽然 HTTP/2 并非强制要求 HTTPS，但在实际应用中，大多数支持 HTTP/2 的服务器都要求使用 HTTPS，因此在发送请求时通常需要使用 `https://` 开头的 URL。

5. **检查 curl 版本**
   旧版本的 curl 可能不支持 HTTP/2，如果遇到问题，先检查并升级 curl 到最新版本。

---

## 7. 总结

HTTP/2 作为 HTTP/1.1 的优化版本，主要提升了**性能和并发能力**，特别是：
- **多路复用**：减少连接数，提升效率
- **头部压缩**：减少重复数据
- **服务器推送**：减少请求延迟
- **二进制格式**：更高效的数据传输

如果你想使用 HTTP/2：
- **确保服务器支持 HTTP/2，并且客户端请求时，需要指定使用 HTTP/2**
- **尽量使用 HTTPS**
- **可以用 `curl` 进行测试**

希望这篇文章能帮你更好地理解 HTTP/2，欢迎留言交流你的问题或经验！🚀