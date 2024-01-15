---
title: 如何运用 systemd 管理 Linux 中的服务
author: Alex
top: false
hide: false
cover: false
toc: true
mathjax: false
categories: Linux
tags:
  - Linux
abbrlink: ad8a576e
date: 2024-01-15 10:05:49
img:
coverImg:
password:
summary:
---

## 什么是 systemd

`systemd` 是 Linux 系统中用于系统和服务管理的一种工具，它是一个初始化系统，用于启动和管理整个系统。从 Linux kernel 完成自我启动后，systemd 是第一个被启动的用户级进程，它的进程 ID 为 1。它的主要职责包括启动系统服务，管理系统资源，以及处理系统关机、重启等操作。

## 系统服务和 systemd

在 Linux 中, 一个服务通常被定义为一个常驻的长时运行程序，比如 web 服务器, 数据库服务器等。在使用 systemd 的系统中，每个服务都被定义为一个独立的 unit，这些 unit 会被 systemd 管理和监视。所有被 systemd 管理的服务的配置文件都在 `/etc/systemd/system` 和 `/usr/lib/systemd/system` 目录下。

## 管理 systemd 服务

你可以通过 `systemctl` 命令来管理被 systemd 控制的服务。以下是一些常用的 `systemctl` 发令：

- 查看所有被 systemd 控制的服务: `systemctl list-units --type=service`
- 启动一个服务: `systemctl start [service-name]`
- 停止一个服务: `systemctl stop [service-name]`
- 重启一个服务: `systemctl restart [service-name]`
- 查看某个服务的状态: `systemctl status [service-name]`
- 使服务在启动时自启动: `systemctl enable [service-name]`
- 禁止服务在启动时自启动: `systemctl disable [service-name]`

## 创建 systemd 服务

创建自定义的 systemd 服务也很简单。主要是创建一个以 `.service` 结尾的文件，这个文件中定义了服务的执行程序，启动方式等信息。以下是一个简单的例子:

Systemd 的 Service 配置在 `/etc/systemd/system/` 目录中，可以创建一个 `echo.service` 文件，实际项目应当改为对应的名称。编辑此文件，添加下列内容：

```bash
[Unit]
Description=Echo Http Server
After=network.target
After=syslog.target

[Service]
Type=simple
LimitNOFILE=65535
ExecStart=/usr/bin/php /opt/servers/echo/server.php
ExecReload=/bin/kill -USR1 $MAINPID
Restart=always

[Install]
WantedBy=multi-user.target graphical.target

```

- After 指令约定了启动的顺序，必须在 network 和 syslog 启动后才启动 echo 服务
- Service 中填写了应用程序的路径信息，请修改为实际项目对应的路径
- Restart=always 表示如果进程挂掉会自动拉起
- WantedBy 约定了在哪些环境下启动，multi-user.target graphical.target 表示在图形界面和命令行环境都会启动

编写完成后需要 reload 守护进程使其生效

```bash

sudo systemctl --system daemon-reload

```

管理服务示例

```bash

# 启动服务
sudo systemctl start echo.service
# reload 服务
sudo systemctl reload echo.service
# 关闭服务
sudo systemctl stop echo.service
# 查看服务状态
sudo systemctl status echo.service

# 查看所有的启动项
sudo systemctl list-unit-files
# 禁用开机启动
sudo systemctl disable echo.service

```

`systemd` 是 Linux 发行版中的标准工具，用于管理系统启动和服务。其 `systemctl` 子命令也为管理服务提供了强大而灵活的工具。理解和掌握 `systemctl` 的使用必将使你能更好的理解和控制你的 Linux 系统。