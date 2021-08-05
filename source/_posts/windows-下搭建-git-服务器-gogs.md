---
title: windows 下搭建 git 服务器 gogs
author: Alex
top: false
hide: false
cover: false
toc: true
mathjax: false
categories: Git
tags:
  - Git
  - Gogs
  - Git 服务器
abbrlink: a90b1e44
date: 2021-08-05 20:26:49
img:
coverImg:
password:
summary:
---

# 本文基于 windows7 64位 搭建 gogs
> gogs 官方文档地址：https://gogs.io/docs    
软件下载地址：https://dl.gogs.io/

# 环境要求
*   数据库（选择以下一项）：
    *   [MySQL](http://dev.mysql.com/)：版本 >= 5.7
    *   [PostgreSQL](http://www.postgresql.org/)
    *   [MSSQL](https://en.wikipedia.org/wiki/Microsoft_SQL_Server)
    *   [TiDB](https://github.com/pingcap/tidb)（实验性支持，使用 MySQL 协议连接）
    *   或者 **什么都不安装** 直接使用 SQLite3
*   [git](http://git-scm.com/)（bash）：
    *   服务端和客户端均需版本 >= 1.7.1
    *   Windows 系统建议使用最新版
*   SSH 服务器：
    *   **如果您只使用 HTTP/HTTPS 的话请忽略此项**
    *   如果您选择在 Windows 系统使用内置 SSH 服务器，请确保添加 `ssh-keygen` 到您的 `%PATH%` 环境变量中
    *   推荐 Windows 系统使用 [Cygwin OpenSSH](http://docs.oracle.com/cd/E24628_01/install.121/e22624/preinstall_req_cygwin_ssh.htm) 或 [Copssh](https://www.itefix.net/copssh)
    *   Windows 系统 请确保 Bash 是默认的 Shell 程序，而不是 PowerShell

# 所需软件
* 必须软件
    * [NSSM](http://nssm.cc/download)
    * [git](https://git-scm.com/downloads)  最好下载最新版
    * [MySQL](https://dev.mysql.com/downloads/mysql/) 官方说的是版本需要大于5.7，我的版本是 5.5.3 发现也并无影响。**但是存储引擎一定要使用：INNODB！**
      ![此时的版本为5.5.3](https://upload-images.jianshu.io/upload_images/14623749-ef2d29505d1077d6.png?imageMogr2/auto-orient/strip%7CimageView2/2/w/1240)
    * [gogs软件包](https://dl.gogs.io/0.10.1/windows_amd64.zip) windows-64位版本

# 安装
1. 将下载的 gogs_0.11.86_windows_amd64.zip 压缩包文件解压。
> 本文解压在 E:\soft-exe 目录下

![解压出来是 gogs 文件夹](https://upload-images.jianshu.io/upload_images/14623749-90a40366e1de646b.png?imageMogr2/auto-orient/strip%7CimageView2/2/w/1240)

2. 创建数据库
> 可以直接去执行 E:\soft-exe\gogs\scripts\mysql.sql 文件，创建 gogs 数据库。（当然也可以自己去创建数据库名为 gogs 的数据库，但是建议还是直接执行 mysql.sql 脚本，毕竟官方建议）

mysql.sql 中的内容为以下：
```
SET GLOBAL innodb_file_per_table = ON,
           innodb_file_format = Barracuda,
           innodb_large_prefix = ON;
DROP DATABASE IF EXISTS gogs;
CREATE DATABASE IF NOT EXISTS gogs CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci;

```
3. 安装 git
   这个貌似没有什么太多需要讲的，直接根据上面提供的链接地址下载 git 最新版，之后一直下一步安装即可。
4. 注册 gogs 服务
* 修改 E:\soft-exe\gogs\scripts\windows\install-as-service.bat ,将其中的
```
SET gogspath=C:/gogs
```
修改成你本地的 gogs 安装路径。

![找到 install-as-service.bat 文件](https://upload-images.jianshu.io/upload_images/14623749-8b353dd27ef518c9.png?imageMogr2/auto-orient/strip%7CimageView2/2/w/1240)

![修改 gogspath 的值为 gogs.exe 所在文件路径](https://upload-images.jianshu.io/upload_images/14623749-9a7e1d01a357ff69.png?imageMogr2/auto-orient/strip%7CimageView2/2/w/1240)

* 解压缩 nssm 压缩包。

![ 以上为解压缩 nssm 之后的状态](https://upload-images.jianshu.io/upload_images/14623749-85a99004000632e8.png?imageMogr2/auto-orient/strip%7CimageView2/2/w/1240)

![将 nssm.exe 文件所在文件绝对路径加入到系统环境变量中](https://upload-images.jianshu.io/upload_images/14623749-eefbd6e6aee7c319.png?imageMogr2/auto-orient/strip%7CimageView2/2/w/1240)

* 以管理员权限运行  install-as-service.bat

![对着 install-as-service.bat 文件鼠标右击，以管理员权限执行](https://upload-images.jianshu.io/upload_images/14623749-0ba62def3fe97198.png?imageMogr2/auto-orient/strip%7CimageView2/2/w/1240)

> 同样也可以使用手动执行命令去执行 install-as-service.bat 文件   
手动执行命令的方法为：在 install-as-service.bat 文件所在文件夹下，随便点击一下空白处，然后按住 shift键，点击鼠标右键，点击 **在此处打开命令窗口** 输入 gogs web 命令，回车即可。

5. 测试
   浏览器访问：127.0.0.1:3000 即可进入配置页面（我只修改了代码仓库存放路径这一项）。（在此页面并不一定非要注册用户，我测试的时候，虽然注册了一个用户，但是最后发现还是需要再重新注册）完成配置后，E:\soft-exe\gogs\custom\conf 目录下会生成一个新的 app.ini 配置文件。

![E:\soft-exe\gogs\custom\conf 路径下生成了 app.ini 配置文件](https://upload-images.jianshu.io/upload_images/14623749-585db4ac286d6008.png?imageMogr2/auto-orient/strip%7CimageView2/2/w/1240)

![打开 app.ini 配置文件，发现里面的内容为在浏览器中输入的内容](https://upload-images.jianshu.io/upload_images/14623749-02ed41e662d49c21.png?imageMogr2/auto-orient/strip%7CimageView2/2/w/1240)

6. 注册用户（系统默认第一个用户为系统管理员）
   再次访问 127.0.0.1:3000 ，点击注册。

![我本地测试的电脑未安装 .net 框架，因此样式乱掉了](https://upload-images.jianshu.io/upload_images/14623749-d80ab87d33fc9f1b.png?imageMogr2/auto-orient/strip%7CimageView2/2/w/1240)

7. 创建测试仓库

![image.png](https://upload-images.jianshu.io/upload_images/14623749-387e54229c30b0c9.png?imageMogr2/auto-orient/strip%7CimageView2/2/w/1240)



