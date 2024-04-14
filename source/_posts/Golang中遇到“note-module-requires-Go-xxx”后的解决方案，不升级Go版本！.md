---
title: Golang中遇到“note module requires Go xxx”后的解决方案，不升级Go版本！
author: Alex
top: false
hide: false
cover: false
toc: true
mathjax: false
categories: Go
tags:
  - Go
  - Golang
abbrlink: 1c839b10
date: 2024-04-14 19:27:26
img:
coverImg:
password:
summary:
---

前几天，需要对一个两年前写的项目添加点儿新功能，需要用到一个 Http 客户端包，于是就用了 `https://github.com/go-resty/resty` 这个插件包。

我先是直接在项目根目录下执行了以下包的安装命令：

```sh
go get -v github.com/go-resty/resty
```

然后，在业务代码中按照官方文档实例化了 `client := resty.New()` 对象，紧接着我想先启动一下项目，看这个包是否正常可用，结果，执行 `go run main.go` 命令时，就报以下错误了。

![报错了](https://upload-images.jianshu.io/upload_images/14623749-0760e7ef8be84cf3.png?imageMogr2/auto-orient/strip%7CimageView2/2/w/1240)

## 问题分析

我们可以看到报错提示是：`note: module requires Go 1.17` 因为，我现在维护的这个项目是两年前写的，那个时候 Go 比较稳定的版本还是 `go 1.16` 因此，这个项目也是基于 `go 1.16` 版本写的。可近期 Go 的版本迭代非常快，截止发稿，现在 Go 已经迭代到 `go 1.22.2` 版本了，并且每个版本之间差异也不小，因此跨版本之后就需要解决一些兼容性问题。

回到刚刚的报错提示：`module requires Go 1.17` 我们大致可见就是**有部分插件包依赖了 Go 1.17** 高版本，而我现在本地版本还是 Go 1.16，版本不兼容导致，那么，应该如何解决这个问题呢？

## 探讨解决方案

当然，我们可以直接升级一下 Go 的版本，这个问题应该就会迎刃而解。但是，这个项目已经在生产环境上跑了两年多了，现在贸然的去升级 Go 版本，感觉还是有些不妥，毕竟没有什么比项目稳定更加重要了。

不能去升级 Go 版本，那就只有一种解决方案了。

**找到有问题的插件包，然后对有问题的依赖包进行降级就好了**。但是，问题是：如何快速且准确的找到这个依赖包呢？

我们再回到刚刚的报错提示，可以仔细查看到，可能跟 `golang.org/x/sys@v0.13.0` 这个包有关系，毕竟错误信息中就含有它。

那么，按照这个思路，我们可以一步一步查到每个包的依赖，方便我们好定位问题。

一般情况下，此时，我们可能会想到直接使用 `go mod graph` 命令来查看项目现有的结构图，但是，如果这个项目依赖的包不算多，我们还可以勉勉强强捋得清楚相关的依赖，如果依赖包比较多了，估计也看麻了吧……

## 借助工具来查找依赖关系

你会发现根本无从看起，那么，我们是否可以借助某些工具来查看呢？其实，我也不确定有没有这样的工具，那还是老套路呗，直接去 Github 上搜一波试了看，结果，一搜还真有！

就是这个包：`https://github.com/PaulXu-cn/go-mod-graph-chart` 看了下包介绍：**“一个能将 `go mod graph` 输出内容可视化的无依赖小工具”** ，这不正是我需要的嘛！

果断用起来！

由于这个小工具是一个二进制文件，于是直接使用 `go install` 命令安装了下。

![使用工具](https://upload-images.jianshu.io/upload_images/14623749-50d8c5f97c29c6f9.png?imageMogr2/auto-orient/strip%7CimageView2/2/w/1240)

然后在项目根目录下执行 `go mod graph | gmchart -keep 1` （设置 `-keep 1` 是为了保证 HTTP 服务永不退出，当然也可以完全不设置，如果不设置的话 gmchart 启动的 HTTP 服务就只会启动一分钟。）

![启动工具](https://upload-images.jianshu.io/upload_images/14623749-2148d071f6be258b.png?imageMogr2/auto-orient/strip%7CimageView2/2/w/1240)

当你执行完以上命令之后，会在浏览器中自动打开一个可视化的展示图表。这个工具其实就是将 `go mod graph` 命令输出的内容以**树状**的形式渲染成 web 页面展示了出来，更加方便我们查看每个包的依赖关系。

![直接搜索依赖](https://upload-images.jianshu.io/upload_images/14623749-a327b7f04b13f9d1.png?imageMogr2/auto-orient/strip%7CimageView2/2/w/1240)

有了这个工具的协助，我们再回过头来去解决刚刚的报错问题。我们可以直接在浏览器中搜索 `sys@v0.13.0` 关键词，看这个包被哪些包依赖，然后依次检查各个依赖包所支持的版本，如果高于 `go 1.16` 那么则直接将这个包降级就好了。

搜到 `x/sys@v0.13.0` 之后，我们对着它点击一下，就可以看见如下，有三个包依赖了 `golang.org/x/sys@v0.13.0` 包。

![确定依赖关系](https://upload-images.jianshu.io/upload_images/14623749-1cc83cbb02dd48f8.png?imageMogr2/auto-orient/strip%7CimageView2/2/w/1240)

现在的思路就是一个一个包去找，把依赖高版本的包找出来。此时我们可以直接访问包的下载地址 `golang.org/x/net` 发现它会自动重定向到 `https://pkg.go.dev/golang.org/x/net` 地址上，这是由于官方包换了域名导致，不用太在意这点。

![找到具体包](https://upload-images.jianshu.io/upload_images/14623749-3cf5b39fd0d37fbf.png?imageMogr2/auto-orient/strip%7CimageView2/2/w/1240)

截至写这篇文章时，`golang.org/x/net` 的 Latest 版本是 v0.22.0，那么，我们如何知道此时的版本依赖于什么版本的 Go 呢？是的，我们可以通过直接去查看这个包的 `go.mod` 文件，就一目了然了。

> 在 Go Modules 模式下，项目根目录的 go.mod 文件中，都会记录当前项目依赖的 Go 最低版本。

![找到指定版本](https://upload-images.jianshu.io/upload_images/14623749-79fe07d35d94c8e8.png?imageMogr2/auto-orient/strip%7CimageView2/2/w/1240)

我们对着 Latest 左侧的版本点一下，以便我们可以查看到所有的版本。因为在我的项目中采用的版本是 `golang.org/x/net@v0.17.0` 因此，我们先优先查看这个版本的 `go.mod` 文件。

![查看go.mod](https://upload-images.jianshu.io/upload_images/14623749-19da41b4a3e7bdde.png?imageMogr2/auto-orient/strip%7CimageView2/2/w/1240)

进入到指定版本页面之后，再点击页面右上方侧的 `go.mod` 处。

![查看指定版本的go.mod内容](https://upload-images.jianshu.io/upload_images/14623749-bb730c130fb98188.png?imageMogr2/auto-orient/strip%7CimageView2/2/w/1240)

我们就可以看到这个版本的 go.mod 文件了，诶…… 是不是一下子就看到了 `go 1.17` 字眼了？有点儿小激动了，是不是？那么，我们只需要往前面的版本中去找，只要找到有一个最大的版本是依赖于 `go 1.16` 就行，然后，我就开始继续找呀找……

![查看各个版本](https://upload-images.jianshu.io/upload_images/14623749-67d9f7360bc3c0ef.png?imageMogr2/auto-orient/strip%7CimageView2/2/w/1240)

点击依赖包的版本位置处，可见所有的版本，然后我们一个版本一个版本的点击后看看。

![最小的版本也是依赖go 1.17](https://upload-images.jianshu.io/upload_images/14623749-bd14b74f47b45584.png?imageMogr2/auto-orient/strip%7CimageView2/2/w/1240)

发现，比较恶心的一幕出现了！**连最低的版本 v0.1.0 就是依赖于 go 1.17 ！**  真是 WTF ！

## 再次碰壁！遇到难题

那……这个问题就没有解了吗？其实并不是，我们忽略了一点，还有一种版本形式，就是类似于 `v0.0.0-20211029224645-99673261e6eb` 这种的版本，为什么会有这样的版本，感兴趣的童鞋可以自行去查一查，今天的重点不在这里。

虽然我们知道，这样的版本肯定是小于 v0.1.0 版本的，但是我们又该如何去找到这种版本号的插件包呢？问题一下子就棘手了起来……

此时，我们再把这个事情好好捋顺一下。我是因为要在项目中用到 `github.com/go-resty/resty` 包，所以我就使用 `go get -v github.com/go-resty/resty` 命令去下载了这个插件包，然后启动项目的时候就直接报错了。报错内容分析得出是某些包依赖了高版本的 go 1.17 我自己本地使用的是 go 1.16 ，好吧，那我先在项目中不使用 `github.com/go-resty/resty` 包，看项目是否能够跑得起来。当我在项目中没有去掉了 resty 包之后，项目跑起来了，没有发生错误，证明我之前的项目是正常的，就是因为下载了 resty 包之后导致无法启动的。可是，我现在就想用 resty 包，并且根据以上的排查，我们还发现了 `golang.org/x/net` 包只能用 v0.0.0 某个具体版本的。

有了这些经验之后，那么，我们来到 resty 包的 go.mod 文件中查看一下。

![查看resty包的依赖情况](https://upload-images.jianshu.io/upload_images/14623749-39c961225bb33afb.png?imageMogr2/auto-orient/strip%7CimageView2/2/w/1240)

我们可以看到，当前最新的 resty v2.12.0 版本是支持 go 1.16 的，但是问题就在于此时用的 `golang.org/x/net` 包是 v0.22.0 版本，通过以上我们的结论来看 v0.22.0 版本是不支持 go 1.16 的，那么，问题就出现在这里了！

![定位到具体resty包](https://upload-images.jianshu.io/upload_images/14623749-9001a183087f44c7.png?imageMogr2/auto-orient/strip%7CimageView2/2/w/1240)

找到问题所在了之后，我们就一个版本一个版本的去找，最后终于锁定了 `resty v2.7.0` 版本！

## 解决方案

1. 直接去项目的 `go.mod` 文件中添加 `require github.com/go-resty/resty/v2 v2.7.0`
2. 再去执行 `go mod download github.com/go-resty/resty/v2` 下载了指定 v2.7.0 版本的 resty 包
3. 再去运行一下项目，就发现跑起来啦！完美！

## 最后总结一下

1. 我本地之前是没有拉取过 `resty` 包的，因此当我使用 `go get` 命令且没有明确指定版本号的情况下，是直接会拉取最新版本的，也就是 `go get -v github.com/go-resty/resty`。
2. 其实不管是 `golang.org/x/sys` 包还是 `golang.org/x/net` 包，在我项目中并没**直接**引用它们，因此，根本就无法在 `go.mod` 文件中去调整它们的版本。另外，哪怕是存在 `go.mod` 文件中，但是后面带了有 `// indirect` 注释的，就算是我们在 `go.mod` 文件中删了这个依赖也没有啥用！

> indirect 标识表示该模块为间接依赖，也就是在当前应用程序中的 import 语句中，并没有发现这个模块的明确引用，有可能是你先手动 go get 拉取下来的，也有可能是你所依赖的模块所依赖的，情况有好几种。

3. 虽然这两个有冲突的包在 `go.mod` 文件中不存在，但是在 `go.sum` 文件中是存在的，可不可以直接在 `go.sum` 文件中将其修改或者删除呢？回答是：**不可以！** 因为此时你只需要执行下 `go mod tidy` 命令，马上又回来了。

其实遇到这个报错时，我也去搜索引擎上查询了一些解决方案，但是多数人的回答都是要升级 go 版本。可是，这个项目已经在生产环境上运行了一段时间了，能够稳定运行的项目，肯定是不能做调整大版本这种大动干戈动作的，不然，出了大问题，可能就直接拿我祭天了。

要是你也遇到了类似的问题，你也不想通过升级 go 版本来解决，也可以试试我说的这个方案。
