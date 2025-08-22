---
title: Jaeger，一个链路追踪神器！
author: Alex
top: false
hide: false
cover: false
toc: true
mathjax: false
categories: 微服务
tags:
  - 微服务
  - jaeger
  - 链路追踪
abbrlink: f542f6b5
date: 2025-08-22 10:49:14
img:
coverImg:
password:
summary:
---

在微服务系统中，一个接口请求可能穿越十几个服务节点，复杂度远超传统单体应用。此时，如何追踪一次请求的全链路信息、快速定位问题、优化性能，成为了运维与开发必须直面的挑战。

> “如果你不能度量它，你就无法改进它。” —— 彼得·德鲁克

**链路追踪（Tracing）技术应运而生，而 Jaeger 作为业界主流的开源分布式链路追踪系统之一，是我们解决上述问题的强有力工具。**

本文将全面介绍 Jaeger 的核心概念、架构原理、使用方式以及在实际项目中的落地方法，助你快速上手并深入掌握链路追踪技术。

* * *

## 一、链路追踪的作用与价值

链路追踪是一种**用于监视和诊断基于微服务的分布式系统的机制**，它的价值体现在：

*   **服务依赖分析**：清晰展现调用链中每个服务的依赖关系；

*   **性能瓶颈定位**：通过调用耗时分析，精准识别性能短板；

*   **可视化调用链**：帮助开发者理解系统交互，辅助系统优化；

*   **排查线上问题**：快速定位服务错误、慢请求、丢失数据等问题。

简而言之，链路追踪是分布式系统“可观测性三剑客”（Metrics、Logs、Tracing）之一，它让我们真正**“看得见”**请求流转的全过程。

* * *

## 二、Jaeger：开源链路追踪解决方案

Jaeger 是由 Uber 开源的分布式追踪系统，具备高性能、高可扩展性，已成为 Cloud Native 领域的重要基础设施之一。

### Jaeger 的关键特性

*   🚀 **高扩展性**：基于微服务架构设计，易于横向扩展；

*   🔌 **原生支持 OpenTracing**：兼容 OpenTelemetry，方便接入各种语言与框架；

*   👁️ **出色的可观察性**：直观的 UI，展示完整调用链和详细追踪数据。

* * *

## 三、Jaeger 追踪模型核心术语解析

链路追踪的底层原理基于分布式上下文传播，其中的核心术语如下：

### 1\. Span：追踪的最小单元

在 Jaeger 中，一个 **Span 表示一个操作单元**（例如一个 HTTP 请求或数据库查询），每个 Span 包含以下信息：

| 字段 | 描述 |
| --- | --- |
| Operation name | 操作名称，也叫 Span name |
| Start timestamp | 起始时间 |
| Finish timestamp | 结束时间 |
| Span tag | 一组 key-value 键值对 |
| Span log | Span 中的事件日志 |
| SpanContext | 当前 Span 的上下文信息 |
| References | 当前 Span 与其他 Span 的引用关系 |

### 2\. Trace：由多个 Span 构成的调用链

一个完整的 Trace 是由多个 Span 组成的，它们通过上下游关系构成了调用链结构，如：

```
Trace
└── Span A（入口服务）
    ├── Span B（服务调用1）
    └── Span C（服务调用2）

```

通过构建 Trace，我们能复原出一次请求完整的调用路径和每个服务的耗时、状态。

* * *

## 四、Jaeger 的原理与调用链分析

来看一下 Jaeger 的工作原理图：

![ jaeger调用原理](https://upload-images.jianshu.io/upload_images/14623749-5bdbe760638c7818.png?imageMogr2/auto-orient/strip%7CimageView2/2/w/1240)

#### 一次典型调用链的生命周期：

![一次调用链分析](https://upload-images.jianshu.io/upload_images/14623749-cc2e26163e132f56.png?imageMogr2/auto-orient/strip%7CimageView2/2/w/1240)

#### 追踪数据在时间维度上的展示：

![调用链时间线](https://upload-images.jianshu.io/upload_images/14623749-1e082a56ccf652fe.png?imageMogr2/auto-orient/strip%7CimageView2/2/w/1240)

这些图形化视图能够帮助我们快速了解请求在哪个节点出现了延迟或者错误。

* * *

## 五、Jaeger 架构与核心组件

![jaeger 的组件](https://upload-images.jianshu.io/upload_images/14623749-9753c2c2fc3fad83.png?imageMogr2/auto-orient/strip%7CimageView2/2/w/1240)

Jaeger 的整体架构采用微服务模式，核心由以下五个组件构成：

### 1\. Jaeger-client（客户端库）

*   集成在服务代码中，生成并上报 Trace 数据；

*   支持多语言：Go、Java、Node.js、Python 等；

*   与 OpenTracing 接口兼容。

### 2\. Agent（客户端代理）

*   作为本地 Daemon 常驻进程运行；

*   接收客户端 Trace 数据并批量转发到 Collector；

*   降低应用程序的性能开销。

### 3\. Collector（收集器）

*   接收 Agent 转发的追踪数据；

*   负责数据的校验、转换和写入后端存储；

*   支持高并发部署与负载均衡。

### 4\. Data Store（后端存储）

*   存储 Trace 数据；

*   支持多种后端，如 Elasticsearch、Kafka、Cassandra 等；

*   可根据业务量和查询需求选择存储类型。

### 5\. UI（前端界面）

*   提供 Web 界面，供开发者查询、检索和分析 Trace 数据；

*   可视化展示调用链、时间线、依赖关系图等信息。

![原理](https://upload-images.jianshu.io/upload_images/14623749-a9af88732ecffe4a.png?imageMogr2/auto-orient/strip%7CimageView2/2/w/1240)

* * *

## 六、Jaeger 各组件的端口说明

为了方便部署和调试，下面是 Jaeger 常用组件及其对应的端口信息：

| 端口号 | 协议 | 所属模块 | 说明 |
| --- | --- | --- | --- |
| 5775 | UDP | Agent | 兼容 Zipkin 的 Thrift 数据 |
| 6831 | UDP | Agent | 接收 Jaeger Thrift 数据 |
| 6832 | UDP | Agent | 接收二进制 Thrift 数据 |
| 5778 | HTTP | Agent | Agent 配置控制接口 |
| 16686 | HTTP | Query/UI | 前端 Web 查询界面 |
| 14268 | HTTP | Collector | 接收客户端 Zipkin 数据 |
| 14267 | HTTP | Collector | 接收客户端 Jaeger 数据 |
| 9411 | HTTP | Collector | Zipkin 兼容接口 |

* * *

## 七、在代码中使用链路追踪（Jaeger）

在实际开发中，我们可以使用 Jaeger 客户端库来手动创建和管理追踪信息。以 Java 为例：

```
Tracer tracer = Configuration.fromEnv("my-service").getTracer();

Span span = tracer.buildSpan("http-request").start();
try {
    // your business logic
} finally {
    span.finish();
}

```

当然，在实际项目中可以借助框架自动注入追踪逻辑（如 Spring Cloud Sleuth + OpenTelemetry），以减少侵入性。

> ✅ **建议**：在系统中统一追踪上下文传递逻辑（如 TraceId），并通过中间件或拦截器封装，降低接入成本。

* * *

## 八、总结：为什么你离不开链路追踪

在分布式系统中，链路追踪技术已不再是“可选项”，而是微服务治理的基础能力。Jaeger 凭借其成熟的架构、良好的性能和活跃的社区，成为业界默认选择之一。

通过本文你应该掌握了：

*   链路追踪的基本作用与应用场景；

*   Jaeger 的核心术语与追踪原理；

*   Jaeger 的架构组件及端口信息；

*   在代码中如何集成 Jaeger 实现追踪。

> 📌 **建议每个中大型微服务系统尽早接入链路追踪系统，尽快构建系统的可观测性能力。**

* * *

最后欢迎大家留言交流你在使用 Jaeger 或其他链路追踪系统（如 Skywalking、Zipkin、OpenTelemetry）中的经验与踩坑故事。

可观测性之路不止一条，但最重要的是：尽早上路！

