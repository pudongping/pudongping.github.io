---
title: 一起来了解下Docker-compose
author: Alex
top: false
hide: false
cover: false
toc: true
mathjax: false
categories: Docker
tags:
  - Linux
  - CentOS
  - Docker
abbrlink: 2d05cf8b
date: 2024-02-26 22:43:52
img:
coverImg:
password:
summary:
---

Docker Compose是一个用于定义和运行多容器Docker应用程序的工具。通过Compose，你可以使用 YAML 文件来配置应用程序的服务。然后，只需一个简单的命令，就可以创建并启动配置中的所有服务。这使得部署多容器应用变得简单快捷。

### 核心特性

- **多容器编排**：Docker Compose 允许你在单个文件中定义一组相关联的容器作为项目。这些容器通过网络连接在一起，形成一个整体服务。

- **易于配置**：使用 YAML 文件定义服务配置，使其易于编写、读取和维护。

- **一键部署**：只需一个`docker-compose up`命令，就可以启动和运行整个应用环境。

- **环境隔离**：Docker Compose 使用项目名称来隔离不同的环境，您可以在同一台机器上运行多个环境，而不会产生冲突。

- **可重复性**：由于服务定义在文件中，因此可以在不同环境中重复部署相同的应用配置，确保环境之间的一致性。

### 使用场景

- **开发环境**：在开发过程中，使用 Docker Compose 来定义和运行应用的所有依赖服务，例如数据库、缓存等，可以快速搭建和拆解开发环境。

- **自动化测试**：通过定义包含应用及其依赖服务的 Compose 文件，可以轻松地在 CI/CD 管道中进行自动化测试。

- **小规模生产部署**：对于小型项目或初期阶段的产品，Docker Compose 提供了一种简单的方式来部署和管理多容器应用。

### 工作原理

1. **定义服务**：在`docker-compose.yml`文件中定义应用所需的服务，包括构建镜像的配置、容器运行时配置（如端口映射、卷挂载）等。

2. **启动服务**：运行`docker-compose up`命令来启动并运行定义的所有服务。Compose 会按照依赖关系顺序启动服务，并确保所需的网络和卷已正确设置。

3. **管理服务**：使用`docker-compose`命令可以管理服务的生命周期，如启动、停止、重建等。

### 开始使用

要开始使用 Docker Compose，您需要：

1. **安装Docker**：确保您的系统上安装了 Docker。

2. **安装Docker Compose**：在大多数情况下，Docker Compose 可以与 Docker 一起安装。

3. **编写`docker-compose.yml`文件**：定义您的多容器应用。

4. **运行`docker-compose up`命令**：启动并运行您的应用。

介于篇幅有限，今天暂时介绍 docker-compose 的安装方式，后面有时间了再讲具体的使用。

### 安装 docker-compose

```shell

# 使用源码安装（推荐方式）
[最新发行的版本地址](https://github.com/docker/compose/releases)
# 安装 1.27.4 版本的 docker-compose （下载了源码，并改名为 docker-compose）
sudo curl -L https://github.com/docker/compose/releases/download/1.27.4/docker-compose-`uname -s`-`uname -m` -o /usr/local/bin/docker-compose

# 增加执行权限
sudo chmod u+x /usr/local/bin/docker-compose

# 建立软连接到 /usr/bin 目录
sudo ln -s /usr/local/bin/docker-compose /usr/bin/docker-compose

# 查看 docker-compose 版本
docker-compose version

---------------------------------

# 如果是在 centos 系统中
# 安装扩展源
sudo yum -y install epel-release
# 安装 python-pip 模块
sudo yum -y install python-pip

# 如果是在 ubuntu 系统中
# 安装 python-pip 模块
apt install python-pip

# 使用 pip 安装 docker-compose
pip install docker-compose


# 删除掉旧的环境变量 docker-compose
sudo rm /usr/bin/docker-compose
sudo rm /usr/local/bin/docker-compose
# 或者使用 pip 命令卸载 docker-compose
pip uninstall docker-compose
```

### 常见命令

```docker

# 启动容器
docker-compose up -d

# 查看容器列表
docker-compose ps

# 查看日志
docker-compose logs

# 关闭容器
docker-compose stop

# 启动容器
docker-compose start

# 重启容器
docker-compose restart

# 关闭并删除容器
docker-compose down

```
