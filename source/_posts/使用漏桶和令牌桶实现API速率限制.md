---
title: 使用漏桶和令牌桶实现API速率限制
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
  - 漏桶
  - 令牌桶
  - 限流
abbrlink: ea016854
date: 2024-11-19 15:06:58
img:
coverImg:
password:
summary:
---

在现代 Web 应用程序中，流量的突增是不可避免的。为防止服务器被过多的请求压垮，**限流（Rate Limiting）** 是一个至关重要的技术手段。

本文将通过 Go 语言的 Gin 框架，演示如何使用**漏桶算法**和**令牌桶算法**来实现 API 的限流。

## 限流的意义

限流的主要目的是保护系统资源，防止因请求量过大导致服务器崩溃。同时，它也能防止恶意用户对系统的攻击，确保服务的稳定性和可用性。

## 两种常见的限流算法

1. **漏桶算法（Leaky Bucket）**

漏桶算法将请求视为水滴，水滴先进入桶中，然后以固定的速率从桶中流出。如果请求的速率超过了桶的流出速率，多余的请求将会被丢弃。

这个算法的优点很明显，就是让请求非常稳定，但是缺点也很明显，因为请求非常稳定，就不适于一些秒杀等一些可能在某一段时间会有洪峰流量的场景。不太好适情况控制流量的进入。

2. **令牌桶算法（Token Bucket）**

令牌桶算法中，系统会以固定的速率向桶中加入令牌，每个请求需要获取一个令牌才能执行。如果桶中没有足够的令牌，请求将被拒绝。

## 代码实现

在这个示例中，我们将展示如何在 Gin 框架中应用这两种算法来实现 API 的限流。

```go
package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	ratelimit2 "github.com/juju/ratelimit" // 令牌桶算法
	ratelimit1 "go.uber.org/ratelimit"     // 漏桶算法
)

func pingHandler(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "pong",
	})
}

func pingHandler2(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "pong2",
	})
}

// rateLimit1 使用漏桶算法来限制请求速率
func rateLimit1() func(ctx *gin.Context) {
	// 漏桶算法，第一个参数为两滴水滴之间的时间间隔。
	// 此时表示两滴水之间的时间间隔是 100 纳秒
	rl := ratelimit1.New(100)

	return func(ctx *gin.Context) {
		// 尝试取出水滴
		if waitTime := rl.Take().Sub(time.Now()); waitTime > 0 {
			fmt.Printf("需要等待 %v 秒，下一滴水才会滴下来\n", waitTime)
			// 这里我们可以让程序继续等待，也可以直接拒绝掉
			// time.Sleep(waitTime)
			ctx.String(http.StatusOK, "rate limit, try again later")
			ctx.Abort()
			return
		}
		// 证明可以继续执行
		ctx.Next()
	}
}

// rateLimit2 使用令牌桶算法来限制请求速率
func rateLimit2() func(ctx *gin.Context) {
	// 令牌桶算法：第一个参数为每秒填充令牌的速率为多少
	// 第二个参数为令牌桶的容量
	// 这里表示每秒填充 10 个令牌
	rl := ratelimit2.NewBucket(time.Second, 10)

	return func(ctx *gin.Context) {
		// 尝试取出令牌
		var num int64 = 1
		// 这里表示需要 num 个令牌和已经取出的令牌数是否相等
		// 不相等，则表示超过了限流
                // 比如，假设每一个请求过来消耗2个令牌，但是从桶中取出的令牌个数为 1 ，那么则认为超过了限流（一般而言是一个请求消耗一个令牌，这里仅为举例）
		if rl.TakeAvailable(num) != num {
			// 此次没有取到令牌，说明超过了限流
			ctx.String(http.StatusOK, "rate limit, try again later")
			ctx.Abort()
			return
		}
		// 证明可以继续执行
		ctx.Next()
	}
}

func main() {
	r := gin.Default()

	// 漏桶算法限流
	r.GET("/ping", rateLimit1(), pingHandler)

	// 令牌桶算法限流
	r.GET("/ping2", rateLimit2(), pingHandler2)

	r.Run()
}
```

## 代码解析

1. **漏桶算法的实现（`rateLimit1` 函数）**
    - 通过 `go.uber.org/ratelimit` 包中的 `ratelimit.New` 方法创建了一个限流器。
    - 当请求速率超过限流器的处理能力时，请求将被拒绝，并返回 "rate limit, try again later"。

2. **令牌桶算法的实现（`rateLimit2` 函数）**
    - 使用 `github.com/juju/ratelimit` 包实现了令牌桶算法。每秒填充一定数量的令牌到桶中。
    - 如果桶中没有足够的令牌，请求将被拒绝。

3. **Gin 路由配置**
    - 在 `main` 函数中，通过 `rateLimit1` 和 `rateLimit2` 中间件为 `/ping` 和 `/ping2` 路由分别设置了漏桶和令牌桶限流。

## 总结

在本文中，我们演示了如何在 Go 中使用漏桶算法和令牌桶算法实现 API 的限流。

这些算法在高并发的 Web 服务中非常有用，可以有效防止服务被大量请求淹没，确保系统的稳定性。希望通过这篇文章，您能更好地理解并应用这些限流技术到您的项目中。
