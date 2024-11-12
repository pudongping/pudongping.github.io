---
title: 如何用Redis高效实现点赞功能？用Set？还是Bitmap？
author: Alex
top: false
hide: false
cover: false
toc: true
mathjax: false
abbrlink: 3adceb1d
date: 2024-11-12 10:51:06
img:
coverImg:
password:
summary:
categories: Redis
tags:
  - Redis
  - Cache
---

在众多软件应用中，**点赞**功能几乎成了所有应用中的“标配”。但实现一个高效的点赞功能并不简单，尤其是在面对大规模的用户量和高并发场景时。

今天，我们就从实际需求出发，探索如何利用 Redis 的数据结构来设计一个点赞系统，从而理解 `Set` 和 `Bitmap` 数据结构的优缺点。

## 需求分析

我们设定这样一个需求场景：在一篇文章的评论下实现点赞功能，每位用户只能对同一条评论点赞一次，再次点赞则视为“取消点赞”。此外，我们还需要统计每条评论的总点赞数。

这听上去不复杂，但当需求量级提升，比如用户量达到千万级别，系统会需要承担巨大的数据存储和高频的读写压力。为了满足高性能和低延迟的需求，我们可以选择 Redis 来管理点赞数据。

## 方案一：使用 Redis 的 `Set` 数据结构

首先考虑使用 Redis 的 `Set` 数据结构。`Set` 非常适合存储一组不重复的元素，因此可以用来记录每条评论下点赞的用户 ID。方案设计如下：

1. **存储设计**：每条评论的点赞数据可以存储为一个 Redis `Set`，键格式为 `comment:likes:{comment_id}`，其中 `{comment_id}` 是评论的唯一标识。
2. **点赞操作**：当用户点赞时，将用户 ID 添加到该评论对应的 `Set` 中。
3. **取消点赞**：如果用户再次点击，则从 `Set` 中移除用户 ID。
4. **统计总点赞数**：直接获取 `Set` 的元素数量，即为当前评论的点赞总数。
5. **检查用户是否点赞过**：可以通过 `SISMEMBER` 指令快速检查某个用户 ID 是否存在于该评论的点赞 `Set` 中。

### Set 方案的代码实现

以下是 PHP 实现点赞、取消点赞和统计点赞数的代码示例：

```php
// 点赞接口
function likeComment($userId, $commentId) {
    $redis = new Redis();
    $redis->connect('127.0.0.1', 6379);

    // 生成唯一键
    $likeSetKey = "comment:likes:{$commentId}";

    // 判断用户是否已点赞
    if ($redis->sIsMember($likeSetKey, $userId)) {
        // 已点赞，取消点赞
        $redis->sRem($likeSetKey, $userId);
        $liked = false;
    } else {
        // 未点赞，添加点赞
        $redis->sAdd($likeSetKey, $userId);
        $liked = true;
    }

    // 获取当前总点赞数
    $totalLikes = $redis->sCard($likeSetKey);

    // 返回点赞状态
    return [
        'status' => 'success',
        'liked' => $liked,
        'totalLikes' => $totalLikes
    ];
}
```

### 使用 Set 方案的优缺点

**优点**：
- **灵活性高**：支持不连续的用户 ID，适合大多数应用场景。
- **操作简单**：Redis 原生支持集合操作，查询、添加、删除等操作性能较高。

**缺点**：
- **存储空间较大**：`Set` 中每个用户 ID 都会占用存储空间，随着点赞用户增多，`Set` 的存储开销也会增长。

## 方案二：使用 Redis 的 `Bitmap` 数据结构

如果用户 ID 是连续的，比如从 `0` 开始顺序增长，那么可以使用 Redis 的 `Bitmap` 数据结构来进一步提升存储效率。`Bitmap` 以每个用户的 ID 作为位（bit）位置，只需 1 位就可以表示每个用户的点赞状态，大大节省存储空间。

1. **存储设计**：每条评论的点赞数据可以存储为一个 `Bitmap`，键的格式为 `comment:likes:{comment_id}`。
2. **点赞操作**：使用 `SETBIT` 将用户的位设置为 `1`。
3. **取消点赞**：再次点击则用 `SETBIT` 将该位设置为 `0`。
4. **统计总点赞数**：通过 `BITCOUNT` 指令统计 `Bitmap` 中位为 `1` 的数量，即为点赞总数。
5. **检查用户是否点赞过**：可以用 `GETBIT` 查询指定用户的点赞状态。

### Bitmap 方案的代码实现

以下是使用 `Bitmap` 实现的 PHP 代码示例：

```php
function likeComment($userId, $commentId) {
    $redis = new Redis();
    $redis->connect('127.0.0.1', 6379);

    // 生成唯一键
    $likeBitmapKey = "comment:likes:{$commentId}";

    // 获取用户当前的点赞状态
    $isLiked = $redis->getBit($likeBitmapKey, $userId);

    if ($isLiked) {
        // 已点赞，取消点赞
        $redis->setBit($likeBitmapKey, $userId, 0);
        $liked = false;
    } else {
        // 未点赞，设置为点赞
        $redis->setBit($likeBitmapKey, $userId, 1);
        $liked = true;
    }

    // 获取当前点赞总数
    $totalLikes = $redis->bitCount($likeBitmapKey);

    return [
        'status' => 'success',
        'liked' => $liked,
        'totalLikes' => $totalLikes
    ];
}
```

### 使用 Bitmap 方案的优缺点

**优点**：
- **存储效率高**：每位只占 1 bit，适合大量用户数据的二元状态存储，极大节省内存。
- **适合批量统计**：通过 `BITCOUNT` 等命令可以快速统计点赞数量，性能极佳。

**缺点**：
- **用户 ID 需连续**：`Bitmap` 适合连续的 ID（如从 `0` 到某个上限），对于离散的 ID 或存在大量空位的 ID，不适用。
- **操作复杂性较高**：当用户 ID 离散或不连续时，使用 `Bitmap` 不仅不节省空间，操作复杂性也会增加。

## 选择合适的数据结构

| 特点                   | Redis `Set`                              | Redis `Bitmap`                      |
|------------------------|------------------------------------------|-------------------------------------|
| **用户 ID 分布**         | 适合不连续的 ID                          | 适合连续的、紧凑的 ID               |
| **存储空间**            | 随点赞数增长而增大                       | 每个用户点赞状态占 1 bit，空间占用小  |
| **统计点赞数**          | 通过 `SCARD` 精确统计                    | 通过 `BITCOUNT` 高效统计             |
| **查询用户状态**         | 支持任意用户 ID 查询点赞状态             | 支持按位查询连续 ID 的状态           |
| **适用场景**            | 离散用户 ID，点赞、关注等集合操作         | 连续 ID，批量二元状态的快速统计       |

## 总结

在实际项目中，选择合适的数据结构至关重要。对于点赞功能，如果用户 ID 是不连续的且规模不大，`Set` 更灵活、易于使用；而对于用户 ID 连续的大规模应用，`Bitmap` 则能极大提升存储和统计效率。在实际应用中，我们可以根据用户 ID 分布、存储需求和性能要求来选择最优方案。

其实 `Bitmap` 这种数据结构更加适合于用作**用户签到、打卡**等场景。