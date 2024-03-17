---
title: Swoole 的底层架构及运行原理
author: Alex
top: false
hide: false
cover: false
toc: true
mathjax: false
categories: PHP
tags:
  - PHP
  - Swoole
abbrlink: 84212e0c
date: 2024-03-17 13:43:05
img:
coverImg:
password:
summary:
---

先看这张底层架构图

[//]: # (![Swoole 的底层架构及运行原理图]&#40;https://laravel.gstatics.cn/wp-content/uploads/2019/07/c6b0f1f7dadffaa5be4f7b6b869acd67.jpg&#41;)

![Swoole 的底层架构及运行原理图](https://upload-images.jianshu.io/upload_images/14623749-176c45bba608fa9a.jpg?imageMogr2/auto-orient/strip%7CimageView2/2/w/1240)

我们可以看到，Swoole 主要包含以下组件：

- **Master：** 当我们运行启动 Swoole 的 PHP 脚本时，首先会创建该进程（它是整个应用的 root 进程），然后由该进程 fork 出 Reactor 线程和 Manager 进程。
- **Reactor：** Reactor 是包含在 Master 进程中的多线程程序，用来处理 TCP 连接和数据收发（异步非阻塞方式）。Reactor 主线程在 Accept 新的连接后，会将这个连接分配给一个固定的 Reactor 线程，并由这个线程负责监听此 socket。在 socket 可读时读取数据，并进行协议解析，将请求投递到 Worker 进程；在 socket 可写时将数据发送给 TCP 客户端。
- **Manager：** Manager 进程负责 fork 并维护多个 Worker 子进程。当有 Worker 子进程中止时，Manager 负责回收并创建新的 Worker 子进程，以便保持 Worker 进程总数不变；当服务器关闭时，Manager 将发送信号给所有 Worker 子进程，通知其关闭服务。
- **Worker：** 以多进程方式运行，每个子进程负责接受由 Reactor 线程投递的请求数据包，并执行 PHP 回调函数处理数据，然后生成响应数据并发给 Reactor 线程，由 Reactor 线程发送给 TCP 客户端。所有请求的处理逻辑都是在 Worker 子进程中完成，这是我们编写业务代码时真正要关心的部分。
- **Task Worker：** 功能和 Worker 进程类似，同样以多进程方式运行，但仅用于任务分发，当 Worker 进程将任务异步分发到任务队列时，Task Worker 负责从队列中消费这些任务（同步阻塞方式处理），处理完成后将结果返回给 Worker 进程。


Swoole 官方对 Reactor、Worker、Task Worker有一个形象的比喻，如果把基于 Swoole 的 Web 服务器比作一个工厂，那么 Reactor 就是这个工厂的销售员，Worker 是负责生产的工人，销售员负责接订单，然后交给工人生产，而 Task Worker 可以理解为行政人员，负责替工人处理生产以外的杂事，比如订盒饭、收快递，让工人可以安心生产。

# Swoole 的生命周期回调函数

当 Master 主进程启动或关闭时会触发下面这两个回调函数：

- onStart
- onShutdown

而 Manager 管理进程启动或关闭时会触发下面这两个回调函数：

- onManagerStart
- onManagerStop

Worker 进程的生命周期中，有多个回调函数：

- onWorkerStart：Worker 进程启动时
- onWorkerStop： Worker 进程关闭时
- onConnect：连接建立时
- onClose：连接关闭时
- onReceive：收到请求数据时
- onFinish：投递的任务处理完成时

Task Worker 进程也有两个回调函数，分别在

- onTask：由新任务投递过来时
- onWorkerStart：Task Worker 进程启动时也会触发

我们日常开发中主要关注的是 Worker 进程的回调函数，只需要在服务器实例上监听相应的事件，并编写对应的回调函数来处理相应的业务逻辑即可。
