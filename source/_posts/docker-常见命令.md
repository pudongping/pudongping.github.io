---
title: docker 常见命令
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
abbrlink: 59b529b7
date: 2021-07-13 09:14:20
img:
coverImg:
password:
summary:
---

# docker 常见命令

- 查看版本

```bash
docker -v
```

- 查看详细版本

```bash
docker version
```

- 查看 docker 基本信息

```bash
docker info
```

---

## 镜像相关的命令

镜像仓库地址：https://hub.docker.com

- 查看本地已经安装的镜像

```bash
docker images
```

- 搜索指定镜像

```bash
docker search <image-name>

# 比如搜索 centos 镜像
docker search centos
```

- 拉取镜像

```bash
docker pull <image-name>

# 比如拉取 centos 镜像（将会拉取最新版本的镜像，即 latest 版本）
docker pull centos

# 拉取指定版本的镜像
docker pull ubuntu:16.04
```

- 删除镜像

```bash
docker rmi <image-name>

# 比如删除 centos 镜像
docker rmi centos
```

- 删除所有的镜像

```bash
docker rmi $(docker images -q)

# 查看所有镜像的镜像 id
docker images -q
```

- 更新镜像

```bash
# 更新镜像前，需要使用镜像创建一个容器
docker run -it ubuntu:15.10 bash

# 在运行的容器内部使用 apt-get update 更新

# 更新完毕之后，输入 exit 命令退出容器

# 使用 docker commit 来提交容器副本
docker commit -m="has update" -a="alex" e218edb10161 alex/ubuntu:v2
# -m 表示提交的描述信息
# -a 表示提交的镜像作者
# e218edb10161 表示容器的 id
# alex/ubuntu:v2 表示指定要创建的目标镜像名

# 查看新的镜像
docker images

# 使用新镜像 alex/ubuntu:v2 来启动一个容器
docker run -it alex/ubuntu:v2 bash
```

- 构建镜像

1. vim  ~/glory/codes/book/demo/Dockerfile 填充以下内容，构建一个 centos 8 系统

```sh

# 指定使用哪个镜像源
FROM centos:8

# 如果写了 3 次 RUN 那么将会在 docker 上新建 3 层，会导致镜像膨胀过大，下面这种方式只会创建 1 层镜像
RUN /bin/echo 'root:123456' | chpasswd; \
useradd alex; \
/bin/echo 'alex:123456' | chpasswd; \
/bin/echo -e "LANG=\"en_US.UTF-8\"" > /etc/default/local

EXPOSE 22
EXPOSE 80
CMD /usr/sbin/sshd -D

```

2. 开始构建镜像，注意不要在 `~/glory/codes/book/demo` 目录下放无用的文件，因为会打包所有该目录下的文件然后发送给 docker 引擎，如果文件过多会造成 build 过程缓慢

```sh
# -t 表示指定要创建的目标镜像名
# ~/glory/codes/book/demo 表示 Dockerfile 文件所在的目录
docker build -t alex/centos:8.0 ~/glory/codes/book/demo

# 查看已经构建好的镜像信息
docker images

# 使用新的镜像来创建新容器
docker run -it alex/centos:8.0 bash
```

- 设置镜像标签

```bash
# 先查看镜像
$ docker images                                                                   

REPOSITORY            TAG                       IMAGE ID       CREATED          SIZE
alex/centos           8.0                       594ab4747ed4   14 minutes ago   210MB

# 设置镜像标签
$ docker tag 594ab4747ed4 alex1/centos1:8.1.1

# 再次查看镜像信息会多一个标签
$ docker images

REPOSITORY            TAG                       IMAGE ID       CREATED          SIZE
alex/centos           8.0                       594ab4747ed4   14 minutes ago   210MB
alex1/centos1         8.1.1                     594ab4747ed4   14 minutes ago   210MB
```


---

## 容器相关的命令

- 查看容器的系统版本信息

```bash
# 进入容器后执行
cat /proc/version
# 比如会输出以下内容
Linux version 4.19.121-linuxkit (root@18b3f92ade35) (gcc version 9.2.0 (Alpine 9.2.0)) #1 SMP Thu Jan 21 15:36:34 UTC 2021
```

- 查看所有的容器

```bash
docker ps -a

# 或者使用以下命令，是一样的效果
docker container ls -a

# 查看最后一次创建的容器
docker ps -l
```

- 查看所有已经运行的容器

```bash
docker ps
```

- 查看容器端口映射

```bash
docker port <container-name or container-id>
```

- 查看 docker 底层信息（比如：查看指定容器的 ip 地址）

```bash
# 查看 docker 容器的配置和状态信息
docker inspect <container-name or container-id> 

# 查看容器的 ip 地址
docker inspect <container-name or container-id> | grep IPAddress

# 比如查看容器 id 为 66204be9fe65 的容器所对应的 ip 地址
docker inspect 66204be9fe65 | grep IPAddress

# 比如查看容器名称为 alex 所对应的 ip 地址
docker inspect alex | grep IPAddress
```

- 创建容器并把镜像恢复到容器当中，且启动容器

```bash
docker run [-i][-t][-v][--name][-d][-p]

# -i 表示 interactive 交互式
# -t 表示得到一个 terminal
# --name 表示修改容器名称
# -d 表示以守护进程的方式运行（默认不会进入容器，想要进入容器则需要使用 docker exec 命令）
# -p 表示 **指定** 映射端口
# -P （大写的字母 p ） 表示 **随机** 映射端口
# /bin/bash 和 bash 等效
docker run -it <image-name> /bin/bash

# 比如创建一个新容器并且进入 ubuntu:16.04 镜像
docker run -it ubuntu:16.04 bash
# 或者
docker run -it ubuntu:16.04 /bin/bash
# 或者
docker run -it ubuntu:16.04

# 以 centos 镜像创建一个新容器，并修改新容器名称为 alex-container
docker run -it --name alex-container centos bash  

# 以守护进程的方式运行 （后台运行）
docker run -d --name alex-container centos  
# 或者
docker run -di --name alex-container centos

# 指定容器绑定的网络地址，这样我们就可以通过访问 127.0.0.1:5001 来访问容器的 5000 端口（默认绑定的都是 tcp 端口）
docker run -d -p 127.0.0.1:5001:5000 centos:8.0 bash
# 如果需要绑定 udp 端口，则
# （还可以进入容器就直接执行 python app.py 命令）
docker run -d -p 127.0.0.1:5001:5000/udp centos:8.0 python app.py

# 比如，安装 hyperf 镜像并启动容器
# 如果 docker 启动时开启了 selinux-enabled 选项，容器内访问宿主机资源就会受限，所以启动容器时可以增加 --privileged -u root 选项
docker run --name hyperf \
-v /workspace/skeleton:/data/project \
-p 9501:9501 -it \
--privileged -u root \
--entrypoint /bin/sh \
hyperf/hyperf:7.4-alpine-v3.11-swoole

# 如果需要开通多个端口时，可以参考
docker run --name api_dfo_hyperf_ws \
-v /Users/pudongping/glory/codes/dfo/api_dfo_hyperf:/api_dfo_hyperf \
-p 9502:9502 \
-p 9503:9503 \
-p 9504:9504 \
-p 9505:9505 -it \
--entrypoint /bin/sh \
alex/alex_api_dfo:v1.0

```

- 启动容器

```bash
docker start <container-name or container-id>

# 比如启动容器名称为 redis-alex 的容器
docker start redis-alex
# 比如启动容器 id 为 c8c0c770ac5b 的容器
docker start c8c0c770ac5b
```

- 直接进入已经创建的容器（不会启动容器）

```bash
docker start -i <container-name or container-id>

# 比如进入容器 id 为 66204be9fe65 的容器
docker start -i 66204be9fe65
# 比如进入容器名称为 alex 的容器
docker start -i alex
```

- 重启容器

```bash
docker restart <container-name or container-id>

# 比如重启容器名称为 redis-alex 的容器
docker restart redis-alex
# 比如重启容器 id 为 c8c0c770ac5b 的容器
docker restart c8c0c770ac5b
```

- 进入已经运行中的容器

```bash
docker exec -it <container-name or container-id> bash

# 比如进入容器名称为 redis-alex 的容器
docker exec -it redis-alex bash
# 比如进入容器 id 为 c8c0c770ac5b 的容器
docker exec -it c8c0c770ac5b bash

# 进入容器之后执行 shell 命令或者执行 shell 脚本
docker exec -it  <container-name or container-id> /bin/sh -c "while true; do echo hello world; sleep 1; done"
# 比如进入容器 id 为 c8c0c770ac5b 的容器，并且进入容器后执行 `bash /portal_api_dfo_hyperf/server.sh restart` 脚本
docker exec -it c8c0c770ac5b /bin/sh -c "bash /portal_api_dfo_hyperf/server.sh restart"
```

- 停止容器

```bash
docker stop <container-name or container-id>

# 比如停止容器名称为 redis-alex 的容器
docker stop redis-alex
# 比如停止容器 id 为 c8c0c770ac5b 的容器
docker stop c8c0c770ac5b
```

- 退出容器

```bash
exit
```

- 删除容器

```bash
docker rm <container-name or container-id>
# 也可以加入 -f 参数，强制移除正在运行中的容器
docker rm -f 1e560fca3906
# 清理掉所有处于终止状态的容器
docker container prune
```

- 修改容器名称

```bash
docker rename <container-name or container-id> <new-container-name>

# 比如将容器 redis-alex 改名为 redis-tt
docker rename redis-alex redis-tt
```

- 查看容器的标准输出

```bash
docker logs <container-name or container-id>

# 比如查看容器 id 为 c8c0c770ac5b 的容器标准输出内容
docker logs c8c0c770ac5b
# 也可以加入 -f 参数，像使用 tail -f 一样来输出容器内部的标准输出
docker logs -f c8c0c770ac5b
```


## 容器与宿主机之间的文件或者目录拷贝

- 从宿主机拷贝文件到容器中

```bash
docker cp <local-directory-or-file> <container-name>:<container-directory-or-file>

# 比如将宿主机中的 /home/alex/test.txt 文件拷贝到 centos1 容器中的 /test.txt
docker cp /home/alex/test.txt centos1:/test.txt
```

- 从容器拷贝到宿主机中

```bash
docker cp <container-name>:<container-directory-or-file> <local-directory>

# 比如将 centos1 容器中的 /test 目录拷贝到宿主机的 /home/alex 目录下
docker cp centos1:/test /home/alex
```


## 目录挂载（创建容器的时候就需要进行目录挂载）

```bash
docker run -di -v <local-directory>:<container-directory> <image-name>

# 在 windows 下挂载（注意路径的书写方式）
# 比如以 centos 镜像创建一个容器，并将本地 D 盘中的 alex 目录，挂载到容器中的 /usr/local/demo 目录
docker run -di -v d:\alex:/usr/local/demo centos

# 在 linux 下挂载
# 比如以 centos 镜像创建一个容器，并将本地的 /home/alex/alex 目录，挂载到容器中的 /usr/local/demo 目录
docker run -di -v /home/alex/alex:/usr/local/demo centos
```


## 导出和导入容器

- 导出容器快照

```bash
docker export <container-id> > <your-backup-name.tar>

# 比如将容器 id 为 7691a814370e 的容器导出快照为 alex.tar
docker export 7691a814370e > alex.tar
```

- 导入容器快照

```bash
cat <your-backup-name.tar> | docker import - <image-author-name>/<your-new-image-name>:<your-new-image-version>

# 比如将容器快照文件 alex.tar 导入到 alex-demo 镜像并定义 alex-demo 镜像的作者为 alex，版本号为 v1.0
cat alex.tar | docker import - alex/alex-demo:v1.0
# 此外，也可以通过指定 url 或者某个目录来导入
docker import http://example.com/example-image.tgz example/image-repo:v1.0
```
