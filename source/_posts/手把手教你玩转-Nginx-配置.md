---
title: 手把手教你玩转 Nginx 配置
author: Alex
top: false
hide: false
cover: false
toc: true
mathjax: false
categories: Nginx
tags:
  - Nginx
abbrlink: 40e3643a
date: 2024-07-03 11:01:00
img:
coverImg:
password:
summary:
---

在现代的互联网应用中，Nginx 已经成为了不可或缺的组成部分。无论是作为静态资源服务器、反向代理服务器、还是负载均衡器，Nginx 的高性能和灵活配置都让它备受青睐。

本文将以简单、易懂的语言和实例，详细介绍几种常见的 Nginx 配置场景，旨在帮助初学者和有一定经验的开发者能更好地理解和使用 Nginx。

## 一、简单配置

让我们从最基本的 Nginx 配置讲起。下面是一个非常基础的静态网站配置示例：

```nginx
server {
    listen          80;
    server_name     alex.com; # 填写你的域名
    location / { 
        root /var/www/wwwroot/myblog; # 静态文件存放目录
        index  index.html index.htm;  # 默认页面
    }
}
```

### 关键点解析：

- `listen 80;`：此行配置 Nginx 监听 80 端口，即 HTTP 标准端口。
- `server_name`：指定当前服务器块处理的域名。
- `location /`：处理根 URL 的请求。在这个位置块里，你定义了请求 `/` 时的行为。
- `root`：指定静态文件的存放目录。
- `index`：当请求目录时，默认返回的文件名。

## 二、配置 SSL - 从 80 端口转发到 443

随着网络安全的日益重要，为网站配置 SSL，即在 HTTP 上实施 TLS/SSL 来加密客户端和服务器之间的通信，已成为一项标准实践。下面的配置展示了如何将 http 流量（80端口）重定向到 https（443端口）：

```nginx
# 80 端口配置，用于重定向到 https
server {
    listen          80;
    server_name     alex.com;
    rewrite ^ https://$http_host$request_uri? permanent;
}
# https 配置
server {
    listen 443 ssl;
    server_name alex.com;     # 绑定证书的域名
    ssl_certificate /etc/nginx/cert/alex_blog.crt;      # 证书路径
    ssl_certificate_key /etc/nginx/cert/alex_blog.key;  # 私钥路径
    ssl_session_timeout 5m;
    ssl_protocols TLSv1 TLSv1.1 TLSv1.2;  # 使用的协议
    ssl_ciphers ECDHE-RSA-AES128-GCM-SHA256:HIGH:!aNULL:!MD5:!RC4:!DHE;
    ssl_prefer_server_ciphers on;
    location / {
        root /var/www/wwwroot/myblog;    # 站点目录
        index  index.html index.htm;
    }
}
```

### 关键点解析：

- 第一个 `server` 块相当于一个跳板，接受 80 端口的请求并通过 301 永久重定向到 https。
- 第二个 `server` 块则是真正处理 https 请求的。
- `ssl_certificate` 和 `ssl_certificate_key` 分别指向了 SSL 证书和私钥的路径，这对于启用 https 是必须的。

## 三、反向代理配置

反向代理是 Nginx 的另一个常用功能，它能让你将客户端的请求转发到其他服务器，并将其响应返回给客户端。这样做的好处包括隐藏服务器真实 IP、负载均衡、缓存静态内容等。以下示例配置描述了如何将对 `api.alex.com` 的请求转发到另一个服务器：

```nginx
server {
    listen 80;
    # 访问 api.alex.com 实际访问到 https://www.alex.com:5200 
    server_name  api.alex.com;
    location / {
        proxy_pass https://www.alex.com:5200; #反向代理的地址
        proxy_http_version 1.1; #配置参数（重要）
        proxy_set_header Upgrade $http_upgrade;
        proxy_set_header Connection "upgrade";
    }
}
```

### 关键点解析：

- `proxy_pass` 指明了被代理服务器的地址。
- `proxy_set_header` 用于设置 HTTP 头信息，以确保正常的 WebSocket 通信。

## 四、配置图片服务器

当你需要专门处理图片或其他静态资源时，Nginx 可以配置成一个高效的图片服务器。以下是一个示例配置，用于处理特定路径下的图片文件请求：

```nginx
# 正则表达式，访问 /home/alex/images/ 路径下的文件
location ~ (images/).+(gif|jpg|jpeg|png)$ { 
                expires 24h;
                root /home/alex/;#指定图片存放路径
                access_log /home/alex/log/images.log;#日志存放路径
                proxy_store on;
                proxy_store_access user:rw group:rw all:rw;
                proxy_temp_path     /home/alex/;#图片访问路径
                proxy_redirect     off;
                proxy_set_header    Host 127.0.0.1;
                client_max_body_size  10m;
                client_body_buffer_size 1280k;
                proxy_connect_timeout  900;
                proxy_send_timeout   900;
                proxy_read_timeout   900;
                proxy_buffer_size    40k;
                proxy_buffers      40 320k;
                proxy_busy_buffers_size 640k;
                proxy_temp_file_write_size 640k;
                if ( !-e $request_filename)
                {
                        proxy_pass http://127.0.0.1;#默认80端口 
                }
        }
```

### 关键点解析：

- 这个位置匹配了以 `images/` 开头，并以 `gif`、`jpg`、`jpeg`、`png` 结尾的请求路径，然后对这些资源进行了缓存和日志记录的配置。

## 五、Vue 应用的 Nginx 配置

将前端项目部署到服务器上时，你通常需要配置一个 Web 服务器来托管它们。Vue 应用就是这样一个案例。下面的 Nginx 配置适用于 Vue 应用的部署：

```nginx
server {
    listen 80;
    server_name h5.alex.com;

    error_log /var/log/nginx/frontend-h5.err;
    access_log  /var/log/nginx/frontend-h5.log;

    location / {
        root   /opt/h5;
#        rewrite ^/article/([0-9]+)$ /index.html?id=$1 last;
             try_files $uri $uri/ /index.html;
    }

        location /dapi/ {
           proxy_pass http://172.19.0.23:8502/; 
           proxy_set_header Host $host;
           proxy_set_header X-Real-IP $http_x_real_ip;
           proxy_set_header X-Forwarded-For $proxy_add_x_forwarded_for;
        }


    #
    error_page   500 502 503 504  /50x.html;

    location = /50x.html {
        root   /usr/share/nginx/html;
    }
}
```

### 关键点解析：

- `try_files` 命令用于尝试按顺序访问指定的文件或目录，如果都没有找到，最后会重定向到 `/index.html` 文件。这对于单页面应用（SPA）非常重要，因为它们依赖于前端路由。

## 六、Go Gin 应用的配置与负载均衡

Go-Gin 是一个高性能的 Web 框架，适用于构建高效的 Web 应用。部署至生产环境时，你可能需要 Nginx 作为反向代理服务器，并实现负载均衡以提高应用的可用性和响应速度。以下配置展示了如何设置反向代理和负载均衡：

```nginx

worker_processes  1;

events {
    worker_connections  1024;
}


http {
    include       mime.types;
    default_type  application/octet-stream;

    sendfile        on;
    keepalive_timeout  65;

    server {
        listen       8081;
        server_name  api.blog.com;

        # 将所有路径转发到 http://127.0.0.1:8000/ 下
        location / {
            proxy_pass http://127.0.0.1:8000/;
        }
    }
}

```

配置负载均衡

> 使用 gin 启动两个服务，分别监听 8001 端口和 8002 端口

```nginx

worker_processes  1;

events {
    worker_connections  1024;
}


http {
    include       mime.types;
    default_type  application/octet-stream;

    sendfile        on;
    keepalive_timeout  65;

    # 设置其对应的 2 个后端服务
    upstream api.blog.com {
        server 127.0.0.1:8001;
        server 127.0.0.1:8002;
    }

    server {
        listen       8081;
        server_name  api.blog.com;

        location / {
            # 格式为 http:// + upstream 的节点名称
            proxy_pass http://api.blog.com/;
        }
    }
}

```

### 关键点解析：

- `upstream` 指令定义了一个服务器组，可以包含一个或多个服务器。在这个例子中，两个 Gin 应用实例分别运行在 8001 和 8002 端口。
- 当请求到来时，Nginx 会根据配置的策略（默认为轮询）将请求分发到不同的服务器。

## 七、按目录划分项目

- www.blog.com 访问静态文档项目
- www.blog.com/frontend 访问门户页面
- www.blog.com/backend 访问后台页面

```nginx

server {
    listen       80;
    server_name  www.blog.com;

    location / {
        root   /usr/share/nginx/html/www;
        index  index.html index.htm;
    }

    location /frontend {
        alias   /usr/share/nginx/html/frontend;
        index  index.html index.htm;
    }

    location /backend {
        alias   /usr/share/nginx/html/backend;
        index  index.html index.htm;
    }

    error_page   500 502 503 504  /50x.html;
    location = /50x.html {
        root   /usr/share/nginx/html;
    }
}

```

## 八、一些判断

```nginx

location / {
    # 没有后缀的请求才会转发 
    if (!-e $request_filename){
        proxy_pass http://localhost:8080;
        break;
    }
}

```

通过这篇文章，我们简要介绍了 Nginx 的几种常见配置场景，并提供了详细的代码示例和解释。希望这能帮助你更好地理解和运用 Nginx。无论你是完全的新手还是有一定经验的开发者，掌握这样一个强大的工具都将极大地提升你的开发和部署效率。
