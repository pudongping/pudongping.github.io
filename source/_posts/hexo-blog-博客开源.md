---
title: hexo-blog 博客开源
author: Alex
top: true
hide: false
cover: true
toc: true
mathjax: false
img: https://pudongping.com/medias/featureimages/26.jpg
coverImg: https://pudongping.com/medias/banner/0.jpg
summary: >-
  折腾了个把星期左右，终于把博客改的有点儿样子了，秉承着开源精神，现在把博客源码开源出来。如果你也想拥有和我一摸一样的博客，那么赶紧来试试吧！如果你喜欢，请记得给个
  Star！
categories: 开源
tags:
  - GitHub
  - 博客
  - Hexo
abbrlink: e3e08109
date: 2021-06-06 19:23:44
password:
---

# 博客开源

## 前言

折腾了一个星期左右，总算是把我的个人博客给完善了，为了回馈开源，我会把我自己修改后的且完善的 `blog` 网站源码开源。
代码都是源码，您可以直接下载我的博客源码，然后将根目录下的 `/_config.yml` 和 `/themes/hexo-theme-matery/_config.yml`
这两个配置文件中的信息修改成您的信息就可以部署成和我一摸一样的博客了。是不是很方便？这对于一些想写博客，但是又不想太折腾，
并且对美观有一定要求的朋友来说，简直是爽的不要不要的。😜 [点我立即下载我的博客源码](https://github.com/pudongping/pudongping.github.io)

## 关于我的博客项目

这个博客，我是基于 **[Hexo](https://hexo.io/zh-cn/docs/)** 框架搭建，并且用到 **[hexo-theme-matery](https://github.com/blinkfox/hexo-theme-matery)** 主题，
在此基础上做了不少的修改，增加了一些新的特性和功能。我的博客访问地址为：[http://www.pudongping.com](http://www.pudongping.com) 或者访问 [https://pudongping.github.io](https://pudongping.github.io)

## 创建项目时，各个软件版本介绍
- 修改项目时间为 2021 年 6 月 5 日
- 使用的 node.js 版本为：v16.2.0
- 使用的 npm 版本为：7.13.0
- 使用的 hexo 版本为：5.4.0
- 使用的 hexo-theme-matery 皮肤版本为：v2.0.0
- 其他的依赖包见 package.json 文件，具体版本号见 package-lock.json 文件。

## 如何搭建和我一摸一样的博客？

> 前提你需要先安装 `git` 和 `node.js` （安装好 node.js 之后，会自动安装 npm，强烈建议使用 [nvm](https://github.com/nvm-sh/nvm) 安装 node.js 以方便管理多个版本的 node.js）

- 安装 `hexo-cli`

```shell
# 全局安装 hexo
npm install -g hexo-cli

# 局部安装 hexo
npm install hexo

```

- 直接使用 `git` 拉取项目并进入相应的文件目录

```shell
# 使用 GitHub 地址
git clone https://github.com/pudongping/pudongping.github.io.git blog

# 使用 gitee 地址 （要是代码拉取不下来的话，建议使用 gitee 地址）
git clone https://gitee.com/pudongping/pudongping.git blog
```

- 切换分支

建议切换到 `hexo` 分支，方便你自定义你的博客。

```shell
# 如果你想实现和我一摸一样的博客的话，你可以直接切换到 `main` 分支，这是我自己完整的博客源码分支，其中包含了我所有的博客文章
git checkout main

# 如果你只想获取我所有的修改，并基于此基础之上改成自己的博客的话，那么可以切换到 `hexo` 分支，这个分支上是我博客修改的基础分支，
# 在此分支基础之上你可以添加你自己的任何修改。（不包含我所有的博客文章和我自定义的其他修改）
git checkout hexo
```

- 安装依赖插件

```shell
npm install
```

- 修改根目录下的 `/_config.yml` 和 `/themes/hexo-theme-matery/_config.yml` 的这两个配置文件，改成自己的信息

- 编译源代码

```shell
# 清除生成的静态文件
hexo cl

# 生成静态页面至 public 页面 （如果只是本地预览，可不进行此步骤）
hexo g

# 生成本地预览
hexo s

# 如果需要生成 github-pages 的话，则执行以下命令 （需要安装 npm install hexo-deployer-git --save 插件包）
hexo d
```

- 本地访问项目

直接在浏览器中访问 [http://localhost:4000/](http://localhost:4000/]) 即可。部署完毕！

## 特性

### hexo-theme-matery 主题特性

- 简单漂亮，文章内容美观易读
- [Material Design](https://material.io/) 设计
- 响应式设计，博客在桌面端、平板、手机等设备上均能很好的展现
- 首页轮播文章及每天动态切换 `Banner` 图片
- 瀑布流式的博客文章列表（文章无特色图片时会有 `24` 张漂亮的图片代替）
- 时间轴式的归档页
- **词云**的标签页和**雷达图**的分类页
- 丰富的关于我页面（包括关于我、文章统计图、我的项目、我的技能、相册等）
- 可自定义的数据的友情链接页面
- 支持文章置顶和文章打赏
- 支持 `MathJax`
- `TOC` 目录
- 可设置复制文章内容时追加版权信息
- 可设置阅读文章时做密码验证
- [Gitalk](https://gitalk.github.io/)、[Gitment](https://imsun.github.io/gitment/)、[Valine](https://valine.js.org/) 和 [Disqus](https://disqus.com/) 评论模块（推荐使用 `Gitalk`）
- 集成了[不蒜子统计](http://busuanzi.ibruce.info/)、谷歌分析（`Google Analytics`）和文章字数统计等功能
- 支持在首页的音乐播放和视频播放功能
- 支持`emoji`表情，用`markdown emoji`语法书写直接生成对应的能**跳跃**的表情。
- 支持 [DaoVoice](http://www.daovoice.io/)、[Tidio](https://www.tidio.com/) 在线聊天功能。

### 修改或增加的特性功能
- 添加 emoji 表情支持
- 添加 RSS 订阅支持
- 添加站点地图支持
- css、js资源增加 cdn 加速
- 底部增加备案号信息
- 修改了主题颜色
- 处理了图片因防盗链无法显示问题
- 添加了动态改变页面标签的功能
- 添加了鼠标点击烟花爆炸特效
- 添加了鼠标点击出现自定义文字浮层特效
- 添加了页面樱花飘落动态特效
- 添加了页面雪花飘落动态特效
- 优化文章的 url
- 图片懒加载
- 添加了在**关于我**页面写简历功能
- 修改了 banner 图随机轮播

## 博客截图

![首页](https://pudongping.github.io/medias/sample/page.png)
![文章列表](https://pudongping.github.io/medias/sample/articles-list.png)
![文章列表](https://pudongping.github.io/medias/sample/articles-list-1.png)
![文章详情](https://pudongping.github.io/medias/sample/article-detail.png)
![文章详情](https://pudongping.github.io/medias/sample/article-detail-1.png)
![标签页](https://pudongping.github.io/medias/sample/tags.png)
![分类页](https://pudongping.github.io/medias/sample/categories.png)
![归档页](https://pudongping.github.io/medias/sample/archives.png)
![关于我页](https://pudongping.github.io/medias/sample/about-me.png)
![留言板页](https://pudongping.github.io/medias/sample/contact.png)

## 感谢支持

如果你觉得对你有所帮助,请帮忙给个 `Star`。  
如果你想贡献一份力量,欢迎提交 `Pull Request`。

## 赞赏捐助

有问题可以在文章最后评论区进行 **留言和评论** ，如果你喜欢我的博客，请博主喝一杯冰阔乐呀^_^ 😘

<table>
  <tr>
    <td>
        <img width="100" src="https://pudongping.github.io/medias/reward/alipay.png" alt="alipay" >
    </td>
    <td>
        <img width="100" src="https://pudongping.github.io/medias/reward/wechat.png" alt="wechat" >
    </td>
  </tr>
</table>


> 😘 若有共鸣，留言足矣，若有赞赏，何以复加？ 🤞

## LICENSE

MIT
