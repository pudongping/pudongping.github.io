---
title: AI开发实战3、90%人用AI写前端都踩的坑：API层混乱！3步教你标准化
author: Alex
top: false
hide: false
cover: false
toc: true
mathjax: false
categories: AI
tags:
  - AI
  - AI编程
  - 微信小程序
abbrlink: 6ea7c65a
date: 2026-04-22 12:39:55
img:
coverImg:
password:
summary:
---

如果你也用 AI 写前端代码，你就会发现，AI 帮你写的前端请求代码总是五花八门？一会用 `axios`，一会用 `uni.request`，甚至把 API 地址直接写死在页面里？

那么，有同学就说了，这又不是不行，代码也是可以正常跑起来的呀，是的，代码确实是可以跑起来的，但是这就完全丢失了代码的易维护性，也不方便自己写 mock 数据。

本篇文章，就结合我最近写的一个开源项目——**时光账记**小程序，来跟大家一起聊聊，如何让 AI 生成标准、统一、可维护的 API 层代码。

## 时光账记

**时光账记**是一款基于 `Uni-app` + `Vue 3` 开发的个人记账微信小程序，后端接口基于 [go-zero](https://github.com/zeromicro/go-zero) 微服务框架构建。

这是一款专注于个人财务管理与生活记录的应用。它不仅支持非常简洁的方式来管理基础的收支记录，还提供了多账本管理、周期性自动记账、预算控制以及节日倒计时等贴心功能，帮助用户更好地管理个人及家庭财务。

> 现在我已将代码都开源了，感兴趣的朋友可以去观摩观摩，也请帮忙点个 Star 支持一下，谢谢！
>
> 小程序端（Uni-app + Vue3）： https://github.com/pudongping/momento-miniapp   
> API 接口（Go + go-zero）： https://github.com/pudongping/momento-api
>
> 前端部分 AI 占比 100%（自己一行代码都没写），接口部分 AI 占比 80%
> 这也是一套非常不错的 AI 练手项目，如果对你有帮助，希望帮忙点个 Star 支持一下，谢谢！

![homepage.png](https://upload-images.jianshu.io/upload_images/14623749-1ae7ffbb67bd966d.png?imageMogr2/auto-orient/strip%7CimageView2/2/w/1240)

![login.png](https://upload-images.jianshu.io/upload_images/14623749-75ccf883fe8b6132.png?imageMogr2/auto-orient/strip%7CimageView2/2/w/1240)

![profile.png](https://upload-images.jianshu.io/upload_images/14623749-60a548bb48cf25f5.png?imageMogr2/auto-orient/strip%7CimageView2/2/w/1240)

![recurring.png](https://upload-images.jianshu.io/upload_images/14623749-7e06994b19353f86.png?imageMogr2/auto-orient/strip%7CimageView2/2/w/1240)

![transaction.png](https://upload-images.jianshu.io/upload_images/14623749-1a023028d183f588.png?imageMogr2/auto-orient/strip%7CimageView2/2/w/1240)

---

## 😫 痛点：AI 写的代码，能跑但很难阅读

在开发前端项目时，我们通常会有统一的 `api/` 目录来管理所有接口。但是，当你直接告诉 AI：“帮我写一个获取用户信息的接口”时，它可能会给你这样的代码：

错误示范 1：直接在组件里写请求

```javascript
uni.request({
  url: 'https://api.example.com/user/info',
  method: 'GET',
  success: (res) => { ... }
})
```

后果：接口地址散落在各个页面，后续切换环境的话，改动太多。

错误示范 2：自作聪明的封装

```javascript
// AI 自己发明了一套封装
import request from '@/utils/http.js'; // 你的项目里根本没这个文件！
export function getUser() {
  return request.get('/user');
}
```

后果：引用了不存在的文件，或者风格与现有代码完全不符。

---

## 💡 破局：给 AI 立“规矩”

AI 就像一个刚入职的实习生，它不知道你项目的架构和规范。而你需要做的工作就是：**把项目里的“潜规则”变成显性的指令。**

以我的 **Momento** 小程序项目为例，我的 API 目录结构如下：

- `api/request.js`: 核心请求封装（处理 Token、Mock 数据、错误拦截等）。
- `api/index.js`: 统一导出所有业务接口。

我希望 AI 生成的代码必须满足以下条件：
1.  **引用规范**：必须从 `./request.js` 导入 `get`, `post`, `put`, `del` 方法。
2.  **命名规范**：函数名必须以 `Api` 结尾（如 `loginApi`，方便自己可以区分哪些方法是请求接口的方法，哪些不是）。
3.  **代码风格**：使用箭头函数。

效果类似这样

![](https://upload-images.jianshu.io/upload_images/14623749-12d98b4a8a2a6069.png?imageMogr2/auto-orient/strip%7CimageView2/2/w/1240)

---

##  生成 API 方法的 Prompt

为了达到这个目的，我编写了一段 **Prompt（提示词）**。你只需要把这段话发给 AI，它写出来的代码就基本能符合要求。

### 先要求 AI 封装好 `api/request.js` 文件

```markdown
# Role
你是一个高级前端架构师，擅长使用 Uni-app 和 Vue 3 进行移动端应用开发。

# Task
请帮我封装一个功能完善的 HTTP 请求库（类似 Axios 的封装），文件路径为 `api/request.js`。该封装需要基于 `uni.request` 和 `uni.uploadFile`，并包含以下核心功能：

# Requirements

## 1. 基础配置与环境
- 引入配置文件（如 `config/index.js`）中的 `baseURL` 和 `useMock` 开关。
- 支持默认配置对象，包含：超时时间（默认 15s）、错误提示开关、Token 前缀（如 'Bearer '）、Token Header 键名（如 'Authorization'）。
- 提供 `setRequestConfig` 方法允许外部动态修改配置。

## 2. 辅助功能
- **Request ID**：为每个请求生成唯一的 `X-Request-ID`（时间戳 + 随机数），用于日志追踪。
- **设备 ID**：实现 `getDeviceId` 方法，首次获取时生成并存储在本地缓存（`uni.setStorageSync`），后续直接读取，请求头中携带 `X-Device-ID`。
- **用户 ID**：从本地缓存用户信息中尝试获取 `user_id` 或 `uid`，若存在则在请求头中携带 `X-User-ID`。
- **URL 处理**：自动判断 URL 是否为绝对路径，若不是则拼接 `baseURL`。

## 3. Mock 数据支持 (关键)
- 在发起真实网络请求前，检查 `useMock` 开关。
- 如果开启 Mock 且请求 URL 在 `mock/index.js` 导出的 `mockApis` 映射中有对应处理函数，则拦截请求。
- 执行 Mock 函数，模拟网络延迟（如 300ms），并返回 Promise 结果。
- Mock 逻辑需同时支持普通请求和上传请求。

## 4. 拦截器与逻辑处理
- **请求拦截**：
  - 自动在 Header 中添加 Token（如果本地缓存存在）。
  - 添加上述的追踪 ID（Request ID, Device ID, User ID）。
- **响应拦截**：
  - **HTTP 状态码检查**：处理非 2xx 的 HTTP 状态码。
  - **业务状态码检查**：
    - 约定 `code === 0` 或 `code === 200` 为成功，直接返回 `data`。
    - **Token 过期处理**：当 `code === 401` 时，自动执行登出逻辑（清除本地 Token 和用户信息，提示用户，延时跳转回登录页）。
    - **其他错误**：默认弹出 Toast 提示错误信息（可配置关闭）。

## 5. API 方法导出
- 导出通用的 `request` 函数。
- 导出便捷方法：`get`, `post`, `put`, `del` (delete)。
- **文件上传**：导出 `upload` 方法，封装 `uni.uploadFile`，同样需要支持 Mock、Header 注入、Token 校验和统一的响应处理（注意 `uploadFile` 返回的 data 通常是字符串，需要尝试 `JSON.parse`）。
```

如果你需要添加接口的时候，你就可以直接复制下面的提示词给到 AI

### 创建新的接口时

```markdown
# Role
你是一名资深前端架构师，负责维护一个基于 Uni-app 的微信小程序项目。

# Context
我们需要在 `api/index.js` 文件中新增一组 API 接口。
该项目已经封装了核心请求库，位于同级目录下的 `request.js` 中。

# Constraints (必须遵守的规则)
1. **导入依赖**：你只能使用以下语句导入请求方法：
   `import { get, post, put, del, upload } from './request.js';`
2. **命名规范**：所有导出的函数名必须以 `Api` 为后缀（例如 `getUserInfoApi`, `createOrderApi`）。
3. **参数处理**：
   - GET 请求通常接收 params 对象。
   - POST/PUT 请求通常接收 data 对象。
   - DELETE 请求通常接收包含 id 的对象。
4. **代码风格**：必须使用 ES6 箭头函数，保持简洁。

# Code Reference (参考代码)
// api/index.js 现有代码范例
import { get, post, put, del } from './request.js';

// 用户相关
export const getUserInfoApi = () => get('/user/info');
export const loginApi = (code) => post('/user/login', { code });
```

当你将上面的两个提示词给到 AI 之后，如果当你有新的接口需要添加时，你就可以直接跟 AI 说，“保持项目现有风格，帮我添加 xxxx 接口”即可，不需要每次都重复发提示词。

当然了，以上封装 request.js 的代码在同类型项目中只需要写一次，下次有类似的项目的时候，可以直接拷贝过去，完全没有必要每次都重复写。

## 📝 总结

不要怪 AI 笨，很多时候是我们没有给出清晰的指令。

对于前端项目中的 API 层，**“约束”比“创造”更重要**。通过构建这样一个标准化的 Prompt，你可以：

1.  **节省 Code Review 时间**：不用再纠结命名和风格。
2.  **降低维护成本**：所有接口长得都一样，新人接手也能秒懂。
3.  **提高开发效率**：完全可以直接复制粘贴，修修改改后就能直接用。

赶紧把这个 Prompt 保存到你的笔记里，下次写接口时试一试吧！

经常用 AI 辅助开发的同学也发现了，用 AI 编辑器开发项目时，明明已经跟 AI 讲清楚了“规则”为什么写出的代码还是会不在规则之类？总感觉 AI 失忆了。是的，AI 不可能会一次性记住太多的规则，但是我们的项目中就需要要各种规则限定，那么，该怎么处理呢？

下一篇文章，我们就来讲解一下 `CLAUDE.md`。让 AI 每次写代码时，都有“记忆”。