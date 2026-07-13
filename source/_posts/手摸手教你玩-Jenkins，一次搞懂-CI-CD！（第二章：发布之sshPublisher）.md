---
title: 手摸手教你玩 Jenkins，一次搞懂 CI/CD！（第二章：发布之sshPublisher）
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
abbrlink: b67d264e
date: 2026-07-13 13:38:47
img:
coverImg:
password:
summary:
---

上一章节，我讲到了如何通过 docker 去安装 Jenkins，并且在 Jenkins 中安装了一些常用的插件用于做项目部署的前期准备，这一篇文章，咱们就直接开始进行项目发布。

看这篇文章之前，首先需要确定你已经安装了以下插件：

- Git Parameter：拉取源代码分支
- Publish Over SSH：将我们打包好的代码推送到目标服务器上
- AnsiColor：向控制台输出添加 ANSI 颜色（也可以不安装）

> 如果不清楚如何安装的，建议先查看 《手摸手教你玩 Jenkins，一次搞懂 CI/CD！（第一章：部署）》这篇文章。

## 一些约定

在讲解如何通过 Jenkins 发布项目之前，我默认你已经有了自己的 Git 项目仓库。这里我为了演示，假定我的 Git 项目仓库地址为： `http://127.0.0.1:8099/demo/demo-project.git` （实际上根本不存在此 Git 仓库）

## 前期准备

需要明确的一点是：虽然我们可以直接通过在 Jenkins 中编译项目，然后将编译后的成品发布到目标服务器上进行部署，但是我的做法是**不建议这么干**。因为这么干，就太依赖于 Jenkins 了。

我建议的做法是：直接在目标服务器上写好发布脚本，然后通过 Jenkins 去直接执行目标服务器上的发布脚本就好了。

这样的做法好处是，就算是哪天我不用 Jenkins，我也可以手动去目标服务器上执行发布脚本进行发布即可，并且对于发布复杂的项目来说，也方便自己调试发布脚本。

### 发布脚本

我们假设在目标服务器上存在 `deploy.sh` 脚本，这里为了演示方便，我就简单写了一个时间输出。

> 我们假设只要在 Jenkins 中执行了这个脚本，就能够在目标服务器上成功发布项目。当然，实际情况应该要根据自己的项目来决定。我这里仅仅为了最简单的演示。

`deploy.sh` 脚本文件中

```bash
#!/bin/bash

version=`date +%Y%m%d.%H%M`

echo '当前时间：'$version
```

### 独立的登录用户

既然是需要执行脚本来发布项目，那么，是不是需要有登录用户去执行才行呢？是的，总不能用 root 用户去执行吧？

这里我们需要在服务器上创建一个新的登录用户，这个登录用户就专门用户通过 ssh 登录，然后再去执行 `deploy.sh` 脚本。

```bash
# 添加用户名为 jenkinsjob 的用户
sudo useradd -m -s /bin/bash jenkinsjob

# 修改 jenkinsjob 用户的密码
sudo passwd jenkinsjob
# 比如密码就为 123456

# 查看用户是否建成功
cat /etc/passwd
# 如果建错了，可以删除
sudo userdel jenkinsjob
```

并且，还有一个问题十分重要：**如果有些服务需要临时用到 root 权限，该怎么办？**

我们都知道是可以通过在命令前面添加 `sudo` ，但是你是否想过当我们在执行命令前面添加 `sudo` 之后，我们也是需要在交互界面输入当前用户的密码的。可是我们是通过 Jenkins 去执行的 `deploy.sh` 脚本，我们压根儿就无法进入交互界面，那么，我们该怎么处理呢？

答案是：**通过为登录用户添加免密环境**。

### 给登录用户添加免密环境

我们以 jenkinsjob 用户为例，给他设置免密环境。这样当 jenkinsjob 用户执行 `sudo` 命令时，就不需要再输入自己的账号密码了。

编辑 `sudoers` 文件：

```bash
# 默认是用的 nano 编辑器
sudo visudo

# 或者通过 vim 编辑器（这里设置了临时通过 vim 编辑器去编辑）
sudo EDITOR=vim visudo
```

添加以下行

```bash
# jenkinsjob 为用户名
jenkinsjob ALL=(ALL) NOPASSWD: ALL

# 或者只为特定命令设置免密码
jenkinsjob ALL=(ALL) NOPASSWD: /usr/bin/apt, /bin/systemctl
```

编辑免密码环境之后，需要执行以下命令进行测试 sudoers 文件是否有错误

```bash
sudo visudo -c
```

## 开干

### 配置 ssh 登录用户

1. 进入**系统配置**页面

![系统配置](https://upload-images.jianshu.io/upload_images/14623749-32e397c59fc4e97d.png?imageMogr2/auto-orient/strip%7CimageView2/2/w/1240)

2. 找到 **Publish over SSH** 下方的 **SSH Servers** 然后点击「新增」按钮

![添加 ssh](https://upload-images.jianshu.io/upload_images/14623749-0cd6ea1118a4fc21.png?imageMogr2/auto-orient/strip%7CimageView2/2/w/1240)

然后设置 ssh 登录用户

- **Name**：需要给新添加的 ssh server 设置一个名称。**这个名称很重要，在 jenkinsfile 文件中会用到的**
- **Hostname**：目标服务器的 IP 地址
- **Username**：登录目标服务器的账号名称
- **Remote Directory**：登录目标服务器的操作目录

![设置 ssh 登录用户](https://upload-images.jianshu.io/upload_images/14623749-4ec139bdf17a2bd4.png?imageMogr2/auto-orient/strip%7CimageView2/2/w/1240)

3. 验证 ssh 连接是否生效

设置完了之后一定要点击 **Test Configuration** 进行验证，当出现了 **Success** 就表示已经设置成功了。

![验证 ssh 生效](https://upload-images.jianshu.io/upload_images/14623749-a667e49724b840f0.png?imageMogr2/auto-orient/strip%7CimageView2/2/w/1240)

最后一定不要忘记点击 **Save** 按钮进行保存！

### 配置 Git 用户凭据

进入到凭据管理页面

![凭据管理页面](https://upload-images.jianshu.io/upload_images/14623749-a165c3168601af83.png?imageMogr2/auto-orient/strip%7CimageView2/2/w/1240)

我们会发现很难找到添加凭据的按钮

![找不到添加凭据的按钮](https://upload-images.jianshu.io/upload_images/14623749-6bc732b2063d462c.png?imageMogr2/auto-orient/strip%7CimageView2/2/w/1240)

其实添加凭据，一般都是通过 Jenkins UI 进行添加的，但是我演示的步骤不通过 Jenkins UI 来设置发布流程，我们通过编写 Jenkinsfile 来编写发布流程，所以需要通过直接访问下面的 URL Path 来进入添加凭据的页面。

![添加凭据页面](https://upload-images.jianshu.io/upload_images/14623749-66927da2e1046a8f.png?imageMogr2/auto-orient/strip%7CimageView2/2/w/1240)

可以直接通过访问 `/manage/credentials/store/system/domain/_/` URL Path 来进入到 `Add Credentials` 页面。

添加新的凭据时，直接将用户名和密码就设置成自己 Git 服务器的账号和密码。

![设置git账号凭据](https://upload-images.jianshu.io/upload_images/14623749-6b421ba82baccea4.png?imageMogr2/auto-orient/strip%7CimageView2/2/w/1240)

设置凭据的时候，ID 选项，我们可以直接为空。

当我们设置好了账号和密码之后，点击「Create」

![全局凭证页面](https://upload-images.jianshu.io/upload_images/14623749-99f5176acafeedc1.png?imageMogr2/auto-orient/strip%7CimageView2/2/w/1240)

当我们进入到**全局凭据**页面后，我们会发现有一个 **ID** 项，请留意这里的 **ID** 这个在后面我们会用到的。

至此，我们所有的前期准备已经做好了，接下来，我们开始编写 jenkinsfile 文件。

## 编写 jenkinsfile

这里，我直接将我的 jenkinsfile 贴进来，以供各位看官老爷们参考：

```bash
#!/usr/bin/env groovy Jenkinsfile

pipeline {
    agent any

    options {
        ansiColor('xterm')
    }

    parameters {
        booleanParam(
            name: 'Select_demo_project',
            description: '可以自己填写一下备注，需要注意 name 字段中，最好不要出现横杠 -'
        )

        gitParameter(
            name: 'Git_Branch_demo_project',
            branch: '',
            branchFilter: 'demo_project/(.*)',
            defaultValue: 'master',
            quickFilterEnabled: true,
            selectedValue: 'NONE',
            sortMode: 'NONE',
            tagFilter: '*',
            type: 'PT_BRANCH_TAG',
            useRepository: 'demo-project',
            description: '需要注意 gitParameter.name 字段中，最好不要出现横杠 -'
        )
    }

    stages {

        stage('Step1：初始化，确认参数') {
            steps {
                catchError(buildResult: 'SUCCESS', stageResult: 'FAILURE') {
                    checkout([
                        $class: 'GitSCM',
                        branches: [[name: '*/master']],
                        extensions: [],
                        userRemoteConfigs: [[
                            credentialsId: '86b1ba1c-db25-47b2-b50e-e58516641574',
                            name: 'demo-project',
                            url: 'http://127.0.0.1:8099/demo/demo-project.git'
                        ]]
                    ])
                }
            }
        }

        stage('Step2：发布服务') {
            steps {
                script {
                    if (Select_demo_project == 'true') {
                        echo "----------------构建参数--------------------"
                        echo "git分支：${Git_Branch_demo_project} ---> Select_demo_project:${Select_demo_project}"

                        sshPublisher(
                            publishers: [
                                sshPublisherDesc(
                                    configName: 'target_ssh_jenkinsjob',
                                    transfers: [
                                        sshTransfer(
                                            cleanRemote: false,
                                            excludes: '',
                                            execCommand: '''
                                                cd /home/jenkinsjob/deploy_sample && \
                                                sudo ./deploy.sh
                                            ''',
                                            execTimeout: 1200000,
                                            flatten: false,
                                            makeEmptyDirs: false,
                                            noDefaultExcludes: false,
                                            patternSeparator: '[, ]+',
                                            remoteDirectory: '',
                                            remoteDirectorySDF: false,
                                            removePrefix: '',
                                            sourceFiles: ''
                                        )
                                    ],
                                    usePromotionTimestamp: false,
                                    useWorkspaceInPromotion: false,
                                    verbose: true
                                )
                            ]
                        )
                    }
                }
            }
        }

    }


}
```

## 通过 jenkinsfile 来添加 job

进入首页，添加 job

![Create a job](https://upload-images.jianshu.io/upload_images/14623749-b5ab1c8cb8df545f.png?imageMogr2/auto-orient/strip%7CimageView2/2/w/1240)

然后输入任务名称，再点击**流水线 （pipeline）**

![通过 Pipeline script 来添加任务](https://upload-images.jianshu.io/upload_images/14623749-bc2786b74376d602.png?imageMogr2/auto-orient/strip%7CimageView2/2/w/1240)

直接将页面滑到最下面，选择 **Pipeline script** 然后将 jenkinsfile 中的内容直接贴进来（其他的选项都保持默认就好了，仅仅粘贴 jenkinsfile 内容就好）

![pipeline script](https://upload-images.jianshu.io/upload_images/14623749-22aaa872fe853b78.png?imageMogr2/auto-orient/strip%7CimageView2/2/w/1240)

保存之后，点击**立即构建**

![立即构建](https://upload-images.jianshu.io/upload_images/14623749-b1f9ff5222d6ed72.png?imageMogr2/auto-orient/strip%7CimageView2/2/w/1240)

当我们直接构建之后，可以看到构建日志可能报错了

![查看构建日志](https://upload-images.jianshu.io/upload_images/14623749-ee81ccf7ab590d36.png?imageMogr2/auto-orient/strip%7CimageView2/2/w/1240)

查看控制台输出内容

![查看构建日志内容](https://upload-images.jianshu.io/upload_images/14623749-dba315db4a56e40e.png?imageMogr2/auto-orient/strip%7CimageView2/2/w/1240)

其实，这是正常的。因为我提供的 jenkinsfile 中是需要传参数的，我们直接点击**立即构建**没有给脚本传参数，所以脚本无法执行，就报错了。

当我们再次回到构建页面时，我们可以看到**立即构建**变成了**Build with Parameters**

![Build with Parameters](https://upload-images.jianshu.io/upload_images/14623749-95a70a3a1f55b335.png?imageMogr2/auto-orient/strip%7CimageView2/2/w/1240)

之所以在 jenkinsfile 中设定了需要传参，那是因为在实际开发中，我们一次性可能不止需要发某一个项目，可能我们需要同时多个项目一起发布，并且我们还需要发布不同的项目分支。

![灵活的构建项目](https://upload-images.jianshu.io/upload_images/14623749-463869332becefde.png?imageMogr2/auto-orient/strip%7CimageView2/2/w/1240)

当我们选择了我们需要发布的项目和分支之后，我们再次点击「Build」

![选择需要发布的项目和分支](https://upload-images.jianshu.io/upload_images/14623749-bf8d50fd06de27e6.png?imageMogr2/auto-orient/strip%7CimageView2/2/w/1240)

然后我们再查看发布日志，我们就可以看到已经登录到目标服务器上，并且已经执行了 `deploy.sh` 脚本

![执行了 deploy.sh 脚本](https://upload-images.jianshu.io/upload_images/14623749-f33e8d119ae69c67.png?imageMogr2/auto-orient/strip%7CimageView2/2/w/1240)

至此，我们通过 sshPublisher 的方式去发布项目已经顺利成功了！

> 本篇文章为了方便小白用户更好的理解流程，截图较多。只要你跟着步骤一步一步的来，相信你也会很快掌握 Jenkins 的发布流程的。创作不易，希望可以帮忙点赞关注一下，感谢。

下一篇，我会继续讲解通过代理节点登录目标服务器，然后发布项目的步骤。敬请期待～