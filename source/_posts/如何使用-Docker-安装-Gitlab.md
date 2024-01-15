---
title: 如何使用 Docker 安装 Gitlab
author: Alex
top: false
hide: false
cover: false
toc: true
mathjax: false
categories: Git
tags:
  - Git
  - Docker
abbrlink: bd835025
date: 2024-01-15 11:51:08
img:
coverImg:
password:
summary:
---


> [Gitlab 官方 docker](https://hub.docker.com/r/gitlab/gitlab-ce)
[官方 docker 安装文档](https://docs.gitlab.com/ee/install/docker.html)

**内存小于 4GB 的机器最好不要安装！GitLab 超级占内存！**

```bash

mkdir -p /srv/gitlab

sudo docker run -d \
--hostname {your_github_hostname or your_ip} \
--publish 443:443 --publish 80:80 --publish 22:22 \
--name gitlab \
# docker 重启后，容器也会重启
--restart always \
--privileged=true \
# gitlab 配置
--volume /srv/gitlab/config:/etc/gitlab \
# gitlab 日志
--volume /srv/gitlab/logs:/var/log/gitlab \
# gitlab 应用数据
--volume /srv/gitlab/data:/var/opt/gitlab \
# 设置共享内存大小
--shm-size 256m \
gitlab/gitlab-ce:latest

```

查看容器启动日志

```bash
docker logs -f gitlab
```

默认用户名为 `root` ，初始密码可以执行以下命令获取。**24小时内要修改默认密码**

```bash
docker exec -it gitlab grep 'Password:' /etc/gitlab/initial_root_password
```

修改配置文件

> 因为我们将 docker 中的配置映射到宿主机的 `/srv/gitlab/config` 目录下，因此可直接在该目录中编辑配置文件

```bash
# 查看启用的配置
cat /srv/gitlab/config/gitlab.rb | grep -v '^#' | grep -v '^$'
# 编辑配置文件
vim /srv/gitlab/config/gitlab.rb
```

可能需要修改的内容

```bash
# 设置时区
gitlab_rails['time_zone'] = 'Asia/Shanghai'

# 设置访问地址，没有域名则可以直接设置宿主机的 IP
external_url 'http://127.0.0.1'

# 设置 ssh 访问 host，没有域名则可以直接设置宿主机的 IP
gitlab_rails['gitlab_ssh_host'] = '127.0.0.1'
# 设置 ssh 端口
gitlab_rails['gitlab_shell_ssh_port'] = 2222

# 禁用内建的 nginx
nginx['enable'] = false

# 取消掉这里的注释
# 太占用内存了
puma['worker_processes'] = 2
postgresql['shared_buffers'] = "256MB"
```