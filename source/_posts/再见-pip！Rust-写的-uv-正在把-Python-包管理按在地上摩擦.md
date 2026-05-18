---
title: 再见 pip！Rust 写的 uv 正在把 Python 包管理按在地上摩擦
author: Alex
top: false
hide: false
cover: false
toc: true
mathjax: false
categories: Python
tags:
  - Python3
abbrlink: 5cc664da
date: 2026-05-18 11:01:19
img:
coverImg:
password:
summary:
---

如果你问，作为一名 Python 开发者，平时最让你头疼的事情是什么？

我相信 90% 的兄弟都会回答：**依赖管理和环境配置。**

这就是我们的日常：

* `pip install` 慢得像蜗牛，看着进度条发呆。
* 接手一个老项目，`requirements.txt` 装到一半报错，提示各种依赖冲突。
* 电脑里装了 Conda、Poetry、pyenv，乱成一锅粥，最后连自己用的哪个 Python 版本都搞不清了。

苦 `pip` 久矣！

但今天，我要给大家安利一个最近在技术圈火到爆炸的神器——**uv**。

用完它，我只有一种感觉：**以前的日子简直是再也回不去了。**

## ⚡️ 什么是 uv？为什么它这么快？

简单来说，`uv` 是一个极其快速的 Python 包安装器和解析器。

它最大的卖点就写在它的基因里：**它是用 Rust 语言编写的。**

大家都知道，Rust 以内存安全和极致性能著称。之前的代码格式化工具 `Ruff` 也是这家叫 **Astral** 的公司出的，当时就凭速度震惊了业界。现在，他们把魔爪伸向了 `pip`。

在某些测试场景下，`uv` 的速度是 `pip` 的 **10-100 倍**。

![](https://upload-images.jianshu.io/upload_images/14623749-26d0889cec11119c.png?imageMogr2/auto-orient/strip%7CimageView2/2/w/1240)

> **注意：** 这不是简单的网络下载速度快，而是它在“解析依赖关系”（Resolver）这一步快得离谱。当你的项目有几百个包互相依赖时，`pip` 还在算数学题，`uv` 已经交卷了。

## 🛠️ 上手实战：快到飞起

安装 `uv` 非常简单，即使你连 Python 都没装，也可以直接用命令行搞定（支持 macOS, Linux, Windows）。

以 macOS/Linux 为例：

```bash
# 官方推荐的安装方式
curl -LsSf https://astral.sh/uv/install.sh | sh

```

装好之后，我们来体验一下它的“丝滑”。

### 1. 创建虚拟环境

以前我们用 `python -m venv .venv`，现在只需要：

```bash
# 瞬间创建一个虚拟环境
uv venv

```

你会发现，它甚至不需要你手动激活环境（在某些命令下），不过为了符合习惯，我们还是可以激活它：

```bash
# macOS/Linux
source .venv/bin/activate
# Windows
.venv\Scripts\activate

```

### 2. 安装包：快如闪电

重点来了！我们来安装一个常用的库，比如 `requests`。

```bash
# 使用 uv 进行安装
uv pip install requests

```

你可以盯着屏幕看，但我估计你还没看清进度条，它就结束了。

而且，`uv` 的输出非常现代化，清晰地告诉你它干了什么，不像 `pip` 那样吐出一大堆红红白白的日志。

### 3. 替代 pip-compile

如果你以前用 `pip-tools` 来锁定版本，`uv` 也完美支持，而且速度快得多：

```bash
# 假设你有一个 requirements.in 文件
# uv 会瞬间帮你生成锁定的 requirements.txt
uv pip compile requirements.in -o requirements.txt

```

## 🤯 不止是快，它想做“全能王”

如果 `uv` 只是个快速版的 `pip`，那它还不至于让我这么兴奋。

最近 `uv` 的更新显示了它的野心：**它想接管你的整个 Python 工作流。**

* **Python 版本管理：** 它可以像 `pyenv` 一样帮你下载和安装不同版本的 Python。
```bash
# 比如你想用 Python 3.12 跑个脚本
uv run --python 3.12 app.py

```


如果你的电脑没装 Python 3.12，它会自动帮你下载下来，放在独立的目录里，不污染系统！
* **脚本执行：** 以后分享脚本，不用告诉别人“先装这个包，再装那个包”。直接在脚本里声明依赖，用 `uv run` 一键运行。

## ✍️ 总结与建议

**uv 会取代 pip 吗？**

短期内，`pip` 作为官方标准依然会存在。但在实际工程中，我强烈建议大家开始尝试 `uv`，特别是对于：

1. **CI/CD 流水线：** 能节省大量的构建时间，帮公司省钱。
2. **大型项目：** 依赖解析速度的提升能极大改善开发体验。
3. **Docker 镜像构建：** `uv` 的缓存机制非常优秀，能显著减小镜像体积和构建耗时。

技术在进步，千万别抱着旧工具不撒手。赶紧去试试 `uv`，体验一下“飞”一般的感觉！

你平时开发中最讨厌 Python 的哪一点？是环境配置难，还是包安装慢？欢迎在评论区吐槽！