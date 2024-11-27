---
title: PHP 互斥锁：如何确保代码的线程安全？
author: Alex
top: false
hide: false
cover: false
toc: true
mathjax: false
categories: PHP
tags:
  - PHP
  - 锁
  - 互斥锁
abbrlink: f387a51d
date: 2024-11-27 10:05:29
img:
coverImg:
password:
summary:
---

在多线程和高并发的环境下，我们经常会遇到需要确保代码段互斥执行的场景。比如，在电商平台中，当多个用户同时购买同一件商品时，如何确保库存的扣减是线程安全的？

今天，我们将一起探讨这个问题，并介绍一个名为 `wise-locksmith` 的 PHP 互斥锁库，它可以帮助我们轻松地解决这类问题。

## 代码的线程安全

在没有互斥机制的情况下，多个进程或线程可能会同时修改同一个资源，导致数据不一致的问题。例如，在一个简单的库存扣减操作中：

```php
// 假设库存为 10
$stock = 10;

// 多个请求同时到达，每个请求都扣减库存
for ($i = 0; $i < 20; $i++) {
    $stock--;
}
// 最终库存可能不是我们预期的 0，而是负数
```

这种情况在实际开发中是不可接受的。那么，我们如何确保在 PHP 中实现代码的互斥执行呢？

## `wise-locksmith` 库介绍

`wise-locksmith` 是一个 PHP 互斥锁库，它提供了**多种锁**机制来帮助我们解决线程安全问题。并且这个库**不局限于任何框架**，也就是说**只要是在 PHP 环境中，都可以使用**。

下面，我们将详细介绍这个库的安装、特性、基本与高级功能，并结合实际应用场景展示其在项目中的使用。来，继续往下看吧～

## 安装

首先，我们通过 Composer 快速安装 `wise-locksmith`：

```shell
composer require pudongping/wise-locksmith
```

## 特性

`wise-locksmith` 提供了多种锁机制，以适应不同的应用场景：

1. **文件锁（flock）**：适用于单服务器环境。
2. **分布式锁（redisLock）**：适用于需要跨多个服务器或实例的分布式环境。
3. **红锁（redLock）**：适用于 Redis 集群环境，提供更高的可靠性。
4. **协程级别的互斥锁（channelLock）**：适用于 Swoole 协程环境。

## 基本功能

### 文件锁（flock）


文件锁没有任何依赖。可通过可选的第 3 个参数参数设置锁的超时时间，单位：秒。（支持浮点型，比如 1.5 表示 1500ms 也就是最多会等待 1500ms，如果没有抢占到锁，那么则主动放弃抢锁，同时会抛出 `Pudongping\WiseLocksmith\Exception\TimeoutException` 异常）
设置成 `Pudongping\WiseLocksmith\Lock\File\Flock::INFINITE_TIMEOUT` 时，表示永不过期，则当前一直会阻塞式抢占锁，直到抢占到锁为止。默认值为：`Pudongping\WiseLocksmith\Lock\File\Flock::INFINITE_TIMEOUT`。

文件锁是最简单的一种锁，适用于单服务器环境。它通过锁定一个文件来实现互斥。以下是一个简单的文件锁示例：

```php
<?php
require 'vendor/autoload.php';
use Pudongping\WiseLocksmith\Locker;

$path = tempnam(sys_get_temp_dir(), 'wise-locksmith-flock-');
$fileHandler = fopen($path, 'r+');
$locker = new Locker();
try {
    $locker->flock($fileHandler, function () use ($stock) {
        // 这里写你想保护的代码
        $stock--;
        // 确保操作的原子性
    });
} catch (\Exception $e) {
    // 处理异常
}
fclose($fileHandler);
unlink($path);
```

### 分布式锁（redisLock）

需要依赖 `redis` 扩展。可通过可选的第 3 个参数设置锁的超时时间，单位：秒。（支持浮点型，比如 1.5 表示 1500ms 也就是最多会等待 1500ms，如果没有抢占到锁，那么则主动放弃抢锁，同时会抛出 `Pudongping\WiseLocksmith\Exception\TimeoutException` 异常）
默认值为：`5`。第 4 个参数为当前锁的具有唯一性的值，除非有特殊情况下需要设置，一般不需要设置。

在分布式系统中，我们经常需要确保跨多个服务器的操作是互斥的。`redisLock` 提供了这样的功能：

```php
<?php
$redis = new \Redis();
$redis->connect('127.0.0.1', 6379);
$locker = new Locker();
try {
    $locker->redisLock($redis, 'redisLock', function () use ($stock) {
        // 这里写你想保护的代码
        $stock--;
        // 确保操作的原子性
    });
} catch (\Exception $e) {
    // 处理异常
}
```

## 高级功能

### 红锁（redLock）

redLock 锁所需要设置的参数和 redisLock 锁除了第一个参数有区别以外，其他几个参数完全一致。redLock 锁是 redisLock 锁的集群实现。

红锁是分布式锁的一种高级实现，它在 Redis 集群环境中提供更高的可靠性：

```php
<?php
$redisInstances = [
    ['host' => '127.0.0.1', 'port' => 6379],
    // 其他 Redis 实例...
];
$redis = array_map(function ($v) {
    $redis = new \Redis();
    $redis->connect($v['host'], $v['port']);
    return $redis;
}, $redisInstances);
$locker = new Locker();
try {
    $locker->redLock($redis, 'redLock', function () use ($stock) {
        // 这里写你想保护的代码
        $stock--;
        // 确保操作的原子性
    });
} catch (\Exception $e) {
    // 处理异常
}
```

### 协程级别的互斥锁（channelLock）

使用此锁时，需要安装 `swoole` 扩展。且版本必须大于等于 `4.5`。可通过可选的第 3 个参数设置锁的超时时间，单位：秒。（支持浮点型，比如 1.5 表示 1500ms 也就是最多会等待 1500ms，如果没有抢占到锁，那么则主动放弃抢锁，同时直接返回 `false` 表示没有抢占到锁）
设置成 `-1` 时，表示永不过期，则当前一直会阻塞式抢占锁，直到抢占到锁为止。默认值为：`-1`。

在 Swoole 协程环境中，`channelLock` 提供了协程级别的互斥锁：

```php
<?php
$locker = new Locker();
try {
    $locker->channelLock('channelLock', function () use ($stock) {
        // 这里写你想保护的代码
        $stock--;
        // 确保操作的原子性
    });
} catch (\Exception $e) {
    // 处理异常
}
```

## 实际应用场景

假设我们有一个高并发的电商平台，需要在用户下单时扣减库存。使用 `wise-locksmith` 库，我们可以确保在任何时候只有一个请求能够修改库存，从而避免超卖的问题。以下是如何在实际项目中使用 `wise-locksmith` 来实现库存扣减的互斥操作：

```php
<?php
// 假设我们有一个全局的 Redis 连接实例
$redis = new \Redis();
$redis->connect('127.0.0.1', 6379);

// 库存扣减操作
function decreaseStock($productId, $quantity) {
    $locker = new Locker();
    try {
        $locker->redisLock($redis, "stock_lock_{$productId}", function () use ($productId, $quantity) {
            // 这里写你想保护的代码
            // 假设我们从数据库中获取当前库存
            $stock = getStockFromDatabase($productId);
            if ($stock >= $quantity) {
                // 更新库存
                updateStockInDatabase($productId, $stock - $quantity);
            }
        });
    } catch (\Exception $e) {
        // 处理异常
    }
}

// 调用扣减库存函数
decreaseStock(123, 1);
```

## 结语

通过 `wise-locksmith` 库，我们可以轻松地在 PHP 应用中实现代码的互斥执行，无论是单服务器环境还是分布式系统。

希望这篇文章能帮助你更好地理解和使用 `wise-locksmith` 库，确保你的代码在多线程环境下的线程安全。如果你觉得这个库对你有点儿帮助，那就请帮忙点个 Star 呀～
