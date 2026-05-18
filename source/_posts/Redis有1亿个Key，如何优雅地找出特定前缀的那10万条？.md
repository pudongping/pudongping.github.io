---
title: Redis有1亿个Key，如何优雅地找出特定前缀的那10万条？
author: Alex
top: false
hide: false
cover: false
toc: true
mathjax: false
date: 2026-05-18 10:58:31
img:
coverImg:
password:
summary:
categories: Redis
tags:
  - Cache
  - Redis
  - 缓存
---

Redis 是每一位后端工程师最熟悉的中间件之一。
它轻量、高效、灵活，几乎是现代高并发系统中最常见的缓存组件。

可当业务规模上来之后，Redis 的简单世界也不再简单。
有一天你可能突然接到一个看似普通的需求：

> 「我们的 Redis 里大概有 1 亿个 Key，现在想把前缀是 `user:profile:` 的那 10 万条数据找出来。」

看似小事，但稍有不慎，就可能让整个 Redis 实例**直接假死**。

---

## 从“最直觉”的做法说起

很多人第一反应是：

```bash
KEYS user:profile:*
```

对吧？
这行命令在开发环境里跑得飞快，一眨眼结果就出来了。

但如果你在线上执行它，那就等着运维同事来找你“喝茶”吧。

### ❌ 为什么 KEYS 命令不能在线上用？

Redis 是单线程模型，`KEYS` 命令会遍历整个键空间（无论有多少 Key），
在执行期间，它**会完全阻塞主线程**。

这意味着：

* 所有请求（读/写/过期）都会卡住；
* CPU 飙升；
* 响应时间暴涨；
* 严重点可能直接被监控系统判定为宕机重启。

一句话总结：

> `KEYS` 命令在开发环境没问题，但**线上执行就是一颗“核弹”**。

---

## Redis 官方给出的“安全方案”：SCAN

自 2.8 版本起，Redis 提供了一个安全替代命令 —— `SCAN`。
它的思路是：**渐进式遍历**。

```bash
SCAN 0 MATCH user:profile:* COUNT 1000
```

这个命令会做两件事：

1. 返回部分匹配到的 Key；
2. 返回一个游标（cursor），下次从这个游标继续扫描。

当游标返回 `0` 时，表示扫描结束。

也就是说，我们可以用一个循环，不断调用 `SCAN`，直到把所有匹配 Key 都找完。

### ✅ 优点

* 非阻塞，不会影响主线程；
* 支持模式匹配；
* 可分批控制速率；
* 可以边扫描边处理结果。

### ⚠️ 缺点

* 扫描是“近似随机”的，不保证顺序；
* 可能会出现重复 Key（需自行去重）；
* 全库扫描依然耗时较长。

---

## 代码实战：安全地找出目标 Key

理论说再多，还不如撸代码来的直接。下面是一段可直接使用的示例代码（使用 `redis-py`）👇

```python
import redis

# 连接 Redis
r = redis.Redis(host='127.0.0.1', port=6379, decode_responses=True)

cursor = 0
matched_keys = set()  # 使用 set 自动去重

while True:
    # 进行渐进式扫描，每次最多扫描 10000 条
    cursor, keys = r.scan(cursor=cursor, match='user:profile:*', count=10000)
    
    # 累积结果
    matched_keys.update(keys)
    
    # 每扫描一批可以打印进度
    print(f"Scanned batch, total found: {len(matched_keys)} keys")
    
    # 游标返回 0 时表示扫描完成
    if cursor == 0:
        break

print(f"\n✅ 扫描完成，共找到 {len(matched_keys)} 个匹配 key")
```

### 💬 运行结果示例：

```bash
Scanned batch, total found: 3200 keys
Scanned batch, total found: 7580 keys
...
✅ 扫描完成，共找到 100000 个匹配 key
```

**关键点解释：**

* `scan()` 是安全的，不会阻塞；
* `count` 可以调节扫描速率；
* 用 `set()` 去重，防止重复；
* 可在循环中加入 `sleep()` 控制速率，防止 CPU 打满。

---

## SCAN 的原理浅析

理解 `SCAN` 的原理，有助于你判断什么时候它会慢。

* Redis 内部的 key 空间是一个哈希表；
* `SCAN` 通过游标分段遍历哈希槽；
* 每次返回的数量不固定；
* 匹配是**边扫描边匹配**，并非提前过滤。

因此：

* key 总量越大，`SCAN` 遍历越久；
* `COUNT` 值越大，单次响应时间越长；
* 若 Redis 实例被大量写入，游标扫描的视图可能略有“漂移”。

这就是为什么说 `SCAN` 是**近似一致性**遍历，不是完全快照。

---

## 架构层优化思路

如果你频繁需要按前缀查找Key，
那其实问题并不在“怎么找”，而在于“为什么要找”。

这说明数据建模可能需要优化。下面是几种更优雅的方案👇

---

### 1️⃣ 为特定前缀维护一个索引集合（Set）

当你写入数据时，同时维护一个索引集合：

```bash
SADD user:profile:index 1001
SADD user:profile:index 1002
```

这样，将来只需：

```bash
SMEMBERS user:profile:index
```

瞬间就能拿到所有对应 Key，无需遍历全库。

#### 💡 优势：

* 查询O(1)，性能极佳；
* 非阻塞；
* 可实现分页（ZSET）。

#### ⚠️ 注意：

* 需要写入时同步维护索引；
* 索引可能出现脏数据，需定期校验。

---

### 2️⃣ 分库 / 分片存储

如果某类 Key 数量特别大，可以按前缀拆分：

* 不同业务前缀放在不同 Redis 实例；
* 或者在 Cluster 模式下通过**哈希标签**定位（如 `user:{profile}:id`）。

这样每次扫描的范围更小，效率提升明显。

## 写在最后

很多人对 Redis 的理解，停留在“它很快”。
但在真正的生产环境里，Redis 的快，更像是一把双刃剑。

它让你轻松读写上亿数据，却也可能因为一个命令——
让整个系统陷入停顿。

所以，当你下次再想用 `KEYS *` 时，请先想一想：

> 你是在开发环境，还是在生产线上？

成熟的工程师，不是写出最短的命令，而是设计出最稳的系统。
