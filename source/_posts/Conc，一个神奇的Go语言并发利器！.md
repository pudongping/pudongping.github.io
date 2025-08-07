---
title: Conc，一个神奇的Go语言并发利器！
author: Alex
top: false
hide: false
cover: false
toc: true
mathjax: false
categories: Go
tags:
  - Go
  - Golang
  - Conc
abbrlink: '1e620500'
date: 2025-08-07 10:11:13
img:
coverImg:
password:
summary:
---

在 Go 语言中，goroutine 和通道是并发编程的核心，但随着项目复杂度的增加，管理成百上千个 goroutine 并确保它们正确协作变得极具挑战性。`sourcegraph/conc` 扩展包为我们提供了一套结构化并发工具，帮助我们更安全、高效地编写并发代码。

![](https://upload-images.jianshu.io/upload_images/14623749-2223286feb708123.png?imageMogr2/auto-orient/strip%7CimageView2/2/w/1240)


## 一、初识 conc 包

### 并发编程的常见问题

在传统 Go 并发编程中，我们常常面临以下问题：

- goroutine 泄漏：忘记等待子 goroutine 完成，导致程序退出时遗留大量未完成的 goroutine。
- panic 处理困难：子 goroutine 中的 panic 会导致整个程序崩溃，缺乏有效的恢复机制。
- 并发代码可读性差：大量的同步原语（如 sync.WaitGroup）和通道操作让代码变得复杂难懂。

### conc 包

`conc` 是一个为 Go 语言量身打造的结构化并发（structured concurrency）库，目标是让并发代码更易读、更安全、更少样板：

1. 提供结构化的并发原语，确保所有 goroutine 都有明确的所有者，并在适当的时候被清理。
2. 自动处理 panic，避免程序因未捕获的异常而崩溃。
3. 提供高层次的抽象，简化并发代码的编写，提高代码可读性。

## 二、conc.WaitGroup：升级版的并发等待

### 传统 sync.WaitGroup 的局限性

在标准库中，sync.WaitGroup 是管理并发 goroutine 的常用工具，但它存在以下问题：

- 忘记调用 Add() 或 Done() 会导致逻辑错误。
- 无法优雅地处理 goroutine 中的 panic。
- 没有明确的资源拥有者概念，容易导致 goroutine 泄漏。

### conc.WaitGroup 的改进

conc.WaitGroup 在传统 WaitGroup 的基础上进行了改进，提供了更安全的并发管理方式。

**代码示例：**

```go
package main

import (
    "fmt"
    "sync/atomic"

    "github.com/sourcegraph/conc"
)

func main() {
    var count atomic.Int64
    var wg conc.WaitGroup
    // 启动 10 个并发任务
    for i := 0; i < 10; i++ {
        wg.Go(func() {
            // 对 count 原子加 1
            count.Add(1)
        })
    }
    // 等待所有任务完成
    wg.Wait()
    fmt.Println("任务完成，总计:", count.Load()) // 输出 10
}
```

**注释说明：**

- `wg.Go(func())`：启动一个新 goroutine 并自动管理其生命周期。
- `wg.Wait()`：等待所有通过 Go() 启动的 goroutine 完成。

若需在等待时收集 panic 信息而不立即抛出，可使用 `WaitAndRecover()`：
```go
wg.Go(func() { panic("错误演示") })
recovery := wg.WaitAndRecover()
fmt.Println("捕获 panic:", recovery.Value) // 输出 "错误演示"
```

## 三、goroutine 池：pool.Pool 的使用

### 为什么需要 goroutine 池？

在高并发场景下，无限制地启动 goroutine 可能导致系统资源耗尽。goroutine 池可以限制并发任务的数量，确保系统稳定运行。

### pool.Pool 的基本用法

**代码示例：**

```go
package main

import (
    "fmt"
    "github.com/sourcegraph/conc/pool"
)

func main() {
    // 创建一个限制最大 goroutine 数量为 2 的池
    p := pool.New().WithMaxGoroutines(2)

    // 提交 5 个任务到池中
    for i := 0; i < 5; i++ {
        p.Go(func() {
            fmt.Println("任务执行中...")
        })
    }

    // 等待所有任务完成
    p.Wait()
    fmt.Println("所有任务已完成")
}
```

**注释说明：**

- `pool.New()`：创建一个新的 goroutine 池。
- `WithMaxGoroutines(2)`：设置池中最大 goroutine 数量为 2。
- `p.Go(func())`：将任务提交到池中执行。

## 四、带上下文的 goroutine 池：pool.ContextPool

### 什么是上下文？

上下文（context）用于在并发任务之间传递截止时间、取消信号等信息。pool.ContextPool 允许我们在任务中使用上下文，方便控制任务的执行。

**代码示例：**

```go
package main

import (
    "context"
    "fmt"
    "github.com/sourcegraph/conc/pool"
    "time"
)

func main() {
    // 创建一个带上下文的 goroutine 池
    ctx, cancel := context.WithCancel(context.Background())
    defer cancel()

    p := pool.NewWithContext(ctx).WithMaxGoroutines(2)

    // 提交 5 个任务到池中
    for i := 0; i < 5; i++ {
        p.Go(func() {
            fmt.Println("任务执行中...")
            // 模拟长时间运行的任务
            time.Sleep(2 * time.Second)
        })
    }

    // 等待 3 秒后取消所有任务
    time.AfterFunc(3*time.Second, func() {
        cancel()
        fmt.Println("取消所有任务")
    })

    // 等待所有任务完成或被取消
    p.Wait()
    fmt.Println("所有任务已完成")
}
```

**注释说明：**

- `pool.NewWithContext(ctx)`：使用给定的上下文创建 goroutine 池。
- `cancel()`：取消上下文，导致池中的任务停止执行。

## 五、带错误处理的 goroutine 池：pool.WithErrors

### 为什么需要错误处理？

在并发任务中，任务可能出现错误。pool.WithErrors 允许我们捕获任务中的错误并进行处理。

**代码示例：**

```go
package main

import (
    "fmt"
    "github.com/sourcegraph/conc/pool"
)

func main() {
    // 创建一个带错误处理的 goroutine 池
    p := pool.New().WithErrors()

    // 提交 3 个任务到池中
    p.Go(func() error {
        return fmt.Errorf("任务 1 出错")
    })

    p.Go(func() error {
        return nil // 任务 2 成功
    })

    p.Go(func() error {
        return fmt.Errorf("任务 3 出错")
    })

    // 等待所有任务完成并获取错误信息
    if err := p.Wait(); err != nil {
        fmt.Printf("任务执行出错: %v\n", err)
    }
}
```

**注释说明：**

- `pool.New().WithErrors()`：创建一个带错误处理的 goroutine 池。
- `p.Go(func() error)`：提交返回错误的任务到池中。
- `p.Wait()`：等待所有任务完成并返回错误信息。


## 总结：conc 包的适用场景与优势

### 适用场景

- 需要管理大量并发任务的场景（如 Web 服务器、爬虫等）。
- 对并发任务的错误处理和资源清理有严格要求的场景。
- 需要提高并发代码可读性和维护性的场景。

### 核心优势

- 提供结构化的并发管理，避免 goroutine 泄漏。
- 自动捕获和处理 panic，提高程序的健壮性。
- 提供高层次的抽象，简化并发代码的编写。

通过本文的介绍和示例，相信你已经对 `sourcegraph/conc` 扩展包有了深入的了解。它为 Go 并发编程提供了一套强大的工具，帮助我们更安全、高效地编写并发代码。在实际项目中，根据需求灵活选择合适的工具，可以显著提升开发效率和代码质量。