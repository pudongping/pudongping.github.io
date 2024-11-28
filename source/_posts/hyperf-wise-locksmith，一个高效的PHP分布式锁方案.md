---
title: hyperf-wise-locksmith，一个高效的PHP分布式锁方案
author: Alex
top: false
hide: false
cover: false
toc: true
mathjax: false
categories: PHP
tags:
  - PHP
  - Hyperf
  - 锁
  - 分布式锁
abbrlink: a3ff01d2
date: 2024-11-28 10:24:57
img:
coverImg:
password:
summary:
---

在分布式系统中，如何确保多台机器之间不会产生竞争条件，是一个常见且重要的问题。`hyperf-wise-locksmith` 库作为 `Hyperf` 框架中的一员，提供了一个高效、简洁的互斥锁解决方案。

本文将带你了解这个库的安装、特性、基本与高级功能，并结合实际应用场景，展示其在项目中的应用。

## hyperf-wise-locksmith 库简介

`hyperf-wise-locksmith` 是一个适配 Hyperf 框架的互斥锁库，它基于 `pudongping/wise-locksmith` 库构建。它可以帮助我们在分布式环境下进行锁的管理，确保同一时刻只有一个进程能够操作某些共享资源，从而避免数据的竞争和不一致问题。

## 安装

要在你的 Hyperf 项目中使用 `hyperf-wise-locksmith`，你需要通过 Composer 进行安装：

```shell
composer require pudongping/hyperf-wise-locksmith -vvv
```

确保你的环境满足以下要求：
- PHP >= 8.0
- hyperf ~3.0.0。

## 特性

`hyperf-wise-locksmith` 提供了多种锁机制，包括**文件锁**、**分布式锁**、**红锁**和**协程级别的互斥锁**。这些锁机制可以帮助开发者在不同的场景下保护共享资源，避免竞态条件。

## 基本功能

### 文件锁（flock）

文件锁是一种简单的锁机制，它依赖于文件系统。以下是一个使用文件锁的示例：

```php
    private function flock(float $amount)
    {
        $path = BASE_PATH . '/runtime/alex.lock.cache';
        $fileHandler = fopen($path, 'a+');
        // fwrite($fileHandler, sprintf("%s - %s \r\n", 'Locked', microtime()));

        $res = $this->locker->flock($fileHandler, function () use ($amount) {
            return $this->deductBalance($amount);
        });
        return $res;
    }
```

### 分布式锁（redisLock）

分布式锁适用于分布式系统，它依赖于 Redis。以下是一个使用分布式锁的示例：

```php
    private function redisLock(float $amount)
    {
        $res = $this->locker->redisLock('redisLock', function () use ($amount) {
            return $this->deductBalance($amount);
        }, 10);
        return $res;
    }
```

## 高级功能

### 红锁（redLock）

红锁是一种更安全的分布式锁实现，它需要多个 Redis 实例。以下是一个使用红锁的示例：

```php
    private function redLock(float $amount)
    {
        $res = $this->locker->redLock('redLock', function () use ($amount) {
            return $this->deductBalance($amount);
        }, 10);
        return $res;
    }
```

### 协程级别的互斥锁（channelLock）

协程级别的互斥锁适用于协程环境，它提供了一种轻量级的锁机制。以下是一个使用协程锁的示例：

```php
    private function channelLock(float $amount)
    {
        $res = $this->locker->channelLock('channelLock', function () use ($amount) {
            return $this->deductBalance($amount);
        });
        return $res;
    }
```

## 实际应用场景

假设我们有一个在线支付系统，需要在多个请求中扣减用户的余额。如果不使用互斥锁，可能会导致超扣或扣减失败。使用 `hyperf-wise-locksmith` 库，我们可以确保每次扣减操作都是原子性的。

### 代码示例

以下是一个扣减用户余额的示例，使用了 `hyperf-wise-locksmith` 库：

```php
<?php
declare(strict_types=1);

namespace App\Services;

use Hyperf\Contract\StdoutLoggerInterface;
use Pudongping\HyperfWiseLocksmith\Locker;
use Pudongping\WiseLocksmith\Exception\WiseLocksmithException;
use Pudongping\WiseLocksmith\Support\Swoole\SwooleEngine;
use Throwable;

class AccountBalanceService
{

    /**
     * 用户账户初始余额
     *
     * @var float|int
     */
    private float|int $balance = 10;

    public function __construct(
        private StdoutLoggerInterface $logger,
        private Locker                $locker
    ) {
        $this->locker->setLogger($logger);
    }

    private function deductBalance(float|int $amount)
    {
        if ($this->balance >= $amount) {
            // 模拟业务处理耗时
            usleep(500 * 1000);
            $this->balance -= $amount;
        }

        return $this->balance;
    }

    /**
     * @return float
     */
    private function getBalance(): float
    {
        return $this->balance;
    }

    public function runLock(int $i, string $type, float $amount)
    {
        try {
            $start = microtime(true);

            switch ($type) {
                case 'flock':
                    $this->flock($amount);
                    break;
                case 'redisLock':
                    $this->redisLock($amount);
                    break;
                case 'redLock':
                    $this->redLock($amount);
                    break;
                case 'channelLock':
                    $this->channelLock($amount);
                    break;
                case 'noMutex':
                default:
                    $this->deductBalance($amount);
                    break;
            }

            $balance = $this->getBalance();
            $id = SwooleEngine::id();
            $cost = microtime(true) - $start;
            $this->logger->notice('[{type} {cost}] ==> [{i}<=>{id}] ==> 当前用户的余额为：{balance}', compact('type', 'i', 'balance', 'id', 'cost'));

            return $balance;
        } catch (WiseLocksmithException|Throwable $e) {
            return sprintf('Err Msg: %s ====> %s', $e, $e->getPrevious());
        }
    }

}
```

然后我们再写一个控制器进行调用

```php
<?php

declare(strict_types=1);

namespace App\Controller;

use Hyperf\HttpServer\Annotation\AutoController;
use App\Services\AccountBalanceService;
use Hyperf\Coroutine\Parallel;
use function \Hyperf\Support\make;

#[AutoController]
class BalanceController extends AbstractController
{

    // curl '127.0.0.1:9511/balance/consumer?type=noMutex'
    public function consumer()
    {
        $type = $this->request->input('type', 'noMutex');
        $amount = (float)$this->request->input('amount', 1);

        $parallel = new Parallel();
        $balance = make(AccountBalanceService::class);

        // 模拟 20 个并发
        for ($i = 1; $i <= 20; $i++) {
            $parallel->add(function () use ($balance, $i, $type, $amount) {
                return $balance->runLock($i, $type, $amount);
            }, $i);
        }

        $result = $parallel->wait();

        return $this->response->json($result);
    }

}
```

当我们访问 `/balance/consumer?type=noMutex` 地址时，我们可以看到用户的余额会被扣成负数，这明显不符合逻辑。
然而当我们访问下面几个地址时，我们可以看到用户余额不会被扣成负数，则说明很好的保护了竞态下的共享资源的准确性。

- `/balance/consumer?type=flock` ：文件锁
- `/balance/consumer?type=redisLock` ：分布式锁
- `/balance/consumer?type=redLock` ：红锁
- `/balance/consumer?type=channelLock` ：协程级别的互斥锁

## 注意

关于使用到 `redisLock` 和 `redLock` 时：

- 使用 `redisLock` 默认采用的 `config/autoload/redis.php` 配置文件中的第一个 `key` 配置 redis 实例（即 **default**）。可按需传入第 4 个参数 `string|null $redisPoolName` 进行重新指定。
- 使用 `redLock` 默认采用的 `config/autoload/redis.php` 配置文件中的所有 `key` 对应的配置 redis 实例。可按需传入第 4 个参数 `?array $redisPoolNames = null` 进行重新指定。

## 文档

详细文档可见 [pudongping/wise-locksmith](https://github.com/pudongping/wise-locksmith)。

## 结语

`hyperf-wise-locksmith` 库为 Hyperf 框架的开发者提供了强大的互斥锁功能，可以帮助我们在高并发场景下保护共享资源。

通过本文的介绍，希望你能对 `hyperf-wise-locksmith` 有一个全面的了解，并在你的项目中灵活运用。如果你觉得这个库对你有帮助，希望你可以帮忙点个 Star 哟～
