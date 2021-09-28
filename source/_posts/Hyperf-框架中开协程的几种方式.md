---
title: Hyperf 框架中开协程的几种方式
author: Alex
top: false
hide: false
cover: false
toc: true
mathjax: false
categories: PHP
tags:
  - Hyperf
  - Swoole
abbrlink: b5fcf661
date: 2021-09-28 10:06:48
img:
coverImg:
password:
summary:
---

# Hyperf 框架中开协程的几种方式

```php

<?php
/**
 * 开协程做请求
 *
 * Created by PhpStorm
 * User: Alex
 * Date: 2021-09-21 19:22
 * E-mail: <276558492@qq.com>
 */
declare(strict_types=1);

namespace App\Controller;

use Hyperf\Utils\Parallel;
use Hyperf\Utils\WaitGroup;
use Swoole\Coroutine\Channel;
use Hyperf\Guzzle\ClientFactory;
use Hyperf\HttpServer\Annotation\AutoController;
use Hyperf\HttpServer\Contract\RequestInterface;
use Hyperf\Di\Annotation\Inject;

/**
 * @AutoController
 * Class CoController
 * @package App\Controller
 */
class CoController
{

    /**
     * @Inject
     * @var ClientFactory
     */
    private $clientFactory;

    /**
     * 方便测试超时操作
     *
     * @param RequestInterface $request
     * @return mixed
     */
    public function sleep(RequestInterface $request)
    {
        $seconds = $request->query('seconds', 1);
        sleep((int)$seconds);
        var_dump('sleep hello ====> ' . $seconds);
        return $seconds;
    }

    /**
     * 使用 Channel 做协程请求
     *
     * @return array
     * @throws \GuzzleHttp\Exception\GuzzleException
     */
    public function test()
    {
        $channel = new Channel();

        co(function () use ($channel) {
            $client = $this->clientFactory->create();
            $client->get('127.0.0.1:9516/co/sleep?seconds=2');
            $channel->push(123);
        });

        co(function () use ($channel) {
            $client = $this->clientFactory->create();
            $client->get('127.0.0.1:9516/co/sleep?seconds=2');
            $channel->push(321);
        });

        $result[] = $channel->pop();  // 第一个协程返回的结果
        $result[] = $channel->pop();  // 第二个协程返回的结果

        return $result;  // [123, 321]
    }

    /**
     * 使用 WaitGroup 做协程请求
     *
     * @return array
     * @throws \GuzzleHttp\Exception\GuzzleException
     */
    public function test1()
    {
        // 通过子协程并行的发起多个请求实现并行请求
        $wg = new WaitGroup();
        $wg->add(2);  // 因为开了两个协程，因此就要添加 2

        $result = [];
        co(function () use ($wg, &$result) {
            $client = $this->clientFactory->create();
            $client->get('127.0.0.1:9516/co/sleep?seconds=2');
            $result[] = 123;
            $wg->done();
        });

        co(function () use ($wg, &$result) {
            $client = $this->clientFactory->create();
            $client->get('127.0.0.1:9516/co/sleep?seconds=2');
            $result[] = 321;
            $wg->done();
        });

        $wg->wait();  // 等待 add 计数器变为 0

        return $result;  // [321, 123]
    }

    /**
     * 使用 Parallel 类，做协程请求
     *
     * @return array
     */
    public function test2()
    {
        $parallel = new Parallel();

        $parallel->add(function () {
            $client = $this->clientFactory->create();
            $client->get('127.0.0.1:9516/co/sleep?seconds=2');
            return 123;
        }, 'foo');

        $parallel->add(function () {
            $client = $this->clientFactory->create();
            $client->get('127.0.0.1:9516/co/sleep?seconds=2');
            return 321;
        }, 'bar');

        $result = $parallel->wait();

        return $result;  // {"foo":123,"bar":321}
    }

    /**
     * 使用 parallel 助手函数，做协程请求
     *
     * @return array
     */
    public function test3()
    {
        $result = parallel([
            'foo' => function () {
                $client = $this->clientFactory->create();
                $client->get('127.0.0.1:9516/co/sleep?seconds=2');
                return 123;
            },
            'bar' => function () {
                $client = $this->clientFactory->create();
                $client->get('127.0.0.1:9516/co/sleep?seconds=2');
                return 321;
            },
        ]);

        return $result;  // {"foo":123,"bar":321}
    }

}

```
