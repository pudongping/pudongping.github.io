---
title: windows10 专业版 64位系统安装docker并使用 laradock 搭建 laravel 环境
author: Alex
top: false
hide: false
cover: false
toc: true
mathjax: false
categories: Docker
tags:
  - Docker
  - Laravel
  - Laradock
abbrlink: f59f8f31
date: 2021-07-15 10:53:45
img:
coverImg:
password:
summary:
---

# windows10 专业版 64 位系统安装 docker 并使用 laradock 搭建 laravel 环境

> [docker官网](https://www.docker.com/)   
[docker官网安装文档](https://docs.docker.com/install/)


## 安装说明
- ### windows 10 系统需要开启 Hyper-V

![官方文档中有写到，必须开启 Hyper-V](https://upload-images.jianshu.io/upload_images/14623749-1c2cfe770c1a9c08.png?imageMogr2/auto-orient/strip%7CimageView2/2/w/1240)

开启方式如下图：
1. 打开控制面板-程序-程序和功能-启用或关闭 windows 功能
   ![启用或关闭 windows 功能](https://upload-images.jianshu.io/upload_images/14623749-8c02d2b070ab1bfd.png?imageMogr2/auto-orient/strip%7CimageView2/2/w/1240)

2. 找到有关  Hyper-V  的项，全部选中
   ![有关 Hyper-V 的选项，全部勾选](https://upload-images.jianshu.io/upload_images/14623749-4a53340ebf3d7ebc.png?imageMogr2/auto-orient/strip%7CimageView2/2/w/1240)

3. 如果发现关于 Hyper-V 的选项无法开启，那么就需要进入 bios 开启虚拟化。开启方法见如下链接：

> [如何查看自己的Win10电脑是否能运行Hyper-V虚拟机](https://www.windows10.pro/how-to-see-if-the-computer-can-run-hyper-v-virtual-machines/)

> 查看 Hyper-V 固件中启用的虚拟化是否开启的步骤：Win + R 输入 “msinfo32 ” 即可看到“系统信息”窗口。   
进入 bios 开启固件虚拟化的方法步骤：进入 bios 设置界面，切换到 Advanced 标签，选中 CPU Configuration 设置 Intel Virtualization Technology 为 Enabled。（不同的主板可能会有不同的设置方法，主要是将 Intel Virtualization Technology 设置为 Enabled 即可）

4. 再次打开【启用或关闭 windows 功能界面】开启 Hyper-V 所有选项。
   ![不管怎样，主要的步骤是开启 Hyper-V 所有选项](https://upload-images.jianshu.io/upload_images/14623749-0aded33529879bf5.png?imageMogr2/auto-orient/strip%7CimageView2/2/w/1240)

- ### 下载 Docker Desktop for Windows desktop app

![官方文档详细步骤](https://upload-images.jianshu.io/upload_images/14623749-f853b71da33b32c7.png?imageMogr2/auto-orient/strip%7CimageView2/2/w/1240)

1. 下载 Docker Desktop for Windows app
> [Docker Desktop Installer.exe 下载安装地址](https://download.docker.com/win/stable/Docker%20for%20Windows%20Installer.exe)

2. 下载完成之后，直接双击安装
   ![安装过程图01](https://upload-images.jianshu.io/upload_images/14623749-901ae0bb45645f92.png?imageMogr2/auto-orient/strip%7CimageView2/2/w/1240)

![安装过程图02](https://upload-images.jianshu.io/upload_images/14623749-922d3317a0152afc.png?imageMogr2/auto-orient/strip%7CimageView2/2/w/1240)

3. 开启 docker
   直接可以通过小娜助手搜 docker 关键词，然后启动 Docker Desktop，不管怎样，主要是找到 Docker Desktop 应用，打开就好。
   ![本图通过小娜助手搜 docker 关键词打开](https://upload-images.jianshu.io/upload_images/14623749-e6b8377c984ee245.png?imageMogr2/auto-orient/strip%7CimageView2/2/w/1240)

4. 查看 docker 开启状态。
   ![桌面右下角会出现 docker 的图标，鼠标移到图标上面会出现 Docker Desktop is running 字样，即为打开](https://upload-images.jianshu.io/upload_images/14623749-54e121b6e4f24813.png?imageMogr2/auto-orient/strip%7CimageView2/2/w/1240)

> 初次安装时，可能会提示你登录 docker 的账号，如果没有 docker 账号的话，可以去 docker 官网注册一下。

5. 测试安装。
   任意位置打开 Windows PowerShell
```
// 查看 docker 版本
docker --version 

docker-compose --version

docker run hello-world
```
![使用 PowerShell 和 cmd是一样的](https://upload-images.jianshu.io/upload_images/14623749-e644b75b3cb1da7b.png?imageMogr2/auto-orient/strip%7CimageView2/2/w/1240)


- ### 使用 laradock 搭建项目
1. 将 laradock 项目代码克隆到本地：
```
git clone https://github.com/Laradock/laradock.git
```

![下载 laradock 项目到本地](https://upload-images.jianshu.io/upload_images/14623749-22cb4a41faae8ed0.png?imageMogr2/auto-orient/strip%7CimageView2/2/w/1240)


2. 进入 `laradock` 目录将 `env-example` 复制一份并命名为 `.env`
```
cp env-example .env
```

![复制配置文件](https://upload-images.jianshu.io/upload_images/14623749-2ba18bf82bcb896a.png?imageMogr2/auto-orient/strip%7CimageView2/2/w/1240)

3. 运行容器 （此时在 laradock 目录下）
```
docker-compose up -d nginx mysql redis workspace
```
如果指定端口已经被占用，运行上述命令会报错，关闭相应的服务再重新运行上述命令即可。

> 注：安装过程中，由于某些资源需要翻墙才能下载，建议安装并启用 VPN 后再执行上述命令。如果出现需要认证的下载资源无权下载，可以通过 Docker ID/密码 登录到 Docker 应用（点击状态栏 Docker 应用小图标就能看到登录菜单），注意这里必须用 Docker ID，不能用注册邮箱。在 Windows 下如果出现目录挂载失败，可以尝试在 Docker 设置中重新设置 Shared Drives。

4. 打开项目的 `.env` 文件并添加如下配置：
```
DB_HOST=mysql
REDIS_HOST=redis
QUEUE_HOST=beanstalkd
```

5. 在和 laradock 同级目录下新建 wwwroot 目录，用于存放代码
   ![新建 wwwroot 目录](https://upload-images.jianshu.io/upload_images/14623749-ec64df8bf3dd4e16.png?imageMogr2/auto-orient/strip%7CimageView2/2/w/1240)

新建 demo 文件夹，并写入 phpinfo(); 到 index.php 作为测试。
![新建测试文件](https://upload-images.jianshu.io/upload_images/14623749-95be36bbd3077799.png?imageMogr2/auto-orient/strip%7CimageView2/2/w/1240)


6. 此时需要再次在 `.env` 文件中修改 `APPLICATION` 配置项（新版本的 laradock 对应的配置项是 `APP_CODE_PATH_HOST`）

```
APPLICATION=../wwwroot/
```
![配置项目路径](https://upload-images.jianshu.io/upload_images/14623749-826d4ee6853f37e5.png?imageMogr2/auto-orient/strip%7CimageView2/2/w/1240)

这样就相当于为 wwwroot 与 Docker 的 /var/www 目录建立了软连接，然后我们修改 nginx 的配置文件，建立映射关系。

```
// 复制一份配置文件 demo.conf
cp ./laradock/nginx/sites/default.conf ./laradock/nginx/sites/demo.conf
```

修改成以下内容

```
server {

    listen 80;
    listen [::]:80;

    server_name demo.test;
    root /var/www/demo;
    index index.php index.html index.htm;

    location / {
         try_files $uri $uri/ /index.php$is_args$args;
    }

    location ~ \.php$ {
        try_files $uri /index.php =404;
        fastcgi_pass php-upstream;
        fastcgi_index index.php;
        fastcgi_buffers 16 16k;
        fastcgi_buffer_size 32k;
        fastcgi_param SCRIPT_FILENAME $document_root$fastcgi_script_name;
        #fixes timeouts
        fastcgi_read_timeout 600;
        include fastcgi_params;
    }

    location ~ /\.ht {
        deny all;
    }

    location /.well-known/acme-challenge/ {
        root /var/www/letsencrypt/;
        log_not_found off;
    }
}
```

重启 Docker 的 Nginx

```
docker-compose up -d nginx 
```

7. 在 hosts 文件中添加 （Windows 下对应文件路径是 `C:\Windows\System32\drivers\etc\hosts`）
```
127.0.0.1 demo.test
```

8. 在浏览器中访问 demo.test
   ![如图所示，则表示 php 环境搭建成功！](https://upload-images.jianshu.io/upload_images/14623749-af86d3cc7ce05b58.png?imageMogr2/auto-orient/strip%7CimageView2/2/w/1240)

9. 安装多项目。比如搭建 laravel 项目
   在 wwwroot 目录下执行 composer 命令 （需要提前在 Windows 环境中安装 composer）
   可以查看我写的这篇文章 [Windows安装composer](https://www.jianshu.com/p/8eea56dbd246)

```
composer create-project laravel/laravel blog --prefer-dist
```

![搭建 laravel 项目 blog](https://upload-images.jianshu.io/upload_images/14623749-0a826d3e36db460a.png?imageMogr2/auto-orient/strip%7CimageView2/2/w/1240)

目录结构如下：
![项目都在 wwwroot 目录下](https://upload-images.jianshu.io/upload_images/14623749-5a4382fac6349a5b.png?imageMogr2/auto-orient/strip%7CimageView2/2/w/1240)

10. 添加 nginx 配置文件追加 hosts 配置
    ![重新复制一份 default.conf 配置文件，并作相应的配置修改](https://upload-images.jianshu.io/upload_images/14623749-cea8ec0325dba714.png?imageMogr2/auto-orient/strip%7CimageView2/2/w/1240)

11. 重启 docker 中的 nginx
```
docker-compose up -d nginx
```

12. 彩蛋

- 进入Workspace 容器, 执行比如(Artisan, Composer, PHPUnit, Gulp, ...)等命令

```
docker-compose exec workspace bash 或者 docker exec -it laradock_workspace_1 bash
```

- 列出正在运行中的容器

```
docker ps
```

- 关闭所有正在运行的容器

```
docker-compose stop
```

- 进入 mysql 容器

```
docker-compose exec mysql bash
```

- 退出容器

```
exit
```
