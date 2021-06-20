---
title: Linux 管理远程会话 screen
author: Alex
top: false
hide: false
cover: false
toc: true
mathjax: false
categories: Linux
tags:
  - Screen
  - 会话管理
abbrlink: 4fadf1be
date: 2021-06-20 20:27:46
img:
coverImg:
password:
summary:
---

# Linux 管理远程会话 screen

1. 创建一个新的会话窗口 backup

```bash
screen -S backup
```

2. 查看会话会出现 session_id

```bash
screen -ls
```

3. 退出会话

```bash
exit
```

4. 直接使用 screen 命令执行要运行的命令，这样在命令中的一切操作也都会被记录下来，当命令执行结束后 screen 会话也会自动结束。

```bash
screen vim memo.txt
```

5. 会话共享功能

```bash
# 终端 A：创建会话
screen -S backup

# 终端 B：同步终端信息
screen -x

# 或者通过 screen-session-id 进入
screen -x <screen-session-id>

# 比如
screen -x 364490.backup
```

6. 进入会话

```bash
screen -r backup

# 有时候screen异常退出可能会提示状态为Attached，可以执行：screen -D -r backup 进行恢复。

screen -D -r <screen-name or screen-session-id>
# 比如
screen -D -r backup
```
