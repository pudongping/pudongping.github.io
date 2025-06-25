---
title: Go 语言中如何处理并发错误
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
abbrlink: d082e483
date: 2025-06-25 10:14:34
img:
coverImg:
password:
summary:
---

在 Go 语言中，错误处理一直是开发中不可或缺的一部分。尤其在并发编程场景下，由于多个 goroutine 同时运行，错误的传递和处理就变得更为复杂。这篇文章就介绍了一些常见的处理并发错误的方法，以供各位参考。

![](https://upload-images.jianshu.io/upload_images/14623749-ce228dd72dcf74d6.png?imageMogr2/auto-orient/strip%7CimageView2/2/w/1240)


## 一、 panic 只会触发当前 goroutine 中的 defer 操作

很多开发者初次接触 Go 时容易误解 panic 的作用范围。下面我们先来看一个错误的代码示例：

### 1.1 示例代码

```go
package main

import (
	"fmt"
	"time"
)

func main() {
	// 在主 goroutine 中设置 defer，用于捕获 panic
	// 注意：这个 defer 只能捕获发生在主 goroutine 中的 panic
	defer func() {
		// recover() 只能捕获当前 goroutine 内的 panic，
		// 如果 panic 发生在其他 goroutine 中，该 defer 无法捕获
		if e := recover(); e != nil {
			fmt.Println("捕获到 panic：", e)
		}
	}()

	// 启动子 goroutine，演示 panic 的传播范围
	go func() {
		// 输出提示信息，表示子 goroutine 开始执行
		fmt.Println("子 goroutine 开始")
		// 主动触发 panic，注意这里的 panic 发生在子 goroutine 内，
		// 因此主 goroutine 中的 defer 无法捕获该 panic
		panic("Goroutine 发生 panic")
	}()

	// 主 goroutine 等待一段时间，确保子 goroutine 有足够时间执行
	time.Sleep(2 * time.Second)
	// 输出主 goroutine 结束信息
	fmt.Println("主 goroutine 结束")
}

```

运行这段代码，我们会发现，会直接报错了：

```bash
子 goroutine 开始
panic: Goroutine 发生 panic

goroutine 18 [running]:
main.main.func2()
        ~/golang-tutorial/tt.go:25 +0x59
created by main.main in goroutine 1
        ~/golang-tutorial/tt.go:20 +0x3b
exit status 2
```

### 1.2 代码说明

*   **主 goroutine 中的 defer：**
    主函数开始时设置了一个 defer 函数，目的是在发生 panic 时捕获并打印错误信息。然而，由于 recover 只能捕获当前 goroutine 内的 panic，当子 goroutine 内发生 panic 时，这个 defer 不会生效。

*   **子 goroutine 中的 panic：**
    在子 goroutine 中调用 panic 后，由于没有设置独立的 recover 逻辑，该 goroutine 会直接崩溃，panic 信息不会传递到主 goroutine 中。
    这样可以清楚地看到，即使主 goroutine 使用了 defer 进行错误捕获，也无法捕捉到其他 goroutine 中发生的 panic。

*   **延时等待：**
    主 goroutine 使用 `time.Sleep` 等待一定时间，以确保子 goroutine 有机会执行并触发 panic，从而验证 panic 的作用范围。


既然程序会直接崩溃，那么，如何解决这个问题呢？

### 1.3 正确处理

我们只需要在子 goroutine 中使用 recover 就可以了：

```go
package main

import (
	"fmt"
	"time"
)

func main() {
	defer func() {
		if e := recover(); e != nil {
			fmt.Println("捕获到 panic：", e)
		}
	}()

	go func() {
		defer func() {
			if e := recover(); e != nil {
				fmt.Println("子 goroutine 捕获到 panic：", e)
			}
		}()
		fmt.Println("子 goroutine 开始")
		panic("Goroutine 发生 panic")
	}()

	time.Sleep(2 * time.Second)
	fmt.Println("主 goroutine 结束")
}
```

运行以上代码，可以发现，打印出的结果为：

```bash
子 goroutine 开始
子 goroutine 捕获到 panic： Goroutine 发生 panic
主 goroutine 结束
```

这就说明：**panic 只会触发当前 goroutine 内的 defer 操作，不能跨 goroutine 捕获或恢复其他 goroutine 中的 panic。**

* * *

## 二、多 goroutine 中收集错误和结果

假设我们有个需求，需要同时使用多个 goroutine 通过 `http.Get` 去请求以下四个地址，其中只有 `https://httpbin.org/get` 能够正常响应，其余地址均为故意写错的地址：

*   `https://httpbin1.org/get`
*   `https://httpbin.org/get`
*   `https://httpbin2.org/get`
*   `https://httpbin3.org/get`

### 2.1 如何批量收集错误信息？

在并发请求中，可以通过错误通道（ error channel ）来收集各个 goroutine 中发生的错误。例如：

```go
package main

import (
	"fmt"
	"net/http"
	"sync"
)

func main() {
	urls := []string{
		"https://httpbin1.org/get",
		"https://httpbin.org/get",
		"https://httpbin2.org/get",
		"https://httpbin3.org/get",
	}

	var wg sync.WaitGroup
	// 创建一个带缓冲的错误通道，大小为 URL 数量
	errCh := make(chan error, len(urls))

	// 遍历所有 URL，分别启动 goroutine 发起请求
	for _, url := range urls {
		wg.Add(1)
		go func(url string) {
			defer wg.Done() // 保证 goroutine 结束时减少计数
			resp, err := http.Get(url)
			if err != nil {
				// 如果请求出错，将错误发送到错误通道中
				errCh <- fmt.Errorf("请求 %s 失败： %v", url, err)
				return
			}
			defer resp.Body.Close()
			// 打印成功信息
			fmt.Printf("请求 %s 成功，状态码： %d\n", url, resp.StatusCode)
		}(url)
	}

	// 等待所有 goroutine 执行完毕
	wg.Wait()
	// 关闭错误通道
	close(errCh)

	// 遍历错误通道，输出所有错误信息
	for err := range errCh {
		fmt.Println("错误信息：", err)
	}
}
```

在这个示例中，我们通过一个 channel `errCh` 来存储每个 goroutine 产生的错误，待所有 goroutine 执行完毕后，再统一处理错误信息。

### 2.2 那如果也需要结果呢？

如果希望每个请求的结果和可能的错误信息，我们可以定义一个结构体，将请求的结果与错误信息封装在一起，再通过 channel 收集：

```go
package main

import (
	"fmt"
	"io"
	"net/http"
	"sync"
)

// Result 用于封装每个请求的结果和错误信息
type Result struct {
	URL        string // 请求的 URL
	StatusCode int    // 返回的 HTTP 状态码
	Err        error  // 请求过程中发生的错误
	Content    []byte // 返回的内容
}

func main() {
	urls := []string{
		"https://httpbin1.org/get",
		"https://httpbin.org/get",
		"https://httpbin2.org/get",
		"https://httpbin3.org/get",
	}

	var wg sync.WaitGroup
	// 创建带缓冲的结果通道，大小为 URL 数量
	resCh := make(chan Result, len(urls))

	// 遍历 URL，启动 goroutine 进行请求
	for _, url := range urls {
		wg.Add(1)
		go func(url string) {
			defer wg.Done()
			resp, err := http.Get(url)

			result := Result{URL: url}

			if err != nil {
				// 将错误结果封装后发送到结果通道
				result.Err = err
			} else {
				defer resp.Body.Close()
				body, _ := io.ReadAll(resp.Body)
				// 将成功的结果封装后发送到结果通道
				result.StatusCode = resp.StatusCode
				result.Content = body
			}

			resCh <- result
		}(url)
	}

	// 等待所有 goroutine 执行完毕
	wg.Wait()
	close(resCh)

	// 遍历结果通道，输出每个请求的结果和错误信息
	for res := range resCh {
		if res.Err != nil {
			fmt.Printf("请求 %s 失败： %v\n", res.URL, res.Err)
		} else {
			fmt.Printf("请求 %s 成功，状态码： %d, 内容： %s \n", res.URL, res.StatusCode, string(res.Content))
		}
	}
}
```

在这个示例中，每个 goroutine 都会将自己的请求结果封装到 `Result` 结构体中，通过通道传递回来，最后我们可以一一对应地输出结果和错误信息。

* * *

## 三、 errgroup 包

### 3.1 errgroup 包简介

`golang.org/x/sync/errgroup` 包提供了一个便捷的方式来管理一组 goroutine，并能统一收集它们产生的错误。该包的主要功能有：

*   **错误收集与聚合：** 当多个 goroutine 发生错误时，errgroup 会返回**第一个**遇到的错误。
*   **自动等待：** 调用 `g.Wait()` 可以等待所有启动的 goroutine 执行完毕。
*   **与 context 结合：** 通过 `WithContext` 方法，可以为所有 goroutine 传入相同的 context，从而实现统一的取消逻辑。

这些特性使得 errgroup 在需要并发执行多个任务且统一管理错误时非常有用。

### 3.2 用 errgroup 包实战一下

以下示例演示了如何使用 errgroup 包来并发请求多个 URL：

```go
package main

import (
	"fmt"
	"net/http"

	"golang.org/x/sync/errgroup"
)

func main() {
	urls := []string{
		"https://httpbin1.org/get",
		"https://httpbin.org/get",
		"https://httpbin2.org/get",
		"https://httpbin3.org/get",
	}

	// 定义一个存储结果的切片，与 errgroup 共同使用
	results := make([]string, len(urls))
	var g errgroup.Group

	// 遍历所有 URL，启动 goroutine 执行 HTTP 请求
	for i, url := range urls {
		i, url := i, url // 为了避免闭包引用同一个变量
		g.Go(func() error {
			fmt.Println("开始请求：", url)
			resp, err := http.Get(url)
			if err != nil {
				return fmt.Errorf("请求 %s 失败： %v", url, err)
			}
			defer resp.Body.Close()
			results[i] = fmt.Sprintf("请求 %s 成功，状态码： %d", url, resp.StatusCode)
			return nil
		})
	}

	// 等待所有 goroutine 执行完毕
	if err := g.Wait(); err != nil {
		fmt.Println("发生错误：", err)
	}

	// 输出所有请求成功的结果
	for _, res := range results {
		fmt.Println(res)
	}
}
```

通过运行上面的代码，可能会打印出类似以下内容：

```bash
开始请求： https://httpbin3.org/get
开始请求： https://httpbin2.org/get
开始请求： https://httpbin1.org/get
开始请求： https://httpbin.org/get
发生错误： 请求 https://httpbin3.org/get 失败： Get "https://httpbin3.org/get": dial tcp: lookup httpbin3.org: no such host

请求 https://httpbin.org/get 成功，状态码： 200
```

我们可以得出以下重要的结论：**Wait 会阻塞直至由上述 Go 方法调用的所有函数都返回，但是，如果有错误的话，只会记录第一个非 nil 的错误，也就是说，如果有多个错误的情况下，不会收集所有的错误。**

并且，通过源码得知：**当遇到第一个错误时，如果之前设定了 cancel 方法，那么还会调用 cancel 方法**，那么，如何创建带有 cancel 方法的 errgroup.Group 呢？

### 3.3 使用 errgroup 包中的 WithContext 方法

有时我们希望在某个 goroutine 发生错误时，能够通知其他正在执行的任务提前取消。这时可以使用 `errgroup.WithContext` 方法。以下示例展示了如何实现这一点：

```go
package main

import (
	"context"
	"fmt"
	"net/http"

	"golang.org/x/sync/errgroup"
)

func main() {
	urls := []string{
		"https://httpbin1.org/get",
		"https://httpbin.org/get",
		"https://httpbin2.org/get",
		"https://httpbin3.org/get",
	}

	// 使用 context.Background 创建基本上下文，并通过 WithContext 包装 errgroup
	ctx := context.Background()
	g, ctx := errgroup.WithContext(ctx)

	// 定义存储结果的切片
	results := make([]string, len(urls))

	// 遍历所有 URL，启动 goroutine 发起请求
	for i, url := range urls {
		i, url := i, url
		g.Go(func() error {
			fmt.Println("开始请求：", url)
			// 在发起请求前，根据 context 判断是否取消
			req, err := http.NewRequestWithContext(ctx, "GET", url, nil)
			if err != nil {
				return err
			}

			resp, err := http.DefaultClient.Do(req)
			if err != nil {
				return fmt.Errorf("请求 %s 失败： %v", url, err)
			}
			defer resp.Body.Close()

			results[i] = fmt.Sprintf("请求 %s 成功，状态码： %d", url, resp.StatusCode)
			return nil
		})
	}

	// 如果有任一任务返回错误，将自动取消所有依赖于 ctx 的请求
	if err := g.Wait(); err != nil {
		fmt.Println("错误发生：", err)
	}

	for _, res := range results {
		fmt.Println(res)
	}
}
```

运行以上的代码，打印结果如下：

```bash
开始请求： https://httpbin3.org/get
开始请求： https://httpbin.org/get
开始请求： https://httpbin2.org/get
开始请求： https://httpbin1.org/get
错误发生： 请求 https://httpbin1.org/get 失败： Get "https://httpbin1.org/get": dial tcp: lookup httpbin1.org: no such host
```

在这个示例中，我们使用 `errgroup.WithContext` 创建了一个共享的上下文 `ctx`，所有的 HTTP 请求都与此 context 绑定。**一旦某个请求发生错误并返回，其他 goroutine 中绑定该 context 的请求会立即收到取消信号，从而实现整体任务的协同取消。**

* * *

## 四、总结

本文从以下几个方面详细介绍了在 Go 语言中如何处理并发错误：

*   **panic 和 defer：** 通过示例说明 panic 只会触发当前 goroutine 内的 defer 操作，并展示了即使主 goroutine 设置了 defer，也无法捕获子 goroutine 内的 panic。
*   **并发中错误收集：** 通过简单示例展示了如何在多个 goroutine 中分别收集错误信息，以及如何关联请求结果与错误信息。
*   **errgroup 包的使用：** 介绍了 errgroup 包的核心功能，展示了如何用 errgroup 包简化并发错误处理，同时详细演示了 WithContext 方法的使用场景和效果。

通过这些示例和详细解释，希望大家在实际开发中能够更加自信地处理并发任务中的错误问题，从而编写出更加健壮和易维护的代码。

希望这篇文章能对你理解 Go 语言中的并发错误处理有所帮助！
