---
title: 使用 Makefile 管理和部署 Go 项目
author: Alex
top: false
hide: false
cover: false
toc: true
mathjax: false
categories: Go
tags:
  - Go
  - Golang
  - Makefile
abbrlink: 926b9d4b
date: 2024-07-05 16:13:17
img:
coverImg:
password:
summary:
---

在软件开发的世界里，自动化是提升效率的关键。`Makefile` 尽管是一个历史悠久的工具，但它在任务管理和自动化方面的能力依然不可小觑。

对于 `Go` 语言项目来说，利用 `Makefile` 来管理和自动化构建、部署过程能极大地简化开发流程。本文将引导你如何使用 `Makefile` 在本地开发 `Go` 项目后，将其更新到服务器上。

## 为什么使用 Makefile？

`Makefile` 提供了一个中心化的脚本集合，用于自动化执行各种任务，如编译源代码、打包软件、清理临时文件、部署到服务器等。使用 `Makefile` 可以让复杂的任务序列化、一键化，从而减少人为错误，提升工作效率。

## Makefile 基础

Makefile 是一个特殊格式的文件，它被 make 工具使用来管理和自动化软件的构建过程。每个 Makefile 包含一系列的规则和依赖，make 根据这些规则来执行任务。

## 创建 Makefile

首先，在项目根目录下创建 `Makefile` 文件：

```bash
vim Makefile
```

接着，我们定义一系列的任务来管理和部署我们的 Go 项目。

### 任务脚本解析

以下是 `Makefile` 的内容示例：

```bash

# 预定义变量
REMOTE=127.0.0.1
APPNAME=alex-blog

# 声明 .PHONY 目标
.PHONY: deploy-dev

# deploy-dev 任务
deploy-dev:
    @echo "\n--- 开始构建可执行文件 ---"
    # 设置目标操作系统为 linux，架构为 amd64，并构建项目
    GOOS=linux GOARCH=amd64 go build -ldflags="-s -w" -v -o tmp/$(APPNAME)_tmp

    @echo "\n--- 上传可执行文件 ---"
    # 将构建的文件上传到服务器
    scp tmp/$(APPNAME)_tmp root@$(REMOTE):/data/www/blog.com/

    @echo "\n--- 停止服务 ---"
    # 使用 SSH 执行命令，停止服务
    ssh root@$(REMOTE) "supervisorctl stop $(APPNAME)"

    @echo "\n--- 替换新文件 ---"
    # 使用 SSH 执行一系列命令，更新应用程序
    ssh root@$(REMOTE) "cd /data/www/blog.com/ \
                            && rm $(APPNAME) \
                            && mv $(APPNAME)_tmp $(APPNAME) \
                            && chown www-data:www-data $(APPNAME)"

    @echo "\n--- 开始服务 ---"
    # 使用 SSH 执行命令，启动服务
    ssh root@$(REMOTE) "supervisorctl start $(APPNAME)"

    @echo "\n--- 部署完毕 ---\n"
```

### `.PHONY: <任务名称>`

`.PHONY` 用于声明一个目标是“伪目标”，而非文件名。这意呤着即使在当前目录下存在与任务同名的文件，执行 `make <任务名称>` 时，仍会执行该任务。这个声明可以避免由于存在同名文件而导致的任务不被执行。

## 执行任务

要运行上面定义的 `deploy-dev` 任务，只需要在项目根目录下运行以下命令：

```bash
make deploy-dev
```

这条命令会依次执行构建可执行文件、上传至服务器、停止服务、替换新文件并重新启动服务等一系列操作，极大地简化了手动部署的过程。

## 总结

通过使用 `Makefile` 管理和自动化 `Go` 项目的构建和部署过程，我们可以节省大量时间，避免在重复性操作中出错。

本文介绍的 `Makefile` 示例展示了如何利用这种强大工具简化开发工作，但 `Makefile` 的潜力远不止于此。随着你对 `Makefile` 更深入的理解和掌握，你将能够创建更加复杂和强大的自动化脚本，使你的开发流程更加高效、专业。