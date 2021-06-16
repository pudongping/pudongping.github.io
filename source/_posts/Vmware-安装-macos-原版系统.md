---
title: Vmware 安装 macos 原版系统
author: Alex
top: false
hide: false
cover: false
toc: true
mathjax: false
categories: macOS
tags:
  - macOS
  - 虚拟机
abbrlink: 17b85fe6
date: 2021-06-16 09:16:29
img:
coverImg:
password:
summary:
---


# VMware WorkStation 安装 macOS 原版系统

## 文献参考
- [手把手教你win 10 VMware 15 安装 MAC OS 10.15原版系统](https://www.bilibili.com/video/BV1a7411e7Pq)
- [VMware 解锁器](https://github.com/paolo-projects/unlocker/releases)
- [官方系统下载](https://updates.cdn-apple.com/2020/macos/001-48329-20200924-122a4260-eb65-4492-b08e-f69dfa22c2dc/SecUpd2020-005Mojave.dmg)
- [镜像下载地址](https://www.mediafire.com/file/7wm2251an4c2n64/macOS_Catalina_ISO_By_Techbland.iso/file)
- [提供各版本的苹果电脑macOS系统镜像下载 ，支持百度网盘/独立服务器/迅雷地址下载](https://www.applex.net/pages/macos/)

## 安装前必须要准备的软件
- VMware workstation 虚拟机
- VMware 解锁器
- macOS 懒人包系统 （ .cdr 文件）

以上软件可以通过访问我的百度网盘获取，需要注意的是，下载下来的 `.cdr` 懒人包系统，最好要通过 `md5` 校验。


以下为网盘链接

```

链接：https://pan.baidu.com/s/1NMHD6i8FIgYaqcPVljbl8w 
提取码：jcr3 

```

> Windows 系统下可以执行以下命令来核对 md5


```sh

certutil -hashfile "cdr文件的绝对路径" md5

# 比如

certutil -hashfile "F:\Vmware Installer macOS Mojave(18G103).cdr" md5

``` 

## 安装步骤如下

- 下载 VMware 解锁器 `unlocker.zip` 并解压，并切记一定要关闭掉 VMware 的所有进程，如果不知道如何关闭掉 VMware 的所有进程，可以考虑直接关机重启，然后再进行以下操作

![解压 VMware 解锁器](https://upload-images.jianshu.io/upload_images/14623749-5d55ecd7da49b4f8.png?imageMogr2/auto-orient/strip%7CimageView2/2/w/1240)

- 对着解压后的 `win-install.cmd` 文件，右键 - 以管理员身份运行。

![右键 - 以管理员身份运行](https://upload-images.jianshu.io/upload_images/14623749-c1094950a7bcb3ea.png?imageMogr2/auto-orient/strip%7CimageView2/2/w/1240)

- 此时解锁器会自动下载相应的文件，不用去理会

![解锁器正在下载相应的解锁文件](https://upload-images.jianshu.io/upload_images/14623749-b55be1c48fe20e9b.png?imageMogr2/auto-orient/strip%7CimageView2/2/w/1240)

- 字条走完之后会自动关闭掉 `cmd` 命令窗口

![解锁器软件安装完毕后，会自动关闭掉 cmd 窗口](https://upload-images.jianshu.io/upload_images/14623749-1469e628bb22865d.png?imageMogr2/auto-orient/strip%7CimageView2/2/w/1240)

- 接下来我们来验证解锁器是否安装好，打开 VMware 虚拟机，尝试新建一台虚拟机

![image.png](https://upload-images.jianshu.io/upload_images/14623749-dc120ff66ee2143f.png?imageMogr2/auto-orient/strip%7CimageView2/2/w/1240)

- 选择稍后安装系统

![image.png](https://upload-images.jianshu.io/upload_images/14623749-3bdecc977abad802.png?imageMogr2/auto-orient/strip%7CimageView2/2/w/1240)

- 选择 `Apple Mac OS` 相应的选项，此时我们可以看到能够选择相应的系统版本，此时需要注意的是 **安装什么版本的系统，就要选择相应的版本，否则安装系统时，页面会直接卡死**

![image.png](https://upload-images.jianshu.io/upload_images/14623749-c0383da1fac1d780.png?imageMogr2/auto-orient/strip%7CimageView2/2/w/1240)

- 此时我们能够确定解锁器已经安装好了，如果此时无法看到 `macOS` 系统相关的选项，那么就需要自己反复多次的安装解锁器，直到能够看到 `macOS` 系统相关的选项才算成功！此时，正式开启安装 `macOS` 系统之旅！

- 以下步骤没有太多需要说的，按照截图一步一步的来操作就好了。

![image.png](https://upload-images.jianshu.io/upload_images/14623749-0d5e6bb7fcbf5693.png?imageMogr2/auto-orient/strip%7CimageView2/2/w/1240)

![image.png](https://upload-images.jianshu.io/upload_images/14623749-bdf66cfcb46d334f.png?imageMogr2/auto-orient/strip%7CimageView2/2/w/1240)

![image.png](https://upload-images.jianshu.io/upload_images/14623749-470f7e95996b8e95.png?imageMogr2/auto-orient/strip%7CimageView2/2/w/1240)

- 这里需要注意的是，我们采用的系统是 `.cdr` 格式的懒人包，我们需要选择 `所有文件` 才可以看到 `.cdr` 文件。

![image.png](https://upload-images.jianshu.io/upload_images/14623749-82cc62377031414c.png?imageMogr2/auto-orient/strip%7CimageView2/2/w/1240)

![image.png](https://upload-images.jianshu.io/upload_images/14623749-76561873201b4868.png?imageMogr2/auto-orient/strip%7CimageView2/2/w/1240)

![image.png](https://upload-images.jianshu.io/upload_images/14623749-a999d47eac24da6f.png?imageMogr2/auto-orient/strip%7CimageView2/2/w/1240)

- 我这里解锁之后最高只能达到 `10.14` 的版本，因此最后只能安装 `macOS 10.14` 的系统，但是操作步骤是一样的，虽然我也不知道为什么我这里最高只能是 `10.14` 的版本，如果你知道，欢迎给我留言分享一下。但是一定要切记，在这里你选择的是什么版本，你就必须得安装什么版本的系统，否则无法安装的。

![image.png](https://upload-images.jianshu.io/upload_images/14623749-5db13a5dddd5913e.png?imageMogr2/auto-orient/strip%7CimageView2/2/w/1240)

- 你也可以自己选择安装路径，这个没有多大影响的。

![image.png](https://upload-images.jianshu.io/upload_images/14623749-32a39bb01028420a.png?imageMogr2/auto-orient/strip%7CimageView2/2/w/1240)

![image.png](https://upload-images.jianshu.io/upload_images/14623749-57a02042db810fe1.png?imageMogr2/auto-orient/strip%7CimageView2/2/w/1240)

![image.png](https://upload-images.jianshu.io/upload_images/14623749-63269f2bf437b483.png?imageMogr2/auto-orient/strip%7CimageView2/2/w/1240)

![image.png](https://upload-images.jianshu.io/upload_images/14623749-bd5063e66c347b8b.png?imageMogr2/auto-orient/strip%7CimageView2/2/w/1240)

![image.png](https://upload-images.jianshu.io/upload_images/14623749-27ab34b1c5910913.png?imageMogr2/auto-orient/strip%7CimageView2/2/w/1240)

![image.png](https://upload-images.jianshu.io/upload_images/14623749-4e4c3d84cd2d81d7.png?imageMogr2/auto-orient/strip%7CimageView2/2/w/1240)

![image.png](https://upload-images.jianshu.io/upload_images/14623749-fcd64dfa495db579.png?imageMogr2/auto-orient/strip%7CimageView2/2/w/1240)

![image.png](https://upload-images.jianshu.io/upload_images/14623749-8dec6384e39c5f0a.png?imageMogr2/auto-orient/strip%7CimageView2/2/w/1240)

![image.png](https://upload-images.jianshu.io/upload_images/14623749-f6c011453fa32db5.png?imageMogr2/auto-orient/strip%7CimageView2/2/w/1240)

- 这里的我们需要选择 `磁盘工具` 先将磁盘抹掉一下

![image.png](https://upload-images.jianshu.io/upload_images/14623749-18c19147c0287b0f.png?imageMogr2/auto-orient/strip%7CimageView2/2/w/1240)

- 一定要注意，只能选择刚刚我们分配给 VMware 的那块盘，其他的不要去动

![image.png](https://upload-images.jianshu.io/upload_images/14623749-d6502e6942507cf8.png?imageMogr2/auto-orient/strip%7CimageView2/2/w/1240)

![image.png](https://upload-images.jianshu.io/upload_images/14623749-a56f6d65c8af8e66.png?imageMogr2/auto-orient/strip%7CimageView2/2/w/1240)

- 也可以直接修改硬盘的名称

![image.png](https://upload-images.jianshu.io/upload_images/14623749-29182a6ba51f3d14.png?imageMogr2/auto-orient/strip%7CimageView2/2/w/1240)

![image.png](https://upload-images.jianshu.io/upload_images/14623749-0bb2fe22c24127fb.png?imageMogr2/auto-orient/strip%7CimageView2/2/w/1240)

![image.png](https://upload-images.jianshu.io/upload_images/14623749-dc6da02e40630c46.png?imageMogr2/auto-orient/strip%7CimageView2/2/w/1240)

- 接下来就可以正式安装系统了

![image.png](https://upload-images.jianshu.io/upload_images/14623749-6b4f9b0dc8dd5f59.png?imageMogr2/auto-orient/strip%7CimageView2/2/w/1240)

![image.png](https://upload-images.jianshu.io/upload_images/14623749-20a8500578e77a4f.png?imageMogr2/auto-orient/strip%7CimageView2/2/w/1240)

![image.png](https://upload-images.jianshu.io/upload_images/14623749-4a5e4d2f6f873237.png?imageMogr2/auto-orient/strip%7CimageView2/2/w/1240)

![image.png](https://upload-images.jianshu.io/upload_images/14623749-ebf3a3aa95ad9e23.png?imageMogr2/auto-orient/strip%7CimageView2/2/w/1240)

![image.png](https://upload-images.jianshu.io/upload_images/14623749-10f3e63d7698153b.png?imageMogr2/auto-orient/strip%7CimageView2/2/w/1240)

![image.png](https://upload-images.jianshu.io/upload_images/14623749-6a29ad902ca42f85.png?imageMogr2/auto-orient/strip%7CimageView2/2/w/1240)

![image.png](https://upload-images.jianshu.io/upload_images/14623749-858e85800375aa9d.png?imageMogr2/auto-orient/strip%7CimageView2/2/w/1240)

![image.png](https://upload-images.jianshu.io/upload_images/14623749-4f7d8621579f1e35.png?imageMogr2/auto-orient/strip%7CimageView2/2/w/1240)

![image.png](https://upload-images.jianshu.io/upload_images/14623749-6d1be7fab284223e.png?imageMogr2/auto-orient/strip%7CimageView2/2/w/1240)

![image.png](https://upload-images.jianshu.io/upload_images/14623749-d91052596f1f6a5a.png?imageMogr2/auto-orient/strip%7CimageView2/2/w/1240)

- 安装到这里之后，我们需要到物理机上面设置一下网络，这样才能够保证虚拟机能够访问网络

![image.png](https://upload-images.jianshu.io/upload_images/14623749-89c913b1a93e35d2.png?imageMogr2/auto-orient/strip%7CimageView2/2/w/1240)

![image.png](https://upload-images.jianshu.io/upload_images/14623749-c76b0e8d339fa9b4.png?imageMogr2/auto-orient/strip%7CimageView2/2/w/1240)

![image.png](https://upload-images.jianshu.io/upload_images/14623749-5a6b3cfa2e4b5ac4.png?imageMogr2/auto-orient/strip%7CimageView2/2/w/1240)

![image.png](https://upload-images.jianshu.io/upload_images/14623749-58f85f0769d2d806.png?imageMogr2/auto-orient/strip%7CimageView2/2/w/1240)

![image.png](https://upload-images.jianshu.io/upload_images/14623749-bb7e63edd1ea925b.png?imageMogr2/auto-orient/strip%7CimageView2/2/w/1240)

![image.png](https://upload-images.jianshu.io/upload_images/14623749-f0e5481918b08eae.png?imageMogr2/auto-orient/strip%7CimageView2/2/w/1240)

![image.png](https://upload-images.jianshu.io/upload_images/14623749-b8cbaab5b48a8f5d.png?imageMogr2/auto-orient/strip%7CimageView2/2/w/1240)

![image.png](https://upload-images.jianshu.io/upload_images/14623749-638a71f2bb30082f.png?imageMogr2/auto-orient/strip%7CimageView2/2/w/1240)

![image.png](https://upload-images.jianshu.io/upload_images/14623749-5dc33873f724a835.png?imageMogr2/auto-orient/strip%7CimageView2/2/w/1240)

![image.png](https://upload-images.jianshu.io/upload_images/14623749-0e2db0987876efb6.png?imageMogr2/auto-orient/strip%7CimageView2/2/w/1240)

![image.png](https://upload-images.jianshu.io/upload_images/14623749-537f2118a2179f0e.png?imageMogr2/auto-orient/strip%7CimageView2/2/w/1240)

![image.png](https://upload-images.jianshu.io/upload_images/14623749-0c519f187338c657.png?imageMogr2/auto-orient/strip%7CimageView2/2/w/1240)

![image.png](https://upload-images.jianshu.io/upload_images/14623749-47342c1c550ac55e.png?imageMogr2/auto-orient/strip%7CimageView2/2/w/1240)

![image.png](https://upload-images.jianshu.io/upload_images/14623749-cc5277d7cdcd2443.png?imageMogr2/auto-orient/strip%7CimageView2/2/w/1240)

![image.png](https://upload-images.jianshu.io/upload_images/14623749-d28b63af01954e7d.png?imageMogr2/auto-orient/strip%7CimageView2/2/w/1240)

- 此时我们在 VMware 中就已经将 macOS 系统安装完毕了。此时我们弹出一下我们的系统镜像

![image.png](https://upload-images.jianshu.io/upload_images/14623749-b9b6a3d9fc18f1d1.png?imageMogr2/auto-orient/strip%7CimageView2/2/w/1240)

- 继续安装 VMware tools ，这样能够自动适配我们物理机的分辨率

![image.png](https://upload-images.jianshu.io/upload_images/14623749-cc8ae8f39937e312.png?imageMogr2/auto-orient/strip%7CimageView2/2/w/1240)

![image.png](https://upload-images.jianshu.io/upload_images/14623749-22bf2dda46ea0caf.png?imageMogr2/auto-orient/strip%7CimageView2/2/w/1240)

![image.png](https://upload-images.jianshu.io/upload_images/14623749-33acc9ad3a16c8cd.png?imageMogr2/auto-orient/strip%7CimageView2/2/w/1240)

![image.png](https://upload-images.jianshu.io/upload_images/14623749-faf3da4dd22cb119.png?imageMogr2/auto-orient/strip%7CimageView2/2/w/1240)

![image.png](https://upload-images.jianshu.io/upload_images/14623749-20f5ee8ff0c04453.png?imageMogr2/auto-orient/strip%7CimageView2/2/w/1240)

![image.png](https://upload-images.jianshu.io/upload_images/14623749-0e0e12068dce26c1.png?imageMogr2/auto-orient/strip%7CimageView2/2/w/1240)

按照以上的截图一步一步操作，我相信问题应该不会很大的。
