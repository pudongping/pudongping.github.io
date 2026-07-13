---
title: 手摸手教你玩 Jenkins，一次搞懂 CI/CD！（第一章：部署）
author: Alex
top: false
hide: false
cover: false
toc: true
mathjax: false
categories: Jenkins
tags:
  - Jenkins
  - CI
  - CD
  - 自动化部署
abbrlink: 51baa55e
date: 2026-07-13 13:37:16
img:
coverImg:
password:
summary:
---

对于开发者来说，能自动化部署项目，就像解锁了“自动驾驶”。而 Jenkins，就是让你从“手动上传代码、拷贝文件、重启服务”的地狱中解放出来的神器。

## Jenkins 是什么？

Jenkins是一款开源 CI&CD 软件，用于自动化各种任务，包括构建、测试和部署软件。

Jenkins 支持各种运行方式，可通过系统包、Docker 或者通过一个独立的 Java 程序。

简单来说，**Jenkins 是一个开源的自动化持续集成工具**。
它能帮你自动完成：

* 拉取最新代码；
* 自动打包；
* 自动部署到服务器；
* 甚至自动测试、自动通知。

## 为啥开发者都离不开 Jenkins？

你可以回想一下，每次项目更新你都得：

1. 打包；
2. 上传服务器；
3. 手动部署；
4. 启动服务；
5. 通知测试。

一天下来，不光浪费时间，还容易出错。
有了 Jenkins，你只需要点一下“构建”，剩下的流程全自动搞定。
这就是 **持续集成（CI）+ 持续部署（CD）** 的魅力。

> 本来想一次性讲完使用 Jenkins 从零到发布的整个流程，但是发现篇幅有点儿长了，于是打算分 3 篇来讲，这一篇是 Jenkins 的部署篇，后面两篇将是使用 Jenkins 发布项目篇。感兴趣的小伙伴可以蹲守一下。

## 安装 Jenkins

为了方便安装，我直接使用 docker 来进行安装 Jenkins。

```bash
docker pull jenkins/jenkins:lts-jdk17

docker run -d --name jenkins \
--restart=unless-stopped \
-u root \
-p 8080:8080 \
-p 50000:50000 \
-v /apps/docker_mounts_dir/jenkins/jenkins-data:/var/jenkins_home \
-e TZ=Asia/Shanghai \
-v /etc/localtime:/etc/localtime:ro \
-e JENKINS_UC=https://mirrors.cloud.tencent.com/jenkins/ \
-e JENKINS_UC_DOWNLOAD=https://mirrors.cloud.tencent.com/jenkins/ \
jenkins/jenkins:lts-jdk17
```

**注意：** 在国内一定要加上 `JENKINS_UC` 和 `JENKINS_UC_DOWNLOAD` 参数，不然在 Jenkins 中下载插件会慢到让你怀疑人生……

安装完毕之后，可以通过执行下面的命令来查看 admin 用户的密码

```bash
docker logs jenkins
```

![查看 admin 用户的密码](https://upload-images.jianshu.io/upload_images/14623749-994dea18abbc5d8f.png?imageMogr2/auto-orient/strip%7CimageView2/2/w/1240)

或者可以直接通过以下命令进行查看

```bash
docker exec jenkins cat /var/jenkins_home/secrets/initialAdminPassword
```

## 访问 Jenkins

然后就可以直接在浏览器中访问 `127.0.0.1:8080`

![访问 Jenkins 页面](https://upload-images.jianshu.io/upload_images/14623749-dd70ae7ed0662550.png?imageMogr2/auto-orient/strip%7CimageView2/2/w/1240)

安装插件

![选择插件来安装](https://upload-images.jianshu.io/upload_images/14623749-d3b6a885f5e417be.png?imageMogr2/auto-orient/strip%7CimageView2/2/w/1240)

进入下一个页面之后，不需要选择任何的插件，直接点击「安装」按钮就可以了。

![直接点击“安装”](https://upload-images.jianshu.io/upload_images/14623749-09e97ad62052aa0b.png?imageMogr2/auto-orient/strip%7CimageView2/2/w/1240)

等待插件安装

![等待插件安装](https://upload-images.jianshu.io/upload_images/14623749-11d70a9181fda540.png?imageMogr2/auto-orient/strip%7CimageView2/2/w/1240)

如果有插件安装失败，那么可以直接**重试**

![如果插件安装失败的话](https://upload-images.jianshu.io/upload_images/14623749-8cf922597aef22e0.png?imageMogr2/auto-orient/strip%7CimageView2/2/w/1240)

直到所有的插件安装成功之后，点击「继续」

![所有的插件安装成功之后](https://upload-images.jianshu.io/upload_images/14623749-b832a8320c495170.png?imageMogr2/auto-orient/strip%7CimageView2/2/w/1240)


然后需要创建一个管理员账号，当然了，不创建也可以，就继续使用 admin 账户登录，但是就得每次记录之前的那个长长的密钥。

![创建第一个管理员用户](https://upload-images.jianshu.io/upload_images/14623749-4ac54b505a887526.png?imageMogr2/auto-orient/strip%7CimageView2/2/w/1240)

这里就根据自己的实际情况进行填写就好了。

![配置 Jenkins URL](https://upload-images.jianshu.io/upload_images/14623749-bb23a8bba65ecb4a.png?imageMogr2/auto-orient/strip%7CimageView2/2/w/1240)

根据自己的实际情况去配置 Jenkins URL 就好了，我觉得没有必要配置，就直接选择的默认。

下面 Jenkins 就已经安装成功啦！

![Jenkins 安装成功](https://upload-images.jianshu.io/upload_images/14623749-cee084d19d11e498.png?imageMogr2/auto-orient/strip%7CimageView2/2/w/1240)

现在只能说 Jenkins 已经安装成功了，但是想要使用 Jenkins 来部署项目，我们还得安装以下插件进行配合才可以。

## 必须要安装的插件

- Git Parameter：拉取源代码分支
- Publish Over SSH：将我们打包好的代码推送到目标服务器上
- AnsiColor：向控制台输出添加 ANSI 颜色（也可以不安装）

我们点击右上角的**设置**图标

![去设置页面](https://upload-images.jianshu.io/upload_images/14623749-62129be65da6c8cd.png?imageMogr2/auto-orient/strip%7CimageView2/2/w/1240)

来到插件管理页面

![插件管理](https://upload-images.jianshu.io/upload_images/14623749-56c59cd089f5068e.png?imageMogr2/auto-orient/strip%7CimageView2/2/w/1240)

我们需要先在 `Available plugins` 中搜索 `Git Parameter` 然后进行安装 Git Parameter 插件

![Git Parameter](https://upload-images.jianshu.io/upload_images/14623749-95f4af665788e010.png?imageMogr2/auto-orient/strip%7CimageView2/2/w/1240)

继续安装另外两个插件

![安装 ssh 插件](https://upload-images.jianshu.io/upload_images/14623749-f911b45c614993d3.png?imageMogr2/auto-orient/strip%7CimageView2/2/w/1240)

当我们将以上插件都安装完毕之后，就代表我们部署项目的准备工作已经做完了。下一章节，我将一步一步告诉你如何通过 Pipeline 来进行自动化部署项目，敬请期待吧～