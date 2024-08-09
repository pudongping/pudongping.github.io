---
title: 手摸手教你，从0到1开发一个Chrome浏览器插件
author: Alex
top: false
hide: false
cover: false
toc: true
mathjax: false
categories: 前端
tags:
  - 前端
  - Chrome
abbrlink: '2980e851'
date: 2024-08-09 12:13:02
img:
coverImg:
password:
summary:
---

开发 Chrome 浏览器插件（也称为扩展）是一段有趣且有成就感的过程。在本教程中，我将引导你从零开始，逐步创建一个简单的 Chrome 插件。无论你是编程新手还是有一定基础的用户，我们都将以简单易懂的方式介绍整个过程。

## 1. 什么是 Chrome 插件？

Chrome 插件是可以添加到 Google Chrome 浏览器中的小程序，旨在增强浏览器的功能。它们可以改变网页的外观、增加新的功能、与用户交互等。

我们先从一个最简单的 Chrome 扩展开始，你也可以参照我写的这个微信公众号小助手插件 `https://github.com/pudongping/mp-vx-insight` 来学习，好了，话不多说，直接开撸。

## 2. 准备工作

在开始之前，你需要确保：

- 已安装 Google Chrome 浏览器。
- 有一个简单的文本编辑器（如 Notepad、VS Code、Sublime Text）。
- 对 HTML、CSS 和 JavaScript 有基本了解。

## 3. 插件的基本结构

一个 Chrome 插件通常由以下几个基本文件组成：

- `manifest.json`：插件的配置文件，定义插件的基本信息和权限。
- `background.js`：插件的后台脚本，负责执行后台任务。
- `popup.html`：用户点击插件图标时显示的界面。
- `style.css`：用于美化插件界面的样式表。

## 4. 创建你的第一个插件

### 步骤 1：创建项目文件夹

在你的计算机上创建一个新的文件夹，例如 `my_first_extension`。

### 步骤 2：创建 manifest.json 文件

在项目文件夹中创建一个文件 `manifest.json`，并复制以下内容：

```json
{
  "manifest_version": 3,
  "name": "My First Extension",
  "version": "1.0",
  "description": "This is my first Chrome extension!",
  "action": {
    "default_popup": "popup.html",
    "default_icon": {
      "16": "icon16.png",
      "48": "icon48.png",
      "128": "icon128.png"
    }
  },
  "background": {
    "service_worker": "background.js"
  },
  "permissions": ["activeTab"]
}
```

### 步骤 3：添加 Popup 界面

在同一文件夹中，创建 `popup.html` 文件并添加以下代码：

```html
<!DOCTYPE html>
<html>
<head>
    <title>My First Extension</title>
    <link rel="stylesheet" href="style.css">
</head>
<body>
    <h1>Hello, Chrome!</h1>
    <button id="clickMe">Click Me!</button>
    <script src="popup.js"></script>
</body>
</html>
```

### 步骤 4：创建样式文件

在同一文件夹中，创建 `style.css` 文件，添加如下内容：

```css
body {
    width: 200px;
    text-align: center;
}

h1 {
    font-size: 16px;
    color: #333;
}

button {
    padding: 10px 15px;
    font-size: 14px;
}
```

### 步骤 5：添加 JavaScript 功能

接下来，创建 `popup.js` 文件，实现按钮的点击事件：

```javascript
document.getElementById('clickMe').addEventListener('click', function() {
    alert('Button clicked!');
});
```

### 步骤 6：添加后台脚本

为了展示后台功能，创建一个 `background.js` 文件，内容可以是简单的 console.log：

```javascript
console.log('Background service worker running!');
```

### 步骤 7：添加图标

为了使你的插件更美观，我们需要为其添加图标。在项目文件夹中添加三种不同尺寸的图标：`icon16.png`、`icon48.png` 和 `icon128.png`。你可以使用在线图标生成器生成图标，或从网络下载适合的图标。

## 5. 在 Chrome 中加载扩展

1. 打开 Chrome 浏览器，输入 `chrome://extensions/` 并按回车。
2. 在右上角启用开发者模式。**一定要开启**
3. 点击 “加载已解压的扩展”，选择你的项目文件夹（`my_first_extension`）。
4. 你会在扩展列表中看到你的插件。

## 6. 测试插件

点击浏览器工具栏中的插件图标，会弹出你定义的 Popup 界面。点击按钮，你应该会看到提示框弹出，显示 “Button clicked!”。

现在，你已经成功创建了一个简单的 Chrome 插件！是不是非常简单呢？当然还有更多的一些特性由于篇幅的原因就不细讲了，强烈建议你可以直接找一个 Chrome 插件源码看看，比如我写的这个微信公众号小助手 Chrome 扩展 `https://github.com/pudongping/mp-vx-insight` 这样可以学得更快！

## 7. 总结

通过这个简单的教程，你已经了解了如何从零开始开发一个基本的 Chrome 插件。插件的开发不仅能提升你的编程技能，更能让你在日常浏览中享受到便利。希望你在这个过程中感受到乐趣。

如需进一步学习，可以参考 [Chrome 扩展文档](https://developer.chrome.com/docs/extensions/mv3/getstarted/)，深入了解不同的 API 和功能。