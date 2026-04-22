---
title: AI开发实战6、抄作业吧！我优化了N遍的go-zero项目AI协作规范文件，一字不差全给你
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
abbrlink: 3c4544a6
date: 2026-04-22 12:42:20
img:
coverImg:
password:
summary:
---

这篇文章是从 0 到 1 使用 AI 开发完整项目的第 6 篇文章，这也是这个系列的最后一篇文章。

今天主要讲解的是，如何在使用 go-zero 框架写 API 服务时，写 `CLAUDE.md` 规范文件。

还不知道什么是 CLAUDE.md 的童鞋可以翻一翻前面的几篇文章了解一下。

其实今天这篇文章就有点儿水，没有太多要讲的内容，主要我会将我写好的 CLAUDE.md 文件直接分享给到大家，以供各位大佬们参考。

继我上一篇文章，我有说到，当我想用 AI 来辅助我编程时，我并没有一上来就让 AI 来帮我写代码，而是，我自己先封装好了一些通用方法，并且我还写好了几个常用的接口，比如：登录、登出……

![通用方法](https://upload-images.jianshu.io/upload_images/14623749-b2a68491e83389ba.png?imageMogr2/auto-orient/strip%7CimageView2/2/w/1240)


其实当我们写好了一两个接口的时候，我们就可以启动 `claude code` 来帮我们**自动**生成 CLAUDE.md 文件。

## 如何自动生成 CLAUDE.md？

我们直接通过命令行切换到我们的项目根目录下，然后执行 `claude` 命令，启动 `Claude Code` AI 编辑器，然后直接执行 `/init` 命令即可。

> 如何使用 `Claude Code` 也可以翻一翻我前面的教程，都是手把手教程，相信聪明的你，一看就会。要是不会的话，就多看几遍～

## 我的 CLAUDE.md 文件

当然，默认生成的 CLAUDE.md 文件可能并不能完全覆盖你项目中的所有要求，那么，你完全可以自行修修改改即可。

好了，下面我就要开始水文章了，直接将我写好的 CLAUDE.md 内容分享给各位看官老爷们

```markdown
# CLAUDE.md

本文件旨在为 Claude Code (claude.ai/code) 在此仓库中的工作提供指导。

## 项目概述

**Momento API** 是一个基于 [go-zero](https://go-zero.dev/) 微服务框架构建的微信小程序后端。项目名称为"时光账记"，是一款个人财务管理应用，支持交易记录、周期性账单管理和重要日期提醒。

**技术栈：**

- Go 1.25.5
- go-zero v1.9.4（高性能 REST 微服务框架）
- MySQL 5.7（数据持久化）
- Redis（缓存）
- JWT（身份认证）
- goctl v1.9.2（代码生成工具）

## 核心架构模式

### 请求流程：Handler → Request → Logic → Model

每个 API 接口都遵循以下标准流程：
1. **Handler**（`internal/handler/{module}/`）- 接收 HTTP 请求，验证参数，调用 Logic
2. **Logic**（`internal/logic/{module}/`）- 包含业务逻辑规则，调用 Model
3. **Model**（`model/`）- 从 MySQL Schema 自动生成，处理数据库操作（可选 Redis 缓存）

### 服务上下文与依赖注入

`internal/svc/serviceContext.go` 初始化所有依赖项（MySQL 连接、Redis 客户端、Model、中间件），并将它们注入到整个应用中。这是 go-zero 的标准模式。

### 代码生成工作流

项目大量使用 **goctl**（go-zero 的代码生成器）：
- `.api` 文件（DSL）定义 API 契约，存放在 `dsl/` 目录
- goctl 根据 `.api` 文件生成 handler 存根和路由
- goctl 根据 MySQL Schema 生成数据库 Model
- `goctlTemplates/1.9.2/` 中的自定义模板确保生成的代码符合项目规范

**重要：** 重新生成代码会覆盖 `internal/handler/routes.go` 文件和 `internal/types/` 下的文件，所以请不要手动修改这些自动生成的文件 - 始终编辑 `.api` 文件。

### 项目结构

项目遵循 go-zero 标准的 REST API 架构：


dsl/                      # API 定义文件（.api DSL 格式）
├── miniapp.api          # 主 API 入口（所有 @server 与 service 定义）
├── user/user.api        # 用户类型定义
├── tag/tag.api          # 标签类型定义
├── transaction/         # 交易类型定义
├── festival/            # 节日类型定义
└── accountBook/         # 账本类型定义

internal/
├── config/              # 配置结构体
├── handler/             # HTTP 处理器（文件自动生成）
├── logic/               # 业务逻辑层（自定义代码的地方）
├── middleware/          # HTTP 中间件（认证检查等）
├── svc/                 # 服务上下文（依赖注入容器）
├── service/             # 共享服务层（跨模块业务逻辑）
├── constant/            # Redis Key 和其他常量
├── requests/            # 请求验证规则
└── types/               # 请求/响应结构体（自动生成）

coreKit/                 # 可复用工具库（可用于其他 go-zero 项目）
├── errcode/             # 错误代码定义
├── httpRest/            # HTTP 辅助工具（错误处理、CORS 等）
├── responses/           # 响应格式化工具
├── jwtToken/            # JWT Token 处理
├── validator/           # 请求验证引擎
└── ctxData/             # 上下文数据工具

etc/                     # 配置文件（YAML 格式）
├── momentoapi.yaml      # 主配置文件（从 .local 复制）
└── momentoapi.yaml.local # 配置模板

sql/                     # 数据库 Schema 和迁移脚本
goctlTemplates/          # 自定义 goctl 模板（覆盖默认模板）
model/                   # 自动生成的数据库 Model（由 goctl 生成）
local_run.sh             # 开发辅助脚本（执行 goctl 命令）
momentoapi.go            # 应用程序入口点


## 常用开发命令

### 代码生成


# 从所有 .api 文件生成 Go 代码
make api
# 或者
./local_run.sh genapi

# 从 MySQL 表生成 Model（存放在 model/ 目录）
./local_run.sh model <table_name>
# 示例：./local_run.sh model users

# 格式化 .api 文件
docker run --rm -it -v $(pwd):/app kevinwan/goctl:1.9.2 api format --dir ./dsl/<filename>.api

# 生成 Markdown API 文档
./local_run.sh mddoc

# 初始化 goctl 模板（一次性设置）
./local_run.sh tplinit

# 直接运行任意 goctl 命令
./local_run.sh goctl <args>


### 构建与运行


# 构建应用
go build -o momento-api momentoapi.go

# 运行应用
go run momentoapi.go -f etc/momentoapi.yaml

# 查看 goctl 配置
make goctlenv

### 配置文件

配置从 `etc/momentoapi.yaml` 加载：
- **MySQL**：数据库连接字符串和凭证
- **Redis**：缓存连接配置
- **JWTAuth**：JWT 密钥和过期时间（秒为单位）
- **Server**：服务器主机和端口配置

## 关键规范与约定

### .api 文件结构（DSL）

所有 `.api` 文件必须遵循 `dsl/API_STYLE_GUIDE.md`：

1. **模块文件**（如 `dsl/user/user.api`）：
   - 仅包含 `type` 定义（请求/响应结构体）
   - 禁止包含 `@server` 或 `service` 块
   - 禁止包含 `syntax` 或 `info` 声明

2. **主入口文件**（`dsl/miniapp.api`）：
   - 包含所有 `@server` 和 `service` 定义
   - 从模块文件中导入类型
   - 定义路由、处理器和中间件

3. **类型命名**：
   - 使用大驼峰命名：`TagListReq`、`UserInfoResp`
   - 每个接口都要同时定义 Req 和 Resp（即使为空）
   - 响应中的 ID（如果明确指定为雪花算法 ID ） 则使用 `string` 类型（防止前端精度丢失）

4. **字段规范**：
   - 必须包含 `json` 标签
   - 可选字段：在 json 标签中添加 `,optional`，同时添加 `valid` 标签用于验证
   - 字段注释写在行尾（不是上方）
   - 示例：`Type string `json:"type,optional" valid:"type"` // expense-支出 income-收入`

5. **Handler 和路由名称**：
   - Handler：小驼峰命名（如 `tagList`、`userInfo`）
   - 路由：全小写，单词用 `/` 分隔（如 `/tags/list`、`/user/info`）

### 身份认证（JWT）

- Token 验证通过 `internal/middleware/authCheckMiddleware.go` 中的 `AuthCheckMiddleware` 进行
- JWT 工具位于 `coreKit/jwtToken/`
- 用户信息（包括用户 ID）存储在 JWT Claims 中，可通过上下文访问
- 在 `.api` 文件的 `@server` 块中添加 `AuthCheckMiddleware` 来保护接口

### 请求验证

- 请求验证规则定义在 `internal/requests/` 目录
- 在结构体字段上使用 `coreKit/validator/` 和 `valid` 标签
- 使用 `govalidator` 库执行验证
- 字符串长度验证的特殊规则：使用 `min_cn` 和 `max_cn` 标签（支持中文字符）
- 详见 `coreKit/validator/README.md`

### 错误处理

- 在 `coreKit/errcode/` 中定义自定义错误
- 所有错误转换为 `{"code": ..., "msg": "..."}` JSON 格式
- `momentoapi.go` 中的主错误处理器将错误转换为 HTTP 200 及相应的错误代码
- 切勿返回 HTTP 错误状态码；使用 200 和响应体中的错误代码

### 数据库 Model

Model 从 MySQL Schema 自动生成：

./local_run.sh model <table_name>


生成的 Model 支持：
- Redis 缓存（启用 `--cache=true` 时）
- 缓存 Key 前缀：`momento_api:cache:`
- 标准 CRUD 方法

不要手动编辑生成的 Model 文件 - 如果 Schema 变更，重新生成即可。

## 常见开发任务

### 添加新的接口

1. 在 `dsl/` 中编辑相关模块的 `.api` 文件，添加类型定义
2. 在 `dsl/miniapp.api` 中导入模块类型
3. 在 `dsl/miniapp.api` 中添加路由定义，包含 `@doc`、`@handler` 和中间件
4. 执行 `./local_run.sh genapi` 自动生成 handler 存根
5. 在 `internal/logic/{module}/` 中实现业务逻辑
6. Model 和 Handler 自动生成；只需实现 Logic

### 添加新的数据库表

1. 在 `sql/` 目录中创建 SQL 迁移脚本
2. 将迁移应用到数据库
3. 执行 `./local_run.sh model <table_name>` 生成 Model
4. Model 生成在 `model/` 目录，支持 Redis 缓存

### 修改请求/响应类型

1. 编辑 `.api` 文件（不是生成的类型文件）
2. 执行 `./local_run.sh genapi` 重新生成

### 跨模块业务逻辑

对于多个模块共享的业务逻辑，在 `internal/service/` 中创建服务，以便提升代码的复用性。

## 重要文件与模式

- `momentoapi.go` - 应用程序入口点和错误处理器配置
- `internal/svc/serviceContext.go` - 依赖注入容器；在此添加新的依赖
- `dsl/miniapp.api` - 中央路由和服务定义
- `dsl/API_STYLE_GUIDE.md` - 强制性的 API 规范
- `coreKit/errcode/` - 集中管理的错误代码定义
- `backend_api_specification.md` - API 契约和业务逻辑规范
- `local_run.sh` - goctl 操作辅助脚本；如需更改，调整 User_Path 和 Go_Version

## 开发贴士

- 添加/移除依赖后，始终运行 `go mod tidy`
- 在业务逻辑层使用 `coreKit/` 工具库 - 它们被设计为可在其他 go-zero 项目中复用
- 生成 Model 时，对性能关键的表使用 `--cache=true` 需要临时开启一下 `local_run.sh` 中的 `--cache=true` 参数
- 项目使用 sonyflake 进行分布式 ID 生成
- 参考 `backend_api_specification.md` 获取详细的 API 契约和业务规则
```

以上就是我用 AI 开发完整项目的全部教程了，希望对你也有所帮助～

以上项目成果我已经开源，就是下面这个**时光账记**小程序，希望各位看官老爷们帮忙点个 Star 支持一下吧。

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