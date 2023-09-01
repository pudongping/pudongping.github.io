---
title: 使用 systemd 管理服务
author: Alex
top: false
hide: false
cover: false
toc: true
mathjax: false
abbrlink: 9a5708c9
date: 2023-09-01 15:56:36
img:
coverImg:
password:
summary:
categories: Linux
tags:
  - Linux
  - systemd
---

# 使用 systemd 管理服务

## 编写 Service 脚本

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

## 管理服务

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