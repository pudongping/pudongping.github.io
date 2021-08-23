---
title: laravel 性能优化
author: Alex
top: false
hide: false
cover: false
toc: true
mathjax: false
categories: PHP
tags:
  - PHP
  - Laravel
abbrlink: 183ed0df
date: 2021-08-23 19:39:01
img:
coverImg:
password:
summary:
---

# laravel 性能优化

## composer 优化

此命令会把 `PSR-0` 和 `PSR-4` 转化为一个映射表，来提高类的加载速度

```
composer dump-autoload --optimize
```

## 安装 Debugbar

```sh
composer require "barryvdh/laravel-debugbar:~3.2" --dev

# 生成配置文件
php artisan vendor:publish --provider="Barryvdh\Debugbar\ServiceProvider"

# 编辑 config/debugbar.php 将 enable 得值设置为
'enable' => env('APP_DEBUG', false),

# 并在 .env 配置文件中开启
APP_DEBUG=true
```

## laravel 配置缓存

> 注意：如果你的配置信息里存在闭包，执行以上命令时将会报错：Your configuration files are not serializable 。解决办法是改写闭包函数为一般的函数，或者改写为类方法。

```sh
# 设置配置缓存，会生成 bootstrap/cache/config.php 文件
php artisan config:cache

# 清空配置缓存，也可以直接删除 bootstrap/cache/config.php 文件来清除
php artisan config:clear
```

测试

```sh
# 手动创建 200 个配置信息
for i in $(seq -f "app%03g.php" 1 200); do cp ./config/app.php ./config/$i; done

# 清除掉 bootstrap/cache/config.php 配置缓存文件
php artisan config:clear

```

## laravel 路由缓存

> 基于闭包的路由无法被缓存。要使用路由缓存。你需要将任何闭包路由转换成控制器路由。

```sh
# 设置路由缓存
php artisan route:cache

# 清除路由缓存
php artisan route:clear
# 也可以直接删除 bootstrap/cache/routes.php 缓存文件
```

路由统计

```sh
php artisan route:list | wc -l | awk '{print $1 - 4}'
```

生成 500 个测试路由
```sh
for i in $(seq -f "%03g" 1 500); do echo "Route::get('test$i/{id}/{query}','Auth\LoginController@testtest$i')->name('test$i')->middleware('auth');" >> routes/web.php; done
```

## 类映射加载优化

在 `laravel 6.x` 中，会生成 `bootstrap/cache/config.php` 和 `bootstrap/cache/packages.php` 和 `bootstrap/cache/routes.php` 和 `bootstrap/cache/services.php` 这四个文件。

services.php 的作用，是把 laravel 在启动过程中需要加载的文件，命名空间和文件路径一个个列举在一个数组中，当 laravel 启动的时候，直接通过这个数组来读取文件。

packages.php 则是将 laravel 扩展包的 provider 和 facade 整理到一个文件里面，在 laravel 启动加载的时候，直接按照这里的数组进行加载。

```sh
php artisan optimize

# 清空类映射
php artisan optimize:clear
```

## 合理划分子视图

> 建议不要子模版嵌套子模板，因为可能会出现 N + 1 的问题

## 会话驱动选择 redis，并且必须选择 phpredis 类库

## 索引优化

### 如何在 laravel 中添加索引？
命令 | 描述
--- | ---
$table->primary('id'); | 添加主键
$table->primary(['id', 'parent_id']); | 添加复合键
$table->unique('email'); | 添加唯一索引
$table->index('state'); | 添加普通索引
$table->spatialIndex('location'); | 添加地理位置信息索引（不支持 SQLite）

添加全文索引时，使用

```php
DB::statement('ALTER TABLE posts ADD FULLTEXT full(name, content)');
```

修改索引名称

```php
$table->renameIndex('from', 'to')
```

### 如何在 laravel 中删除索引？

命令 | 描述
--- | ---
$table->dropPrimary('users_id_primary'); | 从 users 表中删除主键
$table->dropUnique('users_email_unique'); | 从 users 表中删除唯一索引
$table->dropIndex('geo_state_index'); | 从 geo 表中删除基本索引
$table->dropSpatialIndex('geo_location_spatialindex'); | 从 geo 表中删除空间索引（不支持 SQLite）

## 资源压缩（JS/CSS）

- 开发环境时，执行

```sh
npm run dev
```

- 生产环境时，执行 （只有在生产环境中才会对内容进行压缩）

```sh
npm run prod
```

## 压缩 HTML

```
# 安装 laravel-page-speed 扩展包，来去除 html 注释、回车换行符和多余的空格
composer require renatomarinho/laravel-page-speed

# 发布配置文件
php artisan vendor:publish --provider="RenatoMarinho\LaravelPageSpeed\ServiceProvider"

# 在 Kenel.php => $middlewareGroups => web 中添加中间件

\RenatoMarinho\LaravelPageSpeed\Middleware\ElideAttributes::class, // 移除无用的 HTML 属性
\RenatoMarinho\LaravelPageSpeed\Middleware\RemoveComments::class, // 移除注释
\RenatoMarinho\LaravelPageSpeed\Middleware\TrimUrls::class,  // 移除不必要的 URL 前缀
\RenatoMarinho\LaravelPageSpeed\Middleware\CollapseWhitespace::class, // 处理换行符和空格
```

## swoole 加速

```
# 安装 laravles
composer require hhxsv5/laravel-s

# 发布配置
php artisan laravels publish

# 启动 laravels
php bin/laravels start
# 后台守护进程执行
php bin/laravels start -d

```

nginx 配置 laravels

```

upstream swoole {
    # 如果是使用 laradock ，请将 127.0.0.1 更改为 workspace
    server 127.0.0.1:5200 weight=5 max_fails=3 fail_timeout=30s;
    keepalive 16;
}

server {
    listen 80;
    listen 443 ssl http2;
    server_name larablog.test;
    root /var/www/larablog/public;

    index index.html index.htm index.php;

    charset utf-8;

    location / {
        try_files $uri @laravels;
    }

    location = /favicon.ico { access_log off; log_not_found off; }
    location = /robots.txt  { access_log off; log_not_found off; }

    error_log /var/log/nginx/larablog_error.log;
    access_log /var/log/nginx/larablog_access.log;

    sendfile off;

    client_max_body_size 100m;

    location @laravels {
        # proxy_connect_timeout 60s;
        # proxy_send_timeout 60s;
        # proxy_read_timeout 120s;
        proxy_http_version 1.1;
        proxy_set_header Connection "";
        proxy_set_header X-Real-IP $remote_addr;
        proxy_set_header X-Real-PORT $remote_port;
        proxy_set_header X-Forwarded-For $proxy_add_x_forwarded_for;
        proxy_set_header Host $http_host;
        proxy_set_header Scheme $scheme;
        proxy_set_header Server-Protocol $server_protocol;
        proxy_set_header Server-Name $server_name;
        proxy_set_header Server-Addr $server_addr;
        proxy_set_header Server-Port $server_port;
        proxy_pass http://swoole;
    }

    location ~ /\.ht {
        deny all;
    }

    ssl_certificate     /etc/nginx/ssl/larablog.test.crt;
    ssl_certificate_key /etc/nginx/ssl/larablog.test.key;
}

```

如果是使用 laradock 的话，还需要将 `.env` 添加监听地址为 `workspace`

```
LARAVELS_LISTEN_IP=workspace

# 设置后台启动 laravelS 服务，如果需要查看则执行 ps -ef|grep laravels 命令
LARAVELS_DAEMONIZE=true
```
