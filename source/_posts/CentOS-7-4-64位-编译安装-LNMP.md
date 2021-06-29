---
title: CentOS 7.4 64位 编译安装 LNMP
author: Alex
top: false
hide: false
cover: false
toc: true
mathjax: false
categories: CentOS
tags:
  - CentOS
  - LNMP
abbrlink: 2ccaa736
date: 2021-06-29 10:06:34
img:
coverImg:
password:
summary:
---

# CentOS 7.4 64位 编译安装 LNMP

## 查看 Linux 版本

```shell
cat /etc/redhat-release

# CentOS Linux release 7.4.1708 (Core)
```

## 1. 安装 nginx

1-1. 安装 nginx 源

```shell
yum localinstall http://nginx.org/packages/centos/7/noarch/RPMS/nginx-release-centos-7-0.el7.ngx.noarch.rpm
```

1-2. 安装 nginx

```shell
yum install nginx
```

1-3. 启动 nginx
```shell
systemctl start nginx
```

## 2. 安装 MySQL

2-1. 安装 MySQL 源

```shell
yum localinstall  http://dev.mysql.com/get/mysql57-community-release-el7-7.noarch.rpm
```

2-2. 安装 MySQL

```shell
yum install mysql-community-server
```

安装 MySQL 开发包 （*）

```shell
yum install mysql-community-devel
```

2-3. 启动 MySQL

```shell
systemctl start mysqld
```

2-4. 查看 MySQL 是否启动

```shell
systemctl status mysqld
```

2-5. 查看 MySQL 默认密码

```shell
# 首次启动,会把密码放在 /var/log/mysqld.log 里面
# 2018-10-13T15:51:47.482124Z 1 [Note] A temporary password is generated for root@localhost: r)eS,gjku4ts

grep 'temporary password' /var/log/mysqld.log
```

2-6. 更改 MySQL 密码

```shell

# 01、登入数据库
mysql -u root -p 
# 输入密码
r)eS,gjku4ts
# 这里的密码 r)eS,gjku4ts 是从以上 mysqld.log 中查询出来的

# 02、修改 root 账号密码 （密码安全级别要稍微高一点，不然更新不会成功）
ALTER USER 'root'@'localhost' IDENTIFIED BY 'QAZwsx123!@#';
```

2-7. 将 root 用户更改为外网也可以访问

```shell

# 01、打开 mysql 数据表
use mysql;

# 02、查看 mysql 数据表中数据
select user,host from user;

# 03、 % 代表任何 ip 都能访问
UPDATE user SET host = '%' WHERE user = 'root';

# 04、再次查看 mysql 数据表中的数据
select user,host from user;

# 05、刷新权限
flush privileges;
```

2-8. 新建一个用户并且赋予权限（因为 root 用户开放外面使用毕竟不安全）

```shell

# 01、新建 mysql 账户

grant all privileges on *.* to  alex@"%" identified by "QAZwsx123!@#" with grant option;

# grant 是授权命令，其中 alex 是我们连接用的用户名、"AZwsx123!@#"是连接密码，用户名后面的 "%" 通用符表示允许各 host 操作。

# 上面这条命令是指
# 自动创建用户 alex ,密码 AZwsx123!@#
# 格式：grant 权限 on 数据库名.表名 to 用户@登录主机 identified by "用户密码"; 
# @ 后面是访问mysql的客户端IP地址（或是 主机名） % 代表任意的客户端，如果填写 localhost 为本地访问（那此用户就不能远程访问该mysql数据库了）。
	
# 02、刷新权限
flush privileges;	
```	

## 3. 编译安装 php7.2

3-1. 下载安装包，一般情况下,我们都会下载到 /usr/local/src 下面

```shell

# 01、进入 src 目录
cd /usr/local/src

# 02、下载 php7.2 源码包
wget -O php72.tar.gz http://cn2.php.net/get/php-7.2.0.tar.gz/from/this/mirror

# 03、解压安装包
tar -zxvf php72.tar.gz

# 04、进入解压后的包
cd php-7.2.0

# 05、安装 php 的依赖
yum install libxml2 libxml2-devel openssl openssl-devel bzip2 bzip2-devel libcurl libcurl-devel libjpeg libjpeg-devel libpng libpng-devel freetype freetype-devel gmp gmp-devel libmcrypt libmcrypt-devel readline readline-devel libxslt libxslt-devel

# 06、新建 php 这个文件夹（编译配置会用到）
mkdir -p /usr/local/php

# 07、编译配置
	./configure \
--prefix=/usr/local/php \
--with-config-file-path=/etc \
--enable-fpm \
--with-fpm-user=nginx  \
--with-fpm-group=nginx \
--enable-inline-optimization \
--disable-debug \
--disable-rpath \
--enable-shared  \
--enable-soap \
--with-libxml-dir \
--with-xmlrpc \
--with-openssl \
--with-mcrypt \
--with-mhash \
--with-pcre-regex \
--with-sqlite3 \
--with-zlib \
--enable-bcmath \
--with-iconv \
--with-bz2 \
--enable-calendar \
--with-curl \
--with-cdb \
--enable-dom \
--enable-exif \
--enable-fileinfo \
--enable-filter \
--with-pcre-dir \
--enable-ftp \
--with-gd \
--with-openssl-dir \
--with-jpeg-dir \
--with-png-dir \
--with-zlib-dir  \
--with-freetype-dir \
--enable-gd-native-ttf \
--enable-gd-jis-conv \
--with-gettext \
--with-gmp \
--with-mhash \
--enable-json \
--enable-mbstring \
--enable-mbregex \
--enable-mbregex-backtrack \
--with-libmbfl \
--with-onig \
--enable-pdo \
--with-mysqli=mysqlnd \
--with-pdo-mysql=mysqlnd \
--with-zlib-dir \
--with-pdo-sqlite \
--with-readline \
--enable-session \
--enable-shmop \
--enable-simplexml \
--enable-sockets  \
--enable-sysvmsg \
--enable-sysvsem \
--enable-sysvshm \
--enable-wddx \
--with-libxml-dir \
--with-xsl \
--enable-zip \
--enable-mysqlnd-compression-support \
--with-pear \
--enable-opcache

# 08、编译与安装 （此处需要时间，耐心等待）
make && make install
```

3-2. 到这里已经算是安装完成了，查看 php 版本，就会出现熟悉的php 7.2.0 xxxxxx

```shell
# 查看 php 版本
/usr/local/php/bin/php -v
```

3-3. 但是这样,我们没有添加环境变量,太麻烦了,接下来把 php 放到环境变量里面

```shell
# 01、打开文件
vim /etc/profile

# 02、在 profile 文件最底部加入
PATH=$PATH:/usr/local/php/bin
export PATH

# 03、让修改立即生效
source /etc/profile  或者  ./etc/profile
```

3-3. 此时我们查看 PHP 版本 php -v 就行了

```shell

# 将php.ini复制到/etc/下面
cp php.ini-production /etc/php.ini

cp /usr/local/php/etc/php-fpm.conf.default /usr/local/php/etc/php-fpm.conf

cp /usr/local/php/etc/php-fpm.d/www.conf.default /usr/local/php/etc/php-fpm.d/www.conf

cp sapi/fpm/init.d.php-fpm /etc/init.d/php-fpm

chmod +x /etc/init.d/php-fpm

```

3-4. 启动 php-fpm

```shell
/etc/init.d/php-fpm start
```

## 4. 配置 nginx ，使得 nginx 能够解析 php

4-1. 打开 nginx 配置文件

```shell
vim /etc/nginx/conf.d/default.conf
```

默认配置文件中的内容如下：

```nginx

  1 server {
  2     listen       80;
  3     server_name  localhost;
  4 
  5     #charset koi8-r;
  6     #access_log  /var/log/nginx/host.access.log  main;
  7 
  8     location / {
  9         root   /usr/share/nginx/html;
 10         index  index.html index.htm;
 11     }
 12 
 13     #error_page  404              /404.html;
 14 
 15     # redirect server error pages to the static page /50x.html
 16     #
 17     error_page   500 502 503 504  /50x.html;
 18     location = /50x.html {
 19         root   /usr/share/nginx/html;
 20     }
 21 
 22     # proxy the PHP scripts to Apache listening on 127.0.0.1:80
 23     #
 24     #location ~ \.php$ {
 25     #    proxy_pass   http://127.0.0.1;
 26     #}
 27 
 28     # pass the PHP scripts to FastCGI server listening on 127.0.0.1:9000
 29     #
 30     #location ~ \.php$ {
 31     #    root           html;
 32     #    fastcgi_pass   127.0.0.1:9000;
 33     #    fastcgi_index  index.php;
 34     #    fastcgi_param  SCRIPT_FILENAME  /scripts$fastcgi_script_name;
 35     #    include        fastcgi_params;
 36     #}
 37 
 38     # deny access to .htaccess files, if Apache's document root
 39     # concurs with nginx's one
 40     #
 41     #location ~ /\.ht {
 42     #    deny  all;
 43     #}
 44 }
```

4-2. 复制 nginx 默认配置文件 default.conf ，写自定义配置文件 `www.drling.xin.conf`

```shell
cp /etc/nginx/conf.d/default.conf /etc/nginx/conf.d/www.drling.xin.conf
```

4-3. 将以下内容写进 `www.drling.xin.conf` 配置文件中去

```nginx

server {
    listen       80;
    server_name  www.drling.xin;


    location / {
        root   /usr/share/nginx/html;
        index index.php index.html index.htm;
    }


    location ~ \.php$ {
        root           /usr/share/nginx/html;
        fastcgi_pass   127.0.0.1:9000;
        fastcgi_index  index.php;
        #fastcgi_param  SCRIPT_FILENAME  /scripts$fastcgi_script_name;
        fastcgi_param  SCRIPT_FILENAME  $document_root$fastcgi_script_name;
        include        fastcgi_params;
    }


}
```

4-4. 查看 nginx 配置写入是否正确

```shell
nginx -t
```

4-5. 重启 nginx

```shell
systemctl restart nginx
```

4-6. 在 `/usr/share/nginx/html` 下面新建一个 index.php 文件，写入以下内容

```php
<?
phpinfo();
?>
```

4-7. 在浏览器中访问配置文件中 server_name 后写的域名，就可以看到 phpinfo 信息了，我这里是直接在浏览器中访问 `www.drling.xin`

4-8. 后续

这里提供 tp5 的 nginx 配置文件写法

```nginx

server {
    listen       80;
    server_name  www.drling.xin;
    access_log    /var/log/nginx/www.drling.xin_access.log;
    error_log    /var/log/nginx/www.drling.xin_error.log;

    location / {

        root   /var/www/my_cake_test/public;

        index index.php index.html index.htm;

        if (!-e $request_filename) {
                rewrite ^(.*)$ /index.php?s=$1 last;
                break;
        }
        
    }
    
    location ~ \.php$ {
        root           /var/www/my_cake_test/public;
        fastcgi_pass   127.0.0.1:9000;
        fastcgi_index  index.php;
        #fastcgi_param  SCRIPT_FILENAME  /scripts$fastcgi_script_name;
        fastcgi_param  SCRIPT_FILENAME  $document_root$fastcgi_script_name;
        include        fastcgi_params;
    }


}

```

以下为 laravel5.5 的 nginx 配置文件写法

```nginx

server {

    listen       80;
    server_name  sample.drling.xin;
    access_log    /var/log/nginx/sample.drling.xin_access.log;
    error_log    /var/log/nginx/sample.drling.xin_error.log;
    root   /var/www/sample/public;
    index index.php index.html index.htm;

    location / {
        try_files $uri $uri/ /index.php?$query_string;
    }   

        
        
    location ~ \.php$ {
        root           /var/www/sample/public;
        fastcgi_pass   127.0.0.1:9000;
        fastcgi_index  index.php;
        #fastcgi_param  SCRIPT_FILENAME  /scripts$fastcgi_script_name;
        fastcgi_param  SCRIPT_FILENAME  $document_root$fastcgi_script_name;
        include        fastcgi_params;
    }


}
```
