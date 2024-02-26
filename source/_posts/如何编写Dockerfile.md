---
title: 如何编写Dockerfile
author: Alex
top: false
hide: false
cover: false
toc: true
mathjax: false
categories: Docker
tags:
  - Docker
  - Dockerfile
abbrlink: 1a6d22ba
date: 2024-02-26 23:45:27
img:
coverImg:
password:
summary:
---


# 编写 dockerfile

Dockerfile 文件是用于定义 Docker 镜像生成流程的配置文件，文件内容是一条条指令，每一条指令构建一层，因此每一条指令的内容，就是描述该层应当如何构建；这些指令应用于基础镜像并最终创建一个新的镜像
你可以认为用于快速创建自定义的 Docker 镜像。


部署一个简单的 gin 项目，在项目根目录下写 Dockerfile

```docker

# FROM 指定基础镜像（必须有的指令，并且必须是第一条指令）
FROM golang:latest

ENV GOPROXY https://goproxy.cn,direct

# 将工作目录设置为 $GOPATH/src/github.com/Alex/go-gin-example
# 使用 WORKDIR 指令可以来指定工作目录（或者称为当前目录），
# 以后各层的当前目录就被改为指定的目录，如果目录不存在，WORKDIR 会帮你建立目录
WORKDIR $GOPATH/src/github.com/Alex/go-gin-example

# 将当前上下文目录的内容复制到 $GOPATH/src/github.com/Alex/go-gin-example
# COPY 源路径 目标路径
COPY . $GOPATH/src/github.com/Alex/go-gin-example
RUN go build .

# EXPOSE 指令是声明运行时容器提供服务端口，这只是一个声明，在运行时并不会因为这个声明应用就会开启这个端口的服务
# 但是在 Dockerfile 中写入这样的声明有两个好处：
# 1. 帮助镜像使用者理解这个镜像服务的守护端口，以方便配置映射
# 2. 运行时使用随机端口映射时，也就是 docker run -P 时，会自动随机映射 EXPOSE 的端口
EXPOSE 8000

# 将容器启动程序设置为 `./go-gin-example`
ENTRYPOINT ["./go-gin-example"]

```

然后在项目根目录下执行以下命令

```shell

# -t 指定镜像名称为 blog-service:v1.0.0
# . 构建内容为当前上下文目录
docker build -t blog-service:v1.0.0 .

```

查看当前构建的镜像是否存在

```shell
docker images
```

创建并运行一个新容器

```shell
docker run -p 8000:8000 --name my-blog-service blog-service:v1.0.0
```

如果想要两个容器之间关联起来，互相访问的时候可以在容器内直接使用其关联的容器别名进行访问，而不是通过 IP，比如此时初始化的时候将 `my-blog-service` 容器和 `mysql` 容器进行关联

```
docker run --link mysql:mysql -p 8000:8000 --name my-blog-service blog-service:v1.0.0
```
