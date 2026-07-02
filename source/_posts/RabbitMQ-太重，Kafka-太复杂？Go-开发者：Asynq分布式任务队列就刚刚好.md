---
title: RabbitMQ 太重，Kafka 太复杂？Go 开发者：Asynq分布式任务队列就刚刚好
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
  - Asynq
  - 分布式消息队列
abbrlink: 3a7f5eac
date: 2026-07-02 10:56:59
img:
coverImg:
password:
summary:
---

在后端服务中，我们经常需要**异步执行耗时任务** —— 比如发送邮件、生成缩略图、导出报表、调用外部 API 等。让请求等待这些操作完成，会浪费用户体验，也浪费服务器资源。

这时异步任务队列就派上用场了。

而在 Go 语言中，`asynq` 是一个非常流行且易用的异步任务队列库，它基于 Redis 构建，提供了可靠的任务执行、调度、重试等能力。

## 一、什么是 Asynq？为什么要用它？

简单来说：

> **Asynq 是一个 Go 语言的分布式任务队列库，基于 Redis 实现，支持高并发执行、异步队列、延迟队列、失败重试、优先级队列、定时任务等功能。**

它适合我们在生产环境中处理需要异步执行的任务，例如：

*   发送用户激活邮件
*   图片或视频处理
*   定时任务
*   调度某些延迟执行操作，如：订单超时取消

## 二、Asynq 的基本工作原理

在 Asynq 中，涉及三个主要角色：

1.  **Client（生产者）** —— 创建任务放入队列
2.  **Server（消费者/worker）** —— 从队列取出任务并执行
3.  **Redis** —— 作为消息代理，存储任务以及协同客户端和服务端通信

执行流程大致如下：

```bash
Client -> Redis (任务入队) -> Server 从 Redis 取出任务 -> Worker 处理任务
```

## 三、准备工作：Redis 环境与依赖安装

要使用 Asynq，首先需要有一个 Redis 服务器（版本 4.0+）。

可以通过 Docker 快速启动：

```bash
docker run -d -p 6379:6379 redis:latest
```

初始化 Go 项目：

```bash
mkdir asynq-demo
cd asynq-demo
go mod init github.com/yourname/asynq-demo
```

安装库：

```bash
go get github.com/hibiken/asynq
```

## 四、定义任务与创建任务

### 1\. 任务的概念

在 Asynq 中，**任务由类型和负载（payload）组成**。

*   *类型* 用于区分任务，例如 `"email:send"`
*   *payload* 是任务数据，一般使用 JSON 编码携带信息（用户 ID、文件 URL 等）

### 2\. 示例：创建一个发送邮件的任务

在 `tasks/tasks.go` 文件中：

```go
package tasks

import (
	"encoding/json"
	"github.com/hibiken/asynq"
)

// 任务类型
const TypeEmailSend = "email:send"

// 邮件发送所需数据
type EmailPayload struct {
	To      string
	Subject string
	Body    string
}

// 创建一个发送邮件任务
func NewEmailTask(to, subject, body string) (*asynq.Task, error) {
	p := EmailPayload{To: to, Subject: subject, Body: body}
	data, err := json.Marshal(p)
	if err != nil {
		return nil, err
	}
	// NewTask 接收类型和负载
	return asynq.NewTask(TypeEmailSend, data), nil
}

```

在这个函数中，我们将数据序列化为 JSON，然后创建一个 `*asynq.Task` 对象。

## 五、将任务加入队列（Client 端）

编写 `main.go`

```go
package main

import (
	"log"
	"github.com/hibiken/asynq"
	"yourmodule/tasks"
)

func main() {
	// 创建 Client，连接 Redis
	client := asynq.NewClient(asynq.RedisClientOpt{Addr: "127.0.0.1:6379"})
	defer client.Close()

	// 创建任务
	task, err := tasks.NewEmailTask("alice@example.com", "Hello", "Welcome to Asynq!")
	if err != nil {
		log.Fatalf("无法创建任务: %v", err)
	}

	// 入队
	info, err := client.Enqueue(task)
	if err != nil {
		log.Fatalf("无法入队任务: %v", err)
	}

	log.Printf("任务已入队 ID=%s 队列=%s", info.ID, info.Queue)
}

```

运行：

```bash
go run main.go
```

如果一切正常，会看到日志输出说明任务已成功放入队列。

## 六、处理任务（Server 端）

任务入队后，需要有 worker 从队列取出并执行任务逻辑。

### 1\. Handler 定义

在 `worker/worker.go` 中：

```go
package main

import (
	"context"
	"encoding/json"
	"log"

	"github.com/hibiken/asynq"
	"yourmodule/tasks"
)

// 处理 EmailSend 任务
func handleEmailSend(ctx context.Context, t *asynq.Task) error {
	var payload tasks.EmailPayload
	if err := json.Unmarshal(t.Payload(), &payload); err != nil {
		log.Println("Payload 解析失败:", err)
		return err
	}

	log.Printf("发送邮件给: %s 标题: %s", payload.To, payload.Subject)
	// TODO: 真实邮件发送逻辑
	return nil
}

func main() {
	// 创建 Server
	srv := asynq.NewServer(
		asynq.RedisClientOpt{Addr: "127.0.0.1:6379"},
		asynq.Config{
			Concurrency: 5,                  // 并发 worker 数量
			Queues: map[string]int{"default": 1}, // 设置队列优先级
		},
	)

	// ServeMux 用于注册任务处理函数
	mux := asynq.NewServeMux()
	mux.HandleFunc(tasks.TypeEmailSend, handleEmailSend)

	// 启动服务
	if err := srv.Run(mux); err != nil {
		log.Fatal("Worker 运行失败:", err)
	}
}

```

运行：

```bash
go run worker/worker.go
```

此时 worker 会从 Redis 中取出任务并执行 `handleEmailSend`。

## 七、调度任务（延迟 / 定时）

Asynq 支持任务延迟执行：

```go
info, err := client.Enqueue(task, asynq.ProcessIn(10*time.Second))
if err != nil {
	// ...
}
```

上面的代码让任务在 *10 秒后* 执行。
也可以使用 `asynq.ProcessAt(time)` 指定绝对时间。

## 八、失败重试与超时控制

默认情况下，如果 handler 返回错误，任务会重试若干次。
你可以自定义：

```go
task, _ := tasks.NewEmailTask(...)
client.Enqueue(task,
	asynq.MaxRetry(5),       // 最大重试次数
	asynq.Timeout(30*time.Second), // 执行超时
)
```

这样，对于失败的任务，可以进行自动重试，并设置 task 执行的最长时间。

## 九、任务优先级队列

你可以定义多个队列，并为它们设置不同权重：

```go
srv := asynq.NewServer(
	asynq.RedisClientOpt{Addr: "127.0.0.1:6379"},
	asynq.Config{
		Queues: map[string]int{
			"critical": 6,
			"default": 3,
			"low": 1,
		},
	},
)
```

权重高的队列会被更频繁地消费。适合区分紧急与普通任务。


## 十、可视化与监控

Asynq 提供了专门的 **Web UI 和 CLI 工具**：

命令行：

```bash
go install github.com/hibiken/asynq/tools/asynq@latest
asynq stats      # 显示队列统计信息
asynq dash       # 启动交互式监控界面
```

Web UI 可以直观查看队列和任务状态。

![queues](https://upload-images.jianshu.io/upload_images/14623749-37f7dfa928a91a45.png?imageMogr2/auto-orient/strip%7CimageView2/2/w/1240)

## 十一、其他常用功能

1.  **任务重试 (Retry)**
    *   默认情况下，如果 Handler 返回错误，Asynq 会自动重试（默认 25 次，指数退避）。
    *   可以通过 `asynq.MaxRetry(3)` 修改重试次数。
    *   可以通过 `asynq.SkipRetry` 错误来跳过重试（例如数据格式错误，重试也没用）。

2.  **唯一任务 (Unique)**
    *   防止任务重复入队。例如：确保 1 分钟内不要给同一个用户发送两封相同的邮件。
    *   `asynq.Unique(1 * time.Minute)`

3.  **超时控制 (Timeout)**
    *   如果任务执行时间过长，可以强制超时。
    *   `asynq.Timeout(10 * time.Second)`


如果你正在构建微服务、后台任务系统，或者需要提升系统性能和可伸缩性，Asynq 是一个非常值得尝试的工具。

相关示例代码可见 https://github.com/pudongping/golang-tutorial/tree/main/project/asynq_learn 提供了非常详细的异步队列、延迟队列、优先级队列（加权队列）、定时任务相关示例，希望对你有所帮助～