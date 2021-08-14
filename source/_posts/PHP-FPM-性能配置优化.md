---
title: PHP-FPM 性能配置优化
author: Alex
top: false
hide: false
cover: false
toc: true
mathjax: false
categories: PHP
tags:
  - PHP
  - PHP-FPM
abbrlink: 6643f286
date: 2021-08-14 19:05:08
img:
coverImg:
password:
summary:
---

# PHP-FPM 性能配置优化

> 4 核 8 G 服务器大约可以开启 500 个 PHP-FPM，极限吞吐量在 580 qps （Query Per Second 每秒查询数）左右。

## Nginx + php-fpm 是怎么工作的？
php-fpm 全称是 **PHP FastCGI Process Manager** 的简称，从名字可得知，是一个 FastCGI 的管理器。

### 什么是 FastCGI？

FastCGI 是 **Fast Common Gateway Interface** 的简称，是一种交互程序（此处是 PHP）与 Web 服务器之间的 通信协议。FastCGI 是早期通用网关接口（CGI）的增强版本。

注意 FastCGI 和 CGI 都是一种 **通信协议**，独立于任何语言。Web 服务器无须对语言有任何了解。除 PHP 有 php-fpm 外，像 Python, Ruby, Perl, Tcl, C/C++, 和 Visual Basic 都有其各自的 CGI 和 FastCGI 实现。

### CGI 和 FastCGI 的区别？

CGI 程序运行在独立的进程中，并对每个 Web 请求创建一个进程，这种方法非常容易实现，但效率很差，难以扩展。 **面对大量请求，进程的大量创建和消亡使操作系统性能大大下降。** 此外，由于地址空间无法共享，也限制了资源重用。

FastCGI 致力于减少网页服务器与 CGI 程序之间交互的开销，从而使服务器可以同时处理更多的网页请求。与为每个请求创建一个新的进程不同，FastCGI 使用持续的进程来处理一连串的请求。这些进程由 FastCGI 服务器管理（FPM），而不是 Web 服务器。当进来一个请求时，Web 服务器把环境变量和这个页面请求通过一个 Socket 或者 TCP Connection 传递给 FastCGI 进程：

![nginx + php-fpm 工作模式](https://cdn.learnku.com/uploads/images/201909/23/1/3PvNql1Ri5.png!large)


## php-fpm 进程数调优

fpm 服务启动初始化时，会根据配置信息里设置的运行模式，来选择是否创建、以及创建多少 CGI 进程，这些进程随时待命，等待处理从 Web 服务器传送过来的请求：

![nginx 配合 php-cgi 的运行](https://cdn.learnku.com/uploads/images/201909/23/1/6J4kGKBzVL.png!large)


PHP 7.2 FPM 进程池的配置信息位于：

```
/etc/php/7.2/fpm/pool.d/www.conf
```

搜寻下 pm 运行模式的配置，默认是 dynamic ：

```
pm = dynamic
```

fpm 的运行模式有三种：

- ondemand 按需创建
- dynamic 动态创建
- static 固定数量

**ondemand**

ondemand 初始化时不会创建待命的进程。并且会在空闲时将进程销毁，请求进来时再开启。一般是在共享的 VPS 上使用。是一种比较 节省内存 的 FPM 运行方式，不过因为其频繁创建和销毁进程，性能表现不佳。

相关参数：

```

; 默认是 10 秒，超过 10 秒即销毁
pm.process_idle_timeout = 10s;

; 最大并存进程数，超过此值将不再创建
pm.max_children = 50

; 每个进程最多处理多少个请求，超过此值将自动销毁
pm.max_requests = 1000

```


**dynamic**

动态创建，这个是默认选项，也是比较灵活的选项。兼顾稳定和快速响应。同时有四个参数会影响此配置：

```

; FPM 启动时创建的进程数
pm.start_servers = 10

; 最大并存进程数，超过此值将不再创建
pm.max_children = 50

; 空闲进程数最小值，如果空闲进程小于此值，则创建新的子进程
pm.min_spare_servers = 10

; 空闲进程数最大值，如果空闲进程大于此值，则进行清理
pm.max_spare_servers = 40

; 每个进程最多处理多少个请求，超过此值将自动销毁
pm.max_requests = 1000

```


上面的注释已经很释义，空闲进程的概念需要讲下。按照上面的设置，fpm 启动时会有 10 个进程启动，此时这 10 个进程都属于「空闲进程」，随时待命。

进来了一个请求，一个进程前往处理，此时剩下 9 个「空闲进程」，fpm 发现少于 min_spare_servers 设置的值 10 ，就会新建一个进程作为「空闲进程」，此时系统存在 11 个进程，还是 10 个空闲进程。

在第一个请求还未处理完成时，突然一波流量进来，一口气进来了 50 个请求，因为 max_children 设置了 50 个封顶，所以 FPM 会新建 39 个进程，加上 10 个进行进程一起处理这波请求，此时系统中总共 50 个进程共存，50 个进程都属于繁忙中，未分配到进程的请求会等待着。

等所有的请求处理完成后，系统中共存的 50 个进程变成「空闲进程」，超过了 max_spare_servers 值 40 个的限制，超出的 10 个会被销毁，系统此时存在 40 个「空闲进程」，随时待命。

因为一直保证有「空闲进程」可供使用，所以 **dynamic** 的配置，相比 **ondemand** 进程要同时创建，响应速度还是比较快的。然而在还是避免不了频繁创建和销毁进程对系统造成的消耗。

**static**

固定进程数量是性能最好，资源利用率最高的运行方式，一般在要求单机性能最高的时候使用，例如你准备创建 PHP 服务器集群，希望每台机器都能物尽其用。

相关配置：

```

; FPM 启动时创建的进程数，并且会一直保持这个数
pm.max_children = 50

; 每个进程最多处理多少个请求，超过此值将自动销毁
pm.max_requests = 1000

```

`pm.max_children` 的设置，需要我们每一个进程运行我们的程序，需要消耗多少内存，以及机器上有多少内存可供使用。

计算公式：

```

pm.max_children = 可用内存 / 每个进程暂用内存大小

```

注意可用内存不是本机所有内存，要除去其他程序运行，例如说你应该除去 Elasticsearch 占用了 2G 内存。

调试期间，要在生产环境中实战观察，一般建议使用 80% 的内存使用率，留 20% 给内存泄露的空间和其他软件运行。

最后是 pm.max_requests 值，需要我们观察应用是否有 **内存泄漏**。现代的 PHP 程序，尤其是 Laravel ，会依赖于非常多的扩展包，这些扩展包代码质量参差不齐，多少会出现内存泄漏的问题。如果内存泄露不严重，那么把值设置高一点，让单个进程在存活期间 **多处理一些请求**，如果内存泄露比较严重，应当酌情设置低一点。否则会出现系统 **内存不够用**，然后去使用 Swap 物理内存 的窘境。

修改后请记得重启 FPM：

```
sudo service php7.2-fpm restart
```

## Unix Socket 和 TCP Socket

Nginx 连接 FPM 有 Unix Socket 和 TCP Socket 两种方式：

Unix Socket：

```nginx

location ~ \.php$ {
    fastcgi_pass unix:/run/php/php7.2-fpm.sock
}

```

TCP Socket：

```nginx

fastcgi_pass 127.0.0.1:9000;

```

**如何选择？**

Unix Socket 方式 **比较快** 而且 **消耗资源少**（差距不会太大 0.1% ~ 5% 的差别），但是缺点是只能本机使用，没有 TCP 灵活。

如果 Nginx 和 FPM 都在同一台服务器上，推荐使用 Unix Socket。如果是做 PHP 服务器集群，使用 Nginx 做负载均衡的话，只能采用 TCP 的链接方式。

**如何设置成 TCP Socket 的连接方式？**

以 PHP-FPM 7.2 为例：

/etc/php/7.2/fpm/pool.d/www.conf

```

listen = 127.0.0.1:9000

```

同时配置 Nginx 里的 `fastcgi_pass 127.0.0.1:9000;` ，并重启 FPM 和 Nginx：

```sh

sudo service php7.2-fpm restart
sudo service nginx restart

```

**如何设置成 Unix Socket 的连接方式？**

修改 FPM 进程池配置：

/etc/php/7.2/fpm/pool.d/www.conf

```

listen = /run/php/php7.2-fpm.sock

```

同时配置 Nginx 里的 fastcgi_pass unix:/run/php/php7.2-fpm.sock; ，并重启 FPM 和 Nginx：

```sh

sudo service php7.2-fpm restart
sudo service nginx restart

```

文件 `/var/run/php/php7.2-fpm.sock` 会在 FPM 启动时创建。


## 生产环境中一定要关闭掉 Xdebug 扩展

检查生产环境的 PHP 原生扩展文件夹里是否存在，以 PHP 7.2 为例存放路径为：

```
/etc/php/7.2/mods-available/xdebug.ini
```

找到以后确保使用 `;` 符注释掉：

/etc/php/7.2/mods-available/xdebug.ini

```
;zend_extension=xdebug.so
```

重启 fpm

```sh

sudo service php7.2-fpm restart

```

检查 php-fpm 的进程数量

```
# 会多出两个进程数，是因为有一个不负责处理请求的 php-fpm master 进程和一个 grep 进程
ps -aux | grep php-fpm | wc -l
```

## 开启 Slow log 定位慢脚本

### 如何开启？

PHP-FPM 提供一个叫 **慢日志** (slowlog) 的功能，来帮助我们定位执行慢的脚本。以 PHP 7.2 为例，FPM 的配置信息位于：

```
/etc/php/7.2/fpm/pool.d/www.conf
```

相关配置项：

```
; 慢日志的存储路径，默认 `$pool` 设置为 `www`
slowlog = /var/log/$pool.slow.log

; 设置慢日志超时标准，设置为 0 代表关闭慢日志
request_slowlog_timeout = 1s

; 慢日志记录脚本堆栈的深度
request_slowlog_trace_depth = 20
```

以上的配置翻译过来：指定 FPM 当发现有请求执行超过 1 秒钟的时候，将整个调用堆栈记录到 `/var/log/www.slow.log` 文件里，堆栈的深度不超过 20。

你可以把 1s 改成其他值，如 10s。有了以上的设置，裁剪图像尺寸的方法、 网络 I/O 相关的一些请求都经常出现在 PHP 慢日志中。你可以根据自己的情况来选择调整或者忽略。

### 如何分析？

可以使用 grep 命令来快速定位某个函数调用、或者脚本名称被记录的次数，记录的次数越多，优化的优先级就越高。以下是简单的 示例

```
grep -o 'fetch_github_user' /var/log/www.slow.log | wc -l

grep -o 'sendEmail' /var/log/www.slow.log | wc -l

```

**需要注意的是，监控 Slowlog 和记录日志的过程会对 PHP 造成消耗， 切记 调试结束后，务必将其关闭。**


## 开启 OPcache

OPcache 是由 PHP 官方公司 Zend 开发的一款免费使用的 PHP 优化加速拓展。他可以将 PHP 脚本编译后的 bytecode 缓存在共享内存中供以后反复使用，从而避免了从磁盘读取代码再次编译的消耗。同时，它还应用了一些代码优化模式，使得代码执行更快。从而加速 PHP 应用响应。

PHP 自 5.5 版开始，就已经内置了 OPcache 扩展。不过默认是关闭状态的。

### 开启 OPcache

PHP 7.2 FPM 的配置信息位于：

```
/etc/php/7.2/fpm/php.ini
```

编辑以上文件，搜索 `opcache.enable` 将值设为 `1` 即为开启

```
opcache.enable=1
```

`php.ini` 里相关的配置以下，注释里包括说明和推荐设置的值，请详细阅读：

```
; 是否在命令行开启，这里默认设置为 0 ，暂且关闭
;opcache.enable_cli=0

; 这个内存是用来存储编译后的字节码的，视你的程序
; 代码量而定，Laravel 应用一般建议设置为 256，单位 MB，
; 默认是 128
opcache.memory_consumption=256

; 会对程序所有的字符串进行统一存储以加快存取速度，
; 默认是 8m，建议 32 或者不超过 64。
opcache.interned_strings_buffer=32

; 最大加速多少个脚本文件，视项目脚本文件数而定，
; 合理区间 200~1000000 ，默认是 10000 ，建议 500000
opcache.max_accelerated_files=1000000

; 最大作废比例百分比，到达这个比例会重启，默认是 5 ，建议 10
opcache.max_wasted_percentage=10

; 开启情况下会在脚本名称前加上当前目录信息做为缓存的 Key，关闭可以
; 提高性能，但是会面临出错的风险（文件名一致时），建议开启，关闭使用 0
opcache.use_cwd=1

; 开启的话，会按照 opcache.revalidate_freq 设置的频率去检查文件
; 是否修改以便重新缓存，默认开启，生产环境下请设置为关闭，然后
; 写自动化脚本，在每次更新代码后自动重启 OPcache
opcache.validate_timestamps=0

; 文件更新检测频率，单位秒，只有在 opcache.validate_timestamps 
; 开启时才有效。默认为 2，意味着 2 秒钟检查一次，会对文件系统造
; 成负担，如果是在开发环境中请酌情使用，生产环境随意设置，因为
; 我们会设置 validate_timestamps 为关闭。
opcache.revalidate_freq=2200

; 文件加载的逻辑，默认关闭，无需修改
;opcache.revalidate_path=0

; 开启的话会把代码注释一起缓存，关闭可减低内存使用，但是
; 如果有一些代码依赖于注释里的指令，例如 Doctrine, 
; Zend Framework 2 和 PHPUnit，将会出现问题。建议开启
opcache.save_comments=1
```

修改完成后，需要重启 FPM 生效：

```
sudo service php7.2-fpm restart 
```

生产环境下，我们一般会将 `opcache.validate_timestamp` 设置为 0 以获取最大性能。然后在代码变更时候，再重置 OPcache。

有两种重置 OPcache 的方法，一种是重启 FPM。此方法虽然很有效，但是会中断正在处理的请求，用户体验较差，不建议使用。

另一个方法是调用 `opcache_reset()` 方法，此方法会重置 OPcache 缓存并且不需要重启 FPM。然而，OPcache 是运行在 FPM 环境中的，在命令行环境中调用此函数无效。必须是一个可以通过 HTTP 访问到的脚本上来调用 `opcache_reset()` 才行。无法在命令行中执行。

### 在 laravel 中使用 OPcache ，可以直接使用 [laravel-opcache](https://github.com/appstract/laravel-opcache)

> OPcache 是对 PHP 脚本的缓存，每次更改任何 PHP 代码时你都需要清除缓存

```

# 安装
composer require appstract/laravel-opcache

# 清空 fpm 里的 OPcache
php artisan opcache:clear

# 查看 OPcache 的配置信息
php artisan opcache:config

# 查看 OPcache 运行状态（内存使用、缓存了多少文件等）
php artisan opcache:status

# 提前编译文件
php artisan opcache:compile {--force}

```
