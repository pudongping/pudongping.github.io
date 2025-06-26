---
title: 如何为 gRPC Server 编写本地测试代码
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
  - gRPC
  - RPC
abbrlink: e9966a07
date: 2025-06-26 09:52:58
img:
coverImg:
password:
summary:
---

在微服务架构中，gRPC 已成为主流的通信协议之一。但许多开发者在面对 gRPC 服务测试时，常常会遇到需要启动真实网络服务、管理端口占用等烦恼。

本文将介绍如何利用 Go 语言中 gRPC 提供的测试工具 —— bufconn，通过构建内存级别的网络连接，来实现 gRPC Server 的本地测试，而无需占用实际端口。

## 一、测试环境简介

在传统的测试场景中，我们通常需要启动一个 gRPC Server 并通过实际网络端口进行测试。但这种方式可能会引起端口冲突、环境依赖等问题。为了避免这些问题，我们可以使用 **bufconn** 包创建一个内存中的网络连接，从而实现完全隔离的测试环境。

> **bufconn** ：一个基于内存缓冲区的网络连接模拟器，它能够让我们在本地直接测试 gRPC Server，而无需实际监听端口。

## 二、代码实现解析

下面我们以示例代码为基础，详细讲解每个部分的作用和实现细节。

### 2.1 导入相关依赖

首先，我们需要导入 gRPC、测试框架和 bufconn 包。示例代码中还引入了 `context`、`log`、`net` 和 `testing` 等标准包，保证测试环境的完整性。

```go
package main

import (
	"context"
	"log"
	"net"
	"testing"

	"go-tutorial/project/gokit_learn/sample_grpc_srv/pb" // 生成的 pb 文件

	"github.com/stretchr/testify/assert"  // 测试断言库
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"  // 明文传输，不启用加密
	"google.golang.org/grpc/test/bufconn"  // bufconn 实现内存中的网络连接
)
```

> **说明**：在此示例中，我们使用了 [ testify ] 断言库来方便地对测试结果进行验证，确保代码的正确性。


### 2.2 初始化内存连接

在 `init` 函数中，我们通过 `bufconn.Listen` 创建一个内存缓冲区，并启动一个 gRPC Server。通过这种方式，服务会在内存中运行，而无需实际监听端口。

```go
// 设置缓冲区大小为 1 MB
const bufSize = 1024 * 1024

var bufListener *bufconn.Listener

func init() {
	// 创建内存中的网络监听器
	bufListener = bufconn.Listen(bufSize)
	
	// 创建一个新的 gRPC Server 实例
	s := grpc.NewServer()
	// 初始化业务逻辑，例如 addService 实现了对应接口
	gs := NewGRPCServer(addService{})
	// 注册 gRPC 服务
	pb.RegisterAddServer(s, gs)
	
	// 异步启动服务
	go func() {
		if err := s.Serve(bufListener); err != nil {
			log.Fatalf("Server exited with error: %v", err)
		}
	}()
}
```

> **注释**：
> 1. `bufconn.Listen` 返回一个内存连接监听器，该监听器模拟网络监听。
> 2. 使用 `go` 关键字启动一个 goroutine 异步执行 `s.Serve(bufListener)`，保证不会阻塞主线程。

---

### 2.3 自定义拨号器

由于服务运行在内存中，我们需要自定义一个拨号器函数，使得 gRPC 客户端能够通过 bufconn 连接到服务端。

```go
// bufDialer 实现了自定义拨号器，返回内存中的网络连接
func bufDialer(context.Context, string) (net.Conn, error) {
	return bufListener.Dial() // 使用内存连接代替真实网络
}
```

> **说明**：此函数会在测试时通过 `grpc.WithContextDialer` 传入 gRPC 客户端，使得客户端请求都能走内存连接。

---

### 2.4 编写测试用例

接下来，我们提供了一个测试用例来测试 gRPC 服务中的 `Sum` 方法。

> 我们不需要理会 `Sum` 方法的具体实现，我们重点关注如何在不启动真实网络服务的情况下进行 GRPC 本地测试。

#### 测试 Sum 方法

```go
func TestSum(t *testing.T) {
	// 使用自定义拨号器建立连接
	conn, err := grpc.DialContext(
		context.Background(),
		"bufnet",  // 虚拟地址
		grpc.WithContextDialer(bufDialer),
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	
	// 创建 gRPC 客户端
	c := pb.NewAddClient(conn)

	// 发送请求，计算 10 + 2
	resp, err := c.Sum(context.Background(), &pb.SumRequest{
		A: 10,
		B: 2,
	})
	// 断言：错误应为空，返回结果不为空，且计算结果为 12
	assert.Nil(t, err)
	assert.NotNil(t, resp)
	assert.Equal(t, int64(12), resp.V)
}
```

> **关键点**：
> 1. 使用 `grpc.DialContext` 与自定义的 `bufDialer` 建立连接，避免真实网络调用。
> 2. 通过断言库验证返回结果是否符合预期。

---

## 三、总结

通过本文的介绍，我们可以了解到：

- 利用 **bufconn** 可以构建一个内存中的网络环境，从而避免实际网络端口冲突；
- 自定义拨号器使得客户端能够轻松连接内存中的 gRPC Server；
- 利用 **bufconn** 进行测试，仍然需要完整的 GRPC 服务实现，只是通信方式不同；
- 使用断言库（例如 [ testify ]）可以方便地对测试结果进行验证，提高测试的可读性与准确性。

这种方式不仅适用于单元测试，也为后续的集成测试提供了良好的基础。希望本文能帮助大家更好地理解如何为 gRPC Server 编写本地测试代码，并在实际项目中加以应用。

那么，在工作中，你一般是怎么处理的呢？你也有遇到过类似的问题吗？欢迎一起讨论讨论呀～