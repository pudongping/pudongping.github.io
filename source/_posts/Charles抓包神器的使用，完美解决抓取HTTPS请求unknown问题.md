---
title: Charles抓包神器的使用，完美解决抓取HTTPS请求unknown问题
author: Alex
top: false
hide: false
cover: false
toc: true
mathjax: false
categories: 抓包工具
tags:
  - 抓包工具
  - MacOS
  - Charles
abbrlink: 1eac0236
date: 2024-06-12 18:06:34
img:
coverImg:
password:
summary:
---

在 Windows 上可能比较好用的抓包工具是 fidder，当然，在 Mac 上也有一款抓包神器不输 fiddler，那就是今天的主角—— Charles。

软件的安装过程就不介绍了，只要自己下载好了软件，安装过程就是傻瓜式的操作，非常简单。

今天主要介绍的是如何配置 HTTP 和 HTTPS。有不少童鞋在抓取 HTTPS 请求时，会出现 `unknown` 无法解析的情况，那么，遇到这种情况，我们该如何处理呢？

这篇文章将为你介绍整个配置过程，来，继续往下看吧！

> 我这里演示的是：
> Charles 版本为：4.6.6  
MacOS 系统版本为：Sonoma 14.5  
iOS 系统版本为：17.3

## 安装证书

安装证书这一步是抓取 HTTPS 请求的关键所在，**包括 PC 端和手机端都需要安装证书**。

### PC 端

我们需要先打开 `Charles` 软件，然后在菜单栏中依次选择：`Help` -> `SSL Proxying` -> `Install Charles Root Certificate` 然后直接点击，将证书安装到我们的电脑上。

![安装证书](https://upload-images.jianshu.io/upload_images/14623749-8a490511ae8a2661.png?imageMogr2/auto-orient/strip%7CimageView2/2/w/1240)

将证书安装完毕之后，我们需要打开**钥匙串访问**

![打开「钥匙串访问」](https://upload-images.jianshu.io/upload_images/14623749-a26501fc5f4396e5.png?imageMogr2/auto-orient/strip%7CimageView2/2/w/1240)

当我们打开**钥匙串访问**后，找到「系统钥匙串」-> 「系统」-> 「证书」-> 「Charles Proxy CA……」

![找到证书](https://upload-images.jianshu.io/upload_images/14623749-c389e595b039a474.png?imageMogr2/auto-orient/strip%7CimageView2/2/w/1240)

我们可以看到这个证书默认是**不被信任的**，此时我们需要将其设置为信任。

![设置证书为“始终信任”](https://upload-images.jianshu.io/upload_images/14623749-e1684b7fff48320e.png?imageMogr2/auto-orient/strip%7CimageView2/2/w/1240)

1. 我们直接对着“Charles Proxy ……”开头的证书直接**双击**然后就会出现一个弹窗；
2. 此弹窗中有一个**信任**，默认是闭合的，此时我们需要点击一下，进行展开；
3. 展开后我们可以看到有一项“使用此证书时”，我们将其改为**始终信任**；
4. 然后关闭此弹窗就可以了。不放心的童鞋可以再次打开这个弹窗做一个验证。

![需要输入密码](https://upload-images.jianshu.io/upload_images/14623749-7bec8f2e8ed07af8.png?imageMogr2/auto-orient/strip%7CimageView2/2/w/1240)

当我们关闭弹窗的时候，需要我们验证密码。

![已经标记为“受所有用户信任”](https://upload-images.jianshu.io/upload_images/14623749-91aab2a8c9c1560d.png?imageMogr2/auto-orient/strip%7CimageView2/2/w/1240)

当我们再次查看证书时，我们可以看到此时已经是**受所有用户信任**，此时电脑端的证书就已经安装完成了。

接下来，我们安装手机端的证书：

### 手机端

> 我这里以 iOS 设备为例，Android 手机可能会有所差异，但是我想应该安装步骤都是大差不差的，各位做一个参考也行。

在手机上安装证书，依然需要先打开 `Charles` 软件，然后在菜单栏中依次选择：`Help` -> `SSL Proxying` -> `Install Charles Root Certificate on a Mobile Device or Remote Browser` 然后直接点击。

![手机上安装证书](https://upload-images.jianshu.io/upload_images/14623749-0fdd6d440d4b36f0.png?imageMogr2/auto-orient/strip%7CimageView2/2/w/1240)

当我们点击之后会有一个弹框提示，大致内容如下：

![手机端安装时，会有弹窗提示](https://upload-images.jianshu.io/upload_images/14623749-eb69b8372e7e6d68.png?imageMogr2/auto-orient/strip%7CimageView2/2/w/1240)

大致的意思是说需要我们在手机上设置 HTTP 代理，代理地址为 `192.168.0.102:8888` 设置好后再通过浏览器访问 `chls.pro/ssl` 地址并下载安装证书。

并且还需要注意：**如果是 iOS 10 及其以后的版本时，还需要进入「设置」-> 「通用」->「关于本机」-> 滑到最底部「证书信任设置」，并启用 Charles 证书为可信任证书。**

并且还需要注意的是，你**一定需要将手机和电脑连接在同一个局域网内**，如果你的电脑是笔记本的话，就是说你的电脑和手机连接的是同一个 Wi-Fi。另外这里我的电脑端局域网 IP 地址为 `192.168.0.102` 你的可能和我的不一样，这个也是正常现象，以弹窗中的 IP 地址为准。

接下来，我们需要在手机上操作：

1. 打开「设置」找到「无线局域网」这里一定需要注意，此时你手机上连接的 Wi-Fi 一定和电脑所在的网络在同一个局域网内，不然就白搭了。
2. 点击你所连接的 Wi-Fi，滑到最底部会有一个「配置代理」，点击打开，选择「手动」。
3. 然后在“服务器”位置输入**192.168.0.102**，“端口”位置输入**8888**，然后点击右上角的「存储」。

![手机上配置代理](https://upload-images.jianshu.io/upload_images/14623749-248604b868222b90.png?imageMogr2/auto-orient/strip%7CimageView2/2/w/1240)

然后去打开 **Safari** 浏览器（其他浏览器**不一定**可以唤起安装证书的弹窗），输入地址 `chls.pro/ssl` 打开页面，此时会自动唤起安装描述文件的弹窗。

如果此时你发现并没有唤起安装描述文件的弹窗时，你需要回到你的电脑上，可能会有一个类似以下的弹窗，你需要点击一下**Allow**，这是你的手机连接到了 Charles 。当你点击了 **Allow** 之后再去手机 Safari 浏览器上刷新一下应该就会有安装描述文件的弹窗了。

![需要在电脑上允许一下](https://upload-images.jianshu.io/upload_images/14623749-90de545c3e0bb400.png?imageMogr2/auto-orient/strip%7CimageView2/2/w/1240)

当出现安装描述文件的弹窗后，会告知你现在正在下载一个配置描述文件，你直接点击**允许**就好了。

![手机上下载配置描述文件](https://upload-images.jianshu.io/upload_images/14623749-c86485bab631efaa.png?imageMogr2/auto-orient/strip%7CimageView2/2/w/1240)

当下载好证书之后，我们需要前往：「设置」->「通用」-> 「VPN与设备管理」然后就可以看到“已下载的描述文件” `Charles Proxy CA……` 进行点击进去，然后看到右上角会有一个「安装」按钮，点击安装即可。

![安装配置描述文件](https://upload-images.jianshu.io/upload_images/14623749-ab64c916867c3eda.png?imageMogr2/auto-orient/strip%7CimageView2/2/w/1240)

之前也说到了，如果是 iOS 10 及其以后的版本时，还需要在「证书信任设置」，并启用 Charles 证书为可信任证书，这点你按照你的系统版本按需来做就好了。

不过一般来说，现在很少有低于 iOS 10 的手机还在跑了吧？**其实这一点也是很多人虽然安装完了证书，但是发现依然抓取 HTTPS 请求时认为 unknown 的原因。** 所以，一定要记得去信任一下证书，这点非常重要！

![开启信任证书](https://upload-images.jianshu.io/upload_images/14623749-a650be3665c57d25.png?imageMogr2/auto-orient/strip%7CimageView2/2/w/1240)

现在我们证书已经安装好了，但是还需要简单配置一下 Charles。

## 配置 Charles

### 配置代理端口

直接点击菜单栏中的「Proxy」 ->「Proxy Settings」

![配置代理](https://upload-images.jianshu.io/upload_images/14623749-0a36814f9dfecb62.png?imageMogr2/auto-orient/strip%7CimageView2/2/w/1240)

还记得上面我们在手机上设置代理的时候那个 8888 端口吗？如果你想自定义代理端口，可以直接在这个位置进行更改。不改问题也不大，自己按照实际情况而定。

![更改代理端口](https://upload-images.jianshu.io/upload_images/14623749-58134297b5b4207b.png?imageMogr2/auto-orient/strip%7CimageView2/2/w/1240)

### 配置 SSL 代理设置

直接点击菜单栏中的「Proxy」->「SSL Proxying Settings」

![配置SSL代理设置](https://upload-images.jianshu.io/upload_images/14623749-9c01ebf614410cc0.png?imageMogr2/auto-orient/strip%7CimageView2/2/w/1240)

这里有一个非常重要的配置 **SSL Proxying Settings**，我们需要确保勾选上了「Enable SSL Proxying」并且还需要添加「Include」，**否则即使我们添加了证书，抓取 HTTPS 时还是会出现 unknown**。

![配置 SSL 代理设置](https://upload-images.jianshu.io/upload_images/14623749-457f3b5df5b88fff.png?imageMogr2/auto-orient/strip%7CimageView2/2/w/1240)

Include 和 Exclude 就是字面上的含义，分别代表包含和不包含，这里我在 Include 处设置两个星号 `*` 的含义是：包含所有的域名和端口。你也可以按照你自己的实际情况来设置。

好了，现在该配置的都已经配置完了，剩下的就可以愉快的玩耍啦。

> 另外，听说 Android 7.0 之后默认不信任用户添加到系统的 CA 证书，也就是说对基于 SDK24 及以上的 APP 来说，即使你在手机上安装了抓包工具的证书也无法抓取 HTTPS 请求。那么如何解决这个问题呢？当前我手上没有 Android 手机，也无法去测试，这个问题就留给有 Android 手机的用户来解决吧。不过，我想这个问题应该也已经有了解决方案，如果你知道解决方案，也希望一起分享分享。
