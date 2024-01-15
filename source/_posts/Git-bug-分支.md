---
title: Git bug 分支
author: Alex
top: false
hide: false
cover: false
toc: true
mathjax: false
categories: Git
tags:
  - Git
abbrlink: 56cac605
date: 2024-01-15 14:15:19
img:
coverImg:
password:
summary:
---

# Git bug 分支

> 修复 bug 时，我们会通过创建新的 bug 分支进行修复，然后合并，最后删除；   
当手头工作没有完成时，先把工作现场 git stash 一下，然后去修复 bug，修复后，再 git stash pop，回到工作现场；   
在 master 分支上修复的 bug，想要合并到当前 dev 分支，可以用 git cherry-pick <commit id> 命令，把 bug 提交的修改“复制”到当前分支，避免重复劳动。


应用场景如下：   
当你需要修改一个 bug，但是你当前的开发功能还不想提交到远程版本库中，又必须要先将 bug 提交。

1. 查看当前工作区的状态（假设当前在 `dev` 分支）

```git 

# 工作区的文件还没有提交
git status

```

2. 存储当前的工作现场

```git

git stash

```

3. 再次查看工作区的状态，应该是干净的
4. 如果需要在 `master` 分支上修复 bug，那么就需要从 `master` 分支上创建临时分支

```git

# 切换到 master 分支
git checkout master

# 创建修复 bug 的分支为 issue-1
git checkout -b issue-1
```

5. 在 `issue-1` 分支上修复 bug
6. 在 `issue-1` 分支上提交代码

```git

git add -A && git commit -m 'fixed bug 1' 

```

7. 切换到 `master` 分支

```git

git switch master

```

8. 从 `master` 分支上合并 `issue-1` 分支的代码，并添加了合并信息

```git

git merge --no-ff -m "merged bug fix 1" issue-1

```

9. 原来是在 `dev` 分支上干活，现在切换到 `dev` 分支上继续干活

```git

git checkout dev

# 如果此时用 git status 命令查看的话，此时的工作区是干净的

```

10. 查看之前临时存储的工作现场

```git

git stash list

```

11. 恢复工作现场（有两种方法）

```git

# 第一种方法：
# 恢复工作现场
git stash apply
# 删除之前临时存储的工作现场
git stash drop

# 第二种方法：
# 恢复的同时把 stash 内容也删了
git stash pop

```

12. 如果需要恢复指定的工作现场

```git

# 查看所有的工作现场
git stash list

# 恢复名称为 「stash@{0}」 的工作现场
git stash apply stash@{0}

```

13. 将修复好的 bug 同步到其他的分支

> 可以通过直接将修复 bug 的文件修改内容复制到其他分支，比如在 `issue-1` 分支上修复的提交版本号为 `c55ae16d5b1a`，现在只需要将这个版本号的所有修改内容复制到 dev 分支上即可。

```git

# 当前分支为 dev
git checkout dev

# 复制一个特定的提交到当前分支（此操作会自动提交一个版本号）
git cherry-pick c55ae16d5b1a

```