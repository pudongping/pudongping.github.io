---
title: '深入理解Vim保存命令“:w !sudo tee %”'
author: Alex
top: false
hide: false
cover: false
toc: true
mathjax: false
categories: Linux
tags:
  - Vim
abbrlink: c1413261
date: 2023-12-19 14:57:15
img:
coverImg:
password:
summary:
---

在 Linux 世界中，Vim 编辑器是一个广受欢迎的文本编辑器，它强大的功能和丰富的命令集合赋予了它无与伦比的生命力。今天我们要来探讨的是 Vim 中的一个**很常用但可能初学者易于忽略**的命令——`:w !sudo tee %`。这个命令整体上的含义是当你打开一个需要 root 权限**才能修改**的文件，而你修改文件后又不想退出 Vim 编辑器，然后再使用 sudo 命令来启动 Vim，这时你就可以使用这个命令来获取临时的权限，保存你的更改。

我们首先拆分这个命令的每一个部分来理解它是如何工作的。

1. `:`：这在 Vim 中意味着我们将要输入一条命令。在命令模式下，你可以在标题栏中用`: `来开始。

2. `w`：":w" 是一个常见的 Vim 命令，代表着 “write”。这个命令单独使用的话，会将你在 Vim 中所做的改动写入到文件中。这里的写入并不直接涉及文件系统，而是写入给 Vim 的缓冲区。

3. `!`：在 Vim 中，感叹号('!')表示我们要在 shell 环境下执行接下来的命令。它告诉 Vim 去调用一个外部的程序。

4. `sudo`：在 Linux 命令中，`sudo` 命令（Su"do"（superuser; substitute user; switch user do））通常表示位"以超级用户身份去做"的意义。之所以需要 `sudo` 是因为原本你可能没有写入这个文件需要的权限。

5. `tee`：`tee` 是一个标准的 UNIX 命令，它可以接受输入，并将结果重定向到文件和屏幕（stdout）。

6. `%`：在 Vim 中，'%' 符号表示当前正在编辑的文件。

所以，这个命令的整体含义就代表着： **把当前的缓冲区内容输出到屏幕，并以超级用户的身份通过 `tee` 命令写入到当前正在编辑的文件中**。

主要用于：**在没有文件写入权限却需要保存的情况下，赋予用户以超级用户的身份保存正在编辑的文件。**

事实上，这个命令背后的原理更为高深。当你在命令行界面中输入一条带有管道（|）的命令，Shell 会按照顺序一一处理你的命令。在我们的这个命令中，首先它被`:`解释为 Vim 命令，此时的 `w`  指定我们将缓冲区的数据写出。而参数 `!sudo tee %` 被构造为一个传给 Shell 的命令并被执行。`%` 在 Shell 语境下又被 Vim 具体化为当前编辑的文件名，于是 Shell 会执行 `tee` 命令，将缓冲区的数据简单复制到指定的文件和控制台。

这个在 Vim 中绝妙的命令背后的原理包含了多个琐碎的信息，每个都是程序员在日复一日的操作中形成的最佳实践。

掌握了这个命令，也就是说我们对 Vim，对 Shell，对 Linux 文件权限管理有了更深一层的理解。这也是 Vim 为何如此迷人的原因之一，在 Vim 中，有着无数这样的小技巧等你去发掘，正如一位大师曾经说过，“Vim就像一个无尽宝藏，总有你未曾发现的新世界。” 是的，这是我刚刚说的，哈哈……