---
title: 用 Go 写桌面应用？试试 Wails 吧！
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
  - Wails
abbrlink: f3f6160c
date: 2025-08-07 10:12:34
img:
coverImg:
password:
summary:
---

在前端开发中，提起桌面应用，很多人第一反应是 **Electron**。虽然它很流行，但不可否认也“很重” —— 内存占用高、打包体积大。

有没有更轻量的选择呢？当然有！今天我们来聊一聊一个非常有意思的项目 —— **Wails**。

Wails 是一个用 **Go + 前端技术** 开发桌面应用的框架，和 Electron 类似，但更轻、更快、更 Go。

> 其实这些内容本来不打算单独写一篇文章来介绍的，奈何 wails 官网上的教程实在是看得辣眼睛，因此，在我学习过程中，我简单整理一下，供有缘人参考吧。另外，如果你想快速把 wails 用起来，我建议你可以直接拿我的这个开源项目 `https://github.com/pudongping/wx-graph-crawl` 去修修改改，可能会更快！

![https://github.com/pudongping/wx-graph-crawl](https://upload-images.jianshu.io/upload_images/14623749-90e6c3a665503b84.png?imageMogr2/auto-orient/strip%7CimageView2/2/w/1240)


* * *

## 一、Wails 是什么？

Wails 是一个 Go 开源项目，它的核心理念是：

> 用 Go 写后端逻辑，用 HTML/CSS/JS 写前端界面，最终打包成一个原生桌面应用。

简单来说，你可以用你熟悉的 Go 来写后端逻辑，配合 Vue、React 或原生 HTML/CSS，快速开发桌面应用，而且性能很不错！

* * *

## 二、为什么选择 Wails？

*   ✅ **跨平台支持**：Windows、macOS、Linux 都能跑

*   ✅ **原生性能**：启动速度快，资源占用小

*   ✅ **开发体验好**：前后端联调方便，开发者工具好用

*   ✅ **打包方便**：只需一个命令，轻松构建

* * *

## 三、安装 Wails

### 1\. 环境准备

在开始之前，请确保你已经安装了以下内容：

*   Go（推荐版本：1.20 及以上）
*  npm（用于前端构建）
*   Git

你可以通过以下命令检查版本：

```bash
go version
node -v
npm -v
```

### 2\. 安装 Wails CLI

```bash
go install github.com/wailsapp/wails/v2/cmd/wails@latest
```

安装完成后，确保 `$GOPATH/bin` 已加入环境变量。

测试是否安装成功：

```bash
wails doctor
```

![wails doctor](https://upload-images.jianshu.io/upload_images/14623749-8d811f2203d0938d.png?imageMogr2/auto-orient/strip%7CimageView2/2/w/1240)


如果一切正常，会看到一些“ SUCCESS  Your system is ready for Wails development!”字样。

* * *

## 四、创建你的第一个 Wails 项目

Wails 提供了多个模板，我们以最基础的 Vue 模板为例：

使用 JavaScript 生成一个 vue 项目

```bash
wails init -n myapp -t vue
```

你也可以使用 TypeScript

```bash
wails init -n myapp -t vue-ts
```

进入目录：

```bash
cd myapp
```

* * *

## 五、项目结构简析

Wails 项目结构大致如下：

```bash
myapp/
├── app.go              # 应用主逻辑
├── main.go            # 程序入口
├── frontend/          # 前端代码
│   ├── index.html
│   ├── src/
│   └── dist/
├── build/             # 构建相关文件
└── wails.json         # Wails 配置文件
```

* * *

## 六、前后端联动示例

> 当初始化项目之后，wails 会提供一个默认的应用代码，你完全可以直接在项目根目录下执行 `wails dev` 命令，将项目跑起来。

我们写一个简单的例子：点击按钮，调用后端的 Go 函数，并在前端显示返回结果。

### 1\. 编写 Go 方法（后端）

在 `app.go` 中添加如下代码：

```go
package main

import "context"

// App 结构体会被 Wails 自动注册为后端服务
type App struct{}

// Hello 方法供前端调用
func (a *App) Hello(ctx context.Context, name string) string {
	return "你好，" + name + "！欢迎使用 Wails～"
}
```

* * *

### 2\. 前端调用 Go 方法

打开 `frontend/src/views/Home.vue`（或者根据模板结构找到首页组件），添加如下代码：

```vue.js
<script setup>
import { ref } from 'vue'
import { Hello } from '../../wailsjs/go/main/App'

const name = ref('小明')
const result = ref('')

const greet = async () => {
  result.value = await Hello(name.value)
}
</script>

<template>
  <div>
    <input v-model="name" placeholder="请输入你的名字" />
    <button @click="greet">打个招呼</button>
    <p>{{ result }}</p>
  </div>
</template>

```

这里通过 `wailsjs` 提供的封装文件，直接调用 Go 后端的方法，前端体验非常丝滑。

* * *

## 七、运行与打包

### 1\. 启动开发模式

```bash
wails dev
```

Wails 会自动构建前端和后端，并启动一个桌面窗口。你可以在里面看到刚刚写的页面和功能。

### 2\. 构建生产包

```bash
wails build
```

打包完成后，`build/bin` 目录下就会生成可执行文件，可以直接运行！

* * *

## 八、进阶建议

如果你熟悉 Go，可以在后端做很多事情，比如：

*   文件系统操作
*   系统调用
*   网络请求
*   与数据库交互

而前端就专注于界面和交互，完全可以复用你在 Web 项目中的经验。

Wails 的核心设计理念就是：**让前端更前端，让后端更后端**。

* * *

## 九、常见问题与小贴士

### Q1：前端 JS 怎么调用新的 Go 方法？

每次你新增或修改了 Go 方法后，运行：

```bash
wails generate module
```

或者也可以直接运行

```bash
wails dev
```

它会自动生成对应的 `wailsjs` 模块。

### Q2：可以用 React 吗？

当然可以！Wails 支持 Vue、React、Svelte、Lit 等现代前端框架，喜欢哪个用哪个。

### Q3：打包后的应用体积大吗？

比 Electron 小很多，简单应用打完包大约在 **20 ~ 40 MB**，比同类产品优秀不少。

### Q4：在 go 方法中可以任意返回数据吗？

答案是：**不可以**！ 🙅

我们先来看一下文档中介绍的内容：

![https://wails.io/zh-Hans/docs/howdoesitwork/#%E8%B0%83%E7%94%A8%E7%BB%91%E5%AE%9A%E7%9A%84-go-%E6%96%B9%E6%B3%95](https://upload-images.jianshu.io/upload_images/14623749-eb0a655c091b2ea9.png?imageMogr2/auto-orient/strip%7CimageView2/2/w/1240)

从上面红框中框出来的内容可以得知：我们从 Go 程序中返回出的数据，默认第一个返回值会被作为 Promise 的 resolve 去处理，第二个返回值会被作为 reject 去处理，因此，一定要注意，在 Go 中，我们的返回值最多只能返回 2 个，第一个返回值作为正常值，第二个返回值默认为错误 error。

那么，有的同学就会问了，如果万一就是非要返回多个返回值怎么办呢？

那就将第一个返回值用结构体包起来。总之：**在 wails 中，go 中的返回值只能返回 2 个，第一个返回值作为正常返回值，第二个返回值作为错误 error**。

* * *

## 十、结语

Wails 是一个非常值得关注的新生代桌面应用框架，特别适合：

*   想用 Go 写本地工具的小伙伴
*   有前端基础，又不想搞 Electron 的人
*   喜欢轻量、原生、性能好的桌面应用

它让我们用熟悉的技术栈，做出不一样的东西。

👉 **用 Wails 做桌面工具，你会发现原来这么爽！**

![](https://upload-images.jianshu.io/upload_images/14623749-de632c6646009187.png?imageMogr2/auto-orient/strip%7CimageView2/2/w/1240)

> 想要快速学习 wails 可以直接参考我的这个项目 `https://github.com/pudongping/wx-graph-crawl`


* * *

如果你觉得这篇文章对你有帮助，欢迎点赞、转发或留言讨论，我们下次见！

