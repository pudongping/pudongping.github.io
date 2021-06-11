---
title: hyperf 框架常用的助手函数
author: Alex
top: false
hide: false
cover: false
toc: true
mathjax: false
categories: PHP
tags:
  - Hyperf
  - PHP
  - Swoole
abbrlink: 79e7b1a4
date: 2021-06-09 10:07:33
img:
coverImg:
password:
summary: 以下列出部分 hyperf 常用的助手函数，还是比较实用的，仅供各位大佬参考。其中要是想获取当前请求的路由所对应的控制器和方法，请直接调用 `get_current_action()` 方法
---

以下列出部分 hyperf 常用的助手函数

> 其中要是想获取当前请求的路由所对应的控制器和方法，请直接调用 `get_current_action()` 方法

```php

<?php

declare(strict_types=1);

/**
 * 自定义助手函数
 */

use Hyperf\Utils\ApplicationContext;
use Hyperf\Redis\Redis;
use Hyperf\Contract\StdoutLoggerInterface;
use Hyperf\Logger\LoggerFactory;
use Hyperf\HttpServer\Contract\RequestInterface;
use Hyperf\HttpServer\Contract\ResponseInterface;
use Hyperf\HttpServer\Router\Dispatched;
use Hyperf\Server\ServerFactory;
use Swoole\Websocket\Frame;
use Swoole\WebSocket\Server as WebSocketServer;
use Psr\SimpleCache\CacheInterface;

if (! function_exists('container')) {
    /**
     * 获取容器对象
     *
     * @return \Psr\Container\ContainerInterface
     */
    function container()
    {
        return ApplicationContext::getContainer();
    }
}

if (! function_exists('redis')) {
    /**
     *  获取 Redis 协程客户端
     *
     * @return Redis|mixed
     */
    function redis()
    {
        return container()->get(Redis::class);
    }
}

if (! function_exists('std_out_log')) {
    /**
     * 控制台日志
     *
     * @return StdoutLoggerInterface|mixed
     */
    function std_out_log()
    {
        return container()->get(StdoutLoggerInterface::class);
    }
}

if (! function_exists('logger')) {
    /**
     * 文件日志
     *
     * @return \Psr\Log\LoggerInterface
     */
    function logger()
    {
        return container()->get(LoggerFactory::class)->make();
    }
}

if (! function_exists('request')) {
    /**
     * request 实例
     *
     * @return RequestInterface|mixed
     */
    function request()
    {
        return container()->get(RequestInterface::class);
    }
}

if (! function_exists('response')) {
    /**
     * response 实例
     *
     * @return ResponseInterface|mixed
     */
    function response()
    {
        return container()->get(ResponseInterface::class);
    }
}

if (! function_exists('server')) {
    /**
     * 基于 swoole server 的 server 实例
     *
     * @return \Swoole\Coroutine\Server|\Swoole\Server
     */
    function server()
    {
        return container()->get(ServerFactory::class)->getServer()->getServer();
    }
}

if (! function_exists('frame')) {
    /**
     * websocket frame 实例
     *
     * @return mixed|Frame
     */
    function frame()
    {
        return container()->get(Frame::class);
    }
}

if (! function_exists('websocket')) {
    /**
     * websocket 实例
     *
     * @return mixed|WebSocketServer
     */
    function websocket()
    {
        return container()->get(WebSocketServer::class);
    }
}

if (! function_exists('cache')) {
    /**
     * 简单的缓存实例
     *
     * @return mixed|CacheInterface
     */
    function cache()
    {
        return container()->get(CacheInterface::class);
    }
}

if (! function_exists('get_current_action')) {
    /**
     * 获取当前请求的控制器和方法
     *
     * @return array
     */
    function get_current_action() : array
    {
        $obj = request()->getAttribute(Dispatched::class);

        if (property_exists($obj, 'handler')
            && isset($obj->handler)
            && property_exists($obj->handler, 'callback')
        ) {
            $action = $obj->handler->callback;
        } else {
            throw new \Exception('The route is undifined! Please check!');
        }

        $errMsg = 'The controller and method are not found! Please check!';
        if (is_array($action)) {
            list($controller, $method) = $action;
        } elseif (is_string($action)) {
            if (strstr($action, '::')) {
                list($controller, $method) = explode('::', $action);
            } elseif (strstr($action, '@')) {
                list($controller, $method) = explode('@', $action);
            } else {
                list($controller, $method) = [false, false];
                logger()->error($errMsg);
                std_out_log()->error($errMsg);
            }
        } else {
            list($controller, $method) = [false, false];
            logger()->error($errMsg);
            std_out_log()->error($errMsg);
        }
        return compact('controller', 'method');
    }
}

```
