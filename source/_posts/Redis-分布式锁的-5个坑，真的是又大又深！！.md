---
title: Redis 分布式锁的 5个坑，真的是又大又深！！
author: Alex
top: false
hide: false
cover: false
toc: true
mathjax: false
categories: Redis
tags:
  - Redis
  - 分布式锁
abbrlink: d47f830
date: 2026-07-06 09:59:18
img:
coverImg:
password:
summary:
---

在微服务架构中，分布式锁是我们解决“并发抢占资源”（如秒杀扣库存、防重复提交等）的利器。很多开发者认为，用 Redis 实现分布式锁不就是一行 `SETNX` 的事吗？但实际上，从“能跑的玩具代码”到“生产级的高可用组件”，中间隔着无数个暗坑。

![](https://upload-images.jianshu.io/upload_images/14623749-78c76a73b9345ddb.jpg?imageMogr2/auto-orient/strip%7CimageView2/2/w/1240)


今天，我们将重新审视 Redis 分布式锁。我们不仅要填平那 5 个著名的“血坑”，还要引入 Go 语言特有的并发哲学（如 `context` 控制、Goroutine 泄漏防范、无 Goroutine ID 的重入锁设计），最后甚至要探讨分布式系统泰斗 Martin Kleppmann 对 Redis 锁的灵魂拷问。

准备好了吗？发车！

## 演进一：基础实现与原子性陷阱 (The Basics)

**新手的做法**：先 `SetNX`，再 `Expire`。中间如果程序崩溃（OOM 或重启），锁永远无法释放，直接死锁。
**进阶的做法**：使用 `SET key value NX EX time` 原子命令。

但在 Go 的工程实践中，我们绝不能写面条代码。一个中高级的 Go 开发者，首先想到的是**封装与面向对象设计**。我们会使用 Go 经典的 **Functional Options 模式** 来构建锁对象。

```go
package redislock

import (
	"context"
	"crypto/rand"
	"encoding/hex"
	"errors"
	"time"

	"github.com/redis/go-redis/v9"
)

var ErrLockFailed = errors.New("failed to acquire lock")

// Client 抽象 Redis 客户端，方便 mock 测试
type Client interface {
	SetNX(ctx context.Context, key string, value interface{}, expiration time.Duration) *redis.BoolCmd
	Eval(ctx context.Context, script string, keys []string, args ...interface{}) *redis.Cmd
}

type Options struct {
	Expiration time.Duration
	RetryWait  time.Duration
	MaxRetries int
}

type Option func(*Options)

// RedisLock 生产级分布式锁结构体
type RedisLock struct {
	client Client
	key    string
	token  string // 锁的唯一标识，防误删
	opts   Options
}

// NewRedisLock 构造函数
func NewRedisLock(client Client, key string, optFuncs ...Option) *RedisLock {
	opts := Options{
		Expiration: 30 * time.Second,
		RetryWait:  50 * time.Millisecond,
		MaxRetries: 3,
	}
	for _, f := range optFuncs {
		f(&opts)
	}
	
	return &RedisLock{
		client: client,
		key:    key,
		token:  generateToken(), // 生成唯一 Token
		opts:   opts,
	}
}

// generateToken 生成 16 字节的随机 hex 字符串
func generateToken() string {
	b := make([]byte, 16)
	rand.Read(b)
	return hex.EncodeToString(b)
}
```

---

## 演进二：安全解锁与 Lua 脚本 (Safe Unlock)

**踩坑点**：协程 A 的锁超时了，协程 B 抢到了锁。协程 A 此时跑完业务，执行 `DEL key`，把协程 B 的锁删了，导致系统雪崩。
**专家解法**：解锁时必须校验 `token`，且“判断+删除”必须是原子的。

在 Go 中，我们通常会将 Lua 脚本预加载（`SCRIPT LOAD`）或者直接使用 `Eval`。

```go
const unlockScript = `
if redis.call("get", KEYS[1]) == ARGV[1] then
    return redis.call("del", KEYS[1])
else
    return 0
end
`

// Unlock 安全解锁
func (l *RedisLock) Unlock(ctx context.Context) error {
	res, err := l.client.Eval(ctx, unlockScript, []string{l.key}, l.token).Result()
	if err != nil {
		return err
	}
	if n, ok := res.(int64); !ok || n == 0 {
		return errors.New("unlock failed: lock not held by this token")
	}
	return nil
}
```

---

## 演进三：锁超时与看门狗机制 (Watchdog)

**踩坑点**：锁 TTL 设为 30 秒，但遇到慢查询，业务跑了 40 秒。第 30 秒时锁自动失效，并发冲突产生。
**专家解法**：引入 Watchdog（看门狗）后台协程自动续期。

作为中高级 Go 开发者，写后台 Goroutine 必须考虑**生命周期管理**和**防 Goroutine 泄漏**。千万不能直接 `go func() { for {} }()`。必须结合 `context` 和 `chan` 优雅退出。

```go
// 续期 Lua 脚本
const renewScript = `
if redis.call("get", KEYS[1]) == ARGV[1] then
    return redis.call("pexpire", KEYS[1], ARGV[2])
else
    return 0
end
`

// startWatchdog 启动看门狗
func (l *RedisLock) startWatchdog(ctx context.Context, done chan struct{}) {
	// 续期周期通常为超时时间的 1/3
	ticker := time.NewTicker(l.opts.Expiration / 3)
	defer ticker.Stop()

	for {
		select {
		case <-ticker.C:
			// 执行续期
			res, err := l.client.Eval(ctx, renewScript, []string{l.key}, l.token, int64(l.opts.Expiration/time.Millisecond)).Result()
			// 如果锁已经被释放（返回0），或者发生错误，看门狗退出
			if err != nil || res.(int64) == 0 {
				return 
			}
		case <-done:
			// 业务主动解锁，通知看门狗退出
			return
		case <-ctx.Done():
			// 父 Context 取消，看门狗退出
			return
		}
	}
}
```

注意这里的 `select` 监听了 `ctx.Done()`，这是生产级代码的标配。如果整个 HTTP 请求被用户 abort，业务 Context 会被 cancel，看门狗必须能够感知并安全退出，防止 Goroutine 泄漏。

---

## 演进四：Go 语言中的“可重入锁”悖论

**踩坑点**：方法 A 拿到了锁，内部调用方法 B，方法 B 也去拿同一把锁，导致自己把自己死锁。
**专家解法**：在 Java 中，可重入锁非常简单，因为 Java 有 ThreadLocal 和 ThreadID。但在 Go 中，**Goroutine 是没有暴露 ID 的**！

很多新手尝试用汇编去 Hack 获取 Goroutine ID，这在工程上是极度危险且不被官方推荐的。在 Go 中实现可重入锁，通常有两种高级姿势：

1. **Token 传递法**：将加锁返回的 `token` 传递给下游函数，下游函数凭此 token 证明自己是“同一调用链”。
2. **Context 传递法**（推荐）：将锁的标识存入 `context.Context`，在调用链中透传。

如果一定要在 Redis 层面做重入（参考 Redisson 做法），我们需要使用 Redis 的 **Hash** 结构。
- Key：锁的名称
- Field：`唯一的协程标识（可以用 uuid + 链路追踪 TraceID）`
- Value：重入次数

```lua
-- 可重入加锁 Lua 脚本
local key = KEYS[1]
local thread_id = ARGV[1]
local ttl = ARGV[2]

if (redis.call('exists', key) == 0) then
    redis.call('hset', key, thread_id, 1)
    redis.call('pexpire', key, ttl)
    return 1
end

if (redis.call('hexists', key, thread_id) == 1) then
    redis.call('hincrby', key, thread_id, 1)
    redis.call('pexpire', key, ttl)
    return 1
end

return 0
```

---

## 演进五：终极拷问 —— 主从切换与 Fencing Token

上面所有的努力，在单机 Redis 上已经无懈可击了。但生产环境全是 Redis Cluster 或主从哨兵架构。

**终极坑点**：Master 节点挂掉，异步复制导致锁没同步到 Slave。新 Master 上任，锁丢失，导致两个协程同时拿到锁！

很多人会说：“用 Redlock（红锁）啊！”

但分布式系统专家 Martin Kleppmann 曾撰文狠狠批判过 Redlock。因为 Redlock 严重依赖**服务器时钟同步**。如果某台 Redis 服务器发生时钟跳跃（Clock Jump），或者进程发生了长时间的 GC Pause，Redlock 依然会崩溃。

**终极专家的破局之道**：
对于要求**绝对强一致性**的金融级场景，专家的选择是：
1. **换掉 Redis**：使用基于 Raft 协议、CP 模型的 `etcd` 或 `ZooKeeper`。`etcd` 的 Lease（租约）和 Watch 机制天生就是为分布式协调而生的。
2. **引入 Fencing Token（击剑令牌）**：
   无论锁多么完美，我们都不信任它。我们在获取锁的同时，从发号器获取一个**单调递增的 Token**。
   当协程带着数据去操作数据库（或下游系统）时，必须带上这个 Token。下游系统通过乐观锁或唯一索引拒绝低版本 Token 的请求。

*(Martin Kleppmann 提出的 Fencing Token 机制，是解决分布式锁超时/失效的最终防御底线)*

---

## 总结与选型建议

从写下一行 `SetNX` 到构建一个生产级的分布式锁，体现的是一个 Go 开发者对系统边界的敬畏：

1. **基础**：原子操作（SetNX EX）与 Lua 脚本是底线。
2. **健壮**：用 `context` 控制超时，用 Watchdog 防止提前释放，坚决杜绝 Goroutine 泄漏。
3. **架构**：明白 Go 并发模型的特殊性（无 GID），通过 Context 传递上下文实现重入。
4. **视野**：跳出 Redis 的局限，理解 AP 模型与 CP 模型的差异。

**技术选型建议**：
- **常规业务**（如限制用户操作频率、普通电商扣库存）：用 Redis 锁足够，建议直接使用开源库 `go-redsync/redsync`。
- **核心资金链路**（如账务结算）：请老老实实上 `etcd`，并配合数据库的乐观锁（版本号机制）进行兜底。

如果你是使用 PHP 进行开发项目的话，那么也可以直接使用 `https://github.com/pudongping/wise-locksmith` 库，实现了分布式锁和红锁。如果你是使用 `Hyperf` 框架的话，还可以直接使用 `https://github.com/pudongping/hyperf-wise-locksmith` 库，作者都做了特定的兼容。

希望这篇文章能够对你有所帮助～