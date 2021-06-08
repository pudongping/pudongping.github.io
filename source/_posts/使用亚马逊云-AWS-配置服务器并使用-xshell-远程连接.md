---
title: 使用亚马逊云 AWS 配置服务器并使用 xshell 远程连接
author: Alex
top: false
hide: false
cover: false
toc: true
mathjax: false
categories: 服务器
tags:
  - Linux
  - 服务器
abbrlink: 34efcb18
date: 2021-06-08 09:19:56
img:
coverImg:
password:
summary: 使用亚马逊云 AWS 配置服务器并使用 xshell 远程连接，如果之前使用亚马逊云配置服务器没有配置成功，那么请按照下面的方式删除掉实例
---

# 使用亚马逊云 AWS 配置服务器并使用 xshell 远程连接

## 如果之前使用亚马逊云配置服务器没有配置成功，那么请按照下面的方式删除掉实例

- [打开当前实例列表](https://us-east-2.console.aws.amazon.com/ec2/home?region=us-east-2#Instances:)

![你自己的实例列表](https://upload-images.jianshu.io/upload_images/14623749-bd64bb68e37736b9.png?imageMogr2/auto-orient/strip%7CimageView2/2/w/1240)

- 终止你想要删除的实例（我看文档说的是，终止实例其实就表示删除了实例，终止实例会删除掉服务器中的文件，但是停止实例不会）

![终止实例](https://upload-images.jianshu.io/upload_images/14623749-fa7c851c14f1977f.png?imageMogr2/auto-orient/strip%7CimageView2/2/w/1240)

![终止实例会删除掉服务器中的所有文件](https://upload-images.jianshu.io/upload_images/14623749-688dde3c0f03d746.png?imageMogr2/auto-orient/strip%7CimageView2/2/w/1240)

![终止后的状态](https://upload-images.jianshu.io/upload_images/14623749-6d0c8540ccaaf4e7.png?imageMogr2/auto-orient/strip%7CimageView2/2/w/1240)

- 如果想彻底删除掉实例，还需要删除掉 「卷」（不手动删除的话，亚马逊云也会过一段时间自动删除）

![选择卷组](https://upload-images.jianshu.io/upload_images/14623749-16fbeb3feda581a5.png?imageMogr2/auto-orient/strip%7CimageView2/2/w/1240)

![实例列表也会过一段时间自动删除](https://upload-images.jianshu.io/upload_images/14623749-8273f1d73ddf2641.png?imageMogr2/auto-orient/strip%7CimageView2/2/w/1240)

## 接下来开始配置新的服务器

- 创建一个新的实例
  ![直接在实例控制台上面启动一个新的实例](https://upload-images.jianshu.io/upload_images/14623749-94887efe4899baba.png?imageMogr2/auto-orient/strip%7CimageView2/2/w/1240)

- 配置服务器
  ![按照这些步骤一步一步的来，选择你自己的自定义配置就可以了](https://upload-images.jianshu.io/upload_images/14623749-a0a7f77087175125.png?imageMogr2/auto-orient/strip%7CimageView2/2/w/1240)

- 选择硬盘空间大小（按需选择就好，如果超过了套餐外的硬盘空间会收费，建议还是先检查自己的套餐最大硬盘空间）

![设置硬盘空间大小](https://upload-images.jianshu.io/upload_images/14623749-15d94915c460e4bf.png?imageMogr2/auto-orient/strip%7CimageView2/2/w/1240)


- 最最最重要的步骤就是，一定要选择密钥对或者直接在这里生成密钥对，并且一定要下载密钥对，因为密钥对只允许下载一次，错过了，等于你这个实例就无法登录了。亚马逊云默认关闭了账号密码连接 ssh 服务，初次连接只允许密钥对。

![选择已生成的密钥对或者直接在这里生成](https://upload-images.jianshu.io/upload_images/14623749-e95c5b055cfa982b.png?imageMogr2/auto-orient/strip%7CimageView2/2/w/1240)

![我没有提前生成密钥对，因此我在这里直接生成一个密钥对](https://upload-images.jianshu.io/upload_images/14623749-b0479e4003995376.png?imageMogr2/auto-orient/strip%7CimageView2/2/w/1240)

![下载密钥对](https://upload-images.jianshu.io/upload_images/14623749-7a22bfdf90b62c72.png?imageMogr2/auto-orient/strip%7CimageView2/2/w/1240)

![打开实例](https://upload-images.jianshu.io/upload_images/14623749-b4845a5e820c5f57.png?imageMogr2/auto-orient/strip%7CimageView2/2/w/1240)

![已经安装好了服务器](https://upload-images.jianshu.io/upload_images/14623749-bc31ca359b59438a.png?imageMogr2/auto-orient/strip%7CimageView2/2/w/1240)

## 接下来讲解使用 xshell 连接亚马逊云服务器

> 建议还是先看一下文档：https://docs.aws.amazon.com/zh_cn/AWSEC2/latest/UserGuide/managing-users.html

- 根据你选择的系统寻找到亚马逊云默认给你创建的账户名

![我选择的是 redhat 系统，因此用户名是 ec2-user 或 root](https://upload-images.jianshu.io/upload_images/14623749-1bfd9833871ee335.png?imageMogr2/auto-orient/strip%7CimageView2/2/w/1240)

- 打开 xshell ，配置用户身份验证

![配置用户身份验证](https://upload-images.jianshu.io/upload_images/14623749-a62480d370da6f50.png?imageMogr2/auto-orient/strip%7CimageView2/2/w/1240)

![导入密钥对](https://upload-images.jianshu.io/upload_images/14623749-7823047ee33c0a5b.png?imageMogr2/auto-orient/strip%7CimageView2/2/w/1240)

![选中密钥对](https://upload-images.jianshu.io/upload_images/14623749-d8f8286d1d98330a.png?imageMogr2/auto-orient/strip%7CimageView2/2/w/1240)

![效果如下](https://upload-images.jianshu.io/upload_images/14623749-78459d37d9fa3827.png?imageMogr2/auto-orient/strip%7CimageView2/2/w/1240)

- 配置连接

![复制共有 DNS](https://upload-images.jianshu.io/upload_images/14623749-46b2bcf9a3277ade.png?imageMogr2/auto-orient/strip%7CimageView2/2/w/1240)

![配置连接](https://upload-images.jianshu.io/upload_images/14623749-3125ba8d0bec81fb.png?imageMogr2/auto-orient/strip%7CimageView2/2/w/1240)

![连接成功](https://upload-images.jianshu.io/upload_images/14623749-6eee19e6b21804b6.png?imageMogr2/auto-orient/strip%7CimageView2/2/w/1240)


## 亚马逊云默认关闭了 root 连接 ssh，接下来讲解如何使用 root 用户登录 ssh

1. 更改 root 用户密码

```bash
sudo passwd root
```

2. 切换到 root 用户

```bash
su - 
或者
su root
```

![修改 root 用户密码](https://upload-images.jianshu.io/upload_images/14623749-31aaa46c330d2e75.png?imageMogr2/auto-orient/strip%7CimageView2/2/w/1240)


3. 修改 sshd_config 配置文件

```bash
vim /etc/ssh/sshd_config
# 如果提示没有 vim 编辑器，则可以直接使用 vi 编辑器
vi /etc/ssh/sshd_config
```

![编辑 sshd 服务主配置文件](https://upload-images.jianshu.io/upload_images/14623749-2e587674dfadf7b9.png?imageMogr2/auto-orient/strip%7CimageView2/2/w/1240)

4. 开启密码验证

```bash
PasswordAuthentication yes
```

![查看是否允许密码验证](https://upload-images.jianshu.io/upload_images/14623749-0cd7175dbad6da90.png?imageMogr2/auto-orient/strip%7CimageView2/2/w/1240)

5. 设定是否允许root管理员直接登录

```bash
PermitRootLogin yes
```

![查看是否允许 root 管理员直接登录](https://upload-images.jianshu.io/upload_images/14623749-76801d3921d2d1b3.png?imageMogr2/auto-orient/strip%7CimageView2/2/w/1240)

6. 重启 sshd 服务

```bash
# 重启 sshd 服务
systemctl restart sshd

# 将 sshd 服务加入到开机启动项中
systemctl enable sshd
```

7. 测试连接

![使用公网 ip 和 root 用户连接 ssh](https://upload-images.jianshu.io/upload_images/14623749-096bb85b62b21053.png?imageMogr2/auto-orient/strip%7CimageView2/2/w/1240)

![使用 root 账号密码连接 ssh](https://upload-images.jianshu.io/upload_images/14623749-8bdb83fbbe69a83e.png?imageMogr2/auto-orient/strip%7CimageView2/2/w/1240)

![至此，使用 root 用户连接成功](https://upload-images.jianshu.io/upload_images/14623749-d29f7e82e877c3e5.png?imageMogr2/auto-orient/strip%7CimageView2/2/w/1240)
