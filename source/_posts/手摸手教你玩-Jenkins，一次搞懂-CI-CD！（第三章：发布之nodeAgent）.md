---
title: 手摸手教你玩 Jenkins，一次搞懂 CI/CD！（第三章：发布之nodeAgent）
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
abbrlink: c6b6ecb5
date: 2026-07-13 13:39:24
img:
coverImg:
password:
summary:
---

Jenkins 相关的教程，我已经出了两期了。

第一章，我讲到了通过 docker 快速部署 Jenkins 以及安装 Jenkins 的必要插件。

第二章，我讲到了通过 sshPublisher 去目标服务器上发布项目。

今天，我来讲讲如何通过代理节点去目标服务器上发布项目。其实这里实现的效果和第二章中实现的效果是一样的，都是去目标服务器上执行发布项目的脚本，进行项目的编译和发布。

个人觉得，二选一，掌握一种方式就好了。

下面，我默认你已经看过了上面我写的两篇文章，因此，一些相同的部分，我就不再赘述。

## 添加节点

1. 来到**系统管理**页面，点击「节点和云管理」

![节点和云管理](https://upload-images.jianshu.io/upload_images/14623749-f428ced91513712e.png?imageMogr2/auto-orient/strip%7CimageView2/2/w/1240)

默认已经有了一个节点，我们点击「New Node」添加新的节点

![添加新的节点](https://upload-images.jianshu.io/upload_images/14623749-9b31e4f9a1db67e5.png?imageMogr2/auto-orient/strip%7CimageView2/2/w/1240)

2. 填写节点名称

![添加节点](https://upload-images.jianshu.io/upload_images/14623749-b5c30a5845e7be5f.png?imageMogr2/auto-orient/strip%7CimageView2/2/w/1240)

3. 填写节点的详细信息

![填写参数](https://upload-images.jianshu.io/upload_images/14623749-38e56e5f56ae0f9d.png?imageMogr2/auto-orient/strip%7CimageView2/2/w/1240)

是不是看到这里的 `Credentials` 有一种莫名的亲切感？是的，这就是在第二章节中，在凭据管理页面死活找不到的添加凭据的地方。

![继续填写参数](https://upload-images.jianshu.io/upload_images/14623749-6a30b9eb499d9690.png?imageMogr2/auto-orient/strip%7CimageView2/2/w/1240)

添加新的凭据

![添加凭据](https://upload-images.jianshu.io/upload_images/14623749-529c31ccdca934d5.png?imageMogr2/auto-orient/strip%7CimageView2/2/w/1240)

这里需要添加的是目标服务器上，用于 ssh 登录的用户 jenkinsjob 的账号和密码

![添加账号和密码](https://upload-images.jianshu.io/upload_images/14623749-50439b2f8c58ca21.png?imageMogr2/auto-orient/strip%7CimageView2/2/w/1240)

选择刚刚新添加的凭据

![选中凭据](https://upload-images.jianshu.io/upload_images/14623749-87b1ef8bcf6e73c7.png?imageMogr2/auto-orient/strip%7CimageView2/2/w/1240)

填写完了之后就直接点击保存。

## 查看节点

回到节点列表，我们可以看到刚刚我们添加的节点已经添加成功了。

![新添加的节点](https://upload-images.jianshu.io/upload_images/14623749-defd54737dcae04a.png?imageMogr2/auto-orient/strip%7CimageView2/2/w/1240)

我们点击刚刚新添加的节点，进入节点详情，然后点击右上角的「上线节点」，看节点是否能够上线成功

![上线节点](https://upload-images.jianshu.io/upload_images/14623749-d87f67449b29bf82.png?imageMogr2/auto-orient/strip%7CimageView2/2/w/1240)

## 上线节点

一般情况下而言，可能多半就会报如下类似的错误：

```bash
bash: line 1: java: command not found
Agent JVM has terminated. Exit code=127
```

原因是：Jenkins Agent 必须要能运行 Java，但你的服务器上没有 java 命令。

当 Jenkins 通过 SSH 连接 Agent 时，它会自动执行：

```bash
java -jar remoting.jar
```

也就是让 Agent 使用 Java 运行 Jenkins 的 remoting 进程。**如果服务器上没有 Java，Agent 永远无法上线。**

那么，如何解决这个问题呢？

直接在目标服务器上安装 java 环境就好了。

这里假设目标服务器为 ubuntu 系统：

### 安装 java 环境

```bash
sudo apt update
sudo apt install -y openjdk-17-jdk
```

验证安装

```bash
java -version

# 类似以下输出即可
#openjdk version "17.0.16" 2025-07-15
#OpenJDK Runtime Environment (build 17.0.16+8-Ubuntu-0ubuntu124.04.1)
#OpenJDK 64-Bit Server VM (build 17.0.16+8-Ubuntu-0ubuntu124.04.1, mixed mode, sharing)

javac -version
# 类似以下输出即可
#javac 17.0.16
```

然后我们重新上线节点，可以发现节点能够正常上线了。

并且我们在目标服务器上通过 jenkinsjob 账号登录之后，我们可以看到当时设置的远程目录下会存在

![远程目录下会存在这两个](https://upload-images.jianshu.io/upload_images/14623749-877d9c7421ed37ae.png?imageMogr2/auto-orient/strip%7CimageView2/2/w/1240)

## 编写 jenkinsfile

这里还是老规矩，直接贴上我写好的 jenkinsfile 文件内容供各位看官老爷们参考

```bash
#!/usr/bin/env groovy Jenkinsfile

pipeline {
    agent { label 'node_ssh_jenkinsjob' }

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
                        sh 'cd /home/jenkinsjob/deploy_sample && sudo ./deploy.sh'
                    }
                }
            }
        }

    }


}
```

我们可以观察一下，第三章中的 jenkinsfile 和第二章中的 jenkinsfile 的差异性

差别最大的莫过于两点。

### 第一点

agent 发生了改变，这里由 `agent any` 改成了 `agent { label 'node_ssh_jenkinsjob' }`

需要注意的是，`label` 后面填写的是上面添加的节点的名称。

### 第二点

执行脚本的时候，这里直接在 `sh '……………………'` 中间就写了命令。相对来说简化了不少。

---

我们已经写好了 jenkinsfile 那么，添加任务流就和第二章节中是一样的了。赶紧来试试看吧。


![控制台输出](https://upload-images.jianshu.io/upload_images/14623749-ab069b90091c966b.png?imageMogr2/auto-orient/strip%7CimageView2/2/w/1240)

并且当我们执行工作流时，我们发现这个位置确实就是在我们设置的节点上运行的。

![执行成功](https://upload-images.jianshu.io/upload_images/14623749-7493db7cd34061e6.png?imageMogr2/auto-orient/strip%7CimageView2/2/w/1240)


## 总结

如果采用第二章中 sshPublisher 的方式去访问目标服务器，那么就一定得安装 `Publish Over SSH` 插件。并且在 Jenkins 上写执行命令时会稍微复杂一点点。

如果采用本章中添加代理节点的方式去访问目标服务器，那么在目标服务器上就必须得安装 Java 运行环境。

两种方式都可以实现访问目标服务器，并直接在目标服务器上执行部署项目的脚本命令。

## 其他

Q：我们怎么可以查看到一个工作流的 jenkinsfile 内容呢？  
A：可以直接在 `/var/jenkins_home/jobs` 目录下找到对应的工作名称，然后查看 `config.xml` 文件

Q：如何查看到代理节点数据？
A：直接查看 `/var/jenkins_home/nodes` 目录


好了，以上就是 Jenkins 教程中的全部内容了。这 3 篇教程中，完全不会讲解枯燥无味的概念性东西，所有的内容都是可以直接上手去用的。希望对你有所帮助。

创作不易，希望可以帮忙点赞并关注支持一下，谢谢。