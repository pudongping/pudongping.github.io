---
title: Git 常用命令
author: Alex
top: false
hide: false
cover: false
toc: true
mathjax: false
categories: Git
tags:
  - Git
abbrlink: 5fddf106
date: 2021-08-12 20:26:01
img:
coverImg:
password:
summary:
---


# Git 命令

## 生成 SSH Key

```git
ssh-keygen -t rsa -C "youremail@xxx.com"
# 输入之后直接两次回车即可
```

## 查看 git 用户名和邮箱地址

```git
# 查看用户名
git config user.name
# 查看邮箱地址
git config user.email
# 查看配置信息
git config --list

# 当前用户全局
git config --global
# 当前系统全局
git config --system
# 修改用户名
git config --global user.name "username"
# 修改邮箱地址
git config --global user.email "email"

# 重新输入账号密码
git config --system --unset credential.helper

# 查看 git 配置信息
git config --list 或者 git config -l
```

## 初始化一个 git 仓库

```git
git init
```

## 添加文件到 git 仓库

### 1. 添加文件到缓存区

```git
git add <filename>
# 比如添加 file.txt 文件： git add file.txt
```

### 2. 将缓存区文件提交到本地仓库

```git
git commit -m "add_file_message"
# 比如提交 file.txt 文件：git commit -m "add file.txt"
```

## 查看修改内容对比，提交到缓存区或者已经提交到本地仓库，使用 git diff 会没有内容显示，也就是说只要修改了，在提交到缓存区之前使用 git diff 才有用 （查看工作区的改变）

```git
git diff

# 还可以查看具体哪个文件做了什么修改
# 比如查看 readme.txt 做了什么修改
git diff HEAD -- readme.txt
```

## 查看当前仓库状态，任何时候都可以使用

```git
git status
```

## 查看提交日志

```git
# 查看全部 
git log

# 查看最后一次提交
git show

# 查看倒数5条
git log -5

# 简化日志显示方式，并含有提交版本号
git log --pretty=oneline

# 比如：
# $ git log --pretty=oneline
# 44c9beb4c58543b89181829755be2c5e9781ba28 (HEAD -> master) append GPL
# 03112bdf101655c30df9b61e4bd325b2cbe3c090 add distributed
# 8a1386bd0fe677bca99d5a4ef26e87772a3eca71 wrote a readme file
```

## 设置别名 Alias

```git
# 创建/查看本地分支
git config --global alias.br "branch"
# 切换分支
git config --global alias.co "checkout"
# 创建并切换到新分支
git config --global alias.cb "checkout -b"
# 提交
git config --global alias.cm "commit -m"
# 查看状态
git config --global alias.st "status"
# 拉取分支
git config --global alias.pullm "pull origin master"
# 提交分支
git config --global alias.pushm "push origin master"
# 单行、分颜色显示记录
git config --global alias.log "log --oneline --graph --decorate --color=always"
# 复杂显示
git config --global alias.logg "log --graph --all --format=format:'%C(bold blue)%h%C(reset) - %C(bold green)(%ar)%C(reset) %C(white)%s%C(reset) %C(bold white)— %an%C(reset)%C(bold yellow)%d%C(reset)' --abbrev-commit --date=relative"
git config --global alias.lg "log --color --graph --pretty=format:'%Cred%h%Creset -%C(yellow)%d%Creset %s %Cgreen(%cr) %C(bold blue)<%an>%Creset' --abbrev-commit"
```

## 回退版本

```git
# 回退当前版本的上一个版本
git reset --hard HEAD^

# 回退当前版本的上三个版本
git reset --hard HEAD^^^

# 回退当前版本的上 100 个版本
git reset --hard HEAD~100

# 回退指定版本，并将 「回退」动作作为一个版本提交
git revert <commit id>
```

## 回退指定版本（多半是因为回退版本回退多了）

```git
git reset --hard 版本号前几位

# 比如回退到 以上查看日志中的 add distributed 这个提交版本
git reset --hard 44c9beb4c
```

## 查看提交版本号 commit id

```git
# 要重返未来，用 `git reflog` 查看命令历史，以便确定要回到未来的哪个版本。
git reflog
```

## 撤销修改

### 1. 修改只在工作区，还没有添加到缓存区 （还没有 git add）

```git
git checkout -- filename
# 比如 readme.txt 文件只是在工作区修改了，想回退到修改之前的提交的版本
git checkout -- readme.txt
```

### 2. 修改从工作区已经提交到了缓存区 （已经 git add）

```git
# 将缓存区的修改回退到工作区
git reset HEAD filename

# 比如 readme.txt 文件的修改已经提交到缓存区了，但是想撤销修改分为两步
# 01.将缓存区的修改回退到工作区
git reset HEAD readme.txt
# 02.将在工作区的修改回退到和上一个版本一样
git checkout -- readme.txt
```

### 3. 修改已经从缓存区提交到了本地仓库 （已经 git commit）

```git
# 回退当前版本的上一个版本
git reset --hard HEAD^
```

### 4. 将在暂存区的更改文件进行强制撤销。（想让之前已经提交到缓存区的文件覆盖工作区的文件）

```git
git checkout -f
```

### 5. 命令 git clean 作用是清理项目，-f 是强制清理文件的设置，-d 选项命令连文件夹一并清除

```git
git clean -f -d
```

### 6. 假如你想要丢弃你所有的本地改动与提交，可以到服务器上获取最新的版本并将你本地主分支指向到它：

```git
git fetch origin 

git reset --hard origin/master
```

## 删除文件

### 1. 确实要删除该文件

```git
# 比如以 test.txt 为例子

# 01 本地手动删除 test.txt 文件
rm test.txt 或者 rm -rf test.txt
# 02 添加被删除的状态缓存区
git rm test.txt 或者 git add test.txt
# 03 提交状态到本地仓库
git commit -m "remove test.txt"
```

### 2. 工作区误删了文件

```git
# 01 工作区误删了 test.txt 文件
rm test.txt
# 02 找回被误删的文件（撤销修改）
git checkout -- test.txt
# 用版本库中的版本替换掉工作区的版本
```

## 创建分支并切换到该分支

```git
git checkout -b alex  或者 git switch -c alex

# 创建 alex 分支并切换到 alex 分支
# 等同于以下两个命令
# 创建 alex 分支
git branch alex
# 切换到 alex 分支
git checkout alex  或者 git switch alex
```

## 合并指定分支到当前分支

```git
# 比如当前位于 master 分支，欲将 alex 分支合并到 master 分支
git merge alex
# 如果有冲突，解决步骤如下：
# 01 先查看冲突文件
git status
# 02 手动解决冲突文件
# 03 再次合并分支
git merge alex
# 04 添加修改到缓存区
git add .
# 05 提交到本地仓库
git commit -m "merge fixed"
```

## 查看分支的合并情况

```git
git log --graph --pretty=oneline --abbrev-commit
或者直接使用 git log --graph 命令可以看到分支合并图。

# 设置别名查看所有的提交记录
git config --global alias.lg "log --color --graph --pretty=format:'%Cred%h%Creset -%C(yellow)%d%Creset %s %Cgreen(%cr) %C(bold blue)<%an>%Creset' --abbrev-commit"
```

## 删除分支

```git
# 合并代码之后删除分支 
git branch -d 分支名
# 没有合并代码删除分支 
git branch -D 分支名
# 删除远程分支（本地分支需要再次手动删除）
git push origin -d 分支名

# 比如删除 alex 分支
git branch -d alex
# 删除远程 alex 分支
git push origin -d alex
```

## 查看分支

```git
# 只查看远程分支
git branch -r
# 只查看本地分支
git branch
# 查看所有远程分支和本地分支
git branch -a
```

## 拉取远程分支并创建同名本地分支

```git
# 方法一（此方法建立的本地分支和远程分支会有映射关系）
git checkout -b [本地分支名] origin/[远程分支名]
举例：git checkout -b alex origin/alex

# 方法二（此方法建立的本地分支和远程分支没有映射关系）
git fetch origin [远程分支名]:[本地分支名]
举例：git fetch origin alex:alex
```

## 查看本地分支和远程分支的映射关系（远程有的分支而本地没有的分支不会出现）

```git
git branch -vv
```

## 手动建立本地分支和远程分支的映射关系

```git
git branch -u origin/[分支名]
或者
git branch --set-upstream-to origin/[分支名]
或者
git branch --set-upstream-to=origin/[远程分支名] [本地分支名]
```

## 推送分支

```git
git push -u origin <branch name> 
# 第一次推送的时候添加 -u 参数，给本地分支和远程分支创建连接关系，当第二次再次推送时，则不需要添加 -u 参数

# 比如：推送 dev 分支
git push origin dev
```

## 查看远程库的信息

```git
git remote

# 查看远程库更加详细的信息
# 这里可以看到抓取和推送的 origin 的地址
git remote -v （小写的v）

# 删除远程库连接
git remote rm <origin name>

# eg：删除远程 origin 连接
git remote rm origin

# 添加远程库连接
git remote add <origin name> <ssh or http>

# eg: 关联本人 GitHub 连接，并连接名为 myGitHub
git remote add myGitHub git@github.com:Alex66668888/demo.git
```

## 多处备份

### 一个 pull + 多个 push

```git

# 添加远程连接
$ git remote add origin git@gitee.com:pudongping/tt.git

# 查看远程连接地址，可以看到只有一个 fetch 地址和一个 push 地址
$ git remote -v
origin  git@gitee.com:pudongping/tt.git (fetch)
origin  git@gitee.com:pudongping/tt.git (push)

# 添加 github 推送地址
$ git remote set-url --add origin git@github.com:pudongping/tt.git

# 查看远程连接地址，可以看到只有一个 fetch 地址和两个 push 地址
$ git remote -v
origin  git@gitee.com:pudongping/tt.git (fetch)
origin  git@gitee.com:pudongping/tt.git (push)
origin  git@github.com:pudongping/tt.git (push)

# 添加 gitlab 推送地址
$ git remote set-url --add origin git@gitlab.com:pudongping/t1.git

# 查看远程连接地址，可以看到只有一个 fetch 地址和三个 push 地址
$ git remote -v
origin  git@gitee.com:pudongping/tt.git (fetch)
origin  git@gitee.com:pudongping/tt.git (push)
origin  git@github.com:pudongping/t1.git (push)
origin  git@gitlab.com:pudongping/tt.git (push)

# 推送到远程分支，就可以看到已经同步到多个平台上去了，一次 push 到多个远程仓库
git push
```


## git add 三种状态命令比较

命令 | 说明
--- | ---
git add . | 提交所有修改的和新建的数据暂存区 (提交当前文件夹下所有修改)
git add -u | 提交所有被删除和修改的文件到数据暂存区（等同于git add -update）
git add -A | 提交所有被删除、被替换、被修改和新增的文件到数据暂存区（等同于git add –all）

## 已经推送（push）过的文件，想从 git 远程库中删除，并在以后的提交中忽略，但是却还想在本地保留这个文件

```git
git rm --cached [file-path]
git rm --cached config/pay.php

# 如果是目录的话，则需要
git rm -r --cached [dir_name]

git push
```

## 已经推送（push）过的文件，想在以后的提交时忽略此文件，即使本地已经修改过，而且不删除 git 远程库中相应的文件

```git
# 只对文件有效
git update-index --assume-unchanged [file-path]
git update-index --assume-unchanged config/pay.php

# 如果需要恢复提交
git update-index --no-assume-unchanged [file-path]
```

## 合并一个分支上的修改到当前分支

```git
# 比如说 test1 分支上有一个提交 512d725 现在想将这个提交合并到 test2 分支上

# 先切换到 test2 分支上
git checkout test2

# 择优挑选（此时还是在 test2 分支上）
git cherry-pick 512d725 
```
