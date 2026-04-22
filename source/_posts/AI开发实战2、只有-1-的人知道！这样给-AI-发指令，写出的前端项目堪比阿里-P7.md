---
title: AI开发实战2、只有 1% 的人知道！这样给 AI 发指令，写出的前端项目堪比阿里 P7
author: Alex
top: false
hide: false
cover: false
toc: true
mathjax: false
categories: AI
tags:
  - AI
  - AI 编程
  - 微信小程序
abbrlink: b23f5e1a
date: 2026-04-22 12:36:02
img:
coverImg:
password:
summary:
---

为什么别人用 AI 写出的代码结构清晰、风格统一，而你写出来的却是“屎山”堆积？区别不在于 AI 的智商，而在于你会不会“调教”。

本文将为你提供一套经过实战验证的 **“从 0 到 1 前端项目启动 Prompt 清单”**，直接复制粘贴过去，修修改改，让 AI 瞬间变身你的资深技术合伙人。

## 为什么你用 AI 写的代码很乱？

很多小白用户在开始一个新项目时，习惯直接对 AI 说：
> “帮我写一个记账小程序的首页。”

这就好比你对装修队说：“帮我装修一下房子。”
结果可想而知：客厅是欧式的，卧室是中式的，厕所是战损风的。

如果你想让 AI 写出 **代码风格统一、色彩搭配和谐、交互逻辑一致** 的专业级项目，你必须学会**分步骤、模块化**地给指令。

下面就是我总结的一套**“四步走”标准化流程**，以我的开源项目 **Momento**（Uni-app + Vue3）为例，教你如何从 0 到 1 驾驭 AI。

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

## 🛠️ 第一步：定基调 —— 技术栈与架构约束

在写第一行代码之前，先发这条指令，把 AI 的“思维”固定在正确的轨道上。

### 📋 复制这条 Prompt：

```markdown
# Role
你是一名拥有 10 年经验的前端架构师，擅长 Uni-app 和 Vue 3 开发。

# Project Goal
我们要从零开始开发一个名为 "Momento" 的个人记账微信小程序。

# Tech Stack (强制执行)
1. **框架**：Uni-app + Vue 3 (Composition API / <script setup>)。
2. **样式**：SCSS。
3. **构建工具**：Vite。
4. **代码风格**：
   - 必须使用 ES6+ 语法。
   - 变量命名使用 camelCase（小驼峰）。
   - 组件命名使用 PascalCase（大驼峰）。
   - 禁止使用 TypeScript（根据项目实际情况调整）。

# File Structure
请先帮我规划项目的目录结构，参考标准的 Uni-app 项目结构，并解释每个目录的作用。
```

**✅ 效果**：AI 会给出一个标准的目录结构，并且在后续的对话中，它会时刻记得“我要用 Vue 3 Setup 语法”，而不会突然给你蹦出个 Vue 2 的写法。

---

## 🎨 第二步：定颜值 —— 注入设计基因

不要让 AI 自由发挥颜色！一旦它放飞自我，你的 App 就会变成“霓虹灯”。你需要先把“调色板”喂给它。

### 📋 复制这条 Prompt：

```markdown
# Design System Definition
我们需要定义一套全局的 UI 设计规范，请帮我生成一个 `uni.scss` 文件。

# Requirements
1. **主色调**：温暖的橙色 (#FF9A5A)。
2. **辅助色**：
   - 成功：#4CD964
   - 警告：#F0AD4E
   - 错误：#DD524D
3. **中性色**：
   - 主要文字：#333333
   - 次要文字：#999999
   - 边框颜色：#EEEEEE
   - 背景色：#F8F8F8
4. **圆角**：
   - 小圆角（按钮）：100rpx
   - 大圆角（卡片）：24rpx
5. **间距**：
   - 基础间距：20rpx
   - 倍数间距：10rpx, 30rpx, 40rpx

# Action
请生成代码，并告诉我如何在后续开发中引用这些变量。
```

**✅ 效果**：你得到了一份包含 `$uni-color-primary` 等变量的 SCSS 文件。以后你只需要说“用主色”，AI 就会自动填入变量名，而不是硬编码颜色值。

---

## 🧩 第三步：造轮子 —— 封装基础组件

这是最关键的一步！**千万不要直接开始写页面！** 先让 AI 把通用的按钮、弹窗、卡片封装好。

### 📋 复制这条 Prompt：

```markdown
# Component Development
为了保持 UI 统一，我们需要先封装几个基础组件。

# Tasks
请帮我编写以下 Vue 3 组件（存放在 `components/` 目录下）：

1. **BaseButton.vue**：
   - 支持 `type` 属性 (primary/default/danger)。
   - 默认圆角为全局变量中的小圆角。
   - 支持点击事件。
   
2. **BaseCard.vue**：
   - 白色背景，带有轻微阴影。
   - 圆角为全局变量中的大圆角。
   - 内部有 20rpx 的 padding。

3. **BaseModal.vue**：
   - 居中弹窗，带有半透明遮罩。
   - 包含标题栏、内容区、底部按钮区。
   - 动画效果：淡入淡出。

# Constraint
所有样式必须使用上一步定义的 `uni.scss` 变量。
```

**✅ 效果**：你拥有了一套可复用的“积木”。以后写页面时，AI 就会直接调用 `<BaseButton>`，而不是每次都重写一遍 `<button class="btn">`。

---

## 🚀 第四步：写业务 —— 标准化开发流

现在基础打好了，可以开始写页面了。但别急，发指令时带上“参考系”。

### 📋 复制这条 Prompt：

```markdown
# Feature Request
请帮我开发“记账页面” (pages/record/index.vue)。

# UI Requirements
1. 页面背景色使用全局背景色变量。
2. 顶部是一个“金额输入框”，样式风格参考 `BaseCard`。
3. 底部有一个“保存”按钮，**必须使用我们封装的 BaseButton 组件**。
4. 点击保存后，如果金额为空，弹出一个提示框（使用 `uni.showToast`）。

# Code Style
- 使用 Vue 3 `<script setup>`。
- 引入 `uni.scss` 变量。
- 保持代码简洁，关键逻辑添加注释。
```

**✅ 效果**：AI 会乖乖地复用你之前定义的组件和变量，生成出的页面不仅代码整洁，而且和整个 App 的风格完美融合。

---

## 📝 总结

发现了吗？**用 AI 编程，核心不在于“写代码”，而在于“定规矩”。**

如果你能按照：
1.  **定技术栈** (Context)
2.  **定设计变量** (Tokens)
3.  **定基础组件** (Components)
4.  **写业务逻辑** (Features)

这四个步骤来引导 AI，哪怕你完全不懂代码，也能做出一个专业级的前端项目。

万一，万一，你也不知道做成什么样，你也可以直接先跟 AI 进行交流，比如，可以这样问 AI “我想做一个记账的小程序，我应该采用什么样的色系……” 然后你就可以继续让它帮你先写几个 demo 页面出来，你自己再根据 demo 页面去做选择。

总之，如果你要借助 AI 来帮你完成项目，你一定要学会做项目拆解。

下一篇文章，我会继续讲解如何要求 AI 来帮我们给前端项目封装好请求方法，希望对你有所帮助～
