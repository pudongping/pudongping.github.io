---
title: 如何在 Go 中获取 goroutine 的 ID？
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
abbrlink: 2dcc2ce
date: 2025-06-12 10:31:29
img:
coverImg:
password:
summary:
---

在使用 Go 语言进行并发编程时，Goroutine 是一种轻量级线程，具有很高的性能优势。然而，Go 语言**并未直接提供**获取 Goroutine ID 的官方 API。这是 Go 语言设计的一部分，目的是避免开发者依赖 Goroutine ID 进行不必要的复杂操作。然而，在某些场景下，获取 Goroutine ID 可能会有助于调试和日志跟踪。

本文将详细介绍在 Go 语言中获取 Goroutine ID 的几种方法。

---

## 为什么需要 Goroutine ID？

在调试并发程序时，了解某段代码是由哪个 Goroutine 执行的，有助于：
- **调试问题**：清楚地知道问题来源。
- **日志追踪**：区分不同 Goroutine 的执行过程。
- **性能分析**：理解并发任务的执行情况。

虽然这些需求合理，但 Go 语言希望开发者更专注于 Goroutine 的逻辑而非它的标识。因此，官方并未直接提供获取 Goroutine ID 的功能。

---

## 获取 Goroutine ID 的实现原理

其实 Go 的每个 Goroutine 都有一个唯一的标识符，存储在其运行时的内部结构中。这个 ID 不直接对外暴露，但我们可以通过间接手段获取。

Go 的运行时包 `runtime` 提供了一些工具来帮助我们了解 Goroutine 的状态，其中最常用的是 `runtime.Stack`。

`runtime.Stack` 可以生成当前 Goroutine 的调用栈信息，这些信息中包含了 Goroutine 的 ID。通过解析调用栈的内容，就能提取出 Goroutine 的 ID。

---

## 获取 Goroutine ID

以下是一个获取当前 Goroutine ID 的简单实现：

```go
package main

import (
	"bytes"
	"fmt"
	"runtime"
	"strconv"
)

// GetGoroutineID 返回当前 Goroutine 的 ID
// 通过 runtime.Stack 获取当前 Goroutine 的栈信息，然后提取出 Goroutine ID
// 这种方式可以获取到当前 Goroutine 的 ID，但是性能较差
func GetGoroutineID() uint64 {
	var buf [64]byte
	// runtime.Stack(buf[:], false) 会将当前 Goroutine 的栈信息写入 buf 中
	// 第二个参数是 false 表示只获取当前 Goroutine 的栈信息，如果为 true 则会获取所有 Goroutine 的栈信息
	n := runtime.Stack(buf[:], false)
	stack := string(buf[:n])
	// fmt.Println("========")
	// fmt.Println(stack)
	// fmt.Println()
	// stack 样例: "goroutine 7 [running]:\n..."
	// 提取 goroutine 后面的数字
	fields := bytes.Fields([]byte(stack))
	id, err := strconv.ParseUint(string(fields[1]), 10, 64)
	if err != nil {
		panic(fmt.Sprintf("无法解析 Goroutine ID: %v", err))
	}
	return id
}

func main() {
	fmt.Printf("Main Goroutine ID: %d\n", GetGoroutineID())

	var wg sync.WaitGroup
	for i := 0; i < 10; i++ {
		i := i
		wg.Add(1)
		go func() {
			defer wg.Done()
			fmt.Printf("Child [%d] Goroutine ID: [%d]\n", i, GetGoroutineID())
		}()
	}

	wg.Wait()
}
```

---

### 代码解析

1. **`runtime.Stack` 获取调用栈**  
   `runtime.Stack` 会返回当前 Goroutine 的调用栈信息，包括 Goroutine 的 ID。

2. **解析 Goroutine ID**  
   调用栈信息是字符串形式，例如：
   ```
   goroutine 7 [running]:
   ```
   我们只需提取 `goroutine` 后面的数字，即可获取 ID。

3. **类型转换**  
   使用 `strconv.ParseUint` 将字符串 ID 转换为数值类型，便于后续操作。

4. **主 Goroutine 与子 Goroutine 的对比**  
   在 `main` 函数中，我们分别打印主 Goroutine 和子 Goroutine 的 ID，以观察它们的不同。

---

## 注意事项

1. **性能影响**  
   使用 `runtime.Stack` 获取 Goroutine ID 的代价相对较高，仅适用于调试或日志场景，不建议在性能敏感的代码中频繁使用。

2. **不依赖 ID 进行业务逻辑**  
   Goroutine ID 是一个内部实现细节，不应在业务逻辑中依赖它，例如用于锁定资源或同步任务。Go 鼓励使用通道（channel）等高级并发原语来管理任务。

既然使用 `runtime.Stack` 先获取堆栈信息的方式获取 Goroutine ID 性能不高，那么有没有更加高效的方式呢？

## 使用第三方包高效获取

我们可以采用第三方包 `github.com/petermattis/goid` 来高效的获取当前 goroutine 的 ID

首先我们先安装这个包

```bash
go get -u github.com/petermattis/goid
```

这个包使用起来也非常简单，直接

```go
// goid 库使用了 C 和 汇编来获取 Goroutine ID，性能更好
func GetGoroutineID1() int64 {
	id := goid.Get()
	return id
}
```

`goid` 库使用了 C 和汇编来获取 goroutine ID，所以性能更好。并且 `goid` 对多个 go 版本做了兼容，而且为了保证兼容性，我们通过查看 `https://github.com/petermattis/goid/blob/master/goid.go` 也可以发现提供了一个 Go 语言版本的实现。这个 Go 版本的实现也是通过使用 `runtime.Stack()` 来实现的。所以，如果你真的需要获取 goroutine ID，那么还是比较推荐使用这个包的。

---

## 总结

Goroutine 是 Go 并发编程的核心，而 Goroutine ID 在某些场景下可以帮助我们更好地理解和调试代码。尽管 Go 官方没有提供直接的 API，但通过 `runtime.Stack`，我们可以间接获取到 Goroutine 的 ID。但是由于通过 `runtime.Stack` 的方式去获取 Goroutine ID 性能不高，因此如果你确确实实想要获取 Goroutine ID 时，就建议你直接使用 `goid` 包来获取。

然而，获取 ID 应仅限于调试场景，在实际开发中更应关注 Goroutine 的行为和通道通信。

希望这篇文章能帮助您在编程过程中更好地掌握 Goroutine 的使用！如果有任何疑问或其他主题想了解，欢迎留言讨论 😊！