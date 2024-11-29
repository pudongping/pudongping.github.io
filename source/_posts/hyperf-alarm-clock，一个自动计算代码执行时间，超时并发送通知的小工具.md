---
title: hyperf-alarm-clock，一个自动计算代码执行时间，超时并发送通知的小工具
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
abbrlink: 982b031b
date: 2024-11-29 10:07:09
img:
coverImg:
password:
summary:
---

在软件开发中，我们经常需要对代码执行时间进行监控，以确保系统的性能和稳定性。在 PHP 的世界里，Hyperf 框架以其高性能和丰富的组件生态而闻名，而今天我要介绍的是 Hyperf 生态中的一个小巧的插件包 —— `hyperf-alarm-clock` 库。它是一个计算代码执行时间，并在超时后发送通知的小工具。

本文将带你了解这个库的安装、特性、基本与高级功能，并结合实际应用场景，展示其在项目中的应用。

## hyperf-alarm-clock 库简介

`hyperf-alarm-clock` 是一个为 Hyperf 框架设计的库，它可以帮助开发者监控代码的执行时间，并在**代码执行时间超过预设阈值时发送通知**。这对于需要长时间运行的任务或者需要确保及时响应的系统尤为重要。

## 安装

要在你的 Hyperf 项目中使用 `hyperf-alarm-clock`，你需要通过 Composer 进行安装：

> 这个库最近已经支持了 `hyperf 3.1` 版本。

```shell
composer require pudongping/hyperf-alarm-clock:^3.0 -vvv
```

确保你的环境满足以下要求：
- PHP >= 8.1
- Composer
- Hyperf ~3.1.0

## 配置

安装完成后，你需要发布配置文件并进行相应的配置。在项目根目录下执行以下命令，生成配置文件：

```shell
php bin/hyperf.php vendor:publish pudongping/hyperf-alarm-clock
```

这将在`config/autoload`目录下生成`hyperf_alarm_clock.php`配置文件。接下来，你可以根据需要修改该配置文件。

大致配置项如下：

```php
return [

    /*
    |--------------------------------------------------------------------------
    | 是否开启代码执行耗时，超时后通知提醒
    |--------------------------------------------------------------------------
    |
    | 只有当此参数设定为 `true` ，并且指定的通知通道所需要的必要参数都存在时，才会触发通知提醒。
    | 比如，假设设定了通知通道包含 `feishu` 因为发送飞书 webhook 通知必须要有密钥，因此只有当
    | ALARM_CLOCK_ENABLE 为 true 并且 ALARM_CLOCK_CHANNEL_FEISHU_WEBHOOK_SECRET 设定
    | 有密钥时，才会发送通知提醒
    |
    */
    'enable' => env('ALARM_CLOCK_ENABLE', false),

    /*
    |--------------------------------------------------------------------------
    | 通知标题
    |--------------------------------------------------------------------------
    |
    */
    'title' => env('ALARM_CLOCK_TITLE', 'alarm-clock'),

    /*
    |--------------------------------------------------------------------------
    | 默认的通知通道
    |--------------------------------------------------------------------------
    |
    | 默认设定为 **标准输出 stdout** 想要同时设定多个通知通道，请使用英文逗号 `,` 进行分隔。比如
    | `feishu,logging`
    |
    */
    'default' => env('ALARM_CLOCK_CHANNEL', 'stdout'),

    /*
    |--------------------------------------------------------------------------
    | 通知通道设置
    |--------------------------------------------------------------------------
    |
    | 键名为通知通道的名称，值为通知通道的必要参数
    | 不同的通道可能所需要的必要参数不一样
    |
    */
    'channels' => [
        'feishu' => [
            'webhook_secret' => env('ALARM_CLOCK_CHANNEL_FEISHU_WEBHOOK_SECRET', ''),
        ],
        'logging' => [
            'group' => env('ALARM_CLOCK_CHANNEL_LOGGING_GROUP', 'default'),
            'name' => env('ALARM_CLOCK_CHANNEL_LOGGING_NAME', 'hyperf'),
        ],
        'stdout',
    ],

    /*
    |--------------------------------------------------------------------------
    | 超时时间及通知类型设置，超时时间支持浮点数，比如：5.5 表示 5s 500ms
    |--------------------------------------------------------------------------
    |
    | 以下默认设置表示为：
    | 当代码执行时间超过 10s 时，以 `warning` 的形式进行通知
    | 当代码执行时间超过 5s 且不足以 10s 时，以 `notice` 的形式进行通知
    | 当代码执行时间不超过 5s 时，不通知。
    |
    */
    'timeout' => [
        'notice' => env('ALARM_CLOCK_TIMEOUT_NOTICE', 5),
        'warning' => env('ALARM_CLOCK_TIMEOUT_WARNING', 10),
    ],

];
```

## 使用

要使用 `hyperf-alarm-clock`，你需要将 `\Pudongping\HyperfAlarmClock\AlarmClockMiddleware` 中间件添加到 `config/autoload/middlewares.php` 配置文件中，即可。

比如：

```php
<?php

declare(strict_types=1);

return [
    'http' => [
        \Pudongping\HyperfAlarmClock\AlarmClockMiddleware::class,
    ],
];

```

## 基本功能

`hyperf-alarm-clock` 提供了多种通知通道，包括标准输出、日志文件和飞书等。你可以根据项目需求选择合适的通知方式。

## 具体使用

设定好了中间件之后，就不需要做任何改动了。正常写自己的业务代码，如果自己的接口执行时间超过了在配置文件中定义的时间阀值之后，就会根据设定的通知通道进行发送通知，比如，我们随便在任意一个接口中睡眠了几秒，以模拟代码耗时执行。

如果使用**标准输出**作为通知通道时，会出现类似以下内容

![stdout.png](https://upload-images.jianshu.io/upload_images/14623749-f7f185dfc8df26a7.png?imageMogr2/auto-orient/strip%7CimageView2/2/w/1240)

如果使用**日志文件**作为通知通道时，会出现类似以下内容

![logging.png](https://upload-images.jianshu.io/upload_images/14623749-75c95f714f8e7fbf.png?imageMogr2/auto-orient/strip%7CimageView2/2/w/1240)

如果使用**飞书**作为通知通道时，就会出现以下类似内容

![feishu.png](https://upload-images.jianshu.io/upload_images/14623749-a1dbb3040791c44d.png?imageMogr2/auto-orient/strip%7CimageView2/2/w/1240)


## 实际应用场景

假设你正在开发一个电商平台，需要监控订单处理流程的执行时间。如果订单处理时间超过预设阈值，`hyperf-alarm-clock` 可以及时通知开发团队，以便快速定位问题并采取措施。

## 结语

`hyperf-alarm-clock` 是一个强大的工具，它可以帮助 Hyperf 框架的开发者监控代码的执行时间，并在必要时发送通知。

通过本文的介绍，希望你能对 `hyperf-alarm-clock` 有一个全面的了解，并在你的项目中灵活运用。
