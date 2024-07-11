---
title: Docker 安装 KONG 带你玩转 API 网关
author: Alex
top: false
hide: false
cover: false
toc: true
mathjax: false
categories: 微服务
tags:
  - Docker
  - KONG
  - 微服务
abbrlink: b8eafe4f
date: 2024-07-11 23:16:39
img:
coverImg:
password:
summary:
---

在当今的软件开发中，API 网关已成为微服务架构中不可或缺的一环。它不仅简化了服务间的通信、提供了统一的入口，还能在安全、监控、限流等方面发挥巨大作用。

今天，我们就来聊聊如何通过 Docker 快速部署 KONG —— 一个流行的开源 API 网关。

## KONG 简介

KONG 是基于 [Nginx](nginx.org) 和 [OpenResty](https://openresty.org/cn/) （Nginx + Lua）的云原生、高性能、可扩展的微服务 API 网关。它以插件形式提供丰富的功能，包括但不限于：

- 身份认证（如 JWT、basic-auth）
- 安全性（如 IP 黑白名单）
- 监控
- 流量控制

## KONG 与 Nginx 和 OpenResty 的关系

为了更好地理解 KONG，我们首先要明白它与 Nginx、OpenResty 的关系：

- **Nginx**：是一个高性能的 HTTP 服务器和反向代理，以及一个 IMAP/POP3 代理服务器。
- **OpenResty**：是在 Nginx 上集成了 Lua-nginx-module，允许使用 Lua 脚本语言扩展 Nginx 的能力。
- **KONG**：则是在 OpenResty 基础上，加入了自定义框架和一系列企业级功能实现的 API 网关。

## Docker 安装 KONG

安装 KONG 的一种便捷方式是使用 Docker。下面，我们将详细介绍如何通过 Docker 进行安装。

### 步骤 1：创建容器网络

首先创建一个 Docker 网络，以便容器间可以互相通信。

```bash
docker network create kong-net
```

### 步骤 2：搭建数据库

KONG 支持 Cassandra 和 PostgreSQL 数据库。这里，我们以 PostgreSQL 为例。

```bash
# 创建一个数据卷
docker volume create kong-volume
# 查看所有的数据卷
docker volume ls

docker run -d --name kong-database \
--network=kong-net \
-p 5432:5432 \
-e "POSTGRES_USER=konguser" \
-e "POSTGRES_DB=kong" \
-e "POSTGRES_PASSWORD=kongpwd" \
# 挂载卷以便持久化数据到宿主机
-v kong-volume:/var/lib/postgresql/data \
postgres:9.6
```

### 步骤 3：数据库初始化

使用 `docker run --rm` 来初始化数据库，该命令执行后会退出容器而保留内部的数据卷。并且，要注意：**一定要跟你声明的网络、数据库类型、host 名称一致。**

```bash
docker run --rm --network=kong-net \
# 如果使用的是 Cassandra 时，则需要设定为 cassandra 我这里使用的是 PostgreSQL 因此则设定为 postgres
-e "KONG_DATABASE=postgres" \
-e "KONG_PG_HOST=kong-database" \
-e "KONG_PG_USER=konguser" \
-e "KONG_PG_PASSWORD=kongpwd" \
# 企业版才会用到这个配置，非企业版设定此参数也无所谓
-e "KONG_PASSWORD=test" \
kong/kong-gateway:3.4.1.1 kong migrations bootstrap
```

### 步骤 4：启动 KONG

一切准备就绪后，我们可以启动 KONG 容器了。

```bash
docker run -d --name kong-gateway \
--network=kong-net \
-e "KONG_DATABASE=postgres" \
-e "KONG_PG_HOST=kong-database" \
-e "KONG_PG_USER=konguser" \
-e "KONG_PG_PASSWORD=kongpwd" \
-e "KONG_PROXY_ACCESS_LOG=/dev/stdout" \
-e "KONG_ADMIN_ACCESS_LOG=/dev/stdout" \
-e "KONG_PROXY_ERROR_LOG=/dev/stderr" \
-e "KONG_ADMIN_ERROR_LOG=/dev/stderr" \
-e "KONG_ADMIN_LISTEN=0.0.0.0:8001" \
-e "KONG_ADMIN_GUI_URL=http://localhost:8002" \
# 仅仅企业版需要设定证书
-e KONG_LICENSE_DATA \
-p 8000:8000 \
-p 8443:8443 \
-p 8001:8001 \
-p 8444:8444 \
-p 8002:8002 \
-p 8445:8445 \
-p 8003:8003 \
-p 8004:8004 \
kong/kong-gateway:3.4.1.1
```

### 步骤 5：检查 KONG 是否运行正常

最后，我们来检查 KONG 是否成功运行。

```bash
curl -i -X GET --url http://{YOUR_SERVER_IP}:8001/services

# 你也可以通过浏览器访问以下地址查看 KONG 管理界面
http://{YOUR_SERVER_IP}:8002
```

至此，你已经成功地在 Docker 上部署了 KONG。

## 安装 Konga

Konga 是一个开源的 KONG 管理界面，可以帮助我们更方便地管理和监控 KONG。

```bash
# 拉取最新版本的 Konga
docker pull pantsel/konga
```

通过以下命令预装 Konga 所需的数据库：

```bash
# 这里的 172.18.0.2 要是 kong-database 容器的 IP 地址
# konguser 和 kongpwd 是前面安装 PostgreSQL 是用到的账号和密码，konga 数据库专门为 Konga 设定（执行这条命令时，如果 konga 数据库不存在，则会自动创建）
docker run --rm pantsel/konga \
--network=kong-net \
-c prepare \
-a postgres \
-u postgresql://konguser:kongpwd@172.18.0.2:5432/konga
```

安装 Konga：

```bash
docker run -d --name konga \
-p 1337:1337 \
--network kong-net \
-e "NODE_ENV=production" \
-e "DB_ADAPTER=postgres" \
-e "DB_HOST=kong-database" \
-e "DB_PORT=5432" \
-e "DB_USER=konguser" \
-e "DB_PASSWORD=kongpwd" \
-e "DB_DATABASE=konga" \
pantsel/konga
```

可以通过浏览器访问 `http://{YOUR_SERVER_IP}:1337` 来检测安装是否成功。第一次访问时，需要注册管理员账号，通过注册之后，然后在 CONNECTIONS 中添加 Kong 服务的管理路径 `http://kong-gateway:8001` （因为这几个容器都连接了 kong-net 网络，因此这里可以通过容器名称作为 IP 地址，亦即这里的 `kong-gateway`） 即可管理 Kong。

## 如果需要清除所有容器时

```bash
docker kill kong-gateway
docker kill kong-database
docker kill konga
docker container rm kong-gateway
docker container rm kong-database
docker container rm konga
docker network rm kong-net
docker volume rm kong-volume
```

通过本文的步骤，你应该能够成功地在 Docker 中安装和运行 KONG API 网关以及 Konga 管理界面。KONG 提供了强大的 API 管理功能，而 Konga 则提供了一个用户友好的界面来管理 KONG 的各个方面。

希望这篇文章能帮助你快速入门 KONG，赶快动手试试吧！