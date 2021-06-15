---
title: Nginx/Tengine 服务器安装 SSL 证书
author: Alex
top: false
hide: false
cover: false
toc: true
mathjax: false
categories: 服务器
tags:
  - 免费 HTTPS 证书
  - CentOS
  - Linux
abbrlink: 2078b83a
date: 2021-06-15 09:47:36
img:
coverImg:
password:
summary:
---

# Nginx/Tengine 服务器安装 SSL 证书

## 背景

> 使用阿里云下载免费版 DigiCert Inc 证书，并基于 Nginx web 服务器安装 SSL 证书

## 安装步骤

在证书控制台下载 Nginx 版本证书。下载到本地的压缩文件包解压后包含：

- `.crt` 文件：是证书文件，crt 是 pem 文件的扩展名。
- `.key` 文件：证书的私钥文件（申请证书时如果没有选择``自动创建 CSR``，则没有该文件）。

`友情提示：`  `.pem` 扩展名的证书文件采用 Base64-encoded 的 PEM 格式`文本文件`，可根据需要修改扩展名。

以 Nginx 标准配置为例，假如证书文件名是 a.pem，私钥文件是 a.key。

在 Nginx 的安装目录下创建 cert 目录，并且将下载的全部文件拷贝到 cert 目录中。如果申请证书时是自己创建的 CSR 文件，请将对应的私钥文件放到 cert 目录下并且命名为 a.key；

打开 Nginx 安装目录下 conf 目录中的 nginx.conf 文件，找到：

```nginx

# HTTPS server
# #server {
# listen 443;
# server_name localhost;
# ssl on;
# ssl_certificate cert.pem;
# ssl_certificate_key cert.key;
# ssl_session_timeout 5m;
# ssl_protocols SSLv2 SSLv3 TLSv1;
# ssl_ciphers ALL:!ADH:!EXPORT56:RC4+RSA:+HIGH:+MEDIUM:+LOW:+SSLv2:+EXP;
# ssl_prefer_server_ciphers on;
# location / {
#
#
#}
#}

```

将其修改为 (以下属性中 ssl 开头的属性与证书配置有直接关系，其它属性请结合自己的实际情况复制或调整) :

```nginx

server {
 listen 443; # 在使用宝塔面板搭建的 LNMP 环境中，推荐 `listen 443 ssl http2;` 这么写
 server_name localhost;
 ssl on;
 root html;
 index index.html index.htm;
 ssl_certificate   cert/a.pem; # 这里推荐写绝对路径
 ssl_certificate_key  cert/a.key; # 这里推荐写绝对路径
 ssl_session_timeout 5m;
 ssl_ciphers ECDHE-RSA-AES128-GCM-SHA256:ECDHE:ECDH:AES:HIGH:!NULL:!aNULL:!MD5:!ADH:!RC4;
 ssl_protocols TLSv1 TLSv1.1 TLSv1.2;
 ssl_prefer_server_ciphers on;
 location / {
     root html;
     index index.html index.htm;
 }
}

```

## 自己使用宝塔面板 6.9.8 免费版，搭建 LNMP 环境时的 nginx 配置文件

```nginx

server
{
    listen 80; 
    
    # 宝塔中建议这么写监听 443 端口
	listen 443 ssl http2;  
	
    server_name www.drling.xin;
    index index.php index.html index.htm default.php default.htm default.html;
    root /www/wwwroot/larablog/public;
    
    
    # 强制开启 https ，当访问 http 时，直接强制跳转到 https
    #HTTP_TO_HTTPS_START
    if ($server_port !~ 443){
        rewrite ^(/.*)$ https://$host$1 permanent;
    }
    #HTTP_TO_HTTPS_END
    limit_conn perserver 3000;
    limit_conn perip 25;
    limit_rate 512k;    
    
    # 这里直接采用宝塔面板，将 ssl 信息通过 [其他证书] 栏，导入进去，因此 nginx 配置文件由宝塔面板自行写入。
    # 若需要自己写入配置文件，以下参数可以这样写，如下：
    # ssl_certificate    /www/server/panel/vhost/aliyun-ssl/1939426_drling.xin.pem;
    # ssl_certificate_key    /www/server/panel/vhost/aliyun-ssl/1939426_drling.xin.key;    
    
    #SSL-START SSL相关配置，请勿删除或修改下一行带注释的404规则
    #error_page 404/404.html;    
    ssl_certificate    /www/server/panel/vhost/cert/www.drling.xin/fullchain.pem;
    ssl_certificate_key    /www/server/panel/vhost/cert/www.drling.xin/privkey.pem;
    ssl_protocols TLSv1 TLSv1.1 TLSv1.2 TLSv1.3;
    ssl_ciphers ECDHE-RSA-AES128-GCM-SHA256:HIGH:!aNULL:!MD5:!RC4:!DHE;
    ssl_prefer_server_ciphers on;
    ssl_session_cache shared:SSL:10m;
    ssl_session_timeout 10m;
    error_page 497  https://$host$request_uri;
    #SSL-END
    
    
    #ERROR-PAGE-START  错误页配置，可以注释、删除或修改
    #error_page 404 /404.html;
    #error_page 502 /502.html;
    #ERROR-PAGE-END
    
    #PHP-INFO-START  PHP引用配置，可以注释或修改
    include enable-php-73.conf;
    #PHP-INFO-END
    
    #REWRITE-START URL重写规则引用,修改后将导致面板设置的伪静态规则失效
    include /www/server/panel/vhost/rewrite/www.drling.xin.conf;
    #REWRITE-END
    
    #禁止访问的文件或目录
    location ~ ^/(\.user.ini|\.htaccess|\.git|\.svn|\.project|LICENSE|README.md)
    {
        return 404;
    }
    
    #一键申请SSL证书验证目录相关设置
    location ~ \.well-known{
        allow all;
    }
    
    location ~ .*\.(gif|jpg|jpeg|png|bmp|swf)$
    {
        expires      30d;
        error_log off;
        access_log /dev/null;
    }
    
    location ~ .*\.(js|css)?$
    {
        expires      12h;
        error_log off;
        access_log /dev/null; 
    }
    access_log  /www/wwwlogs/www.drling.xin.log;
    error_log  /www/wwwlogs/www.drling.xin.error.log;
}

```

## 结束语
安装完毕之后你可以使用[SSL Server Test -- 安全测试工具](https://www.ssllabs.com/ssltest/index.html) 去测试下你的 HTTPS 是否够安全

## 参考文档

> [阿里云官方介绍 Nginx 服务器安装 SSL 证书](https://help.aliyun.com/knowledge_detail/95491.html)
