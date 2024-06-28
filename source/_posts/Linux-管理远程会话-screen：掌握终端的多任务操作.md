---
title: Linux 管理远程会话 screen：掌握终端的多任务操作
author: Alex
top: false
hide: false
cover: false
toc: true
mathjax: false
categories: Linux
tags:
  - Linux
  - screen
abbrlink: 1330f742
date: 2024-06-28 23:50:08
img:
coverImg:
password:
summary:
---

在日常开发和服务器管理工作中，特别是当我们通过 SSH 连接到远程服务器时，通常需要同时执行多个任务。

Linux 的 `screen` 命令为此提供了一种简便的解决方案。`screen` 允许用户创建多个会话窗口，并在它们之间自由切换，即便与远程服务器的连接断开，这些会话仍然会在后台运行。

接下来，我们将使用简单易懂的语言，探索如何通过 `screen` 提高我们的工作效率。

## 初识 Screen

### 创建新的会话窗口

如果你想运行一个长时间执行的命令（比如备份操作），并不希望因为网络问题而导致命令中断，可以创建一个新的 `screen` 会话窗口。

```bash
screen -S backup
```

这里 `-S` 参数后面跟着的是我们给这个会话窗口的名字，这里名字是 `backup`。

### 查看当前所有会话窗口

如果想要查看当前所有的 `screen` 会话窗口，可以使用下面的命令：

```bash
screen -ls
```

执行这个命令后，你会看到类似于 `session_id` 的信息，其中包含了会话的名称和状态信息，帮助你识别和管理多个会话。

### 退出会话窗口

当你完成工作，想要退出某个 `screen` 会话时，可以简单地输入：

```bash
exit
```

这会结束当前的会话，并关闭相关的窗口。

## 高级操作

### 在 screen 中运行命令

有时候，我们希望直接在创建 `screen` 会话的同时执行某个命令，并且在该命令执行完毕后自动结束会话。可以这样做：

```bash
screen vim memo.txt
```

这个命令会在一个新的 `screen` 会话中打开 `vim` 编辑器编辑 `memo.txt` 文件，当你退出 `vim` 时，该 `screen` 会话也会自动结束。

### 会话共享功能

`screen` 的一个强大功能是支持会话共享，意味着多个用户可以实时共享和操作同一个会话视图。

```bash
# 终端 A：创建会话
screen -S backup

# 终端 B：同步终端信息
screen -x

# 或者通过指定具体的 session-id 来共享会话
screen -x <screen-session-id>
# 比如
screen -x 364490.backup
```

这对于协作调试和教学非常有用。

### 重新连接断开的会话

当 `screen` 会话因为网络问题或其它原因断开时，你可以很容易地重新连接到这个会话：

```bash
screen -r backup
```

如果 `screen` 提示会话状态为 Attached，说明会话仍被另一个连接占用，你可以使用下面的命令强制回收：

```bash
screen -D -r backup
# 或
screen -D -r <screen-session-id>
# 比如
screen -D -r backup
```

## 小结

`screen` 是一个强大的工具，通过它，我们可以在远程服务器上高效地管理多个会话，保证关键任务的持续运行，甚至在不同用户之间共享会话，以便于协作和教学。

使用简单的操作，却能极大地提升我们的工作效率和协作能力。

希望本文能帮助你入门并实践使用 `screen`，让你的 Linux 经验更上一层楼。