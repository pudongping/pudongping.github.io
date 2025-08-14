---
title: 3分钟，手摸手教你用OpenResty搭建高性能隧道代理（附完整配置！）
author: Alex
top: false
hide: false
cover: false
toc: true
mathjax: false
categories: 科普
tags:
  - OpenResty
  - Nginx
abbrlink: c2b7a51a
date: 2025-08-14 12:00:44
img:
coverImg:
password:
summary:
---

经常写爬虫的小伙伴们对代理 IP 应该不会很陌生了吧？

通常，我们为了让爬虫更加稳定，一般我们都会去购买一些代理 IP 用在我们的爬虫服务上。常规的做法，我们一般会去某个代理网站上面购买服务，然后我们会得到一个获取代理 IP 的请求地址，之后我们再写一个请求去获取这些代理 IP。

一般来说，这些代理 IP 的有效期都不会太长，当然和你购买的套餐有一定的关系，常规来说，一般每个代理 IP 的有效期就只有 1-5分钟。我们还需要在爬虫应用程序中去维护这些代理 IP，可能我们的代码就会这样去写

```go
package main

import (
	"crypto/tls"
	"io"
	"net/http"
	"net/url"
	"time"
)

func main() {
   // 通过请求代理IP服务获得一些可用的代理 IP
   // ips := []string{"192.168.0.1:8080", "192.168.0.1:8081", "192.168.0.1:8082"}
   ips := fetchProxyIPs()
   proxyIP := ips[1]

	proxyUrl, err := url.Parse("http://"+proxyIP)
	if err != nil {
		panic(err)
	}
	tr := &http.Transport{
		Proxy:           http.ProxyURL(proxyUrl),
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	client := &http.Client{
		Transport: tr,
		Timeout:   15 * time.Second,
	}
	resp, err := client.Get("https://httpbin.org/ip")
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		panic("Failed to get a valid response")
	}
	content, err := io.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}
	println("Response:", string(content))

}
```

如果我们的爬虫程序只有一个，那么上面的代码完全没有啥问题。但是，如果我们的爬虫程序不止一个呢？是不是 `fetchProxyIPs()` 的代码逻辑就得复制粘贴多次？ 如果哪天我想更换代理服务商岂不是还得一个一个的去改代码？

那么，有没有一种方式，可以在我设置代理 IP 的时候，就设置一个固定的 IP，然后这个固定的 IP 再帮我“自动”去使用代理 IP 呢？

是的，**隧道代理**就是干这事儿的。

> 在软件开发中，没有什么是不能通过加一层中间件来解决问题的，如果有，那么就再加一层……

可能，我们最终需要写的代码，就类似这样：

```go
package main

import (
	"crypto/tls"
	"io"
	"net/http"
	"net/url"
	"time"
)

func main() {
    // 只需要配置隧道代理地址，无需管理代理池
	proxyUrl, err := url.Parse("http://127.0.0.1:9527")
	if err != nil {
		panic(err)
	}
	tr := &http.Transport{
		Proxy:           http.ProxyURL(proxyUrl),
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	client := &http.Client{
		Transport: tr,
		Timeout:   15 * time.Second,
	}
	resp, err := client.Get("https://httpbin.org/ip")
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		panic("Failed to get a valid response")
	}
	content, err := io.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}
	println("Response:", string(content))

}
```

`http://127.0.0.1:9527` 服务就是我们设定的隧道代理，当我们通过 `http://127.0.0.1:9527` 去设置代理时，`http://127.0.0.1:9527` 会自动帮我们切换代理 IP。

现在有很多代理 IP 服务商都有提供隧道代理服务的，但是，价格一般都不会太便宜。感兴趣的小伙伴们可以去了解了解。

其实，自己动手搭建一个隧道代理服务也不会太复杂，用 go 写一个代理转发程序也是可以的，但是，在这个应用场景下，还有更好的选择：**OpenResty**。

OpenResty 其实是 Nginx + Lua JIT。Nginx 本身就擅长处理 TCP 连接，性能高，稳定成熟。

有小伙伴这时候就说了，不太会 Lua 脚本怎么办？

没关系，这里我将整个配置都贴出来，以供各位参考：

```nginx
worker_processes  16;

error_log  /usr/local/openresty/nginx/logs/error.log debug;

events {
    worker_connections  1024;
}


stream {

    # 自定义 TCP 日志格式定义
    # 包含连接的 IP、时间、协议、状态、流量、会话时长、上游地址及流量等
    log_format tcp_proxy '$remote_addr [$time_local] '
                         '$protocol $status $bytes_sent $bytes_received '
                         '$session_time "$upstream_addr" '
                         '"$upstream_bytes_sent" "$upstream_bytes_received" "$upstream_connect_time"';
    # 启用日志记录到指定文件，并使用自定义格式
    access_log /usr/local/openresty/nginx/logs/tcp-access.log tcp_proxy;
    open_log_file_cache off;

    # TCP 代理配置
    # upstream 块中定义一个占位 server
    # 注意：0.0.0.0:1101 实际不会使用，真正地址会被 balancer_by_lua_block 动态覆盖
    upstream real_server {
        server 0.0.0.0:1101;

        # 使用 balancer_by_lua_block 动态设置后端目标主机和端口
        balancer_by_lua_block {
            -- 检查 preread 阶段是否已经设置了 proxy_host 和 proxy_port
            -- 从 ngx.ctx 中获取代理服务器地址
            if not ngx.ctx.proxy_host or not ngx.ctx.proxy_port then
                ngx.log(ngx.ERR, "====>proxy_host or proxy_port is not set in ngx.ctx<====")
                return
            end

            -- 初始化 balancer
            local balancer = require "ngx.balancer"
            local host = ""
            local port = 0

            -- 从上下文中提取目标 IP 和端口
            host = ngx.ctx.proxy_host
            port = ngx.ctx.proxy_port
            -- 设置代理服务器地址
            local ok, err = balancer.set_current_peer(host, port)
            if not ok then
                ngx.log(ngx.ERR, "====>failed to set current peer: " .. tostring(err) .. "<====")
                return
            end
        }
    }

    # 定义 TCP server 模块（stream）监听端口和代理逻辑
    server {
        # preread_by_lua_block 在客户端连接建立时就会触发，用于预处理逻辑
        preread_by_lua_block {
            -- https://github.com/openresty/lua-resty-redis
            local redis = require "resty.redis"
            local redis_instance = redis:new()

            -- 设置 Redis 操作超时时间（毫秒）
            redis_instance:set_timeout(5000)

            -- 一些 redis 连接配置
            local rdb_host = "192.168.1.208"
            local rdb_port = 6379
            local rdb_pwd = ""
            local rdb_db = 1
            -- 存放代理服务器地址的 zset 表名称
            local zset_table_name = "tunnel_proxy_pool"

            -- 连接到 Redis
            local ok, err = redis_instance:connect(rdb_host, rdb_port)
            if not ok then
                ngx.log(ngx.ERR, "====>failed to connect to Redis: [" .. tostring(ok) .. "] err msg ==> " .. tostring(err) .. "<====")
                return
            end

            -- 选择数据库
            local ok, err = redis_instance:select(rdb_db)
            if not ok then
                ngx.log(ngx.ERR, "====>failed to select Redis DB: [" .. tostring(ok) .. "] err msg ==> " .. tostring(err) .. "<====")
                return
            end

            -- 如果设置了密码，则进行认证
            if rdb_pwd and rdb_pwd ~= "" then
                local ok, err = redis_instance:auth(rdb_pwd)
                if not ok then
                    ngx.log(ngx.ERR, "====>failed to auth Redis: [" .. tostring(ok) .. "] err msg ==> " .. tostring(err) .. "<====")
                    return
                end
            end

            -- 先检查 zset 表是否存在或者是否有数据
            local hosts_count, err = redis_instance:zcard(zset_table_name)
            if not hosts_count or hosts_count <= 0 then
                ngx.log(ngx.ERR, "====>no available proxy servers in Redis zset table: " .. tostring(zset_table_name) .. " ==> " .. tostring(err) .. "<====")
                return
            end
            -- 获取分数最低的前 1 个代理服务器地址
            local res, err = redis_instance:zrange(zset_table_name, 0, 0, "WITHSCORES")
            if not res or #res == 0 then
                ngx.log(ngx.ERR, "====>failed to get proxy server from Redis zset table: " .. tostring(zset_table_name) .. "<====")
                return
            end
            -- 解析结果（假设之前存入 zset 的元素类似 127.0.0.1:8080、127.0.0.1:8181 分数为使用次数）
            local proxy_ip, proxy_port = res[1]:match("([^:]+):(%d+)")
            if not proxy_ip or not proxy_port then
                ngx.log(ngx.ERR, "====>failed to parse proxy server address ==> " .. tostring(res[1]) .. "<====")
                return
            end
            -- 获取了当前代理服务器地址后，给其分数加 1，表示当前已经使用过一次
            local ok, err = redis_instance:zincrby(zset_table_name, 1, res[1])
            if not ok then
                ngx.log(ngx.ERR, "====>failed to increment proxy server score in Redis zset table: " .. tostring(zset_table_name) .. " ==> " .. tostring(err) .. "<====")
                return
            end

            -- 将获取到的代理服务器地址存入 ngx.ctx 中，供 balancer_by_lua_block 使用
            ngx.ctx.proxy_host = proxy_ip
            ngx.ctx.proxy_port = tonumber(proxy_port)
            ngx.log(ngx.INFO, "====>using proxy server ==> " .. tostring(proxy_ip) .. ":" .. tostring(proxy_port) .. "<====")

            -- 释放 Redis 连接，否则连接池将保留不完整的连接状态
            ok, err = redis_instance:set_keepalive(10000, 100)
            if not ok then
                ngx.log(ngx.ERR, "====>failed to set Redis keepalive: " .. tostring(err) .. "<====")
            end
        }

        # 对外暴露的监听端口
        listen 0.0.0.0:9527;
        # 设置代理的目标 upstream 名称
        proxy_pass real_server;
        proxy_connect_timeout 5s;
        proxy_timeout 15s;
    }

}
```

以上，其实我们就是借用 OpenResty 做了一层代理转发，你可以结合流程图来看看

![流程图](https://upload-images.jianshu.io/upload_images/14623749-2713e0b65ed0a415.png?imageMogr2/auto-orient/strip%7CimageView2/2/w/1240)

那么，如何部署 OpenResty 呢？

可以直接使用下面的 `docker-compose.yaml` 文件：

```yaml
services:
  openresty:
    container_name: openresty_server
    image: openresty/openresty:1.25.3.2-5-centos7
    ports:
      - "9527:9527"
    volumes:
      - ./conf/tunnel_proxy_redis.conf:/usr/local/openresty/nginx/conf/nginx.conf:ro
      - ./logs:/usr/local/openresty/nginx/logs
```

文件写好之后，直接在和 `docker-compose.yaml` 文件同级目录下执行 `docker-compose up` 即可启动 OpenResty 服务。

另外，还忘记说了一点：你需要自己写一个脚本，定时将可用的代理 IP 同步到 redis 中，上面的 Lua 脚本只是会从 redis 中取出可用的代理 IP 进行转发。

自动脚本干的活儿类似写入这样的数据

```bash
zadd tunnel_proxy_pool 0 127.0.0.1:9001 0 4127.0.0.1:9002 0 127.0.0.1:9003
```

大家感兴趣的，可以通过访问 `https://github.com/pudongping/tunnel-proxy` 获得源码。