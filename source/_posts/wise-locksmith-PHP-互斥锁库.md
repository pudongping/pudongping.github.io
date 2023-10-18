---
title: wise-locksmith PHP 互斥锁库
author: Alex
top: false
hide: false
cover: true
toc: true
mathjax: false
categories: PHP
tags:
  - PHP
  - Mutex
  - Lock
  - 锁
  - 互斥锁
  - 红锁
abbrlink: 1e08f87d
date: 2023-10-18 15:36:06
img:
coverImg:
password:
summary: wise-locksmith 是一个不局限于框架的互斥锁库，用于在高并发场景下提供 PHP 代码的互斥执行。
---

# PHP 互斥锁库

[wise-locksmith](https://github.com/pudongping/wise-locksmith) 是一个不局限于框架的互斥锁库，用于在高并发场景下提供 PHP 代码的互斥执行。
如果你是使用 [hyperf](https://hyperf.wiki/) 框架，那么你可以直接使用 [hyperf-wise-locksmith](https://github.com/pudongping/hyperf-wise-locksmith) 适配库。


## 要求

- PHP >= 7.1 或以上版本
- Redis >= 2.6.12 或以上版本（如果需要使用到分布式锁或者红锁的情况下）
- Swoole >= 4.5 或以上版本 （如果需要使用协程级别的互斥锁的情况下）

## 安装

```shell
composer require pudongping/wise-locksmith
```

## 快速开始

```php
<?php

require 'vendor/autoload.php';

use Pudongping\WiseLocksmith\Locker;

$redisHosts = [
    [
        'host' => '127.0.0.1',
        'port' => 6379
    ],
    [
        'host' => '127.0.0.1',
        'port' => 6380
    ],
    [
        'host' => '127.0.0.1',
        'port' => 6381
    ],
    [
        'host' => '127.0.0.1',
        'port' => 6382
    ],
    [
        'host' => '127.0.0.1',
        'port' => 6383
    ],
];

// 如果需要使用到分布式锁或者红锁时，则需要初始化 redis 实例，否则可跳过这一步
$redisInstances = array_map(function ($v) {
    $redis = new \Redis();
    $redis->connect($v['host'], $v['port']);
    return $redis;
}, $redisHosts);

// 创建一个锁实例
$locker = new Locker();
```

### flock - 文件锁

文件锁没有任何依赖。可通过可选的第 3 个参数参数设置锁的超时时间，单位：秒。（支持浮点型，比如 1.5 表示 1500ms 也就是最多会等待 1500ms，如果没有抢占到锁，那么则主动放弃抢锁，同时会抛出 `Pudongping\WiseLocksmith\Exception\TimeoutException` 异常）
设置成 `Pudongping\WiseLocksmith\Lock\File\Flock::INFINITE_TIMEOUT` 时，表示永不过期，则当前一直会阻塞式抢占锁，直到抢占到锁为止。默认值为：`Pudongping\WiseLocksmith\Lock\File\Flock::INFINITE_TIMEOUT`。

```php

$path = tempnam(sys_get_temp_dir(), 'wise-locksmith-flock-');
$fileHandler = fopen($path, 'r');

$res = $locker->flock($fileHandler, function () {
    // 这里写你想保护的代码
});

unlink($path);

return $res;
```

### redisLock - 分布式锁

需要依赖 `redis` 扩展。可通过可选的第 3 个参数设置锁的超时时间，单位：秒。（支持浮点型，比如 1.5 表示 1500ms 也就是最多会等待 1500ms，如果没有抢占到锁，那么则主动放弃抢锁，同时会抛出 `Pudongping\WiseLocksmith\Exception\TimeoutException` 异常）
默认值为：`5`。第 4 个参数为当前锁的具有唯一性的值，除非有特殊情况下需要设置，一般不需要设置。

```php

$res = $locker->redisLock($redisInstances[0], 'redisLock', function () {
    // 这里写你想保护的代码
});

return $res;
```

### redLock - 红锁（redis 集群环境时，分布式锁的实现）

redLock 锁所需要设置的参数和 redisLock 锁除了第一个参数有区别以外，其他几个参数完全一致。redLock 锁是 redisLock 锁的集群实现。

```php

$res = $locker->redLock($redisInstances, 'redLock', function () {
    // 这里写你想保护的代码
});

return $res;
```

### channelLock - 协程级别的互斥锁

使用此锁时，需要安装 `swoole` 扩展。且版本必须大于等于 `4.5`。可通过可选的第 3 个参数设置锁的超时时间，单位：秒。（支持浮点型，比如 1.5 表示 1500ms 也就是最多会等待 1500ms，如果没有抢占到锁，那么则主动放弃抢锁，同时直接返回 `false` 表示没有抢占到锁）
设置成 `-1` 时，表示永不过期，则当前一直会阻塞式抢占锁，直到抢占到锁为止。默认值为：`-1`。

```php

$res = $locker->channelLock('channelLock', function () {
    // 这里写你想保护的代码
});

return $res;
```