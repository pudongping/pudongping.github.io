---
title: Go Web服务中如何优雅关机？
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
abbrlink: 2e2c4d21
date: 2024-11-14 11:54:11
img:
coverImg:
password:
summary:
---

在构建 Web 服务时，我们往往会遇到一个棘手的问题：**当我们想要停止服务时，如何确保正在处理的请求能够顺利完成，而不是突然中断？** 这种技术被称为“优雅关机”，它可以确保在服务关闭时，所有的请求都被妥善处理。

在这篇文章中，我们将通过一个简单的例子来演示如何在 Go 语言中使用 Gin 框架实现优雅关机。

## 什么是优雅关机？

优雅的关机是指在关闭服务之前，先让服务处理完当前正在处理的请求，然后再关闭服务。这样可以保证服务不会丢失请求，也不会影响到正在处理的请求。这种方式可以提高用户体验，防止服务中断造成的数据丢失或不一致。
而执行 `Ctrl + C` 或者 `kill -2 pid` 命令关闭服务，是不会等待服务处理完请求的，这样就会导致服务丢失请求。

## 如何实现优雅的关机？

Go 1.8 版本之后，`http.Server` 内置的 `Shutdown()` 方法就支持优雅地关机。

## 代码实现

我们来看一个具体的代码示例，通过这个例子我们将展示如何实现优雅关机。

```go
package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	// 定义一个简单的路由
	router.GET("/ping", func(c *gin.Context) {
		// 模拟一个耗时操作，比如数据库查询或外部API调用
		time.Sleep(5 * time.Second)
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	// 配置 HTTP 服务器
	srv := &http.Server{
		Addr:    ":8080",
		Handler: router,
	}

	// 启动服务器
	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("Failed to start server: %v", err)
		}
	}()

	// 等待中断信号来优雅地关闭服务器，为关闭服务器操作设置一个 5 秒的超时
	quit := make(chan os.Signal, 1) // 创建一个接收信号的通道

	// kill 默认会发送 syscall.SIGTERM 信号
	// kill -2 发送 syscall.SIGINT 信号，我们常用的 `Ctrl+C` 就是触发系统 SIGINT 信号
	// kill -9 发送 syscall.SIGKILL 信号，但是不能被捕获，所以不需要添加它
	// signal.Notify 把收到的 syscall.SIGINT 或 syscall.SIGTERM 信号转发给 quit
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM) // 此处不会阻塞

	<-quit                                               // 阻塞在此，当接收到上述两种信号时才会往下执行
	log.Println("Shutdown Server ...")

	// 创建一个 5 秒超时的 context
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	// 5 秒内优雅关闭服务（将未处理完的请求处理完再关闭服务），超过 5 秒就超时退出
	if err := srv.Shutdown(ctx); err != nil {
		log.Fatal("Server Shutdown: ", err)
	}

	log.Println("Server exiting")

}
```

## 代码解析

### 1. 路由定义和服务启动

```go
router := gin.Default()
router.GET("/ping", func(c *gin.Context) {
	time.Sleep(5 * time.Second)
	c.JSON(http.StatusOK, gin.H{
		"message": "pong",
	})
})
```

首先，我们创建了一个简单的 Gin 路由，并定义了一个 `/ping` 接口。当访问这个接口时，服务器会模拟一个耗时 5 秒的操作，然后返回一个 JSON 响应。这段代码展示了一个可能需要优雅关机的典型场景：服务器可能正在处理耗时的请求，如果此时直接关机，请求会被中断。

### 2. HTTP 服务器配置和启动

```go
srv := &http.Server{
	Addr:    ":8080",
	Handler: router,
}

go func() {
	if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		log.Fatalf("Failed to start server: %v", err)
	}
}()
```

我们使用 `http.Server` 结构体配置并启动了一个 HTTP 服务器。服务器在一个单独的 goroutine 中运行，这样主程序可以继续执行，而不必等待服务器启动完成。

### 3. 捕获系统信号

```go
quit := make(chan os.Signal, 1)
signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)

<-quit
```

为了实现优雅关机，我们需要捕获系统信号。这里使用了 `os/signal` 包来监听 `syscall.SIGINT` 和 `syscall.SIGTERM` 信号。当用户按下 `Ctrl+C` 或者通过 `kill` 命令发送信号时，这些信号会被捕获并发送到 `quit` 通道，程序会随即从阻塞状态中恢复，继续执行后续代码。

### 4. 实现优雅关机

```go
ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
defer cancel()

if err := srv.Shutdown(ctx); err != nil {
	log.Fatal("Server Shutdown: ", err)
}
```

在捕获到关机信号后，我们使用 `http.Server` 的 `Shutdown` 方法来实现优雅关机。`Shutdown` 方法接受一个 `context` 参数，这个 `context` 设置了一个超时时间。在这里，我们设置了一个 5 秒的超时时间，意味着服务器将在 5 秒内等待未完成的请求处理完毕，然后关闭。如果超过了设定的超时时间，服务器将退出，程序也会正常结束。

## 如何验证优雅关机的效果？

要验证优雅关机的效果，可以按照以下步骤操作：

1. 打开终端，运行 go run gin_shutdown.go
2. 打开浏览器，并访问 `http://127.0.0.1:8080/ping` 此时浏览器应该会白屏等待服务端返回响应
3. 在刚刚打开的终端上迅速按下 Ctrl+C 命令，此时会自动给程序发送 syscall.SIGINT 信号
4. 此时程序并不会立即退出，而是会等上面的第 **2** 步的响应返回之后再退出，从而实现优雅关机的效果

## 总结

优雅关机是构建健壮 Web 服务的一个重要技术点，它确保了在服务关闭时所有正在处理的请求都能被妥善完成。在本文中，我们通过 Gin 框架演示了如何在 Go 中实现优雅关机。通过这种方式，我们可以提升用户体验，减少由于服务中断导致的各种潜在问题。

希望这篇文章能够帮助你更好地理解和实现 Go 服务中的优雅关机。如果你有任何问题或建议，欢迎在评论区与我讨论！