---
title: Linux 用户、用户组、权限
author: Alex
top: false
hide: false
cover: false
toc: true
mathjax: false
categories: Linux
tags:
  - Linux
abbrlink: 41bc9f3e
date: 2023-12-19 15:43:10
img:
coverImg:
password:
summary:
---

# Linux 用户、用户组、权限

## 用户

命令介绍 | 命令行
--- | ---
查看所有用户 | cat /etc/passwd  或者 cat /etc/shadow
添加账号 | useradd 比如添加 alex 账号: useradd alex
修改密码 | passwd 比如修改 alex 账号的密码： passwd alex
修改账号 | usermod
删除账号 | userdel
查询账号 | id 比如查询 alex 的账号：id alex
切换账号 | su 比如切换到 alex 账号：su alex

## 用户组

命令介绍 | 命令行
--- | ---
查询所有用户组 | cat /etc/group 或者 cat /etc/gshadow
添加用户组 | groupadd
修改组信息 | groupmod
删除用户组 | groupdel

命令 | 含义
--- | ---
groupadd -g 4000 adminuser | 新建 adminuser 用户组，且 id 为 4000
useradd -G adminuser natasha | 新建 natasha 用户，且 adminuser 作为其附属组
useradd -s /sbin/nologin sarah | 新建 sarah 用户，且在系统中没有任何可交互的 shell （不允许该用户登录系统）
useradd -u 123 jay | 新建 jay 用户，且 id 为 123
echo redhat \| passwd --stdin harry | 新建 harry 用户，且密码为 redhat

## 权限

- ### 读、写、执行
数字表达形式：读（4） 写（2） 执行（1）   
字母表达形式：读（r） 写（w） 执行（x）

- ### 修改文件权限
chmod 比如修改a文件的权限为最高权限chmod  0777 a    
修改aaa文件夹下面所有的文件为最大权限chmod 0777 -R aaa (-R代表递归)   
(或者使用chmod -R 777 aaa)

- ### 修改资源的所有者
  chown
- ### 修改资源所属的用户组
  chgrp
- ### 查看权限信息
  ll

```bash

[root@iZuf6aig35m8ho0xq75ijnZ wwwroot]# ll
total 8
drwxr-xr-x 4 www www 4096 Jun 23 01:55 default
drwxrwxrwx 9 www www 4096 Jun 23 17:26 www.drling.xin
# 对于 default 文件夹权限的解读
# drwxr-xr-x
# d 表示 default 为文件夹，如果是文件的话前面会是 -
# rwx 第2个字母到第4个字母 代表着文件所有者的权限，也就是说 可读可写可执行
# r-x 第5个字母到第7个字母 表示着文件所在的用户组的其他用户的权限，也就是说 可读不可写可执行
# r-x 第8个字母到第10个字母 表示用户组其他的人的权限，也可以理解成陌生人的权限，也就是说 可读不可写可执行
# 4 代表连接数
# www 第一个 www 代表 default 这个文件所在的用户
# www 第二个 www 代表 default 这个文件所在的用户组
# 4096 代表档案容量
# Jun 23 01:55 代表档案最后被修改的时间
# default 文件夹对应的数字权限为 755

```

- 更改 `/var/www/test` 文件所属者为 harry，所属组为 alex

```bash
chown harry:alex /var/www/test
```

- 所有人都不能执行 `/var/www/test` 文件

```bash
chmod a-x /var/www/test
```

- 为特定用户设定特殊权限

```bash
setfacl -Rm u:natasha:rw,u:harry:- /var/www/test

# 查看特定权限
getfacl /var/www/test
```

- 切换用户并执行命令

```bash
# 切换成 www-data 用户，并执行 php artisan tinker 命令
sudo -H -u www-data sh -c 'php artisan tinker'
```