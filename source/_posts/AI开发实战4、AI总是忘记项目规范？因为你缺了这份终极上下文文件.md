---
title: AI开发实战4、AI总是忘记项目规范？因为你缺了这份终极上下文文件
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
abbrlink: 66f0cba6
date: 2026-04-22 12:40:53
img:
coverImg:
password:
summary:
---

这篇文章是从 0 到 1 使用 AI 开发完整项目的第 4 篇文章，这也是讲前端相关的最后一篇文章，下一篇文章开始讲解如何通过 AI 来用 go-zero 写后端接口。

这个系列的文章不是基础编程文章，不会告诉你编程语法是怎样的，而是告诉你用 AI 来辅助写代码，重点偏向于提示词部分。

## 什么是 CLAUDE.md？

`CLAUDE.md` 是一个专门为 Claude Code（或者 TRAE 等 AI 编辑器）准备的**项目指南文件**。它的主要作用是帮助 AI 快速理解你的项目结构、开发规范和常用指令，从而更准确、更高效地协助你编写代码。

简单来说，它就像是给 AI 看的“入职手册”。

## CLAUDE.md 在 AI 编程中有什么作用？

1.  **快速建立上下文 (Context Awareness)**
    *   **项目概览**：告诉 AI 这个项目是做什么的（例如：`Momento` 是一个 Uni-app 开发的记账小程序），使用了什么技术栈（Vue 3, SCSS, Mock 等）。
    *   **目录结构**：解释各个文件夹和关键文件的作用（例如：`api/` 是接口层，`mock/` 是模拟数据），让 AI 知道去哪里找代码，去哪里改代码。

2.  **统一代码风格与规范 (Consistency)**
    *   **命名约定**：规定变量、函数、文件的命名方式（例如：API 方法必须以 `Api` 结尾）。
    *   **架构模式**：说明数据流是如何传递的（例如：API -> Request -> Mock -> Storage），确保 AI 生成的代码符合现有的架构设计，而不是“自作主张”地引入新模式。

3.  **提高开发效率 (Efficiency)**
    *   **常用命令**：列出启动、构建、测试项目的具体命令（例如：如何运行到微信开发者工具），AI 可以直接参考这些命令来协助你操作。
    *   **环境配置**：说明如何切换 Mock 和真实环境，避免 AI 在环境配置上犯错。

4.  **减少幻觉与错误 (Accuracy)**
    *   通过明确列出“重要文件说明”和“关键实现细节”（如雪花算法生成设备 ID），AI 不会去瞎猜核心逻辑的实现方式，而是直接遵循文档中的描述。

## 举个例子🌰

在我的**时光账记**小程序项目中 `CLAUDE.md` 文件中，有这样一段：

```markdown
### API 集成模式
所有 API 方法遵循命名约定: <operation><Entity>Api(). 示例:
- getTransactionsApi(params) → GET /transactions/list
```

如果没有这段话，当我们让 AI “写一个获取交易列表的接口”时，可能会写成 `fetchTransactionList` 或者 `getTransactionData`。
但因为有了 `CLAUDE.md`，AI 就会自动将其命名为 `getTransactionsApi`，并遵循你项目中 `api/request.js` 的封装方式，从而保证代码风格的高度一致。

当我们每次跟 AI 编辑器进行交流时，他就会默认先阅读 `CLAUDE.md` 文件中的内容，可以看下面这张图所示

![](https://upload-images.jianshu.io/upload_images/14623749-3e553b944f534340.png?imageMogr2/auto-orient/strip%7CimageView2/2/w/1240)

## 如何使用？

使用方式非常简单，直接在项目根目录下创建一个文件名为 `CLAUDE.md` 的 markdown 文件即可，如下是我**时光账记**小程序的 `CLAUDE.md` 内容，以供参考：

```markdown
# CLAUDE.md

此文件为 Claude Code (claude.ai/code) 在处理本仓库代码时提供指南。

## 项目概览

**Momento** 是一个使用 Uni-app 和 Vue 3 构建的个人财务管理微信小程序。它允许用户记录收入/支出、管理多个账本、设置预算、处理周期性交易，并查看节日倒计时。该应用支持与朋友/家人协作管理账本。

- **框架**: Uni-app (Vue 3 配合 Composition/Options API)
- **目标平台**: 微信小程序 (通过 Uni-app 支持其他平台)
- **样式**: SCSS 并在 `uni.scss` 中定义全局变量
- **后端**: REST API (在 `config/index.js` 中配置)

## 开发设置

### 环境配置

在 `config/index.js` 中切换环境：
- **开发/测试环境**: 注释掉生产配置，启用本地 `baseURL` 并设置 `useMock = true`

### 运行应用

1. 安装 **HBuilderX** 编辑器 (Uni-app 官方 IDE)
2. 安装 **微信开发者工具**
3. 将项目导入 HBuilderX: 文件 → 导入 → 从本地目录
4. 配置 `manifest.json`:
   - 在 `mp-weixin.appid` 下设置微信小程序 AppID
5. 运行: HBuilderX 菜单 → 运行 → 运行到小程序模拟器 → 微信开发者工具
6. 项目将自动编译并在模拟器中加载

### 开发用 Mock 数据

`mock/` 目录包含用于无后端开发的模拟 API 响应。在 `config/index.js` 中设置 `useMock: true` 以启用。Mock 结构与 `api/index.js` 中的 API 端点一一对应。

## 架构与代码组织

### 目录结构


api/                    # API 层
  ├── index.js         # 统一 API 导出 (所有方法以 "Api" 后缀)
  └── request.js       # 集成 Mock 和认证的 HTTP 请求处理
config/
  ├── index.js         # 环境配置与基础 URL
  └── permission.js    # 认证守卫的路由白名单
utils/
  ├── auth.js          # 认证辅助函数 (checkToken, logout 等)
  ├── time.js          # 日期/时间工具
  ├── account-book.js  # 账本计算逻辑
  └── snowflake.js     # 用于设备 ID 的雪花算法生成器
pages/
  ├── home/            # 仪表盘 (月度概览, 节日)
  ├── record/          # 交易记录页面
  ├── profile/         # 用户资料与 UID 显示
  ├── login/           # 微信授权登录
  ├── budget/          # 预算配置
  ├── festivals/       # 节日倒计时管理
  ├── account-books/   # 多账本管理与协作
  ├── edit-transaction/# 交易编辑
  └── ...
components/           # 可复用 Vue 组件
mock/                  # 开发用 Mock API 数据
styles/               # 全局 SCSS 文件
static/               # 图片, 图标, 底部标签栏资源
uni_modules/          # Uni-app 插件 (uni-icons, uni-scss)
App.vue               # 根组件与生命周期钩子
main.js               # 应用入口点
manifest.json         # 项目配置 (微信 AppID, 平台配置)
pages.json            # 路由与底部标签栏配置

### 数据流架构

1. **API 层** (`api/index.js`): 集中管理的 API 方法 (如 `getTransactionsApi`, `addTransactionApi`)
2. **请求处理器** (`api/request.js`):
   - 集成 Mock 的 HTTP 请求
   - 处理认证头 (Bearer token)
   - 设备追踪 (使用雪花算法生成 device_id)
   - 带有唯一请求 ID 的请求日志
3. **存储**: `uni.setStorageSync()` / `uni.getStorageSync()` 用于本地持久化
   - `token`: 认证令牌
   - `userInfo`: 用户资料数据
   - `device_id`: 设备标识符

### 认证流程

- `App.vue` 中的路由守卫拦截导航
- 白名单 (`config/permission.js`) 允许未认证访问: 登录页, 用户协议, 隐私政策
- 受保护路由需要存储中有有效的 `token`
- 微信 OAuth 登录后存储 Token
- `utils/auth.js` 提供辅助函数: `checkToken()`, `clearAuth()`, `logout()`, `checkLoginStatus()`

### 路由与页面配置

路由在 `pages.json` 中定义:
- **标签栏页面** (底部导航): home, record, profile
- **模态页面**: login, user agreement, privacy policy (自定义导航样式)
- **栈页面**: edit-transaction, budget, festivals, account-books
- 部分页面通过 `navigationStyle: "custom"` 使用自定义导航栏

## API 集成模式

所有 API 方法遵循命名约定: `<operation><Entity>Api()`. 示例:
- `getTransactionsApi(params)` → GET `/transactions/list`
- `addTransactionApi(data)` → POST `/transactions/add`
- `updateTransactionApi(data)` → PUT `/transactions/update`
- `deleteTransactionApi(id)` → DELETE `/transactions/delete`
- `uploadFileApi(filePath, type, business)` → POST `/upload/file` (multipart form)

使用 `api/request.js` 中的 `setRequestConfig()` 自定义请求行为 (基础 URL, 超时, 错误处理)。

## 常见开发任务

### 添加新 API 端点

1. 按照命名模式在 `api/index.js` 中添加方法:

   export const getNewDataApi = (params) => get('/path/to/endpoint', params);
   

2. 在组件中导入并使用:

   import { getNewDataApi } from '@/api/index.js';
   

### 添加开发用 Mock 数据

1. 在 `mock/` 中创建/编辑匹配端点的 mock 文件
2. 导出返回 mock 数据的处理函数
3. 在 `mock/index.js` 中导入并添加到 `mockApis` 对象
4. 在 `config/index.js` 中启用 `useMock: true`

### 修改路由与导航

- 更新 `pages.json` 以添加/修改页面和标签栏
- 使用 `uni.navigateTo()` 进行栈导航, `uni.switchTab()` 切换标签栏
- 受保护路由通过 `App.vue` 路由守卫和 `config/permission.js` 白名单进行检查

### 样式与全局变量

- 全局 SCSS 变量在 `uni.scss` 中
- 组件作用域样式在 `<style scoped>` 块中
- 使用 `uni.scss` 变量保持一致性 (如 `$uni-primary-color`)
- 响应式单位: `rpx` (响应式像素, 相对于 750px 屏幕宽度)

## 关键实现细节

### 设备追踪

- 雪花算法 ID 生成器 (`utils/snowflake.js`) 创建唯一设备 ID
- 作为 `device_id` 存储在本地存储中
- 通过 `api/request.js` 包含在所有 API 请求中

### 多账本管理

- 用户可以创建多个账本并邀请协作者
- 账本通过 `getAccountBooksApi()`, `createAccountBookApi()` 追踪
- 协作者通过邀请系统管理: `inviteUserApi()`, `acceptInvitationApi()`, `rejectInvitationApi()`
- 通过 `setDefaultAccountBookApi()` 设置默认账本

### 周期性交易

- 在创建交易时通过频率设置进行配置
- 单独管理: `getRecurringTransactionsApi()`, `deleteRecurringTransactionApi()`
- Mock 数据在 `mock/recurring-transactions.js`

### 预算与预测

- 通过 `updateUserSettingsApi()` 设置月度预算
- 用于预算追踪的交易统计: `getTransactionStats(params)`
- 预算 UI 在 `pages/budget/index`

## 测试与调试

### 启用控制台日志

日志显示在微信开发者工具控制台中:
- 应用启动时检查 `App.vue` 生命周期日志
- `api/request.js` 中记录请求 ID 以追踪 API 调用
- 首次应用启动时记录设备 ID 生成

### 切换真实 API 与 Mock

在 `config/index.js` 中:

// 用于 mock 数据 (开发环境):
export const useMock = true;

// 用于生产 API:
export const useMock = false;
export const baseURL = '用生产环境的接口地址';


### 存储检查

通过微信开发者工具检查持久化数据:
- Storage 面板显示 `token`, `userInfo`, `device_id` 等
- 开发期间使用 `uni.clearStorageSync()` 重置所有存储

## 微信小程序特性

- 在登录页获取微信 OAuth code，发送至后端交换 token
- 授权后自动同步用户信息 (头像, 昵称)
- 小程序能力: 文件上传, 设备信息, 存储, 推送通知
- 在 `manifest.json` 的 `mp-weixin` 部分配置

## 重要文件说明

| 文件 | 用途 |
|------|------|
| `api/request.js` | 核心 HTTP 逻辑，含 Mock 路由与认证 |
| `config/index.js` | 环境切换 (Mock/真实 API) |
| `config/permission.js` | 认证守卫的路由白名单 |
| `App.vue` | 应用生命周期与路由拦截设置 |
| `pages.json` | 所有路由定义与标签栏配置 |
| `utils/auth.js` | 登录/登出/认证状态辅助函数 |
| `mock/index.js` | Mock API 注册表 |
```

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

## 总结

`CLAUDE.md` 是连接我们开发者意图和 AI 执行能力的重要桥梁，维护好这个文件，能让你和 AI 的协作事半功倍。现在不仅仅只有 Claude Code 支持 CLAUDE.md 新版本的 TRAE 编辑器也已经支持读取 CLAUDE.md 了，赶紧去试试吧～

下一篇文章继续讲解借助 AI 来使用 go-zero 框架写后端 API 接口，敬请期待～
