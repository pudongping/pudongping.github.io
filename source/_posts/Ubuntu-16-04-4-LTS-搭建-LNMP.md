---
title: Ubuntu 16.04.4 LTS 搭建 LNMP
author: Alex
top: false
hide: false
cover: false
toc: true
mathjax: false
categories: Ubuntu
tags:
  - Ubuntu
  - LNMP
abbrlink: 7c353a38
date: 2021-07-06 23:38:33
img:
coverImg:
password:
summary:
---

# Ubuntu 16.04.4 LTS 搭建 LNMP

- ## 安装 mysql

```sh
sudo apt-get install mysql-server mysql-client
```
> 测试：mysql -u root -p

查看 Mysql 状态：

```sh
service mysql status/start/stop/retart
```

查看监听端口的情况：

```sh
netstat -tunpl 
或 netstat -tap
```

### 如果之前有装 APACHE 要改下端口,或者直接移除 apache2

```sh
apt-get remove apache2
```

- ##  安装 Nginx
1. 为确保获取最新的 Nginx 先更新源列表（有些文档说的是更新系统）
```sh
sudo apt-get update
```

2. 安装 nginx
```sh
apt-get install nginx
```

3. 检查 nginx 安装是否成功
```sh
# 方法01、 直接输入以下命令可以看到 nginx 的版本号
/usr/sbin/nginx -v

# 方法02、使用浏览器访问 IP 地址，出现 Nginx 的欢迎页面
```

> 可以使用 dpkg -S nginx 命令来搜索 nginx 的相关文件  
> nginx 的配置目录位于：/etc/nginx  
> 其他文档说的是 Nginx 的默认网站目录在 /usr/share/nginx/html/ 下，但是我安装的在 /var/www/html 下，和 apache2 默认的网站目录一样  
> 默认Nginx网站配置文件为 /etc/nginx/sites-available/default


### Nginx 其他命令

```
/etc/init.d/nginx status|stop|start|restart
```

### Nginx 虚拟主机配置
1. 修改 user
```sh
sudo vim /etc/nginx/nginx.conf
# 将 user 改为 www-data 因为 php 默认是这个 user
# 但是我看我自己的 nginx.conf 配置文件第一行直接是
# user www-data; 因此我没有修改
```

2. Nginx 与 php-fpm 的集成  
   这里采用UNIX domain socket方式：（这里一定要注意location的位置！！！）

> 在 /etc/nginx/sites-available/default 配置文件中（网站根目录也在是这里更改）， Nginx已经为与 PHP-FPM的整合准备好了，只需要将下面这部分改好就可以了。sock文件路径为 /run/php/php7.2-fpm.sock 。

```nginx
 51         location ~ \.php$ {
 52                 include snippets/fastcgi-php.conf;
 53 
 54                 # With php7.0-cgi alone:
 55                 #fastcgi_pass 127.0.0.1:9000;
 56                 # With php7.0-fpm:
 57                 fastcgi_pass unix:/run/php/php7.2-fpm.sock;
 58         }
```

> 然后再修改 PHP-FPM 的配置文件 vim  /etc/php/7.2/fpm/pool.d/ www.conf
，如下：

```nginx
 # 与 Nginx监听同一个 sock
 36 listen = /run/php/php7.2-fpm.sock
```

3. 端口-代码映射，方法有三种：

+ 方法01：在 /etc/nginx/sites-available/default 文件末尾处直接添加以下内容
```sh
sudo vim /etc/nginx/sites-available/default
```

提供了两种框架的配置方法

```nginx
# 以下是针对 thinkphp5.0 框架配置
server {
    listen       80;
    # 如果希望二级域名指向同一个项目的话，直接在 server_name 后面添加即可
    server_name  www.drling.xin drling.xin;
    root     /var/www/html/www.drling.xin/public;
    access_log    /var/log/nginx/www.drling.xin_access.log;
    error_log    /var/log/nginx/www.drling.xin_error.log;
    index index.html index.php;

## 在Nginx低版本中，是不支持PATHINFO的，但是可以通过在Nginx.conf中配置转发规则实现

location / {
        if (!-e $request_filename) {
                rewrite ^(.*)$ /index.php?s=$1 last;
                break;
        }
}

    location ~ \.php?.*$ {
        fastcgi_pass   unix:/run/php/php7.2-fpm.sock;
        fastcgi_param  SCRIPT_FILENAME  $document_root$fastcgi_script_name;
        include        fastcgi_params;
    }
}


# 以下是针对 Laravel5.5 框架配置

server {
    listen       80;
    server_name  test.drling.xin;
    root     /var/www/html/test.drling.xin/public;
    access_log    /var/log/nginx/test.drling.xin_access.log;
    error_log    /var/log/nginx/test.drling.xin_error.log;    
    index index.html index.php;

## 如果你使用的是 Nginx，使用如下站点配置指令就可以支持 URL 美化：

location / {
    try_files $uri $uri/ /index.php?$query_string;
}

    location ~ \.php?.*$ {
        fastcgi_pass   unix:/run/php/php7.2-fpm.sock;
        fastcgi_param  SCRIPT_FILENAME  $document_root$fastcgi_script_name;
        include        fastcgi_params;
    }
}
```

+ 方法02、直接在 /etc/nginx/sites-available 目录下，单独写自己的项目配置文件，操作方法如下

```
# 01、比如我写的文件是 www.drling.xin 不需要写 .conf 后缀名
sudo vim /etc/nginx/sites-available/www.drling.xin

# 02、将以上 thinkphp5.0 的配置文件放进去

# 以下是针对 thinkphp5.0 框架配置
server {
    listen       80;
    server_name  www.drling.xin;
    root     /var/www/html/www.drling.xin/public;    
    access_log    /var/log/nginx/www.drling.xin_access.log;
    error_log    /var/log/nginx/www.drling.xin_error.log; 
    
    index index.html index.php;

## 在Nginx低版本中，是不支持PATHINFO的，但是可以通过在Nginx.conf中配置转发规则实现

location / {
        if (!-e $request_filename) {
                rewrite ^(.*)$ /index.php?s=$1 last;
                break;
        }
}

    location ~ \.php?.*$ {
        fastcgi_pass   unix:/run/php/php7.2-fpm.sock;
        fastcgi_param  SCRIPT_FILENAME  $document_root$fastcgi_script_name;
        include        fastcgi_params;
    }
}

# 03、建立软连接
sudo ln -s /etc/apache2/sites-available/www.drling.xin /etc/apache2/sites-enabled/www.drling.xin

```

+ 方法03、直接在 /etc/nginx/conf.d/ 目录下写以 .conf 为后缀的配置文件

```
# 01、比如我写的文件是 test.drling.xin 需要写 .conf 后缀名
sudo vim /etc/nginx/conf.d/test.drling.xin.conf

# 02、将以上 Laravel5.5 的配置文件放进去

server {
    listen       80;
    server_name  test.drling.xin;
    root     /var/www/html/test.drling.xin/public;
    access_log    /var/log/nginx/test.drling.xin_access.log;
    error_log    /var/log/nginx/test.drling.xin_error.log;    
    index index.html index.php;

## 如果你使用的是 Nginx，使用如下站点配置指令就可以支持 URL 美化：

location / {
    try_files $uri $uri/ /index.php?$query_string;
}

    location ~ \.php?.*$ {
        fastcgi_pass   unix:/run/php/php7.2-fpm.sock;
        fastcgi_param  SCRIPT_FILENAME  $document_root$fastcgi_script_name;
        include        fastcgi_params;
    }
}
```

> 其实 nginx 的配置放置位置多半可以参考 /etc/nginx/nginx.conf 文件中61行和62行  
>  61         include /etc/nginx/conf.d/*.conf;  
>  62         include /etc/nginx/sites-enabled/*;

- ### 添加完了之后一定要重启服务器
```
sudo service nginx restart
```

- ### Nginx 配置文件需要注意位置讲解
1. 如果不确定 fastcgi_pass 这一段怎么填写
```sh
#fastcgi_pass为fpm地址，可查看/etc/php/7.2/fpm/pool.d/www.conf中的listen确定（www.conf配置文件中第36行）
```
2. 安装 php 的时候一定要安装 php-fpm 扩展（也称为：PHPFastCGI管理器），要连接数据库则必须安装 php-mysql 扩展
```sh
# 我在这里安装的 php7.2 因为 Laravel5.5核心代码中需要php7.1以上
sudo apt-get install php7.2-fpm php7.2-mysql
```
3. 可以使用如下命令查看端口监听情况
```sh
netstat -anp

// 直接查看80端口可以使用命令：
sudo lsof -i :80
```

4. 一定要在 php.ini 配置文件中修改 cgi.fix_pathinfo=0
```sh
sudo vim /etc/php7.2/fpm/php.ini
```
查找到 cgi.fix_pathinfo 字段（第776行），将值设为0
> 这个参数用来对设置cgi模式下为php是否提供绝对路径信息或PATH_INFO信息

5. Nginx 与 PHP-FPM 集成详细讲解

> PHP-FPM 与 Nginx 通信方式有两种，一种是基于TCP的 Internet domain socket 方式，一种是 UNIX domain socket 方式。
>
> UNIX domain socket 可以使同一台操作系统上的两个或多个进程进行数据通信。UNIX domain socket 的接口和 Internet domain socket 很像，但它不使用网络底层协议来通信。
>
> 服务器压力不大的情况下，这两种方式性能差别不大，但在压力比较满的时候，用UNIX domain socket方式，效果确实比较好。
> 接下来分别讲解一下两种方式的配置方式：

- 使用默认的 UNIX Socket方式:
1. 修改 fpm 的配置文件
```sh
sudo vim /etc/php/7.2/fpm/pool.d/www.conf
```
修改成如下：
```sh
# 第36行
 listen = /run/php/php7.2-fpm.sock
```
> 可以使用如下方式检查下配置文件是否有错误

```sh
sudo php-fpm7.2 -t
```
> 修改完了之后重启一下 php-fpm7.2

```sh
sudo service php-fpm7.2 restart
```
2. 修改 nginx 的配置文件

```sh
sudo vim /etc/nginx/sites-enabled/default
```
修改成如下：
```nginx
 51         location ~ \.php$ {
 52                 include snippets/fastcgi-php.conf;
 53 
 54                 # With php7.0-cgi alone:
 55                 #fastcgi_pass 127.0.0.1:9000;
 56                 # With php7.0-fpm:
 57                 fastcgi_pass unix:/run/php/php7.2-fpm.sock;
 58         }

```
> 修改完了之后重启下 nginx

```sh
sudo service nginx restart
```
> 再次检查 nginx 的配置文件是否正确

```sh
sudo php-fpm7.2 -t
```

- 使用 TCP 方式:

1. 修改 fpm 的配置文件
```sh
sudo vim /etc/php/7.2/fpm/pool.d/www.conf
```
修改成如下：
```sh
 # 配置文件第36行
 listen = 127.0.0.1:9000
```

2. 修改 nginx 的配置文件

```sh
sudo vim /etc/nginx/sites-enabled/default
```
修改成如下：
```nginx
 51         location ~ \.php$ {
 52                 include snippets/fastcgi-php.conf;
 53 
 54                 # With php7.0-cgi alone:
 55                 #fastcgi_pass 127.0.0.1:9000;
 56                 # With php7.0-fpm:
 57                 fastcgi_pass 127.0.0.1:9000;
 58         }

```
