---
title: 深入理解 Nginx 与 Kong 的配置与实践
author: Alex
top: false
hide: false
cover: false
toc: true
mathjax: false
categories: 微服务
tags:
  - Nginx
  - KONG
  - 微服务
abbrlink: 52ffe35e
date: 2024-07-11 23:16:54
img:
coverImg:
password:
summary:
---

在现代的微服务架构中，服务之间的通信和负载坐标成为了关键环节。这篇文章将深入探讨如何通过 Nginx 配置实现服务的负载均衡，同时如何利用 Kong API 网关管理微服务，包括但不限于服务发现、路由、和负载坐标。我们将使用具体代码示例，确保即便是编程新手也能轻松领悟。

## 一、理解 Nginx 配置

Nginx 是一个高性能的 HTTP 和反向代理服务器，也被广泛用于负载均衡。首先，让我们来看一个基本的 Nginx 配置例子，这将帮助我们理解如何实现负载均衡。

### Nginx 配置示例

```nginx
upstream pay-service {
    server 127.0.0.1:5501 weight=2;
    server 127.0.0.1:5502 weight=4;
    server 127.0.0.1:5503 weight=8;
}

server {
    listen 80;
    server_name pay-service.xxx.com;
    
    location /paymanger {
        proxy_redirect             off;
        proxy_set_header           Host             $http_host;
        proxy_set_header           X-Real-IP        $remote_addr;
        proxy_set_header           X-Forwarded-For  $proxy_add_x_forwarded_for;
        
        proxy_cookie_path / "/; secure; HttpOnly; SameSite=strict";
        
        proxy_pass                 http://pay-service;
    }   
}
```
上面的配置定义了一个名为 `pay-service` 的 upstream，它包括三个后端服务节点，通过不同的权重分配处理请求。这里的权重意味着 `127.0.0.1:5503` 的请求处理能力是 `127.0.0.1:5501` 的四倍。

### 核心概念解释：

- **upstream**: 用来定义一组后端服务节点。
- **server**: 监听客户端请求的配置。
- **location**: 匹配请求URI，并定义处理请求的配置。
- **proxy_pass**: 指定请求转发的后端服务群组。

| Kong 组件 | 说明 |
| --- | --- |
| service | service 对应服务，可以直接指向一个 API 服务节点（host 参数设置为 ip + port），也可以指定一个 upstream 实现负载均衡。简单来说，服务用于映射被转发的后端 API 的节点集合。 **对应的就是以上的 proxy_pass http://pay-service; 这一行配置** |
| route | route 对应路由，它负责匹配实际的请求，映射到 service 中。 **对应的就是以上的  /paymanger 配置** |
| upstream | upstream 对应一组 API 节点，实现负载均衡 |
| target | target 对应一个 API 节点 **对应的就是以上的 127.0.0.1:5501 127.0.0.1:5502 127.0.0.1:5503 这三个节点** |

---

![请求示例图](https://upload-images.jianshu.io/upload_images/14623749-d4fae85689ca7f09.png?imageMogr2/auto-orient/strip%7CimageView2/2/w/1240)


## 二、使用 Kong 进行服务管理

Kong 是一个云原生、快速、可扩展的微服务抽象层（API 网关），用于管理微服务的请求如路由、身份验证、监控等。

### Kong 组件简介

- **service**: 映射后端 API 节点集合。
- **route**: 匹配并映射到 service。
- **upstream**: 一组 API 节点，用于负载均衡。
- **target**: 一个 API 节点。

### 创建 upstream 和 target

#### 创建 upstream

```bash
curl -X POST http://127.0.0.1:8001/upstreams  --data "name=pay-service"
curl http://127.0.0.1:8001/upstreams
```

#### 创建 target

```bash
curl -X POST http://127.0.0.1:8001/upstreams/pay-service/targets  \
--data "target=127.0.0.1:5501" \
--data "weight=2"

curl -X POST http://127.0.0.1:8001/upstreams/pay-service/targets  \
--data "target=127.0.0.1:5502" \
--data "weight=4"

curl -X POST http://127.0.0.1:8001/upstreams/pay-service/targets  \
--data "target=127.0.0.1:5503" \
--d...
```

以上命令将创建三个带有不同权重的 target，对应到之前的 Nginx 配置的三个服务节点。

### 创建 service 和 route

```bash
# 创建 service
curl -X POST http://127.0.0.1:8001/services  \
--data "name=payment-service" \
--data "host=pay-service"

# 创建 service 对应的 route
curl -X POST http://127.0.0.1:8001/services/payment-service/routes  \
--data "name=payment-service-route" \
--data "paths[]=/paymanger"
```

通过上述配置，当用户请求 `/paymanger` 时，Kong 会映射这个请求到 `payment-service`，并通过 upstream `pay-service` 实现负载均衡。

## 结语

Nginx 和 Kong 在现代微服务架构中起到了至关重要的角色。Nginx 擅长处理静态内容、负载均衡和反向代理，而 Kong 提供了一个强大的 API 管理平台，让你可以更容易地管理和监控你的 API。希望这篇文章能够帮助你深入理解它们的工作原理和配置方法。
