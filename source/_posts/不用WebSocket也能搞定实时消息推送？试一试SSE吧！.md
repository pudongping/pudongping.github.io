---
title: 不用WebSocket也能搞定实时消息推送？试一试SSE吧！
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
abbrlink: 5a7767cc
date: 2025-06-18 09:49:34
img:
coverImg:
password:
summary:
---

在现代 Web 开发中，实现实时数据更新是一个常见的需求。比如股票行情、聊天消息、体育比赛比分等场景，都需要服务器能够主动将数据推送给客户端，而不是客户端频繁轮询服务器来获取最新数据。

今天，我们就来学习如何使用 Go 语言和 Gin 框架实现 SSE（服务器发送事件）来完成这样一个实时时间推送的功能。

## 一、SSE 技术简介

### 1.1 什么是 SSE？

SSE（Server-Sent Events）是一种简单的服务器向客户端推送实时通知的技术。它基于 HTTP 协议，使用文本格式传输数据。相比 WebSocket，SSE 的实现更为简单，因为它只需要服务器端发送数据到客户端，而不需要处理客户端到服务器的双向通信。SSE 使用 `text/event-stream` MIME 类型，并且通过 `EventSource` JavaScript 接口在浏览器中使用。

与传统的轮询机制相比，SSE 具有以下优势：

- 单向实时通信（服务端 → 客户端）
- 基于 HTTP 协议，无需复杂握手
- 自动重连机制
- 轻量级且易于实现

### 1.2 与 WebSocket 的对比
| 特性               | SSE                     | WebSocket          |
|--------------------|-------------------------|--------------------|
| 通信方向           | 单向（服务器到客户端） | 双向               |
| 协议               | HTTP                   | 独立协议           |
| 数据格式           | 文本                   | 二进制/文本        |
| 断线重连           | 自动支持               | 需手动实现         |
| 浏览器兼容性       | IE 不支持              | 较新浏览器均支持   |

## 二、服务端实现（Go + Gin 框架）

> 新建 `sse.go` 文件，并复制一下 go 代码。

### 2.1 完整代码结构
```go
package main

import (
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	// 静态页面路由（用于展示前端）
	r.StaticFile("/", "./index.html")

	// SSE 事件流路由
	r.GET("/events", func(c *gin.Context) {
		// 设置SSE必要响应头
		c.Header("Content-Type", "text/event-stream")
		c.Header("Cache-Control", "no-cache")
		c.Header("Connection", "keep-alive")

		// 创建通道用于接收关闭通知
		clientClosed := c.Writer.CloseNotify()

		// 无限循环发送事件
		ticker := time.NewTicker(1 * time.Second)
		defer ticker.Stop()

		for {
			select {
			case <-clientClosed:
				fmt.Println("客户端断开连接")
				return
			case t := <-ticker.C:
				// SSE数据格式要求（重要！）
				event := fmt.Sprintf("data: %v\n\n", t.Format("2006-01-02 15:04:05"))

				// 发送数据到客户端
				c.SSEvent("message", event)
				c.Writer.Flush() // 立即刷新缓冲区
			}
		}
	})

	r.Run(":8080")
}
```

### 2.2 关键代码解析

#### 响应头设置
```go
c.Header("Content-Type", "text/event-stream")
c.Header("Cache-Control", "no-cache")
c.Header("Connection", "keep-alive")
```
这三个响应头是 SSE 正常工作所必需的：

- `text/event-stream` 声明事件流类型
- `no-cache` 禁用缓存确保实时性
- `keep-alive` 保持长连接

#### 心跳机制实现
```go
ticker := time.NewTicker(1 * time.Second)
defer ticker.Stop()

for {
	select {
	case <-clientClosed:
		// 处理客户端断开
	case t := <-ticker.C:
		// 发送时间数据
	}
}
```
通过 `time.Ticker` 创建定时器，每秒触发一次数据推送，形成持续的数据流。

## 三、客户端实现（HTML + JavaScript）

> 需要在 `sse.go` 统计目录下创建 `index.html` 文件。

### 3.1 前端页面代码
```html
<!DOCTYPE html>
<html>
<body>
	<h1>SSE 实时时间推送演示</h1>
	<div id="output"></div>

	<script>
		const eventSource = new EventSource('/events');
		const output = document.getElementById('output');

		// 消息接收处理
		eventSource.onmessage = (e) => {
			output.innerHTML += e.data + '<br>';
		};

		// 错误处理
		eventSource.onerror = (e) => {
			console.error('连接异常:', e);
			eventSource.close();
		};
	</script>
</body>
</html>
```

### 3.2 客户端注意事项
- `EventSource` 会自动处理连接建立和重连
- 消息接收使用 `onmessage` 回调函数
- 实际生产环境需要添加错误处理和状态监控
- 支持以下事件类型：

```js
eventSource.addEventListener('customEvent', (e) => {
  console.log('自定义事件:', e.data);
});
```

## 四、运行与测试

### 4.1 启动服务
```bash
# 安装依赖
go get github.com/gin-gonic/gin

# 运行服务
go run sse.go
```

### 4.2 访问测试
1. 浏览器访问 `http://localhost:8080`
2. 打开开发者工具 → Network 面板
3. 观察 EventStream 数据流：

![SSE 网络请求示例](https://upload-images.jianshu.io/upload_images/14623749-9ee92e77539e16ad.png?imageMogr2/auto-orient/strip%7CimageView2/2/w/1240)

**以上的示例，正常在浏览器中访问时，会持续不断的每隔 1 秒打印出当前服务器的时间。**

## 五、扩展应用场景

### 5.1 常见使用案例
1. 实时监控仪表盘（服务器状态、在线人数）
2. 新闻/股票实时行情推送
3. 聊天应用的在线状态提示
4. 长耗时任务的进度更新

### 5.2 性能优化建议
```go
// 示例：连接数控制
var clients = make(map[string]chan struct{})
func sseHandler(c *gin.Context) {
	// 创建唯一客户端标识
	clientID := uuid.New().String()
	clients[clientID] = make(chan struct{})
	
	defer func() {
		delete(clients, clientID)
	}()
	
	// ...原有代码...
}
```

- 增加客户端数量限制
- 实现消息广播功能
- 添加心跳超时检测
- 使用连接池管理资源

## 六、注意事项

1. **浏览器兼容性：** SSE 在大多数现代浏览器中都得到了很好的支持，但在一些老旧的浏览器（如 IE）中可能不支持。如果你需要支持这些浏览器，可能需要考虑使用其他技术，如 WebSocket 或者轮询。
2. **服务器性能：** 由于 SSE 需要保持连接打开，如果连接数量过多，可能会对服务器性能产生影响。在实际生产环境中，需要考虑使用连接池、负载均衡等技术来优化服务器性能。
3. **数据格式：** SSE 的数据格式有一定的规范，除了 `data` 字段外，还可以使用 `event`、`id` 等字段来扩展功能。在发送复杂数据时，可以考虑使用 JSON 格式，并在前端进行解析处理。


通过本文的学习，我们了解了如何使用 Go 语言和 Gin 框架实现 SSE 来完成实时时间推送的功能。

SSE 是一种简单高效的服务器到客户端单向通信技术，在许多实时性要求不高的场景中非常实用。希望本文能够帮助你掌握 SSE 的基本用法，并在实际项目中应用这一技术来提升用户体验。