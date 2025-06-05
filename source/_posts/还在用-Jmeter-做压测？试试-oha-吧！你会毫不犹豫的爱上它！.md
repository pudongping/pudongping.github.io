---
title: 还在用 Jmeter 做压测？试试 oha 吧！你会毫不犹豫的爱上它！
author: Alex
top: false
hide: false
cover: false
toc: true
mathjax: false
categories: 科普
tags:
  - Jmeter
  - oha
  - 压力测试
  - 压测
abbrlink: 869fb74d
date: 2025-06-05 10:44:44
img:
coverImg:
password:
summary:
---

在进行 Web 服务和 API 性能测试时，选择合适的工具至关重要。市面上有很多工具可以帮助开发者进行负载测试，其中 **OHA** 和 **JMeter** 是两个常见的选择。

今天，我们将一起探讨 **OHA**，一个基于 Rust 语言的高效 HTTP 性能测试工具，并与传统的性能测试工具 **JMeter** 进行简单对比，帮助你了解它们各自的优势和适用场景。

## 什么是 OHA？

Ohayou(おはよう) 是一个用 Rust 编写的轻量级 HTTP 负载生成器，它能够向 Web 应用发送负载，并实时显示 TUI。这个程序由 tokio 提供动力，并使用了由 ratatui 创建的美观的 TUI。

相比于许多传统的性能测试工具，OHA 的特点是简单易用，且性能非常高，能够在高并发场景下提供良好的测试体验。

## 为什么选择 OHA？

1. **高性能**：OHA 使用 Rust 语言编写，Rust 在性能和内存管理方面表现优异，这使得 OHA 在高并发负载下依然能够稳定运行，且资源占用较低。
2. **易于使用**：OHA 提供简单的命令行接口，只需要几行命令就可以开始性能测试，适合开发者快速上手。
3. **简洁的报告**：OHA 提供清晰的性能报告，显示关键性能指标，如吞吐量、响应时间、请求成功率等，帮助开发者迅速了解系统的瓶颈。
4. **开源**：作为开源项目，OHA 可以自由使用和修改，适合各种开发需求。

![来源自 oha 官网](https://upload-images.jianshu.io/upload_images/14623749-17fcb2b84c74cd5f.gif?imageMogr2/auto-orient/strip)


## OHA 和 JMeter 的对比

**JMeter** 是另一个广泛使用的性能测试工具，它同样能够模拟并发请求、生成负载并进行压力测试。JMeter 支持 HTTP、FTP、JDBC 等协议，并且提供了丰富的插件，适用于多种复杂的性能测试场景。以下是 OHA 和 JMeter 的一些关键对比：

| 特性          | OHA                                      | JMeter                                    |
| ------------- | ---------------------------------------- | ---------------------------------------- |
| **语言**      | Rust                                     | Java                                     |
| **性能**      | 高效，适用于高并发负载测试               | 性能较为一般，尤其在大规模并发时占用较高  |
| **使用难度**  | 简单，命令行工具，易于快速上手           | 需要较多配置和学习，图形化界面较为复杂   |
| **资源占用**  | 轻量级，内存占用低，适合小型服务器和 CI 环境 | 较重，特别是在执行大规模测试时资源占用高 |
| **报告方式**  | 简洁明了，适合快速查看测试结果           | 可视化丰富，支持图形化报告和日志分析    |
| **扩展性**    | 限制较多，主要集中于 HTTP 性能测试        | 支持多种协议，功能丰富，插件扩展多      |

### 1. 性能和资源占用

OHA 使用 Rust 编写，Rust 在内存管理和并发执行方面的优势使得 OHA 在高并发负载下表现尤为出色。与 Java 编写的 JMeter 相比，OHA 的资源占用显著较低，这意味着在小型服务器和 CI/CD 环境中，OHA 更加高效和节省资源。

JMeter 在进行大规模负载测试时，由于其基于 Java 的架构，可能会遇到内存占用高、启动慢等问题，尤其是在模拟几千、几万请求时。

### 2. 易用性和上手难度

OHA 是一个命令行工具，配置简洁，开发者只需要指定目标 URL、并发数和请求次数，就可以快速开始测试。这使得 OHA 非常适合那些需要快速进行性能测试的开发者，尤其是在 CI/CD 流程中集成自动化测试时，OHA 的简洁性是一个非常大的优势。

与此相比，JMeter 提供了图形化的用户界面，虽然功能强大，但对于初学者来说，图形界面的配置和选项较多，需要一定的学习曲线。JMeter 适合那些需要多协议、多维度测试的复杂场景，而 OHA 更适合快速、单一的 HTTP 性能测试。

### 3. 测试报告和结果

OHA 提供了简洁的命令行输出，测试结果包括吞吐量、响应时间、请求成功率等关键指标。这些数据帮助开发者快速评估 HTTP 服务的性能瓶颈。

JMeter 提供了更加丰富的可视化报告，支持图表、表格和日志文件分析，适合那些需要详细性能分析报告的用户。对于复杂的性能测试，JMeter 的图形化报告可以帮助分析各种请求的细节和响应情况。

### 4. 适用场景

- **OHA 适用场景**：如果你需要进行快速、简单的 HTTP 性能测试，OHA 是一个非常合适的选择。它可以帮助开发者在短时间内获得测试数据，适用于单一 HTTP 服务的性能评估、API 测试等。

- **JMeter 适用场景**：JMeter 更加适合那些复杂的性能测试需求，比如多协议（HTTP、JDBC、FTP 等）的负载测试、大规模压力测试、以及需要图形化报告的详细性能分析。对于企业级的性能测试，JMeter 提供了更多的定制化选项和功能。

## 如何使用 OHA？

### 安装 OHA

OHA 的安装非常简单，你可以通过以下方式进行安装：

- 使用 `cargo` 安装：

```bash
cargo install oha
```

- 直接使用 Homebrew（适用于 macOS 或 Linux）安装：

```bash
brew install oha
```

- 在 Windows 上还可以直接使用 winget 来安装：

```bash
winget install hatoo.oha
```

Ohayou 的使用非常灵活，它提供了丰富的命令行选项来满足不同的测试需求。以下是一些常用的操作和示例：

### 使用 OHA

```shell
oha https://example.com
```

这个命令会向 `https://example.com` 发送 200 个请求，使用 50 个并发连接。

#### 指定请求数量和并发数

```shell
oha -n 1000 -c 100 https://example.com
```

这个命令会向 `https://example.com` 发送 1000 个请求，使用 100 个并发连接。

#### 指定压测时间

```shell
oha -z 10s https://example.com
```

这个命令会持续向 `https://example.com` 发送请求 10 秒钟。

#### QPS 限速

```shell
oha -q 100 https://example.com
```

这个命令会以每秒 100 个请求的速度向 `https://example.com` 发送请求。

#### 动态 URL

```shell
oha --rand-regex-url http://127.0.0.1/[a-z][a-z][0-9]
```

这个命令会生成类似于 `http://127.0.0.1/ab1` 的随机 URL 进行压测。

#### 更真实的压测场景

```shell
oha -z 10s -c 100 -q 100 --latency-correction --disable-keepalive https://example.com
```

这个命令会进行 10 秒的压测，使用 100 个并发连接，每秒 100 个请求的速率，并启用延迟修正，关闭 keep-alive 连接。

#### 从文件中读取 URL

```shell
oha --urls-from-file urls.txt
```

这个命令会从 `urls.txt` 文件中读取 URL 列表进行压测，每个 URL 占一行。

#### 输出 JSON 格式结果

```shell
oha -j -n 1000 -c 100 https://example.com
```

这个命令会以 JSON 格式输出压测结果，方便后续分析。

#### 自定义 HTTP 方法、Header、Body

```shell
oha -m POST -H "Content-Type: application/json" -d '{"key":"value"}' https://example.com
```

这个命令会发送 POST 请求到 `https://example.com`，包含自定义的 Header 和 Body。

#### 设置超时时间

```shell
oha -t 5s https://example.com
```

这个命令会设置每个请求的超时时间为 5 秒。

## 性能比较：Ohayou vs JMeter

在性能测试领域，JMeter 是一个老牌的工具，而 Ohayou 作为一个新星，在基准测试中显示出了优异的性能。在与 rakyll/hey 的对比测试中，Ohayou 被发现运行速度约为 1.32 ± 0.48 倍于 rakyll/hey。这得益于其底层的 Rust 和 tokio 框架，使得 Ohayou 在处理高并发请求时更加高效。

JMeter 是 Apache 组织开发的基于 Java 的压力测试工具，具有开源免费、框架灵活、多平台支持等优势。JMeter 除了压力测试外，也可以应用于接口测试上。JMeter 的性能也非常出色，但是相比 Ohayou，它可能在某些场景下显得稍微笨重一些。

## 总结

Ohayou 以其轻量级和高性能的特点，在性能测试领域展现出了强大的竞争力。与 JMeter 相比，Ohayou 在某些场景下可能更加适合需要快速、轻量级负载测试的场合。当然，选择哪个工具还需要根据具体的测试需求和环境来决定。

Ohayou 作为一个新兴的 HTTP 负载生成器，以其轻量级、高性能和实时 TUI 界面，为性能测试提供了新的选择。无论你是初学者还是经验丰富的开发者，Ohayou 都是一个值得尝试的工具。当你使用上了它，可能你就会毫不犹豫的爱上它！

