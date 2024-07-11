---
title: Go语言中的map为什么默认不是并发安全的？
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
abbrlink: 3761f660
date: 2024-07-12 00:51:29
img:
coverImg:
password:
summary:
---

今天我们聊一个 Go 语言中的 “热门” 话题——为什么 **Go 语言中的 map 默认不是并发安全**的呢？

对于广大 Go 程序员来说，尤其是那些刚跨入 Go 世界的新朋友们，这个问题或许让你们摸不着头脑。别急，让我们一起慢慢揭开这层神秘的面纱。

## Go 语言中 map 的基本使用

首先，我们得知道 map 是什么。在 Go 中，map 是一种内置的数据结构，它提供了 “键值对”（Key-Value）的存储机制。使用 map，你可以通过 Key 快速找到对应的 Value，这让我们在处理一些需要快速查询的场景时如虎添翼。

一个简单的 map 示例：

```go
package main

import "fmt"

func main() {
    // 创建一个map
    myMap := make(map[string]int)
    // 向map中添加键值对
    myMap["apple"] = 1
    myMap["banana"] = 2

    // 从map中获取值
    fmt.Println(myMap["apple"]) // 输出: 1
}
```

## 那为什么 map 默认不是并发安全的呢？

难不成 Go 官方觉得太复杂了？性能太差了？还是为了什么？

### 典型使用场景

Go 官方认为，map 的典型使用场景并不需要从多个 goroutine 中安全地访问。因此，在设计时，优先考虑了性能和简单性，而没有将并发安全作为默认特性。这是一种基于使用案例进行权衡的结果。

### 性能考量

引入并发安全意味着每次操作 map 时都需要进行加锁和解锁，这无疑会增加额外的性能开销。为了大多数程序的性能考虑，Go 没有将 map 设计为并发安全的，因为这会导致即使在不需要并发访问的场景下，也要付出不必要的性能代价。

### 官方方案

从 `Go 1.6` 开始，引入了并发访问 map 的检测机制，如果检测到并发读写，程序会直接崩溃，而不是隐瞒问题。Go 官方倾向于让问题显露出来（"let it crash"），这样可以迫使开发者正视并发问题，采取正确的方法来解决。

## 如何安全地在多个 goroutine 中操作 map？

虽然原生的 map 不是并发安全的，但 Go 提供了其他机制来解决并发访问的问题。最直接的方法是使用互斥锁 `sync.Mutex`，来确保同一时间只有一个 goroutine 能访问 map。

> 当然现在不止这一个方法保证 map 并发安全，由于篇幅有限，这里仅以此为例。

例子如下：

```go
package main

import (
    "fmt"
    "sync"
)

var (
    myMap = make(map[string]int)
    lock  sync.Mutex
)

func main() {
    // 启动一个 goroutine 写入数据
    go func() {
        for {
            lock.Lock()
            myMap["apple"] = 1
            lock.Unlock()
        }
    }()
    
    // 在主 goroutine 中读取数据
    for {
        lock.Lock()
        fmt.Println(myMap["apple"])
        lock.Unlock()
    }
}
```

## 结论

通过以上的探讨，我们了解了为什么 Go 语言中的 map 默认不是并发安全的，其实就是一句话概括：**Go 官方觉得大部分场景都不需要支持并发，从性能上做的考虑**。

Go 语言的设计哲学之一就是简单而有效，通过让开发者显式地处理并发问题，既保证了性能，也让代码的行为更加透明。

也有网友讨论说，可以像 Java 那样提供两个 map，一个支持并发，性能差些，一个不支持并发，性能好。但是 Go 官方为什么不提供两个，那就不得而知了，可能是为了符合 Go 语言“少就是多”的理念？

你有什么看法呢？一起聊聊……