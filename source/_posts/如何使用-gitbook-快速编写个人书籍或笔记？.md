---
title: 如何使用 gitbook 快速编写个人书籍或笔记？
author: Alex
top: false
hide: false
cover: false
toc: true
mathjax: false
summary: >-
  使用 `gitbook` 写书籍很 nice，但是如果再搭配合适的插件，那简直就是如虎添翼之举，当然你也可以直接参照我的 `gitbook` 配置文件
  `book.json`，喜欢就拿去。
categories: 开源
tags:
  - GitBook
  - 博客
  - GitHub
abbrlink: b00f410f
date: 2021-06-07 07:53:30
img:
coverImg:
password:
---

# [GitBook](https://github.com/GitbookIO) 使用

## 参考文献

> [GitBook 官网](https://www.gitbook.com/)   
[GitBook 文档](https://github.com/GitbookIO/gitbook)  
[GitBook 使用教程](https://blankj.com/gitbook/gitbook/)  
[推荐12个实用的gitbook插件](https://juejin.im/post/6844903865146441741)  
[GitBook 和它有趣的插件](http://jartto.wang/2020/02/02/about-gitbook/)

## 示例图

![页面截图](https://pudongping.github.io/notes/resources/images/sample/page.png)

## 在线访问

[蒲东平的编程笔记](https://notes.pudongping.com) 或者 [蒲东平的编程笔记-GitHub Pages](https://pudongping.github.io/notes)

## 使用

- 安装

> 安装 GitBook 之前需要安装 `Node.js` ， `GitBook` 是一个基于 `Node.js` 的命令行工具，因此需要先下载安装 [Node.js](https://nodejs.org/en)

```

# 查看 node.js 是否安装成功
node -v

# 安装 GitBook 
npm install gitbook-cli -g

# 查看 GitBook 是否安装成功
gitbook -V （大写的 V ）

```

更多详情可以参考 [GitBook 官方安装文档](https://github.com/GitbookIO/gitbook/blob/master/docs/setup.md) 来安装 GitBook


- 创建项目

```

# 切换到项目文件夹并创建项目
mkdir project-directory && cd project-directory && gitbook init

# 或者直接使用以下命令
gitbook init ./project-directory

```

- 启动项目

```

cd project-directory && gitbook serve

// 如果不想使用 4000 端口，想要使用 9520 端口时
gitbook serve -p 9520

```

然后在浏览器地址栏中输入 `http://localhost:4000` 便可预览书籍，至此，gitbook 安装完毕。

- 编译项目 （生成网页而不开启服务器）

```

cd project-directory && gitbook build

```

- 查看所有可用的 gitbook 版本

```

gitbook ls-remote

```

- 安装指定的 gitbook 版本

```

gitbook fetch beta（版本号）

```

- 编译时，输出目录详细的记录包括 debug

```

gitbook build ./ --log=debug --debug

```

## 关于配置

### 配置文件

需要在项目根目录下手动创建 `book.json` 或者 `book.js` 文件

### 配置文件变量
可以参考这篇文章中的介绍：[GitBook文档（中文版）](https://chrisniael.gitbooks.io/gitbook-documentation/content/format/templating.html)

### 如果需要导出 pdf 等文件
可以参考这篇文章中的介绍： [导出电子书](https://snowdreams1006.github.io/myGitbook/advance/export.html)


## 安装插件

安装插件有两种方式： [点我搜索更多 gitbook 插件](https://www.npmjs.com/search?q=gitbook-plugin)

1. 在 `/book.json` 配置文件中写入相应的插件和配置后，使用 `gitbook install` 命令安装插件
2. 直接使用 `npm install gitbook-plugin-pluginname` 命令安装指定的插件，然后在 `/book.json` 配置文件中写入配置，比如安装 `highlight` 插件时，需要执行 `npm install gitbook-plugin-highlight`

```npm

npm install gitbook-plugin-search-pro --registry=https://registry.npm.taobao.org/

```

### GitBook 初始化时，会默认自动下载 7 个插件
1. [highlight](https://github.com/GitbookIO/plugin-highlight)  ： 代码高亮
2. [search](https://github.com/GitbookIO/plugin-search)  ： 导航栏查询功能 （不支持中文）
3. [lunr](https://www.npmjs.com/package/gitbook-plugin-lunr)  ：  为 `search` 插件提供后端支持
4. [sharing](https://github.com/GitbookIO/plugin-sharing)  ： 右上角分享功能
5. [fontsettings](https://github.com/GitbookIO/plugin-fontsettings)  ： 字体设置（最上方的 "A" 符号）
6. [livereload](https://github.com/GitbookIO/plugin-livereload)  ： 为 GitBook 实时重新加载
7. [theme-default](https://www.npmjs.com/package/gitbook-plugin-theme-default)  ：  默认的 3.0.0 版本之后的主题插件

### 推荐插件

- [advanced-emoji](https://github.com/codeclou/gitbook-plugin-advanced-emoji)  ： 支持 emoji 表情包  [emoji 表情包地址](https://www.webfx.com/tools/emoji-cheat-sheet/)
- [anchor-navigation-ex](https://github.com/zq99299/gitbook-plugin-anchor-navigation-ex/blob/master/doc/config.md)  ： 悬浮按钮目录  （在页面中增加 <extoc></extoc> 标签，会在此处生成 TOC 目录）、还会增加标题锚
- [auto-scroll-table](https://www.npmjs.com/package/gitbook-plugin-auto-scroll-table)  ：  表格滚动条
- [audio_image](https://www.npmjs.com/package/gitbook-plugin-audio_image)  ： 播放音频
- [back-to-top-button](https://github.com/stuebersystems/gitbook-plugin-back-to-top-button) ： 返回顶部
- [baidu-tongji-with-multiple-channel](https://snowdreams1006.github.io/gitbook-plugin-baidu-tongji-with-multiple-channel/)  ：  百度统计插件,支持多渠道独立统计,一份源码多处部署独立统计
- [chapter-fold](https://www.npmjs.com/package/gitbook-plugin-chapter-fold)  ： 左侧目录可折叠，建议和 `expandable-chapters` 插件一起使用可以互补相互的 bug
- [code](https://github.com/TGhoul/gitbook-plugin-code) ： 代码添加行号和复制按钮
- [change_girls](https://github.com/zhenchao125/gitbook-plugin-change_girls)  ： 可自动切换背景
- [click-reveal](https://github.com/c4software/gitbook-plugin-click-reveal)  ： 点击显示内容，默认把内容已经隐藏
- [custom-favicon](https://github.com/Bandwidth/gitbook-plugin-custom-favicon)  ：  修改标题栏图标
- [copyright](https://www.npmjs.com/package/gitbook-plugin-copyright)  ：  用于复制内容时追加版权信息以及文章末尾添加版权小尾巴
- [donate](https://github.com/willin/gitbook-plugin-donate) ： 打赏功能
- [diff](https://github.com/snowdreams1006/gitbook-plugin-diff/blob/HEAD/README_zh.md)  ： 在 markdown 文档中显示代码之间的差异
- [expandable-chapters](https://github.com/chrisjake/gitbook-plugin-expandable-chapters-small)  ： 左侧目录可折叠，和 `expandable-chapters-small` 的区别是： `expandable-chapters-small` 的折叠图标要小一些 （建议和 `chapter-fold` 插件一起使用可以互补相互的 bug）
- [edit-link-plus](https://www.npmjs.com/package/gitbook-plugin-edit-link-plus)  ： 在线编辑文件
- [flexible-alerts](https://github.com/fzankl/gitbook-plugin-flexible-alerts) ：警告提示框
- [favicon](https://github.com/menduo/gitbook-plugin-favicon)  ： 修改网站的 favicon.ico  [点我在线制作ico图标，建议尺寸32*32](http://www.bitbug.net/)
- [ga](https://www.npmjs.com/package/gitbook-plugin-ga) ： Google 分析
- [github](https://github.com/GitbookIO/plugin-github) ： 右上角添加 Github 图标
- [github-buttons](https://github.com/azu/gitbook-plugin-github-buttons)  ： 右上角添加 GitHub 的按钮
- [gitbook-plugin-charts](https://www.npmjs.com/package/gitbook-plugin-charts)  ：  在 gitbook 中使用图表插件，目前支持 echarts
- [google-tongji-with-multiple-channel](https://github.com/snowdreams1006/gitbook-plugin-google-tongji-with-multiple-channel)  ：  Google 统计插件,支持多渠道独立统计,一份源码多处部署独立统计
- [hide-element](https://www.npmjs.com/package/gitbook-plugin-hide-element)  ：  隐藏元素
- [image-captions](https://github.com/todvora/gitbook-plugin-image-captions) ： 将图片的 alt 或 title 属性转换为标题
- [insert-logo](https://github.com/matusnovak/gitbook-plugin-insert-logo) ： 插入 Logo
- [icp](https://github.com/snowdreams1006/gitbook-plugin-icp/blob/HEAD/README_zh.md)  ：  在首页页脚区域添加 icp 网站备案信息
- [klipse](https://github.com/brian-dawn/gitbook-plugin-klipse)  ： 嵌入类似 IDE 的功能
- [lightbox](https://www.npmjs.com/package/gitbook-plugin-lightbox)  ：  页面弹窗查看图片 （支持在弹层切换上下图）
- [local-video](https://github.com/PacktPublishing/gitbook-local-video)  ：  播放本地视频，[点我查看使用教程](http://gitbook.zhangjikai.com/plugins.html#local-video)
- [multipart](https://www.npmjs.com/package/gitbook-plugin-multipart) ： 将左侧的目录分章节展示
- [mygitalk](https://www.npmjs.com/package/gitbook-plugin-mygitalk)  ：  通过 GitHub issues 添加评论框
- [mermaid-gb3](https://www.npmjs.com/package/gitbook-plugin-mermaid-gb3)  ： 支持 markdown 的流程图
- [meta](https://www.npmjs.com/package/gitbook-plugin-meta)  ： 添加 meta 头部信息
- [page-copyright](https://www.npmjs.com/package/gitbook-plugin-page-copyright) ： 页脚版权
- [pageview-count](https://github.com/tinys/gitbook-plugin-pageview-count#readme) ： 阅读统计
- [popup](https://github.com/somax/gitbook-plugin-popup#readme)  ：  打开新的页面查看图片
- [prism](https://github.com/gaearon/gitbook-plugin-prism)  ：  代码块颜色插件 （使用的时候需要禁用掉 gitbook 自带的 `highlight` 插件并且和 `code` 插件一起使用时，需要放到 `code` 插件后，否则样式会被覆盖掉）
- [rss](https://www.npmjs.com/package/gitbook-plugin-rss)  ： 生成 rss [如何使用 RSS ](https://www.ruanyifeng.com/blog/2006/01/rss.html)
- [readmore](https://www.npmjs.com/package/gitbook-plugin-readmore)  ： 实现博客的每一篇文章自动增加阅读更多效果,关注公众号后方可解锁全站文章
- [sharing-plus](https://github.com/GitbookIO/plugin-sharing) ： 分享当前页面
- [splitter](https://github.com/yoshidax/gitbook-plugin-splitter) ： 侧边栏宽度可调节
- [search-pro](https://github.com/gitbook-plugins/gitbook-plugin-search-pro) ： 高级搜索，支持中文 （使用的时候需要禁用掉 `lunr` 和 `search` 插件）
- [sectionx](https://github.com/manchiyiu/gitbook-plugin-sectionx)  ： 折叠模块(页面内容可折叠)
- [sitemap-general](https://github.com/CyberZHG/gitbook-plugin-sitemap-general)  ： 生成站点地图
- [simple-mind-map](https://snowdreams1006.github.io/gitbook-plugin-simple-mind-map/zh/)  ：  在 markdown 中生成并导出思维导图
- [todo](https://github.com/ly-tools/gitbook-plugin-todo) ： 待做项
- [theme-fexa](https://github.com/tonyyls/gitbook-plugin-theme-fexa)  :  网站主题，使用这个主题之后不能够在页面上进行上下文翻页
- [theme-comscore](https://www.npmjs.com/package/gitbook-plugin-theme-comscore)  ： 标题和正文颜色有所区分的主题，表格也有颜色
- [tbfed-pagefooter](https://github.com/zhj3618/gitbook-plugin-tbfed-pagefooter)  ： 添加页脚版权信息，这个感觉没有 `page-copyright` 好用


## 我自己的 book.json 配置信息

> 注意：需要删除掉所有的注释信息

```json

{
    "title": "Alex's Notes",  // 设置书本的标题
    "author": "Alex",  // 作者的相关信息
    "description": "live and learn",  // 本书的简单描述
    "language": "zh-hans",  // 可使用的语言：en, ar, bn, cs, de, en, es, fa, fi, fr, he, it, ja, ko, no, pl, pt, ro, ru, sv, uk, vi, zh-hans, zh-tw 这里我选择的是简体中文 zh-hans
    "gitbook": "3.2.3",  // 指定使用的 gitbook 版本
    "styles": {  // 自定义页面样式
        "website": "./resources/styles/website.css",  // 当此时的 gitbook 输出为站点模式时使用的 css 样式
        // "ebook": "styles/ebook.css",  // 当此时的 gitbook 输出为 ebook 时使用的 css 样式
        // "pdf": "styles/pdf.css",  // 当此时的 gitbook 输出为 pdf 时使用的 css 样式
        // "mobi": "styles/mobi.css",
        // "epub": "styles/epub.css"        
    },
    "structure": {  // 指定 Readme、Summary、Glossary 和 Languages 对应的文件名
        "readme": "README.md",  // 该书的介绍 （默认会创建）
        "summary": "SUMMARY.md",  // 该书的章节结构 （默认会创建）
        // "glossary": "GLOSSARY.md",  // 多语言书籍
        // "languages": "LANGS.md",  // 术语描述的清单
    },
    "links": {  // 在左侧导航栏添加链接信息
        "sidebar": {
            "我的博客": "https://drling.xin/",
            "GitHub": "https://github.com/pudongping"
        }
    },
    "plugins": [  // 需要使用的插件列表，注释插件的时候直接在插件名称前加 “横杠 -”，比如注释 “highlight” 插件为 “-highlight”
        "-highlight",
        "-lunr",
        "-search",
        "search-pro",
        "-sharing",
        "sharing-plus",
        "fontsettings",
        "livereload",
        "expandable-chapters-small",
        "chapter-fold",
        "splitter",
        "hide-element",
        "back-to-top-button",
        "favicon",
        "insert-logo",
        "pageview-count",
        "code",
        "prism",
        "lightbox",
        "github",
        "github-buttons",
        "donate",
        "anchor-navigation-ex",
        "meta",
        "mygitalk",
        "change_girls",
        "simple-mind-map",
        "image-captions",
        "todo",
        "edit-link-plus",
        "sitemap-general",
        "rss",
        "icp",
        "theme-comscore",
        "page-copyright"
    ],
    "pluginsConfig": {
        "sharing": {
            "douban": true,
            "facebook": false,
            "google": true,
            "hatenaBookmark": false,
            "instapaper": false,
            "line": false,
            "linkedin": false,
            "messenger": false,
            "pocket": false,
            "qq": true,
            "qzone": true,
            "stumbleupon": false,
            "twitter": false,
            "viber": false,
            "vk": false,
            "weibo": true,
            "whatsapp": false,
            "all": [
                "douban",
                "facebook",
                "google",
                "hatenaBookmark",
                "instapaper",
                "line",
                "linkedin",
                "messenger",
                "pocket",
                "qq",
                "qzone",
                "stumbleupon",
                "twitter",
                "viber",
                "vk",
                "weibo",
                "whatsapp"
            ]
        },
        "hide-element": {
            "elements": [
                "a.gitbook-link[href='https://www.gitbook.com']"
            ]
        },
        "favicon": {
            "shortcut": "./resources/images/favicon.ico",
            "bookmark": "./resources/images/favicon.ico",
            "appleTouch": "./resources/images/favicon.png",
            "appleTouchMore": {
                "120x120": "./resources/images/apple-touch-icon-120x120.png",
                "180x180": "./resources/images/apple-touch-icon-180x180.png"
            }
        },
        "insert-logo": {
            "url": "/resources/images/favicon.png",
            "style": "background: none; max-height: 30px; min-height: 30px"
        },
        "prism": {
            "css": [
                "prismjs/themes/prism-solarizedlight.css"
            ],
            "lang": {
                "shell": "bash"
            }
        },
        "lightbox": {
            "sameUuid": true  // 开启了这个属性之后支持在弹层，左右切换图片
        },
        "github": {
            "url": "https://github.com/pudongping"  // 在右上角会显示很小的 github 的官方图标
        },
        "github-buttons": {  // 在右上角会显示 github 图标的 button
            "buttons": [
                {
                    "user": "pudongping",
                    "repo": "glory",
                    "type": "star",
                    "count": true,
                    "size": "small"
                }
            ]
        },
        "donate": {
            "alipay": "/resources/images/donate.png",
            "title": "\"若有共鸣，留言足矣，若有赞赏，何以复加？\"",
            "button": "赞赏",
            "alipayText": "微信/支付宝/QQ"
        },
        "anchor-navigation-ex": {
            "showLevel": true,  // 右上角浮层目录显示序号
            "showGoTop": false  // 不显示回到顶部的图标，不建议开启这个属性，因为这个图标总是显示，不像 “back-to-top-button” 插件还可以自动显示和隐藏
        },
        "meta": {
            "data": [
                {
                    "name": "referrer",
                    "content": "never"
                }
            ]
        },
        "mygitalk": {
            "clientID": "",  // GitHub 开发者设置，客户端连接标识
            "clientSecret": "",  // GitHub 开发者设置，客户端秘钥
            "repo": "",  // GitHub 仓库名
            "owner": "",  // GitHub 仓库所有者
            "admin": [  // GitHub 仓库管理者，支持多个管理者
                "admin-1",
                "admin-2"
            ],
            "distractionFreeMode": false  // 类似 Facebook 评论框的全屏遮罩效果,默认值: false
        },
        "change_girls": {
            "time": 5,  // 每 5 秒切换一次背景
            "urls": [
                "https://timgsa.baidu.com/timg?image&quality=80&size=b9999_10000&sec=1605033246957&di=d6170f1a9f0466f270ad1baee847eab9&imgtype=0&src=http%3A%2F%2Fpic1.win4000.com%2Fwallpaper%2Fe%2F55f26f55e9138.jpg",
                "https://timgsa.baidu.com/timg?image&quality=80&size=b9999_10000&sec=1605033334002&di=4ef73db6c98fb737f5a3068670160056&imgtype=0&src=http%3A%2F%2Fww3.sinaimg.cn%2Flarge%2Fd2e27164gw1fbmwbgf0mij21hc0u0487.jpg"
            ]
        },
        "simple-mind-map": {
            "type": "markdown",
            "preset": "colorful",
            "linkShape": "diagonal",
            "autoFit": true
        },
        "image-captions": {
            "caption": "sugaryesp 的笔记 - _PAGE_LEVEL_._PAGE_IMAGE_NUMBER_ - _CAPTION_"
        },
        "edit-link-plus": {
            "base": {
                "edit-link-name-1": "edit-link-1",
                "edit-link-name-2": "edit-link-2"
            },
            "defaultBase": "",  // 这里填写链接地址
            "label": "编辑本页"
        },
        "sitemap-general": {
            "prefix": "http://notes.drling.xin/"
        },
        "rss": {
            "title": "sugaryesp 的笔记",
            "description": "削个椰子皮_给个梨的笔记",
            "author": "Alex",
            "site_url": "http://notes.drling.xin",
            "feed_url": "http://notes.drling.xin/rss",
            "managingEditor": "276558492@qq.com (Alex Pu)",
            "webMaster": "276558492@qq.com (Alex Pu)",
            "categories": [
                "markdown",
                "git",
                "gitee",
                "github",
                "php",
                "python",
                "vue.js"
            ]
        },
        "icp": {
            "number": "鄂ICP备18004705号",
            "link": "https://beian.miit.gov.cn/"
        },
        "page-copyright": {
            "description": "modified at:",
            "signature": "Alex",
            "wisdom": "Artisan, Backend Developer & overall web enthusiast",
            "format": "YYYY-MM-dd hh:mm:ss",
            "copyright": "Copyright &#169; Alex",
            "timeColor": "#666",
            "copyrightColor": "#666",
            "utcOffset": "8",
            "style": "normal",
            "noPowered": true,
            "baseUri": "http://notes.drling.xin/"
        }
    }
}

```


## 如果使用 gitbook install 安装插件太慢，可以使用 npm init 初始化项目，然后再使用 npm install 安装插件

```

cd project-directory && npm init

// 比如安装 livereload 插件

npm install gitbook-plugin-livereload

```

### 这是我的 package.json

```json

{
  "name": "glory",
  "version": "1.0.0",
  "description": "Alex's notes",
  "main": " ",
  "dependencies": {
    "gitbook-plugin-anchor-navigation-ex": "^1.0.14",
    "gitbook-plugin-back-to-top-button": "^0.1.4",
    "gitbook-plugin-change_girls": "^2.2.1",
    "gitbook-plugin-chapter-fold": "^0.0.4",
    "gitbook-plugin-code": "^0.1.0",
    "gitbook-plugin-donate": "^1.0.2",
    "gitbook-plugin-edit-link-plus": "^0.1.1",
    "gitbook-plugin-expandable-chapters-small": "^0.1.7",
    "gitbook-plugin-favicon": "^0.0.2",
    "gitbook-plugin-fontsettings": "^2.0.0",
    "gitbook-plugin-github": "^2.0.0",
    "gitbook-plugin-github-buttons": "^3.0.0",
    "gitbook-plugin-hide-element": "^0.0.4",
    "gitbook-plugin-icp": "^0.1.2",
    "gitbook-plugin-image-captions": "^3.1.0",
    "gitbook-plugin-insert-logo": "^0.1.5",
    "gitbook-plugin-lightbox": "^1.2.0",
    "gitbook-plugin-livereload": "0.0.1",
    "gitbook-plugin-meta": "^0.1.12",
    "gitbook-plugin-mygitalk": "^0.2.6",
    "gitbook-plugin-page-copyright": "^1.0.8",
    "gitbook-plugin-pageview-count": "^1.0.1",
    "gitbook-plugin-prism": "^2.4.0",
    "gitbook-plugin-rss": "^3.0.2",
    "gitbook-plugin-search-pro": "^2.0.2",
    "gitbook-plugin-sharing-plus": "^0.0.2",
    "gitbook-plugin-simple-mind-map": "^0.2.4",
    "gitbook-plugin-sitemap-general": "^0.1.1",
    "gitbook-plugin-splitter": "^0.0.8",
    "gitbook-plugin-theme-comscore": "0.0.3",
    "gitbook-plugin-todo": "^0.1.3"
  },
  "devDependencies": {},
  "scripts": {
    "test": " "
  },
  "repository": {
    "type": "git",
    "url": "git+https://github.com/pudongping/glory.git"
  },
  "keywords": [
    "Artisan"
  ],
  "author": "Alex",
  "license": "MIT",
  "bugs": {
    "url": "https://github.com/pudongping/glory/issues"
  },
  "homepage": "https://github.com/pudongping/glory#readme"
}


```
