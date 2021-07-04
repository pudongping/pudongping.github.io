---
title: CentOS 7.4 64位 搭建LAMP
author: Alex
top: false
hide: false
cover: false
toc: true
mathjax: false
categories: CentOS
tags:
  - CentOS
  - LAMP
abbrlink: 10ab7ab8
date: 2021-07-04 23:56:20
img:
coverImg:
password:
summary:
---

# CentOS 7.4 64位 搭建LAMP

- ## 安装php7.2
若直接采用centos中的yum安装：sudo yum -y install php，版本是5.4，远远不够，因此我们要手动更新rpm即可。

### 1. 首先获取 rpm （添加 php 的 yum 仓库 ）：

```sh
rpm -Uvh https://dl.fedoraproject.org/pub/epel/epel-release-latest-7.noarch.rpm   
rpm -Uvh https://mirror.webtatic.com/yum/el7/webtatic-release.rpm    
```

然后可以利用 sudo yum list php* 查看目前都有 php 的什么版本了，可以发现从 4-7.2 的版本都有，7.2 版本名为 72w，因此安装该版本即可：

### 2. 安装 php7.2

```sh
sudo yum -y install php72w
```

但安装完毕后，输入 `php -v` 发现并没有该命令，因为 php72w 只是安装了 php 最小的库，一些应用还未安装，因此安装一些拓展包即可：

### 3. 安装 php7.2 其它扩展（安装过程中全部选 yes 即可）

```sh
# 以下扩展中有部分和上面的扩展重复，安装的时候请注意区分

sudo yum install php72w.x86_64 php72w-cli.x86_64 php72w-common.x86_64 php72w-gd.x86_64 php72w-ldap.x86_64 php72w-mbstring.x86_64 php72w-mcrypt.x86_64 php72w-mysql.x86_64 php72w-pdo.x86_64 php72w-devel.x86_64
```

### 4. 安装 PHP 7.2 的 fpm

```sh
sudo yum install php72w-fpm.x86_64
```

- ## 安装 MySQL5.7 或 MySQL5.6

### 1. 配置 YUM 源

```sh
wget http://dev.mysql.com/get/mysql57-community-release-el7-8.noarch.rpm

sudo yum localinstall mysql57-community-release-el7-8.noarch.rpm

# 检查 mysql 源是否安装成功
yum repolist enabled | grep "mysql.*-community.*"
```

### 2. 修改安装 mysql 版本配置（现在默认安装是 mysql5.7）

可以修改 `vim /etc/yum.repos.d/mysql-community.repo`  源，改变默认安装的 mysql 版本。比如要安装 5.7 版本，将 5.6 源的 `enabled=1` 改成 `enabled=0` 。然后再将 5.7 源的 `enabled=0` 改成 `enabled=1` 即可。改完之后的效果如下所示：

```sh
 15 # Enable to use MySQL 5.5
 16 [mysql55-community]
 17 name=MySQL 5.5 Community Server
 18 baseurl=http://repo.mysql.com/yum/mysql-5.5-community/el/7/$basearch/
 19 enabled=0
 20 gpgcheck=1
 21 gpgkey=file:///etc/pki/rpm-gpg/RPM-GPG-KEY-mysql
 22 
 23 # Enable to use MySQL 5.6
 24 [mysql56-community]
 25 name=MySQL 5.6 Community Server
 26 baseurl=http://repo.mysql.com/yum/mysql-5.6-community/el/7/$basearch/
 27 enabled=0
 28 gpgcheck=1
 29 gpgkey=file:///etc/pki/rpm-gpg/RPM-GPG-KEY-mysql
 30 
 31 [mysql57-community]
 32 name=MySQL 5.7 Community Server
 33 baseurl=http://repo.mysql.com/yum/mysql-5.7-community/el/7/$basearch/
 34 enabled=1
 35 gpgcheck=1
 36 gpgkey=file:///etc/pki/rpm-gpg/RPM-GPG-KEY-mysql
 ```

### 3. 安装 MySQL

 ```sh
 yum install mysql-community-server
 ```

### 4. 启动 MySQL

 ```sh
 systemctl start mysqld
 ```

### 5. 配置开机启动

 ```sh
 systemctl enable mysqld
 systemctl daemon-reload
 ```

### 6. 修改 root 本地登录密码
mysql 安装完成之后，在 /var/log/mysqld.log 文件中给 root 生成了一个默认密码。通过下面的方式找到 root 默认密码，然后登录 mysql 进行修改：

 ```sh
 # 查找默认生成的密码
 grep 'temporary password' /var/log/mysqld.log
 
 # 用默认生成的密码登录 mysql
 mysql -uroot -p
 
 # 修改新密码为 MyNewPass4! 
 # 用这种方式修改密码的时候需要把密码修改的稍微复杂一点，不然老是提示你创建的密码不安全
ALTER USER 'root'@'localhost' IDENTIFIED BY 'MyNewPass4!';
```

### 7. 配置默认编码为 utf8

```sh
# 直接命令行下敲命令，不需要登录 mysql 之后
character_set_server=utf8
init_connect='SET NAMES utf8'
```

- ## 安装 Apache

### 1. 安装 httpd

```sh
yum -y install httpd
```

### 2. 修改 apache 配置文件

主配置文件的路径为 /etc/httpd/conf/httpd.conf    
扩展配置文件路径为 /etc/httpd/conf.d/*.conf

```sh
vim /etc/httpd/conf.d/sample.conf

# 配置文件编写如下

 <VirtualHost *:80>
   # 项目文件目录，默认项目目录在 /var/www/html/ 下
   DocumentRoot /var/www/html/sample/public
   # 虚拟域名
   ServerName test.drling.xin
   # 多个虚拟域名
   ServerAlias sample.drling.xin
   # 如果还需要配置第三个虚拟域名
   ServerAlias www.drling.xin
   # 错误日志目录
   ErrorLog /var/log/httpd/sample-error_log
   # 访问日志目录
   Customlog /var/log/httpd/sample-access_log common
 </VirtualHost>
```

### 3. 启动 httpd

```sh
# 开启/重启/停止/状态
systemctl start/restart/stop/status httpd 

# 开机启动 httpd
systemctl enable httpd
```

**以下步骤我搭建服务器的时候没有操作，发现项目也运行起来了**

### 4. 修改 httpd.conf 配置

```sh
# line 86: 改变管理员的邮箱地址
ServerAdmin root@linuxprobe.org
# line 95: 改变域名信息
ServerName www.linuxprobe.org:80
# line 151: none变成All
AllowOverride All
# line 164: 添加只能使用目录名称访问的文件名
DirectoryIndex index.html index.cgi index.php
# add follows to the end
# server's response header（安全性）
ServerTokens Prod
# keepalive is ON
KeepAlive On
```


### 5. 修改防火墙配置

```sh
如果 Firewalld 正在运行，请允许 HTTP 服务。HTTP 使用 80/TCP
firewall-cmd --add-service=http --permanent
firewall-cmd --reload
```

### 6. 运行 tp5 项目的时候提示没有写 session 的权限，默认 session 文件夹如下

```sh
# session 文件夹路径： /var/lib/php/session/

# 给最大权限 
chmod 0777 -R /var/lib/php/session
```
