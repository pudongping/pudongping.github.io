---
title: Go Web服务中如何优雅平滑重启？
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
abbrlink: cf5ac086
date: 2024-11-15 10:30:01
img:
coverImg:
password:
summary:
---

在生产环境中，当我们需要对正在运行的服务进行升级时，如何确保不影响当前未处理完的请求，同时又能应用新的代码，是个极具挑战性的问题。

传统的做法通常是停止当前服务，部署新代码后再重启服务，但这种方式会导致正在处理的请求被强制中断，用户体验会受到很大的影响。

在这篇文章中，我将带大家一起探索如何在 Go 语言中通过使用 `endless` 包来实现服务的**优雅重启**，即在不影响当前正在处理的请求的情况下，完成服务的无缝升级。

## 什么是优雅重启？

优雅重启的核心思想是：在服务启动新的进程处理新请求的同时，允许旧的进程继续完成其手头未完成的工作，然后再优雅地退出。这种方式可以确保服务在升级的过程中不会出现中断，提升用户体验的同时，也降低了在服务切换过程中的风险。

## 实现优雅重启的代码示例

下面的代码演示了如何使用 `endless` 包来实现 Gin 服务的优雅重启。

```go
package main

import (
	"log"
	"net/http"
	"time"

	"github.com/fvbock/endless"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.GET("/ping", func(c *gin.Context) {
		// 模拟程序处理请求需要 5 秒
		time.Sleep(5 * time.Second)
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	// 默认 endless 服务器会监听下列信号：
	// syscall.SIGHUP，syscall.SIGUSR1，syscall.SIGUSR2，syscall.SIGINT，syscall.SIGTERM 和 syscall.SIGTSTP
	// 接收到 SIGHUP 信号将触发 `fork/restart` 实现优雅重启（kill -1 pid 会发送 SIGHUP 信号）
	// 接收到 syscall.SIGINT 或 syscall.SIGTERM 信号将触发优雅关机
	// 接收到 SIGUSR2 信号将触发 HammerTime
	// SIGUSR1 和 SIGTSTP 被用来触发一些用户自定义的 hook 函数
	if err := endless.ListenAndServe(":8080", router); err != nil {
		log.Printf("Server err: %v\n", err)
	}

	log.Println("Server exiting")
}
```

## 如何验证优雅重启？

为了验证这个实现，我们可以通过以下步骤进行测试：

1. 打开终端，执行 `go build -o gin_graceful_restart gin_graceful_restart.go` 编译程序，然后执行 `./gin_graceful_restart` 启动服务。终端会输出当前服务的 PID（例如：`[pid] 12345`）。

2. 修改代码中的 `/ping` 接口的响应内容，比如将 `pong` 修改为 `pong1`。

3. 再次执行 `go build -o gin_graceful_restart gin_graceful_restart.go` 编译程序。

4. 在浏览器中访问 `http://127.0.0.1:8080/ping`，此时浏览器会等待服务返回响应（由于接口模拟了 5 秒的延迟）。

5. 在另一个终端中执行 `kill -1 12345` 命令，向服务发送 `syscall.SIGHUP` 信号，12345 为第一步中的 PID。

6. 依旧在第 4 步的浏览器中等待，等响应到 `pong` 信息之后，再次刷新页面，你会发现响应内容变成了 `pong1`，这意味着服务已经应用了新的代码，同时之前的请求也得到了正确的处理。

## Endless 的工作原理

`endless` 包实现优雅重启的原理非常简单：它会 fork 一个新的子进程来处理新的请求，而旧的进程则继续处理已经接收的请求。当旧的进程处理完所有的请求后才会退出。因此，虽然 PID 发生了变化，但服务依旧保持了无缝衔接。

值得注意的是，由于 `endless` 的这种机制，当你的项目是通过类似 `supervisor` 的软件来管理进程时，这种方式就不再适用了。因为 `supervisor` 会根据 PID 来管理进程，而在优雅重启过程中 PID 是会变化的，这会导致 `supervisor` 认为服务已经崩溃。

## 总结

在实际的生产环境中，优雅重启是非常实用的一项技术，它可以帮助我们在不影响用户体验的前提下，对服务进行升级和维护。

如果你也在开发一个长期运行的服务，希望本文的介绍能对你有所帮助，让你的服务更加健壮和可靠。