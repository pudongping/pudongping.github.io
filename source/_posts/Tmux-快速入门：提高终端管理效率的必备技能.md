---
title: Tmux 快速入门：提高终端管理效率的必备技能
author: Alex
top: false
hide: false
cover: false
toc: true
mathjax: false
categories: 科普
tags:
  - Tmux
abbrlink: bb0f40eb
date: 2024-07-02 15:31:02
img:
coverImg:
password:
summary:
---

在日常的软件开发过程当中，程序员经常需要同时操作多个终端窗口。不论是编写代码、运行测试、监控日志，还是远程登录服务器，多个窗口的切换不仅繁琐，而且降低了工作效率。

此时，一个叫作 **Tmux**（terminal multiplexer）的工具就能大显身手了。Tmux 允许你在一个终端窗口中，通过多个窗格（pane）和窗口（window）同时进行多项操作，极大地提高了使用终端的效率。

接下来，我们用简单易懂的语言介绍一下 Tmux 的基础使用方法，让你迅速上手这个强大的工具。

---

## 快捷键

快捷键是 Tmux 操作的核心，通过它们，你可以快速实现会话管理、窗格和窗口的操作。以下是一些基础且常用的快捷键列表。

### 会话管理
会话（Session）是 Tmux 的顶层组织单位，每个会话可以包含多个窗口。

- 分离当前会话： Ctrl+b d
- 列出所有会话： Ctrl+b s
- 重命名当前会话： Ctrl+b $

### 窗格操作
窗格是分屏的基本单位，一个窗口可以包含一个或多个窗格。

- 划分左右两个窗格： Ctrl+b %
- 划分上下两个窗格： Ctrl+b " （左双引号）
- 上下-左右窗格互切： Ctrl+b <space>  （空格）
- 光标切换到其他窗格： Ctrl+b <arrow key> （方向键）
- 当前窗格与上一个窗格互换位置： Ctrl+b {
- 当前窗格与下一个窗格互换位置： Ctrl+b }
- 关闭当前窗格： Ctrl+b x
- 将当前窗格拆分为一个独立窗口： Ctrl+b !
- 当前窗格全屏显示： Ctrl+b z （再次使用恢复原来大小）
- 调整窗格大小： Ctrl+b Ctrl+<arrow key>  （以 1 个单元格为单位调整）
- 调整窗格大小： Ctrl+b Alt+<arrow key>  （以 5 个单元格为单位调整）
- 显示窗格编号： Ctrl+b q
- 顺时针旋转当前窗口的窗格： Ctrl+b Ctrl+o  （字母 o）
- 逆时针旋转当前窗口的窗格：  Ctrl+b Alt+o
- 显示时钟：Ctrl+b t

### 窗口操作
窗口可以视为多个工作空间，每个窗口可以包含多个窗格。

- 创建新窗口： Ctrl+b c
- 切换到上一个窗口： Ctrl+b p
- 切换到下一个窗口： Ctrl+b n
- 切换到指定编号的窗口: Ctrl+b <number>
- 从列表中选择窗口： Ctrl+b w
- 窗口重命名： Ctrl+b ,

---

## 安装

在不同的操作系统上安装 Tmux 的命令如下：

```shell
# Ubuntu 或 Debian
sudo apt-get install tmux

# CentOS 或 Fedora
sudo yum install tmux

# Mac
brew install tmux
```

## 启动与退出

要开始使用 Tmux，只需打开终端并输入 `tmux` 来启动。

- **启动**：直接在终端中输入 `tmux`。
- **退出**：可以使用 `Ctrl + d` 或者输入 `exit` 命令。

## 会话管理

会话让你能在单一窗口中管理多个项目或任务，非常适合多任务操作。

1. **新建会话**

新建一个默认名称的会话：

```shell
tmux
```

或新建一个指定名称的会话：

```shell
tmux new -s <session-name>  # 新建一个名称叫做 alex 的会话
```

2. **分离会话**

让当前会话在后台运行，你可以安全地关闭终端连接，之后再重新接入：

```shell
Ctrl + b d  # 或者 tmux detach
```

3. **查看当前所有的 tmux 会话**

查看有哪些会话正在运行：

```shell
tmux ls  # 或者 tmux list-session
```

## 配置相关

- 系统配置文件 `/etc/tmux.conf`
- 用户级配置文件 `~/.tmux.conf`

配置文件实际上就是 `tmux` 的命令集合，也就是说每行配置均可在进入命令行模式后输入生效

- 将 `Ctrl + r` 设置为加载配置文件，并显示 `Refresh configure!`

```shell

bind C-r source-file ~/.tmux.conf \; display "Refresh configure!"

```

- 将 `prefix` 快捷键前缀由 `Ctrl+b` 更改为 `Ctrl+a`

```shell

# Send prefix
set-option -g prefix C-a
unbind-key C-a
bind-key C-a send-prefix

```

- 不用按快捷键前缀，直接使用 `alt + 方向键` 在 pane 之间 switch

```shell

# Use Alt-arrow keys to switch panes
bind -n M-Left select-pane -L
bind -n M-Right select-pane -R
bind -n M-Up select-pane -U
bind -n M-Down select-pane -D

```

- 不用按快捷键前缀，直接使用 `shift + 方向键` 在 window 之间 switch

```shell

# Shift arrow to switch windows
bind -n S-Left previous-window
bind -n S-Right next-window

```

- 开启鼠标模式，用鼠标就能切换 window 、 pane、还能调整 pane 的大小

```shell

# Mouse mode
set -g mouse on

```

- `Ctrl+b v` 竖着分屏， `Ctrl+b h` 横着分屏


```shell

# Set easier window split keys
bind-key v split-window -h
bind-key h split-window -v

```

## 总结

tmux 是一个功能强大的终端复用器，它通过丰富的快捷键和灵活的配置选项，极大地提升了我们使用终端的效率。无论你是开发者、系统管理员还是普通用户，掌握 tmux 都能帮助你更好地管理终端会话。

操作 Tmux 的精髓在于灵活地管理和切换会话、窗口和窗格。掌握以上基本命令，就可以让你的终端操作效率大大提高。继续探索 Tmux，你会发现更多强大功能，比如自定义快捷键、脚本自动化等，让你的工作更加得心应手。