---
title: 如何借助Redis更高效统计UV？——Hyperloglog篇
author: Alex
top: false
hide: false
cover: false
toc: true
mathjax: false
categories: Redis
tags:
  - Cache
  - Redis
  - 缓存
abbrlink: ceb09f0b
date: 2024-07-18 18:22:51
img:
coverImg:
password:
summary:
---

在今天的互联网时代，数据如潮水般汹涌而来。从用户行为数据、系统日志到实时交互数据，如何高效、准确地统计这海量数据中的唯一元素数量，成为了一个不小的挑战。

今天，我们要一起探索的是 Redis 中一个非常强大但可能被忽视的数据类型——HyperLogLog，它如何在牺牲极少的准确度前提下，实现对大规模数据集的快速去重计数。

## 什么是 HyperLogLog？

HyperLogLog 是一种用于基数统计的算法，基数指的是**一个集合中不重复元素的数量**。

想象一下，当我们面对数亿级别的数据时，传统的去重统计方法不仅计算量大，而且消耗大量的存储空间。HyperLogLog 则以一种非常节省空间的方式，解决了这个问题，虽然它的计算结果是一个估计值，但是准确率非常高，通常**误差率仅为 0.81%**。

## 使用 Redis 的 HyperLogLog 实现去重计数

### 添加元素：pfadd

首先，我们看一下如何向 HyperLogLog 中添加元素。Redis 提供了 `pfadd` 命令用于添加元素到 HyperLogLog 中。

```bash
127.0.0.1:6379> pfadd p1 a b c d e f g
(integer) 1
```

这里，我们添加了七个元素（a, b, c, d, e, f, g）到 HyperLogLog p1 中。返回值 `(integer) 1` 表示操作成功。

### 统计元素数量：pfcount

接下来，我们利用 `pfcount` 命令来获取 HyperLogLog 中的元素个数。

```bash
127.0.0.1:6379> pfcount p1
(integer) 7
```

通过 `pfcount p1`，我们得知 HyperLogLog p1 中有七个唯一元素，和我们之前添加的元素数量一致。

### 合并 HyperLogLog：pfmerge

如果我们有多个 HyperLogLog，想要合并它们的统计结果，该怎么做呢？Redis 的 `pfmerge` 命令能够帮助我们实现这一点。

```bash
127.0.0.1:6379> pfadd p2 d e f g h i j k
(integer) 1
127.0.0.1:6379> pfcount p2
(integer) 8

127.0.0.1:6379> pfmerge p3 p1 p2
OK

127.0.0.1:6379> pfcount p3
(integer) 11
```

从以上例子中，我们可以看到 `pfadd` 命令分别向 p1 和 p2 中添加了不同的元素。通过 `pfmerge p3 p1 p2`，我们将 p1 和 p2 合并到了 p3 中，并用 `pfcount p3` 确认了合并后 HyperLogLog p3 中共有 11 个唯一元素。

### 添加相同元素时

值得一提的是，如果我们尝试再次添加相同的元素到 HyperLogLog，它将不会增加计数，因为 HyperLogLog 本质上是一种去重计数工具。

刚刚我们往 `p1` 中添加了 (a, b, c, d, e, f, g) 七个元素，那么如果此时我们添加几个相同的元素以及少量的不同元素，会怎样呢？

```bash
127.0.0.1:6379> pfadd p1 a a b c h
(integer) 1
127.0.0.1:6379> pfcount p1
(integer) 8
127.0.0.1:6379>
```

我们可以看到，`p1` 只会统计不同的元素的个数，会自动过滤掉相同的元素。依据这个特性，是不是在脑海中就里面想到了一个常见的应用场景了？是的，依据这个去重特性，我们可以非常方便的做 UV 统计。

> UV（Unique visitor）：是指通过互联网访问、浏览这个网页的自然人。访问的一个电脑客户端为一个访客，一天内同一个访客仅被计算一次。

传统的做法是使用 set 保存用户的 ID，然后统计 set 中元素的数量作为判断标准。但是这种方式保存了大量的用户 ID，用户 ID 一般比较长，这就占用空间，还很麻烦。我们的目的是计数，不是保存数据，所以这样做有弊端。但是如果使用 hyperloglog 就比较合适了。

hyperloglog 的优点是**占用内存小**，并且是**固定的**。存储 2^64 个不同元素的基数，只需要 12 KB 的空间。

### 数据类型：type

最后，我们来确认一下 HyperLogLog 的数据类型。

```bash
127.0.0.1:6379> type p1
string
```

虽然 `type p1` 返回的是 `string`，但在 Redis 中，HyperLogLog 是作为一种概率型数据结构实现的，它通过一种特殊的字符串格式来存储数据。因此，虽然它在 Redis 中表现为字符串类型，但它用于实现基数统计的功能。

## 小结

HyperLogLog 提供了一种非常高效的方式来对大规模数据集进行去重计数。虽然其结果是估计值，但其高效性和准确度使其在处理大数据统计时表现出色。

通过上述的简单示例，相信你已经对 Redis 的 HyperLogLog 有了基本的了解。无论是实时数据分析、日志统计还是用户行为分析，HyperLogLog 都是一个值得尝试的利器。