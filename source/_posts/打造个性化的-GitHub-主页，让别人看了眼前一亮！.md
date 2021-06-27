---
title: 打造个性化的 GitHub 主页，让别人看了眼前一亮！
author: Alex
top: true
hide: false
cover: false
toc: true
mathjax: false
categories: 开源
tags:
  - GitHub
abbrlink: bd18a856
date: 2021-06-27 20:36:35
img:
coverImg:
password:
summary:
---

# 打造个性化的 GitHub 主页，让别人看了眼前一亮！

> 首先可以看一下我的 [GitHub 首页](https://github.com/pudongping)

![首页1](https://upload-images.jianshu.io/upload_images/14623749-e33225263f52d150.png?imageMogr2/auto-orient/strip%7CimageView2/2/w/1240)

![首页2](https://upload-images.jianshu.io/upload_images/14623749-e938fade8ae7a1c5.png?imageMogr2/auto-orient/strip%7CimageView2/2/w/1240)

![首页3](https://upload-images.jianshu.io/upload_images/14623749-3570ccfce03fb0cb.png?imageMogr2/auto-orient/strip%7CimageView2/2/w/1240)

## 如何构建属于自己个性化的 GitHub 首页？

- 先创建一个和自己 GitHub 同用户名的仓库，比如我的 GitHub 账号为 `pudongping` ，我的 GitHub 地址为：`https://github/pudongping` 那么我就需要创建一个名为 `pudongping` 的仓库，如果你填写的是正确的，会出现绿色的提示框，提示你正在创建一个 GitHub profile，如下图所示。

![当仓库名称和自己的 GitHub 同名时，会提示这是一个特别的仓库](https://upload-images.jianshu.io/upload_images/14623749-9b4928a3dc53bbd0.png?imageMogr2/auto-orient/strip%7CimageView2/2/w/1240)

> 需要注意的是：仓库的访问权限一定要设置为 `public` ，创建仓库的时候，你可以勾选一下 `Add a README file` 也可以不勾选，然后自己重新创建一个 `README.md` 文件，如果是勾选了的话，则默认的 `README.md` 文件中会有如下代码：

```html
### Hi there 👋 

<!--
**pudongping/pudongping** is a ✨ _special_ ✨ repository because its `README.md` (this file) appears on your GitHub profile.

Here are some ideas to get you started:

- 🔭 I’m currently working on ...
- 🌱 I’m currently learning ...
- 👯 I’m looking to collaborate on ...
- 🤔 I’m looking for help with ...
- 💬 Ask me about ...
- 📫 How to reach me: ...
- 😄 Pronouns: ...
- ⚡ Fun fact: ...
-->

```

![仓库名称需要设置成和自己 GitHub 账号同名](https://upload-images.jianshu.io/upload_images/14623749-91e84d1a05fa4228.png?imageMogr2/auto-orient/strip%7CimageView2/2/w/1240)

![访问和自己 GitHub 账号同名的仓库时内容](https://upload-images.jianshu.io/upload_images/14623749-eb2d92510ffa6cb3.png?imageMogr2/auto-orient/strip%7CimageView2/2/w/1240)

![默认的 GitHub profile README.md 文件](https://upload-images.jianshu.io/upload_images/14623749-e93baf911b63ba8a.png?imageMogr2/auto-orient/strip%7CimageView2/2/w/1240)

这时候你可以访问一下你自己的 GitHub 首页，就可以看到自己的个性化 `GitHub profile` 了，就是这么简单 ✌️

![查看自己的 GitHub profile](https://upload-images.jianshu.io/upload_images/14623749-8da90386857abfe2.png?imageMogr2/auto-orient/strip%7CimageView2/2/w/1240)

## 接下来说一说个性化定制，打造更加 `炫酷` 的 GitHub profile

### 每周代码统计

如果你需要统计你每周的代码的话，你可以先去 [WakaTime](https://wakatime.com) 网站，使用 GitHub 授权登录一下，获取 `wakatime` 的密钥，然后访问 [waka-readme](https://github.com/athul/waka-readme)  插件页面，根据说明文档设置一下就可以了。

> 代码统计需要运行 `GitHub Action` ，如果你不知道如何使用 `GitHub Action` ，那么你可以参考一下阮一峰的 [GitHub Actions 入门教程](http://www.ruanyifeng.com/blog/2019/09/getting-started-with-github-actions.html) ，这里就需要用到上面获取的 `wakatime` 的密钥。

展示效果如下：

```text
Week: 19 June, 2021 - 25 June, 2021

Markdown   1 hr 38 mins    ███████████████████▓░░░░░   79.03 % 
CSS        17 mins         ███▓░░░░░░░░░░░░░░░░░░░░░   14.17 % 
YAML       7 mins          █▒░░░░░░░░░░░░░░░░░░░░░░░   05.74 % 
```

以下以截图的形式展示如何使用：[WakaTime](https://wakatime.com)

先去 wakatime 中使用 github 注册一个账号

![注册账号](https://upload-images.jianshu.io/upload_images/14623749-992eeba288611344.png?imageMogr2/auto-orient/strip%7CimageView2/2/w/1240)

直接使用 GitHub 授权登录

![直接使用 GitHub 授权登录](https://upload-images.jianshu.io/upload_images/14623749-edc612cf27e4b612.png?imageMogr2/auto-orient/strip%7CimageView2/2/w/1240)

访问 https://wakatime.com/settings/account 获取 wakatime 的密钥

![获取 wakatime 的密钥](https://upload-images.jianshu.io/upload_images/14623749-9b34e7147d31b4ed.png?imageMogr2/auto-orient/strip%7CimageView2/2/w/1240)

在 GitHub 中设置刚刚获取到的 wakatime secret，注意密钥的名称设置为 WAKATIME_API_KEY

![设置密钥](https://upload-images.jianshu.io/upload_images/14623749-1995498683c19ab5.png?imageMogr2/auto-orient/strip%7CimageView2/2/w/1240)

![设置成功如下所示](https://upload-images.jianshu.io/upload_images/14623749-0659cec75b88f8c2.png?imageMogr2/auto-orient/strip%7CimageView2/2/w/1240)

然后在自己的仓库下面创建 `.github/workflows/waka-time.yml` 文件，填入对应的配置信息，这些配置可以在 [waka-time](https://github.com/athul/waka-readme) 仓库中找到示例。

![填入配置信息](https://upload-images.jianshu.io/upload_images/14623749-e1a749790225578a.png?imageMogr2/auto-orient/strip%7CimageView2/2/w/1240)

此时，修改好配置信息后，我们将代码推送到 GitHub，然后点击 `Action` 我们可以看到有工作流，我们可以找到 `Waka Readme ` 这一个工作流程，我们可以点一下 `Run workflow` 按钮手动执行一下

![查看一下 GitHub Action](https://upload-images.jianshu.io/upload_images/14623749-84d0a1130196df10.png?imageMogr2/auto-orient/strip%7CimageView2/2/w/1240)

![选中指定的工作流程](https://upload-images.jianshu.io/upload_images/14623749-7b11660f35badf18.png?imageMogr2/auto-orient/strip%7CimageView2/2/w/1240)

执行完毕之后，我们可以看到执行日志记录

![可以查看到执行日志记录](https://upload-images.jianshu.io/upload_images/14623749-762a711dcb3abda1.png?imageMogr2/auto-orient/strip%7CimageView2/2/w/1240)

此时我们再回到自己的 GitHub 首页，查看代码统计结果，我们可以看到红色框中的内容（此时还没有统计结果，我们可以为我们的代码编辑器下载对应的 `wakatime` 插件，不知道如何下载插件的话，可以[点击这里](https://wakatime.com/plugins)）

![查看结果](https://upload-images.jianshu.io/upload_images/14623749-6a20665d519221fe.png?imageMogr2/auto-orient/strip%7CimageView2/2/w/1240)

[WakaTime](https://wakatime.com/) 是一款很有用的代码统计工具，它不仅仅可以统计代码编辑器上面的活动状态，还可以统计其他的软件上面的活动，比如 `WPS Office` 、`Chrome` 等等，我这里就统计了 `WebStorm` 和 `PhpStorm` 等几个编辑器

![活动统计](https://upload-images.jianshu.io/upload_images/14623749-bd50fe828ce3c2bc.png?imageMogr2/auto-orient/strip%7CimageView2/2/w/1240)

### 其他好玩且有用的个性化定制
- [展示 GitHub Stats](https://github.com/anuraghazra/github-readme-stats)
- [展示最近一个月的贡献统计数据](https://github.com/ashutosh00710/github-readme-activity-graph)
- [徽章制作工具](https://shields.io/)
- [怪异英文生成器](https://www.dute.org/weird-fonts) - 生成比较好看的英文字体，支持复制粘贴

### 其他
- 如果你喜欢我的 [GitHub profile](https://github.com/pudongping/pudongping) 你可以克隆代码之后，自己自定义修改。
- 制作不易，很多内容都是我一步一步通过搜索引擎获取到的，比如说徽章的背景色，我都是打开各大官网吸取的背景色。如果你喜欢的话，麻烦给个 Star 哇！
