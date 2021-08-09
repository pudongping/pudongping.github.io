---
title: Git 半小时入门 <第一章>
author: Alex
top: false
hide: false
cover: false
toc: true
mathjax: false
categories: Git
tags:
  - Git
abbrlink: 89044c18
date: 2021-08-09 12:08:25
img:
coverImg:
password:
summary:
---


# Git 半小时入门 <第一章>

## 邂逅 Git
1. git 是什么？
   Git 是目前世界上最先进的分布式版本控制系统（没有之一）。
2. git 可以干什么？
   对文件进行版本控制。
3. git 的优点？
   高端大气上档次！（速度最快，操作简单）
4. git 和 CVS、SVN 这些版本控制系统有什么区别？
   最大的一个区别在于 CVS、SVN ……这些版本控制系统是 「集中式」的版本控制系统，而 Git 是 「分布式」的版本控制系统。
5. git 起源原因
   很多人都知道，Linus 在 1991 年创建了开源的 Linux 系统，起初 Linus 通过手工合并代码，但随后随着代码量越来越大，Linus 无法再通过手工合并代码，因此采用了商业的版本控制系统 BitKeeper，起初 BitKeeper 的东家免费将该版本控制系统提供给 Linux 社区使用，但开发 Samba 的 Andrew 试图破解 BitKeeper 的协议，这下把 BitKeeper 的东家给惹怒了，要收回 Linux 社区的免费使用权，结果 Linus 花了两周时间自己用 C 写了一个分布式版本控制系统，这就是 Git！随后在 2008年，GitHub 网站上线了，它为开源项目免费提供 Git 存储，无数开源项目开始迁移至 GitHub，包括 jQuery，PHP，Ruby 等等。

### 集中式和分布式的区别
**集中式版本控制系统**，版本库是集中存放在中央服务器的，
而干活的时候，用的都是自己的电脑，所以要先从中央服务器
取得最新的版本，然后开始干活，干完活了，再把自己的活
推送给中央服务器。

![集中式版本控制系统](https://upload-images.jianshu.io/upload_images/14623749-2aa9d28c6ff90e8f.png?imageMogr2/auto-orient/strip%7CimageView2/2/w/1240)

> 最大的毛病就是必须联网才能工作，并且如果中央服务器出现了问题，那么所有人都没法干活儿了

**分布式版本控制系统**根本没有“中央服务器”，每个人的电脑上都是一个完整的版本库，这样，你工作的时候，就不需要联网了，因为版本库就在你自己的电脑上。

![分布式版本控制系统](https://upload-images.jianshu.io/upload_images/14623749-8265e227f4d5f25b.png?imageMogr2/auto-orient/strip%7CimageView2/2/w/1240)

## 安装 Git
1. 在 Linux 上安装 Git
   Debian 或 Ubuntu 系列： sudo apt-get install git 或者 sudo apt-get install git-core（因为以前有个软件也叫 GIT（GNU Interactive Tools），结果 Git 就只能叫 git-core 了。由于 Git 名气实在太大，后来就把 GNU Interactive Tools 改成 gnuit，git-core 正式改为 git。）
- Redhat 或 CentOS 系列：yum install git
- Mac OS X 上安装 git有两种方法
    - 一是安装homebrew，然后通过homebrew安装Git，具体方法请参考homebrew的文档：http://brew.sh/
    - 二是直接从AppStore安装Xcode，Xcode集成了Git，不过默认没有安装，你需要运行Xcode，选择菜单“Xcode”->“Preferences”，在弹出窗口中找到“Downloads”，选择“Command Line Tools”，点“Install”就可以完成安装了。
2. 在 Windows 上安装 git
   直接在 git 官网 https://git-scm.com/downloads 上面下载 .exe 软件，然后一直点击下一步即可。

### 以下讲解在 Windows 系统下安装

![官网下载 Git ](https://upload-images.jianshu.io/upload_images/14623749-abc8c279ef435fc6.png?imageMogr2/auto-orient/strip%7CimageView2/2/w/1240)

![一直点击“Next”即可安装](https://upload-images.jianshu.io/upload_images/14623749-76eb47be74fdfd67.png?imageMogr2/auto-orient/strip%7CimageView2/2/w/1240)

![鼠标右键或者任务栏中出现 “Git Bash”则表示安装成功](https://upload-images.jianshu.io/upload_images/14623749-7e64397183b9823b.png?imageMogr2/auto-orient/strip%7CimageView2/2/w/1240)

#### 配置 GIT

> 因为 Git 是分布式版本控制系统，所以，每个机器都必须自报家门：你的名字和 Email 地址。

- 配置用户名

  ```git

  git config --global user.name "<Your Name>"

  ```

![配置用户名](https://upload-images.jianshu.io/upload_images/14623749-2cab46c99eac9e15.png?imageMogr2/auto-orient/strip%7CimageView2/2/w/1240)


- 配置 Email 地址

```git

git config --global user.email "<email@example.com>"

```

![配置 Email 地址](https://upload-images.jianshu.io/upload_images/14623749-912cdc2fb3e06ec5.png?imageMogr2/auto-orient/strip%7CimageView2/2/w/1240)

### 使用版本控制系统须知

**所有的版本控制系统，其实只能跟踪文本文件的改动，** 比如 TXT 文件，网页，所有的程序代码等等，Git 也不例外。
版本控制系统可以告诉你每次的改动，比如在第 5 行加了一个单词“Linux”，在第 8 行删了一个单词“Windows”。
而 **图片、视频这些二进制文件，虽然也能由版本控制系统管理，但没法跟踪文件的变化，** 只能把二进制文件每次改动串起来，
也就是只知道图片从 100KB 改成了 120KB，但到底改了啥，版本控制系统不知道，也没法知道。

不幸的是，Microsoft 的 Word 格式是二进制格式，因此， **版本控制系统是没法跟踪 Word 文件的改动的，** 如果要真正使用版本控制系统，
就要以纯文本方式编写文件。

因为文本是有编码的，比如中文有常用的 GBK 编码，日文有 Shift_JIS 编码，如果没有历史遗留问题，强烈建议使用标准的 UTF-8 编码，
所有语言使用同一种编码，既没有冲突，又被所有平台所支持。

使用 Windows 的童鞋要特别注意：

**千万不要使用 Windows 自带的记事本编辑任何文本文件。** 原因是 Microsoft 开发记事本的团队使用了一个非常弱智的行为来保存 UTF-8 编码的文件，
他们自作聪明地在每个文件开头添加了 0xefbbbf（十六进制）的字符，你会遇到很多不可思议的问题，比如，网页第一行可能会显示一个“?”，
明明正确的程序一编译就报语法错误，等等，都是由记事本的弱智行为带来的。建议你下载 Notepad++ 代替记事本，不但功能强大，而且免费！
记得把 Notepad++ 的默认编码设置为 UTF-8 without BOM 即可：

### Git 工作流

你的本地仓库由 git 维护的三棵“树”组成。
第一个是你的 工作目录，它持有实际文件；
第二个是 缓存区（Index），它像个缓存区域，临时保存你的改动；
最后是 HEAD，指向你最近一次提交后的结果。并且 git 为我们自动创建的
第一个分支 master，以及指向 master 的一个指针叫 HEAD。

![Git工作流](https://upload-images.jianshu.io/upload_images/14623749-8fb97ab7eaf4a408.png?imageMogr2/auto-orient/strip%7CimageView2/2/w/1240)

![Git工作流](https://upload-images.jianshu.io/upload_images/14623749-963c15ac783353df.png?imageMogr2/auto-orient/strip%7CimageView2/2/w/1240)
