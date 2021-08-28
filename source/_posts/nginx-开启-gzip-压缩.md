---
title: nginx 开启 gzip 压缩
author: Alex
top: false
hide: false
cover: false
toc: true
mathjax: false
categories: Nginx
tags:
  - Nginx
  - gzip
abbrlink: 63bed376
date: 2021-08-28 22:11:09
img:
coverImg:
password:
summary:
---

# nginx 开启 gzip 压缩

> 在服务器 Nginx 开启 gzip 压缩是优化网站性能的方法之一，可以有效减少服务器带宽的消耗，缺点是会增大 CPU 的占用率，但是很多时候 CPU 往往是空闲最多的。

在 nginx 中开启 gzip 压缩，需要编辑 nginx.conf 文件，添加如下

```nginx

gzip on;
gzip_min_length 1k;
gzip_buffers 16 64k;
gzip_http_version 1.1;
gzip_comp_level 9;
gzip_types text/plain text/css text/javascript application/json application/javascript application/x-javascript application/xml application/x-httpd-php image/jpeg image/gif image/png font/ttf font/otf image/svg+xml;
gzip_vary on;   

```

## gzip 参数介绍

### gizp on|off;
开启或者关闭 gzip 模块

### gzip_min_length 1k;
设置允许压缩的页面最小字节数，页面字节数从 header 头中的 Content-Length 中进行获取。默认值是 0，不管页面多大都压缩。建议设置成大于 1k 的字节数，小于 1k 可能会越压越大。 即: gzip_min_length 1024

### gzip_proxied expired no-cache no-store private auth;
Nginx 作为反向代理的时候启用，开启或者关闭后端服务器返回的结果，匹配的前提是后端服务器必须要返回包含”Via” 的 header 头。

### gzip_types text/plain application/xml;
匹配 MIME 类型进行压缩，（无论是否指定）”text/html” 类型总是会被压缩的。

## 如何判断是否已经开启了 gzip 压缩？

打开浏览器，按住 f12 查看 Content-Encoding 字段如果是 gzip，表示该网页的 body 数据是经过 gzip 压缩的。
