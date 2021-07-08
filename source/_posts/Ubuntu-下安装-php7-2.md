---
title: Ubuntu 下安装 php7.2
author: Alex
top: false
hide: false
cover: false
toc: true
mathjax: false
categories: Ubuntu
tags:
  - Ubuntu
  - Linux
abbrlink: 98509b56
date: 2021-07-08 09:29:25
img:
coverImg:
password:
summary:
---

# Ubuntu 下安装 php7.2

1. 安装软件源拓展工具：

```sh
sudo apt-get install software-properties-common python-software-properties
```

2. 更新软件源缓存并且添加 `Ondřej Surý` 的 `PHP PPA` 源，需要按一次回车：

```sh
sudo add-apt-repository ppa:ondrej/php && sudo apt-get update
```

3. 安装 php7.2

```sh
sudo apt-get -y install php7.2
```

4. 如果之前有其他版本PHP，在这边禁用掉

```sh
# 禁用掉 php5.6 的版本
sudo a2dismod php5.6
# 开启 php7.2 的版本
sudo a2enmod php7.2
```

5. 安装 php7.2 常用扩展

```sh
apt install php7.2-fpm php7.2-mysql php7.2-curl php7.2-gd php7.2-mbstring php7.2-xml php7.2-xmlrpc php7.2-zip php7.2-opcache -y
```

查看可以安装的扩展指令

```sh
apt-cache search php7.2

# 比如如下所示：
root@iZuf6aig35m8ho0xq75ijnZ:/# apt-cache search php7.2

php-amqp - AMQP extension for PHP
php-apcu - APC User Cache for PHP
php-geoip - GeoIP module for PHP
php-igbinary - igbinary PHP serializer
php-imagick - Provides a wrapper to the ImageMagick library
php7.2-fpm - server-side, HTML-embedded scripting language (FPM-CGI binary)
php7.2-gd - GD module for PHP
php7.2-gmp - GMP module for PHP
php7.2-imap - IMAP module for PHP
```

6. 设置 php

安装完成后，编辑 `/etc/php/7.2/fpm/php.ini` 替换成 `;cgi.fix_pathinfo=1` 为 `cgi.fix_pathinfo=0` 快捷命令：

```sh
# 直接命令行输入
sed -i 's/;cgi.fix_pathinfo=1/cgi.fix_pathinfo=0/' /etc/php/7.2/fpm/php.ini
```

7. 管理 php

```sh
systemctl restart php7.2-fpm #重启
systemctl start php7.2-fpm #启动
systemctl stop php7.2-fpm #关闭
systemctl status php7.2-fpm #检查状态
```

8. 重启 apache2

```sh
sudo service apache2 restart
```
