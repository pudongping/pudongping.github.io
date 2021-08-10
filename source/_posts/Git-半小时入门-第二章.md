---
title: Git 半小时入门 <第二章>
author: Alex
top: false
hide: false
cover: false
toc: true
mathjax: false
categories: Git
tags:
  - Git
abbrlink: 13c0012
date: 2021-08-10 17:57:51
img:
coverImg:
password:
summary:
---

# Git 半小时入门 <第二章>

## 使用 Git
1. 新项目如何使用 git ？
2. 已有项目如何使用 git 管理代码？
3. 如何回退版本？
4. 如何撤销修改？
5. 分支管理
6. 多人协作
7. 忽略文件的写法

### 新项目如何使用Git
1. 创建 gitdemo 目录（创建版本库）

  ```git
    mkdir gitdemo
  ```
2. 切换到版本库中，并初始化 git （初始化的意思是将当前目录交给 git 管理）

  ```git
    cd gitdemo && git init
  ```

![可以看到当前文件夹出现 .git 目录](https://upload-images.jianshu.io/upload_images/14623749-f250d2a5e2802918.png?imageMogr2/auto-orient/strip%7CimageView2/2/w/1240)

3. 写入测试文件
  ```git
    echo 123 > demo.txt
  ```
4.  查看当前 git 仓库状态 （此时的状态为 「未提交」）

  ```git
    git status
  ```

![此时的状态为没有将文件添加到 git 中](https://upload-images.jianshu.io/upload_images/14623749-525075aa4e3a14fd.png?imageMogr2/auto-orient/strip%7CimageView2/2/w/1240)

5. 提交当前 『当前修改』到暂存区
- 提交单个文件
  ```git
    git add demo.txt
  ```
- 提交多个文件
   ```git
     git add demo.txt other.txt text.txt
   ```
- 提交所有修改
     ```git
       git add -A # 提交所有修改
     ```
6. 提交到仓库中（将暂存区中的内容同步到本地仓库中）
    ```git
      git commit -m “<description current commit>”
    ```

![image.png](https://upload-images.jianshu.io/upload_images/14623749-a2a7007346846470.png?imageMogr2/auto-orient/strip%7CimageView2/2/w/1240)

7. 创建远程库（这里使用全球最大的 git 仓库 GitHub 作为示范）

![image.png](https://upload-images.jianshu.io/upload_images/14623749-a3d1feac668ed82a.png?imageMogr2/auto-orient/strip%7CimageView2/2/w/1240)

![image.png](https://upload-images.jianshu.io/upload_images/14623749-ad92163aeb466d05.png?imageMogr2/auto-orient/strip%7CimageView2/2/w/1240)

![至此，在 GitHub 上面创建远程仓库已经完毕。](https://upload-images.jianshu.io/upload_images/14623749-8103d08391732d57.png?imageMogr2/auto-orient/strip%7CimageView2/2/w/1240)

8. 本地仓库和远程仓库创建关联关系
  ```git
    git remote add origin git@github.com:<Your name>/demogit.git
  ```
9. 将本地仓库中的代码推送到远程仓库中
  ```git
    git push -u origin master
  ```
10. 从远程仓库中拉取代码到本地仓库中
  ```git
    git fetch origin master && git pull 
  ```
也可以直接使用 git pull

![image.png](https://upload-images.jianshu.io/upload_images/14623749-a9df74b0b3da3cc7.png?imageMogr2/auto-orient/strip%7CimageView2/2/w/1240)

11. 生成 SSH Key。

我们知道 GitHub 远程库支持 HTTPS 和 SSH 两种方式通讯，如果你就想采用 HTTPS 的话，那么请忽略这一步，但是如果采用 HTTPS 通讯，每次提交的时候就需要输入一次账号和密码，也挺烦的。接下来讲解，如何采用 SSH 协议通讯。

生成 SSH Key 密钥对

  ```git
  ssh-keygen -t rsa -C "youremail@example.com" 
  ```

如果一切顺利的话，可以在用户主目录里找到 .ssh 目录，里面有 id_rsa 和 id_rsa.pub 两个文件，这两个就是 SSH Key 的秘钥对，id_rsa 是私钥，不能泄露出去，id_rsa.pub 是公钥，可以放心地告诉任何人。
将生成的 id_rsa.pub 公钥内容添加到 GitHub 中

生成 SSH 密钥对

![image.png](https://upload-images.jianshu.io/upload_images/14623749-11d85bb5fe6494d1.png?imageMogr2/auto-orient/strip%7CimageView2/2/w/1240)

查看生成的密钥对，生成的路径在 「用户家目录/.ssh」

![image.png](https://upload-images.jianshu.io/upload_images/14623749-7e60b43f07906de8.png?imageMogr2/auto-orient/strip%7CimageView2/2/w/1240)

![image.png](https://upload-images.jianshu.io/upload_images/14623749-8d9715611944f0a6.png?imageMogr2/auto-orient/strip%7CimageView2/2/w/1240)

![image.png](https://upload-images.jianshu.io/upload_images/14623749-17283add08b5bad4.png?imageMogr2/auto-orient/strip%7CimageView2/2/w/1240)

![image.png](https://upload-images.jianshu.io/upload_images/14623749-a79d366fc59768a4.png?imageMogr2/auto-orient/strip%7CimageView2/2/w/1240)

将公钥添加到 GitHub 中，至此我们再次去推送一次本地代码到远程库，即可看见可以推送成功了！

![image.png](https://upload-images.jianshu.io/upload_images/14623749-1d6bd88c7c40f6b4.png?imageMogr2/auto-orient/strip%7CimageView2/2/w/1240)

### 新项目如何使用 Git （总结）

一个新项目如何使用 git 来管理，并创建远程仓库？
1. 使用 GitHub、gitee、coding 创建好远程仓库
2. 如果采用 ssh 协议通讯的话，那么需要先在本地生成 ssh key 密钥对，并且将公钥添加到远程库账号 ssh 相关设置中，如果采用 https 协议通讯的话，则跳过此步骤。
   `ssh-keygen -t rsa -C "youremail@example.com" `
3. 将代码通过远程库来管理
   可以直接将远程仓库克隆到本地，此种方法最为简单。
   `git clone git@github.com:<Your GitHub account name>/demogit.git`
   可以先在本地创建一个仓库之后，再将本地仓库和远程仓库创建关联关系，稍许复杂。
- 创建项目目录并 git 初始化项目
  ```git
    mkdir -p ~/demogit && cd ~/demogit && git init 
  ```

- 创建远程连接
  ```git
    git remote add origin git@github.com:<Your GitHub account name>/demogit.git
  ```

- 查看远程连接是否创建成功
  ```git
    git remote -vv
  ```
4.  创建测试文件并提交代码到本地仓库
  ```git
    cd ~/demogit && echo 123 > demo.txt && git add -A && git commit -m “commit for test” 
  ```
5. 将本地代码推送到远程库
  ```git
    # 第一次推送的时候需要加上 「-u」参数，后续推送的时候，可以不用加
    git push -u origin master
  ```


### 已有项目如何使用git管理代码？

1. 先创建远程仓库，然后将远程空仓库拉到本地，之后将已有的项目全部复制到空本地仓库中
2. 本地直接在项目中 git 初始化项目，然后创建远程连接，之后将本地代码全部添加到 「暂存区」，再提交到本地仓库中，之后再推送。

### 如何回退版本？

1. reset （直接回退版本）
```git
  git reset --hard HEAD^  # 回退到上一个版本
  git reset --hard HEAD^^^  # 回退到上三个版本
  git reset --hard HEAD~100  # 回退到上 100 个版本
  git reset --hard <commit id> # 回退到指定版本
```
2.   revert （将回退过程当成一个版本提交）
```git
git revert <commit id>
```

如何查看提交版本号？
  ```git
    git log -2  # 查看倒数 2 条提交记录，提交记录中含有版本号
    git reflog  # 查看所有的提交版本号，如果要重返未来，用git reflog查看命令历史，以便确定要回到未来的哪个版本。
    git log --pretty=oneline
  ```


### 如何撤销修改？

因为 git 主要有三种状态，因此就有三种可能性，工作区、暂存区、版本库
1. 当修改只在工作区，还没有 git add 时： （一定要注意有 「--」）
  ```git
    git checkout -- <filename>  比如： git checkout -- example.txt
  ```

2. 当修改在暂存区时，已经 git add：
  ```git
  2.1. git reset HEAD <filename>  比如：git reset HEAD example.txt
  2.2. git checkout -- <filename>  比如：git checkout -- example.txt
  ```
3. 当修改已经从暂存区提交到了本地仓库，已经 git commit：
  ```git
  git reset --hard HEAD^
  ```

### 分支管理

#### 什么是分支？
分支是用来将特性开发绝缘开来的。在你创建仓库的时候，master 是“默认的”。在其他分支上进行开发，完成后再将它们合并到主分支上。

![你正在学习git，同时异步空间的你又在学习svn，最后合并，二者你都会了](https://upload-images.jianshu.io/upload_images/14623749-1ebabe114f060966.png?imageMogr2/auto-orient/strip%7CimageView2/2/w/1240)

- 创建一个叫做 「feature_x」 分支，并切换到 「feature_x」 分支
  ```git
  git checkout -b feature_x  或者 git switch -c feature_x 
  ```

- 创建分支
  ```git
    git branch feature_x
  ```
- 从当前分支切换到新的分支
  ```git
    git checkout feature_x  或者 git switch feature_x
  ```
- 查看分支
  ```git
    git branch
  ```
- 删除分支
  ```git
    # 将新建的分支删除（已经合并后）
    git branch -d feature_x

    # （未合并，强删）
    git branch -D feature_x
  ```
- 合并某分支到当前分支
  ```git
  git merge <branch name> 
  ```

- 除非将分支推送到远程仓库，不然该分支其他人不可见
  ```git
    git push origin <branch>  
  ```
- 查看分支合并图
  ```git
    git log --graph 
  ```

![image.png](https://upload-images.jianshu.io/upload_images/14623749-7a1ec719f624521e.png?imageMogr2/auto-orient/strip%7CimageView2/2/w/1240)

### 多人协作

1. 首先，可以试图用 **git push origin <branch-name>** 推送自己的修改；
2. 如果推送失败，则因为远程分支比你的本地更新，需要先用 **git pull** 试图合并；
3. 如果合并有冲突，则解决冲突，并在本地提交；
4. 没有冲突或者解决掉冲突后，再用 **git push origin <branch-name>** 推送就能成功！

如果 git pull 提示 no tracking information，则说明本地分支和远程分支的链接关系没有创建，用命令 **git branch --set-upstream-to <branch-name> origin/<branch-name>**。
直接从远程仓库拉取指定分支，并在本地创建和远程同名分支
```git
  git branch --set-upstream branch-name origin/branch-name
```

### 忽略文件的写法 （.gitignore）

1. 忽略文件的作用
   忽略不想提交的文件
2. 忽略文件的写法
   直接在 .gitignore 文件中写 test.txt  # 表示忽略和 .gitignore 平级的 test.txt 文件
   直接写 upload/  # 表示忽略和 .gitignore 平级的 upload 下所有的文件等同于 upload/*
   直接写 upload/*.txt  # 表示忽略和 .gitignore 平级的 upload 文件夹下所有的以 .txt 命名的文件
   直接写 *.py[cod] hu'lue # 表示忽略和 .gitignore 平级的 所有文件名是 .pyc 或 .pyo 或 .pyd 的文件
   忽略所有 .txt 的文件，但是除了 a.txt 的文件的写法 *.txt !a.txt
3. 如果某个文件已经在忽略文件中忽略，但是又想要提交该文件
   git  add -f <file name>  # 强制提交 force.txt 文件 => git add -f force.txt
4. 检查某个文件是否被忽略文件所忽略
   git check-ignore -v <file name> # 查看 force.txt 是否被忽略 =>  git check-ignore -v force.txt
