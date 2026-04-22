---
title: 全网都在推 Claude Code，但只有这篇文章教你如何“真正”能用
author: Alex
top: false
hide: false
cover: false
toc: true
mathjax: false
categories: AI
tags:
  - AI
  - Claude Code
abbrlink: 293ebadc
date: 2026-04-22 12:33:17
img:
coverImg:
password:
summary:
---

Claude Code 这个 AI 神器想必已经不用过多介绍了吧，但是身边有很多朋友都说安装上了，但是总是没办法丝滑的使用，额～由于一些懂得都懂的原因，很多人都卡在网络上，这篇文章就教大家如何一步一步的从安装到能正常使用。

![](https://upload-images.jianshu.io/upload_images/14623749-b16e28c4ae033844.png?imageMogr2/auto-orient/strip%7CimageView2/2/w/1240)


话不多说，直接开撸！

首先，你得有科学上网。还不会的小伙伴们可以翻翻我以前的文章，也有类似的介绍。或者直接问问 AI 有很多成熟的方案，这里就不多说了。

## 安装 Claude Code

如果你的环境是：macOS, Linux, WSL，你可以直接执行下面的命令一键安装

```bash
curl -fsSL https://claude.ai/install.sh | bash
```

在 macOS 上，你还可以通过 Homebrew 工具进行安装

```bash
brew install --cask claude-code
```

Windows 用户可以通过 `WinGet` 进行安装

```bash
winget install Anthropic.ClaudeCode
```

Windows 用户也可以通过 `Windows PowerShell` 进行安装

```bash
irm https://claude.ai/install.ps1 | iex
```

就是如此丝滑，到此为止，就已经成功安装了 Claude Code！

然而，然而，当你满怀期待的去命令行中敲击 `claude` 命令时，你会发现确实安装成功了，但是随之而来还有其他的问题。

## 提示网络无法连接

这个问题并不是每一个人都会遇到，但是一旦遇到，你可以参考下面我的解决方案：

1. clashx 需要开启 **增强模式**
2. 在 `~/.claude/settings.json` 配置文件中增加命令行代理（需要根据你自身情况而定）比如：

```bash
{
  "env": {
    "HTTP_PROXY": "http://127.0.0.1:7890",
    "HTTPS_PROXY": "http://127.0.0.1:7890",
  }
}
```

## 免登录配置

当你好不容易解决了网络无法连接的问题，当你再次使用 `claude` 命令去启动 Claude Code 时，你会发现此时它要你登录，然而，你就是不想要去注册 Anthropic 的账号，但是就是想使用 Claude Code 那，怎么办呢？

![](https://upload-images.jianshu.io/upload_images/14623749-772afd2bdbd46a5c.png?imageMogr2/auto-orient/strip%7CimageView2/2/w/1240)

开启免登录配置！

为了跳过登录，我们需要在 `~/.claude.json` 中，加入设置 `"hasCompletedOnboarding": true`  然后重新打开终端即可。

## 其他模型接入 Claude Code

我们也可以使用其他的模型（比如：DeepSeek、Kimi、智谱大模型，包括一些中间模型服务商提供的服务）来接入 Claude Code

使用方式也非常简单

直接编辑 `vim ~/.claude/settings.json` 配置文件，这里以接入 DeepSeek 为例

```bash
{
  "env": {
    "ANTHROPIC_AUTH_TOKEN": "你自己申请的 API key",
    "ANTHROPIC_API_KEY": "你自己申请的 API key",
    "ANTHROPIC_BASE_URL": "https://api.deepseek.com/anthropic",
    "API_TIMEOUT_MS": "3000000",
    "CLAUDE_CODE_DISABLE_NONESSENTIAL_TRAFFIC": "1",
    "HTTP_PROXY": "http://127.0.0.1:7890",
    "HTTPS_PROXY": "http://127.0.0.1:7890",
    "ANTHROPIC_MODEL": "deepseek-chat"
  }
}
```

其中，参数说明

- ANTHROPIC_AUTH_TOKEN：需要填写你自己申请到的 API Key
- ANTHROPIC_API_KEY：需要填写你自己申请到的 API Key
- ANTHROPIC_BASE_URL：模型 api 请求地址
- CLAUDE_CODE_DISABLE_NONESSENTIAL_TRAFFIC：禁用非必要的流量
- ANTHROPIC_MODEL：指定使用的模型名称

设置成功之后，你可以通过运行 claude 启动 Claude Code 然后输入 `/status` 来确认模型的状态，如果不是，可以通过输入 `/config` 来切换模型。
