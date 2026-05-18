---
title: Celery 太重了？这可能是你一直在找的 asyncio 任务队列
author: Alex
top: false
hide: false
cover: false
toc: true
mathjax: false
categories: Python
tags:
  - Python3
  - Celery
  - arq
abbrlink: 35e2046b
date: 2026-05-18 11:03:30
img:
coverImg:
password:
summary:
---

你是否在配置 Celery 时，被复杂的 Broker、Backend、Worker 搞得头皮发麻？你是否在 Python `async/await` 的世界里，发现老牌的任务队列总是显得“格格不入”？

在 Python 异步编程（Asynchronous Programming）日益普及的今天，我们需要一个**原生支持 asyncio**、**极度轻量**且**性能强悍**的任务队列。

今天，我要向你推荐一个“小而美”的神器——**arq**。

💡 **读完本文，你将获得：**
1.  彻底理解 arq 的核心设计理念（小白也能懂）。
2.  掌握 arq 的安装、配置与实战代码（含优雅写法）。
3.  深入理解其基于 Redis Streams 的底层原理（高手进阶）。

---

## 🧐 What & Why

### 什么是 arq？
简单来说，`arq` 是一个基于 **Python asyncio** 和 **Redis** 的作业队列（Job Queue）。它的作者是 `pydantic` 的大神 Samuel Colvin（没错，就是那个写数据验证库的大佬）。

### 💡 一个生活化的比喻
为了让你秒懂 arq 和 Celery 的区别，我们想象一下**“餐厅后厨”**的场景：

*   **Celery（老牌霸主）**：
    就像一个**五星级酒店的行政总厨**。功能极其强大，能做中餐、西餐、日料；有专门的切菜部、炒菜部、传菜部。但是，如果你只想煮一碗泡面，启动这一整套流程就显得非常笨重、繁琐，启动慢，资源消耗大。

*   **arq（新晋网红）**：
    就像一个**智能全自动炒菜机器人**。它只做一件事：高效地处理订单。它直接连接“菜篮子”（Redis），用最快的速度（asyncio）并行处理几百个订单。它没有复杂的层级，**轻便、极速、开箱即用**。

### 为什么要选 arq？
1.  **原生异步**：天生支持 `async/await`，与 FastAPI、Sanic 等异步框架是绝配。
2.  **极简依赖**：只依赖 `redis-py`，没有复杂的依赖树。
3.  **高性能**：利用 Redis Streams 和 asyncio，并发能力极强。

---

## 🛠️ How

### 1. 快速上手

首先，安装它：
```bash
pip install arq
```

### 2. 定义 Worker（消费者）

arq 的 worker 定义非常简单，通常我们创建一个 `worker.py`：

```python
# worker.py
import asyncio
from arq import create_pool
from arq.connections import RedisSettings

# 1. 定义具体的任务函数
async def say_hello(ctx, name: str):
    """
    一个简单的异步任务
    :param ctx: 上下文，包含 redis 连接等信息
    :param name: 任务参数
    """
    await asyncio.sleep(1) # 模拟耗时操作
    print(f"👋 Hello {name}, 任务完成！")
    return f"Hello {name}"

# 2. 定义 Worker 配置
class WorkerSettings:
    # Redis 配置
    redis_settings = RedisSettings(host='localhost', port=6379)
    # 注册任务函数
    functions = [say_hello]
    # 并发数
    max_jobs = 10 
    
    # 💡 优雅写法：生命周期管理
    async def on_startup(self):
        print("🚀 Worker 启动啦！")

    async def on_shutdown(self):
        print("🛑 Worker 停止啦！")
```

启动 Worker：
```bash
arq worker.WorkerSettings
```

### 3. 发布任务（生产者）

在你的业务代码（例如 FastAPI 接口）中这样调用：

```python
# main.py
import asyncio
from arq import create_pool
from arq.connections import RedisSettings

async def main():
    # 1. 创建 Redis 连接池
    redis = await create_pool(RedisSettings(host='localhost', port=6379))
    
    # 2. 入队任务
    # 这里的 'say_hello' 必须和 worker 中注册的函数名一致
    print("📤 正在发送任务...")
    await redis.enqueue_job('say_hello', name='World')
    
    # 3. 关闭连接
    await redis.close()

if __name__ == '__main__':
    asyncio.run(main())
```

### ⚠️ 避坑指南

1.  **不要在任务中执行同步阻塞代码**：
    arq 是基于 `asyncio` 的，如果你在任务里写了 `time.sleep(10)` 或者同步的 `requests.get()`，整个 Worker 的事件循环会被卡死！
    *   ✅ 正确：`await asyncio.sleep(10)` 或 `httpx.AsyncClient`
    *   ❌ 错误：`time.sleep(10)`
2.  **参数序列化陷阱**：
    arq 默认使用 `pickle` 进行序列化，这很方便，支持几乎所有 Python 对象。但在跨语言或安全性要求高的场景，建议配置为 `JSON` 序列化。
3.  **任务去重**：
    arq 支持 `_job_id` 参数。如果你希望同一个任务 ID 在同一时间只执行一次（防止重复扣款等），可以在入队时指定 `_job_id`。

---

## 多聊一聊

对于资深玩家，你可能想问：**arq 到底是怎么利用 Redis 实现队列的？**

### 1. Redis Streams 的妙用
在 arq 的早期版本中，它使用 Redis List (`LPUSH`/`BRPOP`)。但从 v0.16 开始，它全面拥抱了 **Redis Streams**（Redis 5.0+ 新特性）。

*   **可靠性**：Streams 支持 **Consumer Groups**（消费者组）。这意味着如果一个 Worker 拿走了任务但崩溃了（没有 ACK），任务不会丢失，会保留在 Pending List 中，等待被其他 Worker 认领（Claim）。
*   **持久化**：相比 Pub/Sub 的“发后即焚”，Streams 是持久化的日志结构。

### 2. Lua 脚本保证原子性
arq 大量使用了 Redis **Lua 脚本**来保证操作的原子性。
比如“入队”操作，不仅要写入 Stream，可能还需要检查是否有延迟任务（ZSET），或者是否有唯一性约束。arq 将这些逻辑封装在 Lua 脚本中，一次网络往返即可完成，既快又安全。

### 3. 延迟任务 (Deferred Jobs)
arq 如何实现 `enqueue_job('task', _defer_by=60)`（延迟60秒执行）？
*   它并不是让 Worker `sleep` 60秒。
*   它是将任务放入 Redis 的 **Sorted Set (ZSET)** 中，score 是执行时间戳。
*   Worker 内部有一个协程轮询这个 ZSET，时间一到，立马把任务“搬运”到 Streams 队列中供消费。

---

## 总结

*   **定位**：轻量、异步、基于 Redis。
*   **核心组件**：WorkerSettings, enqueue_job, RedisSettings。
*   **底层**：Redis Streams (队列) + ZSET (延迟/定时) + Lua (原子性)。
*   **适用场景**：高并发 I/O 任务、轻量级微服务、FastAPI 背景任务。

如果你的业务中既有**CPU 密集型任务**（如视频转码，会阻塞 EventLoop），又有**I/O 密集型任务**（如发邮件），你会如何设计 arq 的 Worker 架构？是混合部署还是拆分部署？

欢迎在评论区留下你的架构方案！

