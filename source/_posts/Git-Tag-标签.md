---
title: Git Tag 标签
author: Alex
top: false
hide: false
cover: false
toc: true
mathjax: false
categories: Git
tags:
  - Git
abbrlink: '68649176'
date: 2021-08-06 23:33:49
img:
coverImg:
password:
summary:
---

# Git Tag 标签

标签总是和某个 commit 挂钩。如果这个 commit 既出现在 master 分支，又出现在 dev 分支，那么在这两个分支上都可以看到这个标签。

- 创建标签

```git

# 给当前提交版本打标签
git tag <tag name>

# eg: 打一个名称为 v1.0 的标签，此时默认将标签打到最新提交的 commit 上
git tag v1.0

```

- 给指定提交版本打标签

```git

# 查看历史提交的 commit id
git log --pretty=oneline --abbrev-commit

# 给指定提交版本打标签
git tag <tag name> <commit id>

# eg: 给提交版本号为 c53b867 的版本，打一个名称为 v0.8 的标签
git tag v0.8 c53b867

```

- 创建带有说明的标签，用 `-a` 指定标签名，`-m` 指定说明文字：

```git

git tag -a <tag name> -m <tag description> <commit id>

# eg：给版本号为 c53b867 的版本，打一个名称为 v0.8 的标签，并对 v0.8 这个标签添加说明文字为 "add v0.8 tag"
git tag -a v0.8 -m "add v0.8 tag" c53b867

```

- 查看所有的标签 （标签不是按时间顺序列出，而是按字母排序的。）

```git

git tag

```

- 查看标签信息

```git

git show <tag name>

# eg：查看标签名得 v0.9 的标签信息
git show v0.9

```

- 删除标签

```git

git tag -d <tag name>

# eg：删除标签名为 v0.1 的标签
git tag -d v0.1

```

- 推送本地标签到远程库

```git

git push origin <tag name>

# eg： 推送本地标签名为 v0.1 的标签到远程库中
git push origin v0.1

```

- 一次性推送全部尚未推送到远程的本地标签

```git

git push origin --tags

```

- 如果标签已经推送到远程，想要删除远程标签

```

# 第一步：删除本地标签
git tag -d v0.1

# 第二步：从远程删除
git push origin :refs/tags/v0.1

# 第三步：在远程库中查看是否被删除
```
