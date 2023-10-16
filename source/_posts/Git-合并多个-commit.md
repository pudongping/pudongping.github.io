---
title: Git 合并多个 commit
author: Alex
top: false
hide: false
cover: false
toc: true
mathjax: false
categories: Git
tags: Git
abbrlink: d05f7a67
date: 2023-10-16 11:11:37
img:
coverImg:
password:
summary:
---

# Git 合并多个 commit

## 查看提交历史 `git log`

```bash

# 最近的第 1 条
commit ddc9e34424e8764357d086ad103219fa2c87e2dd

# 最近的第 2 条
commit 3cccd5d8696b91163b47a8be045e8bbf9c443ddd

# 最近的第 3 条
commit d70ae2c4c1c6d2edd16c6b11a8e334663a3dade5

# 最近的第 4 条
commit fed4fe30dbc89ccc7ae72b917b2a600ecf249d4e

```

## git rebase

如果想要合并最近的 1 到最近的第 3 条，有两个方法：

1. 从 `HEAD` 版本开始往过去数 3 个版本，比如：

```bash
git rebase -i HEAD~3
```

2. 指定要合并的版本之前的一个版本号，比如：

```bash
# 比如我想要将最近的 1 条 commit 和最近的第 3 条 commit 进行合并，那么
# 我这里是需要写最近第 4 条 commit 的 commit_id `fed4fe30db` 
# fed4fe30db 不参与合并
git rebase -i fed4fe30db
```

## 选取要合并的提交

1. 执行了 `git rebase` 命令之后，会弹出一个窗口，比如大致如下：

```bash

pick 3cccd5d 对词库加上读写锁
pick ddc9e34 debug

# Rebase d70ae2c..ddc9e34 onto d70ae2c (2 commands)
#
# Commands:
# p, pick <commit> = use commit
# r, reword <commit> = use commit, but edit the commit message
# e, edit <commit> = use commit, but stop for amending
# s, squash <commit> = use commit, but meld into previous commit
# f, fixup <commit> = like "squash", but discard this commit's log message
# x, exec <command> = run command (the rest of the line) using shell
# b, break = stop here (continue rebase later with 'git rebase --continue')
# d, drop <commit> = remove commit
# l, label <label> = label current HEAD with a name
# t, reset <label> = reset HEAD to a label
# m, merge [-C <commit> | -c <commit>] <label> [# <oneline>]
# .       create a merge commit using the original merge commit's
# .       message (or the oneline, if no original merge commit was
# .       specified). Use -c <commit> to reword the commit message.
#
# These lines can be re-ordered; they are executed from top to bottom.
#
# If you remove a line here THAT COMMIT WILL BE LOST.
#
# However, if you remove everything, the rebase will be aborted.
#
# Note that empty commits are commented out

```

- pick：正常选中
- reword：选中，并且修改提交信息
- edit：选中，rebase 时会暂停，允许你修改这个 commit
- squash：选中，会将当前 commit 与上一个 commit 合并
- fixup：与 squash 相同，但不会保存当前 commit 的提交信息
- exec：执行其他 shell 命令

2. 需要将 commit_id 前面的 `pick` 改为 `s` 或者 `squash` 之后保存并关闭文本编辑窗口，改完之后的内容如下：（这里仅仅展示了内容变动情况）

```bash

pick 3cccd5d 对词库加上读写锁
s ddc9e34 debug # 这一行做了改动

```

3. 然后保存并退出，git 就会压缩提交历史，如果有冲突，则需要解决冲突，解决冲突的时候需要注意，保留最新的历史，不然我们的修改就丢弃了，修改之后要记得敲下面的命令

```bash

git add .

# 确认 rebase
git rebase --continue

```

如果你想放弃这次压缩的话，那么可以执行以下命令

```bash
# 取消 rebase
git rebase --abort
```

4. 如果没有冲突，或者冲突已经解决，则会出现如下编辑窗口，比如：

```bash

pick 3cccd5d 对词库加上读写锁
# This is a combination of 2 commits.
# This is the 1st commit message:

对词库加上读写锁

# This is the commit message #2:

debug

# Please enter the commit message for your changes. Lines starting
# with '#' will be ignored, and an empty message aborts the commit.

```

5. 输入 `:wq` 保存并退出，然后查看 `git log --oneline` 查看 `commit` 历史信息，你就会发现 commit 已经被合并了
6. 强制推送到远程服务器 `git push -f`