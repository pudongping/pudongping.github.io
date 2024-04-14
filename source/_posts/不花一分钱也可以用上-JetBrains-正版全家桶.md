---
title: 不花一分钱也可以用上 JetBrains 正版全家桶
author: Alex
top: true
hide: false
cover: false
toc: true
mathjax: false
categories: 科普
tags:
  - JetBrains
abbrlink: b3af447c
date: 2023-11-24 16:18:45
img:
coverImg:
password:
summary:
---

昨天，发现正在使用的 idea 要到期了，于是自己得马上去续约，免得影响自己的工作。

其实，我体验过不少的编辑器，Notepad++、Sublime Text、Apache NetBeans、Vim、Neovim、再就是大名鼎鼎的 Visual Studio Code 和巨强大的 JetBrains IDE。最终体验下来，在工作中的主要编辑器还是使用了 IDE。除了确实比较占用内存以外，没有其他可挑剔的。

![价格不菲](https://upload-images.jianshu.io/upload_images/14623749-dcae6b5c4b3eafc8.png?imageMogr2/auto-orient/strip%7CimageView2/2/w/1240)

![](https://img-blog.csdnimg.cn/img_convert/db879ed091daa570863771c473f61c70.png)

但是，要是直接去购买正版授权的话，价格确实不便宜，光一个 IDE 每月也要 24.9 美元，折合人民币小两百来块了，更别说全家桶了，要 77.9 美元。对于个人用户而言，估计一般都会出门右转去某宝上花个 9.9 买账号了。用确实可以用，但是就是不太稳定。

那么，有没有办法既想使用正版授权，又不想掏钱呢？还真有，**开源开发许可证——JetBrains 通过为核心项目贡献者免费提供一套一流的开发者工具来支持非商业开源项目。**

别看他写的条件那么多，其实总结下来也就几点：

1.  你得有一个开源项目，且近期三个月有一定的活跃度，也就是在近三个月有定期提交代码就行。
2. 在你的项目根目录下得有一个开源许可证，  项目为公开代码。
3. 并且此仓库还不能是博客和一些示例代码。（这一点好像是今年加上的吧，去年我续约的时候，貌似都还没有这一点）

具体满足条件可通过访问 https://www.jetbrains.com/zh-cn/community/opensource/#support 地址进行查看。

![满足条件](https://upload-images.jianshu.io/upload_images/14623749-eff4930bcc518842.png?imageMogr2/auto-orient/strip%7CimageView2/2/w/1240)

接下来讲如何申请这个免费的许可证。

## 注册账号

先访问 https://account.jetbrains.com/licenses 地址，在 JetBrains 官网上注册一个账号。

## 申请

直接访问官方的申请地址 https://www.jetbrains.com/shop/eform/opensource

这个页面有三个模块，我们大致讲解一下。

1.  Do we know you?

**No**：之前从来没有申请过的，就选这个。**Yes**：之前弄过开源许可的，就直接选这个，然后填上自己的 License ID 即可。

2. Tell us about your project

这里就根据你项目的实际情况填写就行了，没有太多的滑头。

3. Tell us about yourself

- Email address: 比较重要的是这一项，这个邮箱要和你 GitHub 主页上的邮箱地址一致。
- A link to your profile on GitHub,etc: 这个填你 GitHub 的主页地址即可。

![11.png](https://upload-images.jianshu.io/upload_images/14623749-ebea11272aecf962.png?imageMogr2/auto-orient/strip%7CimageView2/2/w/1240)

![22.png](https://upload-images.jianshu.io/upload_images/14623749-fb0937c35ce4e7aa.png?imageMogr2/auto-orient/strip%7CimageView2/2/w/1240)

![33.png](https://upload-images.jianshu.io/upload_images/14623749-7330fa398a713aa3.png?imageMogr2/auto-orient/strip%7CimageView2/2/w/1240)

![44.png](https://upload-images.jianshu.io/upload_images/14623749-4bf839e7eda63932.png?imageMogr2/auto-orient/strip%7CimageView2/2/w/1240)

好了，填写完毕之后，直接提交就可以等着了。一般需要等两周左右，其实很快，基本上在一周内就会有邮件通知你具体的申请结果。

![55.png](https://upload-images.jianshu.io/upload_images/14623749-7b6bb5cd1b97bf7c.png?imageMogr2/auto-orient/strip%7CimageView2/2/w/1240)

## 通过申请后

申请之后，你要留心你的邮箱里面标题为 **License Certificate for JetBrains** 开头的邮件。然后你需要点击 **Take me to my license(s)** 这个许可证可以直接和你的 JetBrains 账号直接绑定，使用 IDE 的时候直接通过你的账号登录即可。你也可以直接下载离线激活码，通过激活码的形式去激活 IDE。这个就看自己的喜好了。

## 注意事项

**在你的 GitHub 首页，需要把你的邮箱显示出来，** 这里主要是为了方便审查人员判断这是否为你自己的账号。人家都免费给你使用了，肯定也是希望你不要太鬼呀。

另外这个许可证是可以反复申请的，每一次的有效期是一年。也就是说，第一次申请成功之后，如果第二年你的仓库还满足申请条件，那么你还可以继续申请。

感叹开源的伟大！感叹 JetBrains 的伟大！
