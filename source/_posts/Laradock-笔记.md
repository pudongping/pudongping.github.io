---
title: Laradock 笔记
author: Alex
top: false
hide: false
cover: false
toc: true
mathjax: false
categories: Docker
tags:
  - Docker
  - Linux
  - CentOS
  - Laradock
abbrlink: 17c35873
date: 2021-07-14 09:45:56
img:
coverImg:
password:
summary:
---

# laradock 笔记

- [官方文档](http://laradock.io/documentation/)
- [中文文档](https://laradock.linganmin.cn/zh/getting-started/)


## 安装步骤

1. 首先将 Laradock 项目代码克隆到本地：

```bash
git clone https://github.com/Laradock/laradock.git
```

2. 进入 laradock 目录将 env-example 重命名为 .env：

```bash
cp env-example .env
```

然后在 .env 中修改镜像构建过程中 Linux 软件源为国内镜像以免镜像构建过程中出现网络超时问题：

```bash
CHANGE_SOURCE
# If you need to change the sources (i.e. to China), set CHANGE_SOURCE to true
CHANGE_SOURCE=true
# Set CHANGE_SOURCE and UBUNTU_SOURCE option if you want to change the Ubuntu system sources.list file.
UBUNTU_SOURCE=aliyun
```

3. 构建镜像 & 启动容器：

```bash
docker-compose up -d nginx mysql redis beanstalkd
```

nginx 镜像构建在 php-fpm 之上，php-fpm 构建在 workspace 之上，所以启动 nginx 会先启动 workspace 和 php-fpm。

如果指定端口已经被占用，运行上述命令会报错，关闭相应的服务再重新运行上述命令即可。

如果在 Windows 系统中上述指令构建镜像过程中报错：

```bash
/tmp/sources.sh: not found
```

可参考这个 issue 解决：https://github.com/laradock/laradock/issues/2450。

4. 打开 Laravel 项目的 .env 文件并添加如下配置：

```bash
DB_HOST=mysql
# 这里填写容器的名称（比如：laradock_redis_1）或者容器的 ip 地址（比如：172.28.0.5）也可以
REDIS_HOST=redis
QUEUE_HOST=beanstalkd
```

## 其他命令

- 构建镜像 & 启动容器

```shell
docker-compose up -d nginx mysql redis beanstalkd
```

- [重新构建容器](https://laradock.linganmin.cn/zh/documentation/#%E6%9E%84%E5%BB%BA%E6%88%96%E9%87%8D%E6%9E%84%E5%AE%B9%E5%99%A8)

```shell
# 比如：重新构建 mysql 容器
docker-compose build mysql
```

- 增加一个项目之后,重启 Docker 的 nginx

```shell
docker-compose up -d nginx
```

- 列出容器

```shell
# 列出正在运行的容器
docker ps

# 如果你只想看当前这个项目的容器，你也可以执行下面这个命令
docker-compose ps
```

- 重启当前这个项目中的所有容器（如果你不想一个一个的开启每一个容器，可以方便的执行这条命令）

```shell
docker-compose restart
```

- 关闭容器

```shell
# 关闭所有正在运行的容器
docker-compose stop

# 停止单个容器
docker-compose stop {container-name}
```

- 进入容器

```shell
# 使用下面的命令进入任意容器
docker-compose exec {container-name} bash

# 进入 mysql 容器
docker-compose exec mysql bash

# 进入 mysql 并在 mysql 容器中使用命令提示符
docker-compose exec mysql mysql -u homestead -psecret

# 进入 workspace 容器,执行比如(Artisan, Composer, PHPUnit, Gulp, ...)等命令
docker-compose exec workspace bash
```

- 删除所有现有容器

```shell
docker-compose down
```

- 查看日志文件  
  NGINX 日志文件存储在logs/nginx目录中  
  但是要查看其他容器（Mysql，PHP-FPM,...）的日志，可以运行以下命令

```shell
docker-compose logs {container-name}

docker-compose logs -f {container-name}
```

## 在 laradock 中安装 swoole

在本地安装的话，以 Laradock 为例，需要在 laradock 目录下的 .env 中将下面两行配置值设置为 true：

```bash
WORKSPACE_INSTALL_SWOOLE=true
PHP_FPM_INSTALL_SWOOLE=true
```

然后运行 `docker-compose build php-fpm workspace` 重新构建 Docker 容器，构建完成后重启这两个容器，进入 workspace 容器，运行 `php -m` 查看 Swoole 是否安装成功，如果扩展列表包含 swoole 则表示安装成功。


## 其他需要注意事项

### 在 `laravel` 框架中，如果配置不生效，请注意清理下 `laravel` 的缓存

```bash
php artisan config:clear
```

### 检查下是否开启了 5200 端口

> 安装 netstat 命令，查看端口  
apt-get update  
apt-get install net-tools  
netstat -ntlp

```bash
netstat -ant | grep 5200
```

### 端口映射开启方式

1. 进入 `laradock/docker-compose.yml`

在 `workspace` 下的 `ports` 中新增

```bash
  ports:
    - "${WORKSPACE_SSH_PORT}:22"
    - "${WORKSPACE_VUE_CLI_SERVE_HOST_PORT}:8080"
    - "${WORKSPACE_VUE_CLI_UI_HOST_PORT}:8000"
    - "${WORKSPACE_PORT}:5200"     /**这一行为新增的行，也可以直接在这里加 5200:5200 这样加了之后，就不再需要在 .env 中设置变量了**/
```

2. 进入 `.env` 在 `WORKSPACE` 下最后一行增加

```bash

WORKSPACE_AST_VERSION=1.0.3
WORKSPACE_VUE_CLI_SERVE_HOST_PORT=8080
WORKSPACE_VUE_CLI_UI_HOST_PORT=8001
WORKSPACE_INSTALL_GIT_PROMPT=false
WORKSPACE_PORT=5200  /**这一行为新增的行，其实就是设置步骤 1 中的变量**/

```

> 有些博客说还需要在 laradock/workspace/Dockerfile 文件的最后添加一行，申明开放端口： EXPOSE 5200，这里，我并没有做这一步，同样也成功了，如果你没有成功，你加上去之后再试试看吧

3. 强制重新创建 workspace 容器

```bash
docker-compose up -d --force-recreate workspace
```

4. 重启 `docker-compose`

```bash
docker-compose restart 

docker ps

```

5. 测试端口是否开通成功

```bash
telnet 127.0.0.1 5200

# 或者直接查看容器的端口列表中是否含有你所需要开通的端口
docker port {container-name}
# 比如，如下
docker port laradock_workspace_1
```
