---
title: 全新服务器使用 lnmp 搭建 laravel 项目
author: Alex
top: false
hide: false
cover: false
toc: true
mathjax: false
categories: PHP
tags:
  - PHP
  - LNMP
  - Laravel
abbrlink: 73f4393e
date: 2021-08-22 16:19:05
img:
coverImg:
password:
summary:
---

# 使用 lnmp 搭建 laravel 项目

- 打包源码

```shell

# 先在旧服务器上面打包源代码
tar -czvf hello-world.tar.gz hello-world

# 然后将源码直接复制到新的服务器上面
scp hello-world.tar.gz root@192.168.1.1:/home/wwwroot/

```

- 安装 screen

```shell

yum install -y screen

screen -S lnmp

```

- 下载安装 lnmp 环境

```shell

wget http://soft.vpser.net/lnmp/lnmp1.7-full.tar.gz

tar -xzvf lnmp1.7-full.tar.gz

cd lnmp1.7-full/

# 安装 lnmp
./install.sh lnmp

# 选择了 php7.4 mysql5.7

# 检查 lnmp 是否安装成功
lnmp status

```

- 安装 swoole 扩展

```shell

# 安装 swoole 源码
wget https://github.com/swoole/swoole-src/archive/v4.6.4.tar.gz

# 解压缩源码
tar -xzvf v4.6.4.tar.gz

# 生成 configure 文件
/usr/local/php/bin/phpize

# 编译配置项
./configure --with-php-config=/usr/local/php/bin/php-config --enable-openssl --enable-http2 --enable-sockets --enable-mysqlnd

# 编译安装
make && make install

# 查看 php.ini 配置文件
php --ini

# 开启 swoole 扩展
vim /usr/local/php/etc/php.ini

# 在 php.ini 配置文件中开启 swoole 扩展
extension=swoole.so
# 或者写绝对路径
extension=/usr/local/php/lib/php/extensions/no-debug-non-zts-20190902/swoole.so

# 重启 php
/etc/init.d/php-fpm restart

# 查看扩展是否已经开启
php -m  
php --ri swoole

```

- 安装 redis

```shell

# 下载最新稳定版 6.2.1 源码
wget https://download.redis.io/releases/redis-6.2.1.tar.gz

# 解压缩
tar xzf redis-6.2.1.tar.gz
cd redis-6.2.1

# 编译安装
make

# 启动 redis 服务端
src/redis-server （开启后台任务 src/redis-server &）

# 使用 redis 客户端作为测试
src/redis-cli 或者 cd src && ./redis-cli

# 将 redis-cli 加入到环境变量中
cp ~/software/redis-6.2.1/src/redis-cli /usr/local/bin/redis-cli

# 设置执行权限
chmod u+x redis-cli

```

- 安装 php redis 扩展

```shell

# 安装 redis 扩展
wget https://github.com/phpredis/phpredis/archive/5.3.3.tar.gz
tar -xzvf 5.3.3.tar.gz
cd phpredis-5.3.3

# 生成 configure 文件
/usr/local/php/bin/phpize

# 设置配置项
./configure --with-php-config=/usr/local/php/bin/php-config

# 编译安装
make && make install

# 重启 php
/etc/init.d/php-fpm restart

# 查看扩展是否已经开启
php -m  
php --ri redis

```

- 安装 npm

```shell

yum -y install npm

npm config set registry=https://registry.npm.taobao.org

```

- 配置项目

vim /usr/local/nginx/conf/vhost/hello-world.com.conf

添加 nginx 配置

```shell

server
    {
        listen 80;
        #listen [::]:80;
        server_name www.hello-world.com ;
        index index.html index.htm index.php default.html default.htm default.php;
        root  /home/wwwroot/hello-world/public;

        include rewrite/none.conf;

        include enable-php.conf;

        location ~ .*\.(gif|jpg|jpeg|png|bmp|swf)$
        {
            expires      30d;
        }

        location ~ .*\.(js|css)?$
        {
            expires      12h;
        }

        location ~ /.well-known {
            allow all;
        }

        location ~ /\.
        {
            deny all;
        }

        location / {
          try_files $uri $uri/ /index.php?$query_string;
        }

        access_log  /home/wwwlogs/hello-world.com.access.log;
        error_log  /home/wwwlogs/hello-world.com.error.log;
    }
    
server {
  listen 80;
  server_name hello-world.com;
  rewrite ^/(.*) http://www.hello-world.com/$1 permanent;
}

```

删除防跨文件夹设置  
vim /usr/local/nginx/conf/fastcgi.conf  
注释掉  
`fastcgi_param PHP_ADMIN_VALUE "open_basedir=$document_root/:/tmp/:/proc/";`

```shell

chattr -i /home/wwwroot/hello-world/public/.user.ini

rm -rf /home/wwwroot/hello-world/public/.user.ini

# 重启 php-fpm
/etc/init.d/php-fpm restart

# 重新加载 nginx
/etc/init.d/nginx reload

```

vim .env

更改 mysql 数据库、redis连接信息、以及配置域名

```shell

php artisan jwt:secret

php artisan key:generate

```
