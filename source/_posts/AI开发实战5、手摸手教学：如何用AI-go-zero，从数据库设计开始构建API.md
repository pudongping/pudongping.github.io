---
title: AI开发实战5、手摸手教学：如何用AI+go-zero，从数据库设计开始构建API
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
abbrlink: 2f6d4246
date: 2026-04-22 12:41:38
img:
coverImg:
password:
summary:
---

这篇文章是从 0 到 1 使用 AI 开发完整项目的第 5 篇文章，也是讲解用 AI 开发后端 API 接口的第一篇文章。

关于用 AI 来开发 API 接口的文章不会讲太多，因为其实大部分的框架部分最好还是自己写，用 AI 来辅助为好，当你把项目框架写清楚之后，简单的业务逻辑部分就可以直接交给 AI，让 AI 在你限定的框架之下去编程，这样不仅可以高效率的写代码而且还可以很好的保证代码风格的一致性。

好了，废话就不多说了。

直接拿我最近刚刚撸完的一个开源项目——**Momento（时光账记）** 后端 API 为例，手把手教大家如何利用 AI（Claude 3.5 Sonnet / Gemini Pro 等模型）从 0 到 1 构建一个高性能的 Go 微服务后端。

在讲解之前，先简单介绍一下**时光账记**小程序：

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

## 项目背景与技术选型

我们要做的项目是一个**记账小程序后端**，核心功能包括：记账、多账本管理、多人协作、报表统计等。

技术栈：
- 语言：Go 1.25+
- 框架：go-zero（高性能微服务框架，自带代码生成神器 goctl）
- 数据库：MySQL 5.7+
- 缓存：Redis
- 工具库：Squirrel (SQL 构建), Cast (类型转换)等

## 💡 第一步：架构设计与数据库建模

看过我这个系列文章的同学应该都知道，当时我写这个小程序的时候，**我是先写的小程序端，然后才写的 API 接口**。这么做的好处就是：

- 整个项目本来就只有我一个人开发，先写页面，我完全可以确定好哪些功能需要，哪些功能可以后面再加
- 用 AI 写的这个项目能不能符合我的设定，如果不能符合，可能我放弃的也可以早一点儿，就不那么折腾了

当我把页面确定好了之后，所有的数据，我都让 AI 帮我填充了 mock 数据，并且是模拟的接口数据，这样当我觉得功能上没有什么大的问题的时候，我就直接让 ai 以 mock 数据为基础，帮我写好接口文档，并备注清楚特殊逻辑。

有了这份接口文档，我就可以让 ai 根据接口文档帮我设计数据库和开发真实的接口了。

### 🤖 提示词 (Prompt) 1：数据库设计

> 我正在开发一个基于微信小程序的家庭记账应用“时光账记”。请根据我提供的接口文档帮我设计 MySQL 数据库 Schema。
>
> **核心需求：**
> 1. **用户体系**：支持微信 OpenID 登录，记录昵称、头像。
> 2. **账本 (AccountBooks)**：一个用户可以创建多个账本，支持多人协作（多对多关系）。
> 3. **交易记录 (Transactions)**：核心表，记录金额、类型（收入/支出）、分类标签、备注、时间、图片凭证等。
> 4. **标签体系 (Tags)**：支持系统默认标签和用户自定义标签。
> 5. **周期性账单 (Recurring)**：支持每天/每周/每月自动记账。
>
> **技术要求：**
> - 使用 MySQL 5.7 语法。
> - 所有表包含 `created_at`, `updated_at` 字段，类型为 int(11) (存秒级时间戳)。
> - 字段要有详细注释。
> - 使用 snake_case 命名规范。
> - 考虑到查询性能，请适当添加索引。

当然了，你也可以更加详细的给 AI 提要求，比如：

```sql
-- =============================================
-- 「时光账记」数据库建表语句
-- 数据库版本: MySQL 5.7
-- 字符集: utf8mb4
-- 排序规则: utf8mb4_unicode_ci

-- 规范如下：
-- 1. 时间戳: 所有时间戳字段均为秒级时间戳（INT(11) UNSIGNED）
-- 2. 每张表都会有创建时间（created_at）和更新时间（updated_at）字段，类型均为 INT(11) UNSIGNED NOT NULL DEFAULT 0
-- 3. 每张表除了特殊说明，均会有自增主键 id 字段，并且类型为 BIGINT(20) UNSIGNED NOT NULL AUTO_INCREMENT
-- 主键命名规范为 “表名单数_id” 比如用户表 表名为 users 主键则为 user_id
-- 4. 如果表中需要自定义排序字段时，统一使用 sort_num 字段，类型为 INT(11) UNSIGNED NOT NULL DEFAULT 0
-- 5. 如果表中需要自定义状态字段时，统一使用 status 字段，类型为 TINYINT(1) NOT NULL DEFAULT 0
-- 6. 如果表中需要用到枚举类型时，统一使用 smallint(1) NOT NULL DEFAULT 0 并且一般起始值不为 0，比如 1-未支付 2-已支付
-- 7. 如果表中需要用到金额类型时，统一使用 decimal(10,2) NOT NULL DEFAULT 0
-- 8. 如果表中需要用到时间、日期类型时，统一使用 int(11) unsigned NOT NULL DEFAULT 0
-- 9. 如果表中需要用到布尔类型时，统一使用 tinyint(1) NOT NULL DEFAULT 1 一般使用 1-是 2-否
-- 10. 软删除时，统一使用 deleted_at 字段，类型为 INT(11) UNSIGNED NOT NULL DEFAULT 0
-- 11. 表字段一般不用 not null，需要设置默认值
-- =============================================

-- =============================================
-- 索引优化说明

-- 1. 所有外键字段都建立了索引
-- 2. 常用查询条件字段建立了索引
-- 3. 时间字段建立了索引，便于按时间范围查询
-- 4. 组合索引遵循最左前缀原则
-- 5. 唯一索引用于保证数据唯一性
-- =============================================
```

AI 会直接给你一套完整的 `sql/momento.sql` 文件，包含 `users`, `account_books`, `account_book_members`, `transactions`, `tags` 等表的建表语句，甚至连索引都帮你建好了。

然后你就可以根据自己的实际情况，将 AI 帮你生成的 sql 文件修修改改调整成符合自己设想的。

下一步，就结合 sql 文件以及接口文档去定义 go-zero 框架的 `api` 文件

## 第二步：API 接口定义 (DSL) —— 让 AI 帮你“写契约”

go-zero 的核心优势是 **Design First（设计优先）**。我们需要先写 `.api` 文件，然后自动生成代码。

### 🤖 提示词 (Prompt) 2：API DSL 生成

> 基于你刚刚设计的数据库 Schema 和我提供的接口文档，请帮我编写 go-zero 框架的 `.api` DSL 文件。
>
> **项目结构要求：**
> - 主文件 `miniapp.api` 直接放在项目根目录下的 `dsl` 目录下，模块文件需要根据模块进行划分，放在 `dsl/` 模块目录下（如 `dsl/user/user.api`, `dsl/transaction/transaction.api`）。
> - 遵循 RESTful 风格。
>
> **规范要求：**
> - 每个接口都要有 Request 和 Response 结构体定义。
> - 使用 `@doc` 注解添加接口文档说明。
> - 列表接口需要包含分页参数 (`page`, `per_page`)。
> - 所有的 ID 字段在 JSON 中使用 string 类型（防止前端精度丢失），在 Go 结构体中使用 int64。

但是你会发现 AI 生成的 `.api` 文件也并不能百分百符合自己的要求，那么，此时就需要给 AI 更多的限定了。

我的做法是，直接在 `dsl` 目录下创建了一个 `API_STYLE_GUIDE.md` 文件，在这个文件中写清楚了我的各种要求。感兴趣的童鞋可以把项目拉下来之后进行查看。

![API_STYLE_GUIDE.md](https://upload-images.jianshu.io/upload_images/14623749-db18628551b38abe.png?imageMogr2/auto-orient/strip%7CimageView2/2/w/1240)

## 写业务代码

准备工作都已经做好了，那么如何让 AI 开始写符合自己代码风格的业务代码呢？

我的方法是，自己先写好“模版”代码。

1. 首先，我自己简单封装了一些适合 go-zero 框架的基础方法，可以见 `https://github.com/pudongping/momento-api/tree/master/coreKit` 统一了相应数据格式、错误处理、参数验证等（但是，没有过度的封装）
2. 自己去写一些基础的接口，比如登录、登出接口
3. 然后直接让 AI 根据我的代码风格写出接口文档中的其他接口即可
4. 当然了，还是得自定义 `CLAUDE.md` 文件，否则 AI 还是会“乱写”（下一篇文章就介绍写接口的 CLAUDE.md）

## 总结

通过上面这个项目的实战，我们可以总结出用 AI 编程的三个核心秘诀：

1.  Context is King（上下文为王）：不要上来就让 AI 写代码。先给它数据库 Schema，先给它项目结构介绍。AI 懂你的项目越多，写的代码越准。
2.  Step by Step（分步执行）：不要试图用一个 Prompt 生成整个项目。拆解任务：数据库 -> API 定义 -> 核心逻辑 -> 优化。
3.  Review & Refine（审核与微调）：AI 也会犯错（比如类型转换错误、包引用错误）。你需要扮演 Tech Lead 的角色，Review 它的代码，并指出错误让它修正。

在 AI 时代，编程的门槛确实已经降低了不少，但**架构设计能力**和**将业务转化为技术需求的能力**反而变得更重要了。掌握了这套方法论，你一个人就是一个开发团队！

🚀 **关注我，下期教大家如何用 AI 写接口的 CLAUDE.md，敬请期待～**

*本文基于真实项目 Momento API 编写，项目已开源，欢迎 Star！*

